# Incremental steps taken

This file records each implemented step of the local Chapter 8 Collector
exercise.

## Step 1: Create the smallest complete trace pipeline

### What changed

Added a local Docker Compose exercise containing:

- An OpenTelemetry Collector Contrib container
- A Collector configuration file
- A one-shot `telemetrygen` test container
- Instructions for running and inspecting the exercise

The Collector and test generator are pinned to version `0.153.0` so the
exercise does not change silently when newer images are released.

### Concept demonstrated

A Collector signal pipeline needs an input and an output:

```text
receiver -> pipeline -> exporter
```

The OTLP receiver accepts trace data over gRPC. The traces pipeline connects
that receiver to the debug exporter. The debug exporter prints the resulting
telemetry to the Collector's own output.

The `service.pipelines` entry is essential. Declaring a receiver or exporter in
its top-level configuration section makes it available for use, but does not
enable it.

### Relevant configuration

```yaml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317

exporters:
  debug:
    verbosity: detailed

service:
  pipelines:
    traces:
      receivers: [otlp]
      exporters: [debug]
```

The book calls the console exporter `logging`. In current Collector versions,
the equivalent learning and troubleshooting component is the `debug` exporter.

### Expected behavior

After the Collector starts, running the trace generator sends one trace to the
OTLP/gRPC receiver. The Collector's output then contains detailed trace
information printed by the debug exporter.

The generated trace may contain multiple spans. `--traces=1` means one trace,
not necessarily one span.

### Important distinctions

- The receiver accepts telemetry; it does not create or store the trace.
- The traces pipeline wires components together; it is not a backend.
- The debug exporter prints telemetry; it does not persist it for later
  querying.
- `telemetrygen` is a test data source outside the Collector pipeline.
- OTLP is the telemetry protocol, while gRPC is the transport used in this
  step.

### Deliberately not introduced yet

- OTLP over HTTP
- Metrics and logs pipelines
- An instrumented application
- Processors such as batch, attributes, or memory limiter
- File or network backend exporters
- Extensions
- Alloy or Kubernetes

## Step 2: Accept OTLP over gRPC and HTTP

### What changed

The existing OTLP receiver now listens using both standard network transports:

- OTLP over gRPC on port `4317`
- OTLP over HTTP on port `4318`

Docker Compose now publishes both ports on the host's loopback interface. A
second one-shot test generator sends a trace using HTTP, while the original
generator continues to use gRPC.

### Concept demonstrated

OTLP and its transport are related but distinct:

```text
OTLP
    Defines how OpenTelemetry data is represented and exchanged.

gRPC or HTTP
    Provides a network transport that carries the OTLP data.
```

The receiver can expose both transports without creating separate trace
pipelines:

```text
OTLP/gRPC :4317 --\
                   >-- OTLP receiver -> traces pipeline -> debug exporter
OTLP/HTTP :4318 --/
```

### Relevant configuration

```yaml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317
      http:
        endpoint: 0.0.0.0:4318
```

The HTTP test generator adds the transport-selection flag:

```yaml
command:
  - traces
  - --otlp-endpoint=collector:4318
  - --otlp-insecure
  - --otlp-http
  - --traces=1
```

### Expected behavior

Running each generator produces a separate trace entry in the Collector output:

```bash
docker compose run --rm trace-generator
docker compose run --rm trace-generator-http
docker compose logs collector
```

Both traces contain equivalent resource, scope, and span structures. The
transport changes how the data reaches the Collector, not the meaning of the
trace.

### Important distinctions

- Ports `4317` and `4318` are conventions for OTLP/gRPC and OTLP/HTTP,
  respectively; they are not different telemetry signals.
- Both protocols are configured under the same named `otlp` receiver.
- Both protocols feed the same traces pipeline and debug exporter.
- `--otlp-insecure` disables transport security for this local exercise. It is
  not a recommendation for a production connection.

### Deliberately not introduced yet

- Metrics and logs pipelines
- A real instrumented application
- Processors
- Authentication or TLS
- A telemetry backend
- Alloy or Kubernetes

## Step 3: Send spans from a real Go application

### What changed

Copied the Chapter 4 random-character tracing program into a separate
`trace-sender` module under this exercise. The original Chapter 4 program
remains unchanged.

The copied application's stdout trace exporter was replaced with an OTLP/gRPC
trace exporter pointed at the Collector on `localhost:4317`. Its tracing
structure, context propagation, resource, span processor, and four spans remain
the same.

### Concept demonstrated

An application's trace exporter determines where completed spans leave the
process:

```text
Before:

Go SDK -> stdout exporter -> application terminal

After:

Go SDK -> OTLP/gRPC exporter -> Collector -> debug exporter
```

The application now knows the Collector endpoint and OTLP transport, but it does
not know or care that the Collector currently prints spans through a debug
exporter. The Collector operator can later change the destination without
changing the application's instrumentation.

### Relevant code

```go
exporter, err := otlptracegrpc.New(
    ctx,
    otlptracegrpc.WithEndpoint("localhost:4317"),
    otlptracegrpc.WithInsecure(),
)
```

The new exporter is still passed to the same application-side batch span
processor:

```go
processor := sdktrace.NewBatchSpanProcessor(exporter)
```

This application processor batches spans before sending OTLP. It is distinct
from a Collector processor, which would run after the Collector receives them.

### Expected behavior

With the Collector running:

```bash
cd trace-sender
go run .
```

The application terminal contains only ordinary program output:

```text
Original: <ten characters>
Reversed: <the characters reversed>
Saved the reversed characters to reversed-random-characters.txt
```

The Collector terminal contains one trace with four spans:

- `process random characters`
- `read random characters`
- `reverse characters`
- `save characters`

### Important distinctions

- Instrumentation creates the same spans as before; only their export
  destination changed.
- The application-side OTLP exporter sends telemetry but does not store it.
- The Collector's OTLP receiver accepts the telemetry.
- The Collector's debug exporter prints it.
- The application batch span processor and a future Collector batch processor
  run in different processes and at different stages.
- `WithInsecure` is used only for this localhost learning connection.

### Deliberately not introduced yet

- Metrics and logs pipelines
- Collector processors
- Environment-based endpoint configuration
- TLS or authentication
- A telemetry backend
- Alloy or Kubernetes

## Step 4: Add an independent metrics pipeline

### What changed

Added a metrics pipeline that reuses the configured OTLP receiver and debug
exporter. Also added a one-shot `metrics-generator` Compose service that sends
one metric over OTLP/gRPC.

The existing trace pipeline, trace generators, and Go trace sender were not
changed.

### Concept demonstrated

Collector pipelines are specific to a telemetry signal:

```text
                                      |-- traces pipeline  --\
OTLP receiver ------------------------|                       |--> debug exporter
                                      `-- metrics pipeline --/
```

The same named receiver and exporter can participate in more than one pipeline.
The pipelines remain independent: traces entering the receiver go through the
traces pipeline, while metrics go through the metrics pipeline.

Adding a metrics pipeline does not derive metrics from traces or turn a trace
into another signal. It only provides a configured route for metrics that
arrive at the Collector.

### Relevant configuration

```yaml
service:
  pipelines:
    traces:
      receivers: [otlp]
      exporters: [debug]

    metrics:
      receivers: [otlp]
      exporters: [debug]
```

The one-shot test source uses the metrics mode of `telemetrygen`:

```yaml
command:
  - metrics
  - --otlp-endpoint=collector:4317
  - --otlp-insecure
  - --metrics=1
```

### Expected behavior

With the Collector running:

```bash
docker compose run --rm metrics-generator
docker compose logs collector
```

The debug exporter prints a `Metrics` entry containing one generated gauge data
point. Running either trace generator or the Go trace sender still produces a
separate `Traces` entry.

### Important distinctions

- A receiver describes how data enters; a pipeline selects which signal flows
  through which configured components.
- Reusing `otlp` does not create a second network listener.
- Reusing `debug` does not combine traces and metrics into one signal.
- A metric data point is not a span, even though both can travel over OTLP.
- The generator supplies the metric; the Collector does not create it.

### Deliberately not introduced yet

- A logs pipeline
- Metrics emitted by the Go application
- Collector processors
- A telemetry backend
- Alloy or Kubernetes

## Step 5: Add an independent logs pipeline

### What changed

Added a logs pipeline that reuses the configured OTLP receiver and debug
exporter. Also added a one-shot `logs-generator` Compose service that sends one
OpenTelemetry log record over OTLP/gRPC.

The existing trace and metrics pipelines remain unchanged.

### Concept demonstrated

The Collector routes each OpenTelemetry signal through a pipeline of the
matching type:

```text
                    |-- traces pipeline  --\
OTLP receiver ------|-- metrics pipeline ---> debug exporter
                    `-- logs pipeline    --/
```

All three pipelines can refer to the same named `otlp` receiver and `debug`
exporter configuration. They do not thereby become one combined signal.

### Relevant configuration

```yaml
service:
  pipelines:
    traces:
      receivers: [otlp]
      exporters: [debug]

    metrics:
      receivers: [otlp]
      exporters: [debug]

    logs:
      receivers: [otlp]
      exporters: [debug]
```

The log test source sends one record with a recognizable body:

```yaml
command:
  - logs
  - --otlp-endpoint=collector:4317
  - --otlp-insecure
  - --logs=1
  - --body=chapter 8 test log
```

### Expected behavior

With the Collector running:

```bash
docker compose run --rm logs-generator
docker compose logs collector
```

The debug exporter prints a `Logs` entry containing one log record whose body is
`chapter 8 test log`.

The other sources continue to produce their respective entries:

```text
trace source       -> Traces
metrics generator  -> Metrics
logs generator     -> Logs
```

### Important distinctions

- An OpenTelemetry log record is its own signal, not a span event.
- This generator pushes an OTLP log record directly to the receiver.
- The Collector is not tailing the generator's standard output.
- Enabling a logs pipeline does not automatically capture application or
  container output.
- The debug exporter can print all three signals while keeping their data
  models distinct.

### Deliberately not introduced yet

- Logs emitted by the Go application
- File or container log tailing
- Collector processors
- A telemetry backend
- Alloy or Kubernetes

## Step 6: Batch telemetry inside the Collector

### What changed

Defined one named batch processor with a five-second timeout and enabled it in
the traces, metrics, and logs pipelines.

Each pipeline now follows this shape:

```text
receiver -> batch processor -> exporter
```

The deliberately long timeout makes the behavior visible. It is a learning
setting, not a recommended production value.

### Concept demonstrated

A processor receives telemetry after the receiver and before the exporter:

```yaml
processors:
  batch:
    timeout: 5s
```

Defining the processor only makes it available. Referencing it in a pipeline is
what enables it:

```yaml
service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [debug]
```

The batch processor groups telemetry of the same signal so an exporter can
handle fewer, larger export operations.

### Expected behavior

Send several traces within the five-second window:

```bash
docker compose run --rm trace-generator
docker compose run --rm trace-generator
docker compose run --rm trace-generator
```

The debug exporter does not print each trace immediately. At the next periodic
flush, it can print one trace export containing the spans accumulated during
that window. The five-second timeout is a maximum interval, not a separate
five-second delay measured from each item's arrival.

Metrics and logs experience the same delay in their respective pipelines.

### Important distinctions

- One batch configuration referenced by multiple pipelines produces a separate
  processor instance for each pipeline.
- Traces, metrics, and logs are never combined into one signal batch.
- Batching changes export timing and grouping; it does not alter telemetry
  meaning.
- The processor order written in a pipeline is the processing order.
- The Go application's batch span processor batches spans before OTLP export.
- The Collector batch processor batches telemetry after OTLP reception.

The complete trace route is now:

```text
application spans
    -> application batch span processor
    -> OTLP exporter
    -> Collector OTLP receiver
    -> Collector batch processor
    -> Collector debug exporter
```

### Deliberately not introduced yet

- Attribute or resource modification
- Span-name transformation
- Filtering
- Memory limiting
- A telemetry backend
- Alloy or Kubernetes

## Step 7: Enrich spans with a Collector attribute

### What changed

Defined an attributes processor instance named
`attributes/add-study-chapter`. It upserts this attribute:

```text
study.chapter = "8"
```

The processor is enabled only in the traces pipeline and runs before the batch
processor:

```text
OTLP receiver -> attributes processor -> batch processor -> debug exporter
```

The Go trace sender was not changed.

### Concept demonstrated

The Collector can centrally enrich telemetry after applications send it:

```yaml
processors:
  attributes/add-study-chapter:
    actions:
      - key: study.chapter
        value: "8"
        action: upsert
```

The component identifier has two parts:

```text
attributes/add-study-chapter
^^^^^^^^^^ ^^^^^^^^^^^^^^^^^
type       instance name
```

The suffix lets multiple independently configured instances of the same
processor type coexist.

### Relevant pipeline configuration

```yaml
service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [attributes/add-study-chapter, batch]
      exporters: [debug]
```

Processor list order is execution order. Each span is enriched before the batch
processor groups spans for export.

### Expected behavior

Run the unchanged Go application while the Collector is running. The debug
exporter prints its root span and three child spans, and each contains:

```text
study.chapter: Str(8)
```

The application's own output and source-level span attributes remain otherwise
unchanged.

### Important distinctions

- The application creates the spans; the Collector adds this attribute later.
- `upsert` inserts a missing key or replaces the existing value.
- Enabling this processor only in `traces` leaves metrics and logs unchanged.
- This is a span attribute, not a Resource attribute describing the producing
  service.
- Batching affects grouping and timing; it does not cause the enrichment.

### Deliberately not introduced yet

- Resource attribute modification
- Span-name transformation
- Filtering
- Memory limiting
- A telemetry backend
- Alloy or Kubernetes

## Step 8: Enrich every signal's Resource

### What changed

Defined a Resource processor instance named
`resource/add-local-environment`. It upserts this Resource attribute:

```text
deployment.environment.name = "development"
```

The processor runs near the beginning of the traces, metrics, and logs
pipelines. The existing trace-only attributes processor remains unchanged.

### Concept demonstrated

Resource attributes describe the entity or environment producing telemetry:

```yaml
processors:
  resource/add-local-environment:
    attributes:
      - key: deployment.environment.name
        value: development
        action: upsert
```

This differs from the previous span enrichment:

```text
study.chapter
    Attached to each individual span

deployment.environment.name
    Attached to the Resource shared by a group of telemetry records
```

The book uses `deployment.environment`. That key is now deprecated and has been
replaced by `deployment.environment.name`.

### Relevant pipeline configuration

```yaml
service:
  pipelines:
    traces:
      processors:
        - resource/add-local-environment
        - attributes/add-study-chapter
        - batch

    metrics:
      processors: [resource/add-local-environment, batch]

    logs:
      processors: [resource/add-local-environment, batch]
```

The Resource processor is referenced by all three pipelines so their
environment metadata remains consistent.

### Expected behavior

The debug exporter displays this under `Resource attributes` for traces,
metrics, and logs:

```text
deployment.environment.name: Str(development)
```

The Go application's trace output continues to show this separately inside each
span:

```text
study.chapter: Str(8)
```

### Important distinctions

- A Resource attribute is associated with the telemetry producer, not one
  operation or record.
- Resource enrichment does not modify the Go application's instrumentation.
- Referencing one Resource processor configuration in multiple pipelines gives
  each pipeline its own processor instance.
- The Resource processor does not convert, combine, or correlate signals.
- `upsert` creates the Resource attribute or replaces its existing value.

### Deliberately not introduced yet

- Span-name transformation
- Filtering
- Memory limiting
- Resource detection from the runtime environment
- A telemetry backend
- Alloy or Kubernetes

## Step 9: Transform span names with OTTL

### What changed

Added a trace transform processor that prefixes each span name with the value
of its `study.chapter` attribute:

```yaml
transform/prefix-span-name:
  trace_statements:
    - set(span.name, Concat([span.attributes["study.chapter"], span.name], ":"))
```

For example, `read random characters` becomes
`8:read random characters`.

### Concept demonstrated

The transform processor uses the OpenTelemetry Transformation Language (OTTL)
to modify telemetry as it passes through the Collector. `set` replaces the
span name, while `Concat` joins the chapter attribute and original name with a
colon.

This is more flexible than an attributes processor: the attributes processor
performs predefined operations on attributes, whereas OTTL can read fields and
use functions to construct a new value for another field.

### Relevant pipeline configuration

```yaml
processors:
  - resource/add-local-environment
  - attributes/add-study-chapter
  - transform/prefix-span-name
  - batch
```

The order matters. `attributes/add-study-chapter` must run before
`transform/prefix-span-name` because the transform reads the attribute created
by the preceding processor. A processor sees the results of every processor
that ran before it.

### Expected behavior

The debug exporter should display the Go application's four transformed span
names:

```text
8:read random characters
8:reverse characters
8:save characters
8:process random characters
```

The existing `study.chapter: Str(8)` span attribute and
`deployment.environment.name: Str(development)` Resource attribute remain
present.

### Important distinctions

- OTTL is evaluated inside the Collector; the sending Go application is not
  changed.
- The transform processor changes existing telemetry rather than generating a
  new span.
- Renaming spans does not change their trace IDs, span IDs, parent-child
  relationships, timings, or attributes.
- Prefixing names with a chapter number is an observable study example. Real
  deployments more commonly use transformations for normalization,
  sanitization, redaction, and backend compatibility.
- Only the traces pipeline references this processor, so metrics and logs are
  unaffected.

### Deliberately not introduced yet

- Conditional OTTL statements
- Filtering
- Memory limiting
- A telemetry backend
- Alloy or Kubernetes

## Step 10: Drop one span with a filter processor

### What changed

Added a trace filter that drops the transformed span named
`8:reverse characters`:

```yaml
filter/drop-reverse-span:
  error_mode: ignore
  trace_conditions:
    - span.name == "8:reverse characters"
```

The Go application remains unchanged and still creates all four spans.

### Concept demonstrated

The filter processor evaluates OTTL conditions against incoming telemetry.
When a condition evaluates to `true`, the matching telemetry is discarded:

```text
condition false -> span continues through the pipeline
condition true  -> span is dropped
```

Unlike the previous processors, this processor does not enrich or modify the
matching span. It prevents that span from reaching later processors and
exporters.

### Relevant pipeline configuration

```yaml
processors:
  - resource/add-local-environment
  - attributes/add-study-chapter
  - transform/prefix-span-name
  - filter/drop-reverse-span
  - batch
```

Order matters here too. The filter compares against the name
`8:reverse characters`, so the transform processor must add the `8:` prefix
before the filter runs. If the filter came first, it would see the original
name `reverse characters` and this condition would not match.

### Expected behavior

The application creates and exports four spans to the Collector, but the debug
exporter displays only these three:

```text
8:read random characters
8:save characters
8:process random characters
```

It does not display:

```text
8:reverse characters
```

### Important distinctions

- A filter condition describes what to drop, not what to keep.
- Filtering happens in the Collector and does not stop the application from
  performing or instrumenting the operation.
- The dropped span is not available to processors or exporters that follow the
  filter.
- `error_mode: ignore` controls errors while evaluating conditions. It does not
  mean that matching spans are retained.
- This example removes a leaf span, so it does not orphan a child span.
  Filtering a parent span could leave exported children whose parent is absent.
- Production filters should be narrowly targeted and tested because discarded
  telemetry cannot be recovered by a later processor.

### Deliberately not introduced yet

- Filtering by attributes, Resources, duration, or status
- Multiple filter conditions
- Memory limiting
- A telemetry backend
- Alloy or Kubernetes

## Step 11: Protect the Collector with a memory limiter

### What changed

Added a memory limiter with a fixed local-study limit:

```yaml
memory_limiter:
  check_interval: 1s
  limit_mib: 128
  spike_limit_mib: 32
```

It is referenced first in the traces, metrics, and logs pipelines.

### Concept demonstrated

The memory limiter periodically checks memory used by the Collector process. It
uses two thresholds:

```text
soft limit = limit_mib - spike_limit_mib
           = 128 MiB - 32 MiB
           = 96 MiB

hard limit = 128 MiB
```

Above the soft limit, it temporarily refuses incoming telemetry with a
non-permanent error. A compatible preceding component can retry or apply
backpressure, giving the Collector time to recover. Above the hard limit, the
processor can additionally force Go garbage collection.

### Relevant pipeline configuration

```yaml
service:
  pipelines:
    traces:
      processors:
        - memory_limiter
        - resource/add-local-environment
        - attributes/add-study-chapter
        - transform/prefix-span-name
        - filter/drop-reverse-span
        - batch

    metrics:
      processors: [memory_limiter, resource/add-local-environment, batch]

    logs:
      processors: [memory_limiter, resource/add-local-environment, batch]
```

Order matters. The memory limiter belongs first so it can refuse data before
later processors perform additional work and allocate more memory.

### Expected behavior

Under this exercise's tiny workload, telemetry should behave exactly as before:

- The Collector starts normally.
- Its startup output reports `Memory limiter configured` together with the
  `128 MiB` limit, `32 MiB` spike allowance, and one-second check interval.
- Traces, metrics, and logs continue reaching the debug exporter.
- The trace filter continues exporting three of the application's four spans.

The memory limiter becomes visible through warnings and refused data only if
the configured thresholds are exceeded. This exercise deliberately does not
generate artificial memory pressure.

### Important distinctions

- The memory limiter protects Collector availability; it is not ordinary
  filtering based on telemetry content.
- Refusing telemetry creates backpressure only when the preceding component
  handles the non-permanent error correctly.
- Retry is not guaranteed forever, so prolonged pressure can still result in
  data loss.
- Referencing the configuration in three pipelines creates one processor
  instance per pipeline. Those instances all observe memory used by the same
  Collector process.
- The hard limit targets Go heap allocation rather than total process memory;
  total process usage can be higher.
- A memory limiter complements proper Collector sizing and deployment memory
  limits; it does not replace them.
- These fixed values make the local example concrete and are not universal
  production recommendations.

### Deliberately not introduced yet

- Deliberately exhausting memory to trigger the limiter
- Container memory limits
- `GOMEMLIMIT`
- Percentage-based memory limits
- A telemetry backend
- Alloy or Kubernetes

## Step 12: Fan out traces to two exporters

### What changed

Added a file exporter alongside the debug exporter:

```yaml
exporters:
  debug:
    verbosity: detailed

  file/local-traces:
    path: /var/lib/otelcol-output/traces.json
    format: json
    flush_interval: 1s
```

Only the traces pipeline references both exporters:

```yaml
exporters: [debug, file/local-traces]
```

Docker Compose mounts the local `collector-output` directory into the
Collector so its otherwise read-only container filesystem has a writable,
persistent destination. Generated output is ignored by Git.

### Concept demonstrated

A single pipeline can fan out its processed telemetry to multiple exporters:

```text
                                      -> debug -> terminal
receiver -> processors -> fan-out ---|
                                      -> file  -> traces.json
```

The application sends the trace once. The trace passes through the pipeline's
processors once, and the resulting telemetry is offered to both exporters.

### Relevant pipeline configuration

```yaml
service:
  pipelines:
    traces:
      receivers: [otlp]
      processors:
        - memory_limiter
        - resource/add-local-environment
        - attributes/add-study-chapter
        - transform/prefix-span-name
        - filter/drop-reverse-span
        - batch
      exporters: [debug, file/local-traces]
```

Metrics and logs still reference only `debug`, so they are not written to this
trace file.

### Expected behavior

After running the Go trace sender:

- The debug exporter prints one trace containing three spans.
- `collector-output/traces.json` contains one line of OTLP JSON with the same
  three processed spans.
- Both outputs contain the transformed names and Collector-added attributes.
- Neither output contains the filtered `8:reverse characters` span.

### Important distinctions

- Fan-out creates multiple destinations, not multiple application exports.
- A second exporter does not create a second pipeline.
- Both exporters receive telemetry after all listed processors have run.
- The file exporter serializes OTLP data; it does not provide indexing,
  querying, retention management, dashboards, or distributed storage.
- An exporter failure can affect delivery to that destination. Configuring two
  exporters does not make one a durable backup queue for the other.
- The one-second file flush interval controls buffered file writes. It is
  separate from the five-second batch processor timeout.
- The exporter truncates its file when a new Collector instance starts because
  `append` remains at its default value of `false`. This keeps repeated study
  runs easy to inspect.

### Deliberately not introduced yet

- File rotation, compression, or append mode
- Exporter sending queues and retry tuning
- A real telemetry backend
- Host metrics collection
- Extensions
- Alloy or Kubernetes

## Step 13: Actively collect memory metrics

### What changed

Added a host-metrics receiver with only its memory scraper enabled:

```yaml
host_metrics/local:
  initial_delay: 1s
  collection_interval: 10s
  scrapers:
    memory:
```

The existing metrics pipeline now accepts metrics from both receivers:

```yaml
metrics:
  receivers: [otlp, host_metrics/local]
```

### Concept demonstrated

Receivers do not all obtain data in the same way:

```text
OTLP receiver
    Passively listens for telemetry pushed by another process.

host-metrics receiver
    Actively runs configured scrapers on a schedule.
```

The host-metrics receiver therefore produces metrics even when no application
or test generator sends anything to the Collector.

### Relevant data flow

```text
OTLP metric sender --------\
                            -> metrics pipeline -> debug exporter
memory scraper every 10s --/
```

Both sources share the metrics pipeline's memory limiter, Resource enrichment,
batch processor, and debug exporter. They remain separate metric streams even
though they travel through the same pipeline.

### Expected behavior

Start only the Collector and do not run `metrics-generator`. After the
one-second initial delay and the batch processor's export interval, the debug
exporter should print system memory metrics. Another scrape occurs every ten
seconds.

The collected metrics also receive:

```text
deployment.environment.name = "development"
```

because they pass through the existing Resource processor.

### Important distinctions

- `collection_interval` controls how often the receiver measures memory.
- The five-second batch timeout controls when accepted metrics are offered to
  exporters. It does not control when the measurements are made.
- Adding `host_metrics/local` under `receivers` only configures it; referencing
  it in the metrics pipeline activates it.
- One pipeline can accept telemetry from multiple receivers.
- The receiver creates system metrics; it does not derive them from traces or
  application metrics.
- If Node Exporter and Prometheus already collect the required node metrics,
  enabling an overlapping host-metrics receiver is usually unnecessary. Node
  Exporter fits the Prometheus scrape ecosystem and its existing dashboards;
  host metrics is useful when system measurements should enter an
  OpenTelemetry/OTLP pipeline directly. This step demonstrates an active
  receiver rather than recommending duplicate production collection.
- Because the Collector runs in Docker without a host-filesystem mount, this
  exercise observes the system view available inside the container. It should
  not be treated as a complete measurement of the Docker host.
- `memory_limiter` protects memory used by the Collector process. The
  host-metrics memory scraper measures system memory exposed to the container.
  They serve different purposes despite both involving memory.

### Deliberately not introduced yet

- CPU, disk, filesystem, network, process, and paging scrapers
- Mounting the host filesystem into the container
- Selecting or excluding individual metrics
- Filtering host metrics after collection
- A telemetry backend
- Extensions
- Alloy or Kubernetes

## Step 14: Scrape visible network interfaces

### What changed

Enabled the network scraper alongside the existing memory scraper:

```yaml
host_metrics/local:
  initial_delay: 1s
  collection_interval: 10s
  scrapers:
    memory:
    network:
```

No interface selection rule has been added yet.

### Concept demonstrated

One host-metrics receiver can run several independent scrapers on the same
schedule:

```text
host_metrics/local
    |-- memory scraper
    `-- network scraper
             |
             v
       metrics pipeline
```

Each scraper creates its own metric families. Enabling the network scraper does
not modify or derive data from the memory scraper.

### Expected behavior

With only the Collector running, the debug exporter continues to print memory
metrics and additionally prints network I/O metrics for every interface visible
from the container.

The network data points include an interface-identifying attribute, allowing us
to inspect the real names before writing a selection rule. A subsequent scrape
occurs every ten seconds.

Local verification exposed these interface values:

```text
device = "lo"
device = "eth0"
```

The output included `system.network.io`, `system.network.packets`,
`system.network.dropped`, `system.network.errors`, and the system-wide
`system.network.connections` metric.

### Important distinctions

- Enabling a scraper selects a category of measurements for collection.
- No processor is responsible for creating these network metrics.
- No application or telemetry generator needs to push them.
- This step collects all visible interfaces deliberately; it does not yet
  demonstrate receiver-level include or exclude rules.
- Interfaces visible inside the Collector container may differ from interfaces
  visible directly on the Docker host.
- As with memory metrics, an existing Node Exporter deployment may make this
  collection redundant in a real Prometheus-based environment.

### Deliberately not introduced yet

- Selecting a network interface at collection time
- Filtering a metric after it has been collected
- Other host-metrics scrapers
- Mounting the host filesystem into the container
- A telemetry backend
- Extensions
- Alloy or Kubernetes

## Step 15: Select an interface at collection time

### What changed

Configured the network scraper to include only the `eth0` interface discovered
in the previous step:

```yaml
network:
  include:
    interfaces: [eth0]
    match_type: strict
```

### Concept demonstrated

This include rule operates inside the receiver's scraper:

```text
eth0 -> measured -> metrics pipeline
lo   -> not measured
```

The unwanted interface data never enters the metrics pipeline.

### Expected behavior

Interface-scoped network metrics such as `system.network.io`,
`system.network.packets`, `system.network.errors`, and
`system.network.dropped` should contain:

```text
device = "eth0"
```

They should no longer contain:

```text
device = "lo"
```

The system-wide `system.network.connections` metric can remain because its data
points describe protocol and connection state rather than a particular network
interface.

Local verification showed only `device: Str(eth0)`. Compared with the previous
all-interface run, the batch decreased from 34 to 26 data points: eight `lo`
points were avoided across four metric families and two traffic directions.

### Important distinctions

- Receiver selection prevents unwanted measurements from being collected.
- A filter processor receives already-created telemetry and decides whether it
  continues downstream.
- `match_type: strict` requires the interface name to equal `eth0`; it is not a
  regular expression or substring match.
- This interface name came from the previous local observation rather than a
  platform-specific example copied from the book.
- Container interface names are environment-specific. A production
  configuration must select names appropriate to that deployment.
- This rule applies only to the network scraper. Memory collection and OTLP
  metrics remain unchanged.

### Deliberately not introduced yet

- Receiver exclusion rules or regular-expression matching
- Filtering collected host metrics with a processor
- Other host-metrics scrapers
- A telemetry backend
- Extensions
- Alloy or Kubernetes

## Step 16: Filter a metric after collection

### What changed

Added a metric filter that drops the network scraper's system-wide connection
metric:

```yaml
filter/drop-network-connections:
  error_mode: ignore
  metric_conditions:
    - metric.name == "system.network.connections"
```

Enabled it in the metrics pipeline after Resource enrichment and before
batching:

```yaml
processors:
  - memory_limiter
  - resource/add-local-environment
  - filter/drop-network-connections
  - batch
```

### Concept demonstrated

The filter receives a metric that the network scraper has already created. A
matching OTTL condition then prevents that metric from reaching subsequent
processors and exporters:

```text
network scraper
    -> creates system.network.connections
    -> filter condition matches
    -> metric is discarded
```

### Expected behavior

The debug exporter should no longer contain:

```text
system.network.connections
```

It should continue to contain:

```text
system.network.dropped
system.network.errors
system.network.io
system.network.packets
system.memory.usage
```

The previous batch contained six metric families and 26 data points. Removing
the connection metric and its 12 state data points should leave five metric
families and 14 data points.

Local verification produced exactly `metrics: 5, data points: 14`, with no
`system.network.connections` descriptor in the debug output.

### Important distinctions

- The `eth0` receiver include rule prevents `lo` measurements from being
  created.
- This processor filter discards `system.network.connections` only after the
  receiver has created it.
- A matching filter condition means drop, not retain.
- `error_mode: ignore` preserves valid telemetry when condition evaluation
  fails; it does not retain telemetry that successfully matches.
- The filter is enabled only in the metrics pipeline. It is independent from
  the trace filter that drops `8:reverse characters`.
- Filtering after collection can be useful when the receiver cannot express
  the desired condition or when centralized policy should apply to metrics
  arriving from several receivers.

### Deliberately not introduced yet

- Multiple metric filter conditions
- Filtering individual data points rather than a whole metric
- A telemetry backend
- Extensions
- Alloy or Kubernetes

## Step 17: Add a Collector health-check extension

### What changed

Defined a health-check extension:

```yaml
extensions:
  health_check:
    endpoint: 0.0.0.0:13133
```

Then activated it under `service`:

```yaml
service:
  extensions: [health_check]
```

Docker Compose publishes the endpoint only on the host loopback interface:

```yaml
- 127.0.0.1:13133:13133
```

### Concept demonstrated

An extension adds an operational capability to the Collector itself. It does
not receive, modify, or export telemetry:

```text
telemetry components:
    receivers -> processors -> exporters

Collector capability:
    health-check extension -> HTTP status endpoint
```

As with pipeline components, defining an extension only makes its configuration
available. Referencing it in `service.extensions` activates it.

### Expected behavior

With the Collector running:

```bash
curl http://localhost:13133/
```

returns HTTP status `200`, indicating that the Collector reports itself ready.
No telemetry generator is required, and the request does not pass through a
telemetry pipeline.

Local verification returned `HTTP 200`, while the Collector startup output
identified `health_check` as component kind `extension` and reported its state
as `ready`.

### Important distinctions

- In a Kubernetes deployment, this endpoint would mostly be used for Collector
  liveness or readiness probes.
- A liveness probe can cause Kubernetes to restart an unhealthy Collector.
- A readiness probe can keep traffic away from a Collector that is not ready.
- HTTP `200` does not prove that telemetry is reaching a backend or that every
  exporter is delivering successfully.
- If the Collector process is not running, the endpoint cannot be reached at
  all; that differs from a running endpoint returning an unhealthy status.
- The endpoint is bound to every interface inside the container so Docker can
  reach it, but Compose publishes it only on host loopback.
- This exercise deliberately does not enable the unreliable
  `check_collector_pipeline` option.

### Deliberately not introduced yet

- Kubernetes probe manifests
- TLS or authentication for operational endpoints
- End-to-end telemetry health monitoring
- A telemetry backend
- Other extensions
- Alloy

## Step 18: Synthesize the final Collector and Alloy concept map

### What changed

Created `ai_final_collector_concept_map.md` and refreshed the exercise README
to describe the completed configuration rather than its earlier Step 8 state.

### Concept demonstrated

The Collector is best understood as connected signal-specific pipelines:

```text
sources -> receivers -> ordered processors -> exporters
```

Extensions add operational capabilities outside those pipelines. Grafana Alloy
preserves much of the same OpenTelemetry processing model but expresses it as
direct component-to-component references rather than a central
`service.pipelines` section.

### Relevant Alloy distinction

Many components map directly:

```text
OTLP receiver    -> otelcol.receiver.otlp
memory limiter   -> otelcol.processor.memory_limiter
transform        -> otelcol.processor.transform
filter           -> otelcol.processor.filter
batch            -> otelcol.processor.batch
debug exporter   -> otelcol.exporter.debug
file exporter    -> otelcol.exporter.file
```

Other capabilities are architectural alternatives rather than direct wrappers:

```text
Collector host metrics -> Alloy prometheus.exporter.unix or existing Node Exporter
Collector health check -> Alloy built-in HTTP readiness/health endpoints
fixed Resource action  -> Alloy OTTL transform when no direct wrapper exists
```

### Expected behavior

This step changes documentation only. The already-verified Collector behavior
remains unchanged. The final map can now be used to trace each signal from its
source to its destination and to decide which parts would map directly—or need
an alternative—when designing an Alloy deployment.

### Important distinctions

- The map describes the completed learning configuration, not a ready-made
  production architecture.
- Alloy uses many upstream Collector components, but the configuration models
  are not identical.
- No one-to-one Alloy component is assumed when the current official component
  catalog does not list one.
- Existing Node Exporter collection is likely preferable to duplicating these
  study host metrics in the user's Kubernetes environment.
- Backend and Kubernetes decisions remain deliberately separate from this local
  chapter exercise.

### Deliberately not introduced yet

- An executable Alloy configuration
- Kubernetes manifests or placement decisions
- Tempo, Loki, Prometheus, or Mimir configuration
- TLS, authentication, persistent queues, or production retry tuning
- Sampling
