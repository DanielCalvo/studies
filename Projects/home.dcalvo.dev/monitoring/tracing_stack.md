# Kubernetes application tracing stack

The first tracing pipeline accepts OpenTelemetry Protocol (OTLP) traces through
Grafana Alloy, batches them, stores them in Grafana Tempo, and makes them
queryable through Grafana.

```text
instrumented application -> Alloy -> Tempo -> Grafana
                             OTLP     storage   query
```

All components are internal `ClusterIP` services. This trusted homelab does not
use TLS or authentication between Alloy and Tempo.

## Pinned releases

- Alloy: `grafana/alloy` chart `1.10.0`, Alloy `v1.17.0`
- Tempo: `grafana-community/tempo` chart `2.2.3`, Tempo `2.10.7`
- Grafana: `grafana/grafana` chart `10.5.15`, Grafana `12.3.1`

The single-binary Tempo chart moved from the Grafana chart repository to the
Grafana Community chart repository in 2026. Add and update that repository with:

```bash
helm repo add grafana-community https://grafana-community.github.io/helm-charts
helm repo update
```

The pinned Tempo image publishes both `linux/amd64` and `linux/arm64` variants.

## Data flow and ports

Applications send OTLP to the internal Alloy Service:

- OTLP gRPC: `alloy.monitoring.svc.cluster.local:4317`
- OTLP HTTP: `http://alloy.monitoring.svc.cluster.local:4318`

Alloy sends batched traces over unencrypted OTLP gRPC to:

```text
tempo.monitoring.svc.cluster.local:4317
```

Grafana queries Tempo over HTTP at:

```text
http://tempo.monitoring.svc.cluster.local:3200
```

Tempo uses a 5 GiB `local-path` PVC and retains trace blocks for seven days.
The PVC is node-local and is the only copy of the trace data. This is suitable
for learning but is not highly available.

## Install or reconcile

Install Tempo before enabling Alloy's exporter, then reconcile Alloy and
Grafana:

```bash
helm upgrade --install tempo grafana-community/tempo \
  --version 2.2.3 \
  --namespace monitoring \
  --values monitoring/tempo/values.yaml \
  --wait \
  --timeout 10m

helm upgrade --install alloy grafana/alloy \
  --version 1.10.0 \
  --namespace monitoring \
  --values monitoring/alloy/values.yaml \
  --wait \
  --timeout 10m

helm upgrade --install grafana grafana/grafana \
  --version 10.5.15 \
  --namespace monitoring \
  --values monitoring/grafana/values.yaml \
  --wait \
  --timeout 10m
```

## Verify

Inspect the workloads, storage, Services, and ServiceMonitors:

```bash
kubectl -n monitoring get pod,service,pvc
kubectl -n monitoring get servicemonitor alloy tempo
kubectl -n monitoring logs deployment/alloy -c alloy
kubectl -n monitoring logs statefulset/tempo
```

Forward Tempo's internal HTTP port and check readiness:

```bash
kubectl -n monitoring port-forward service/tempo 13200:3200
curl -fsS http://127.0.0.1:13200/ready
```

Send five synthetic traces through Alloy. This pinned telemetry generator image
is a multi-architecture manifest that includes `linux/arm64`:

```bash
kubectl -n monitoring run tracing-smoke-test \
  --image=ghcr.io/open-telemetry/opentelemetry-collector-contrib/telemetrygen@sha256:cdee2a230cca4263f6595eee7350ef6908e70a6a66e92301a2f21fa8747dd9ba \
  --restart=Never \
  -- traces \
  --otlp-endpoint=alloy.monitoring.svc.cluster.local:4317 \
  --otlp-insecure \
  --service=tracing-smoke-test \
  --traces=5 \
  --child-spans=3 \
  --span-duration=250ms

kubectl -n monitoring logs pod/tracing-smoke-test
kubectl -n monitoring delete pod tracing-smoke-test
```

With the Tempo port-forward still running, search for the smoke-test traces:

```bash
curl -fsS --get \
  --data-urlencode 'q={ resource.service.name = "tracing-smoke-test" }' \
  http://127.0.0.1:13200/api/search
```

In Grafana at `http://192.168.1.221`, select **Explore**, choose the provisioned
Tempo data source, and search for traces by service name or TraceQL.

## Initial live verification

The complete pipeline was verified on 2026-07-24 with five synthetic traces and
three child spans per trace:

- Alloy accepted 20 spans over OTLP gRPC.
- Alloy exported all 20 spans to Tempo.
- Alloy reported zero failed or refused spans.
- Tempo returned all five traces in a TraceQL service-name search.
- Fetching one trace through Grafana returned its root and three child spans.
- Grafana reported its Prometheus, Loki, and Tempo data sources healthy.
- Prometheus reported Alloy, Loki, and Tempo with `up == 1`.

The instrumented image-resizer was then deployed and verified with its
post-deployment smoke suite:

- Nine resize requests created twenty application spans.
- A successful resize created the `POST /v1/resize` server span and child spans
  for upload reading, JPEG decoding, resizing, and encoding.
- Rejected requests recorded their stable application error codes.
- Loki completion logs contained the same trace and span IDs.
- Grafana returned a complete application trace through the Tempo data source
  and had bidirectional Tempo/Loki correlation provisioned.
- Alloy accepted and exported every application span without refusal.
- Prometheus reported both image-resizer scrape targets with `up == 1`.

## Current boundaries

- Tempo is one process and one replica, not an HA deployment.
- Trace storage is a node-local PVC.
- Trace retention is seven days.
- Metrics-generator, service graphs, span metrics, and sampling experiments are
  intentionally deferred.
- Grafana correlates image-resizer Tempo spans with Loki completion logs through
  the shared trace ID. The trace ID remains a JSON field rather than a Loki
  stream label.

## Image-resizer tracing dashboard

The version-controlled Grafana payload at
`image-resizer-api/grafana/image-resizer-tracing.json` provides this
metrics-to-traces learning workflow:

```text
Prometheus trends -> suspicious behavior -> Tempo trace -> Loki log
```

It combines request rates, latency percentiles, and processing-stage histograms
from Prometheus with searchable, clickable Tempo trace tables. TraceQL metrics
are not used because Tempo's metrics-generator remains disabled; this also
avoids duplicating the application's existing Prometheus metrics.

Publish or update it from the repository root:

```bash
curl --fail-with-body --silent --show-error \
  --user admin:admin \
  --header 'Content-Type: application/json' \
  --data-binary @image-resizer-api/grafana/image-resizer-tracing.json \
  http://192.168.1.221/api/dashboards/db
```

Open:

```text
http://192.168.1.221/d/image-resizer-tracing/image-resizer-tracing
```

The published dashboard was validated on 2026-07-24 while the traffic generator
was active:

- Grafana stored all eleven dashboard panels.
- The recent-trace table returned its configured limit of fifty traces.
- The default `500ms` slow-trace table returned thirty traces.
- The rejected/failed table returned ten traces.
- Grafana attached an internal Tempo link to every returned Trace ID.
- Prometheus observed about two successful requests per second and returned
  decode, resize, and encode p95 latency series.
