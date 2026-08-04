# kube-state-metrics install notes

Installed `kube-state-metrics` with the `prometheus-community/kube-state-metrics` Helm chart.

## Commands used

```bash
helm --kube-context hp repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update prometheus-community
helm install kube-state-metrics prometheus-community/kube-state-metrics \
  --version 8.1.3 \
  --namespace monitoring \
  --create-namespace
```

## Validation

Checked the rollout:

```bash
kubectl -n monitoring rollout status deployment/kube-state-metrics --timeout=120s
```

The deployment rolled out successfully with one ready pod.

Checked the exposed service and endpoint:

```bash
kubectl -n monitoring get pods,svc,deploy -o wide
kubectl -n monitoring get endpoints kube-state-metrics
```

The service endpoint was present on port `8080`.

Checked the metrics endpoint through the Kubernetes API service proxy:

```bash
kubectl get --raw /api/v1/namespaces/monitoring/services/http:kube-state-metrics:8080/proxy/metrics
```

The endpoint returned Prometheus metrics such as `kube_configmap_info`.

## Notes

- The default chart install exposes metrics but does not create a `ServiceMonitor`.
- Prometheus Operator discovery still needs to be wired separately, either by enabling the chart's `prometheus.monitor.enabled` value or by creating a matching `ServiceMonitor`.
