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

Current load balancers in use: 
```
container-registry   registry                   LoadBalancer   10.43.102.211   192.168.1.242   5000:31250/TCP                 21d
image-resizer        image-resizer-api          LoadBalancer   10.43.119.23    192.168.1.222   80:32268/TCP                   21h
kube-system          traefik                    LoadBalancer   10.43.34.169    192.168.1.220   80:32337/TCP,443:30100/TCP     25d
monitoring           grafana                    LoadBalancer   10.43.40.120    192.168.1.221   80:31362/TCP                   15d
monitoring           prometheus-lb              LoadBalancer   10.43.124.164   192.168.1.223   80:30596/TCP                   16d
nginx                my-second-lb               LoadBalancer   10.43.1.121     192.168.1.241   80:31395/TCP                   22d
```

## Container Registry
- Registry endpoint: `192.168.1.242:5000`
- Kubernetes namespace: `container-registry`
- Transport/auth: plain HTTP, no authentication
- Storage: PVC using the default `local-path` storage class
- Images must support `linux/arm64`

## Monitoring
- Namespace: `monitoring`
- Prometheus Operator is installed from the upstream getting-started `bundle.yaml`.
- Prometheus instance: `monitoring/prometheus`.
- Prometheus is exposed through `monitoring/prometheus-lb` at `192.168.1.223`.
- Prometheus uses ServiceMonitor label selector `prometheus: homelab`.
- Prometheus scrapes kube-state-metrics, node-exporter, kubelet/cAdvisor, itself, and the Prometheus Operator.
- kube-state-metrics is installed by Helm and scraped through `monitoring/kube-state-metrics`.
- node-exporter is installed by Helm as a DaemonSet, one pod per node.
- Grafana is installed by Helm, exposed by MetalLB at `192.168.1.221`.
- Grafana uses a `local-path` PVC, datasource UID `prometheus`, and pinned Grafana.com dashboards for kube-state-metrics, node-exporter, Prometheus, and the Kubernetes Views Global, Namespaces, Nodes, and Pods dashboards.
- Kubelet `/metrics` and `/metrics/cadvisor` are scraped on both nodes through `monitoring/prometheus_operator/kubelet-servicemonitor.yaml`.

## Image Resizer API

- Namespace: `image-resizer`
- Images use local-time tags in `vYYYY-MM-DD-HH-MM-SS` format. The checked-in Deployment contains `REPLACE_WITH_TAG`, and the build/deploy script substitutes the generated tag in a temporary manifest before applying it.
- Current image: `192.168.1.242:5000/image-resizer-api:v2026-07-21-16-50-15`
- Deployment: two ARM64 replicas, normally spread across `opi1` and `opi2`
- LoadBalancer address: `192.168.1.222`
- Application port: `8080`; LoadBalancer port: `80`
- Endpoints: `/v1/resize`, `/livez`, `/readyz`, and `/metrics`
- ServiceMonitor: `monitoring/image-resizer-api`, selected by
  `prometheus: homelab`; scrapes both replicas at `/metrics` every 15 seconds
- Declarative resources and the build/deploy script live in `image-resizer-api/`.

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
