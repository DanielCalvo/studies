# Homelab Cluster Context
This is a small local k3s cluster used for study and experimentation.

## Cluster
- Runtime: k3s
- Hardware: Orange Pi arm64 nodes
- Network: local trusted home LAN
- Primary use: learning, experiments, and lightweight internal services
- Live state last verified: 2026-07-24

## Nodes
- `opi1`: `192.168.1.201`, k3s server/control-plane
- `opi2`: `192.168.1.202`, k3s agent
- Both nodes run k3s `v1.36.2+k3s1`, Ubuntu 24.04.4 LTS, kernel
  `6.1.31-sun50iw9`, and containerd `2.3.2-k3s2`.

## Node SSH Access
- Root SSH access is available to both nodes: `root@192.168.1.201` and `root@192.168.1.202`.
- If a user request can be answered by inspecting the nodes over SSH, it is acceptable to SSH into the relevant node and gather the information directly.
- SSH operations are read-only by default. Do not modify files, services, packages, Kubernetes state, or other node state while logged in unless the user explicitly asks for a change.

## Networking
- MetalLB provides `LoadBalancer` service IPs on the home LAN.
- MetalLB is installed from Helm chart `metallb` `0.16.1`.
- MetalLB uses layer 2 advertisement for the automatically assigned pool
  `192.168.1.220-192.168.1.250`.
- The local container registry is exposed through MetalLB.
- Since this runs on a trusted home network, plain HTTP and no authentication are acceptable when they keep experiments simple.

Current load balancer assignments:

| Namespace | Service | LAN address | Service ports |
| --- | --- | --- | --- |
| `kube-system` | `traefik` | `192.168.1.220` | 80, 443 |
| `monitoring` | `grafana` | `192.168.1.221` | 80 |
| `image-resizer` | `image-resizer-api` | `192.168.1.222` | 80 |
| `monitoring` | `prometheus-lb` | `192.168.1.223` | 80 |
| `nginx` | `my-second-lb` | `192.168.1.241` | 80 |
| `container-registry` | `registry` | `192.168.1.242` | 5000 |

### Load Balancer Caveat

- k3s ServiceLB (`svclb-*`/Klipper LB) is still enabled alongside MetalLB.
- ServiceLB pods for Grafana, Image Resizer, Prometheus, and the example nginx
  service are Pending because their requested host port 80 is unavailable.
  MetalLB still assigns their LAN IPs and the application workloads are healthy.
- The ServiceLB pods for Traefik and the registry are Running. If MetalLB is to
  be the only load balancer implementation, disable k3s ServiceLB explicitly
  during a future k3s configuration change.

## Container Registry
- Registry endpoint: `192.168.1.242:5000`
- Kubernetes namespace: `container-registry`
- Transport/auth: plain HTTP, no authentication
- Storage: 20 GiB PVC using the default `local-path` storage class
- Images must support `linux/arm64`

## Persistent Storage

- The cluster uses k3s `local-path` storage. These volumes are node-local rather
  than shared, so a pod using one must run on the node that owns its PV.
- `container-registry/registry-storage`: 20 GiB, pinned to `opi1`.
- `monitoring/grafana`: 5 GiB, pinned to `opi2`.
- `monitoring/storage-loki-0`: 10 GiB, pinned to `opi2`.
- `monitoring/storage-tempo-0`: 5 GiB, pinned to `opi2`.

## Monitoring
- Namespace: `monitoring`
- Prometheus Operator `v0.92.1` is installed in the `default` namespace from the
  upstream getting-started `bundle.yaml`.
- Prometheus instance: `monitoring/prometheus`.
- Prometheus server version: `v3.12.0`.
- Prometheus is exposed through `monitoring/prometheus-lb` at `192.168.1.223`.
- Prometheus uses ServiceMonitor label selector `prometheus: homelab`.
- Prometheus scrapes kube-state-metrics, node-exporter, kubelet/cAdvisor, itself, and the Prometheus Operator.
- kube-state-metrics is installed from Helm chart `7.5.1` (app `2.19.1`) and
  scraped through `monitoring/kube-state-metrics`.
- node-exporter is installed from Helm chart `4.55.0` (app `1.11.1`) as a
  DaemonSet, one pod per node.
- Grafana is installed from Helm chart `10.5.15` (app `12.3.1`) and exposed by
  MetalLB at `192.168.1.221`.
- Grafana uses a `local-path` PVC, datasource UIDs `prometheus`, `loki`, and
  `tempo`, and pinned Grafana.com dashboards for kube-state-metrics,
  node-exporter, Prometheus, and the Kubernetes Views Global, Namespaces, Nodes,
  and Pods dashboards.
- Kubelet `/metrics` and `/metrics/cadvisor` are scraped on both nodes through `monitoring/prometheus_operator/kubelet-servicemonitor.yaml`.
- Loki is installed from pinned Helm chart `grafana/loki` `7.0.0` in monolithic mode. It uses a 10 GiB `local-path` PVC, TSDB indexes, filesystem chunks, seven-day retention, and the internal `monitoring/loki` ClusterIP Service.
- Grafana Alloy is installed from pinned Helm chart `grafana/alloy` `1.10.0` as
  one Deployment without host mounts. It tails logs through the Kubernetes API
  and currently keeps only `image-resizer/image-resizer-api` pods. It also
  accepts OTLP gRPC on internal Service port `4317` and OTLP HTTP on `4318`,
  batches traces, and forwards them to Tempo.
- Tempo is installed from pinned Helm chart `grafana-community/tempo` `2.2.3`
  (Tempo `2.10.7`) as one monolithic replica. It uses a 5 GiB `local-path` PVC,
  local filesystem trace storage, seven-day retention, and the internal
  `monitoring/tempo` ClusterIP Service. Metrics-generator is disabled.
- Prometheus scrapes Loki, Alloy, and Tempo through ServiceMonitors labeled
  `prometheus: homelab`.

## Image Resizer API

- Namespace: `image-resizer`
- Images use local-time tags in `vYYYY-MM-DD-HH-MM-SS` format. The checked-in Deployment contains `REPLACE_WITH_TAG`, and the build/deploy script substitutes the generated tag in a temporary manifest before applying it.
- Current image: `192.168.1.242:5000/image-resizer-api:v2026-07-24-11-09-13`
- Deployment: two ARM64 replicas, normally spread across `opi1` and `opi2`
- LoadBalancer address: `192.168.1.222`
- Application port: `8080`; LoadBalancer port: `80`
- Endpoints: `/v1/resize`, `/livez`, `/readyz`, and `/metrics`
- ServiceMonitor: `monitoring/image-resizer-api`, selected by
  `prometheus: homelab`; scrapes both replicas at `/metrics` every 15 seconds
- `POST /v1/resize` emits an OpenTelemetry HTTP server span plus child spans for
  upload reading, JPEG decoding, resizing, and encoding. The application exports
  OTLP gRPC to Alloy, and its completion logs include trace and span IDs for
  Tempo/Loki correlation in Grafana.
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
