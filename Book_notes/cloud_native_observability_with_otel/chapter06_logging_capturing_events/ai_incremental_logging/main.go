package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/log/global"
	sdklog "go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.41.0"
)

const (
	listenAddress          = ":8080"
	instrumentationName    = "ai-incremental-logging"
	instrumentationVersion = "1.0.0"
)

func main() {
	if err := run(); err != nil {
		// We need some way to report a startup failure before introducing a
		// logging library. For now, write the error directly to standard error.
		fmt.Fprintf(os.Stderr, "server failed: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	ctx := context.Background()

	// The exporter determines where completed OpenTelemetry log records go.
	// Pretty-printed JSON makes their fields easy to inspect while learning.
	logExporter, err := stdoutlog.New(stdoutlog.WithPrettyPrint())
	if err != nil {
		return fmt.Errorf("create log exporter: %w", err)
	}

	// The batch processor queues log records and exports them asynchronously.
	// This keeps exporter work, especially future network I/O, away from the
	// request path. LoggerProvider shutdown flushes records still in the queue.
	logProcessor := sdklog.NewBatchProcessor(logExporter)

	// A Resource identifies the entity producing every log record from this
	// provider. Merge our service identity with the default Resource to retain
	// useful OpenTelemetry SDK metadata.
	serviceResource, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("greeting-service"),
			semconv.ServiceVersion("1.0.0"),
		),
	)
	if err != nil {
		return fmt.Errorf("create service resource: %w", err)
	}

	// The LoggerProvider owns the OpenTelemetry logging pipeline. Registering
	// it globally makes it available to log bridges and instrumentation.
	loggerProvider := sdklog.NewLoggerProvider(
		sdklog.WithProcessor(logProcessor),
		sdklog.WithResource(serviceResource),
	)
	global.SetLoggerProvider(loggerProvider)
	defer func() {
		if err := loggerProvider.Shutdown(ctx); err != nil {
			fmt.Fprintf(os.Stderr, "shut down logger provider: %v\n", err)
		}
	}()

	// The bridge is an slog.Handler backed by an OpenTelemetry Logger. Making
	// it slog's default connects existing slog calls to the OTel pipeline.
	slog.SetDefault(
		otelslog.NewLogger(
			instrumentationName,
			otelslog.WithVersion(instrumentationVersion),
			otelslog.WithSchemaURL(semconv.SchemaURL),
		),
	)

	// This familiar tracing pipeline exists only to create and export the span
	// that will be correlated with our log record.
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())
	if err != nil {
		return fmt.Errorf("create trace exporter: %w", err)
	}
	traceProcessor := sdktrace.NewSimpleSpanProcessor(traceExporter)
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

	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", greet)

	// otelhttp creates a server span before calling mux and places that span in
	// each request's Context. Application logs can then use r.Context() for
	// trace correlation without manually starting a span in greet.
	instrumentedHandler := otelhttp.NewHandler(mux, "greeting-server") //oh wow, this returns an http handler!

	// This remains a direct human-readable startup message rather than an OTel
	// log record, so starting the service does not demonstrate correlation.
	fmt.Printf("Greeting service listening on http://localhost%s\n", listenAddress)
	return http.ListenAndServe(listenAddress, instrumentedHandler)
}

// greet produces an application log correlated with the server span created
// by the otelhttp instrumentation around this handler.
func greet(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")

	// This deliberate failure gives us an ERROR log to compare with the normal
	// INFO record. otelhttp has already placed the server span in r.Context().
	if name == "fail" {
		slog.ErrorContext(
			r.Context(),
			"greeting failed",
			slog.String("name", name),
			slog.String("reason", "simulated failure"),
		)
		http.Error(w, "failed to produce greeting", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	_, _ = fmt.Fprintf(w, "Hello, %s!\n", name)

	// The otelslog bridge extracts the server span's TraceID, SpanID, and trace
	// flags from the request Context supplied by otelhttp.
	slog.InfoContext(
		r.Context(),
		"greeting produced",
		slog.String("name", name),
	)
}
