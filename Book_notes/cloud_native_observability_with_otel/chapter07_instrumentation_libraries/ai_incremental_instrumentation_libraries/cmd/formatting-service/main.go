package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
)

const (
	listenAddress                 = ":8081"
	formattingInstrumentationName = "ai-incremental-instrumentation-libraries/formatting-service"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "formatting service failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	// This service needs its own telemetry pipeline because it is a separate
	// process from the greeting service.
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return fmt.Errorf("create trace exporter: %w", err)
	}
	traceProcessor := sdktrace.NewSimpleSpanProcessor(traceExporter)

	serviceResource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("formatting-service"),
			semconv.ServiceVersion("1.0.0"),
		),
	)
	if err != nil {
		return fmt.Errorf("create service resource: %w", err)
	}

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

	// On the receiving side, TraceContext parses the traceparent header and
	// reconstructs the remote span context for NewHandler.
	otel.SetTextMapPropagator(propagation.TraceContext{})

	mux := http.NewServeMux()
	mux.HandleFunc("GET /format/{name}", formatGreeting)

	// The wrapper extracts the incoming trace context, creates a server span
	// beneath the remote client span, and places the new span in r.Context().
	instrumentedHandler := otelhttp.NewHandler(mux, "formatting-server")

	fmt.Printf("Formatting service listening on http://localhost%s\n", listenAddress)
	return http.ListenAndServe(listenAddress, instrumentedHandler)
}

// formatGreeting performs the small piece of work owned by the formatting
// service: it turns a name into a greeting and returns it to its HTTP caller.
func formatGreeting(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	greeting := buildGreeting(r.Context(), name)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = fmt.Fprintln(w, greeting)
}

// buildGreeting is application-specific work that an HTTP instrumentation
// library cannot identify automatically. Starting one manual span gives that
// work its own timing beneath the automatic formatting server span.
func buildGreeting(ctx context.Context, name string) string {
	// The TracerProvider is already configured globally in run. Passing the
	// handler Context makes this span a child of the active server span.
	_, span := otel.Tracer(formattingInstrumentationName).Start(ctx, "build greeting")
	defer span.End()

	return fmt.Sprintf("Hello, %s!", name)
}
