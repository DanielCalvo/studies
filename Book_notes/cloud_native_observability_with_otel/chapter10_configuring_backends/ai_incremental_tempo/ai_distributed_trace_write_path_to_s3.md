# How a trace travels through Alloy and distributed Tempo to S3

This walkthrough assumes:

- The application uses an OpenTelemetry SDK.
- The trace is sampled.
- The application exports OTLP data to Grafana Alloy.
- Alloy forwards that data to Tempo.
- Tempo 3 runs in microservices mode.
- Tempo uses a Kafka-compatible system for ingestion and S3 for long-term
  object storage.

The principal write path is:

```text
Application
  -> OpenTelemetry SDK
  -> Alloy
  -> optional ingress gateway
  -> Tempo distributor
  -> Kafka
  -> Tempo block-builder
  -> local scratch disk
  -> S3
```

Kafka also sends the same data to independent consumers:

```text
Kafka
  |-> block-builder      -> S3
  |-> live-store         -> recent trace queries
  `-> metrics-generator  -> Prometheus-compatible metrics backend
```

Only the block-builder branch creates the long-term trace blocks in S3.

## 1. The application creates spans

Suppose one request passes through three services:

```text
frontend -> orders-service -> payment-service
```

Each instrumented service creates spans. Context propagation ensures that
those spans share a trace ID and have the correct parent-and-child
relationships.

A span can contain:

- Trace ID and span ID
- Parent span ID
- Service and operation names
- Start time and duration
- Status
- Resource and span attributes
- Events and links

The application normally exports completed **spans**, not one magical,
complete trace object. Different services can finish and export their spans
at different times. Tempo reconstructs the logical trace from the spans later.

Sampling occurs before export. A trace that is not selected by the sampling
decision does not enter this pipeline.

## 2. The OpenTelemetry SDK batches and exports spans

When a span ends, the SDK gives it to its span processor. A production
application commonly uses a batch span processor:

```text
ended spans
  -> in-memory SDK queue
  -> batch
  -> OTLP exporter
```

The exporter serializes the spans as OTLP and sends them to Alloy, usually
through:

- OTLP/gRPC on port `4317`
- OTLP/HTTP on port `4318`

Batching reduces the number of network requests made by the application.

## 3. Alloy receives and processes the OTLP data

Alloy's OTLP receiver accepts the batches. The spans then follow the component
graph defined in the Alloy configuration:

```text
otelcol.receiver.otlp
  -> processors
  -> otelcol.exporter.otlp
```

Depending on its configuration, Alloy can:

- Batch spans
- Add, remove, or transform attributes
- Add Kubernetes metadata
- Apply memory limits
- Filter spans
- Perform tail sampling
- Retry failed exports
- Buffer data in an exporter queue

These operations happen only when their Alloy components are configured and
connected into the graph. Alloy's OTLP exporter eventually sends the spans to
Tempo's ingestion endpoint.

See [Grafana Alloy tracing documentation][alloy-tracing].

## 4. An optional gateway receives the request

A production deployment commonly places an ingress gateway or load balancer
in front of the Tempo distributors:

```text
Alloy -> gateway -> Tempo distributors
```

That gateway can provide:

- TLS
- Authentication
- Tenant identification
- Load balancing
- A network-policy boundary

In a multitenant installation, a request commonly carries an
`X-Scope-OrgID` tenant header. The gateway is operational infrastructure; it
does not permanently store the trace.

## 5. A Tempo distributor validates and partitions the spans

The gateway forwards the OTLP request to a Tempo distributor. The distributor
is Tempo's ingestion entry point. It:

1. Decodes the OTLP request.
2. Determines the tenant.
3. Validates the data against configured ingestion limits.
4. Applies the configured rejection or truncation behavior.
5. Hashes the trace ID.
6. Uses the result to select a Tempo/Kafka partition.
7. Writes records into Kafka.

Hashing by trace ID routes spans from the same trace to the same partition.
The distributor is horizontally scalable and does not permanently own the
trace.

See the [Tempo distributor documentation][distributor].

## 6. Kafka becomes the immediate durable copy

Kafka stores the accepted span records in a partition of Tempo's trace topic:

```text
Tempo distributor
  -> topic: tempo-traces
      -> partition N
```

Kafka is Tempo's durable write-ahead log in microservices mode. It is not the
long-term trace database.

The distributor reports a successful write to Alloy only after Kafka
acknowledges the data. At that boundary, a Tempo component can crash and the
downstream consumers can replay the records from Kafka.

The guarantee observed by the application also depends on the Alloy
configuration. For example, if Alloy has accepted data into a non-persistent
in-memory batch queue, the application can receive success before that data
has reached Tempo and Kafka.

Kafka's partition count also controls the maximum parallelism available to
several downstream Tempo components.

See the [Tempo Kafka documentation][tempo-kafka].

## 7. Independent consumers read the same records

Tempo uses separate Kafka consumer groups. Each group reads the trace records
independently and maintains its own offsets.

### Block-builder

The block-builder produces long-term Tempo blocks and uploads them to S3.
This is the branch that ultimately persists the trace.

### Live-store

The live-store holds recent trace data in memory and temporary local blocks.
It makes new traces queryable before the block-builder has placed them in S3.

Consequently, a trace queried only seconds after ingestion will probably be
served from the live-store rather than S3.

See the [Tempo live-store documentation][live-store].

### Metrics-generator

When enabled, the metrics-generator derives information such as:

- Span request counts
- Span duration histograms
- Service graph metrics

It sends those metrics to a Prometheus-compatible backend. It does not store
them as trace blocks in S3.

The consumer groups proceed at their own pace. For example, a delayed
metrics-generator does not inherently stop the block-builder.

## 8. The block-builder consumes a window of records

The block-builder does not upload every span as an individual S3 object. It
works in cycles:

1. Select a range of Kafka records.
2. Read the records in that range.
3. Group spans by tenant and partition.
4. Deduplicate spans when necessary.
5. Organize them into Tempo blocks.
6. Build the blocks on local scratch disk.
7. Upload the completed blocks to S3.
8. Commit the consumed Kafka offset.

The scratch disk is temporary working space, not the long-term trace store.

The block-builder does not know when a distributed trace is complete. If late
spans arrive during another consumption cycle, one logical trace can be spread
across multiple blocks. Tempo's query path joins those spans when the trace is
read.

See the [Tempo block-builder documentation][block-builder].

## 9. The spans become a Tempo Parquet block

A simplified block layout in S3 looks like this:

```text
tempo-bucket/
`-- tenant-a/
    `-- block-id/
        |-- data.parquet
        |-- meta.json
        |-- bloom-0
        |-- bloom-1
        `-- index
```

The files have different responsibilities:

- `data.parquet` contains the span data.
- `meta.json` describes the block, tenant, time range, and other metadata.
- Bloom filters accelerate trace ID lookups.
- Index data helps locate traces inside the Parquet data.

Tempo therefore does not normally store one JSON object per trace. It packs
many traces and spans into query-efficient blocks.

See the [Tempo block-format documentation][block-format].

## 10. The block-builder uploads the block to S3

The block-builder uploads the important block files in this order:

1. Bloom filters and indexes
2. `data.parquet`
3. A temporary `nocompact.flg`
4. `meta.json`

Writing `meta.json` is effectively the publication point. Until it exists,
Tempo's read path does not regard the block as live.

After a successful upload, the block-builder commits its Kafka offset. If the
block-builder crashes before completing the operation, it can replay the Kafka
records and reconstruct the block. Deterministic block IDs and span
deduplication make that replay safe.

At this point, the durability roles are:

```text
Kafka: immediate durable ingestion buffer
S3:    long-term durable trace storage
```

Kafka must retain the records long enough for block-builders to process them.
Once blocks are safely stored in S3, Kafka can eventually delete its older
copy according to its retention policy.

See the [Tempo object-storage documentation][object-storage].

## What happens after the first S3 upload

The initial ingestion journey is complete, but Tempo continues maintaining the
stored data:

- Backend workers maintain block metadata.
- Compaction combines smaller blocks into larger blocks.
- Retention eventually removes expired trace blocks.
- Queriers read the blocks when searches cover data older than the live-store
  window.

Grafana, query frontends, and queriers are not part of the ingestion route.
They participate when someone reads or searches the trace later.

## Compact mental model

```text
The SDK creates and exports spans.
Alloy processes and forwards them.
The distributor validates and partitions them.
Kafka makes ingestion immediately durable.
The live-store makes recent data queryable.
The block-builder turns spans into Parquet blocks.
S3 provides long-term trace storage.
```

For the complete current architecture, see Grafana's
[Tempo architecture overview][tempo-architecture].

[alloy-tracing]: https://grafana.com/docs/tempo/latest/set-up-for-tracing/instrument-send/set-up-collector/grafana-alloy/
[block-builder]: https://grafana.com/docs/tempo/latest/reference-tempo-architecture/components/block-builder/
[block-format]: https://grafana.com/docs/tempo/latest/reference-tempo-architecture/block-format/
[distributor]: https://grafana.com/docs/tempo/latest/reference-tempo-architecture/components/distributor/
[live-store]: https://grafana.com/docs/tempo/latest/reference-tempo-architecture/components/live-store/
[object-storage]: https://grafana.com/docs/tempo/latest/reference-tempo-architecture/object-storage/
[tempo-architecture]: https://grafana.com/docs/tempo/latest/reference-tempo-architecture/about-tempo-architecture/
[tempo-kafka]: https://grafana.com/docs/tempo/latest/reference-tempo-architecture/components/kafka/
