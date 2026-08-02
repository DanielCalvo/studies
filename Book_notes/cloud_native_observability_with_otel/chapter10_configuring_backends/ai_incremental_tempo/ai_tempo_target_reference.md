# Tempo `target` reference

The Tempo container contains the code for many internal components. The
`target` setting determines which component or group of components the Tempo
process starts.

For Tempo 3.0, `target` accepts the following values:

| Target | Role |
| --- | --- |
| `all` | Runs the complete monolithic Tempo backend in one process. This is the default. |
| `distributor` | Receives and validates incoming traces, then forwards them into the write path. |
| `block-builder` | Consumes traces from Kafka and writes persistent trace blocks. |
| `live-store` | Holds and serves recently ingested traces. |
| `metrics-generator` | Generates span metrics and service-graph metrics. |
| `query-frontend` | Accepts queries, divides the work, and combines the results. |
| `querier` | Searches recent data and stored trace blocks. |
| `backend-scheduler` | Creates and coordinates backend work. |
| `backend-worker` | Executes scheduled backend work such as compaction and retention. |

The individual targets correspond roughly to these functional areas:

```text
Write path:
  distributor
  block-builder
  live-store
  metrics-generator (optional)

Read path:
  query-frontend
  querier

Storage maintenance:
  backend-scheduler
  backend-worker
```

## One binary, different jobs

A distributed deployment can use the same Tempo image for several Kubernetes
workloads and give each workload a different target:

```yaml
# Distributor Deployment
target: distributor
```

```yaml
# Live-store StatefulSet
target: live-store
```

```yaml
# Querier Deployment
target: querier
```

Conceptually, those workloads run:

```text
grafana/tempo:3.0.2 + target: distributor
grafana/tempo:3.0.2 + target: querier
grafana/tempo:3.0.2 + target: live-store
```

The image is the same in each case. The `target` selects which subset of the
Tempo binary starts in that particular process and Pod.

## Things that are not targets

Some parts of a complete Tempo system are dependencies or configuration rather
than executable targets:

- Kafka is an external dependency used by Tempo 3 microservices mode.
- Object storage is a configured storage backend.
- Memberlist is a supporting internal module.
- Overrides are configuration for limits and tenant-specific behavior.
- HTTP and internal gRPC servers start as dependencies of the relevant targets.

## Removed Tempo 2 targets

Tempo 3 removed these older target names:

- `ingester`
- `compactor`
- `scalable-single-binary`

Their responsibilities were reorganized among Tempo 3's live store, block
builder, backend scheduler, and backend worker. Older Tempo examples may still
contain these names, so the example's version must be checked before reusing
such configuration.

## Why this exercise uses `all`

The current exercise uses:

```yaml
target: all
```

This is the appropriate choice for learning because it provides a complete
Tempo backend in one process without requiring Kafka or separately deployed
Tempo components.

Individual targets become relevant when deliberately constructing a
microservices deployment with Kafka, shared object storage, Kubernetes
Services, and components that can be scaled independently.

## Current reference

- [Tempo command-line flags and valid targets](https://grafana.com/docs/tempo/latest/set-up-for-tracing/setup-tempo/command-line-flags/)
- [Tempo deployment modes](https://grafana.com/docs/tempo/latest/reference-tempo-architecture/deployment-modes/)
