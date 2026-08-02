# Grafana Alloy

Alloy discovers the image-resizer pods through the Kubernetes API, tails their
container stdout/stderr, attaches a small set of Kubernetes labels, and writes
the original log lines to Loki. A separate OpenTelemetry pipeline receives OTLP
traces, batches them, and forwards them to Tempo.

For a broader survey of Helm values, controller topologies, telemetry pipelines,
state, scaling, and non-Helm installation choices, see
[`ai_deployment_options.md`](ai_deployment_options.md).

## Pinned release

- Helm repository: `https://grafana.github.io/helm-charts`
- Helm release: `alloy`
- Namespace: `monitoring`
- Chart: `grafana/alloy`
- Chart version: `1.10.0`
- Alloy version: `v1.17.0`

Install or reconcile the release:

```bash
helm upgrade --install alloy grafana/alloy \
  --version 1.10.0 \
  --namespace monitoring \
  --values monitoring/alloy/values.yaml \
  --wait \
  --timeout 10m
```

The first pipeline deliberately keeps only pods in namespace `image-resizer`
whose `app.kubernetes.io/name` label is `image-resizer-api`. It stores these
bounded labels in Loki:

- `cluster`
- `namespace`
- `app`
- `service_name`
- `pod`
- `container`
- `node`
- `job`

Fields inside the application's JSON, including `request_id`, `duration_ms`,
dimensions, and byte counts, remain in the log body. They must not become Loki
labels because their high cardinality would create too many streams.

The tracing pipeline exposes these ports through Alloy's internal ClusterIP
Service:

- OTLP gRPC: `4317`
- OTLP HTTP: `4318`

It forwards traces over unencrypted OTLP gRPC to
`tempo.monitoring.svc.cluster.local:4317`. Both links remain inside the trusted
home-lab cluster.

Inspect Alloy and its discovered components:

```bash
kubectl -n monitoring get deployment,pod,service -l app.kubernetes.io/instance=alloy
kubectl -n monitoring logs deployment/alloy -c alloy
kubectl -n monitoring port-forward service/alloy 12345:12345
```

Then open `http://127.0.0.1:12345` to inspect component health and discovered
targets.
