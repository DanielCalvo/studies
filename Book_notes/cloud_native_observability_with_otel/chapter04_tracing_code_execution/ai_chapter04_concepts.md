# Chapter 4: Distributed Tracing – Concepts to Learn

Chapter 4 covers printed pages 75–125 of *Cloud Native Observability with
OpenTelemetry*. A large part of the chapter is a worked Python instrumentation
exercise. The code demonstrates the ideas, but it does not all need to be memorized.

The central lesson is:

> Distributed tracing works by recording units of work as spans, maintaining the
> active span in context, and propagating that context across service boundaries.

Everything else in the chapter supports that model.

## 1. The tracing pipeline

The tracing pipeline can be pictured as:

```text
Application code
      |
      v
Tracer -> completed Span -> SpanProcessor -> SpanExporter -> destination
           ^
           |
     TracerProvider
     + Resource
```

- `TracerProvider` owns the tracing configuration and produces tracers.
- `Tracer` creates spans. It is associated with an instrumentation scope, such as
  a library or module name and version.
- `Resource` describes the entity producing the telemetry, such as its service
  name, service version, host, or deployment environment.
- `SpanProcessor` receives spans and controls how they are processed for export.
- `SpanExporter` sends spans to the console, an OpenTelemetry Collector, or a
  tracing backend.

An important architectural distinction is:

- The OpenTelemetry API is what application and library code calls.
- The SDK provides the working implementation and configuration.
- Without an SDK and provider configuration, API calls can safely behave as
  no-ops.

The separation between these components is important. The exact Python
configuration used in the book does not need to be memorized.

## 2. Span lifecycle and relationships

A span represents one unit of work. It contains information such as:

- A name
- A span ID
- A trace ID
- An optional parent span ID
- Start and end timestamps
- Attributes
- Events
- Status
- Resource information

The first span without a parent is the root span.

Spans are not connected merely because one runs after another. A new span
normally becomes a child of whichever span is current in the active context.

For example:

```text
visit store
└── browse
    └── web request
```

All three spans share a trace ID, but each has its own span ID. A child's parent
ID points to its parent's span ID.

Ending a span:

- Fixes its duration.
- Marks its recording as complete.
- Makes it available for processing and export.

## 3. Context versus SpanContext

This is one of the most important distinctions in the chapter.

`Context` is the general mechanism OpenTelemetry uses to carry state through the
current execution. It can hold the current span and other cross-cutting
information. Within an application, frameworks and language runtimes help move
context through function calls, threads, tasks, and request handlers.

`SpanContext` is the small piece of tracing identity associated with a span:

- Trace ID
- Span ID
- Trace flags
- Trace state

A useful mental model is:

```text
Context
└── current Span
    └── SpanContext
        ├── trace ID
        ├── span ID
        ├── trace flags
        └── trace state
```

The chapter demonstrates manual `attach` and `detach` operations. The important
part is understanding why they exist. Their syntax does not need to be memorized,
and framework instrumentation usually manages them automatically.

## 4. Context propagation

Context propagation turns several independent local traces into one distributed
trace.

At a network boundary:

1. The caller obtains the current context.
2. A propagator serializes the relevant information.
3. The serialized values are injected into a carrier, such as HTTP headers.
4. The receiving service extracts those values.
5. It makes the extracted context current.
6. Its server span becomes a child of the caller's client span.

```text
Service A                          Service B

CLIENT span                       SERVER span
trace: 123                        trace: 123
span: A1        --headers-->      span: B1
                                  parent: A1
```

Without propagation, Service B starts a new root span with a different trace ID.
Both services may be instrumented, but the distributed trace will still be
broken.

This gives an SRE an important diagnostic rule:

> If a trace stops and another trace begins at a service boundary, investigate
> context injection, extraction, and propagation-format compatibility.

## 5. Propagation formats and carriers

A carrier is whatever transports context across a boundary. Examples include:

- HTTP headers
- Message metadata
- RPC metadata
- Queue headers

A propagator understands how to inject context into and extract it from a
carrier.

The chapter discusses:

- W3C Trace Context
- B3
- Jaeger
- `ot-trace`
- Composite propagators

The individual Python classes do not need to be memorized. The important problem
is interoperability: if two services expect incompatible propagation formats,
the trace breaks.

A composite propagator can support multiple formats during migrations or when
integrating legacy services.

## 6. SpanKind

`SpanKind` describes a span's role in a communication relationship:

- `INTERNAL`: work contained inside a process.
- `CLIENT`: sends a synchronous request and waits for a response.
- `SERVER`: handles a synchronous request.
- `PRODUCER`: initiates asynchronous work, such as publishing a message.
- `CONSUMER`: processes asynchronous work.

`SpanKind` does not describe the type of application. A server application can
create `CLIENT`, `SERVER`, `PRODUCER`, `CONSUMER`, and `INTERNAL` spans during
one request.

For a synchronous call, the usual relationship is:

```text
Service A CLIENT span
└── Service B SERVER span
```

For messaging, think in terms of producer and consumer operations rather than
pretending the interaction is synchronous.

## 7. Resource attributes versus span attributes

These describe different things.

Resource attributes answer:

> Who produced this telemetry?

Examples include:

- Service name
- Service version
- Host
- Instance
- Region
- Deployment environment

They generally apply to all telemetry generated by that resource rather than to
one operation.

Span attributes answer:

> What was true about this particular operation?

Examples include:

- HTTP request method
- Route
- Response status code
- Database operation
- Peer service
- Application-specific operation information

A concise rule is:

```text
Resource = identity of the producer
Span attributes = details of the operation
```

Resource detectors automatically discover resource information from the runtime
environment. Their purpose is important, but the custom detector code in the
chapter is not.

## 8. Semantic conventions

Semantic conventions define standard names and meanings for common telemetry
fields. They allow different services, languages, instrumentation libraries,
Collectors, and observability backends to interpret telemetry consistently.

The underlying lesson is:

> Consistent telemetry is queryable telemetry.

If one service records an HTTP method one way and another service uses a
different name, fleet-wide queries and dashboards become unreliable.

When selecting attributes:

- Avoid personally identifiable information, credentials, and secrets.
- Record information useful for filtering, grouping, or diagnosing.
- Do not attach arbitrary data merely because it is available.

The book uses an old OpenTelemetry Python release. Its exact
semantic-convention names and import paths should be treated as historical
examples rather than current implementation guidance.

## 9. Span processors and observability overhead

The chapter compares two processing models:

- Simple processing exports each span synchronously when it ends.
- Batch processing queues spans and exports them asynchronously in groups.

The conceptual trade-off is:

```text
Synchronous export
  simpler and immediate
  but adds export latency to application execution

Batch export
  reduces request-path overhead
  but needs queues, background work, flushing, and shutdown handling
```

For production systems, telemetry must not introduce excessive latency or
resource consumption. This trade-off is more important than the particular
Python constructors shown in the chapter.

## 10. Attributes, events, exceptions, and status

These are related but are not interchangeable.

| Data | What it communicates |
| --- | --- |
| Attribute | A property of the span or operation |
| Event | Something that happened at a particular time during the span |
| Exception | A specially structured event describing an exception |
| Status | The final interpreted outcome of the span |

An event occurs within the duration of a span and has its own timestamp. It is
appropriate for significant occurrences such as a retry, cache miss, or state
transition.

An exception is normally represented as an event containing information such as:

- Exception type
- Message
- Stack trace
- Whether it escaped the span

A crucial distinction is:

> Recording an exception does not necessarily mean that the span failed.

An operation might encounter an exception, retry successfully, and complete
normally. The exception remains useful diagnostic evidence, while the final span
status describes the operation's ultimate outcome.

The statuses covered in the chapter are:

- `UNSET`: no explicit outcome has been assigned.
- `OK`: the operation explicitly succeeded.
- `ERROR`: the operation failed according to the operation's semantics.

`UNSET` should not be treated as synonymous with failure.

## What can be skimmed

The following parts are mainly implementation practice:

- Virtual-environment and package-installation commands
- Repeated console JSON
- Python decorators and context-manager syntax
- Flask request-hook mechanics
- The custom `ResourceDetector` implementation
- Repetitive client/server example code
- Exact import paths and semantic-convention constants
- Detailed composite-propagator configuration

These examples are useful for seeing the concepts take shape, but they should
not be studied as current implementation references.

## A gap in the chapter: span links

The chapter introduction says that it will cover span links, but it does not
meaningfully teach them. It only displays empty `"links": []` fields in span
output.

Links should be investigated separately because they are useful when a strict
parent-child relationship does not model the work well. Examples include:

- Batch processing
- Fan-in
- Processing a message independently from its producer
- Associating a new trace with one or more earlier traces

This omission is particularly relevant to asynchronous systems.

## Mastery questions

The chapter's concepts are understood if these questions can be answered:

1. What turns separate spans into a single trace?
2. How does a new span determine its parent?
3. What is the difference between `Context` and `SpanContext`?
4. What must happen when a request crosses a process boundary?
5. Why can two fully instrumented services still produce disconnected traces?
6. What is the difference between a resource and a span attribute?
7. Why should telemetry use semantic conventions?
8. When should a span be `CLIENT` rather than `INTERNAL`?
9. Why is batching preferable on most production request paths?
10. Why can a span contain an exception without having an error status?
11. How would broken propagation appear in a tracing backend?
12. Which data describes the telemetry-producing service, and which describes
    one operation?

If these can be explained comfortably, the important material has been
extracted from Chapter 4. The remaining pages are primarily implementation
practice.
