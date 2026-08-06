# Grafana Alloy

Alloy discovers the image-resizer pods through the Kubernetes API, tails their
container stdout/stderr, attaches a small set of Kubernetes labels, and writes
the original log lines to the Alloy gateway in the HP cluster. A separate
OpenTelemetry pipeline receives OTLP traces, batches them, and forwards them to
the same gateway.

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
  --kube-context opi \
  --version 1.10.0 \
  --namespace monitoring \
  --values values.yaml \
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

New logs use `cluster="opi"` to identify their source. This label must remain
the originating cluster name; HP Alloy forwards it unchanged to HP Loki.

## Cross-cluster metrics

The `prometheus.operator.servicemonitors.image_resizer` component selects
ServiceMonitors labeled `alloy: opi` in `monitoring` and `kube-system`. It
scrapes both Image Resizer application pods every 15 seconds, adds
`cluster="opi"`, and sends remote write to HP Alloy at
`http://192.168.1.232:9091/api/v1/metrics/write`.

Verification on 2026-08-05 discovered two targets and forwarded 12,874 samples
with zero failed, retried, or pending samples. OPI has no Prometheus server;
Alloy is the sole local scraper and forwards the samples to HP Prometheus via
the HP Alloy gateway. The preserved `prometheus: opi` label is retained only
for compatibility with the existing ServiceMonitor metadata.

The remote-write WAL is stored on ephemeral pod storage. A pod replacement can
lose buffered samples that have not yet reached HP Alloy.

The selected `kube-system/kubelet` ServiceMonitor scrapes each OPI kubelet's
authenticated HTTPS port `10250` at `/metrics` and `/metrics/cadvisor`. Those
sources provide kubelet volume/capacity metrics and cAdvisor pod/container
resource metrics respectively. The ServiceMonitor preserves the endpoint node
name as `node` and adds `cluster="opi"`; Alloy's ServiceMonitor discovery
includes both `monitoring` and `kube-system` namespaces.

On 2026-08-05, Alloy reported two healthy `job="kubelet"` targets. HP
Prometheus contained `machine_cpu_cores` and `machine_memory_bytes` for both
`opi1` and `opi2`, plus 31 and 26 `container_cpu_usage_seconds_total` series
for `opi1` and `opi2` respectively. This supplies the previously empty
Kubernetes Views Nodes kubelet/cAdvisor panels.

Alloy's chart-generated ServiceMonitor is also labeled `alloy: opi`, so the
same pipeline forwards Alloy self-metrics to HP Prometheus. This makes scrape
health and remote-write delivery observable without restoring an OPI
Prometheus server. On 2026-08-05, HP Prometheus reported one healthy
`job="alloy"` target, one `alloy_build_info` series, and one
`prometheus_remote_storage_samples_total` series for `cluster="opi"`.

Fields inside the application's JSON, including `request_id`, `duration_ms`,
dimensions, and byte counts, remain in the log body. They must not become Loki
labels because their high cardinality would create too many streams.

The tracing pipeline exposes these ports through Alloy's internal ClusterIP
Service:

- OTLP gRPC: `4317`
- OTLP HTTP: `4318`

It forwards traces over unencrypted OTLP gRPC to the HP Alloy gateway at
`192.168.1.232:4317`. The HP gateway then forwards them to Tempo through Tempo's
private ClusterIP Service. The cross-cluster link stays on the trusted home LAN.

The log pipeline sends Loki Push API requests to
`http://192.168.1.232:3100/loki/api/v1/push`. HP Alloy forwards them to its
private Loki Service. OPI's historical logs remain in OPI Loki; newly ingested
logs are stored in HP Loki after the Helm upgrade.

Inspect Alloy and its discovered components:

```bash
kubectl --context opi -n monitoring get deployment,pod,service \
  -l app.kubernetes.io/instance=alloy
kubectl --context opi -n monitoring logs deployment/alloy -c alloy
kubectl --context opi -n monitoring port-forward service/alloy 12345:12345
```

Then open `http://127.0.0.1:12345` to inspect component health and discovered
targets.
