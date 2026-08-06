# Grafana Alloy telemetry gateway

This Alloy instance is the LAN-facing telemetry gateway for the central
observability stack in the HP cluster. It accepts OTLP traces, Loki log streams,
and selected OPI Prometheus remote-write samples, then forwards them to the
private Tempo, Loki, and Prometheus Services.

## Pinned release

- Helm repository: `https://grafana.github.io/helm-charts`
- Helm release: `alloy`
- Namespace: `monitoring`
- Chart: `grafana/alloy`
- Chart version: `1.10.0`
- Alloy version: `v1.17.0`

## Endpoints

- OPI applications continue sending traces to their local OPI Alloy Service.
- OPI Alloy forwards traces to `192.168.1.232:4317`.
- OPI Alloy forwards Loki Push API log requests to
  `http://192.168.1.232:3100/loki/api/v1/push`.
- HP Alloy forwards traces to
  `tempo.monitoring.svc.cluster.local:4317` inside the HP cluster.
- HP Alloy forwards logs to
  `http://loki.monitoring.svc.cluster.local:3100/loki/api/v1/push` inside the
  HP cluster.
- `monitoring/alloy-otlp` is the only LAN-facing Alloy Service.
- The Alloy UI and self-metrics stay on the internal `monitoring/alloy`
  ClusterIP Service.

The LoadBalancer exposes only ports `4317`, `3100`, and `9091`, accepts only the
OPI node addresses and the development workstation, and uses
`externalTrafficPolicy: Local` to preserve the source address. Verify the
allowlist from both an allowed and a denied LAN client after installation.

## Cross-cluster metrics

OPI Alloy sends only the `image-resizer-api` ServiceMonitor, selected by
`alloy: opi`, to `http://192.168.1.232:9091/api/v1/metrics/write`. The HP
receiver fans every batch to both stable Prometheus pod DNS names:

- `prometheus-prometheus-0.prometheus-operated.monitoring.svc.cluster.local:9090`
- `prometheus-prometheus-1.prometheus-operated.monitoring.svc.cluster.local:9090`

The HP Prometheus replicas have `enableRemoteWriteReceiver: true`. Verification
on 2026-08-05 sent 12,874 samples to each replica; both queues drained to zero
with zero failed or retried samples. The `cluster="opi"`,
`namespace="image-resizer"`, and `job="image-resizer-api"` labels are retained.

Alloy remote-write WALs remain on ephemeral `/tmp/alloy` storage. A pod
replacement can lose buffered samples that have not reached the next hop; this
is a known limitation of the bounded home-lab experiment.

## Install

Before changing the HP cluster, verify its identity:

```bash
kubectl --context hp config view --minify \
  -o jsonpath='{.clusters[0].cluster.server}{"\n"}'
kubectl --context hp get nodes -o wide
```

Install or reconcile Alloy:

```bash
helm upgrade --install alloy grafana/alloy \
  --kube-context hp \
  --version 1.10.0 \
  --namespace monitoring \
  --values values.yaml \
  --wait \
  --timeout 10m
```

Do not change the OPI Alloy destination until HP Alloy is healthy and the
LoadBalancer has received `192.168.1.232`.

## Inspect

```bash
kubectl --context hp -n monitoring get deployment,pod,service \
  -l app.kubernetes.io/instance=alloy -o wide
kubectl --context hp -n monitoring get service alloy-otlp -o wide
kubectl --context hp -n monitoring logs deployment/alloy -c alloy
```
