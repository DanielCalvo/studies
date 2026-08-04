For now, installed the operator with the basic/easy steps described here: https://prometheus-operator.dev/docs/getting-started/installation/

```
LATEST=$(curl -s https://api.github.com/repos/prometheus-operator/prometheus-operator/releases/latest | jq -cr .tag_name)
curl -sL https://github.com/prometheus-operator/prometheus-operator/releases/download/${LATEST}/bundle.yaml | kubectl create -f -
```

We can try a more refined/production like install later.

## Prometheus access

Prometheus is exposed on the home LAN by the `monitoring/prometheus-lb`
LoadBalancer Service at `http://192.168.1.223`.

The Service originally requested `192.168.1.220`, but that address is assigned to
the k3s Traefik LoadBalancer. MetalLB therefore left Prometheus pending until its
requested address was changed to the unused `192.168.1.223` on 2026-07-21.

## Kubelet and cAdvisor metrics

`kubelet-servicemonitor.yaml` discovers the Prometheus Operator-managed
`kube-system/kubelet` Service and scrapes both nodes every 30 seconds:

- `/metrics` provides kubelet health and volume metrics.
- `/metrics/cadvisor` provides pod and container CPU, memory, network, and
  filesystem metrics.

Both endpoints use the kubelet's authenticated HTTPS port (`https-metrics`,
10250). Prometheus authenticates with its mounted service-account token. TLS
certificate verification is disabled because kubelet certificates are issued
for the nodes rather than the discovered endpoint addresses; this is acceptable
for this trusted home network.

The endpoint node name is copied into the Prometheus `node` label so Kubernetes
Grafana dashboards can join cAdvisor metrics to kube-state-metrics data.

The kubelet and kube-state-metrics ServiceMonitors preserve labels exposed by
the scrape targets with `honorLabels: true`. This keeps the real Kubernetes
`namespace` and `pod` labels instead of moving them to `exported_namespace` and
`exported_pod`. Both ServiceMonitors also attach `cluster="homelab"` for Grafana
dashboard filtering.

Apply and verify it with:

```bash
kubectl apply -f kubelet-servicemonitor.yaml
kubectl -n monitoring get servicemonitor kubelet
```
