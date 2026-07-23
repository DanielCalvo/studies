# Grafana

Grafana is installed in the `monitoring` namespace with the `grafana/grafana` Helm chart.

It is exposed on the home LAN with MetalLB:

```text
http://192.168.1.221/
```

## Commands

```bash
helm repo add grafana https://grafana.github.io/helm-charts
helm repo update grafana

helm upgrade --install grafana grafana/grafana \
  --version 10.5.15 \
  --namespace monitoring \
  --values ../grafana/values.yaml
```

## Service

The chart creates a `LoadBalancer` Service:

```yaml
service:
  type: LoadBalancer
  port: 80
  targetPort: 3000
  annotations:
    metallb.io/loadBalancerIPs: 192.168.1.221
```

## Persistence

Grafana uses a `ReadWriteOnce` PVC from the cluster default `local-path` StorageClass for `/var/lib/grafana`.

```yaml
persistence:
  enabled: true
  type: pvc
  accessModes:
    - ReadWriteOnce
  size: 5Gi
```

The deployment strategy is `Recreate` because this setup uses a single local-path PVC and Grafana stores state in SQLite. This avoids two Grafana pods touching the same database during chart upgrades.

The chart's default `initChownData` init container is disabled because it fails against the existing local-path PVC. Grafana can already read/write the PVC.

```yaml
deploymentStrategy:
  type: Recreate

initChownData:
  enabled: false
```

## Resources

```yaml
resources:
  requests:
    cpu: 25m
    memory: 128Mi
  limits:
    memory: 256Mi
```

## Login

Initial credentials:

```text
username: admin
password: admin
```

This is intentionally simple for the homelab. Move this to an existing Kubernetes Secret later if desired.

## Datasource

The Prometheus datasource points at the in-cluster Prometheus LoadBalancer Service:

```text
http://prometheus-lb.monitoring.svc.cluster.local
```

The datasource UID is pinned to `prometheus` so downloaded dashboards can reference it reliably:

```yaml
datasources:
  datasources.yaml:
    apiVersion: 1
    datasources:
      - name: Prometheus
        type: prometheus
        uid: prometheus
        access: proxy
        url: http://prometheus-lb.monitoring.svc.cluster.local
        isDefault: true
```

## Dashboards

Dashboards are pinned by Grafana.com ID and revision rather than storing large JSON exports locally.

Configured dashboards:

```yaml
dashboards:
  default:
    kube-state-metrics-overview:
      # Grafana dashboard: Kube-State-Metrics Overview
      gnetId: 25091
      revision: 1
      datasource:
        - name: DS_PROMETHEUS
          value: prometheus
    node-exporter-full:
      # Grafana dashboard: Node Exporter Full
      gnetId: 1860
      revision: 45
      datasource:
        - name: ds_prometheus
          value: prometheus
    prometheus-2-overview:
      # Grafana dashboard: Prometheus 2.0 Overview
      gnetId: 3662
      revision: 2
      datasource:
        - name: DS_THEMIS
          value: prometheus
    kubernetes-views-global:
      # Grafana dashboard: Kubernetes / Views / Global
      gnetId: 15757
      revision: 43
    kubernetes-views-namespaces:
      # Grafana dashboard: Kubernetes / Views / Namespaces
      gnetId: 15758
      revision: 46
    kubernetes-views-nodes:
      # Grafana dashboard: Kubernetes / Views / Nodes
      gnetId: 15759
      revision: 40
    kubernetes-views-pods:
      # Grafana dashboard: Kubernetes / Views / Pods
      gnetId: 15760
      revision: 39
```

Dashboard URLs:

```text
Kube-State-Metrics Overview  /d/ksm-overview/kube-state-metrics-overview
Node Exporter Full           /d/rYdddlPWk/node-exporter-full
Prometheus 2.0 Overview      /d/a088dcaa-fc31-4c96-9c28-a4ed8fe2b648/prometheus-2-0-overview
Kubernetes / Views / Global      /d/k8s_views_global/kubernetes-views-global
Kubernetes / Views / Namespaces  /d/k8s_views_ns/kubernetes-views-namespaces
Kubernetes / Views / Nodes       /d/k8s_views_nodes/kubernetes-views-nodes
Kubernetes / Views / Pods        /d/k8s_views_pods/kubernetes-views-pods
```

Source dashboards on Grafana.com:

```text
https://grafana.com/grafana/dashboards/25091-kube-state-metrics-overview/
https://grafana.com/grafana/dashboards/1860-node-exporter-full/
https://grafana.com/grafana/dashboards/3662-prometheus-2-0-overview/
https://grafana.com/grafana/dashboards/15757-kubernetes-views-global/
https://grafana.com/grafana/dashboards/15758-kubernetes-views-namespaces/
https://grafana.com/grafana/dashboards/15759-kubernetes-views-nodes/
https://grafana.com/grafana/dashboards/15760-kubernetes-views-pods/
```

## Useful Checks

```bash
kubectl -n monitoring get pod,svc,pvc -l app.kubernetes.io/instance=grafana
curl -fsS -u admin:admin http://192.168.1.221/api/datasources/name/Prometheus | jq '{name, uid, url}'
curl -fsS -u admin:admin 'http://192.168.1.221/api/search?query=' | jq '[.[] | {title, uid, url}]'
```
