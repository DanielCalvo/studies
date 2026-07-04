# Homelab Cluster Context
This is a small local k3s cluster used for study and experimentation.

## Cluster
- Runtime: k3s
- Hardware: Orange Pi arm64 nodes
- Network: local trusted home LAN
- Primary use: learning, experiments, and lightweight internal services

## Nodes
- `opi1`: `192.168.1.201`, k3s server
- `opi2`: `192.168.1.202`, k3s agent

## Networking
- MetalLB provides `LoadBalancer` service IPs on the home LAN.
- The local container registry is exposed through MetalLB.
- Since this runs on a trusted home network, plain HTTP and no authentication are acceptable when they keep experiments simple.

## Container Registry
- Registry endpoint: `192.168.1.242:5000`
- Kubernetes namespace: `container-registry`
- Transport/auth: plain HTTP, no authentication
- Storage: PVC using the default `local-path` storage class
- Images must support `linux/arm64`

## Monitoring
- Prometheus Operator is installed for initial experimentation.
- Install method: upstream getting-started `bundle.yaml` applied directly with `kubectl create -f -`.
- Current posture: basic/default install; a more refined setup can come later.
- kube-state-metrics is installed with the `prometheus-community/kube-state-metrics` Helm chart.
- kube-state-metrics release: `kube-state-metrics` in namespace `monitoring`.
- kube-state-metrics service: `kube-state-metrics.monitoring.svc.cluster.local:8080`, `ClusterIP`.
- Current kube-state-metrics posture: metrics endpoint is running, but no `ServiceMonitor` was created by the default chart install.

## Practical Constraints
- Prefer lightweight components and simple deployments.
- Do not assume amd64 images will run; this cluster is arm64.
- Prefer direct Kubernetes manifests and clear notes over production-grade platform complexity unless explicitly needed.
- TLS, authentication, and hardening can be relaxed for internal experiments, but call out the tradeoff when it matters.

## Related Repo Areas

- `k3s.md`: install and access notes for the k3s cluster.
- `metallb/`: MetalLB configuration and validation examples.
- `container-registry/`: local Docker Registry v2 manifests and operational notes.
- `prometheus_operator/`: brief Prometheus Operator install notes for initial monitoring experiments.
- `kube_state_metrics/`: Helm install notes and validation for kube-state-metrics.
