# kube-state-metrics

kube-state-metrics runs in OPI and exposes Kubernetes object-state metrics for
the central HP monitoring stack. OPI Alloy scrapes it locally and remote-writes
the samples to HP; OPI does not run a Prometheus server.

## Pinned release

- Helm release: `kube-state-metrics`
- Namespace: `monitoring`
- Chart: `prometheus-community/kube-state-metrics` `7.5.1`
- App/image: kube-state-metrics `v2.19.1`
- Workload: one Deployment replica
- Service: `kube-state-metrics.monitoring.svc.cluster.local:8080`

The chart creates a ServiceMonitor labeled `alloy: opi`. Its HTTP endpoint uses
`honorLabels: true`, a 30-second interval, and attaches `cluster="opi"`.

## Reconcile

```bash
helm upgrade --install kube-state-metrics \
  prometheus-community/kube-state-metrics \
  --kube-context opi \
  --version 7.5.1 \
  --namespace monitoring \
  --values opi-cluster/monitoring/kube_state_metrics/values.yaml \
  --wait \
  --timeout 10m
```

## Verify

```bash
kubectl --context opi -n monitoring get \
  deployment,pod,service,servicemonitor \
  -l app.kubernetes.io/name=kube-state-metrics
```

Both HP Prometheus replicas returned two
`kube_node_info{cluster="opi"}` series on 2026-08-05.
