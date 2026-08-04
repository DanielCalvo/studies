# Prometheus Node Exporter

Node exporter is installed in the `monitoring` namespace as a Kubernetes DaemonSet using the `prometheus-community/prometheus-node-exporter` Helm chart.

It runs one pod per node and exposes host-level metrics such as CPU, memory, filesystem, disk, load, and network counters.

## Install

```bash
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update prometheus-community

helm upgrade --install prometheus-node-exporter prometheus-community/prometheus-node-exporter \
  --version 4.55.0 \
  --namespace monitoring \
  --values ../node-exporter/values.yaml
```

## Prometheus Discovery

The chart creates a `ServiceMonitor` with this label:

```yaml
prometheus: homelab
```

The local Prometheus instance selects it through:

```yaml
serviceMonitorSelector:
  matchLabels:
    prometheus: homelab
```

The ServiceMonitor selects the chart-created Service by Service labels. The Service exposes the `metrics` port on `9100`.

The ServiceMonitor attaches `cluster="homelab"` to node-exporter samples so the
Kubernetes Grafana dashboards can use their cluster selector consistently.

## Resources

Node exporter is lightweight, so the chart is configured with small resources:

```yaml
resources:
  requests:
    cpu: 10m
    memory: 32Mi
  limits:
    memory: 64Mi
```

Current observed usage after install was about `1m` CPU and `3Mi` memory per pod.

## Current Shape

Expected DaemonSet shape on this two-node cluster:

```text
desired/current/ready: 2/2/2
nodes: opi1, opi2
image: quay.io/prometheus/node-exporter:v1.11.1
```

Prometheus active targets should be:

```text
192.168.1.201:9100
192.168.1.202:9100
```

## Useful Checks

```bash
kubectl -n monitoring get daemonset,pod,service,servicemonitor -l app.kubernetes.io/name=prometheus-node-exporter
kubectl -n monitoring top pod -l app.kubernetes.io/name=prometheus-node-exporter
curl -fsS 'http://192.168.1.220/api/v1/query?query=up%7Bjob%3D%22prometheus-node-exporter%22%7D' | jq '.data.result'
```

PromQL examples:

```promql
up{job="prometheus-node-exporter"}
node_memory_MemAvailable_bytes
node_filesystem_avail_bytes
rate(node_cpu_seconds_total[5m])
```
