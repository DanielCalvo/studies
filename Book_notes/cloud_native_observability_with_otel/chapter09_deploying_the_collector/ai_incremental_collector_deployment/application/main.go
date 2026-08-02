package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
)

func main() {
	if err := run(context.Background()); err != nil {
		slog.Error("greeting application stopped", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context) error {
	tracerProvider, err := newTracerProvider(ctx)
	if err != nil {
		return err
	}
	defer tracerProvider.Shutdown(ctx)

	mux := http.NewServeMux()
	mux.HandleFunc("/", greet)

	address := ":8080"

	// The application still writes an ordinary log to its own standard output.
	// Adding tracing does not automatically turn slog records into OTel logs.
	slog.Info("greeting application started", "address", address)

	// otelhttp creates a server span for each request. The span is exported to
	// the Collector selected through the OTLP exporter configuration.
	instrumentedHandler := otelhttp.NewHandler(mux, "greeting-http-server")

	return http.ListenAndServe(address, instrumentedHandler)
}

func newTracerProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	// With no endpoint options in code, the exporter reads the standard
	// OTEL_EXPORTER_OTLP_ENDPOINT environment variable. This lets each
	// Kubernetes deployment method select its Collector without rebuilding or
	// changing the application.
	endpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	slog.Info("configuring OTLP trace exporter", "endpoint", endpoint)

	exporter, err := otlptracegrpc.New(ctx)
	if err != nil {
		return nil, fmt.Errorf("create OTLP trace exporter: %w", err)
	}

	res, err := resource.New(
		ctx,
		// WithFromEnv reads standard OTEL_RESOURCE_ATTRIBUTES values. In the
		// agent deployment, Kubernetes supplies the Pod IP as an association
		// hint so the Collector can look up the rest of the workload metadata.
		resource.WithFromEnv(),
		resource.WithAttributes(
			semconv.ServiceName("chapter9-greeting"),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("create OpenTelemetry resource: %w", err)
	}

	provider := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
	)
	otel.SetTracerProvider(provider)

	return provider, nil
}

func greet(writer http.ResponseWriter, request *http.Request) {
	name := request.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}

	// This remains an ordinary structured application log. The HTTP span and
	// this log are separate signals in this step.
	slog.Info("greeting requested", "name", name)

	fmt.Fprintf(writer, "Hello, %s!\n", name)
}
