# Loki

Loki stores and queries logs shipped by Grafana Alloy. This homelab uses one
monolithic Loki replica with TSDB indexes and filesystem storage on a
`local-path` PVC.

For a broader survey of Helm values, deployment modes, storage backends, and
non-Helm configuration choices, see
[`ai_deployment_options.md`](ai_deployment_options.md).

## Pinned release

- Helm repository: `https://grafana.github.io/helm-charts`
- Helm release: `loki`
- Namespace: `monitoring`
- Chart: `grafana/loki`
- Chart version: `7.0.0`
- Loki version: `3.6.7`

Install or reconcile the release:

```bash
helm upgrade --kube-context hp --install loki grafana/loki \
  --version 7.0.0 \
  --namespace monitoring \
  --values values.yaml \
  --wait \
  --timeout 10m
```

The `loki` service is intentionally a cluster-internal `ClusterIP`. Grafana and
Alloy reach it at `http://loki.monitoring.svc.cluster.local:3100`.

## Storage and retention

The PVC is 10 GiB and the global retention period is seven days (`168h`).
Compactor retention is enabled. Loki does not delete data in response to disk
pressure, so check usage periodically:

```bash
kubectl -n monitoring get pvc
kubectl -n monitoring exec loki-0 -- df -h /var/loki
```

The `local-path` volume is tied to one node. It survives pod restarts, but this
deployment is neither highly available nor protected from loss of that node's
disk.

## Health and direct queries

```bash
kubectl -n monitoring get statefulset,pod,service,pvc
kubectl -n monitoring port-forward service/loki 3100:3100
curl -fsS http://127.0.0.1:3100/ready
curl -fsS -G http://127.0.0.1:3100/loki/api/v1/query_range \
  --data-urlencode 'query={namespace="image-resizer",app="image-resizer-api"}' \
  --data-urlencode 'limit=20'
```
