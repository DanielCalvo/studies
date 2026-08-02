# Incremental steps taken

## Step 1: Create the ordinary greeting service

### What changed

Created a small Go HTTP service with one route:

```text
GET /hello/{name}
```

The handler reads the name captured by the route and returns a plain-text greeting:

```go
func greet(w http.ResponseWriter, r *http.Request) {
    name := r.PathValue("name")

    w.Header().Set(
        "Content-Type",
        "text/plain; charset=utf-8",
    )
    _, _ = fmt.Fprintf(w, "Hello, %s!\n", name)
}
```

### Concept demonstrated

This establishes the ordinary application behavior before introducing the technology
being studied. Future log records will describe events in a program whose behavior is
already easy to understand.

### Expected behavior

Start the server:

```bash
go run .
```

Then request a greeting:

```bash
curl http://localhost:8080/hello/Daniel
```

The response is:

```text
Hello, Daniel!
```

### Important distinction

The startup message is written directly with `fmt.Printf`, and a fatal startup error is
written directly to standard error. These are ordinary text output calls, not structured
log records and not OpenTelemetry telemetry.

The greeting is the HTTP response body. Returning it to the caller is application
behavior, not logging.

### Deliberately not introduced yet

There is no standard `log/slog` logger, structured log attribute, severity level,
OpenTelemetry logging API or SDK, LoggerProvider, processor, exporter, Resource, trace,
span, or trace/log correlation.

## Step 2: Emit one standard-library log record

### What changed

Imported Go's standard `log/slog` package and emitted one INFO record after producing a
greeting:

```go
slog.Info("greeting produced")
```

### Concept demonstrated

A basic log record describes an event that occurred at a point in time. Go's default
`slog` logger automatically supplies:

- A timestamp
- The `INFO` severity level
- The message body, `greeting produced`

The application supplies the event message but does not need to construct those standard
fields manually.

### Expected behavior

The HTTP behavior remains unchanged:

```text
GET /hello/Daniel -> Hello, Daniel!
```

The server terminal now also shows a record resembling:

```text
time=2026-07-28T12:00:00.000+02:00 level=INFO msg="greeting produced"
```

The exact timestamp is determined at the moment the log call executes.

### Important distinction

The greeting response and the log record have different audiences:

- `Hello, Daniel!` is application data returned to the HTTP client.
- `greeting produced` is operational telemetry written by the server.

This is a standard Go log record, not yet an OpenTelemetry log record.

### Deliberately not introduced yet

The record has no structured attributes and does not receive the HTTP request context.
There is no custom `slog.Handler`, OpenTelemetry LoggerProvider, processor, exporter,
Resource, trace, span, or trace/log correlation.

## Step 3: Add a structured log attribute

### What changed

Added the requested name to the existing log record as a typed `slog` attribute:

```go
slog.Info(
    "greeting produced",
    slog.String("name", name),
)
```

### Concept demonstrated

Structured logging stores contextual information as named fields instead of requiring a
backend to extract it from prose. The log event still has the stable message
`greeting produced`, while `name` carries the value that varies between requests.

This gives a logging backend a field it can filter, group, or display directly:

```text
name = "Daniel"
```

### Expected behavior

Requesting:

```bash
curl http://localhost:8080/hello/Daniel
```

still returns:

```text
Hello, Daniel!
```

The standard logger emits a record resembling:

```text
time=... level=INFO msg="greeting produced" name=Daniel
```

### Important distinction

These two approaches may look similar to a human but produce different data:

```go
slog.Info("greeting produced for Daniel")
slog.Info("greeting produced", slog.String("name", "Daniel"))
```

The first stores everything in an unstructured message body. The second keeps the event
description stable and stores `name` as a separately addressable attribute.

### Deliberately not introduced yet

This is still Go's default `slog` output. The log call does not receive the request
context, and there is no custom Handler, OpenTelemetry LoggerProvider, processor,
exporter, Resource, trace, span, or trace/log correlation.

## Step 4: Pass request context to the logger

### What changed

Changed the log call from `slog.Info` to `slog.InfoContext` and passed the HTTP request
context:

```go
slog.InfoContext(
    r.Context(),
    "greeting produced",
    slog.String("name", name),
)
```

### Concept demonstrated

A `context.Context` can carry request-scoped information through a Go application.
Passing it to the logger makes that information available to the `slog.Handler` that
eventually processes the record.

Later, tracing instrumentation can store the active span in this same request context.
An OpenTelemetry-aware logging Handler can then obtain the corresponding trace and span
identifiers while processing the log record.

### Expected behavior

The greeting response and default log output remain essentially unchanged:

```text
Hello, Daniel!
time=... level=INFO msg="greeting produced" name=Daniel
```

The change prepares the call for request-scoped enrichment; it does not yet add a new
visible field.

### Important distinction

Passing `r.Context()` does not turn the complete Context into attributes and does not
automatically produce a trace ID. It only gives the Handler access to the Context.

The default `slog` Handler does not know about OpenTelemetry spans. Trace correlation
will require both an active span in the Context and an OpenTelemetry-aware Handler.

As a convenience, `slog.Info` invokes the default logger using a background Context.
`slog.InfoContext` lets the application choose which Context is passed to the Handler.
This is why both calls currently print the same fields while only `InfoContext` preserves
the request-scoped information that a future OpenTelemetry Handler can inspect.

### Deliberately not introduced yet

The request context does not currently contain an application span. There is no custom
Handler, OpenTelemetry LoggerProvider, processor, exporter, Resource, trace, span, or
trace/log correlation.

## Step 5: Configure the OpenTelemetry logging pipeline

### What changed

Added an OpenTelemetry stdout log exporter, simple processor, and LoggerProvider:

```go
exporter, err := stdoutlog.New(
    stdoutlog.WithPrettyPrint(),
)

processor := sdklog.NewSimpleProcessor(exporter)

provider := sdklog.NewLoggerProvider(
    sdklog.WithProcessor(processor),
)
global.SetLoggerProvider(provider)
defer provider.Shutdown(context.Background())
```

The real code checks errors from both exporter creation and provider shutdown.

### Concept demonstrated

These components form the basic OpenTelemetry logging pipeline:

```text
LoggerProvider -> LogRecordProcessor -> LogExporter
```

- The LoggerProvider owns the pipeline and supplies OTel Loggers.
- The processor receives emitted OTel log records.
- The stdout exporter encodes those records as readable JSON.

The simple processor exports each record immediately. This makes the data flow easy to
observe before considering asynchronous batching.

### Expected behavior

The application still emits only the standard `slog` text record:

```text
time=... level=INFO msg="greeting produced" name=Daniel
```

No pretty-printed OpenTelemetry JSON record appears yet. Configuring an OTel pipeline
does not automatically intercept calls made through Go's standard logger.

### Important distinction

The program currently contains two disconnected paths:

```text
slog.InfoContext -> default slog Handler -> text output

OTel LoggerProvider -> simple processor -> stdout exporter
         ^
         no records enter this path yet
```

Setting the global OTel LoggerProvider makes it discoverable to bridges and
instrumentation, but it does not replace `slog`'s default Handler.

Provider shutdown is still important even with a simple processor because it releases
the pipeline and exporter cleanly. It will become essential when a later processor holds
records in a batch.

### Deliberately not introduced yet

There is no `slog`-to-OpenTelemetry bridge, so no application log reaches this pipeline.
There is also no Resource, active span, trace ID, span ID, batch processor, OTLP export,
or trace/log correlation.

## Step 6: Bridge standard slog records into OpenTelemetry

### What changed

Created a standard `slog.Logger` backed by OpenTelemetry's `otelslog` bridge and made it
the default logger:

```go
const instrumentationName = "ai-incremental-logging"

global.SetLoggerProvider(provider)
slog.SetDefault(
    otelslog.NewLogger(instrumentationName),
)
```

The bridge uses the globally registered OTel LoggerProvider by default. The existing
`slog.InfoContext` call in the handler did not change.

### Concept demonstrated

A log bridge adapts an established logging API to the OpenTelemetry log data model.
`otelslog` implements the standard `slog.Handler` interface and converts each
`slog.Record` into an OTel log record:

```text
slog.InfoContext
    -> otelslog Handler
    -> OTel LoggerProvider
    -> simple processor
    -> stdout exporter
```

The bridge maps the `slog` timestamp, level, message, and attributes to the corresponding
OpenTelemetry timestamp, severity, body, and attributes.

### Expected behavior

A greeting request still returns:

```text
Hello, Daniel!
```

The compact default `slog` line is replaced by one pretty-printed OpenTelemetry JSON
record containing fields resembling:

```text
SeverityText = INFO
Body = greeting produced
Attributes = [{name: Daniel}]
```

The instrumentation scope name is `ai-incremental-logging`, identifying the application
code whose `slog` records the bridge converted.

### Important distinction

The bridge does not require application code to replace `slog.InfoContext` with direct
calls to the OpenTelemetry Logs API. It replaces the Handler behind the familiar
standard-library API.

The log is not intentionally duplicated:

- Before this step, the default `slog` Handler wrote compact text.
- After this step, the OTel Handler sends the record through the OTel pipeline.

Trace and span identifiers will still be invalid or empty because the request context
does not contain an active span yet.

### Deliberately not introduced yet

The LoggerProvider still uses its default Resource. There is no active span, valid trace
ID, valid span ID, batch processor, OTLP exporter, Collector, or trace/log correlation.

## Step 7: Identify the service with a Resource

### What changed

Created a Resource containing an explicit service name and version, merged it with the
default Resource, and attached it to the LoggerProvider:

```go
serviceResource, err := resource.Merge(
    resource.Default(),
    resource.NewWithAttributes(
        semconv.SchemaURL,
        semconv.ServiceName("greeting-service"),
        semconv.ServiceVersion("1.0.0"),
    ),
)

provider := sdklog.NewLoggerProvider(
    sdklog.WithProcessor(processor),
    sdklog.WithResource(serviceResource),
)
```

### Concept demonstrated

A Resource identifies the entity producing telemetry. Because it belongs to the
LoggerProvider, its attributes are associated with every OTel log record emitted through
that provider.

Merging with `resource.Default()` preserves metadata such as the OpenTelemetry SDK
name, language, and version while replacing the default unknown service identity.

### Expected behavior

The JSON Resource section changes from an automatically generated value resembling:

```text
service.name = unknown_service:ai-incremental-logging-server
```

to:

```text
service.name = greeting-service
service.version = 1.0.0
```

The SDK metadata remains present.

### Important distinction

Resource attributes describe the producer and apply broadly:

```text
service.name = greeting-service
```

Log attributes describe one particular event:

```text
name = Daniel
```

The `greet` handler does not need to repeat the service identity on every log call because
the LoggerProvider supplies it.

### Deliberately not introduced yet

The instrumentation scope still has no version or schema URL. There is no active span,
valid trace ID, valid span ID, batch processor, OTLP exporter, Collector, or trace/log
correlation.

## Step 8: Complete the instrumentation scope metadata

### What changed

Added an instrumentation version and semantic-convention schema URL when creating the
bridged `slog.Logger`:

```go
const instrumentationVersion = "1.0.0"

slog.SetDefault(
    otelslog.NewLogger(
        instrumentationName,
        otelslog.WithVersion(
            instrumentationVersion,
        ),
        otelslog.WithSchemaURL(
            semconv.SchemaURL,
        ),
    ),
)
```

### Concept demonstrated

The instrumentation scope identifies the code or library that created a telemetry record.
Its name, version, and schema URL allow a backend to distinguish records produced by
different instrumentation within the same service.

The schema URL declares which OpenTelemetry schema is used to interpret standardized
telemetry fields.

### Expected behavior

The exported Scope changes from:

```text
Name = ai-incremental-logging
Version = ""
SchemaURL = ""
```

to:

```text
Name = ai-incremental-logging
Version = 1.0.0
SchemaURL = https://opentelemetry.io/schemas/1.41.0
```

The log body, INFO severity, `name` attribute, and service Resource remain unchanged.

### Important distinction

Although both versions are `1.0.0` in this study application, they identify different
things:

- Resource attribute `service.version` identifies the running greeting service version.
- Scope `Version` identifies the version of the instrumentation producing the logs.

A service can use several instrumentation libraries with versions that evolve
independently from the service itself.

### Deliberately not introduced yet

The scope has no custom scope attributes. There is no active span, valid trace ID, valid
span ID, batch processor, OTLP exporter, Collector, or trace/log correlation.

## Step 9: Correlate a log record with an active span

### What changed

Added the familiar tracing pipeline using a stdout trace exporter, simple span processor,
TracerProvider, and Tracer. The logging and tracing providers share the same service
Resource.

The greeting handler now starts one span:

```go
ctx, span := tracer.Start(
    r.Context(),
    "produce greeting",
)
defer span.End()
```

The existing context-aware logging call now receives the Context returned by
`tracer.Start`:

```go
slog.InfoContext(
    ctx,
    "greeting produced",
    slog.String("name", name),
)
```

### Concept demonstrated

`tracer.Start` returns a new Context containing the active span. When the bridged
`slog.InfoContext` call passes that Context to OpenTelemetry, the log SDK extracts the
span context and places its trace ID, span ID, and trace flags into the log record.

```text
tracer.Start
    -> Context containing active span
    -> slog.InfoContext
    -> otelslog bridge
    -> OTel log with trace and span identifiers
```

The log record can therefore point to the exact trace operation that produced the event.

### Expected behavior

One greeting request produces both a log record and a completed span. Their identifiers
match:

```text
log.TraceID == span.SpanContext.TraceID
log.SpanID  == span.SpanContext.SpanID
```

Unlike earlier steps, the log fields are no longer all-zero identifiers:

```text
TraceID = 8a9f...
SpanID = 37bc...
TraceFlags = 01
```

The exact identifiers are generated independently for every request.

### Important distinction

The log record is not a span and is not stored inside the span. They are separate
telemetry records connected by identifiers.

Passing the original `r.Context()` to the log call would still produce zero identifiers
in this manually instrumented example because the new span is stored in the Context
returned by `tracer.Start`. The variable `ctx` must be passed onward.

The tracing scaffolding is not itself a new logging mechanism. Its purpose in this step is
to create the active span needed to demonstrate log correlation.

### Deliberately not introduced yet

There is only one manual span and no automatic HTTP instrumentation or remote context
propagation. Both logs and spans use simple processors and stdout exporters. There is no
batching, OTLP, Collector, backend, or clickable trace-to-log navigation.

## Step 10: Record a failed event with ERROR severity

### What changed

Added a deliberately simulated failure for this request:

```text
GET /hello/fail
```

The handler emits a structured ERROR record using the existing span-containing Context
and returns HTTP status 500:

```go
if name == "fail" {
    slog.ErrorContext(
        ctx,
        "greeting failed",
        slog.String("name", name),
        slog.String("reason", "simulated failure"),
    )
    http.Error(
        w,
        "failed to produce greeting",
        http.StatusInternalServerError,
    )
    return
}
```

### Concept demonstrated

Log severity communicates the importance and nature of an event. A successful greeting
uses INFO, while the simulated inability to produce one uses ERROR.

The bridge maps Go's `slog.LevelError` to the corresponding OpenTelemetry severity
number and text while preserving the message and structured attributes.

### Expected behavior

A normal request continues to return HTTP 200 and emit:

```text
SeverityText = INFO
Body = greeting produced
name = Daniel
```

The simulated failure returns HTTP 500 and emits:

```text
SeverityText = ERROR
Body = greeting failed
name = fail
reason = simulated failure
```

The ERROR log's trace ID and span ID match the span created for that failed request.

### Important distinction

Log severity and span status belong to different telemetry records:

- `slog.ErrorContext` marks the log record as ERROR.
- It does not automatically set the correlated span's status to Error.

The span in this step still has its default unset status. If we want the trace itself to
represent failure, application code must set the span status separately.

The HTTP status code is also separate. Calling `http.Error` returns a 500 response, while
the ERROR log records the operational event that explains why.

### Deliberately not introduced yet

The failure is simulated and does not wrap a real Go error. We have not set span status,
recorded a span exception, added automatic HTTP instrumentation, switched to batch
processors, or configured OTLP, a Collector, or a backend.

## Step 11: Batch log export away from the request path

### What changed

Replaced the simple log processor with a batch processor:

```go
logProcessor := sdklog.NewBatchProcessor(
    logExporter,
)
```

The LoggerProvider configuration and all `slog` calls remain unchanged.

### Concept demonstrated

The batch processor queues OTel log records and exports them asynchronously in groups:

```text
slog call
    -> otelslog bridge
    -> batch processor queues record
    -> request continues

batch processor
    -> exports queued records later
```

This keeps exporter work away from the request path. That is particularly useful when a
production exporter performs network I/O.

### Expected behavior

The exported log JSON contains the same body, severity, attributes, Resource, scope, and
trace correlation fields as before. It may appear shortly after the request rather than
during the `slog.InfoContext` or `slog.ErrorContext` call.

Tracing deliberately retains its simple processor. Consequently, the completed span may
be printed before its correlated batched log even though the log event occurred before
`span.End`.

### Important distinction

Batching changes when records are exported, not when their events occurred. Each record
retains its original event `Timestamp` and its `ObservedTimestamp`.

Preserving timestamps does not guarantee arrival order. Concurrent requests, separate
service instances, batch boundaries, network delays, and retries can all cause a backend
to receive records in a different order from their event timestamps.

LoggerProvider shutdown flushes records that remain queued when `run` returns cleanly.
The current study server does not yet handle operating-system termination signals
gracefully, so forcibly terminating it can still lose a final queued record.

### Deliberately not introduced yet

We use the batch processor's default queue, batch, and export timing settings. We have
not tuned them, added graceful signal handling, batched spans, configured OTLP, or added
a Collector or backend.

## Step 12: Let HTTP instrumentation create the request span

### What changed

Removed the manual `tracer.Start` call from `greet` and wrapped the HTTP router with
OpenTelemetry's `otelhttp` instrumentation:

```go
instrumentedHandler := otelhttp.NewHandler(
    mux,
    "greeting-server",
)

http.ListenAndServe(
    listenAddress,
    instrumentedHandler,
)
```

The application log calls now use the request Context directly:

```go
slog.InfoContext(
    r.Context(),
    "greeting produced",
    slog.String("name", name),
)
```

### Concept demonstrated

Instrumentation middleware can create telemetry around framework or standard-library
operations without requiring each application handler to implement that telemetry.

`otelhttp` starts a server span before calling our router, stores the span in a new request
Context, and passes that enriched request to `greet`. The existing OTel `slog` bridge then
extracts the server span's identifiers from `r.Context()`.

```text
HTTP request
    -> otelhttp starts server span
    -> otelhttp enriches r.Context()
    -> greet calls slog.InfoContext(r.Context())
    -> log is correlated with server span
    -> otelhttp ends server span
```

### Expected behavior

A request still produces one correlated log and span with matching identifiers:

```text
log.TraceID == server span.TraceID
log.SpanID  == server span.SpanID
```

The trace record now represents an HTTP server operation and includes HTTP attributes
captured by the instrumentation library. Its instrumentation scope identifies the
`otelhttp` package rather than our manual tracer.

### Important distinction

Tracing the HTTP request is now automatic, but application logging remains manual:

- `otelhttp` decides when to start and end the HTTP server span.
- Our code still decides that `greeting produced` and `greeting failed` are meaningful
  events and explicitly calls `slog`.

Automatic HTTP instrumentation does not know which business events are worth recording
as application logs.

The `greet` function could still create a manual child span if producing a greeting
became a substantial business operation. This step removes it only to isolate the
middleware-provided server span.

### Deliberately not introduced yet

We use the instrumentation library's default span naming and HTTP attributes. We have
not configured incoming remote-context propagation explicitly, added custom server-span
attributes, tuned batching, configured OTLP, or added a Collector or backend.

## Step 13: Final Chapter 6 logging concept map

### What changed

No program behavior changed. This final study step maps the completed example to the
logging concepts introduced throughout Chapter 6.

### Log record field map

The `otelslog` bridge converts fields from Go's standard logging model into the
OpenTelemetry log data model:

| Source | OpenTelemetry log field | Example |
| --- | --- | --- |
| Time on `slog.Record` | `Timestamp` | Time when `slog.InfoContext` was called |
| Time observed by OTel | `ObservedTimestamp` | Time when the SDK received the record |
| `slog.LevelInfo` or `slog.LevelError` | `Severity` and `SeverityText` | `INFO` or `ERROR` |
| `slog` message | `Body` | `greeting produced` |
| `slog.Attr` values | `Attributes` | `name=Daniel` |
| Active span in the supplied Context | `TraceID`, `SpanID`, and `TraceFlags` | Identifiers shared with the server span |
| LoggerProvider Resource | `Resource` | `service.name=greeting-service` |
| Bridge name and options | `Scope` | `ai-incremental-logging`, version `1.0.0` |

### Logging pipeline map

```text
slog.InfoContext or slog.ErrorContext
    -> default slog.Logger
    -> otelslog Handler
    -> OTel Logger from the LoggerProvider
    -> batch LogRecordProcessor
    -> stdout log exporter
    -> pretty-printed OTel JSON
```

Each component has a distinct responsibility:

| Component | Responsibility |
| --- | --- |
| `slog.Logger` | Provides the application-facing standard Go logging API. |
| `otelslog.Handler` | Converts `slog.Record` values into OTel log records. |
| `LoggerProvider` | Owns the logging SDK pipeline and supplies OTel Loggers. |
| Batch processor | Queues records and exports them away from the request path. |
| Stdout exporter | Encodes completed OTel log records as JSON. |
| Resource | Identifies the service producing all records from the provider. |
| Instrumentation scope | Identifies the code or library that created the records. |

### Request and correlation map

```text
HTTP request
    -> otelhttp starts an HTTP server span
    -> otelhttp places that span in a new request Context
    -> mux calls greet with the enriched request
    -> greet passes r.Context() to slog
    -> otelslog passes the Context into the OTel Log SDK
    -> Log SDK copies the active span identifiers into the log record
```

The resulting log and span are separate telemetry records connected by identifiers:

```text
log.TraceID == span.TraceID
log.SpanID  == span.SpanID
```

The request Context contains request-scoped state such as the active span, cancellation,
and deadlines. It does not contain the exporters, processors, providers, or Resource;
those belong to the telemetry SDK configuration.

### Metadata scope map

| Metadata | Scope | Example |
| --- | --- | --- |
| Resource attribute | All telemetry from the provider | `service.name=greeting-service` |
| Instrumentation-scope metadata | Records created by one instrumentation source | `Scope.Name=ai-incremental-logging` |
| Log attribute | One log event | `name=Daniel` |
| Trace and span identifiers | Correlation with one trace operation | `TraceID=...`, `SpanID=...` |

The same semantic-convention schema URL can appear on both a Resource and an
instrumentation scope, but it describes the schema of each separate metadata location.

### Severity and failure map

The example deliberately produces two event types:

| Request | HTTP status | Log severity | Log body |
| --- | --- | --- | --- |
| `/hello/Daniel` | 200 | `INFO` | `greeting produced` |
| `/hello/fail` | 500 | `ERROR` | `greeting failed` |

HTTP status, log severity, and span status are independent values. An ERROR log does not
by itself mark a span as failed, and returning HTTP 500 does not replace the need for a
useful application log explaining the event.

### Manual and automatic instrumentation map

- `otelhttp` automatically creates the HTTP server span and captures generic HTTP
  information such as method, route, status, and request duration.
- Application code manually emits `greeting produced` and `greeting failed` because
  only the application understands which business events are meaningful.
- Automatic request instrumentation measures the complete handler operation but cannot
  explain time spent inside application-specific file, database, or downstream-service
  operations. Those would require additional manual child spans.

### Processing and ordering map

The batch log processor delays export but preserves each record's event timestamp.
Export order is not guaranteed to equal event order:

```text
log event occurs inside span
    -> span ends and its simple processor exports immediately
    -> batch log processor exports the correlated log later
```

Backends use timestamps and correlation identifiers rather than relying solely on
ingestion order.

### Expected final behavior

Run the server and make both requests:

```bash
curl http://localhost:8080/hello/Daniel
curl -i http://localhost:8080/hello/fail
```

The server exports an INFO log and an ERROR log, each correlated with its automatically
created HTTP server span. Both signals identify the same `greeting-service` Resource.

### Deliberately not introduced

This example stops at local stdout export. It does not configure OTLP, an OpenTelemetry
Collector, Loki or another log backend, incoming remote trace-context propagation,
graceful operating-system signal handling, sampling policy, automatic application-event
logging, or dashboard navigation between logs and traces.

The exercise also does not call the low-level OpenTelemetry Logs API directly.
Application code continues to use Go's established `slog` API, while the bridge handles
conversion into the OTel data model.
