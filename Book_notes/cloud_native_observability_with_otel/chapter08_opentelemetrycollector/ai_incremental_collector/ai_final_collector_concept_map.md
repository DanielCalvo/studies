# Final OpenTelemetry Collector concept map

This document summarizes the completed local Chapter 8 exercise. It describes
the configuration we built, not a universal production architecture.

## The Collector's central model

Telemetry pipelines have three principal stages:

```text
receiver -> zero or more processors -> one or more exporters
```

- A receiver gets telemetry into the Collector.
- A processor modifies, rejects, groups, or otherwise handles accepted
  telemetry.
- An exporter sends processed telemetry out of the Collector.
- `service.pipelines` activates and connects those configured components for a
  particular signal.

An extension is different. It adds a capability to the Collector process but
does not occupy a position in a telemetry pipeline.

## Complete local architecture

```text
                                   OpenTelemetry Collector

Go application ---------\
telemetrygen gRPC --------> OTLP receiver -------------------------------\
telemetrygen HTTP -------/                                                |
                                                                          |
                           traces                                         |
                           -------                                        |
                           memory limiter                                 |
                           -> Resource environment enrichment              |
                           -> span chapter attribute                       |
                           -> OTTL span-name transform                     |
                           -> drop reverse-operation span                  |
                           -> batch                                       |
                           -> fan-out -> debug terminal                    |
                                      -> OTLP JSON file                    |
                                                                          |
telemetrygen metrics ---> OTLP receiver --\                               |
                                         |                                |
memory and network ----> host metrics ----+-> metrics                     |
                                             memory limiter               |
                                             -> Resource enrichment       |
                                             -> drop connection metric    |
                                             -> batch                     |
                                             -> debug terminal            |
                                                                          |
OTLP log record -------> OTLP receiver -----> logs                        |
                                             memory limiter               |
                                             -> Resource enrichment       |
                                             -> batch                     |
                                             -> debug terminal            |
                                                                          |
Kubernetes/manual probe -> health-check extension -> HTTP :13133          |
                           (outside every telemetry pipeline)              |
```

Traces, metrics, and logs can share component configurations while remaining
separate signals with different data models and pipeline instances.

## The three pipelines

### Traces

```text
OTLP
  -> memory_limiter
  -> resource/add-local-environment
  -> attributes/add-study-chapter
  -> transform/prefix-span-name
  -> filter/drop-reverse-span
  -> batch
  -> debug
  -> file/local-traces
```

The Go application creates four spans. The Collector enriches and renames all
four, then removes one. Both exporters therefore receive these three processed
spans:

```text
8:read random characters
8:save characters
8:process random characters
```

The file and debug exporters are two destinations for the same post-processor
telemetry. The application exports the trace only once.

### Metrics

```text
OTLP -------------------\
                         -> memory_limiter
host metrics receiver --/   -> resource/add-local-environment
                              -> filter/drop-network-connections
                              -> batch
                              -> debug
```

The receiver actively scrapes memory plus network measurements for `eth0`.
The network scraper never creates `lo` data points. It does create
`system.network.connections`, but the metric filter removes that metric later.

The verified output contains five metric families and 14 data points:

```text
system.memory.usage
system.network.dropped
system.network.errors
system.network.io
system.network.packets
```

### Logs

```text
OTLP
  -> memory_limiter
  -> resource/add-local-environment
  -> batch
  -> debug
```

Enabling the logs pipeline allows OTLP log records to flow through the
Collector. It does not tail container standard output automatically.

## Passive and active receivers

```text
OTLP receiver
    Passive: another process pushes telemetry to ports 4317 or 4318.

host-metrics receiver
    Active: the Collector invokes configured scrapers every ten seconds.
```

OTLP over gRPC on `4317` and OTLP over HTTP on `4318` are two transports for
the same protocol and telemetry models. They are not different signals.

If Node Exporter and Prometheus already provide the required node metrics,
running an overlapping host-metrics receiver is usually unnecessary. In that
environment, the host-metrics work in this exercise should be understood as a
demonstration of an active receiver.

## Processor ordering

Processors execute in their listed order. Each sees the result produced by the
processors before it:

```text
create study.chapter
  -> read study.chapter to construct the span name
  -> compare the constructed name in the filter
```

Reversing those processors would change the result. The transform could not
read an attribute that had not been created, and a filter looking for
`8:reverse characters` would not match the original name
`reverse characters`.

The memory limiter belongs first so it can reject new telemetry before later
processors allocate more memory. Batch belongs near the end so reductions and
enrichment happen before telemetry is grouped for export.

## Selection, transformation, and filtering

These operations are related but not interchangeable:

```text
Receiver selection
    Decides what the receiver collects.
    Example: collect eth0 but never create lo measurements.

Transformation
    Changes telemetry that already exists.
    Example: rename a span with OTTL.

Processor filtering
    Drops telemetry after it has entered the pipeline.
    Example: discard system.network.connections.
```

A filter condition describes what to drop. It is not an allow-list unless the
condition is deliberately written to drop everything outside an allowed set.

## Resource attributes and record attributes

```text
deployment.environment.name = development
    Resource attribute describing the producing entity or environment.

study.chapter = 8
    Span attribute describing an individual operation in this exercise.
```

Resources are shared descriptions attached to groups of telemetry. Span, log
record, and metric data-point attributes describe lower-level records.

## Batching and memory limiting

Batching improves export efficiency:

```text
many small telemetry items -> fewer, larger export operations
```

It does not combine traces, metrics, and logs into one signal.

The memory limiter protects Collector availability:

```text
below 96 MiB           -> accept normally
above 96 MiB soft limit -> refuse new data with retryable errors
above 128 MiB hard limit -> also force Go garbage collection
```

It does not flush telemetry faster. Refused data is preserved only when the
preceding component and sender retry successfully; sustained overload can
still cause data loss.

## Exporter fan-out

One pipeline can have several exporters:

```text
processed trace -> debug exporter -> terminal
                `-> file exporter  -> traces.json
```

Fan-out does not make the file a durable backup queue for the debug exporter.
The file exporter serializes telemetry but supplies no indexing, querying,
dashboarding, or distributed retention.

## Extensions

The health-check extension exposes a Collector-level HTTP endpoint:

```text
GET :13133/ -> HTTP 200 while ready
```

It would mostly be used by Kubernetes readiness or liveness probes. A successful
response confirms basic Collector readiness, not end-to-end delivery to Tempo,
Loki, Prometheus, or another backend.

Defining the extension does not activate it. The reference in
`service.extensions` does.

## Mapping the concepts to Grafana Alloy

Alloy uses the upstream Collector engine for many `otelcol.*` components, but
its configuration is a component graph rather than a YAML
`service.pipelines` section. Components explicitly forward signals to the
`.input` fields of downstream components.

| Collector exercise | Current Alloy equivalent or alternative |
| --- | --- |
| `otlp` receiver | `otelcol.receiver.otlp` |
| `memory_limiter` | `otelcol.processor.memory_limiter` |
| attributes processor | `otelcol.processor.attributes` |
| transform processor | `otelcol.processor.transform` |
| filter processor | `otelcol.processor.filter` |
| batch processor | `otelcol.processor.batch` |
| debug exporter | `otelcol.exporter.debug` |
| file exporter | `otelcol.exporter.file` |
| Resource processor with a fixed value | No standalone `otelcol.processor.resource` is listed; use an OTTL resource statement in `otelcol.processor.transform`, or a specialized detector when appropriate |
| host-metrics receiver | No `otelcol.receiver.hostmetrics` is listed; Alloy commonly uses `prometheus.exporter.unix` and `prometheus.scrape`, or scrapes an existing Node Exporter |
| health-check extension | Alloy provides built-in `/-/ready` and `/-/healthy` HTTP endpoints rather than an `otelcol.extension.health_check` component |
| `service.pipelines` wiring | Direct `output` references between Alloy components |
| several exporters | Put several downstream `.input` references in an `output` list |

The absence of a one-to-one component does not mean Alloy lacks the capability.
It means migration may use a native Alloy or Prometheus component instead of
an upstream Collector wrapper.

Alloy's `/-/ready` endpoint reports whether the initial configuration has
loaded. `/-/healthy` summarizes component health, but Grafana explicitly warns
that `/-/healthy` is not suitable as a Kubernetes liveness probe: an unhealthy
external dependency may not be fixed by restarting Alloy.

For this user's existing Kubernetes design, a likely division is:

```text
application OTLP -> Alloy otelcol components -> Tempo or another OTLP backend
container logs   -> Alloy Loki components    -> Loki
Node Exporter    -> Alloy/Prometheus scrape  -> Prometheus-compatible storage
```

That is an architectural direction, not a generated production configuration.
Backend authentication, TLS, queues, retry policy, tenancy, and Kubernetes
placement still need deployment-specific decisions.

## What remains deliberately deferred

- Tempo, Loki, Prometheus, or Mimir backend configuration
- OTLP TLS and authentication
- Exporter sending queues, persistent storage, and retry tuning
- Kubernetes sidecar, DaemonSet, agent, and gateway placement
- Collector horizontal scaling
- Tail or probabilistic sampling
- Production memory sizing and `GOMEMLIMIT`
- Kubernetes Resource enrichment and RBAC
- An executable Alloy conversion
- End-to-end delivery monitoring

## Current references

- [Grafana Alloy OpenTelemetry components](https://grafana.com/docs/alloy/latest/reference/components/otelcol/)
- [Migrate from the OpenTelemetry Collector to Alloy](https://grafana.com/docs/alloy/latest/tasks/migrate/from-otelcol/)
- [Alloy HTTP readiness and health endpoints](https://grafana.com/docs/alloy/latest/reference/http/)
- [Alloy Unix exporter based on Node Exporter](https://grafana.com/docs/alloy/latest/reference/components/prometheus/prometheus.exporter.unix/)
- [Alloy OTLP receiver](https://grafana.com/docs/alloy/latest/reference/components/otelcol/otelcol.receiver.otlp/)
- [Alloy transform processor](https://grafana.com/docs/alloy/latest/reference/components/otelcol/otelcol.processor.transform/)
- [Alloy filter processor](https://grafana.com/docs/alloy/latest/reference/components/otelcol/otelcol.processor.filter/)
- [Alloy batch processor](https://grafana.com/docs/alloy/latest/reference/components/otelcol/otelcol.processor.batch/)

These links describe the current Alloy component model. Component availability
and stability can change, so a future production conversion should recheck the
documentation and validate the generated Alloy configuration.
