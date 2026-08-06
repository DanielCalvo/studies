# Kubernetes Production Monitoring and Alerting

## Question

For a generic production SaaS web service running on Kubernetes, what conditions should be monitored and which of them should generate alerts? The list should cover more than pod failures and include workload controllers, nodes, autoscaling, capacity, storage, networking, the control plane, and the monitoring system itself.

How should alerts such as crash-looping pods, `OOMKilled` containers, failed rollouts, and high resource use be prioritized? Which conditions should page someone immediately, which should generate an urgent Slack notification, and which are better treated as working-hours warnings or capacity-planning signals?

Finally, what additional monitoring and alerting is needed for an on-premises Kubernetes cluster, where machines, disks, control-plane components, storage systems, networks, and physical infrastructure may remain in service for years instead of being replaced frequently by a cloud autoscaler such as Karpenter?

## Answer

The most important design principle is:

> Page on customer-visible symptoms or imminent loss of service. Send lower-severity notifications for Kubernetes causes that need investigation.

A crash-looping pod is important, but one crash-loop among twenty healthy replicas should not necessarily wake someone. If that crash loop leaves the service with zero healthy endpoints, the availability alert should page immediately. Prometheus similarly recommends symptom-based paging and cause-based diagnostic alerts.

### Suggested severity model

| Severity | Response | Typical delivery |
|---|---|---|
| **P1 — Critical** | Acknowledge in about five minutes; customer impact, data-loss risk, or complete loss of redundancy | Pager |
| **P2 — High** | Investigate in 15–30 minutes; likely imminent impact or significant degradation | Pager for tier-one services; urgent Slack otherwise |
| **P3 — Warning** | Investigate during working hours, normally within one to three days | Slack or ticket |
| **P4 — Informational** | Trend, capacity-planning, or forensic signal | Dashboard or report |

These are starting points. Every workload should have an importance label such as `tier=critical|standard|batch`, because the same condition has different consequences for a login service and a development report generator.

### Customer-facing and application alerts

These are not strictly Kubernetes alerts, but they should sit above all Kubernetes alerts:

- **P1:** External synthetic journey fails, such as login, checkout, or creation of a customer resource. Require persistence for two to five minutes or failures from multiple locations.
- **P1/P2:** HTTP availability or error-budget burn.
- **P1/P2:** Application latency violates its SLO, preferably measured at the ingress or application edge.
- **P1:** No successful requests or business transactions when traffic is expected.
- **P1/P2:** Critical background processing is late enough to threaten a customer promise.
- **P1/P2:** Queue backlog or oldest-message age approaches a business deadline.
- **P1/P2:** Critical dependencies such as a database, identity provider, payment system, object store, or third-party API are unavailable.
- **P2/P3:** TLS certificate approaches expiry; for example, P2 at seven days and P3 at 30 days, adjusted for the expected renewal automation.

The external synthetic monitor should live outside the cluster or monitoring failure domain. It can detect a complete cluster or Prometheus outage that internal monitoring cannot see.

### Pods and containers

#### Availability and lifecycle

- **P1:** A critical service has no ready endpoints, normally after one or two minutes.
- **P1/P2:** Available replicas fall below the service's minimum safe capacity, such as fewer than two ready replicas across failure domains.
- **P2:** A pod remains in a crash loop for approximately five to ten minutes. Escalate to P1 if this removes service capacity or affects a critical singleton.
- **P2:** A container is repeatedly `OOMKilled`. Record every production OOM kill; page when it repeats, affects multiple replicas, or reduces ready capacity.
- **P2:** A pod remains `Pending` for five to fifteen minutes because of insufficient resources, affinity or topology rules, an untolerated taint, an unbound PVC, exhausted pod or IP capacity, a missing extended resource, or an exceeded quota.
- **P2:** A running pod remains non-ready. Running is not equivalent to serving.
- **P2:** A container is stuck in `ImagePullBackOff`, `ErrImagePull`, `CreateContainerConfigError`, `CreateContainerError`, a volume-mount failure, or another waiting condition.
- **P2/P3:** The restart rate increases abnormally. Detect increases rather than testing `restart_count > 0`, because the counter persists for the pod's lifetime.
- **P2:** Liveness or startup probes fail repeatedly.
- **P3:** Readiness-probe failures increase without current availability loss.
- **P2:** A pod remains terminating substantially beyond its termination grace period, commonly ten to fifteen minutes.
- **P2/P3:** A pod is unexpectedly evicted or preempted. Page only if service capacity or an important job is affected.
- **P2:** A critical pod disappears and its controller does not replace it.
- **P3:** A service container exits unexpectedly or completes when it should remain running.

Common crash-loop causes include application misconfiguration, invalid probes, missing Secrets or ConfigMaps, unavailable dependencies, image problems, resource constraints, permissions, security contexts, and volume failures.

#### Resource behavior

- **P2:** A container is repeatedly OOM-killed.
- **P3:** Memory working set remains above approximately 85–90% of its limit.
- **P3:** Memory-limit exhaustion is forecast within hours or days.
- **P3:** Sustained CPU throttling is accompanied by latency or backlog.
- **P4:** CPU utilization is high without a customer or application symptom.
- **P3:** CPU, memory, or I/O pressure-stall time remains high.
- **P3:** A container exceeds expected ephemeral-storage usage.
- **P3:** A pod lacks resource requests or limits where policy requires them.
- **P3:** Requests consistently and materially exceed actual usage.
- **P3:** Limits are consistently too close to normal usage.

A near-memory-limit alert is normally a warning, while an actual OOM kill is a high-severity event. Memory utilization is not a perfect predictor: applications may retain reclaimable cache or intentionally operate near their limit.

### Workload controllers and releases

#### Deployments

- **P1/P2:** Desired replicas are not available.
- **P2:** A rollout exceeds its progress deadline or reports `ProgressDeadlineExceeded`.
- **P2:** Updated replicas do not converge on the desired count.
- **P2:** The observed generation does not match the current generation.
- **P2:** A rollout is stuck because new pods cannot become ready.
- **P2:** A rollout leaves too few old and new replicas available.
- **P3:** A production Deployment unexpectedly has zero desired replicas.
- **P3:** A Deployment remains paused unexpectedly.

Kubernetes reports a stalled Deployment but continues retrying; it does not automatically repair an invalid release. Quota, readiness-probe, image-pull, permission, and configuration failures are common causes.

#### StatefulSets

- **P1/P2:** Ready replicas fall below the workload's safe quorum or availability requirement.
- **P2:** An update does not roll out or an unexpected partition blocks it.
- **P2:** Current and desired revisions do not converge.
- **P2:** An ordinal pod is missing or stuck.
- **P2:** A stateful pod cannot attach or mount its storage.
- **P3:** The StatefulSet generation does not converge.

Quorum-aware alerting is essential: two of three database replicas is not equivalent to nineteen of twenty stateless web replicas.

#### DaemonSets

- **P2:** A required system DaemonSet is missing from an eligible node.
- **P2:** A DaemonSet rollout is stuck.
- **P2:** DaemonSet pods are scheduled where they should not be.
- **P1/P2:** A CNI, CSI, kube-proxy, logging, security, or node-monitoring DaemonSet is unavailable across several nodes.

#### Jobs and CronJobs

- **P2/P3:** A Job fails.
- **P2/P3:** A Job exceeds its expected duration.
- **P2:** A critical CronJob has not succeeded within its business deadline.
- **P3:** A CronJob misses multiple schedules.
- **P3:** Jobs remain active or stuck unexpectedly.
- **P3:** Completed Jobs are not being cleaned up.

Do not page merely because one retryable batch execution failed if the next execution can complete before the business deadline. Alert on time since the last successful completion.

#### Disruption protection

- **P2/P3:** A PodDisruptionBudget has too few healthy pods.
- **P3:** A PDB permits no disruptions for an extended period.
- **P2 during maintenance:** A drain or upgrade cannot progress because of a PDB.
- **P3:** Supposedly highly available replicas all occupy one node, rack, hypervisor, or availability zone.

### Autoscaling and capacity

#### Horizontal and vertical scaling

- **P2:** An HPA remains at maximum replicas while its scaling metric remains above target.
- **P2:** An HPA cannot obtain its CPU, custom, or external metric.
- **P2:** Current replicas cannot reach desired replicas.
- **P3:** An HPA repeatedly oscillates between sizes.
- **P3:** An HPA remains at minimum while the application is overloaded, suggesting a missing or incorrect metric.
- **P3:** VPA recommendations materially exceed configured requests or limits.
- **P3:** VPA repeatedly evicts pods, if it is permitted to do so.

#### Cluster capacity and provisioning

- **P1/P2:** Critical pods are unschedulable and no additional capacity is arriving.
- **P2:** Cluster Autoscaler or Karpenter cannot provision nodes.
- **P2:** Node provisioning or registration exceeds its expected duration.
- **P2:** A node pool reaches its CPU, memory, node-count, or provider limit.
- **P2:** A cloud quota prevents instance, disk, IP, or load-balancer creation.
- **P2/P3:** The cluster cannot tolerate the loss of one node or one failure domain.
- **P3:** CPU, memory, pod, IP, storage, or GPU capacity is forecast to run out.
- **P3:** Request overcommit exceeds organizational policy.
- **P3:** Namespace ResourceQuota exceeds approximately 80–90%.
- **P2:** ResourceQuota is exhausted in an active production namespace.
- **P3:** Topology imbalance undermines expected redundancy.

Capacity alerts should use requests and schedulability as well as actual usage. A cluster can have low CPU utilization and still be unable to schedule a pod because requests, affinity, or topology constraints consume its usable capacity.

### Nodes

- **P1/P2:** A node is `NotReady` or unreachable. Severity depends on whether workloads remain redundant.
- **P2:** Node readiness flaps.
- **P2:** The node reports `MemoryPressure`, `DiskPressure`, `PIDPressure`, or `NetworkUnavailable`.
- **P2:** Kubelet or the container runtime is down or unhealthy.
- **P2:** Pods cannot start or pod sandboxes cannot be created on a node.
- **P2:** The kernel invokes its OOM killer unexpectedly.
- **P2:** A filesystem becomes read-only or produces I/O errors.
- **P2:** A node experiences significant packet loss or interface errors.
- **P3:** Node-level CPU, memory, I/O, or PID pressure remains high.
- **P3:** A node or image filesystem is forecast to fill.
- **P3:** Filesystem inodes, conntrack entries, or file descriptors approach exhaustion.
- **P3:** A node reboots unexpectedly or experiences a kernel panic.
- **P3:** Node clock drift becomes excessive.
- **P3:** A node remains cordoned or unschedulable unexpectedly.
- **P3:** Automation does not remove or repair unhealthy nodes.
- **P3:** Kubernetes or operating-system versions drift outside policy.

Rapid Karpenter replacement reduces the value of long-horizon disk forecasting, but it does not eliminate disk alerts. Image pulls, container logs, or inode exhaustion can push a fresh node into `DiskPressure` quickly, after which kubelet may evict pods.

For a disposable cloud node fleet:

- Treat `DiskPressure` or resulting evictions as **P2**.
- Treat rapid growth likely to affect pods before replacement as **P3**.
- Keep a disposable node merely exceeding 70% disk use on a dashboard.
- Treat repeated short-lived nodes filling their disks as a service-level **P2/P3**, because it indicates a systematic logging, image, or ephemeral-storage problem.

### Kubernetes control plane

With EKS or another managed control plane, the provider owns most component processes, but the platform team still monitors the API as a consumer. A self-managed cluster must alert on each component directly.

#### API server

- **P1:** The Kubernetes API is unavailable from multiple monitoring locations.
- **P1/P2:** API request error-budget burn.
- **P2:** Sustained high API latency or elevated `5xx` responses.
- **P2:** Admission webhooks fail or add excessive latency.
- **P2:** An aggregated API is unavailable.
- **P2:** Writes are rejected or requests terminate abnormally.
- **P3:** Clients are heavily throttled or retrying.
- **P3:** Authorization denials increase unexpectedly, particularly after a release.
- **P3:** API object count or request volume grows unusually quickly.

#### Scheduler

- **P1/P2:** The scheduler is unavailable in a self-managed cluster.
- **P2:** The scheduling queue grows or pods wait too long to be scheduled.
- **P2:** Scheduling fails across unrelated workloads.
- **P3:** Scheduler leadership changes frequently or its workqueue accumulates.

#### Controller managers

- **P1/P2:** The controller manager or cloud controller manager is unavailable.
- **P2:** A leader is absent or changes frequently.
- **P2:** A controller workqueue is stuck or reconciliation errors rise.
- **P2:** Cloud-provider operations fail repeatedly.
- **P2:** Node, endpoint, service, volume, or Job controllers stop converging desired and actual state.

#### Certificates and versions

- **P2:** A control-plane or kubelet certificate expires within seven days.
- **P3:** A certificate expires within 30 days.
- **P2:** Automatic certificate renewal fails.
- **P2/P3:** A CSR remains pending unexpectedly.
- **P3:** Kubernetes components have unsupported or unsafe version skew.
- **P3:** A Kubernetes release approaches the end of support.

### Networking, DNS, ingress, and service discovery

- **P1:** A production ingress or Gateway is unavailable.
- **P1/P2:** Ingress error rate or latency violates the service SLO.
- **P2:** The ingress controller has insufficient ready replicas.
- **P2:** LoadBalancer provisioning or reconciliation fails.
- **P2:** An expected Service has no ready EndpointSlices.
- **P2:** CoreDNS is unavailable or has insufficient replicas.
- **P2:** DNS errors or timeouts increase.
- **P2:** A DNS-resolution synthetic test fails from multiple nodes.
- **P2:** The CNI agent is unavailable on nodes.
- **P2:** Pod sandbox creation fails.
- **P2:** The pod or service IP pool approaches exhaustion.
- **P2:** Cross-node or cross-zone connectivity tests fail.
- **P2:** kube-proxy, an eBPF service agent, or its equivalent is unavailable.
- **P2:** The network-policy engine fails.
- **P2/P3:** Packet drops, retransmissions, interface errors, or MTU-related failures rise substantially.
- **P3:** Conntrack entries or NAT ports approach exhaustion.
- **P3:** ExternalDNS cannot update records.
- **P2:** A service-mesh control plane is unavailable, if used.
- **P2/P3:** Mesh mTLS certificate renewal fails or certificates approach expiry.
- **P3:** EndpointSlice counts or Service-to-endpoint mappings are abnormal.

CoreDNS process health is insufficient by itself; test a real lookup from inside the cluster.

### Persistent storage

- **P1:** A critical volume reports corruption, data loss, or complete unavailability.
- **P2:** A PVC enters `Lost` or another abnormal state.
- **P2:** A PVC remains `Pending`.
- **P2:** Volume attachment, detachment, mounting, or unmounting fails repeatedly.
- **P2:** A CSI controller or node plugin is unavailable.
- **P2:** Provisioning, attachment, or mount latency becomes excessive.
- **P2:** CSI volume health reports an abnormal condition.
- **P2/P3:** A persistent volume is forecast to fill within the relevant response window.
- **P3:** A persistent volume exceeds approximately 80–90% use.
- **P3:** Volume inodes are forecast to run out.
- **P3:** A storage class cannot provision in a failure domain.
- **P3:** Orphaned volumes or snapshots accumulate.
- **P2:** A required snapshot or backup fails.
- **P2/P3:** Stateful production data has no recent successful backup.
- **P2:** A periodic restore test fails.

Time-to-full is more useful than a universal percentage. A volume at 85% but stable for a year may be safe, while one at 60% and growing by 10% per hour is urgent.

### Monitoring, logging, and alerting infrastructure

The monitoring system itself needs monitoring:

- **P1:** An external watchdog or dead-man's switch stops receiving alerts.
- **P1/P2:** All Prometheus replicas or the hosted metrics service become unavailable.
- **P2:** Alertmanager cannot deliver pager or Slack notifications.
- **P2:** The Alertmanager cluster loses redundancy.
- **P2:** Important scrape targets disappear.
- **P2:** Recording or alerting rule evaluation fails.
- **P2:** Remote-write failures, backlog, or dropped samples threaten alert coverage.
- **P2:** Monitoring storage fills or ingestion stops.
- **P2:** kube-state-metrics is unavailable or missing shards.
- **P2:** Node-exporter or kubelet metrics disappear across several nodes.
- **P2:** The logging pipeline drops production logs.
- **P2/P3:** Logging backlog or storage approaches exhaustion.
- **P3:** Metric cardinality or log volume grows unexpectedly.
- **P3:** Monitoring retention falls below requirements.
- **P3:** An alert remains firing implausibly long, suggesting broken ownership or automation.
- **P3:** Dashboards or datasources fail health checks.

Metrics Server supports the lightweight Resource Metrics API and `kubectl top`; it is not a complete historical production-monitoring system. A full pipeline should collect component metrics, kubelet and cAdvisor metrics, Kubernetes object state, application metrics, and external probes.

### Security-relevant operational alerts

These should be coordinated with the security team rather than sent indiscriminately to the normal Kubernetes pager:

- **P1/P2:** Unexpected creation of privileged pods or host-mounted workloads.
- **P1/P2:** Suspicious use of `exec`, `attach`, port-forwarding, or ephemeral containers.
- **P2:** Unexpected cluster-admin binding or RBAC escalation.
- **P2:** Audit logging stops.
- **P2:** Image-signature or admission-policy enforcement stops working.
- **P2:** Unexpected anonymous or unauthenticated API activity.
- **P2:** Authentication or authorization failures spike suddenly.
- **P2:** Runtime security detects escape behavior, host namespace access, or sensitive-file modification.
- **P3:** Workloads use stale or vulnerable images according to organizational policy.
- **P3:** Secret, certificate, or service-account rotation fails.

## Additional monitoring for on-premises Kubernetes

On-premises Kubernetes introduces long-lived machines, physical infrastructure, self-managed control planes, storage systems, and capacity that may take weeks or months to replace.

### Self-managed etcd

- **P1:** etcd loses its leader or quorum.
- **P1/P2:** An etcd member is unavailable.
- **P2:** Leadership changes frequently.
- **P2:** Consensus proposals fail or remain pending.
- **P2:** A follower remains significantly behind the leader.
- **P2:** Fsync or commit latency is excessive.
- **P2:** The etcd database approaches its backend quota.
- **P2:** A `NOSPACE` alarm is active.
- **P3:** The database requires compaction or defragmentation.
- **P2:** A scheduled etcd snapshot fails.
- **P1/P2:** An etcd restore test fails.
- **P3:** Peer or client certificates approach expiry.

An etcd cluster without a leader cannot make progress. Rapid leadership changes commonly indicate network instability or excessive load, while backend-quota exhaustion leads to a `NOSPACE` alarm and rejected writes.

### Long-lived node operating systems

- **P2:** A root, runtime, image, log, or kubelet filesystem approaches exhaustion.
- **P2:** Filesystem inodes approach exhaustion.
- **P2:** A filesystem becomes read-only or reports block-device I/O errors.
- **P2:** RAID degrades or a rebuild fails.
- **P2:** The kernel reports OOM events, hung tasks, soft lockups, or a panic.
- **P2:** Kubelet or the container runtime becomes unhealthy.
- **P3:** Journal or container-log growth becomes excessive.
- **P3:** Stale images and stopped containers consume substantial disk space.
- **P3:** Swap is unexpectedly enabled or used, according to cluster policy.
- **P3:** The OS patch level or kernel falls outside policy.
- **P3:** A security or kernel update requires a reboot.
- **P3:** Node configuration drifts from the managed baseline.

### Physical hardware

- **P1/P2:** SMART or NVMe health predicts disk failure.
- **P2:** ECC errors indicate a failing DIMM.
- **P2:** The CPU reports machine-check errors.
- **P2:** A fan, PSU, voltage, or temperature alarm fires.
- **P2:** One PSU or NIC in a redundant pair fails.
- **P2:** The BMC or IPMI reports a hardware-health alarm.
- **P3:** Corrected hardware errors increase over time.
- **P3:** Firmware, BIOS, RAID-controller, or BMC versions fall outside policy.

Corrected ECC errors, SMART counters, and intermittent NIC errors are especially useful on machines that may remain in service for years.

### Physical and datacenter networking

- **P1/P2:** A top-of-rack switch, router, firewall, or load balancer becomes unavailable.
- **P2:** Switch-port errors, drops, flaps, or speed and duplex mismatches occur.
- **P2:** A bonded interface loses redundancy.
- **P2:** A BGP or BFD peer is down or unstable, if used.
- **P2:** VLAN, routing, overlay, or MTU connectivity fails.
- **P2:** DHCP or PXE fails where nodes depend on it.
- **P2:** Internal DNS or NTP becomes unavailable.
- **P2:** Node or network-device clock drift becomes severe.
- **P3:** Link utilization is forecast to saturate.
- **P3:** A network-device configuration backup fails.

### On-premises storage systems

- **P1:** A storage cluster loses quorum or risks data loss.
- **P2:** A storage array, SAN, NAS, or distributed-storage node degrades.
- **P2:** A storage pool approaches capacity.
- **P2:** Replication or erasure-coding health degrades.
- **P2:** Path or multipath redundancy is lost.
- **P2:** Storage latency or IOPS saturation threatens workloads.
- **P2:** Disks fail, rebuilds fail, or rebuilds take dangerously long.
- **P2:** Snapshot, replication, or off-site backup fails.
- **P3:** A thin-provisioning data or metadata pool is forecast to fill.
- **P3:** Storage-controller cache or battery health degrades.

Monitor both the Kubernetes PVC view and the underlying storage system. Kubernetes may expose only mount timeouts while the storage system knows that a controller, disk group, or path has failed.

### Power and environment

- **P1/P2:** A UPS switches to battery or utility power is lost.
- **P2:** UPS runtime falls below the shutdown requirement.
- **P2:** A generator, PDU, or redundant power feed becomes unavailable.
- **P2:** Rack temperature or humidity leaves the safe range.
- **P2:** Cooling or leak detection reports a problem.
- **P3:** A UPS battery requires replacement or has insufficient capacity.

### Capacity and lifecycle

- **P2/P3:** Remaining schedulable capacity falls below the failure-domain reserve.
- **P3:** Capacity will run out within the hardware procurement lead time.
- **P3:** Spare-disk, spare-node, or replacement-part inventory falls below policy.
- **P3:** Warranty or support contracts approach expiry.
- **P3:** Kubernetes, OS, firmware, CNI, CSI, or runtime versions approach end of support.
- **P3:** Certificates or required licenses approach expiry.

Cloud capacity may require only a two-week forecast. If purchasing and installing on-premises hardware takes three months, warnings need to fire months earlier.

### Virtualization layer

If Kubernetes runs on virtual machines, also monitor:

- **P1/P2:** Hypervisor or virtualization management plane availability.
- **P2:** Datastore capacity and latency.
- **P2:** Hypervisor contention, ballooning, swapping, or CPU ready time.
- **P2:** Failure of anti-affinity, causing control-plane or application replicas to share one physical host.
- **P2:** Snapshot accumulation that affects capacity or performance.
- **P3:** Insufficient N+1 hypervisor capacity.

## Recommended default pager set

The comprehensive catalog should not be enabled wholesale as paging alerts. For a normal SaaS platform, a deliberately small initial pager set would be:

1. External synthetic failure or fast SLO error-budget burn.
2. A critical service has zero endpoints or falls below its minimum safe replicas.
3. The Kubernetes API is unavailable or severely degraded.
4. Critical DNS, ingress, CNI, or service networking fails.
5. Several nodes are unavailable or the cluster loses N+1 capacity.
6. A critical workload is unschedulable and autoscaling cannot provide capacity.
7. A stateful service loses quorum or reports data-loss risk.
8. A critical persistent volume becomes unavailable or abnormal.
9. The monitoring and alert-delivery watchdog fails.
10. On-premises: etcd quorum loss, critical storage failure, major network failure, or power and environmental emergencies.

Crash loops, OOM kills, stalled rollouts, failed Jobs, HPA saturation, disk forecasts, certificate expiry, and node pressure should all produce signals, but many belong in P2 or P3 notification channels unless they coincide with customer impact.

## Alert-noise controls

- Alert on a service or controller rather than separately on every pod.
- Aggregate related pod failures into one alert.
- Inhibit pod alerts when their owning node is down.
- Inhibit lower-level Kubernetes alerts when a higher-level service outage is already paging.
- Add a `for` duration to remove brief rollout and rescheduling noise.
- Use shorter durations for critical singletons and longer durations for large replica sets.
- Suppress expected failures during declared maintenance.
- Attach the owner, environment, cluster, namespace, workload, dashboard, logs, recent deployment, and runbook to every notification.
- Ensure every paging alert has a concrete human action.
- Record events such as OOM kills even when they do not page.
- Test alert delivery periodically, including an end-to-end synthetic alert.

## References

- [Prometheus alerting practices](https://prometheus.io/docs/practices/alerting/)
- [The Zen of Prometheus](https://prometheus.io/docs/practices/the_zen/)
- [kube-prometheus runbook catalog](https://runbooks.prometheus-operator.dev/)
- [Kubernetes observability overview](https://kubernetes.io/docs/concepts/cluster-administration/observability/)
- [Kubernetes system component metrics](https://kubernetes.io/docs/concepts/cluster-administration/system-metrics/)
- [Kubernetes metrics reference](https://kubernetes.io/docs/reference/instrumentation/metrics/)
- [Kubernetes resource monitoring](https://kubernetes.io/docs/tasks/debug/debug-cluster/resource-usage-monitoring/)
- [Node-pressure eviction](https://kubernetes.io/docs/concepts/scheduling-eviction/node-pressure-eviction/)
- [Deployment rollout and progress deadlines](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
- [KubePodCrashLooping runbook](https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubepodcrashlooping/)
- [Persistent-volume capacity runbook](https://runbooks.prometheus-operator.dev/runbooks/kubernetes/kubepersistentvolumefillingup/)
- [CSI volume health monitoring](https://kubernetes.io/docs/concepts/storage/volume-health-monitoring/)
- [Kubernetes DNS troubleshooting](https://kubernetes.io/docs/tasks/administer-cluster/dns-debugging-resolution/)
- [etcd metrics](https://etcd.io/docs/v3.8/metrics/)
- [etcd maintenance](https://etcd.io/docs/v3.5/op-guide/maintenance/)
