# Incremental steps taken

## Step 1: Create two ordinary HTTP services

### What changed

Created two small Go services with no telemetry:

```text
caller
  |
  | GET /hello/{name}
  v
greeting service (:8080)
  |
  | GET /format/{name}
  v
formatting service (:8081)
```

The greeting service makes an ordinary outbound HTTP request:

```go
request, err := http.NewRequestWithContext(
    r.Context(),
    http.MethodGet,
    endpoint,
    nil,
)
response, err := client.Do(request)
```

The formatting service turns the name into the final response:

```go
func formatGreeting(w http.ResponseWriter, r *http.Request) {
    name := r.PathValue("name")
    _, _ = fmt.Fprintf(w, "Hello, %s!\n", name)
}
```

### Concept demonstrated

This establishes the ordinary application behavior before introducing
instrumentation libraries. There are two separate service boundaries and two HTTP
operations:

1. The inbound request handled by the greeting service.
2. The outbound request from the greeting service to the formatting service.
3. The inbound request handled by the formatting service.

These operations already happen, but we cannot yet observe their relationship as a
distributed trace.

### Expected behavior

Start the formatting service in one terminal:

```bash
go run ./cmd/formatting-service
```

Start the greeting service in another terminal:

```bash
go run ./cmd/greeting-service
```

Then call the public endpoint:

```bash
curl http://localhost:8080/hello/Daniel
```

The response should be:

```text
Hello, Daniel!
```

If the formatting service is not running, the greeting service returns HTTP `502 Bad
Gateway` with:

```text
formatting service unavailable
```

### Important distinctions

The two programs are separate services even though they live in the same Go module and
run on the same computer. An HTTP request crosses the boundary between them.

The outbound request inherits `r.Context()` so cancellation can flow from the caller to
the downstream request. This is ordinary Go context propagation. The Context does not
yet contain an OpenTelemetry span, and no trace context is injected into HTTP headers.

The calls to `http.ListenAndServe` and `client.Do` use the standard `net/http`
implementation directly. They do not automatically produce OpenTelemetry telemetry.

### Deliberately not introduced yet

There is no OpenTelemetry API or SDK, TracerProvider, exporter, processor, Resource,
span, propagation header, `otelhttp.NewHandler`, `otelhttp.NewTransport`, metric, log
correlation, Collector, or observability backend.

## Step 2: Instrument the greeting service's inbound HTTP requests

### What changed

Added the familiar stdout tracing pipeline to the greeting service and registered its
`TracerProvider` globally. Then wrapped the existing HTTP multiplexer with the
`otelhttp` instrumentation library:

```go
instrumentedHandler := otelhttp.NewHandler(
    mux,
    "greeting-server",
)

return http.ListenAndServe(
    listenAddress,
    instrumentedHandler,
)
```

The provider has a Resource that identifies spans as coming from
`service.name=greeting-service`.

### Concept demonstrated

An **instrumentation library** knows how a particular library or framework operates.
`otelhttp.NewHandler` understands Go's `net/http` server interfaces. Its returned
handler automatically:

- Starts a server span before calling the wrapped handler.
- Places that span in the request Context.
- Measures the request's duration.
- Adds standard HTTP attributes.
- Ends the span after the wrapped handler returns.

The application does not call `tracer.Start` or `span.End` inside `greet`.

The exporter, processor, and provider do not instrument HTTP themselves. They make it
possible to process and observe the span that `otelhttp` creates.

### Expected output or behavior

The HTTP behavior remains the same:

```text
GET /hello/Daniel -> Hello, Daniel!
```

The greeting service now prints one JSON span resembling:

```text
Name: GET /hello/{name}
SpanKind: server
Resource:
  service.name: greeting-service
Attributes:
  http.request.method: GET
  http.response.status_code: 200
  url.path: /hello/Daniel
```

The exact output contains additional HTTP, network, SDK, timing, trace ID, and span ID
fields supplied automatically.

### Important distinctions

This is **code-based library instrumentation**, not manual span instrumentation:

- We added the `otelhttp` wrapper in our source code.
- Once activated, the wrapper creates each HTTP server span automatically.
- Manual instrumentation would call the tracing API directly inside application code.

It is also not true **zero-code instrumentation**, because activating the library required
changing the program.

Although `r.Context()` now contains the greeting server span, the ordinary
`http.Client` does not inspect that span or inject it into HTTP headers. Passing a Context
to an ordinary client is not by itself distributed trace propagation.

### Deliberately not introduced yet

The outbound HTTP client is not wrapped with `otelhttp.NewTransport`. It creates no
client span and injects no trace context. The formatting service remains completely
uninstrumented, so its work does not appear in the trace.

There are no custom span names, route filters, automatic HTTP metrics under study,
double-instrumentation examples, zero-code mechanisms, Collector, or observability
backend.

## Step 3: Instrument the greeting service's outbound HTTP client

### What changed

Replaced the greeting service's default HTTP transport with an instrumented transport:

```go
client := &http.Client{
    Transport: otelhttp.NewTransport(
        http.DefaultTransport,
    ),
}
```

The request itself is still made with the standard `http.Client` API:

```go
response, err := client.Do(request)
```

### Concept demonstrated

An `http.Client` delegates the network exchange to an implementation of
`http.RoundTripper`, stored in its `Transport` field. `otelhttp.NewTransport` wraps an
ordinary `RoundTripper` and observes requests passing through it.

For each outbound request, the wrapper automatically:

- Starts a span with span kind `CLIENT`.
- Uses the active span in the request Context as its parent.
- Measures the network exchange.
- Adds standard HTTP client attributes.
- Records the response status.
- Ends the span when the response body is closed or reaches end-of-file.

No tracing API calls were added to `greet`.

### Expected output or behavior

The HTTP response remains:

```text
Hello, Daniel!
```

The greeting service now exports two spans:

```text
GET /hello/{name}                  SERVER
  `-- GET                          CLIENT
```

The relationship can be verified from the JSON:

- Both spans have the same `TraceID`.
- The client span's parent span ID is the server span's span ID.
- The server span reports one child.
- Both spans have `service.name=greeting-service` because both were created by that
  service's provider.

The client span is normally printed first because it finishes before the surrounding
server span.

### Important distinctions

`otelhttp.NewHandler` instruments inbound server handling, whereas
`otelhttp.NewTransport` instruments outbound client requests. They are two different
activation points in the same instrumentation package.

The parent/child relationship inside the greeting service works because the outbound
request carries `r.Context()`. The transport can read the active server span from that
in-memory Context.

The transport also contains the mechanism that can inject trace context into outbound
HTTP headers. However, we have not configured a text-map propagator yet. We will study
cross-process propagation separately rather than treating client span creation and
propagator configuration as one concept.

### Deliberately not introduced yet

The formatting service is still uninstrumented and produces no span. A W3C
`traceparent` propagator has not been configured, so the next process cannot yet extract
the greeting service's trace relationship.

There are no custom span names, route filters, automatic HTTP metrics under study,
double-instrumentation examples, zero-code mechanisms, Collector, or observability
backend.

## Step 4: Propagate trace context to the formatting service

### What changed

Configured the greeting service to use the W3C Trace Context propagator before creating
its instrumented HTTP transport:

```go
otel.SetTextMapPropagator(
    propagation.TraceContext{},
)
```

Added one diagnostic print to the formatting service so the injected header was directly
observable:

```go
fmt.Printf(
    "received traceparent: %s\n",
    r.Header.Get("traceparent"),
)
```

### Concept demonstrated

An in-memory Go `Context` cannot cross a process or network boundary. A **propagator**
translates selected telemetry context into a carrier that can cross that boundary.

For HTTP, the carrier is the request headers. The W3C Trace Context propagator writes a
`traceparent` header with this structure:

```text
version-trace-id-parent-id-trace-flags
```

For example:

```text
00-973cbc4930b043b992d7ae83809c4725-3864f846310ac174-01
```

`otelhttp.NewTransport` creates the client span and then asks the configured propagator
to inject that span's context into the outbound request headers.

### Expected output or behavior

The application still returns:

```text
Hello, Daniel!
```

At this step, the formatting service printed a line resembling:

```text
received traceparent: 00-973cbc4930b043b992d7ae83809c4725-3864f846310ac174-01
```

Comparing that header with the client span exported by the greeting service showed:

- The header's trace ID equaled the client and greeting server spans' `TraceID`.
- The header's parent ID equaled the client span's `SpanID`.
- The final `01` indicated the sampled trace flag was set.

### Important distinctions

**Injection** serializes context into an outbound carrier. **Extraction** reads that
serialized data at the receiving service and reconstructs an OpenTelemetry Context.
This step performed injection only.

The `traceparent` field is named `parent-id` by the W3C format because it tells the next
service which span should become the parent of its new span. At the sender, that value is
the current client span's ID.

At this step, the formatting service's `r.Header.Get` call merely read a string. It did
not make that trace context active, create a span, or connect the service to the
distributed trace.

### Deliberately not introduced yet

At this step, the formatting service still had no OpenTelemetry SDK, propagator
configuration, or instrumented handler. It could not extract the header and emitted no
span.

Baggage propagation, custom propagators, custom span names, route filters, automatic
HTTP metrics under study, double instrumentation, zero-code mechanisms, a Collector,
and an observability backend had not been introduced.

## Step 5: Extract context and instrument the formatting service

### What changed

Added a separate stdout tracing pipeline and Resource to the formatting service:

```go
tracerProvider := sdktrace.NewTracerProvider(
    sdktrace.WithSpanProcessor(traceProcessor),
    sdktrace.WithResource(serviceResource),
)
otel.SetTracerProvider(tracerProvider)
```

Configured the same W3C Trace Context propagator on the receiving side and wrapped the
formatting service's mux:

```go
otel.SetTextMapPropagator(
    propagation.TraceContext{},
)

instrumentedHandler := otelhttp.NewHandler(
    mux,
    "formatting-server",
)
```

Removed the diagnostic `r.Header.Get("traceparent")` print. The instrumentation library
now consumes the header for a functional purpose.

### Concept demonstrated

`otelhttp.NewHandler` performs both receiving-side operations:

1. It asks the propagator to **extract** the remote span context from the incoming
   `traceparent` header.
2. It creates a new `SERVER` span using that extracted client span as its parent.

The resulting formatting server span becomes active in `r.Context()` while
`formatGreeting` runs. The handler itself still contains no calls to `tracer.Start`,
`span.End`, or the propagation API.

### Expected output or behavior

The application response remains:

```text
Hello, Daniel!
```

The two service terminals collectively export this distributed trace:

```text
greeting-service:   GET /hello/{name}    SERVER
  `-- greeting-service: HTTP GET         CLIENT
        `-- formatting-service:
            GET /format/{name}           SERVER
```

All three spans share the same `TraceID`. The formatting server span's parent fields
contain:

- The client span's trace ID.
- The client span's span ID.
- `Remote: true`, showing that the parent context arrived from another process.

The formatting span has `service.name=formatting-service`, while the first two spans
have `service.name=greeting-service`.

### Important distinctions

Propagation connects spans; it does not send completed span records between services.
Each service creates and exports its own spans through its own provider:

- The greeting service exports the greeting server and HTTP client spans.
- The formatting service exports the formatting server span.

They form one distributed trace because their IDs and parent relationships connect them.
A tracing backend would assemble these independently received records into the tree.

A **remote parent** does not mean the parent span object or its full telemetry record was
transferred. Only the compact span context required for correlation crossed the HTTP
boundary.

The Resource remains local to the service creating a span. Resource attributes are not
propagated through `traceparent`.

### Deliberately not introduced yet

There are no manual child spans inside either handler, Baggage propagation, custom
propagators, custom span names, route filters, automatic HTTP metrics under study,
double-instrumentation examples, zero-code mechanisms, Collector, or observability
backend.

## Step 6: Filter health checks from automatic instrumentation

### What changed

Added an ordinary health endpoint to the greeting service:

```go
mux.HandleFunc("GET /health", health)
```

Configured its `otelhttp` server instrumentation with a request filter:

```go
instrumentedHandler := otelhttp.NewHandler(
    mux,
    "greeting-server",
    otelhttp.WithFilter(func(r *http.Request) bool {
        return r.URL.Path != "/health"
    }),
)
```

### Concept demonstrated

Instrumentation libraries provide automatic behavior, but that behavior is
configurable. An `otelhttp` filter decides whether the wrapper should instrument each
request:

- `true` means create the automatic HTTP telemetry.
- `false` means call the wrapped handler without creating that telemetry.

The filter excludes `/health` because health probes are typically repetitive and may
produce a large volume of low-value telemetry.

### Expected output or behavior

Calling the application route still returns a greeting and exports the three-span
distributed trace:

```bash
curl http://localhost:8080/hello/Daniel
```

Calling the health route returns a normal response:

```bash
curl http://localhost:8080/health
```

```text
OK
```

No span is exported for the health request.

### Important distinctions

The filter controls instrumentation, not request routing or authorization. Returning
`false` does not reject the request: the `/health` handler still executes and responds.

The filter is attached to the greeting service's `NewHandler`, so it controls only
automatic inbound server telemetry from that wrapper. It does not globally disable
tracing, affect the formatting service, or filter an independently instrumented HTTP
client.

Filtering happens before a server span is created. This differs from sampling, where the
tracing system makes a recording/export decision about a span or trace. We have not
introduced sampling configuration here.

### Deliberately not introduced yet

There are no configurable filter lists, environment-variable exclusions, manual child
spans, Baggage propagation, custom propagators, custom span names, automatic HTTP
metrics under study, double-instrumentation examples, zero-code mechanisms, Collector,
or observability backend.

## Step 7: Export HTTP metrics recorded by the instrumentation library

### What changed

Added a stdout metric exporter and a five-second periodic reader to the greeting service:

```go
metricReader := sdkmetric.NewPeriodicReader(
    metricExporter,
    sdkmetric.WithInterval(5*time.Second),
)

meterProvider := sdkmetric.NewMeterProvider(
    sdkmetric.WithReader(metricReader),
    sdkmetric.WithResource(serviceResource),
)
otel.SetMeterProvider(meterProvider)
```

The provider is registered before `otelhttp.NewTransport` and
`otelhttp.NewHandler` are constructed. No counters or histograms were declared or
recorded in the application handlers.

### Concept demonstrated

An instrumentation library can use more than one OpenTelemetry signal. The same
`otelhttp` wrappers now:

- Create spans through the configured `TracerProvider`.
- Record HTTP measurements through the configured `MeterProvider`.

`otelhttp` defines the appropriate metric instruments, records them at the correct
points in the HTTP lifecycle, and supplies standard HTTP attributes. Our metric SDK
pipeline aggregates and exports those measurements.

### Expected output or behavior

After calling:

```bash
curl http://localhost:8080/hello/Daniel
```

the greeting service continues to export its server and client spans. At the periodic
collection interval, it also exports metrics including:

```text
http.server.request.duration
http.server.request.body.size
http.server.response.body.size
http.client.request.duration
http.client.request.body.size
```

The exact set depends on which applicable HTTP measurements were observed. Metric data
points contain automatically supplied attributes such as the HTTP method, status code,
server address, server port, protocol version, and route where applicable.

The Resource identifies these metric streams as coming from `greeting-service`.

The observed histogram data points also contained exemplars with trace and span IDs:

```text
Exemplars:
  Value: ...
  SpanID: ...
  TraceID: ...
```

Because `otelhttp` recorded the measurements with an active sampled span in the
Context, the metric SDK could retain representative measurements that link the
aggregated metric data back to specific traces. Not every measurement is guaranteed to
be retained as an exemplar; exemplar selection is handled by the SDK.

Calling only:

```bash
curl http://localhost:8080/health
```

does not add server measurements because the handler filter skips all of that wrapper's
telemetry for the health request.

### Important distinctions

The instrumentation library **records measurements** when HTTP operations occur. The
periodic reader **collects and exports aggregates** every five seconds. The reader does
not cause HTTP measurements to happen.

A trace span describes one specific HTTP operation. A metric data point summarizes
measurements from one or more operations with the same attribute set.

An exemplar is neither another metric stream nor another span. It is a representative
measurement stored with a metric data point that carries correlation information for a
trace. We have observed this behavior here but have deliberately not configured or
explored exemplar reservoirs and filtering in this step.

The five-second interval exists only to make the example easy to observe. It is an
export/collection interval, not a promise that every individual request will become its
own metric data point.

The filter is evaluated before both the server span and server metrics are produced.
It is therefore broader than a trace-only exclusion in this wrapper.

### Deliberately not introduced yet

The formatting service still has no metric SDK pipeline, so its `otelhttp.NewHandler`
uses a no-op MeterProvider even though it produces spans.

There are no custom views, custom buckets, exemplars under study, manual HTTP metrics,
custom span names, double-instrumentation examples, zero-code mechanisms, Collector, or
observability backend.

## Step 8: Customize automatic server span names

### What changed

Added a span-name formatter to the greeting service's existing `otelhttp.NewHandler`
configuration:

```go
instrumentedHandler := otelhttp.NewHandler(
    mux,
    "greeting-server",
    // Other options omitted here.
    otelhttp.WithSpanNameFormatter(
        greetingSpanName,
    ),
)
```

The callback uses the resolved route pattern:

```go
func greetingSpanName(
    operation string,
    r *http.Request,
) string {
    if r.Pattern == "" {
        return operation
    }
    return "greeting " + r.Pattern
}
```

### Concept demonstrated

Instrumentation libraries normally provide useful semantic-convention-based defaults,
but they can expose callbacks or options for cases where an application needs controlled
customization. `WithSpanNameFormatter` lets the application choose the names of server
spans that the library still creates and manages automatically.

No calls to `tracer.Start`, `span.SetName`, or `span.End` were added to the handler.

### Expected output or behavior

The greeting server span name changes from:

```text
GET /hello/{name}
```

to:

```text
greeting GET /hello/{name}
```

The client and formatting server span names remain unchanged:

```text
greeting GET /hello/{name}          SERVER
  `-- HTTP GET                      CLIENT
        `-- GET /format/{name}      SERVER
```

HTTP responses, parent/child relationships, propagation, attributes, metrics, and the
health filter behave as before.

### Important distinctions

The callback uses `r.Pattern`, which is `GET /hello/{name}` for this method-aware
`ServeMux` route, rather than `r.URL.Path`, such as `/hello/Daniel`. The pattern already
contains the HTTP method. A route pattern creates a stable, low-cardinality operation
name that can group many requests.

Putting user-controlled identifiers into span names would produce high-cardinality names
and make operations harder to aggregate:

```text
Good: greeting GET /hello/{name}
Bad:  greeting GET /hello/Daniel
Bad:  greeting GET /hello/Alice
```

Span names and metric attributes are separate pieces of telemetry. This formatter does
not rename metrics or change their `http.route` attributes.

The fallback operation name is used before or when routing cannot resolve a pattern.
The `otelhttp` wrapper invokes the formatter again after the mux resolves a route, so
the exported span receives the route-based name for a matched endpoint.

### Deliberately not introduced yet

There are no custom metric views, custom buckets, custom span attributes, manual child
spans, double-instrumentation examples, zero-code mechanisms, Collector, or
observability backend.

## Step 9: Configure Resource identity with standard environment variables

### What changed

Removed the hard-coded greeting service Resource attributes and used the SDK's default
Resource:

```go
serviceResource := resource.Default()
```

Also removed all application-specific service flags and configuration precedence code.

### Concept demonstrated

`resource.Default()` includes the Go SDK's environment Resource detector. It reads the
standard OpenTelemetry variables:

```text
OTEL_SERVICE_NAME
OTEL_RESOURCE_ATTRIBUTES
```

This allows deployment configuration to identify the service without rebuilding the
application or adding application-specific configuration parsing.

### Expected output or behavior

The program still starts with no arguments or environment variables:

```bash
go run ./cmd/greeting-service
```

In that case, the SDK supplies a default service name resembling:

```text
service.name = unknown_service:greeting-service
```

The exact executable suffix can differ when using `go run`, because Go executes a
temporary compiled binary.

For a meaningful identity, start the service with standard OTel environment variables:

```bash
OTEL_SERVICE_NAME=greeting-service \
OTEL_RESOURCE_ATTRIBUTES="service.version=1.0.0,deployment.environment.name=study" \
go run ./cmd/greeting-service
```

The exported trace and metric Resources then contain:

```text
service.name = greeting-service
service.version = 1.0.0
deployment.environment.name = study
```

### Important distinctions

Environment variables are process environment, not command-line arguments. The
application receives zero program arguments in both launch examples.

This step automatically configures only Resource attributes. The application still
constructs its trace and metric exporters, processors, reader, providers, and
instrumentation wrappers in code.

`resource.Default()` reads and caches its detected Resource the first time it is called
in a process. Environment variables should therefore be set before starting the program,
not changed after telemetry initialization.

Resource attributes describe the entity producing telemetry. They do not become span
attributes, metric attributes, propagation headers, or HTTP request parameters.

### Deliberately not introduced yet

Exporter selection, processor configuration, sampling, and propagator selection are
still explicit in code. There is no full automatic SDK configurator, double
instrumentation, manual internal child span, zero-code mechanism, Collector, or
observability backend.

## Step 10: Add manual instrumentation inside an automatic HTTP span

### What changed

Extracted the formatting service's application-specific string construction into a
function that accepts the handler Context:

```go
greeting := buildGreeting(r.Context(), name)
```

Added one manual span inside that function:

```go
func buildGreeting(
    ctx context.Context,
    name string,
) string {
    _, span := otel.Tracer(
        formattingInstrumentationName,
    ).Start(ctx, "build greeting")
    defer span.End()

    return fmt.Sprintf("Hello, %s!", name)
}
```

### Concept demonstrated

Instrumentation libraries understand known library operations, such as an HTTP server
handling a request. They do not understand which pieces of application logic are
meaningful operations.

Manual instrumentation complements the automatic server span by describing the
application-specific `build greeting` operation. Because `buildGreeting` receives
`r.Context()`, the tracing API finds the formatting server span there and assigns it as
the manual span's parent.

No new exporter, processor, provider, Resource, or propagator was required.

### Expected output or behavior

The HTTP response remains:

```text
Hello, Daniel!
```

The distributed trace now contains four spans:

```text
greeting GET /hello/{name}       SERVER   automatic
  `-- HTTP GET                   CLIENT   automatic
        `-- GET /format/{name}   SERVER   automatic
              `-- build greeting INTERNAL manual
```

The manual span:

- Shares the distributed trace ID.
- Uses the formatting server span ID as its parent.
- Has the default `INTERNAL` span kind.
- Has `service.name=formatting-service` because it uses that service's provider.
- Uses `ai-incremental-instrumentation-libraries/formatting-service` as its
  instrumentation scope instead of the `otelhttp` scope.

### Important distinctions

Automatic and manual instrumentation are complementary, not mutually exclusive.
Automatic instrumentation covers common framework boundaries cheaply; manual spans add
domain-specific visibility where it is useful.

The Context returned by `Tracer.Start` is ignored here because `buildGreeting` has no
deeper operation to instrument. If it called another traced function, that returned
Context should be passed to the child operation.

The Tracer's name identifies the instrumentation scope. It is not the service name and
does not replace the Resource.

The span measures only string construction in this deliberately small example. In a real
service, manual spans are most valuable around meaningful operations such as a database
query, algorithm, cache lookup, or external-library call not already instrumented.

### Deliberately not introduced yet

The manual span has no custom attributes, events, links, explicit status, or error
recording. There is no double-instrumentation example, full automatic SDK configurator,
zero-code mechanism, Collector, or observability backend.

## Concept note: Avoid double instrumentation

Double instrumentation happens when more than one instrumentation layer observes the
same operation unintentionally. For example, wrapping our already instrumented greeting
handler again would be suspicious:

```go
// instrumentedHandler already contains the intended otelhttp wrapper.
duplicateHandler := otelhttp.NewHandler(
    instrumentedHandler,
    "duplicate-greeting-server",
)
```

A single request could then produce a misleading trace:

```text
duplicate greeting SERVER
  `-- greeting SERVER
        `-- HTTP CLIENT
              `-- formatting SERVER
                    `-- build greeting INTERNAL
```

Both server wrappers could also record the same HTTP server measurements, causing one
real request to increase metric counts twice. Other consequences include unnecessary
processing, greater telemetry volume and storage cost, and confusing request durations.

Some agent-style instrumentors remember whether they have already patched a library and
warn or refuse to instrument it again. Go HTTP middleware cannot generally assume that
nested wrappers are accidental: middleware layers may legitimately represent different
operations. It is therefore our responsibility to understand where instrumentation is
activated and avoid semantically duplicating the same boundary.

This anti-pattern was deliberately **not implemented**. Keeping the final program
correct is more useful than adding and subsequently removing the mistake, while this
note preserves the lesson in version control.

## Concept note: Finding instrumentation libraries for common SRE and AWS work

The [OpenTelemetry registry](https://opentelemetry.io/ecosystem/registry/) is the
central catalogue, but it contains several different kinds of components. When looking
for automatic application instrumentation, first filter for the application's language
and for **instrumentation**. A Collector receiver, exporter, propagator, or Resource
detector can be valuable, but it does a different job.

A practical search order is:

1. Identify the library actually used by the program, including its major version. For
   example, search for the AWS SDK for Go **v2**, not merely "AWS".
2. Check the OpenTelemetry registry and the
   [OpenTelemetry Go instrumentation documentation](https://opentelemetry.io/docs/languages/go/libraries/).
3. Open the package documentation and verify its repository, latest version, supported
   OpenTelemetry version, emitted signals, and activation point.
4. Look for configuration that affects cardinality, sensitive attributes, propagation,
   filters, and span names.
5. Check that another agent, wrapper, middleware, or library-owned integration is not
   already instrumenting the same operation.

### Useful Go instrumentation examples

| Technology | Package | How it is activated | What it observes |
| --- | --- | --- | --- |
| HTTP clients and servers | [`otelhttp`](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp) | Wrap an `http.Handler` or `http.Transport` | HTTP client/server traces and metrics; this is the library used in our example |
| gRPC clients and servers | [`otelgrpc`](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc) | Install its client or server gRPC stats handler | RPC traces and metrics |
| AWS SDK for Go v2 | [`otelaws`](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-sdk-go-v2/otelaws) | Append middleware to the SDK configuration | AWS API calls, including service, operation, Region, and request metadata |
| AWS Lambda Go handlers | [`otellambda`](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/github.com/aws/aws-lambda-go/otellambda) | Wrap the Lambda handler | Lambda invocations, trace extraction, and handler execution |
| Redis through `go-redis` | [`redisotel`](https://pkg.go.dev/github.com/redis/go-redis/extra/redisotel/v9) | Add tracing and/or metrics hooks to a Redis client | Redis commands and client activity |
| Go `database/sql` | [`otelsql`](https://pkg.go.dev/github.com/XSAM/otelsql) | Open or register a wrapped database driver | Database operations and optional connection-pool metrics |
| The Go runtime itself | [`runtime`](https://pkg.go.dev/go.opentelemetry.io/contrib/instrumentation/runtime) | Start runtime metric reporting | Memory, garbage collection, goroutines, scheduler, and other Go runtime measurements |

The first four and the runtime package live in OpenTelemetry's Go contrib repository.
`redisotel` is maintained with the Redis Go client, and `otelsql` is a community
integration. "Not in the OpenTelemetry contrib repository" does not automatically mean
bad, but it makes ownership, maintenance, version compatibility, and semantic-convention
support especially important to check.

The activation point is normally a small boundary change rather than an alteration to
business logic. Illustrative examples include:

```go
// gRPC: install an OpenTelemetry stats handler at the server boundary.
server := grpc.NewServer(
    grpc.StatsHandler(otelgrpc.NewServerHandler()),
)

// AWS SDK v2: add OpenTelemetry middleware to this SDK configuration.
otelaws.AppendMiddlewares(&cfg.APIOptions)

// go-redis: add OpenTelemetry hooks to this client.
_ = redisotel.InstrumentTracing(redisClient)
_ = redisotel.InstrumentMetrics(redisClient)
```

These snippets are a recognition guide, not additions to this program. Exact APIs and
error handling should always be checked against the dependency version selected by the
application.

### The important AWS distinction

`otelaws` instruments calls **made by the application through the AWS SDK**. For
example, it can create a span for a `PutItem` request to DynamoDB. It does not monitor
the CPU, memory, disk, or health of the EC2 instance, ECS task, or EKS cluster on which
the application runs.

Other OpenTelemetry components cover those different concerns:

| Component | Purpose | Why it is not the same as an instrumentation library |
| --- | --- | --- |
| AWS EC2, ECS, EKS, and Lambda Resource detectors | Attach cloud/platform identity to telemetry, such as where a service is running | They enrich the Resource; they do not create spans around AWS API calls |
| AWS X-Ray propagator | Reads and writes the AWS X-Ray trace header format | It transports trace context; it does not observe an operation |
| Collector Host Metrics receiver | Collects host CPU, load, memory, disk, filesystem, and network metrics | It runs in the Collector, outside the Go application's SDK |
| Collector AWS receivers | Ingest telemetry from services such as CloudWatch, ECS, or X-Ray | They are Collector pipeline inputs, not wrappers around application code |
| OTLP or vendor exporters | Send completed telemetry to another system | They determine the destination, not which application operations are measured |

For an SRE, these pieces are complementary. A realistic setup might use:

```text
otelaws       -> traces the service's calls to AWS APIs
AWS detector  -> says which AWS environment produced those traces
runtime       -> reports the Go process's runtime health
host receiver -> reports the underlying host's operating-system health
exporter      -> sends all resulting telemetry toward a backend
```

### Evaluation checklist

Before adopting an instrumentation library, answer:

- Does it support the exact library and major version used by the service?
- Is it maintained by OpenTelemetry, by the instrumented library's maintainers, or by a
  third party?
- Does it emit traces, metrics, logs, or only some of them?
- Where is it activated: handler, transport, client hook, driver, middleware, or agent?
- Does it propagate Context across the relevant boundary?
- Which semantic conventions and attribute names does it use?
- Can route names, SQL statements, Redis arguments, AWS parameters, or other sensitive
  values be captured?
- Could any attributes create unbounded cardinality?
- Is the same boundary already instrumented elsewhere?

No package was added to this example for this concept note. The purpose is to learn how
to locate and assess the appropriate integration before changing code.

## Final Chapter 7 concept map

Instrumentation is the part of an observability system that **creates telemetry about
operations**. The SDK pipeline and exporter handle that telemetry after it has been
created.

```text
                         APPLICATION
                              |
          +-------------------+-------------------+
          |                   |                   |
          v                   v                   v
 instrumentation         manual API          zero-code
    libraries          instrumentation      instrumentation
          |                   |                   |
 known library          application-         externally observes
  boundaries            specific work        supported boundaries
          |                   |                   |
          +-------------------+-------------------+
                              |
                    spans and measurements
                              |
                     OpenTelemetry providers
                    /          |           \
             TracerProvider MeterProvider LoggerProvider
                    \          |           /
                              |
                        processors/readers
                              |
                           exporters
                              |
                   stdout, Collector, or backend
```

### 1. SDK initialization prepares the telemetry pipeline

Creating a provider, processor, reader, and exporter gives generated telemetry somewhere
to go. It does not by itself discover operations or create spans:

```text
SDK initialization = handles generated telemetry
instrumentation     = generates telemetry
```

In our services, the tracing and metrics providers define the pipelines. The application
would produce no operation spans if neither instrumentation libraries nor manual spans
created them.

### 2. Instrumentation libraries understand known boundaries

An instrumentation library understands the API and lifecycle of a specific library or
framework. Our `otelhttp` wrappers understand HTTP clients and servers:

```go
instrumentedHandler := otelhttp.NewHandler(mux, "greeting-server")

client := &http.Client{
    Transport: otelhttp.NewTransport(http.DefaultTransport),
}
```

From these small boundary changes, the library can create HTTP spans, record standard
HTTP attributes and metrics, propagate trace context, and connect client and server
operations.

It does not understand the application's business intent. An HTTP library can recognize
a request, but it cannot infer that formatting a greeting is an operation worth tracing.

### 3. Manual instrumentation adds application meaning

Manual instrumentation describes work that a general-purpose library cannot identify:

```go
_, span := otel.Tracer(
    formattingInstrumentationName,
).Start(ctx, "build greeting")
defer span.End()
```

Passing the current Context connects the manual span to the automatically created HTTP
server span. Manual and automatic instrumentation therefore form one trace rather than
two separate tracing systems.

Manual instrumentation is appropriate for meaningful business operations, algorithms,
or unsupported library calls. Instrumenting every small function would usually create
noise and unnecessary telemetry.

### 4. Context connects operations within a process

The active span is stored in a `context.Context`. Instrumented code extracts the parent
from the received Context and returns a new Context containing the newly started span:

```text
incoming request Context
          |
          v
automatic HTTP server span
          |
          v
Context passed to buildGreeting
          |
          v
manual child span
```

If the Context is discarded or replaced with `context.Background()` during the request,
the intended parent-child relationship can be lost.

### 5. Propagation connects different processes

A Go Context cannot cross a network directly. A propagator converts trace context into
transport metadata such as HTTP headers:

```text
greeting service Context
          |
          | inject traceparent header
          v
       HTTP request
          |
          | extract traceparent header
          v
formatting service Context
```

The client instrumentation injects the headers and the server instrumentation extracts
them. Configuring `TraceContext` tells both libraries which wire format to use.

Instrumentation creates spans; propagation connects spans across service boundaries.

### 6. Resources identify the telemetry producer

A Resource supplies attributes about the entity producing telemetry, such as:

```text
service.name=greeting-service
service.version=1.0.0
deployment.environment.name=development
```

The provider associates its Resource with all telemetry it produces. A Resource is not
an instrumentation library and does not create operations. It answers **where did this
telemetry come from?**, while span attributes answer **what happened during this
operation?**

### 7. Instrumentation scope identifies the telemetry-producing code

The name passed to `otel.Tracer(...)` identifies the instrumentation scope:

```go
otel.Tracer("ai-incremental-instrumentation-libraries/formatting-service")
```

This says which instrumentation code created a span. It is distinct from
`service.name`, which identifies the service in which that instrumentation ran.

For automatic HTTP spans, the scope belongs to `otelhttp`. For our manual child span,
the scope belongs to our application instrumentation.

### 8. Zero-code instrumentation changes the activation mechanism

Zero-code instrumentation observes supported application or protocol boundaries without
requiring source changes. Depending on the mechanism, instrumentation may be introduced
by an agent, eBPF, or the build process:

```text
code-based library: application explicitly installs a wrapper or hook
zero-code:          an external or build-time mechanism installs instrumentation
```

Zero-code instrumentation can provide broad technical coverage, but it still cannot
infer arbitrary business meaning. A future Kubernetes experiment will explore its
deployment model, discovery, permissions, metadata, and exported telemetry. It was
deliberately not added to the current services.

### 9. Exporters and Collectors do not instrument the application

An exporter sends completed telemetry somewhere. A Collector can receive, process,
batch, filter, enrich, and forward that telemetry:

```text
instrumentation -> SDK -> exporter -> Collector -> backend
```

Changing an exporter changes the destination. It does not cause a previously invisible
application operation to become instrumented.

### 10. Avoid observing the same boundary twice

Two instrumentation mechanisms can accidentally observe the same operation:

```text
otelhttp wrapper
        +
zero-code HTTP instrumentation
        =
possible duplicate HTTP spans and measurements
```

Before enabling another library or agent, identify every existing activation point.
Nested middleware is not inherently wrong, but two layers representing the same
semantic operation make traces misleading and metrics too large.

### The complete trace from this exercise

```text
Resource: service.name=greeting-service
Scope:    otelhttp
Span:     greeting GET /hello/{name}       SERVER
  |
  | Context within the greeting service
  |
  `-- Scope: otelhttp
      Span: HTTP GET                        CLIENT
        |
        | W3C Trace Context over HTTP
        |
        `-- Resource: service.name=formatting-service
            Scope:    otelhttp
            Span:     GET /format/{name}   SERVER
              |
              | Context within the formatting service
              |
              `-- Scope: application instrumentation
                  Span: build greeting     INTERNAL
```

This trace demonstrates the chapter's central idea: instrumentation libraries provide
reusable visibility at known technical boundaries, propagation joins those boundaries
across services, and manual instrumentation fills in the application-specific meaning
between them.

### Chapter 7 decision guide

```text
Do I need visibility into a known library or framework?
  |
  +-- Yes -> Look for a maintained instrumentation library.
  |
  `-- No
       |
       +-- Is this meaningful application-specific work?
       |     |
       |     `-- Yes -> Add a focused manual span or measurement.
       |
       `-- Do I need broad coverage without changing source?
             |
             `-- Yes -> Evaluate zero-code instrumentation and its
                        supported libraries, deployment requirements,
                        and duplicate-instrumentation risk.
```

The purpose is not to choose one instrumentation style for the entire application.
Healthy instrumentation commonly combines maintained libraries for standard boundaries,
a small amount of manual instrumentation for important application behavior, and
carefully evaluated zero-code coverage where its operational trade-offs are worthwhile.
