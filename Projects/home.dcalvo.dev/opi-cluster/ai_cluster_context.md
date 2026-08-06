# Homelab Cluster Context
This is a small local k3s cluster used for study and experimentation.

## Cluster
- Runtime: k3s
- Hardware: Orange Pi arm64 nodes
- Network: local trusted home LAN
- Primary use: learning, experiments, and lightweight internal services
- Live state last verified: 2026-08-05

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
  `192.168.1.220-192.168.1.229`.
- The local container registry is exposed through MetalLB.
- Since this runs on a trusted home network, plain HTTP and no authentication are acceptable when they keep experiments simple.

Current load balancer assignments:

| Namespace | Service | LAN address | Service ports |
| --- | --- | --- | --- |
| `kube-system` | `traefik` | `192.168.1.220` | 80, 443 |
| `monitoring` | `grafana` | `192.168.1.221` | 80 |
| `image-resizer` | `image-resizer-api` | `192.168.1.222` | 80 |
| `monitoring` | `prometheus-lb` | `192.168.1.223` | 80 |
| `nginx` | `my-second-lb` | `192.168.1.224` | 80 |
| `container-registry` | `registry` | `192.168.1.225` | 5000 |

### Load Balancer Caveat

- k3s ServiceLB (`svclb-*`/Klipper LB) is still enabled alongside MetalLB.
- ServiceLB pods for Grafana, Image Resizer, Prometheus, and the example nginx
  service are Pending because their requested host port 80 is unavailable.
  MetalLB still assigns their LAN IPs and the application workloads are healthy.
- The ServiceLB pods for Traefik and the registry are Running. If MetalLB is to
  be the only load balancer implementation, disable k3s ServiceLB explicitly
  during a future k3s configuration change.

## Container Registry
- Registry endpoint: `192.168.1.225:5000`
- Kubernetes namespace: `container-registry`
- Transport/auth: plain HTTP, no authentication
- Storage: 20 GiB PVC using the default `local-path` storage class
- Images must support `linux/arm64`
- Both k3s nodes have `/etc/rancher/k3s/registries.yaml` configured for the
  plain-HTTP `192.168.1.225:5000` endpoint; this was verified on 2026-08-03.
- Workstations that push images must list `192.168.1.225:5000` as an insecure
  Docker registry.

## Persistent Storage

- The cluster uses k3s `local-path` storage. These volumes are node-local rather
  than shared, so a pod using one must run on the node that owns its PV.
- `container-registry/registry-storage`: 20 GiB, pinned to `opi1`.
- `monitoring/grafana`: 5 GiB, pinned to `opi2`.
- `monitoring/storage-loki-0`: 10 GiB, pinned to `opi2`.
- `monitoring/storage-tempo-0`: 5 GiB, pinned to `opi2`.

## Monitoring

The former local Prometheus, Loki, Tempo, and Grafana workloads were
decommissioned on 2026-08-05. Their releases, Prometheus resource/operator
controller, monitoring PVCs, and stale resources were removed to reclaim Orange
Pi disk space. kube-state-metrics and node-exporter were then reinstalled as
lightweight OPI metric producers; HP Prometheus stores their samples.

- The local telemetry workloads are Alloy, kube-state-metrics, and
  node-exporter. ServiceMonitors are retained for Alloy discovery and do not
  require a local Prometheus server.
- Namespace: `monitoring`
- Grafana Alloy is installed from pinned Helm chart `grafana/alloy` `1.10.0` as
  one Deployment without host mounts. It tails logs through the Kubernetes API
  and currently keeps only `image-resizer/image-resizer-api` pods. It also
  accepts OTLP gRPC on internal Service port `4317` and OTLP HTTP on `4318`,
  batches traces, and forwards them over the home LAN to the HP Alloy gateway
  at `192.168.1.232:4317`. Applications continue sending traces to the local
  OPI Alloy ClusterIP Service. Its Image Resizer log pipeline sends Loki Push
  API requests to `http://192.168.1.232:3100/loki/api/v1/push`; HP Alloy then
  forwards the logs to HP Loki. New streams use `cluster="opi"`.

## Cross-cluster metrics

- OPI Alloy selects ServiceMonitors labeled `alloy: opi`. The selected sources
  are Image Resizer, kube-state-metrics, and node-exporter.
- OPI Alloy forwards the samples to HP Alloy at
  `http://192.168.1.232:9091/api/v1/metrics/write`, with `cluster="opi"`.
- kube-state-metrics is pinned to chart `7.5.1` (app `v2.19.1`) as one
  Deployment. node-exporter is pinned to chart `4.55.0` (app `v1.11.1`) as a
  two-pod DaemonSet, one pod per OPI node.
- On 2026-08-05, both HP Prometheus replicas independently reported two healthy
  Image Resizer targets, one healthy kube-state-metrics target, and two healthy
  node-exporter targets. `kube_node_info` and `node_uname_info` each described
  both OPI nodes, and OPI Alloy logged no scrape or remote-write errors.
- OPI's k3s-managed `kube-system/kubelet` headless Service exposes both nodes'
  authenticated HTTPS port `10250`. The declarative kubelet ServiceMonitor in
  `monitoring/kubelet/servicemonitor.yaml` is selected by OPI Alloy through its
  `alloy: opi` label and scrapes `/metrics` plus `/metrics/cadvisor`, attaching
  `cluster="opi"` and the endpoint node name. TLS verification is disabled
  only for these trusted-LAN kubelet endpoint addresses. On 2026-08-05, HP
  Prometheus had kubelet metrics for both OPI nodes and cAdvisor container CPU
  and memory series for both nodes.
- OPI Alloy's remote-write WAL is ephemeral pod storage; buffered samples may
  be lost during pod replacement.
- OPI Alloy's own chart-generated ServiceMonitor is labeled `alloy: opi`, so
  HP Prometheus also receives the collector's self-metrics. This allows the
  cross-cluster scrape and remote-write pipeline to be monitored without a
  local Prometheus server.

## Image Resizer API

- Namespace: `image-resizer`
- Images use local-time tags in `vYYYY-MM-DD-HH-MM-SS` format. The checked-in Deployment contains `REPLACE_WITH_TAG`, and the build/deploy script substitutes the generated tag in a temporary manifest before applying it.
- Current image: `192.168.1.225:5000/image-resizer-api:v2026-07-28-10-50-06`
- Deployment: two ARM64 replicas, normally spread across `opi1` and `opi2`
- LoadBalancer address: `192.168.1.222`
- Application port: `8080`; LoadBalancer port: `80`
- Endpoints: `/v1/resize`, `/livez`, `/readyz`, and `/metrics`
- ServiceMonitor: `monitoring/image-resizer-api`, labeled `alloy: opi`; Alloy
  scrapes both replicas at `/metrics` every 15 seconds
- `POST /v1/resize` emits an OpenTelemetry HTTP server span plus child spans for
  upload reading, JPEG decoding, resizing, and encoding. The application exports
  OTLP gRPC to Alloy, and its completion logs include trace and span IDs for
  Tempo/Loki correlation in Grafana.
- Declarative resources and the build/deploy script live in `image-resizer-api/`.
- The in-cluster traffic generator runs as `image-resizer-traffic-gen` in the
  `image-resizer` namespace with one replica. Its ARM64 image is built from
  `image-resizer-api/traffic-gen/` and sends 30 requests per minute to the
  internal `image-resizer-api` Service. Declarative manifests and the build /
  deploy script are in that directory.

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
