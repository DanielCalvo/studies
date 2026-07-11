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

## Node SSH Access
- Root SSH access is available to both nodes: `root@192.168.1.201` and `root@192.168.1.202`.
- If a user request can be answered by inspecting the nodes over SSH, it is acceptable to SSH into the relevant node and gather the information directly.
- SSH operations are read-only by default. Do not modify files, services, packages, Kubernetes state, or other node state while logged in unless the user explicitly asks for a change.

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
- Namespace: `monitoring`
- Prometheus Operator is installed from the upstream getting-started `bundle.yaml`.
- Prometheus instance: `monitoring/prometheus`, exposed by MetalLB at `192.168.1.220`.
- Prometheus uses ServiceMonitor label selector `prometheus: homelab`.
- Prometheus scrapes kube-state-metrics, node-exporter, itself, and the Prometheus Operator.
- kube-state-metrics is installed by Helm and scraped through `monitoring/kube-state-metrics`.
- node-exporter is installed by Helm as a DaemonSet, one pod per node.
- Grafana is installed by Helm, exposed by MetalLB at `192.168.1.221`.
- Grafana uses a `local-path` PVC, datasource UID `prometheus`, and pinned Grafana.com dashboards for kube-state-metrics, node-exporter, and Prometheus.
- Useful next monitoring target: kubelet/cAdvisor metrics for pod/container resource usage.

## Practical Constraints
- Prefer lightweight components and simple deployments.
- Do not assume amd64 images will run; this cluster is arm64.
- Prefer direct Kubernetes manifests and clear notes over production-grade platform complexity unless explicitly needed.
- TLS, authentication, and hardening can be relaxed for internal experiments, but call out the tradeoff when it matters.

## Related Repo Areas

- `k3s.md`: install and access notes for the k3s cluster.
- `metallb/`: MetalLB configuration and validation examples.
- `container-registry/`: local Docker Registry v2 manifests and operational notes.
- `monitoring/`: monitoring namespace, Grafana, node-exporter, and Prometheus Operator manifests/notes.
- `monitoring/prometheus_operator/`: Prometheus CR, RBAC, LoadBalancer Service, and ServiceMonitors.
- `monitoring/grafana/`: Grafana Helm values and dashboard provisioning notes.
- `monitoring/node-exporter/`: node-exporter Helm values and notes.
