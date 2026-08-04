# Grafana Tempo

Tempo is installed in monolithic mode for a lightweight tracing backend on the
arm64 homelab. Alloy is the OTLP collector; applications should normally send
traces to Alloy rather than directly to Tempo.

## Pinned release

- Helm repository: `https://grafana-community.github.io/helm-charts`
- Helm release: `tempo`
- Namespace: `monitoring`
- Chart: `grafana-community/tempo`
- Chart version: `2.2.3`
- Tempo version: `2.10.7`

Install or reconcile:

```bash
helm upgrade --kube-context hp --install tempo grafana-community/tempo \
  --version 2.2.3 \
  --namespace monitoring \
  --values values.yaml \
  --wait \
  --timeout 10m
```

## Deployment choices

- One monolithic Tempo replica
- OTLP gRPC and HTTP receivers enabled internally
- Local filesystem trace backend on a 5 GiB `local-path` PVC
- Seven-day retention
- No multitenancy, authentication, TLS, or external LoadBalancer
- Usage reporting and metrics-generator disabled
- Prometheus ServiceMonitor labeled `prometheus: homelab`

The PVC is node-local and retained when the StatefulSet is removed. Trace data
is therefore durable across pod restarts on its owning node but is not highly
available.

## Initial verification

On 2026-07-24, an arm64 `telemetrygen` pod sent five traces containing 20 spans
to Alloy over OTLP gRPC. Alloy accepted and exported all 20 with no failed or
refused spans. Tempo returned all five traces, and Grafana successfully queried
one complete four-span trace through the provisioned `tempo` data source.
