# HP Homelab Cluster Context

This is the x86_64 home-lab Kubernetes cluster used for study, experiments,
and the central observability stack. Live state was last verified on
2026-08-05.

## Cluster

- Distribution: RKE2 `v1.35.6+rke2r1`
- API server: `https://192.168.1.211:6443`
- Architecture: `linux/amd64`
- Network: trusted local home LAN
- CNI: RKE2 Canal (Calico and Flannel)
- Ingress: RKE2 ingress-nginx using host ports 80 and 443
- Merged kubeconfig: `/home/daniel/.kube/config`
- Source kubeconfig: `/home/daniel/.kube/config-hp` (its internal records are
  named `default`; the merged file renames the cluster, user, and context)
- Required context: `hp`

Agents must still pass the kubeconfig and `--context hp` (or Helm's
`--kube-context hp`) explicitly on every cluster command.

## Nodes

- `hp1`: `192.168.1.211`, control-plane and worker
- `hp2`: `192.168.1.212`, worker
- The cluster was reinstalled on 2026-08-04. Both nodes were then Ready on
  Ubuntu 26.04 LTS with kernel `7.0.0-28-generic` and containerd
  `2.2.5-k3s2`.
- Each node is an HP ProDesk with an Intel Core i7-4790S, 16 GiB RAM, and a
  250 GB SSD.
- Non-root SSH access is available as `daniel@192.168.1.211` and
  `daniel@192.168.1.212`. SSH inspection is read-only unless the user
  explicitly authorizes a node change.

## Networking

- MetalLB is installed from Helm chart `0.16.1`.
- Its layer-2 pool is `192.168.1.230-192.168.1.239`.
- RKE2 ingress-nginx currently uses `192.168.1.239`.
- Monitoring reserves `192.168.1.231` for Grafana, `192.168.1.232` for the
  cross-cluster Alloy OTLP gateway, and `192.168.1.233` for Prometheus.
- These services use plain HTTP on the trusted LAN. Do not expose them through
  Internet port forwarding without a separate authentication and TLS design.

## Storage

- Rancher Local Path Provisioner `v0.0.36` is installed in the
  `local-path-storage` namespace.
- Its `local-path` StorageClass uses `WaitForFirstConsumer`, a `Delete` reclaim
  policy, and `/opt/local-path-provisioner` as the node storage root.
- The StorageClass is not the default; workloads must set
  `storageClassName: local-path` explicitly.
- Local Path Provisioner storage is node-local and not replicated. A workload
  using a bound volume cannot fail over to the other node while retaining that
  data.
- Requested PVC capacity is not enforced by Local Path Provisioner. Workloads
  such as Prometheus should configure their own retention size to protect the
  nodes' root disks.

## Cross-cluster tracing

- Grafana Alloy is installed from pinned Helm chart `grafana/alloy` `1.10.0`
  (Alloy `v1.17.0`) as one Deployment in `monitoring`.
- The private `monitoring/alloy` ClusterIP Service exposes Alloy's UI and
  self-metrics. The dedicated `monitoring/alloy-otlp` LoadBalancer exposes
  OTLP gRPC `4317` and Loki Push API `3100` at `192.168.1.232`.
- The LoadBalancer allows `192.168.1.201/32` and `192.168.1.202/32` and uses
  `externalTrafficPolicy: Local` to preserve client source addresses.
- OPI Alloy sends Image Resizer traces and logs to this gateway. HP Alloy
  forwards traces to `tempo.monitoring.svc.cluster.local:4317` and logs to
  `http://loki.monitoring.svc.cluster.local:3100/loki/api/v1/push`.
- End-to-end delivery was verified on 2026-08-05: HP Alloy accepted and sent
  250 spans with zero refused spans, and HP Tempo returned recent five-span
  `POST /v1/resize` traces for `image-resizer-api`.

## Cross-cluster metrics

- OPI Alloy selects ServiceMonitors labeled `alloy: opi` for Image Resizer,
  kube-state-metrics, and node-exporter, and sends their samples to HP Alloy at
  `192.168.1.232:9091/api/v1/metrics/write`.
- HP Alloy fans samples to both Prometheus replicas at their stable pod DNS
  names. Prometheus remote-write receiver mode is enabled on both replicas.
- Verification on 2026-08-05 independently found two healthy Image Resizer
  targets, one healthy kube-state-metrics target, and two healthy node-exporter
  targets in both replicas. `kube_node_info{cluster="opi"}` and
  `node_uname_info{cluster="opi"}` each described both OPI nodes. OPI Alloy
  logged no scrape or remote-write errors.
- Forwarded samples carry `cluster="opi"`. OPI no longer runs a Prometheus
  server; its Alloy instance performs the local scrapes.
- The Alloy remote-write WALs use ephemeral pod storage, so buffered samples
  can be lost during pod replacement. Persistent WAL storage is a follow-up,
  not part of this migration.

## Local Kubernetes metrics

- The two-replica HP Prometheus resource selects the `monitoring/kubelet`
  ServiceMonitor through `prometheus: hp`. It scrapes the RKE2-managed
  `kube-system/kubelet` headless Service on authenticated HTTPS port `10250`
  at both `/metrics` and `/metrics/cadvisor`, attaching `cluster="hp"` and
  each endpoint node name. On 2026-08-05, both kubelet targets were healthy
  and kubelet plus cAdvisor CPU and memory metrics were present for `hp1` and
  `hp2`.

## Alerting

- A single operator-managed Alertmanager (`monitoring/alertmanager`) runs on
  HP with configuration in `monitoring/alertmanager/config-secret.yaml`.
- HP Prometheus selects `PrometheusRule` objects labeled `prometheus: hp` and
  sends alerts to the operator-managed `alertmanager-operated` Service.
- The private `monitoring/alert-sink` ClusterIP service accepts Alertmanager
  webhooks at `/alerts` and logs payloads for homelab testing. Its amd64 image
  is stored at `192.168.1.225:5000/alert-sink`; HP nodes need the registry
  configured as insecure because the registry uses plain HTTP.
- On 2026-08-05, a temporary Image Resizer demo rule delivered both firing and
  resolved notifications to the sink. The checked-in rule is inactive.
- Karma `v0.131` is deployed as `monitoring/karma` and reads the HP
  Alertmanager API. Its private MetalLB LoadBalancer address is
  `192.168.1.234`. On 2026-08-06 it connected successfully and reported zero
  alert groups because the temporary demo rule was inactive.
