# Incremental steps taken

This file records each implemented step in the Chapter 10 Tempo exercise.

## Step 1: Run the smallest possible Tempo in Minikube

### What changed

Before deploying Tempo, we deleted and recreated the `minikube` profile to
start from a clean cluster. This removed the Chapter 9 runtime objects; their
manifests and study notes remain in the repository.

We created a dedicated `chapter10-tempo` namespace and added:

- a ConfigMap containing `tempo.yaml`
- a Deployment running one `grafana/tempo:3.0.2` Pod
- a Service exposing Tempo's HTTP API and two OTLP ingestion ports
- temporary `emptyDir` storage for backend data at `/data/tempo`
- a separate temporary runtime workspace at `/var/tempo`

The essential Tempo configuration is:

```yaml
target: all

server:
  http_listen_port: 3200

distributor:
  receivers:
    otlp:
      protocols:
        grpc:
          endpoint: "0.0.0.0:4317"
        http:
          endpoint: "0.0.0.0:4318"

storage:
  trace:
    backend: local
    wal:
      path: /data/tempo/wal
    local:
      path: /data/tempo/blocks
```

### Concept demonstrated

Tempo is a tracing backend rather than a general OpenTelemetry Collector
pipeline.

`target: all` starts the required Tempo components in one process. This is
called monolithic mode. The distributor accepts spans through OTLP, while the
storage configuration tells Tempo where to keep the write-ahead log and
completed trace blocks.

Kubernetes and Tempo have separate configuration responsibilities:

```text
Kubernetes manifests
  -> create and connect the ConfigMap, Pod, Service, and storage

tempo.yaml
  -> controls how the Tempo process receives and stores trace data
```

### Expected output and behavior

The Deployment should report one available replica:

```text
deployment.apps/tempo   1/1
```

After port-forwarding Service port `3200`, this request:

```bash
curl http://localhost:3200/ready
```

should return:

```text
ready
```

This proves that Tempo started with the supplied configuration. It does not
prove that a trace has been ingested yet.

Tempo 3.0.2 can emit idle startup warnings about automatic in-process worker
configuration, the deliberately disabled metrics generator, an unowned
`blocks` entry during WAL replay, and the backend scheduler having no jobs
before trace blocks exist. These messages do not prevent the Pod from becoming
ready.

### Important distinctions

- Port `3200` serves Tempo's HTTP API, including `/ready` and later query APIs.
- Ports `4317` and `4318` receive OTLP over gRPC and HTTP respectively.
- A Kubernetes Service supplies a stable network identity; it does not run
  Tempo or store traces.
- Tempo's distributor uses OpenTelemetry receiver implementations, but Tempo
  does not expose the Collector's general
  `receivers -> processors -> exporters` pipeline.
- Monolithic mode means Tempo components share one process and resource pool.
  It does not mean that Tempo has only one internal responsibility.
- `emptyDir` survives an individual container restart inside the same Pod but
  is deleted when that Pod is removed or replaced.
- `/data/tempo` contains the configured local storage backend, whereas
  `/var/tempo` is Tempo 3's live-store runtime workspace. Both are temporary in
  this step.

### Deliberately not introduced yet

- a trace-producing application or Job
- trace ingestion or querying
- Grafana and TraceQL
- Alloy or an OpenTelemetry Collector
- persistent storage
- object storage
- limits and per-tenant overrides
- the metrics generator and Prometheus
- Tempo microservices mode and Kafka

## Step 2: Ingest and retrieve one trace directly

### What changed

We added:

- a tiny Go trace producer
- a Dockerfile for its local Minikube image
- a Kubernetes Job that runs the producer exactly once

The Job sends OTLP/gRPC directly to the Tempo Service:

```yaml
env:
  - name: TEMPO_OTLP_ENDPOINT
    value: tempo:4317
```

The program creates one span, records its trace ID, ends the span, and shuts
down the tracer provider:

```go
_, span := tracer.Start(ctx, "send one trace to Tempo")
traceID := span.SpanContext().TraceID().String()
span.End()

if err := provider.Shutdown(shutdownCtx); err != nil {
	return fmt.Errorf("flush trace to Tempo: %w", err)
}

fmt.Printf("trace_id=%s\n", traceID)
```

### Concept demonstrated

Tempo's distributor is the write entry point. Its OTLP/gRPC receiver accepts
the span on Service port `4317`.

Printing the trace ID gives us a direct key for Tempo's read API:

```text
producer
  -> OTLP/gRPC
  -> Tempo distributor
  -> recent/storage path
  -> GET /api/traces/<trace-id>
```

This demonstrates both basic paths without a visual frontend:

- write path: send a span to Tempo
- read path: retrieve the trace from Tempo by ID

### Expected output and behavior

The Kubernetes Job should complete:

```text
job.batch/trace-producer   Complete
```

Its logs should contain:

```text
exported one span directly to tempo:4317
trace_id=<32-character trace ID>
```

With Tempo port `3200` forwarded locally, the following should return the trace
as JSON:

```bash
curl -H 'Accept: application/json' \
  "http://localhost:3200/api/traces/<trace-id>"
```

The response should contain:

- `service.name` equal to `chapter10-trace-producer`
- span name `send one trace to Tempo`
- the `study.step` attribute
- the span event added by the producer

### Verified in Minikube

The Job completed successfully and printed:

```text
exported one span directly to tempo:4317
trace_id=609a002f4c4e67599db66c796bdce0d4
```

Tempo's own counters confirmed one accepted span:

```text
tempo_distributor_spans_received_total{tenant="single-tenant"} 1
tempo_distributor_bytes_received_total{tenant="single-tenant"} 341
```

Retrieving that exact trace ID returned JSON containing the expected service,
span name, `study.step` attribute, and span event.

### Important distinctions

- Port `4317` accepts OTLP/gRPC writes; port `3200` serves Tempo query APIs.
- Kubernetes Job completion means the producer exited successfully.
- Provider shutdown is what waits for the batched span to be exported before
  the short-lived Job exits.
- The trace ID identifies the whole trace. This example has one trace
  containing one span, so that relationship is not visually complex yet.
- Direct application-to-Tempo export is useful for isolating the backend in
  this exercise. An Alloy or Collector layer is recommended in production for
  batching, retries, processing, and decoupling.

### Deliberately not introduced yet

- Grafana or TraceQL search
- Alloy or an OpenTelemetry Collector
- persistent storage
- object storage
- ingestion limits or per-tenant overrides
- metrics derived from traces
- Tempo microservices mode

## Step 3: Add Grafana and query Tempo

### What changed

We added:

- a Grafana data-source provisioning ConfigMap
- a single-replica Grafana Deployment
- a Service exposing Grafana on port `3000` inside the cluster

The provisioned data source points Grafana at Tempo's HTTP query endpoint:

```yaml
datasources:
  - name: Tempo
    type: tempo
    uid: tempo
    access: proxy
    url: http://tempo:3200
    isDefault: true
    jsonData:
      streamingEnabled:
        search: false
        metrics: false
```

Grafana reads that YAML file at startup because the ConfigMap is mounted in its
data-source provisioning directory.

### Concept demonstrated

Tempo and Grafana have different responsibilities:

```text
Tempo
  -> receives, stores, searches, and returns trace data

Grafana
  -> sends queries to Tempo and renders the returned traces
```

The URL `http://tempo:3200` works because both Pods are in the
`chapter10-tempo` namespace and Kubernetes DNS resolves the `tempo` Service
name. It is a query connection. Grafana does not send OTLP through ports `4317`
or `4318`.

Trace IDs and TraceQL provide two different ways to query:

- a trace ID directly retrieves one already-known trace
- TraceQL searches for traces whose spans or resources satisfy conditions

The first TraceQL expression searches by an OpenTelemetry resource attribute:

```traceql
{ resource.service.name = "chapter10-trace-producer" }
```

Grafana can optionally stream partial TraceQL results as Tempo finds them. That
feature requires `stream_over_http_enabled: true` in a self-managed Tempo
configuration. We explicitly leave streaming off in this step: normal HTTP
queries still work, but Grafana displays their results after each query
finishes. This keeps the first Grafana step focused on the basic query path.

### Expected output and behavior

After port-forwarding Grafana:

```bash
kubectl port-forward -n chapter10-tempo service/grafana 3000:3000
```

opening `http://localhost:3000/explore` should show the provisioned `Tempo` data
source. Looking up the Step 2 trace ID or running the TraceQL query should
produce a result that can be expanded into a visual trace view.

That view should contain:

- service `chapter10-trace-producer`
- span `send one trace to Tempo`
- resource attributes such as service version and environment
- span attribute `study.step`
- the span event emitted by the producer

### Verified in Minikube

Both the Tempo and Grafana Deployments reached `1/1` ready. Grafana loaded the
provisioned data source with UID `tempo`, and its health API returned:

```json
{"message":"Data source is working","status":"OK"}
```

An exact trace-ID request through Grafana's data-source proxy returned the
original Step 2 trace. The non-streaming TraceQL query:

```traceql
{ resource.service.name = "chapter10-trace-producer" }
```

also returned trace ID `609a002f4c4e67599db66c796bdce0d4`, root service
`chapter10-trace-producer`, and root span `send one trace to Tempo`.

Grafana's data-source health test also probed the optional gRPC streaming path
and logged a failed gRPC connection to Tempo's HTTP port. This does not indicate
failure of the ordinary HTTP query path: the health result was `OK`, and both
the exact trace lookup and TraceQL search succeeded. Enabling streaming later
would require `stream_over_http_enabled: true` in Tempo.

### Important distinctions

- A Grafana data source is a connection to a backend; it is not another copy of
  the trace data.
- Grafana port `3000` serves the user interface, whereas Tempo port `3200`
  serves Tempo's HTTP query API.
- Tempo's OTLP ports receive telemetry, while its query port returns telemetry.
- A trace-ID lookup is exact and requires the caller to know the ID.
- TraceQL searches trace contents and can discover traces whose IDs are not
  already known.
- Streaming changes when partial search results appear; it is not required to
  issue a TraceQL query or display its completed results.
- The anonymous Admin setting is only for this isolated Minikube exercise. A
  real deployment needs authentication and appropriately limited roles.

### Deliberately not introduced yet

- Alloy or an OpenTelemetry Collector
- persistent Tempo storage
- object storage
- ingestion limits or per-tenant overrides
- metrics generated from traces
- Prometheus
- Tempo microservices mode

## Step 4: Persist Tempo's local backend with a PVC

### What changed

We added a 1 GiB `ReadWriteOnce` PersistentVolumeClaim:

```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: tempo-data
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```

The Tempo Deployment now mounts that claim at `/data/tempo`:

```yaml
- name: tempo-data
  persistentVolumeClaim:
    claimName: tempo-data
```

This path contains the configured Tempo WAL and completed local-backend blocks.
The separate `/var/tempo` live-store workspace remains an `emptyDir`.

### Concept demonstrated

Kubernetes storage lifecycle and Tempo storage lifecycle are related but
different:

```text
Tempo
  -> turns recent trace data into WAL data and completed blocks

Kubernetes PVC
  -> keeps those files when the Tempo Pod is deleted or replaced
```

An `emptyDir` belongs to one Pod. A PVC is a separate Kubernetes object whose
bound volume can be mounted by the replacement Pod.

Tempo still uses its `local` storage backend in this exercise. The PVC changes
the durability of the filesystem below that backend; it does not change Tempo
to an object-storage backend.

### Expected output and behavior

The claim should be bound:

```text
tempo-data   Bound   1Gi   RWO
```

After Tempo flushes a trace into a completed backend block, deleting the Tempo
Pod should cause the Deployment to create a different Pod. Querying the same
trace ID should still return its resource, span, attributes, and event.

### Verified in Minikube

The claim bound successfully to Minikube's `standard` StorageClass. After
attaching it, we recreated the producer Job and received:

```text
trace_id=98c19333f8ca70c77d945f19e1219904
```

Tempo then reported:

```text
tempo_live_store_local_blocks_flushed_total 1
```

We retrieved the trace, deleted Pod `tempo-5c89d879bb-d8mbr`, and waited for the
Deployment to create replacement Pod `tempo-5c89d879bb-b5t5t`. The PVC remained
bound to the same volume. Both Tempo's API and Grafana's data-source proxy
retrieved the same trace ID from the replacement Pod, including:

- service `chapter10-trace-producer`
- span `send one trace to Tempo`
- attribute value `tempo-direct-ingestion`

### Important distinctions

- A PVC persists files across Pod replacement; it is not a backup and does not
  by itself provide replication or high availability.
- `ReadWriteOnce` means the volume can be mounted read/write by one node. It
  does not mean only one Pod can ever use the claim during its lifetime.
- The Tempo local backend remains intended for this monolithic learning
  environment. Production Tempo normally uses object storage.
- Recently ingested traces first live in memory and live-store runtime state.
  The persistence test waited for
  `tempo_live_store_local_blocks_flushed_total` to prove that a completed block
  had reached the PVC-backed backend.
- Exact lookup and TraceQL search have different read behavior. Exact lookup
  found the persisted trace immediately after restart.
- Tempo's current default
  `query_frontend.search.query_backend_after: 15m` routes searches for the most
  recent interval to live-store rather than backend blocks. Because
  `/var/tempo` remains temporary, TraceQL did not immediately rediscover this
  very recent trace after restart even though exact lookup proved its backend
  block was durable. It becomes eligible for backend search after that
  boundary.

### Deliberately not introduced yet

- object storage such as S3, GCS, or Azure Blob Storage
- storage replication, backup, or disaster recovery
- retention configuration or compaction tuning
- Alloy or an OpenTelemetry Collector
- ingestion limits or per-tenant overrides
- metrics generated from traces
- Prometheus
- Tempo microservices mode

## Step 5: Put Alloy between the producer and Tempo

### What changed

We added:

- an Alloy configuration ConfigMap
- a single-replica Alloy Deployment
- an Alloy Service exposing OTLP/gRPC and Alloy's HTTP debugging interface
- a backend-neutral `OTLP_ENDPOINT` producer setting

The producer now uses:

```yaml
- name: OTLP_ENDPOINT
  value: alloy:4317
```

Alloy connects three labeled components into a trace pipeline:

```alloy
otelcol.receiver.otlp "applications" {
  grpc {
    endpoint = "0.0.0.0:4317"
  }

  output {
    traces = [otelcol.processor.batch.traces.input]
  }
}

otelcol.processor.batch "traces" {
  output {
    traces = [otelcol.exporter.otlp.tempo.input]
  }
}

otelcol.exporter.otlp "tempo" {
  client {
    endpoint = "tempo:4317"
    tls {
      insecure = true
    }
  }
}
```

### Concept demonstrated

Alloy is now the collection layer:

```text
producer
  -> OTLP/gRPC
  -> Alloy receiver
  -> Alloy batch processor
  -> Alloy exporter
  -> OTLP/gRPC
  -> Tempo
```

The producer knows only the collection address `alloy:4317`. Alloy knows the
backend address `tempo:4317`. Changing the backend or adding another exporter
would therefore be an Alloy configuration concern rather than an application
configuration concern.

Alloy expresses this flow as a component graph. A component's `output` block
refers to the next component's typed `input`. The equivalent vanilla
OpenTelemetry Collector configuration would connect named components in a
`service.pipelines.traces` entry:

```text
receivers: [otlp] -> processors: [batch] -> exporters: [otlp/tempo]
```

### Expected output and behavior

Alloy should become ready and healthy:

```text
Alloy is ready.
All Alloy components are healthy.
```

Running the producer Job should report:

```text
exported one span through Alloy at alloy:4317
trace_id=<32-character trace ID>
```

Alloy's internal metrics should then report an accepted span and an exported
span. Retrieving the printed ID from Tempo should return the same span.

### Verified in Minikube

Alloy v1.18.0 loaded all three components, started its gRPC receiver on port
`4317`, and reached `1/1` ready. Both HTTP checks succeeded.

The recreated producer Job completed and printed:

```text
exported one span through Alloy at alloy:4317
trace_id=215531b693001a0599449cdb76769246
```

Alloy then reported:

```text
otelcol_receiver_accepted_spans_total{...transport="grpc"} 1
otelcol_exporter_sent_spans_total{...server_address="tempo",server_port="4317"} 1
otelcol_receiver_failed_spans_total{...transport="grpc"} 0
```

Tempo returned that exact trace ID with:

- service `chapter10-trace-producer`
- span `send one trace through Alloy to Tempo`
- attribute `study.step=tempo-via-alloy`
- event `the example span is ready for Alloy`

Together, the endpoint, Alloy counters, and Tempo result prove that the trace
followed the intended path.

### Important distinctions

- OTLP is used on both network hops, but Alloy and Tempo do different jobs.
  Alloy receives, processes, and exports telemetry; Tempo stores and queries
  traces.
- The application SDK still batches before export, and Alloy batches received
  telemetry again for its separate outbound connection. These are independent
  process and network boundaries.
- A Kubernetes Service gives Alloy a stable address. It does not itself
  receive or process telemetry.
- Alloy's component graph is conceptually similar to the Collector pipeline,
  but it is wired through component references rather than one central
  `service.pipelines` list.
- This is a centralized traces-only Deployment. Alloy clustering, a DaemonSet,
  and a sidecar would be separate deployment choices.

### Deliberately not introduced yet

- attribute enrichment or transformation
- filtering
- head or tail sampling
- multiple exporters or routing
- durable Alloy queues
- authentication or TLS
- ingestion limits or per-tenant overrides
- metrics generated from traces
- Prometheus
- Tempo object storage or microservices mode

## Step 6: Limit the size of stored attributes

### What changed

We added one default ingestion override to Tempo:

```yaml
overrides:
  defaults:
    ingestion:
      max_attribute_bytes: 64
```

The producer now creates a deliberately oversized span attribute:

```go
attribute.String(
	"study.oversized_attribute",
	strings.Repeat("x", 128),
)
```

The Job continues to send OTLP to Alloy, which continues to batch and export
the span to Tempo.

### Concept demonstrated

An ingestion override tells Tempo how to handle telemetry as it enters the
backend. `max_attribute_bytes` applies a maximum byte length to individual
attribute keys and values.

The behavior in this experiment is:

```text
128-byte span attribute
        |
        v
Tempo distributor limit: 64 bytes
        |
        +-> trace remains accepted
        |
        +-> stored attribute is truncated to 64 bytes
        |
        +-> diagnostic metric and rate-limited log are emitted
```

This protects storage and query components from unexpectedly large values such
as request bodies, large headers, full database statements, or message
payloads.

### Expected output and behavior

The Job should still complete because truncation does not reject its span.
Retrieving the trace should show a `study.oversized_attribute` value containing
64 `x` characters rather than the 128 characters created by the producer.

Tempo should report:

```text
tempo_distributor_attributes_truncated_total{
  scope="span",
  tenant="single-tenant"
} 1
```

### Verified in Minikube

Tempo's effective configuration showed both:

```text
distributor.max_attribute_bytes: 2048
overrides.defaults.ingestion.max_attribute_bytes: 64
```

The override is the effective tenant/default ceiling for this experiment.

The producer completed through Alloy and printed:

```text
trace_id=4fa23ed5f2d4e8b54b6f066350caf55a
```

Retrieving that exact trace showed:

```text
stored_length=64
stored_value=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

Tempo's metric reported exactly one truncated span attribute:

```text
tempo_distributor_attributes_truncated_total{
  scope="span",
  tenant="single-tenant"
} 1
```

Tempo also logged:

```text
attributes truncated
max_size_bytes=64
example_scope=span
example_name=study.oversized_attribute
example_field=value
example_orig_size=128
```

### Important distinctions

- The entire trace was accepted; only the oversized attribute value changed.
- Truncation is different from filtering, which would remove a selected
  attribute or span according to a policy.
- Truncation is also different from rate limiting or `max_bytes_per_trace`,
  which can reject or discard telemetry.
- The OpenTelemetry SDK and Alloy forwarded the original 128-byte value. Tempo
  enforced this backend policy at ingestion.
- `defaults` applies when no more specific tenant override replaces it.
- The `scope="span"` metric label identifies an attribute attached to a span.
  Other possible scopes include resource, instrumentation scope, event, and
  link.
- The limit counts bytes, not user-perceived characters. ASCII `x` makes the
  two counts equal in this example.

### Deliberately not introduced yet

- per-tenant runtime override files
- user-configurable overrides through Tempo's API
- ingestion rate and burst limits
- maximum live-trace or total-trace-size limits
- outright rejection or discarded spans
- Alloy filtering or transformation
- metrics generated from traces
- Prometheus
- Tempo object storage or microservices mode

## Step 7: Generate Prometheus metrics from traces

### What changed

We added:

- a Prometheus ConfigMap, Deployment, and Service
- Prometheus's opt-in remote-write receiver
- Prometheus exemplar storage
- Tempo metrics-generator storage and remote-write configuration
- Tempo's `span-metrics` and `service-graphs` processors
- a provisioned Prometheus data source in Grafana
- links between the Prometheus and Tempo data sources

The core Tempo configuration is:

```yaml
metrics_generator:
  registry:
    collection_interval: 5s
  storage:
    path: /var/tempo/metrics-generator/wal
    remote_write_add_org_id_header: false
    remote_write:
      - url: http://prometheus:9090/api/v1/write
        send_exemplars: true

overrides:
  defaults:
    metrics_generator:
      processors:
        - span-metrics
        - service-graphs
```

Prometheus explicitly accepts pushed remote-write samples:

```yaml
args:
  - --web.enable-remote-write-receiver
  - --enable-feature=exemplar-storage
```

### Concept demonstrated

The application still emits only traces. Tempo inspects those spans and derives
metrics:

```text
producer
  -> trace
  -> Alloy
  -> Tempo
       |
       +-> store the trace
       |
       +-> derive metric samples
             |
             +-> remote write
                   |
                   v
               Prometheus
```

The `span-metrics` processor derives RED-style measurements:

- request rate from a span counter
- errors from the span status label
- duration from a span latency histogram

The `service-graphs` processor looks for cross-service parent-child
relationships and derives service edges. It cannot create an edge from one
internal span belonging to one service.

### Expected output and behavior

After a trace passes through Tempo and the five-second collection interval
elapses, Prometheus should contain:

```promql
traces_spanmetrics_calls_total
traces_spanmetrics_latency_count
traces_spanmetrics_latency_sum
traces_spanmetrics_latency_bucket
```

The histogram should also contain an exemplar carrying a representative trace
ID. Grafana should be able to query Prometheus and use that ID to open Tempo.

### Verified in Minikube

Prometheus v3.13.2 reached `1/1` ready. Its status API confirmed:

```text
web.enable-remote-write-receiver = true
enable-feature = exemplar-storage
```

Tempo's effective configuration showed:

```text
collection_interval: 5s
remote_write URL: http://prometheus:9090/api/v1/write
send_exemplars: true
processors: [span-metrics, service-graphs]
```

The unchanged producer emitted one trace through Alloy:

```text
trace_id=2d7de15d0253a916164d701a635b5b4e
```

Prometheus then returned:

```text
traces_spanmetrics_calls_total = 1
traces_spanmetrics_latency_count = 1
traces_spanmetrics_latency_sum = 0.000527193
```

The series had these useful labels:

```text
service="chapter10-trace-producer"
span_name="demonstrate Tempo attribute truncation"
span_kind="SPAN_KIND_INTERNAL"
status_code="STATUS_CODE_UNSET"
source="tempo"
```

The latency histogram exemplar contained:

```text
traceID="2d7de15d0253a916164d701a635b5b4e"
value="0.000527193"
```

Tempo's own metrics confirmed:

```text
prometheus_remote_storage_exemplars_in_total 1
prometheus_remote_storage_samples_in_total 113
tempo_metrics_generator_registry_active_series{tenant="single-tenant"} 19
```

Grafana reported its provisioned Prometheus data source healthy, linked the
Tempo service map to Prometheus, and mapped the Prometheus `traceID` exemplar
label back to Tempo.

### Important distinctions

- Tempo stores traces; Prometheus stores the metrics derived from those traces.
- The producer did not create or export OpenTelemetry metrics in this step.
- Remote write is a push path from Tempo to Prometheus. It is different from
  Prometheus scraping a `/metrics` endpoint.
- `metrics_generator.storage.path` is the generator's Prometheus Agent WAL. It
  is separate from Tempo's trace WAL and trace blocks.
- Declaring `metrics_generator` does not enable processors. The processors must
  also be selected under the overrides.
- Span metrics describe individual operations even when a trace contains only
  one service.
- Service graphs require relationships between spans from different services.
  Enabling the processor alone cannot invent those relationships.
- An exemplar is not another metric or trace. It is a representative trace ID
  attached to a particular metric observation.
- The generated metric labels are useful dimensions, but adding too many
  dimensions creates high cardinality.

### Deliberately not introduced yet

- application-generated OpenTelemetry metrics
- custom span-metric dimensions
- metrics-generator cardinality tuning
- persistent Prometheus storage
- Prometheus high availability or long-term metrics storage
- Tempo object storage or microservices mode

## Step 7b: Derive a service-graph edge

### What changed

We added a separate service-graph producer under the existing Go module and a
one-shot Kubernetes Job:

```text
trace-producer/service-graph/main.go
k8s/service-graph-producer-job.yaml
```

The existing attribute-limit producer remains unchanged, so each example still
has one focused purpose.

The new producer creates:

```text
greeting-service: GET /greet                  SERVER
└── greeting-service: call formatting service CLIENT
    └── formatting-service: POST /format       SERVER
```

### Concept demonstrated

A service graph is inferred from relationships already present in tracing
data. Tempo needs a parent-child jump between spans belonging to different
services; enabling the processor cannot invent a missing relationship.

The focused propagation code is:

```go
remoteParent := trace.ContextWithRemoteSpanContext(
	context.Background(),
	trace.SpanContextFromContext(clientCtx),
)

_, formattingServer := formattingTracer.Start(
	remoteParent,
	"POST /format",
	trace.WithSpanKind(trace.SpanKindServer),
)
```

The server span receives the client span's trace and parent identifiers. The
remote marker represents the process boundary that HTTP context extraction
would establish in a real distributed application.

Two tracer providers give the spans different resources:

```text
client span resource: service.name="greeting-service"
server span resource: service.name="formatting-service"
```

### Expected output and behavior

The Job prints the new trace ID:

```text
exported one cross-service trace through Alloy at alloy:4317
trace_id=<trace ID>
expected_service_graph_edge=greeting-service -> formatting-service
```

After Tempo's five-second metrics collection interval, Prometheus should
contain:

```promql
traces_service_graph_request_total{
  client="greeting-service",
  server="formatting-service"
}
```

It should also contain service-graph duration histograms and ordinary span
metrics for all three spans.

### Verified in Minikube

The Job completed and emitted:

```text
trace_id=782062d1f5ec21ea66590000ff581394
```

Retrieving the trace from Tempo showed:

```text
greeting-service   GET /greet               SERVER  parent: none
greeting-service   call formatting service  CLIENT  parent: GET /greet
formatting-service POST /format              SERVER  parent: client span
```

Every span had the same trace ID and `STATUS_CODE_OK`.

Prometheus returned:

```text
traces_service_graph_request_total{
  client="greeting-service",
  server="formatting-service",
  source="tempo"
} 1
```

It also returned:

```text
traces_service_graph_request_server_seconds_count{
  client="greeting-service",
  server="formatting-service",
  source="tempo"
} 1
```

The server-latency histogram exemplar was:

```text
traceID="782062d1f5ec21ea66590000ff581394"
value="0.010318462"
```

The span-metrics processor independently produced one calls-counter
observation for each of the three spans.

### Important distinctions

- Tempo derives the edge; neither service emits a
  `traces_service_graph_request_total` metric.
- The edge comes from trace structure plus the two `service.name` resources,
  not merely from `server.address="formatting-service"`.
- The client and server spans belong to the same trace, but different
  services.
- A client span and its corresponding server span describe the two sides of
  one request. They are not duplicate spans.
- The root greeting server span supplies realistic request context but does
  not itself create the cross-service edge.
- One process represents two logical services only to keep this study step
  small. Real services would use separate processes and propagate the context
  through HTTP, gRPC, or messaging headers.
- Service-graph metrics and span metrics are different outputs from the same
  tracing data.

### Deliberately not introduced yet

- two real networked application services
- automatic HTTP context injection and extraction
- failed or retried cross-service requests
- messaging and database virtual nodes
- custom service-graph dimensions
- service-graph cardinality and edge-wait tuning
- production Tempo architecture

## Step 8: Map the learning deployment to production Tempo

### What changed

We added a documentation-only production architecture map:

```text
ai_production_tempo_architecture_map.md
```

It compares every important choice in the Minikube deployment with its likely
production replacement and records the decisions that cannot be copied from an
example.

No Kubernetes resources were changed. Kafka, object storage, and distributed
Tempo components were deliberately not deployed merely to illustrate their
names.

### Concept demonstrated

The monolithic and microservices modes use the same Tempo binary but arrange
its responsibilities differently.

Our learning path is:

```text
Alloy -> one target:all Tempo process -> local PVC
Grafana -------------------------------> same process
```

A representative production microservices path is:

```text
write:
Alloy -> secured gateway -> distributors -> Kafka
                                         ├-> live-stores
                                         ├-> block-builders -> object storage
                                         └-> metrics-generators -> metrics backend

read:
Grafana -> secured gateway -> query-frontends -> queriers
                                                   ├-> live-stores
                                                   └-> object storage
```

The production change is not simply "add more Tempo replicas." It separates
write, recent-data, historical-storage, metrics-generation, query, and backend
maintenance failure domains.

### Expected output and behavior

The artifact is a reviewable decision map rather than new runtime output. It
should let a reader answer:

- whether monolithic or microservices mode fits the requirements
- why Tempo 3 microservices mode needs Kafka
- why a shared object store replaces the trace-block PVC
- which components scale for writes, reads, recent data, block creation,
  metrics generation, and compaction
- where authentication, TLS, and trusted tenant assignment belong
- what drives retention, limits, sizing, availability, and monitoring

### Important distinctions

- Monolithic Tempo can be reasonable for a small production environment when
  its failure domain and scaling limits are explicitly acceptable.
- Microservices mode is required for independent scaling and is the normal
  direction for high availability.
- Kafka is durable ingest transport in microservices mode; object storage is
  the long-term trace database.
- A PVC and object storage are not interchangeable in a distributed
  deployment because all relevant components need shared block access.
- Tempo has no built-in authentication layer. A trusted reverse proxy or
  gateway must secure the boundary.
- `X-Scope-OrgID` identifies a tenant but does not authenticate the caller.
- Helm manages Kubernetes resource composition; it cannot choose business
  retention, security, capacity, or tenancy policy.
- Readiness verifies a component, while a synthetic trace verifies the complete
  system path.
- Sampling and preventive redaction normally happen in Alloy or another
  collection tier before Tempo.

### Deliberately not introduced

- a production-sized Helm values file without production requirements
- a Kafka-compatible cluster
- S3, GCS, Azure Blob, MinIO, or Ceph deployment
- fabricated replica counts or resource requests
- production credentials
- multi-zone failure testing
- a migration of the running Minikube exercise

The chapter exercise now covers all eight planned steps.
