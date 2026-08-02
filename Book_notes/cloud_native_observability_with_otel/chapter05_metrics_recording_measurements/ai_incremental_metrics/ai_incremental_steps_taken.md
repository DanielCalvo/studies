# Incremental steps taken

## Step 1: Create the plain names service

### What changed

Created a small Go HTTP service with uninstrumented CRUD operations. Names are
stored in `names.txt`, one name per line.

The available operations are:

- `POST /names` to append the name supplied in the request body.
- `GET /names` to list all names.
- `GET /names/{name}` to check whether a name exists.
- `PUT /names/{name}` to replace a name with the value in the request body.
- `DELETE /names/{name}` to remove a name.

### Concept demonstrated

This is the ordinary application whose behavior we will measure later. Establishing
the application first keeps the web service separate from the OpenTelemetry concepts
we will introduce.

The server registers its routes using a standard Go HTTP multiplexer:

```go
mux.HandleFunc("POST /names", createName)
mux.HandleFunc("GET /names", listNames)
mux.HandleFunc("GET /names/{name}", readName)
mux.HandleFunc("PUT /names/{name}", updateName)
mux.HandleFunc("DELETE /names/{name}", deleteName)
```

### Expected behavior

Running the program starts an HTTP server on `localhost:8080`. Requests create,
read, update, and delete lines in `names.txt`.

### Important distinction

The mutex is application behavior rather than telemetry. Go's HTTP server can handle
multiple requests concurrently, so the mutex prevents overlapping file operations.

### Deliberately not introduced yet

There is no OpenTelemetry configuration, meter, metric instrument, measurement,
reader, exporter, or metrics-related middleware in the program.

## Step 2: Configure the metrics pipeline and obtain a meter

### What changed

Added a console exporter, periodic metric reader, SDK `MeterProvider`, and meter.
The provider is also registered as OpenTelemetry's global meter provider.

### Concept demonstrated

These components form the basic metrics pipeline:

```text
Meter -> MeterProvider -> periodic reader -> console exporter
```

- The exporter determines where collected metrics are sent.
- The periodic reader triggers collection and export every five seconds.
- The `MeterProvider` owns the reader and coordinates the pipeline.
- The meter will create instruments in later steps.

The meter is obtained from the globally registered provider:

```go
provider := sdkmetric.NewMeterProvider(sdkmetric.WithReader(reader))
otel.SetMeterProvider(provider)
meter := otel.Meter(instrumentationName)
```

### Expected behavior

The CRUD service behaves exactly as it did before. Every five seconds, the console
exporter may print resource metadata followed by `"ScopeMetrics": null`. This confirms
that collection is running, but there is no useful application metric because no
instruments or measurements exist yet.

### Important distinction

A meter does not record measurements by itself. It is a factory used to create
instruments such as counters and histograms. This resembles tracing, where a tracer
creates spans but is not itself a span.

### Deliberately not introduced yet

There are no metric instruments, measurements, metric attributes, resources, views,
Prometheus endpoints, or OTLP exporters.

## Step 3: Count successful name creations

### What changed

Used the meter to create an integer counter named `names.created`. The `POST /names`
handler adds one to the counter after it successfully appends a name to the file:

```go
namesCreated, err = meter.Int64Counter(
    "names.created",
    metric.WithDescription("Number of names successfully created"),
)

namesCreated.Add(r.Context(), 1)
```

### Concept demonstrated

A counter records a total that only increases. The meter creates the counter, and the
counter's `Add` method records individual measurements. The SDK aggregates those
measurements before the periodic reader collects them.

The request's context is passed to `Add` because the measurement belongs to the work
performed for that HTTP request.

### Expected behavior

Each successful `POST /names` request increases `names.created` by one. At the next
five-second export, the console output contains the cumulative value:

```text
1 -> 2 -> 3 -> ...
```

### Important distinction

`names.created` is the number of successful create operations since this process started.
It is not the current number of names in the file. Deleting a name therefore does not
decrement this counter.

The counter is incremented only after the file write succeeds, so failed creates are not
counted as successful operations.

### Deliberately not introduced yet

The counter has no attributes. Other CRUD operations, errors, durations, concurrent
requests, and current file state are not measured yet.

## Step 4: Distinguish CRUD operations with an attribute

### What changed

Evolved `names.created` into a counter named `names.operations`. Each successful
handler records a measurement with an `operation` attribute:

```go
nameOperations.Add(
    r.Context(),
    1,
    metric.WithAttributes(attribute.String("operation", "create")),
)
```

The five possible attribute values are `create`, `list`, `read`, `update`, and `delete`.

### Concept demonstrated

Metric attributes add dimensions to a metric. The SDK aggregates measurements with
different attribute sets into separate data points, even though they were recorded
through the same counter:

```text
names.operations{operation="create"} 3
names.operations{operation="read"}   7
names.operations{operation="delete"} 1
```

These series can be viewed separately or added together to obtain the number of all
successful operations.

### Expected behavior

After making different successful CRUD requests, the next console export contains one
`names.operations` metric with a data point for each operation that occurred. Every data
point has its own cumulative value.

### Important distinction

An attribute does not create another instrument. It creates a distinct dimension, and
therefore a distinct metric series, underneath the same instrument.

The `operation` attribute has only five possible values. The actual name is deliberately
not recorded because user-provided names would create an unbounded number of series
and could also expose personal information.

### Deliberately not introduced yet

Only successful operations are counted. A result or status attribute, error counting,
request duration, concurrent requests, resources, and views have not been introduced.

## Step 5: Measure create-request duration with a histogram

### What changed

Created a floating-point histogram named `names.create.duration` with seconds as its
unit:

```go
createDuration, err = meter.Float64Histogram(
    "names.create.duration",
    metric.WithUnit("s"),
    metric.WithDescription("Duration of create-name operations"),
)
```

The create handler captures its start time and records the elapsed duration when the
handler finishes:

```go
start := time.Now()
defer func() {
    createDuration.Record(r.Context(), time.Since(start).Seconds())
}()
```

### Concept demonstrated

A histogram records a distribution of observations. The SDK aggregates individual
durations into a count, sum, minimum, maximum, and bucket counts before the periodic
reader collects them.

Using `defer` ensures that a duration is recorded for every create attempt, including a
request that returns early because reading or writing failed.

### Expected behavior

After one or more `POST /names` requests, the console output contains
`names.create.duration` with fields such as:

```text
Count
Bounds
BucketCounts
Min
Max
Sum
```

The output describes the distribution rather than listing every original duration.

### Important distinction

The counter and histogram answer different questions:

- `names.operations` answers how many successful operations occurred.
- `names.create.duration` describes how create-request durations were distributed.

The histogram records failed and successful create attempts, while the existing counter
still counts only successful operations.

### Deliberately not introduced yet

Only create-request duration is measured. The histogram has no attributes and uses the
SDK's default bucket boundaries. Durations for other operations, custom boundaries,
result attributes, and percentiles have not been introduced.

## Step 6: Track active requests with an up/down counter

### What changed

Created an integer up/down counter named `http.server.active_requests` and wrapped the
router in a small middleware:

```go
activeRequests.Add(r.Context(), 1)
defer activeRequests.Add(r.Context(), -1)

next.ServeHTTP(w, r)
```

Every request passes through this middleware before reaching a CRUD handler.

### Concept demonstrated

An up/down counter records additive changes that may be positive or negative. A request
adds one when it begins and subtracts one when it finishes:

```text
0 -> 1 -> 2 -> 3 -> 2 -> 1 -> 0
```

The resulting value represents the number of requests currently executing rather than a
total accumulated since startup.

### Expected behavior

When no requests are executing, the exported value is zero. If a collection occurs while
requests overlap or one request remains in progress, the exported value is greater than
zero.

### Important distinction

`names.operations` is a monotonic counter and never decreases.
`http.server.active_requests` is a non-monotonic sum produced by an up/down counter, so
it can increase and decrease.

The middleware counts all HTTP requests, including unmatched routes and failed
requests. It is handwritten application middleware; adding an OpenTelemetry instrument
does not automatically instrument the Go HTTP server.

### Deliberately not introduced yet

The active-request instrument has no attributes. We have not introduced automatic HTTP
instrumentation, status/result dimensions, asynchronous instruments, resources, views,
Prometheus, or OTLP.

## Step 7: Observe the current number of stored names

### What changed

Created an observable up/down counter named `names.stored`. Its callback locks and
reads the names file whenever the metric reader performs a collection:

```go
storedNames, err = meter.Int64ObservableUpDownCounter(
    "names.stored",
    metric.WithInt64Callback(func(
        _ context.Context,
        observer metric.Int64Observer,
    ) error {
        namesFileMu.Lock()
        defer namesFileMu.Unlock()

        names, err := loadNames()
        if err != nil {
            return err
        }

        observer.Observe(int64(len(names)))
        return nil
    }),
)
```

### Concept demonstrated

An observable instrument records measurements asynchronously from application
operations. OpenTelemetry invokes its callback when the reader collects metrics.

The callback reports an absolute current value:

```text
File contains 5 names -> Observe(5)
File contains 7 names -> Observe(7)
File contains 2 names -> Observe(2)
```

It does not call `Add(1)` for creates or `Add(-1)` for deletes.

### Expected behavior

Every five-second collection includes `names.stored`, even if no HTTP request happened
during that interval. Its value matches the number of lines currently stored in
`names.txt`.

### Important distinction

Both `http.server.active_requests` and `names.stored` produce non-monotonic sums, but
they receive measurements differently:

- The synchronous up/down counter receives positive and negative changes inline with
  request processing.
- The observable up/down counter invokes a callback during collection and receives the
  complete absolute value.

Here, *asynchronous* does not mean that creating the instrument immediately launches
its callback as an independent background function. Instrument creation only registers
the callback and then the program continues. At each scheduled collection, the periodic
reader invokes the callback as part of that collection, and collection waits for it to
return. The callback can still run concurrently with HTTP handlers, which is why it uses
the file mutex.

Because `names.stored` reads persisted state, it remains correct after an application
restart without reconstructing the value from earlier operations.

The callback reads a small local file for this exercise. Production callbacks should
finish quickly and avoid expensive or unreliable network operations because collection
waits for them.

### Deliberately not introduced yet

The observable counter has no attributes. We have not introduced an observable gauge,
resource metrics, custom views, Prometheus, or OTLP.

## Step 8: Observe a non-additive value with a gauge

### What changed

Created an observable gauge named `names.longest_length`. During collection, its
callback reads the file, finds the longest stored name, and observes its length in Unicode
characters:

```go
longestNameLength, err = meter.Int64ObservableGauge(
    "names.longest_length",
    metric.WithUnit("{character}"),
    metric.WithInt64Callback(func(
        _ context.Context,
        observer metric.Int64Observer,
    ) error {
        // Load names while holding the file mutex.

        longestLength := 0
        for _, name := range names {
            longestLength = max(
                longestLength,
                utf8.RuneCountInString(name),
            )
        }

        observer.Observe(int64(longestLength))
        return nil
    }),
)
```

### Concept demonstrated

An observable gauge reports a current non-additive value when the reader collects
metrics. The callback is invoked on the collection schedule, just like the callback for
`names.stored`.

Two separate pieces produce this behavior:

1. Creating an **observable** instrument with a callback registers how the SDK can
   obtain the instrument's current value. We do not invoke this callback ourselves from
   the CRUD handlers.
2. The **periodic reader** determines when collection happens. In this example its
   interval is five seconds, so each scheduled collection makes the SDK invoke the
   registered observable callbacks and then export their observations.

The callback has no timer and does not run continuously. Its normal execution sequence
is:

```text
periodic reader starts collection
    -> SDK invokes observable callback
    -> callback observes current value
    -> reader exports the collected metric
```

The periodic reader is also useful without observable instruments. For synchronous
instruments, handlers call methods such as `Add` or `Record`, the SDK aggregates those
measurements, and the reader periodically collects and exports the aggregates. Observable
instruments differ only in that collection also causes the SDK to invoke their callbacks.

Unicode characters are counted rather than raw bytes so names containing multibyte
characters have a human-meaningful length.

### Expected behavior

For a file containing:

```text
Ana
Daniel
Beatrice
```

the next collection reports:

```text
names.stored = 3
names.longest_length = 8
```

### Important distinction

The two observable instruments describe different aggregation semantics:

- `names.stored` is additive. Counts from independent stores can be summed.
- `names.longest_length` is non-additive. Adding the longest lengths from two instances
  produces no meaningful value; a query would normally take their maximum.

The exported gauge contains the last value observed during that collection. It is not
exported as a monotonic or non-monotonic sum.

### Deliberately not introduced yet

The gauge has no attributes. We have not introduced resources, views, customized
histogram boundaries, an observable monotonic counter, Prometheus, or OTLP.

## Step 9: Customize histogram buckets with a View

### What changed

Created a View that matches `names.create.duration` and gives its explicit-bucket
histogram boundaries suited to a fast local operation:

```go
createDurationView := sdkmetric.NewView(
    sdkmetric.Instrument{Name: "names.create.duration"},
    sdkmetric.Stream{
        Aggregation: sdkmetric.AggregationExplicitBucketHistogram{
            Boundaries: []float64{
                0.001, 0.005, 0.010, 0.025, 0.050,
                0.100, 0.250, 0.500, 1.000,
            },
        },
    },
)

provider := sdkmetric.NewMeterProvider(
    sdkmetric.WithReader(reader),
    sdkmetric.WithView(createDurationView),
)
```

The values are expressed in seconds because `names.create.duration` uses `s` as its
unit. For example, `0.005` seconds is 5 milliseconds.

### Concept demonstrated

A View changes how the SDK turns measurements from matching instruments into metric
streams. The handler still records exactly the same duration:

```go
createDuration.Record(r.Context(), time.Since(start).Seconds())
```

The View changes the SDK aggregation configuration without requiring any change at that
recording site.

### Expected behavior

The exported `names.create.duration` histogram now has boundaries at 1 ms, 5 ms, 10 ms,
25 ms, 50 ms, 100 ms, 250 ms, 500 ms, and 1 second. Each recorded duration contributes
to the corresponding bucket, while the histogram still reports its count and sum.

### Important distinction

The histogram instrument records individual duration measurements. The View neither
records measurements nor schedules collection; it selects how the SDK aggregates the
measurements before the reader collects them.

This View matches one instrument by name. Views can also rename streams, filter
attributes, select other aggregation types, or match a wider set of instruments, but
those capabilities are not used in this step.

### Deliberately not introduced yet

We have not added resources, attribute filtering through Views, an observable monotonic
counter, Prometheus, or OTLP.

## Step 10: Identify the service with a Resource

### What changed

Created a Resource containing the service name and version, merged it with
OpenTelemetry's default resource, and attached it to the `MeterProvider`:

```go
serviceResource, err := resource.Merge(
    resource.Default(),
    resource.NewWithAttributes(
        semconv.SchemaURL,
        semconv.ServiceName("names-service"),
        semconv.ServiceVersion("1.0.0"),
    ),
)

provider := sdkmetric.NewMeterProvider(
    sdkmetric.WithReader(reader),
    sdkmetric.WithView(createDurationView),
    sdkmetric.WithResource(serviceResource),
)
```

Merging preserves default metadata such as the OpenTelemetry SDK name, language, and
version while replacing the default `unknown_service` identity with our explicit one.

### Concept demonstrated

A Resource describes the entity producing telemetry. Because it belongs to the
`MeterProvider`, its attributes are associated with every metric produced through that
provider.

### Expected behavior

Each exported metrics collection now contains resource attributes including:

```text
service.name = names-service
service.version = 1.0.0
```

The service identity appears even when no CRUD request has occurred because it describes
the producer, not an individual measurement.

### Important distinction

Resource attributes and metric attributes have different scopes:

- `service.name=names-service` identifies the entity producing all the metrics.
- `operation=create` divides `names.operations` into a particular metric series.

In OpenTelemetry they remain separate kinds of metadata. A Prometheus translation may
map `service.name` to `job`, place other resource attributes on `target_info`, or promote
selected resource attributes to labels according to its configuration.

### Deliberately not introduced yet

We have not added automatic infrastructure resource detectors, attribute filtering
through Views, an observable monotonic counter, Prometheus, or OTLP.

## Step 11: Expose metrics for Prometheus to scrape

### What changed

Replaced the standard-output exporter and `PeriodicReader` with OpenTelemetry's
Prometheus exporter:

```go
prometheusReader, err := otelprometheus.New()

provider := sdkmetric.NewMeterProvider(
    sdkmetric.WithReader(prometheusReader),
    sdkmetric.WithView(createDurationView),
    sdkmetric.WithResource(serviceResource),
)
```

The Prometheus exporter implements the OpenTelemetry metric `Reader` interface and
registers a collector with Prometheus's default registry. A Prometheus HTTP handler now
serves that registry:

```go
mux.Handle("GET /metrics", promhttp.Handler())
```

### Concept demonstrated

Prometheus uses pull-based collection. The application no longer collects and prints
metrics every five seconds. Instead, each request to `/metrics` initiates collection and
returns the result in Prometheus exposition format:

```text
Prometheus or curl requests /metrics
    -> Prometheus handler gathers registered collectors
    -> OpenTelemetry Prometheus reader collects metrics
    -> SDK invokes observable callbacks
    -> handler returns Prometheus text
```

Synchronous measurements are still recorded when handlers call `Add` and `Record`.
Their SDK aggregates are retained until a scrape reads them. Observable callbacks for
`names.stored` and `names.longest_length` now run at scrape time rather than on a
five-second timer.

### Expected behavior

After starting the program, metrics can be inspected directly:

```bash
curl http://localhost:8080/metrics
```

The response includes Prometheus series such as the operations counter, create-duration
histogram, active requests, stored-name count, and longest-name length. OpenTelemetry
metric names and units are translated to Prometheus conventions where appropriate.

### Important distinction

The OpenTelemetry Prometheus exporter and the HTTP handler have related but different
jobs:

- The exporter is a metric reader and Prometheus collector that bridges the
  OpenTelemetry SDK's aggregated data into Prometheus metric families.
- `promhttp.Handler()` exposes registered Prometheus collectors over HTTP.

There is no `PeriodicReader` in this version. The scrape request itself determines when
collection occurs.

Because the current active-request middleware wraps the entire server, the `/metrics`
scrape is itself an active request and may make `http_server_active_requests` report
`1` while that scrape is being served.

### Deliberately not introduced yet

We have not installed or configured a Prometheus server. The endpoint can be inspected
with `curl`, but no external system is scraping or storing it. We have also not added
OTLP, a Collector, automatic HTTP instrumentation, or an observable monotonic counter.

## Step 12: Identify the instrumentation scope

### What changed

Added a version and schema URL when obtaining the Meter:

```go
const instrumentationVersion = "1.0.0"

meter := otel.Meter(
    instrumentationName,
    metric.WithInstrumentationVersion(instrumentationVersion),
    metric.WithSchemaURL(semconv.SchemaURL),
)
```

The Meter's existing name, the new version, and the schema URL together describe the
instrumentation scope that creates our instruments.

### Concept demonstrated

An instrumentation scope identifies the application code or library responsible for
creating telemetry. This lets a service emit metrics from multiple instrumentation
libraries while retaining the origin of each metric.

The schema URL identifies the OpenTelemetry schema used to interpret standardized
telemetry names and attributes. This example uses the schema URL supplied by the
semantic-conventions package imported by the program.

### Expected behavior

The Prometheus output changes from empty scope metadata:

```text
otel_scope_name="ai-incremental-metrics"
otel_scope_version=""
otel_scope_schema_url=""
```

to values resembling:

```text
otel_scope_name="ai-incremental-metrics"
otel_scope_version="1.0.0"
otel_scope_schema_url="https://opentelemetry.io/schemas/1.41.0"
```

### Important distinction

Although both versions are `1.0.0` in this study program, they describe different things:

- Resource attribute `service.version` is the version of the running service.
- Instrumentation-scope version is the version of the code or library that created the
  metric instruments.

Those versions can evolve independently in a production application.

### Deliberately not introduced yet

We have not added instrumentation-scope attributes, a second Meter from another
instrumentation library, an observable monotonic counter, additional View filtering,
OTLP, or a Collector.

## Step 13: Observe an externally maintained monotonic total

### What changed

Created an observable counter that reads the number of major page faults accumulated by
the process:

```go
majorPageFaults, err = meter.Int64ObservableCounter(
    "process.major_page_faults",
    metric.WithUnit("{fault}"),
    metric.WithInt64Callback(func(
        _ context.Context,
        observer metric.Int64Observer,
    ) error {
        var usage syscall.Rusage
        if err := syscall.Getrusage(
            syscall.RUSAGE_SELF,
            &usage,
        ); err != nil {
            return err
        }

        observer.Observe(usage.Majflt)
        return nil
    }),
)
```

The operating system maintains this cumulative process value. Application handlers do
not update it and never call `Add`.

### Concept demonstrated

An observable counter reports a current absolute total during collection. It is suitable
when another system already owns a cumulative, non-decreasing value or when obtaining
that value inline with application work would be inappropriate.

In this program, requesting `/metrics` causes the Prometheus reader to invoke the
callback, which reads the latest total from the operating system.

### Expected behavior

The Prometheus endpoint includes a monotonic counter resembling:

```text
# TYPE process_major_page_faults_total counter
process_major_page_faults_total{
    otel_scope_name="ai-incremental-metrics"
} 0
```

The exact value depends on the operating system's activity. Across scrapes in the same
process, it may remain unchanged or increase, but it must not decrease.

### Important distinction

Both ordinary and observable counters represent monotonic totals, but measurements
enter the SDK differently:

- With an ordinary counter, application code reports each change by calling `Add`.
- With an observable counter, the reader invokes a callback and the callback reports the
  complete cumulative value.

This differs from `names.stored`, whose observable up/down-counter value can decrease,
and `names.longest_length`, whose gauge value is a non-additive snapshot.

### Deliberately not introduced yet

This process measurement uses Go's `syscall` API and is platform-specific rather than
portable runtime instrumentation. We have not added duplicate instruments, further View
filtering, OTLP, or a Collector.

## Step 14: Filter dimensions into a rolled-up metric stream

### What changed

Added two Views matching `names.operations`. One preserves the existing stream and its
`operation` attribute, while the other removes all metric attributes and renames the
result:

```go
operationCriteria := sdkmetric.Instrument{
    Name: "names.operations",
}

operationDetailView := sdkmetric.NewView(
    operationCriteria,
    sdkmetric.Stream{},
)

allOperationsView := sdkmetric.NewView(
    operationCriteria,
    sdkmetric.Stream{
        Name: "names.operations.all",
        AttributeFilter: attribute.NewAllowKeysFilter(),
    },
)
```

Both Views are registered with the provider:

```go
sdkmetric.WithView(
    createDurationView,
    operationDetailView,
    allOperationsView,
)
```

The empty allow-list intentionally rejects every metric attribute from the rolled-up
stream.

### Concept demonstrated

A View can select which dimensions appear in a metric stream. Removing the `operation`
attribute does more than hide its label: measurements that previously belonged to
separate attribute sets are aggregated together.

The existing handler calls remain unchanged:

```go
nameOperations.Add(
    r.Context(),
    1,
    metric.WithAttributes(
        attribute.String("operation", "create"),
    ),
)
```

That one measurement feeds both matching streams inside the SDK.

### Expected behavior

After two creates, one read, and one delete, the detailed stream contains separate
Prometheus series:

```text
names_operations_total{operation="create"} 2
names_operations_total{operation="read"} 1
names_operations_total{operation="delete"} 1
```

The additional stream has no `operation` label and contains their combined total:

```text
names_operations_all_total 4
```

### Important distinction

Metric attributes define dimensions and therefore distinct series. Filtering a dimension
reduces the number of series and also reduces the questions the data can answer:

- The detailed stream can answer how often each operation occurred.
- The rolled-up stream can answer only how many operations occurred altogether.

In the Go SDK, any explicit View matching an instrument suppresses its implicit default
View. We therefore added the explicit pass-through `operationDetailView` to retain the
original detailed stream alongside the filtered one.

### Deliberately not introduced yet

The rolled-up stream removes every metric attribute. We have not demonstrated a View
that retains a selected subset from several attributes, dropped an entire instrument, or
introduced duplicate instruments, OTLP, or a Collector.

## Step 15: Request the same logical instrument from two places

### What changed

Requested a second counter handle using a registration identical to
`nameOperations`:

```go
nameOperations, err = meter.Int64Counter(
    "names.operations",
    metric.WithDescription(
        "Number of successful name operations",
    ),
)

readNameOperations, err = meter.Int64Counter(
    "names.operations",
    metric.WithDescription(
        "Number of successful name operations",
    ),
)
```

The read handler now records through the second handle:

```go
readNameOperations.Add(
    r.Context(),
    1,
    metric.WithAttributes(
        attribute.String("operation", "read"),
    ),
)
```

Other handlers continue using `nameOperations`.

### Concept demonstrated

Separate components can independently request the same logical instrument. When the
Meter, name, instrument type, unit, and description are identical, the SDK recognizes an
identical registration and directs measurements from both handles into the same
aggregations.

This can be useful when components are initialized separately and sharing one Go
variable would create unnecessary coupling.

### Expected behavior

A successful read still produces only the existing detailed series:

```text
names_operations_total{operation="read"} 1
```

It also contributes once to the rolled-up stream introduced in Step 14:

```text
names_operations_all_total 1
```

There is no second Prometheus metric merely because a second instrument handle exists.

### Important distinction

This step demonstrates an **identical duplicate registration**, which the SDK can safely
combine. It does not demonstrate a conflicting registration.

Reusing `names.operations` with a different instrument type, unit, or incompatible
description would create ambiguous metric semantics and may produce an OpenTelemetry
diagnostic. Such conflicts should be corrected rather than used as an application design.

Attributes remain the mechanism for distinguishing meaningful categories such as
`operation="read"` within the shared logical metric.

The duplicate is not the Go variable appearing in two places. The global statement:

```go
var readNameOperations metric.Int64Counter
```

declares the variable once. The statement in `run` uses `=`, not `:=`, so it assigns an
instrument handle to that existing variable. The duplicate registration comes from
calling `meter.Int64Counter` twice with the same instrument definition. Calling
`readNameOperations.Add` then records through the second handle; it does not redeclare
the variable.

This is more useful across independent components than in our small program. For
example, a `names` package and an `audit` package might each receive the same Meter
during initialization and independently request an identical `names.operations`
counter. They do not need access to a shared global Go variable: the SDK recognizes the
identical registrations and safely combines both packages' measurements into one logical
metric instead of creating conflicting metrics.

### Deliberately not introduced yet

Both handles are shown in one file to keep the example easy to read. We have not split
the application into packages, created an intentional conflicting registration, dropped
an instrument with a View, or added OTLP or a Collector.

## Step 16: Final Chapter 5 concept map

### What changed

No program behavior changed. This final study step maps the instruments and metric
streams in the completed example to the concepts they demonstrate.

### Instrument concept map

| Metric or stream | Instrument | Measurement style | Monotonic | Additive | Operational question |
| --- | --- | --- | --- | --- | --- |
| `names.operations` | Counter | Synchronous `Add(1)` after a successful operation | Yes | Yes | How many successful operations occurred for each operation type? |
| `names.operations.all` | Counter stream produced by a View | Receives the same synchronous measurements after removing attributes | Yes | Yes | How many successful operations occurred altogether? |
| `names.create.duration` | Histogram | Synchronous `Record(duration)` when the create handler finishes | Not applicable | Yes | What is the distribution of create-request latency? |
| `http.server.active_requests` | Up/down counter | Synchronous `Add(1)` on entry and `Add(-1)` on exit | No | Yes | How many HTTP requests are executing now? |
| `names.stored` | Observable up/down counter | Callback reports the absolute count during a scrape | No | Yes | How many names are currently stored? |
| `names.longest_length` | Observable gauge | Callback reports the current longest length during a scrape | No | No | What is the longest current name? |
| `process.major_page_faults` | Observable counter | Callback reports the OS-maintained cumulative total during a scrape | Yes | Yes | How many major page faults has this process accumulated? |

Here, **monotonic** means that the value cannot decrease during the lifetime of one
process and time series. A counter can restart from zero when the process restarts.

**Additive** means that values from independent sources can be meaningfully summed. For
example, operation counts from two service instances can be added. The longest name
length is non-additive because adding two maximum lengths has no useful meaning.

### Pipeline concept map

```text
application handlers
    -> call Add or Record on synchronous instruments
    -> SDK aggregates measurements

Prometheus requests /metrics
    -> Prometheus reader starts collection
    -> SDK invokes observable callbacks
    -> Views shape aggregations, dimensions, and streams
    -> Prometheus handler returns exposition text
```

The main components have these responsibilities:

| Component | Responsibility in this program |
| --- | --- |
| `MeterProvider` | Owns the metrics pipeline and connects its Reader, Views, and Resource. |
| `Meter` | Identifies the instrumentation scope and creates instruments. |
| Instrument | Accepts or observes measurements with defined metric semantics. |
| View | Changes SDK aggregation or dimensions without changing recording calls. |
| Prometheus exporter/Reader | Collects SDK metrics when Prometheus scrapes. |
| `promhttp.Handler()` | Serves the Prometheus registry at `/metrics`. |
| Resource | Identifies the service producing all metrics. |
| Instrumentation scope | Identifies the code or library that created the instruments. |

### Important distinctions

- Synchronous instruments are called by application code; observable callbacks are
  invoked by the Reader during collection.
- An instrument defines measurement semantics; a View determines how matching
  measurements become output streams.
- Metric attributes create series dimensions; Resource attributes describe the
  telemetry-producing entity.
- A Reader controls collection; an exporter or HTTP handler controls how collected
  metrics leave the process.
- Identical instrument registrations can share one logical aggregation, but ordinary
  code should share one handle when that is simple.

### Expected behavior

The final program exposes its metrics at:

```bash
curl http://localhost:8080/metrics
```

Synchronous metrics reflect calls made by the CRUD handlers. Observable values are read
when this endpoint is scraped. The detailed and rolled-up operation streams show how a
View can produce different aggregations from the same measurements.

### Deliberately not introduced

This chapter example stops at local Prometheus exposition. It does not configure a
Prometheus server, OTLP export, an OpenTelemetry Collector, long-term storage,
dashboards, alerts, automatic HTTP instrumentation, or exemplar-based trace correlation.
Those are integration and production-observability topics that can be studied separately
without changing the instrument fundamentals summarized here.
