# Trace backend alternatives without mandatory Kafka

Kafka is not an OpenTelemetry requirement. Several tracing backends support
distributed deployments without it.

The important trade-off is that distributed-state complexity does not
disappear. It moves into another system, usually OpenSearch, ClickHouse,
Cassandra, or replicated local storage.

## Practical alternatives

| Backend | Primary storage | Kafka required? | Distributed approach |
|---|---|---:|---|
| Tempo 3 microservices | Object storage | Yes | Tempo components consume a Kafka stream |
| Jaeger v2 | OpenSearch, Elasticsearch, or Cassandra | No | Stateless collectors and query services share a distributed database |
| VictoriaTraces | Local disks or persistent volumes on storage nodes | No | `vtinsert`, `vtselect`, and sharded `vtstorage` nodes |
| OpenSearch Observability | OpenSearch indexes | No | Data Prepper instances write directly into an OpenSearch cluster |
| SigNoz | ClickHouse | No | SigNoz services use a distributed ClickHouse database |
| Uptrace | ClickHouse, PostgreSQL, and Redis | No | The application tier scales around those databases |
| Elastic Observability | Elasticsearch | No | OTLP or APM ingestion writes into Elasticsearch |

## Jaeger v2 with OpenSearch

Jaeger v2 with OpenSearch is a clear, mature alternative to distributed Tempo
without mandatory Kafka:

```text
Applications
    |
    v
Alloy / OpenTelemetry Collectors
    |
    v
Jaeger collectors
    |
    v
OpenSearch cluster
    ^
    |
Jaeger query replicas
    ^
    |
Grafana or Jaeger UI
```

Jaeger collectors and query services are stateless, so multiple instances can
run in parallel while sharing the OpenSearch backend. Kafka is optional:

```text
Direct:

Jaeger collectors -> OpenSearch
```

```text
Optional buffered path:

Jaeger collectors -> Kafka -> Jaeger ingesters -> OpenSearch
```

Current Jaeger recommends OpenSearch over Cassandra for large-scale production
because it provides stronger trace-search functionality.

References:

- [Jaeger architecture](https://www.jaegertracing.io/docs/2.20/architecture/)
- [Jaeger storage backends](https://www.jaegertracing.io/docs/2.20/storage/)

OpenSearch itself is a substantial distributed system. Operating it involves
decisions about:

- data and cluster-manager nodes
- index lifecycle and retention
- shard sizing
- replication
- JVM memory
- disk watermarks
- reindexing and schema upgrades
- snapshots and restoration

Jaeger can therefore remove Kafka without necessarily reducing the total
operational complexity. The result depends heavily on whether the organization
already operates OpenSearch successfully.

## VictoriaTraces

VictoriaTraces has an especially compact distributed architecture:

```text
Alloy
  |
  v
vtinsert replicas
  |
  v
vtstorage-1  vtstorage-2  vtstorage-3
  ^
  |
vtselect replicas
  ^
  |
Grafana
```

Each `vtstorage` node stores trace data directly in its filesystem directory.
VictoriaTraces does not require:

- Kafka
- object storage
- ClickHouse
- OpenSearch
- Cassandra

It accepts OTLP and provides Jaeger-compatible query APIs for Grafana.

Reference:

- [VictoriaTraces cluster architecture](https://docs.victoriametrics.com/victoriatraces/cluster/)

This is one of the closest matches to "distributed trace storage with fewer
external systems."

However, VictoriaTraces does not replicate data between storage nodes. If one
storage node is unavailable, queries requiring its data fail instead of
returning an incomplete result. Its documented stronger-availability approach
is to have a trace shipper such as an OpenTelemetry Collector duplicate spans
to two independent VictoriaTraces clusters:

```text
                       +-> VictoriaTraces cluster A
+------------------+   |
| Alloy/Collector  |---+
| buffer + fan-out |   |
+------------------+   +-> VictoriaTraces cluster B
```

The basic cluster has few dependencies, but this stronger HA topology requires
external duplication. VictoriaTraces is also newer than Jaeger or Tempo, so
production evaluation should explicitly test:

- upgrades and rollback
- storage-node failure
- query behavior during partial failure
- retention
- backup and restoration
- Grafana and API compatibility
- operational dashboards and alerts

## OpenSearch Observability

OpenSearch Data Prepper can receive OTLP traces and write them directly to
OpenSearch:

```text
Applications -> Alloy -> Data Prepper -> OpenSearch
                                         ^
                                         |
                              OpenSearch Dashboards
```

Kafka is optional. Data Prepper supports sources, processors, buffers, and
sinks, following a pipeline model similar to the OpenTelemetry Collector.

References:

- [OpenSearch Data Prepper](https://docs.opensearch.org/latest/data-prepper/)
- [OpenSearch telemetry ingestion](https://docs.opensearch.org/latest/observing-your-data/apm/configuring-telemetry-ingestion/)

This can be attractive when an organization already operates OpenSearch for
logs or search. Deploying an OpenSearch cluster solely to avoid Kafka may
increase rather than decrease the total operational burden.

## ClickHouse-based platforms

### SigNoz

SigNoz sends telemetry through its OpenTelemetry Collector layer into
ClickHouse:

```text
Applications -> OpenTelemetry Collector -> ClickHouse
                                             ^
                                             |
                                           SigNoz
```

It provides traces, metrics, logs, dashboards, alerts, and service maps as a
unified platform. Kafka is not a mandatory component in its documented
self-hosted architecture.

Reference:

- [SigNoz architecture](https://signoz.io/docs/architecture/)

### Uptrace

Uptrace follows a related model with several storage dependencies:

```text
ClickHouse -> high-volume telemetry
PostgreSQL -> users, projects, dashboards, and configuration
Redis      -> cache and background coordination
```

Reference:

- [Self-hosting Uptrace](https://uptrace.dev/get/hosted)

These platforms avoid mandatory Kafka, but a production ClickHouse cluster and
its coordination requirements are another form of distributed-storage
complexity.

## Elastic Observability

Elastic accepts OpenTelemetry data through its supported OTLP and APM ingestion
paths and stores it in Elasticsearch:

```text
Applications -> OpenTelemetry collection layer -> Elasticsearch
                                                   ^
                                                   |
                                              Kibana / Elastic UI
```

Kafka is not mandatory. Elasticsearch handles distributed storage, indexing,
replication, retention, and querying. This can be a good fit where an
organization already operates Elastic, but licensing, resource usage, shard
management, and the desired open-source posture must be evaluated.

Reference:

- [Using OpenTelemetry with Elastic APM](https://www.elastic.co/guide/en/apm/get-started/current/open-telemetry-elastic.html)

## The unavoidable trade-off

A durable distributed tracing system needs something to solve:

- replication
- sharding
- failure recovery
- backpressure
- concurrent writes
- retention
- query coordination

Different systems assign those responsibilities differently:

```text
Tempo:
Kafka + object storage

Jaeger:
OpenSearch, Elasticsearch, or Cassandra

SigNoz:
ClickHouse

VictoriaTraces:
local storage shards plus external duplication for stronger HA
```

Kafka can disappear, but the distributed-systems problem cannot.

## Practical direction for this study environment

Given the existing use of Grafana, Alloy, and the LGTM ecosystem:

1. Start with monolithic Tempo backed by reliable S3-compatible object storage
   if one Tempo failure domain is acceptable.
2. Measure whether independent scaling and high availability are genuinely
   required.
3. If monolithic Tempo is outgrown:
   - accept Kafka-compatible infrastructure and retain the native Tempo
     integration;
   - evaluate Jaeger v2 with OpenSearch when OpenSearch is already an
     organizational capability; or
   - evaluate VictoriaTraces when minimizing dependencies and using local
     storage are priorities, while carefully testing its availability model.

If the requirement is "highly available tracing without operating Kafka or
another distributed database," a managed tracing backend is the simpler
answer. A self-hosted distributed system always leaves the operator responsible
for the difficult stateful layer somewhere.
