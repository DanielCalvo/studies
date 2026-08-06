package main

import (
	"bytes"
	"context"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/sdk/trace/tracetest"
)

func TestSuccessfulResizeCreatesProcessingSpansAndCorrelatedLog(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exporter))
	t.Cleanup(func() {
		_ = provider.Shutdown(context.Background())
	})

	var logs bytes.Buffer
	logger := slog.New(slog.NewJSONHandler(&logs, nil))
	server := newHandlerWithTelemetry(
		defaultConfig(),
		logger,
		provider,
		propagation.TraceContext{},
	)

	healthResponse := httptest.NewRecorder()
	server.ServeHTTP(healthResponse, httptest.NewRequest(http.MethodGet, "/livez", nil))
	if spans := exporter.GetSpans(); len(spans) != 0 {
		t.Fatalf("health request created %d spans, want 0", len(spans))
	}

	request := multipartRequest(t, http.MethodPost, "/v1/resize?width=3", "image", makeJPEG(t, 6, 4))
	request.Header.Set("traceparent", "00-4bf92f3577b34da6a3ce929d0e0e4736-00f067aa0ba902b7-01")
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d; body: %s", response.Code, http.StatusOK, response.Body.String())
	}

	spans := exporter.GetSpans()
	if len(spans) != 5 {
		t.Fatalf("exported spans = %d, want 5", len(spans))
	}
	serverSpan := findSpan(t, spans, "POST /v1/resize")
	if got := serverSpan.SpanContext.TraceID().String(); got != "4bf92f3577b34da6a3ce929d0e0e4736" {
		t.Fatalf("server trace ID = %q, want propagated trace ID", got)
	}
	if !serverSpan.Parent.IsRemote() {
		t.Fatal("server span parent is not marked remote")
	}
	if got := serverSpan.ChildSpanCount; got != 4 {
		t.Fatalf("server child span count = %d, want 4", got)
	}

	for _, name := range []string{"image.read_upload", "image.decode", "image.resize", "image.encode"} {
		child := findSpan(t, spans, name)
		if child.Parent.SpanID() != serverSpan.SpanContext.SpanID() {
			t.Fatalf("%s parent = %s, want server span %s", name, child.Parent.SpanID(), serverSpan.SpanContext.SpanID())
		}
	}

	assertSpanAttribute(t, serverSpan.Attributes, "image.outcome", "succeeded")
	assertSpanAttribute(t, serverSpan.Attributes, "image.input_width", int64(6))
	assertSpanAttribute(t, serverSpan.Attributes, "image.input_height", int64(4))
	assertSpanAttribute(t, serverSpan.Attributes, "image.target_width", int64(3))
	assertSpanAttribute(t, serverSpan.Attributes, "image.output_height", int64(2))

	entry := decodeLogEntry(t, &logs)
	assertLogValue(t, entry, "trace_id", serverSpan.SpanContext.TraceID().String())
	assertLogValue(t, entry, "span_id", serverSpan.SpanContext.SpanID().String())
	assertLogValue(t, entry, "trace_sampled", true)
}

func TestOpenTelemetryResourceMergesEnvironmentAttributes(t *testing.T) {
	t.Setenv("OTEL_RESOURCE_ATTRIBUTES", "service.namespace=image-resizer,deployment.environment.name=opi")

	res, err := newOpenTelemetryResource(context.Background())
	if err != nil {
		t.Fatalf("newOpenTelemetryResource returned an error: %v", err)
	}
	assertSpanAttribute(t, res.Attributes(), "service.name", serviceName)
	assertSpanAttribute(t, res.Attributes(), "service.namespace", "image-resizer")
	assertSpanAttribute(t, res.Attributes(), "deployment.environment.name", "opi")
}

func TestRejectedResizeAnnotatesServerSpan(t *testing.T) {
	exporter := tracetest.NewInMemoryExporter()
	provider := sdktrace.NewTracerProvider(sdktrace.WithSyncer(exporter))
	t.Cleanup(func() {
		_ = provider.Shutdown(context.Background())
	})

	server := newHandlerWithTelemetry(
		defaultConfig(),
		discardLogger(),
		provider,
		propagation.TraceContext{},
	)
	request := multipartRequest(t, http.MethodPost, "/v1/resize?width=0", "image", makeJPEG(t, 4, 2))
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	if response.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusBadRequest)
	}

	spans := exporter.GetSpans()
	if len(spans) != 1 {
		t.Fatalf("exported spans = %d, want 1", len(spans))
	}
	serverSpan := findSpan(t, spans, "POST /v1/resize")
	assertSpanAttribute(t, serverSpan.Attributes, "image.outcome", "rejected")
	assertSpanAttribute(t, serverSpan.Attributes, "image.error_code", "invalid_width")
}

func findSpan(t *testing.T, spans tracetest.SpanStubs, name string) tracetest.SpanStub {
	t.Helper()
	for _, span := range spans {
		if span.Name == name {
			return span
		}
	}
	t.Fatalf("span %q not found", name)
	return tracetest.SpanStub{}
}

func assertSpanAttribute(t *testing.T, attributes []attribute.KeyValue, key string, want any) {
	t.Helper()
	for _, item := range attributes {
		if string(item.Key) == key {
			if got := item.Value.AsInterface(); got != want {
				t.Fatalf("span attribute %s = %#v, want %#v", key, got, want)
			}
			return
		}
	}
	t.Fatalf("span attribute %s not found", key)
}
