# Incremental OpenTelemetry Collector exercise

This directory follows the Chapter 8 study plan one small configuration change
at a time.

## Current state: completed incremental exercise

The Collector accepts traces, metrics, and logs through separate signal
pipelines. It now demonstrates passive and active receivers, ordered
processors, enrichment, OTTL transformation, filtering, batching, exporter
fan-out, memory protection, and a health-check extension:

```text
OTLP receiver
    |-- traces -> memory limiter -> enrich -> transform -> filter -> batch
    |                                                        |-> debug
    |                                                        `-> file
    |-- metrics -> memory limiter -> enrich -> filter -> batch -> debug
    `-- logs -> memory limiter -> enrich -> batch -> debug

host metrics -> metrics pipeline
health check -> Collector extension outside the pipelines
```

`trace-sender` is a real Go application copied from the Chapter 4 tracing
exercise. Its spans leave the application through an OTLP/gRPC exporter:

```text
Go application -> OTLP/gRPC :4317 -> Collector -> debug and file exporters
```

See [ai_final_collector_concept_map.md](ai_final_collector_concept_map.md) for
the complete signal routes and the current Grafana Alloy mapping.

## Run the exercise

Start only the Collector:

```bash
docker compose up --detach collector
```

Send one test trace using OTLP over gRPC:

```bash
docker compose run --rm trace-generator
```

Send another test trace using OTLP over HTTP:

```bash
docker compose run --rm trace-generator-http
```

Send one test metric using OTLP over gRPC:

```bash
docker compose run --rm metrics-generator
```

Send one test log record using OTLP over gRPC:

```bash
docker compose run --rm logs-generator
```

To make Collector-side trace batching visible, run the trace generator several
times within five seconds:

```bash
docker compose run --rm trace-generator
docker compose run --rm trace-generator
docker compose run --rm trace-generator
```

Alternatively, run the real Go application from another terminal:

```bash
cd trace-sender
go run .
```

Inspect what the debug exporter printed:

```bash
docker compose logs collector
```

Check the Collector's operational health endpoint:

```bash
curl http://localhost:13133/
```

Inspect the trace export written as OTLP JSON:

```bash
less collector-output/traces.json
```

Stop and remove the exercise containers and network:

```bash
docker compose down
```

The batch processor waits up to five seconds before forwarding telemetry. The
timeout is a maximum periodic flush interval, not a new five-second countdown
for every item. Items of the same signal that arrive during an interval may
therefore appear in one larger debug export. It never combines traces, metrics,
and logs with one another.

The metrics generator produces a `Metrics` entry, the logs generator produces a
`Logs` entry, and either trace source produces a `Traces` entry after its
pipeline's batch processor forwards the data.

The Go application terminal should still contain only its ordinary
random-character messages. It creates four spans, but the Collector drops the
reverse-operation span. The debug and file exporters therefore receive three
spans. Every surviving span should contain the centrally added attribute:

```text
study.chapter: Str(8)
```

The Go application does not set this attribute.

All three signals should contain this in their Resource section:

```text
deployment.environment.name: Str(development)
```

The Resource value describes the telemetry producer. It is not repeated as an
individual span, metric data-point, or log-record attribute.

The Collector also gathers memory and `eth0` network metrics every ten seconds
without a generator. Its network receiver does not collect `lo`, and a
downstream filter removes `system.network.connections`.

The synthetic generators remain available for comparing gRPC and HTTP. Their
exported trace structure is equivalent regardless of which transport carries
OTLP.

## Deliberately absent

This local exercise deliberately does not contain:

- Logs emitted by the Go application or tailed from standard output
- A telemetry backend
- TLS, authentication, persistent queues, or production retry tuning
- An executable Alloy configuration
- A Kubernetes deployment
- Sampling
