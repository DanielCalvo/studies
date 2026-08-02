// Program ai-distributed-tracing-walkthrough connects the Chapter 4 tracing
// concepts in one executable example.
//
// The program simulates several services in one process. Each simulated service
// has its own TracerProvider and Resource, just as independently deployed
// services would. Requests still cross explicit HTTP-header and message-header
// carriers, so context propagation does not happen accidentally through normal
// Go function calls.
package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"sort"
	"sync"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
	"go.opentelemetry.io/otel/trace"
)

const instrumentationVersion = "1.0.0"

// service contains the telemetry objects owned by one simulated service.
//
// A real executable normally has one TracerProvider. This teaching program has
// several because shopper, grocery-store, and inventory-worker are pretending
// to be independently deployed services.
type service struct {
	name     string
	provider *sdktrace.TracerProvider
	tracer   trace.Tracer
}

// message is a small stand-in for a message sent through a queue.
//
// The body is application data. The headers are the carrier into which the
// producer injects tracing context and from which the consumer extracts it.
type message struct {
	body    string
	headers propagation.MapCarrier
}

func main() {
	if err := run(context.Background()); err != nil {
		fmt.Fprintf(os.Stderr, "walkthrough failed: %v\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	// ---------------------------------------------------------------------
	// SECTION 1: Configure propagation.
	//
	// TraceContext implements the W3C Trace Context format. It serializes a
	// SpanContext into headers such as "traceparent".
	//
	// Baggage is a different cross-cutting concern, but combining it with
	// TraceContext demonstrates a composite propagator. In a migration, a
	// composite could similarly contain W3C and a legacy trace format such as
	// B3. A sender and receiver must understand compatible formats.
	// ---------------------------------------------------------------------
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	// One exporter receives completed spans from every simulated service. Its
	// compact output makes the important IDs and relationships easy to inspect.
	exporter := &teachingExporter{}

	shopper, err := newService(ctx, "shopper", exporter)
	if err != nil {
		return err
	}
	groceryStore, err := newService(ctx, "grocery-store", exporter)
	if err != nil {
		return err
	}
	inventoryWorker, err := newService(ctx, "inventory-worker", exporter)
	if err != nil {
		return err
	}
	services := []*service{shopper, groceryStore, inventoryWorker}

	// The OpenTelemetry API lets application code ask for tracers and create
	// spans. The SDK supplies the concrete provider, processors, and exporters.
	// Setting a global provider connects API calls to that SDK implementation.
	// Without this configuration, the API safely uses a no-op provider.
	//
	// The other simulated services use their provider fields directly because
	// one process can have only one global provider.
	otel.SetTracerProvider(shopper.provider)

	groceryHandler := newGroceryHandler(groceryStore)

	// ---------------------------------------------------------------------
	// SECTION 2: Start a root span and carry it in a Go context.Context.
	//
	// tracer.Start returns:
	//   1. a new context containing the new current span; and
	//   2. the span itself, which records the unit of work.
	//
	// Passing sessionCtx to another function makes new spans children of this
	// span. This is in-process context propagation.
	// ---------------------------------------------------------------------
	sessionCtx, sessionSpan := shopper.tracer.Start(
		ctx,
		"shopper session",
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithAttributes(
			attribute.String("shopper.cart.id", "cart-123"),
			attribute.Int("shopper.cart.item_count", 2),
		),
	)
	sessionSpan.AddEvent("shopper started browsing")

	// Save the small, serializable tracing identity before ending the span.
	// SpanContext contains the trace ID, span ID, trace flags, and trace state.
	// It is not the same thing as context.Context, which carries the current
	// span and other request-scoped values through Go code.
	sessionSpanContext := sessionSpan.SpanContext()

	// This request injects SpanContext into HTTP headers. The server extracts
	// it, so its SERVER span joins this trace as a child of the CLIENT span.
	if err := browseStore(sessionCtx, shopper, groceryHandler, true); err != nil {
		recordFailure(sessionSpan, err)
		sessionSpan.End()
		return err
	}

	// This second request intentionally omits injection. Although both sides
	// are instrumented, the server has no incoming parent information and
	// therefore starts a new trace. This demonstrates broken propagation.
	if err := browseStore(sessionCtx, shopper, groceryHandler, false); err != nil {
		recordFailure(sessionSpan, err)
		sessionSpan.End()
		return err
	}

	// This child operation fails, but the enclosing shopper session can still
	// recover and finish successfully. It demonstrates that status belongs to
	// an individual span: one ERROR child does not require every ancestor to
	// have the same status.
	rejectExpiredCoupon(sessionCtx, shopper)
	sessionSpan.AddEvent("expired coupon ignored; checkout continued")

	// ---------------------------------------------------------------------
	// SECTION 3: Demonstrate asynchronous PRODUCER and CONSUMER spans.
	//
	// The producer injects context into message metadata. The consumer can
	// extract it later, after the producer has ended, and still join the trace.
	// ---------------------------------------------------------------------
	restockMessage := produceRestockRequest(sessionCtx, shopper)
	consumeRestockRequest(ctx, inventoryWorker, restockMessage)

	sessionSpan.AddEvent("shopper completed checkout")
	sessionSpan.SetStatus(codes.Ok, "")
	sessionSpan.End()

	// ---------------------------------------------------------------------
	// SECTION 4: Demonstrate a span link.
	//
	// The audit operation deliberately starts a new trace. A link records its
	// relationship with the earlier shopper span without making that span its
	// parent. Links are useful for batches, fan-in, and other work that is
	// related but is not accurately represented by one parent-child tree.
	// ---------------------------------------------------------------------
	writeLinkedAuditRecord(ctx, groceryStore, sessionSpanContext)

	// BatchSpanProcessor exports in the background. ForceFlush asks each
	// provider to export queued spans now so the lesson's output is complete.
	for _, svc := range services {
		if err := svc.provider.ForceFlush(ctx); err != nil {
			return fmt.Errorf("flush %s traces: %w", svc.name, err)
		}
	}
	for _, svc := range services {
		if err := svc.provider.Shutdown(ctx); err != nil {
			return fmt.Errorf("shut down %s tracing: %w", svc.name, err)
		}
	}

	return nil
}

// newService assembles the tracing pipeline for one service:
//
//	TracerProvider + Resource -> Tracer -> SpanProcessor -> SpanExporter
func newService(ctx context.Context, name string, exporter sdktrace.SpanExporter) (*service, error) {
	// A Resource answers "who produced this telemetry?" These attributes
	// describe the service and are attached to all of its spans.
	//
	// WithHost is a ResourceDetector: it discovers host information rather
	// than requiring application code to provide it manually. WithFromEnv
	// also accepts standard OTEL_RESOURCE_ATTRIBUTES configuration.
	res, err := resource.New(
		ctx,
		resource.WithFromEnv(),
		resource.WithHost(),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceName(name),
			semconv.ServiceVersion("1.0.0"),
			semconv.DeploymentEnvironmentNameKey.String("study-lab"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("create %s resource: %w", name, err)
	}

	// BatchSpanProcessor queues completed spans and exports them outside the
	// request path. This avoids paying exporter latency every time Span.End is
	// called. Shutdown or ForceFlush is important so queued spans are not lost.
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithBatcher(
			exporter,
			sdktrace.WithBatchTimeout(200*time.Millisecond),
			sdktrace.WithMaxExportBatchSize(64),
		),
	)

	// The tracer's name and version identify the instrumentation scope: the
	// code that created the spans. This is separate from service.name, which
	// identifies the entity running that code.
	tracer := provider.Tracer(
		"example.org/chapter4/"+name,
		trace.WithInstrumentationVersion(instrumentationVersion),
	)

	return &service{name: name, provider: provider, tracer: tracer}, nil
}

func browseStore(
	ctx context.Context,
	shopper *service,
	groceryHandler http.Handler,
	propagate bool,
) error {
	// This INTERNAL span represents work that does not cross a process
	// boundary. Because ctx contains "shopper session", this span is its child.
	ctx, browseSpan := shopper.tracer.Start(
		ctx,
		"browse products",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer browseSpan.End()

	requestName := "HTTP GET /products (propagated)"
	if !propagate {
		requestName = "HTTP GET /products (headers omitted)"
	}

	// CLIENT says that this span represents the caller's side of a synchronous
	// remote request. SpanKind describes the operation's role, not the type of
	// application containing it.
	clientCtx, clientSpan := shopper.tracer.Start(
		ctx,
		requestName,
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			// Semantic conventions give common concepts consistent names so
			// different services and backends can query them reliably.
			semconv.HTTPRequestMethodKey.String(http.MethodGet),
			semconv.ServerAddress("grocery-store"),
			semconv.URLFull("http://grocery-store/products"),
		),
	)
	defer clientSpan.End()

	// context.Background intentionally simulates a separate server process.
	// The server cannot inherit clientCtx through an ordinary function call;
	// only values explicitly injected into headers can cross this boundary.
	req := httptest.NewRequestWithContext(
		context.Background(),
		http.MethodGet,
		"http://grocery-store/products",
		nil,
	)
	if propagate {
		otel.GetTextMapPropagator().Inject(
			clientCtx,
			propagation.HeaderCarrier(req.Header),
		)
	}

	recorder := httptest.NewRecorder()
	groceryHandler.ServeHTTP(recorder, req)
	response := recorder.Result()
	defer response.Body.Close()

	clientSpan.SetAttributes(
		semconv.HTTPResponseStatusCode(response.StatusCode),
	)
	if response.StatusCode >= http.StatusBadRequest {
		err := fmt.Errorf("grocery store returned %s", response.Status)
		recordFailure(clientSpan, err)
		return err
	}

	clientSpan.AddEvent("product response received")
	clientSpan.SetStatus(codes.Ok, "")
	return nil
}

func newGroceryHandler(groceryStore *service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract reads the W3C traceparent header and returns a context
		// containing the remote SpanContext. If the header is missing or uses
		// an incompatible format, the returned context has no valid parent.
		extractedCtx := otel.GetTextMapPropagator().Extract(
			r.Context(),
			propagation.HeaderCarrier(r.Header),
		)

		// SERVER represents the receiving side of the synchronous request. With
		// successful extraction it becomes the child of the remote CLIENT span.
		ctx, serverSpan := groceryStore.tracer.Start(
			extractedCtx,
			"GET /products",
			trace.WithSpanKind(trace.SpanKindServer),
			trace.WithAttributes(
				semconv.HTTPRequestMethodKey.String(r.Method),
				semconv.HTTPRoute("/products"),
				semconv.ServerAddress(r.Host),
			),
		)
		defer serverSpan.End()

		lookUpProducts(ctx, groceryStore)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`[{"name":"orange","quantity":10}]`))

		serverSpan.SetAttributes(
			semconv.HTTPResponseStatusCode(http.StatusOK),
		)
		serverSpan.SetStatus(codes.Ok, "")
	})
}

func lookUpProducts(ctx context.Context, groceryStore *service) {
	_, span := groceryStore.tracer.Start(
		ctx,
		"look up products",
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithAttributes(
			attribute.String("inventory.lookup.strategy", "cache-then-database"),
		),
	)
	defer span.End()

	// An event is a timestamped occurrence within a span. Attributes describe
	// the operation; events describe things that happened during it.
	span.AddEvent("inventory cache lookup started")

	// RecordError adds a structured exception event. The simulated cache error
	// does not escape: the database fallback succeeds. An exception therefore
	// does not automatically mean the overall operation failed.
	cacheErr := errors.New("inventory cache temporarily unavailable")
	span.RecordError(cacheErr)
	span.AddEvent(
		"database fallback succeeded",
		trace.WithAttributes(attribute.Int("inventory.product_count", 1)),
	)

	// Setting OK is intentional: the operation ultimately succeeded. Other
	// internal spans in this program retain the default UNSET status, which
	// means "no explicit status" rather than "failure".
	span.SetStatus(codes.Ok, "")
}

func rejectExpiredCoupon(ctx context.Context, shopper *service) {
	_, span := shopper.tracer.Start(
		ctx,
		"apply coupon",
		trace.WithSpanKind(trace.SpanKindInternal),
		trace.WithAttributes(attribute.String("coupon.code", "EXPIRED-DEMO")),
	)
	defer span.End()

	err := errors.New("coupon has expired")
	recordFailure(span, err)
}

func produceRestockRequest(ctx context.Context, producer *service) message {
	// PRODUCER identifies the operation that initiates asynchronous work. It
	// does not wait for the consumer to process the message.
	producerCtx, span := producer.tracer.Start(
		ctx,
		"restock requested",
		trace.WithSpanKind(trace.SpanKindProducer),
		trace.WithAttributes(
			semconv.MessagingSystemKey.String("in-memory-queue"),
			semconv.MessagingDestinationName("restock-requests"),
		),
	)

	headers := propagation.MapCarrier{}
	otel.GetTextMapPropagator().Inject(producerCtx, headers)
	span.AddEvent("message published")
	span.End()

	return message{
		body:    `{"product":"orange","quantity":20}`,
		headers: headers,
	}
}

func consumeRestockRequest(ctx context.Context, consumer *service, msg message) {
	// Extraction recovers the producer's SpanContext from the message carrier.
	// The consumer can join the same trace even though the producer span has
	// already ended.
	extractedCtx := otel.GetTextMapPropagator().Extract(ctx, msg.headers)
	_, span := consumer.tracer.Start(
		extractedCtx,
		"restock request processed",
		trace.WithSpanKind(trace.SpanKindConsumer),
		trace.WithAttributes(
			semconv.MessagingSystemKey.String("in-memory-queue"),
			semconv.MessagingDestinationName("restock-requests"),
			attribute.Int("messaging.message.body_size", len(msg.body)),
		),
	)
	defer span.End()

	span.AddEvent("inventory replenishment scheduled")
	span.SetStatus(codes.Ok, "")
}

func writeLinkedAuditRecord(
	ctx context.Context,
	auditService *service,
	relatedSpan trace.SpanContext,
) {
	_, span := auditService.tracer.Start(
		ctx,
		"write audit record",
		// WithNewRoot deliberately creates a different trace ID.
		trace.WithNewRoot(),
		// The link preserves an explicit relationship to the earlier trace
		// without claiming that the earlier span is this operation's parent.
		trace.WithLinks(trace.Link{
			SpanContext: relatedSpan,
			Attributes: []attribute.KeyValue{
				attribute.String("link.reason", "checkout audit"),
			},
		}),
		trace.WithAttributes(attribute.String("audit.type", "checkout")),
	)
	defer span.End()

	span.AddEvent("audit record persisted")
	span.SetStatus(codes.Ok, "")
}

func recordFailure(span trace.Span, err error) {
	// Exceptions are events; status is the interpreted outcome. A real failure
	// often needs both so operators can filter ERROR spans and inspect details.
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
}

// teachingExporter is a compact SpanExporter written specifically for this
// walkthrough. A production application would normally export OTLP data to an
// OpenTelemetry Collector instead.
//
// Exporters receive read-only, completed spans from a SpanProcessor. The mutex
// protects output because every service has its own background batch processor.
type teachingExporter struct {
	mu sync.Mutex
}

func (e *teachingExporter) ExportSpans(_ context.Context, spans []sdktrace.ReadOnlySpan) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	for _, span := range spans {
		fmt.Println("------------------------------------------------------------")
		fmt.Printf("service: %s\n", resourceAttribute(span, semconv.ServiceNameKey))
		fmt.Printf(
			"scope:   %s@%s\n",
			span.InstrumentationScope().Name,
			span.InstrumentationScope().Version,
		)
		fmt.Printf("span:    %s\n", span.Name())
		fmt.Printf("kind:    %s\n", span.SpanKind())
		fmt.Printf("trace:   %s\n", span.SpanContext().TraceID())
		fmt.Printf("span ID: %s\n", span.SpanContext().SpanID())
		fmt.Printf(
			"flags:   %s (sampled=%t)\n",
			span.SpanContext().TraceFlags(),
			span.SpanContext().IsSampled(),
		)
		fmt.Printf("state:   %s\n", span.SpanContext().TraceState())
		if span.Parent().IsValid() {
			fmt.Printf("parent:  %s\n", span.Parent().SpanID())
		} else {
			fmt.Println("parent:  none (root span)")
		}
		fmt.Printf("status:  %s", span.Status().Code)
		if span.Status().Description != "" {
			fmt.Printf(" (%s)", span.Status().Description)
		}
		fmt.Println()
		fmt.Printf("duration: %s\n", span.EndTime().Sub(span.StartTime()).Round(time.Microsecond))

		printAttributes("attributes", span.Attributes())

		for _, event := range span.Events() {
			fmt.Printf("event:   %s\n", event.Name)
			printAttributes("  event attributes", event.Attributes)
		}
		for _, link := range span.Links() {
			fmt.Printf(
				"link:    trace=%s span=%s\n",
				link.SpanContext.TraceID(),
				link.SpanContext.SpanID(),
			)
			printAttributes("  link attributes", link.Attributes)
		}
	}

	return nil
}

func (*teachingExporter) Shutdown(context.Context) error {
	return nil
}

func resourceAttribute(span sdktrace.ReadOnlySpan, key attribute.Key) string {
	for _, attr := range span.Resource().Attributes() {
		if attr.Key == key {
			return attr.Value.AsString()
		}
	}
	return "<not set>"
}

func printAttributes(label string, attrs []attribute.KeyValue) {
	if len(attrs) == 0 {
		return
	}

	// Sorting only makes the teaching output deterministic and easier to read.
	sorted := append([]attribute.KeyValue(nil), attrs...)
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].Key < sorted[j].Key
	})

	fmt.Printf("%s:\n", label)
	for _, attr := range sorted {
		fmt.Printf("  %s=%v\n", attr.Key, attr.Value.AsInterface())
	}
}
