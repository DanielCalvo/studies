package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

const (
	listenAddress        = ":8080"
	formattingServiceURL = "http://localhost:8081"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "greeting service failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	// This is the familiar tracing pipeline. The instrumentation library will
	// create spans, while this pipeline processes and prints them.
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return fmt.Errorf("create trace exporter: %w", err)
	}
	traceProcessor := sdktrace.NewSimpleSpanProcessor(traceExporter)

	// Default includes the SDK's environment Resource detector. It reads
	// OTEL_SERVICE_NAME and OTEL_RESOURCE_ATTRIBUTES at process startup, so
	// service identity can be configured without application-specific code.
	serviceResource := resource.Default()

	tracerProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSpanProcessor(traceProcessor),
		sdktrace.WithResource(serviceResource),
	)
	otel.SetTracerProvider(tracerProvider)
	defer func() {
		if err := tracerProvider.Shutdown(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "shut down tracer provider: %v\n", err)
		}
	}()

	// TraceContext implements the W3C traceparent header format. NewTransport
	// uses this global propagator to inject the active client span's identity
	// into the outbound HTTP request.
	// this must be set before constructing OpenTelemetry instrumentation that reads the global propagator, such as:
	//  otelhttp.NewTransport(...)
	//  otelhttp.NewHandler(...)
	otel.SetTextMapPropagator(propagation.TraceContext{})

	// The instrumentation library already knows which HTTP measurements to
	// record. This pipeline makes those automatic measurements observable.
	//generates metrics for all our http endpoints, neat
	metricExporter, err := stdoutmetric.New(stdoutmetric.WithPrettyPrint())
	if err != nil {
		return fmt.Errorf("create metric exporter: %w", err)
	}
	metricReader := sdkmetric.NewPeriodicReader(
		metricExporter,
		sdkmetric.WithInterval(5*time.Second),
	)
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(metricReader),
		sdkmetric.WithResource(serviceResource),
	)
	otel.SetMeterProvider(meterProvider)
	defer func() {
		if err := meterProvider.Shutdown(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "shut down meter provider: %v\n", err)
		}
	}()

	// NewTransport wraps the component that performs each outbound HTTP
	// exchange. It automatically creates a client span and records client
	// metrics around client.Do.
	client := &http.Client{
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", func(w http.ResponseWriter, r *http.Request) {
		greet(w, r, client)
	})
	mux.HandleFunc("GET /health", health)

	// NewHandler is the instrumentation-library activation point. It wraps the
	// ordinary mux and automatically creates server spans and metrics.
	// WithFilter allows /health to execute normally while returning false
	// suppresses all of this wrapper's telemetry for that request.
	//So this excludes the health check and the point from having tracing enabled on it
	instrumentedHandler := otelhttp.NewHandler(
		mux,
		"greeting-server",
		otelhttp.WithFilter(func(r *http.Request) bool {
			return r.URL.Path != "/health"
		}),
		// The callback customizes only the automatically generated server span
		// name. It uses the route pattern rather than request-specific values.
		otelhttp.WithSpanNameFormatter(greetingSpanName),
	)

	fmt.Printf("Greeting service listening on http://localhost%s\n", listenAddress)
	return http.ListenAndServe(listenAddress, instrumentedHandler)
}

// greetingSpanName demonstrates a customization callback exposed by the
// instrumentation library. r.Pattern is stable across names such as Daniel and
// Alice. Before routing resolves a pattern, use the fixed operation fallback.
func greetingSpanName(operation string, r *http.Request) string {
	if r.Pattern == "" {
		return operation
	}
	// Our registered pattern is "GET /hello/{name}", so it already contains
	// the HTTP method and does not need r.Method prepended separately.
	return "greeting " + r.Pattern
}

// greet receives the public request and asks the formatting service to produce
// the response. The client's instrumented Transport observes the outbound
// request without requiring tracing calls inside this function.
func greet(w http.ResponseWriter, r *http.Request, client *http.Client) {
	name := r.PathValue("name")
	endpoint := formattingServiceURL + "/format/" + url.PathEscape(name)

	// r.Context() contains the server span created by NewHandler. Passing it to
	// the request lets NewTransport make its client span a child of that span.
	request, err := http.NewRequestWithContext(r.Context(), http.MethodGet, endpoint, nil)
	if err != nil {
		http.Error(w, "could not create formatting request", http.StatusInternalServerError)
		return
	}

	response, err := client.Do(request)
	if err != nil {
		http.Error(w, "formatting service unavailable", http.StatusBadGateway)
		return
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		http.Error(w, "formatting service failed", http.StatusBadGateway)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = io.Copy(w, response.Body)
}

// health provides a lightweight liveness response. The route belongs to the
// application as usual; only the surrounding OTel server span is filtered out.
func health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = fmt.Fprintln(w, "OK")
}
