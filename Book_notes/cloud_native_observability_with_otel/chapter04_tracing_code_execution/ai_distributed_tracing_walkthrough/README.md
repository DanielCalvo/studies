# Chapter 4 distributed tracing walkthrough

This program connects the main concepts from Chapter 4 in one executable Go
example. It simulates several services in the same process so that it can be run
without Docker, an OpenTelemetry Collector, or a tracing backend.

The example demonstrates:

- API and SDK responsibilities
- `TracerProvider`, `Tracer`, `Resource`, `SpanProcessor`, and `SpanExporter`
- Instrumentation scopes
- Parent and child spans
- `Context` and `SpanContext`
- Resource attributes and span attributes
- Semantic conventions
- All five span kinds
- Events, exceptions, and span status
- W3C Trace Context injection and extraction
- Correct and deliberately broken context propagation
- An asynchronous producer and consumer
- A span link between two separate traces

Run it with:

```bash
go run .
```

The custom teaching exporter prints a compact description of every completed
span. Compare the trace and parent IDs in these operations:

- `HTTP GET /products (propagated)` and the corresponding server span share a
  trace ID. The server's parent is the client span.
- `HTTP GET /products (headers omitted)` and the corresponding server span have
  different trace IDs. This is what broken propagation looks like.
- `restock requested` and `restock request processed` share a trace even though
  the producer ends before the consumer begins.
- `write audit record` starts a new trace but contains a link to `shopper
  session`.

The exporter uses `BatchSpanProcessor`, so spans are queued and exported away
from the application path. The program explicitly flushes each provider near
the end so that the output is complete before the process exits.
