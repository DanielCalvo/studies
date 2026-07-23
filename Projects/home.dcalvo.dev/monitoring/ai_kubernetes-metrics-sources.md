# Kubernetes Metrics Sources and Exporters

This is a practical map of metrics sources commonly used with Kubernetes and Prometheus. It is ranked for a small homelab cluster, but the categories are broadly applicable.

## Summary Ranking

| Rank | Source | Commonness | Usefulness | What it answers |
| --- | --- | --- | --- | --- |
| 1 | kube-state-metrics | Very common | Very high | What is Kubernetes trying to run, and is object state healthy? |
| 2 | node-exporter | Very common | Very high | Are the nodes healthy at the OS/hardware level? |
| 3 | kubelet/cAdvisor | Very common | Very high | What are pods/containers actually using? |
| 4 | Prometheus self metrics | Very common | High | Is Prometheus healthy and scraping correctly? |
| 5 | Kubernetes API server metrics | Common | High | Is the API server healthy, slow, or erroring? |
| 6 | CoreDNS metrics | Common | High | Is cluster DNS healthy? |
| 7 | kubelet native metrics | Common | Medium-high | Is kubelet healthy on each node? |
| 8 | controller-manager / scheduler metrics | Common in full control-plane monitoring | Medium | Is the control plane reconciling and scheduling correctly? |
| 9 | etcd metrics | Common in HA/control-plane monitoring | Medium-high | Is Kubernetes storage healthy? |
| 10 | Ingress controller metrics | Common when ingress is important | Medium-high | Are HTTP routes healthy and how much traffic flows through ingress? |
| 11 | CNI/network plugin metrics | Common in larger clusters | Medium | Is pod networking healthy? |
| 12 | CSI/storage plugin metrics | Situational | Medium | Are volumes/provisioning healthy? |
| 13 | kube-proxy metrics | Situational | Medium-low | Are Service proxy rules healthy? |
| 14 | metrics-server | Very common | Low for Prometheus, high for autoscaling | What does `kubectl top` and HPA see? |
| 15 | App-specific exporters | Very common when apps matter | Very high for those apps | Is Postgres/Redis/Nginx/etc. healthy? |
| 16 | Blackbox exporter | Common | Medium-high | Can the cluster reach important endpoints? |
| 17 | GPU/device exporters | Situational | High if hardware exists | Are GPUs/devices healthy and utilized? |

## Already Installed Here

This cluster currently has:

```text
kube-state-metrics
node-exporter
kubelet/cAdvisor metrics
Prometheus self-scrape
Prometheus Operator scrape
Grafana dashboards for kube-state-metrics, node-exporter, and Prometheus
```

Kubelet and cAdvisor are scraped through
`prometheus_operator/kubelet-servicemonitor.yaml`. This fills the gap between
node-level OS metrics and Kubernetes object-state metrics by adding pod and
container usage metrics.

## 1. kube-state-metrics

Usefulness: very high  
Commonness: very common  
Current cluster: installed

kube-state-metrics watches the Kubernetes API and exposes object-state metrics. It tells you about desired/current state, not actual resource usage.

Good for:

```text
deployments with unavailable replicas
pods stuck Pending/Failed
PVC status
node readiness labels/conditions
daemonset rollout state
job completion/failure
resource requests/limits declared on objects
```

Example metrics:

```promql
kube_deployment_status_replicas_available
kube_deployment_status_replicas_unavailable
kube_pod_status_phase
kube_persistentvolumeclaim_status_phase
kube_node_status_condition
kube_daemonset_status_number_unavailable
```

Use this when asking:

```text
Is Kubernetes object state healthy?
Did a rollout fail?
Are workloads scheduled and available?
```

## 2. node-exporter

Usefulness: very high  
Commonness: very common  
Current cluster: installed

node-exporter exposes host OS and kernel metrics from each Linux node.

Good for:

```text
node CPU
node memory
disk space
disk I/O
filesystem health
network interface traffic
load average
kernel and host-level stats
```

Example metrics:

```promql
node_cpu_seconds_total
node_memory_MemAvailable_bytes
node_filesystem_avail_bytes
node_filesystem_size_bytes
node_disk_read_bytes_total
node_disk_written_bytes_total
node_network_receive_bytes_total
node_network_transmit_bytes_total
node_load1
```

Use this when asking:

```text
Are my nodes running out of disk?
Is a node memory constrained?
Is the host CPU busy?
Is the network interface busy?
```

## 3. kubelet/cAdvisor

Usefulness: very high  
Commonness: very common  
Current cluster: scraped directly on both nodes

kubelet embeds cAdvisor-derived container metrics. This is the usual source for pod/container usage in Prometheus.

Good for:

```text
per-pod CPU usage
per-container memory usage
container filesystem usage
container network traffic
container restarts and runtime-level usage signals
```

Common scrape endpoints:

```text
/metrics
/metrics/cadvisor
/metrics/resource
```

Example metrics:

```promql
container_cpu_usage_seconds_total
container_memory_working_set_bytes
container_fs_usage_bytes
container_network_receive_bytes_total
container_network_transmit_bytes_total
kubelet_running_pods
kubelet_running_containers
kubelet_volume_stats_used_bytes
kubelet_volume_stats_capacity_bytes
```

Use this when asking:

```text
Which pod is using CPU?
Which container is using memory?
Which PVC is filling up?
How much network traffic is a pod moving?
```

This is the main missing layer if you want good pod/container resource dashboards.

## 4. Prometheus Self Metrics

Usefulness: high  
Commonness: very common  
Current cluster: installed

Prometheus exposes its own metrics on `/metrics`.

Good for:

```text
scrape health
target counts
query performance
TSDB size and churn
reload health
sample ingestion rate
rule evaluation duration
```

Example metrics:

```promql
prometheus_target_sync_length_seconds
prometheus_tsdb_head_series
prometheus_tsdb_head_samples_appended_total
prometheus_rule_evaluation_duration_seconds
prometheus_config_last_reload_successful
scrape_samples_scraped
up
```

Use this when asking:

```text
Is Prometheus healthy?
Are scrapes failing?
Is cardinality growing?
Did the last config reload succeed?
```

## 5. Kubernetes API Server Metrics

Usefulness: high  
Commonness: common  
Current cluster: not yet scraped directly

The API server exposes control-plane request, latency, auth, watch, and storage-related metrics.

Good for:

```text
API request rate
API latency
error rates
watch behavior
admission latency
API server saturation
```

Example metrics:

```promql
apiserver_request_total
apiserver_request_duration_seconds
apiserver_current_inflight_requests
apiserver_response_sizes
apiserver_admission_controller_admission_duration_seconds
```

Use this when asking:

```text
Is the Kubernetes API slow?
Are clients hammering the API server?
Are requests failing?
```

In k3s, access details can differ from kubeadm-style clusters, so this may need cluster-specific inspection.

## 6. CoreDNS Metrics

Usefulness: high  
Commonness: common  
Current cluster: likely available if CoreDNS exposes metrics

CoreDNS is usually critical even in tiny clusters. DNS issues often look like application issues.

Good for:

```text
DNS request rate
DNS response code rate
DNS latency
cache hits/misses
forwarding issues
```

Example metrics:

```promql
coredns_dns_requests_total
coredns_dns_responses_total
coredns_dns_request_duration_seconds
coredns_cache_hits_total
coredns_cache_misses_total
```

Use this when asking:

```text
Is cluster DNS healthy?
Are DNS requests failing?
Are lookups slow?
```

## 7. kubelet Native Metrics

Usefulness: medium-high  
Commonness: common  
Current cluster: not yet scraped directly

kubelet has metrics beyond cAdvisor container usage.

Good for:

```text
kubelet runtime operations
pod lifecycle operation latency
PLEG health
running pod/container count
volume stats
```

Example metrics:

```promql
kubelet_runtime_operations_total
kubelet_runtime_operations_errors_total
kubelet_runtime_operations_duration_seconds
kubelet_pleg_relist_duration_seconds
kubelet_running_pods
kubelet_running_containers
```

Use this when asking:

```text
Is kubelet having trouble managing pods?
Is the container runtime slow or failing?
```

## 8. Controller Manager and Scheduler Metrics

Usefulness: medium  
Commonness: common in full control-plane monitoring  
Current cluster: not yet scraped

These expose Kubernetes control-loop and scheduling behavior.

Good for:

```text
scheduler latency
pending scheduling attempts
controller queue depth
controller workqueue retries
leader election
reconciliation health
```

Example metric families:

```promql
scheduler_*
workqueue_*
leader_election_*
rest_client_*
```

Use this when asking:

```text
Is Kubernetes scheduling slowly?
Are controllers falling behind?
```

In managed distributions or k3s, these may be embedded or exposed differently than in kubeadm clusters.

## 9. etcd Metrics

Usefulness: medium-high for control plane, lower for single-node experiments  
Commonness: common in production control-plane monitoring  
Current cluster: k3s-specific; may not be relevant if not using external etcd

etcd metrics matter when etcd is the backing store.

Good for:

```text
database size
leader status
disk fsync latency
proposal failures
raft health
backend commit duration
```

Example metrics:

```promql
etcd_server_has_leader
etcd_server_proposals_failed_total
etcd_disk_wal_fsync_duration_seconds
etcd_mvcc_db_total_size_in_bytes
etcd_network_peer_round_trip_time_seconds
```

Use this when asking:

```text
Is Kubernetes storage healthy?
Is etcd disk latency causing cluster issues?
```

## 10. Ingress Controller Metrics

Usefulness: medium-high if ingress is important  
Commonness: common  
Current cluster: Traefik exists in `kube-system`; metrics not reviewed here

Ingress controllers often expose request, latency, response-code, and backend metrics.

For Traefik, useful categories include:

```text
HTTP request totals
response status codes
request duration
entrypoint/router/service metrics
TLS/certificate metrics
```

Use this when asking:

```text
Is ingress routing traffic?
Are HTTP 5xx errors increasing?
Are requests slow?
```

## 11. CNI / Network Plugin Metrics

Usefulness: medium  
Commonness: common in larger clusters  
Current cluster: depends on k3s networking setup

Network plugins may expose metrics for packet drops, policy evaluation, IPAM, tunnel health, or dataplane errors.

Common examples:

```text
Cilium metrics
Calico Felix metrics
Flannel metrics, when exposed
```

Use this when asking:

```text
Are pods failing because of networking?
Are network policies dropping traffic?
Is encapsulation/tunnel traffic unhealthy?
```

For small k3s clusters, this is useful but usually lower priority than kubelet/cAdvisor and CoreDNS.

## 12. CSI / Storage Plugin Metrics

Usefulness: medium  
Commonness: situational  
Current cluster: local-path provisioner is present; metrics not reviewed here

Storage-related metrics vary by provisioner.

Good for:

```text
volume provisioning failures
attach/detach latency
mount errors
storage backend health
capacity signals
```

For PVC usage itself, kubelet volume metrics are often more useful:

```promql
kubelet_volume_stats_used_bytes
kubelet_volume_stats_capacity_bytes
```

## 13. kube-proxy Metrics

Usefulness: medium-low in small clusters  
Commonness: situational  
Current cluster: k3s may use a different networking/proxy mode

kube-proxy metrics can help diagnose Service proxying issues.

Good for:

```text
sync latency
iptables/ipvs rule sync health
network programming errors
```

Example metric families:

```promql
kubeproxy_*
rest_client_*
```

Use this when asking:

```text
Are Kubernetes Services being programmed correctly?
Is kube-proxy falling behind?
```

## 14. metrics-server

Usefulness for Prometheus: low  
Usefulness for Kubernetes autoscaling/debugging: high  
Commonness: very common  
Current cluster: installed

metrics-server is for the Kubernetes resource metrics pipeline. It powers:

```text
kubectl top
Horizontal Pod Autoscaler
Vertical Pod Autoscaler inputs
```

It is not usually scraped by Prometheus as the main source of historical metrics. Prometheus should usually scrape kubelet/cAdvisor directly for historical pod/container resource metrics.

Use this when asking:

```text
What does kubectl top show?
Can HPA get CPU/memory data?
```

## 15. App-Specific Exporters

Usefulness: very high for the app  
Commonness: very common  
Current cluster: add as applications need them

Common exporters:

```text
postgres-exporter
mysql-exporter / mysqld-exporter
redis-exporter
nginx/nginx-vts/exporter variants
blackbox-exporter
snmp-exporter
cert-manager metrics
Argo CD metrics
Longhorn metrics
MetalLB metrics
Traefik metrics
```

Use this when asking:

```text
Is this specific application healthy?
How much work is this service doing?
Are app-level errors increasing?
```

## 16. Blackbox Exporter

Usefulness: medium-high  
Commonness: common  
Current cluster: not installed

Blackbox exporter probes endpoints from Prometheus instead of scraping application internals.

Good for:

```text
HTTP status checks
TLS certificate checks
DNS checks
TCP connectivity checks
ICMP checks, if allowed
```

Example metrics:

```promql
probe_success
probe_duration_seconds
probe_http_status_code
probe_ssl_earliest_cert_expiry
```

Use this when asking:

```text
Can users reach this endpoint?
Is this certificate expiring?
Is DNS working externally?
```

## 17. GPU / Device Exporters

Usefulness: high if hardware exists  
Commonness: situational  
Current cluster: not relevant unless GPU or special devices are added

Examples:

```text
NVIDIA DCGM exporter
Intel GPU exporters
device plugin metrics
```

Good for:

```text
GPU utilization
GPU memory
temperature
power draw
device errors
```

## Recommended Next Steps for This Cluster

1. Add CoreDNS scraping.
2. Consider Traefik metrics if ingress traffic matters.
3. Consider basic alert rules:
   - target down
   - node disk low
   - node memory low
   - pod crash loops
   - Prometheus config reload failed
4. Leave metrics-server alone unless debugging HPA or `kubectl top`.

## Sources

- Kubernetes resource metrics pipeline: https://kubernetes.io/docs/tasks/debug/debug-cluster/resource-metrics-pipeline/
- Kubernetes system component metrics: https://kubernetes.io/docs/concepts/cluster-administration/system-metrics/
- Kubernetes metrics reference: https://kubernetes.io/docs/reference/instrumentation/metrics/
- Kubernetes kube-state-metrics overview: https://kubernetes.io/docs/concepts/cluster-administration/kube-state-metrics/
- kube-state-metrics project: https://github.com/kubernetes/kube-state-metrics
- Prometheus node-exporter guide: https://prometheus.io/docs/guides/node-exporter/
- node-exporter project: https://github.com/prometheus/node_exporter
- cAdvisor Prometheus docs: https://github.com/google/cadvisor/blob/master/docs/storage/prometheus.md
- Kubernetes metrics-server: https://kubernetes-sigs.github.io/metrics-server/
