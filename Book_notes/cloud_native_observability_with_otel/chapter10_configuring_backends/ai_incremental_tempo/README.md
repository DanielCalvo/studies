# Incremental Tempo exercise

This exercise studies Grafana Tempo as the tracing-backend case study for
Chapter 10.

Focused references:

- [Tempo `target` values](ai_tempo_target_reference.md)
- [Tempo storage backends](ai_tempo_storage_backends.md)
- [On-premises object storage for Tempo](ai_tempo_on_premises_object_storage.md)
- [Why use a Collector before Tempo?](ai_why_use_a_collector_before_tempo.md)
- [Alloy's component graph](ai_alloy_component_graph.md)
- [Tempo span metrics explained](ai_tempo_span_metrics_explained.md)
- [Learning-to-production Tempo architecture map](ai_production_tempo_architecture_map.md)
- [Trace backend alternatives without mandatory Kafka](ai_trace_backend_alternatives_without_kafka.md)
- [How a trace travels through Alloy and distributed Tempo to S3](ai_distributed_trace_write_path_to_s3.md)

## Current state: Step 8 complete

The Minikube cluster contains one monolithic Tempo instance, a completed
trace-producer Job, one Grafana instance, and a persistent volume for Tempo's
local backend. Alloy now sits between the producer and Tempo:

```text
ConfigMap (tempo.yaml)
          |
          v
Deployment -> Pod -> Tempo process
                       |
                       +-> PVC: WAL and completed backend blocks
                       |
                       +-> emptyDir: recent live-store runtime

Tempo Service -> HTTP API :3200
              -> OTLP/gRPC :4317
              -> OTLP/HTTP :4318

trace-producer Job -> Alloy Service :4317
                          |
                          | receive -> batch -> export
                          v
                    Tempo Service :4317

Grafana -> Tempo Service :3200 -> Tempo query API

Tempo metrics generator -> remote write -> Prometheus :9090
                                             ^
                                             |
                                      Grafana queries metrics
```

Tempo stores and queries the traces. Grafana is a separate client that asks
Tempo for trace data and renders the results. Alloy receives and processes
telemetry before forwarding it. None of the Services makes its ports directly
available on the host.

## Inspect the deployment

```bash
kubectl get all,configmap,pvc -n chapter10-tempo
kubectl logs -n chapter10-tempo deployment/tempo
kubectl logs -n chapter10-tempo job/trace-producer
kubectl logs -n chapter10-tempo deployment/grafana
kubectl logs -n chapter10-tempo deployment/alloy
kubectl logs -n chapter10-tempo deployment/prometheus
kubectl logs -n chapter10-tempo job/service-graph-producer
```

Temporarily forward the Tempo HTTP API to the host:

```bash
kubectl port-forward -n chapter10-tempo service/tempo 3200:3200
```

In another terminal:

```bash
curl http://localhost:3200/ready
```

Expected response:

```text
ready
```

## Retrieve the Step 2 trace

The trace-producer logs contain:

```text
trace_id=<32-character trace ID>
```

With the port-forward still running, use that ID to retrieve the trace:

```bash
curl -H 'Accept: application/json' \
  "http://localhost:3200/api/traces/<trace-id>"
```

The returned JSON contains the resource, instrumentation scope, and span that
the producer exported. This proves the complete direct path:

```text
produce span -> OTLP/gRPC -> Tempo distributor -> retrieve by trace ID
```

The current persistence-test trace ID is:

```text
98c19333f8ca70c77d945f19e1219904
```

It was retrieved successfully before and after deleting the Tempo Pod. The
replacement Pod mounted the same PVC and found the completed trace block.

## Retrieve the trace sent through Alloy

The current producer sends to `alloy:4317`, and Alloy exports to `tempo:4317`.
Its trace ID is:

```text
215531b693001a0599449cdb76769246
```

With Tempo's port forwarded, retrieve it as before:

```bash
curl -H 'Accept: application/json' \
  "http://localhost:3200/api/traces/215531b693001a0599449cdb76769246"
```

The span is named `send one trace through Alloy to Tempo` and has
`study.step=tempo-via-alloy`.

Alloy also exposes a debugging UI and operational endpoints. Forward its HTTP
port:

```bash
kubectl port-forward -n chapter10-tempo service/alloy 12345:12345
```

Then open `http://localhost:12345` to inspect the component graph, or check:

```bash
curl http://localhost:12345/-/ready
curl http://localhost:12345/-/healthy
curl http://localhost:12345/metrics
```

## Inspect Tempo's attribute-size limit

Tempo now has this default ingestion override:

```yaml
overrides:
  defaults:
    ingestion:
      max_attribute_bytes: 64
```

The current producer adds a 128-byte value under
`study.oversized_attribute`. Tempo accepts the span but stores only the first
64 bytes. The current trace ID is:

```text
4fa23ed5f2d4e8b54b6f066350caf55a
```

Retrieve it through Tempo's API and inspect the attribute to see the truncated
value. Tempo's operational evidence is available on its metrics endpoint:

```text
tempo_distributor_attributes_truncated_total{scope="span",tenant="single-tenant"} 1
```

The `scope` label identifies which part of the OpenTelemetry data contained the
oversized attribute. The trace itself was not rejected.

## Query metrics derived from the trace

Tempo now enables the `span-metrics` and `service-graphs` processors. Its
metrics generator writes samples to Prometheus every five seconds:

```text
span -> Tempo metrics generator -> Prometheus remote write receiver
```

Forward Prometheus to the host:

```bash
kubectl port-forward -n chapter10-tempo service/prometheus 9090:9090
```

Open `http://localhost:9090` or query these metrics through Grafana's
provisioned `Prometheus` data source:

```promql
traces_spanmetrics_calls_total{
  service="chapter10-trace-producer"
}
```

```promql
traces_spanmetrics_latency_count{
  service="chapter10-trace-producer"
}
```

The current derived metric has an exemplar whose `traceID` is:

```text
2d7de15d0253a916164d701a635b5b4e
```

Grafana's Prometheus data source maps that exemplar label to the Tempo data
source, allowing metric-to-trace navigation.

The original attribute-limit trace has one `SPAN_KIND_INTERNAL` span from one
service, so it cannot produce a service-graph edge. The separate
service-graph-producer below supplies the cross-service relationship that this
processor requires.

## Inspect the derived service-graph edge

The service-graph producer emits one trace with this structure:

```text
greeting-service: GET /greet                 SERVER
└── greeting-service: call formatting service CLIENT
    └── formatting-service: POST /format      SERVER
```

Its current trace ID is:

```text
782062d1f5ec21ea66590000ff581394
```

The client span context becomes the remote parent of the formatting server
span. This models the propagation that real HTTP instrumentation would perform
through request headers.

With Prometheus forwarded to `localhost:9090`, run:

```promql
traces_service_graph_request_total{
  client="greeting-service",
  server="formatting-service"
}
```

The expected value is `1`. Tempo also generates a server-latency histogram:

```promql
traces_service_graph_request_server_seconds_count{
  client="greeting-service",
  server="formatting-service"
}
```

Its histogram exemplar contains the same trace ID, providing another
metric-to-trace link. The Grafana Tempo data source is already configured to
read service-graph metrics from Prometheus.

## Explore the trace in Grafana

Forward Grafana's HTTP port to the host:

```bash
kubectl port-forward -n chapter10-tempo service/grafana 3000:3000
```

Then open:

```text
http://localhost:3000/explore
```

The local study instance permits anonymous access. This is deliberately
convenient for Minikube and is not an appropriate production authentication
configuration.

Select the provisioned `Tempo` data source. In its TraceQL query editor, either:

- paste the trace ID from the producer Job to retrieve that exact trace, or
- search by its resource attribute:

```traceql
{ resource.service.name = "chapter10-trace-producer" }
```

Open the result to see its span, resource attributes, `study.step` span
attribute, and span event. The trace still resides in Tempo; Grafana only
queries and presents it.

This first data-source configuration uses ordinary completed HTTP queries.
Optional result streaming remains disabled because it also requires a Tempo
server setting; streaming is not necessary for querying or viewing traces.
Grafana's data-source health test may consequently log a failed optional gRPC
streaming connection to port `3200` even while reporting `Data source is
working`. The successfully returned TraceQL result is the relevant check for
the non-streaming path used in this step.

After a Tempo restart, exact trace-ID lookup can find a newly persisted trace
before TraceQL search does. Tempo's current default
`query_frontend.search.query_backend_after` is `15m`: searches covering the
most recent 15 minutes are sent to the live-store rather than the backend to
avoid querying both paths. Our live-store workspace is deliberately temporary,
but the completed backend block is persistent. Once the trace crosses that
boundary, TraceQL includes the backend block. This does not affect immediate
exact lookup by trace ID.

## Expected idle startup messages

Tempo 3.0.2 may log warnings while this deliberately minimal instance is idle:

- it automatically connects the in-process query and backend workers because
  we did not configure separate microservices addresses
- the backend worker can report `no jobs found` before any trace blocks exist
- WAL replay can ignore its own `blocks` directory as an unowned entry

These messages do not prevent readiness. The important evidence is that the Pod
is `Ready`, both OTLP servers and the metrics-generator module started,
`/ready` responds, and traces and derived metrics can be retrieved.

Object storage has not been introduced yet. Prometheus uses `emptyDir`, so its
metrics disappear with its Pod. Alloy currently performs only OTLP reception,
batching, and export; enrichment, filtering, sampling, and fan-out have not
been added.
