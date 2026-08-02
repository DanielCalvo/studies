# Grafana Alloy deployment and configuration options

This is a reconnaissance guide to the flexibility available in Grafana Alloy.
The homelab's desired collector configuration remains in `values.yaml`.

The Helm value names are based on the versions pinned in this repository:

- Helm chart: `grafana/alloy` `1.10.0`
- Alloy: `v1.17.0`

Check the values for the exact version before copying examples to another
release:

```bash
helm show values grafana/alloy --version 1.10.0
```

## Mental model

Alloy is a programmable telemetry collector. Its configuration creates a graph
of components:

```text
discover or receive -> relabel/process/sample -> batch/queue -> write/export
```

Each component exports values or receivers that other components consume. Alloy
evaluates those dependencies automatically, so configuration blocks do not have
to be written in execution order.

One Alloy process can handle:

- Prometheus discovery, scraping, relabeling, and remote write
- Loki log collection, parsing, relabeling, and writing
- OpenTelemetry metrics, logs, and traces
- Pyroscope profiling data
- Host and application exporters
- eBPF-based instrumentation where permissions allow it
- Routing one signal to multiple destinations

Alloy is a collector, not the long-term backend. Loki, Prometheus/Mimir,
Tempo/OTLP destinations, and Pyroscope provide storage and querying.

## Helm chart value families

### Alloy configuration delivery

| Value | What it controls |
| --- | --- |
| `alloy.configMap.create` | Whether the chart creates the configuration ConfigMap |
| `alloy.configMap.content` | Inline Alloy configuration language |
| `alloy.configMap.name` and `key` | Use an externally managed ConfigMap instead |
| `configReloader.enabled` | Adds a sidecar that reloads Alloy when its ConfigMap changes |
| `alloy.extraEnv`, `alloy.envFrom` | Inject environment values or Kubernetes Secrets |
| `alloy.extraArgs` | Pass flags to `alloy run` |
| `alloy.stabilityLevel` | Permit generally available, public-preview, or experimental components |

The homelab keeps the pipeline inside `alloy.configMap.content`, so the Helm
release owns the ServiceAccount, RBAC, workload, and exact running pipeline.

For larger configurations, an existing ConfigMap can be generated from separate
`.alloy` files using another declarative tool. That avoids embedding hundreds of
lines inside Helm values while keeping the result reproducible.

### Controller type and replicas

The pinned chart supports:

```yaml
controller:
  type: daemonset | deployment | statefulset
  replicas: 1
```

It also exposes:

- Horizontal and vertical autoscaling
- Update strategy and rollout history
- Pod disruption budgets
- Host networking and host PID access
- Affinity, topology spreading, node selectors, tolerations, and priority
- Extra volumes, init containers, sidecars, and StatefulSet claim templates

The correct controller depends on what Alloy must reach and what state it must
preserve.

### Process and UI settings

Useful `alloy` values include:

- `listenAddr`, `listenPort`, and `enableHttpServerPort` for the health/debug UI
- `storagePath` for component state
- `resources` for CPU and memory requests/limits
- `securityContext` for user, capabilities, and filesystem restrictions
- `mounts.varlog` and `mounts.dockercontainers` for node-local log files
- `extraPorts` for OTLP, syslog, Faro, or other receivers
- `clustering.enabled` and cluster name/settings
- anonymous usage reporting control

### Kubernetes access and security

The chart can manage:

- ServiceAccount creation and token mounting
- ClusterRole/Role rules for Kubernetes discovery and logs
- Separate cluster-scoped rules for Nodes and node metrics
- NetworkPolicies
- Pod and container security contexts
- Secrets and environment injection
- Image registry, tag, digest, pull policy, and pull secrets

RBAC should match the configured components. A collector that only receives OTLP
does not need broad Kubernetes discovery permissions. A collector using
Kubernetes discovery, Pod logs, Events, or node metrics does.

Some integrations need elevated privileges. For example, host file access,
journald, cAdvisor, eBPF profiling, and Beyla may require host mounts, root, or
Linux capabilities. Do not apply a universal restrictive security context
without checking the selected components.

### Networking

The chart can create:

- A ClusterIP or externally exposed Service
- An Ingress for the HTTP UI or receiver ports
- Extra named ports
- NetworkPolicies
- A ServiceMonitor for Alloy's Prometheus metrics

A logs-only shipper normally needs inbound access only for its UI/health and
outbound access to Kubernetes and Loki. An OTLP gateway needs stable inbound
gRPC/HTTP ports from applications.

### CRDs and additional resources

The chart can install Grafana monitoring CRDs and accepts `extraObjects`.
CRD-backed sources such as `loki.source.podlogs` let teams select log workloads
with `PodLogs` resources instead of editing the central Alloy pipeline.

The homelab does not currently need those CRDs, so their installation is
disabled.

## Common deployment topologies

### Central Deployment

A small number of Alloy replicas centrally discover or receive application
telemetry.

Good for:

- Kubernetes API-based Pod log collection on a small cluster
- Prometheus scraping across reachable services
- OTLP gateways for applications
- Central processing and export

Benefits:

- Fewer collectors to operate
- Independent scaling
- Straightforward meta-monitoring

Tradeoffs:

- More network traffic to the central collectors
- Cannot read host-local files unless they are remotely available
- A single replica needs availability consideration

The homelab currently uses one Deployment because
`loki.source.kubernetes` tails the selected containers through the Kubernetes
API and does not need host mounts.

### DaemonSet / host daemon

One Alloy runs on every node or machine.

Good for:

- `/var/log/pods` and container-runtime files
- systemd journal logs
- host metrics and hardware exporters
- node-local cAdvisor or kubelet access
- local buffering before forwarding

Benefits:

- Direct access to local telemetry
- Lower traffic to the Kubernetes logs API
- Failure or pressure is isolated per node

Tradeoffs:

- Replica count follows node count
- Every collector opens backend connections
- Host mounts and privileges require care
- Configuration and state exist on every node

This is the likely topology when the homelab adds host journal logs.

### StatefulSet

Stable pod identity and persistent claims make this useful for central
collectors that maintain important state.

Good for:

- Prometheus scraping with a remote-write WAL
- Stable clustering membership
- Predictable claim-to-replica attachment
- Pipelines where replay after restart matters

A StatefulSet does not automatically make a pipeline highly available. Sources
still need to be sharded or load-balanced correctly, and the destination must
handle duplicate delivery where applicable.

### Sidecar

Alloy runs beside an application in the same Pod or task.

Good for:

- Short-lived or specialized applications
- Telemetry that is only reachable through a shared localhost/network context
- Per-workload credentials or isolation

Tradeoffs:

- Repeated configuration and resource overhead
- More collectors to upgrade
- Usually unnecessary when a node or central collector can reach the workload

### Multi-tier

Node-local or sidecar Alloys forward to central Alloy gateways, which then
sample, enrich, batch, and export.

This is useful for:

- Large networks with controlled egress
- Centralized credentials
- Tail sampling or consistent processing
- Reducing the number of direct backend connections

It is more operationally complex and unnecessary for this two-node cluster.

## Clustering and scaling

Alloy clustering lets participating components distribute targets among Alloy
replicas using a shared membership view.

Possible uses:

- Shard Prometheus scrape targets
- Divide Kubernetes log targets
- Scale a central collector horizontally
- Reassign work when a replica leaves

Important distinctions:

- Enabling Alloy clustering does not automatically shard every component.
  Individual components must support and opt into clustering.
- Stateless receivers may instead need an ordinary network load balancer.
- A source moving between replicas can sometimes replay a bounded amount of
  data, depending on how positions are stored.
- Scaling a DaemonSet is achieved by adding nodes, not by increasing an
  arbitrary replica count.

For central Prometheus collection, persistent identity and WAL storage make a
StatefulSet a strong default. For a traces-only gateway with no important local
state, a Deployment can be enough.

## State and persistence

`alloy.storagePath` is where components keep local state. Depending on the
pipeline, this may include:

- Prometheus remote-write WAL data
- Log file read positions
- Component-specific queues or caches
- State needed to resume after restart

The chart defaults to an ephemeral path. That is acceptable when:

- The source can safely replay
- Small duplicate or missing windows are tolerable
- The pipeline is effectively stateless

Use persistent volumes when a WAL or positions state is operationally
important. With a StatefulSet, configure `controller.volumeClaimTemplates` and
mount the claim at the selected `storagePath`.

## Log collection choices

Alloy offers multiple ways to acquire logs.

| Source | Best fit | Main tradeoff |
| --- | --- | --- |
| `loki.source.kubernetes` | Pod logs through the Kubernetes API | No host mount, but more kubelet/API CPU and network use |
| `loki.source.file` | Node-local files such as `/var/log/pods` | Efficient local reads, but requires DaemonSet and mounts |
| `loki.source.podlogs` | CRD-driven Pod selection | Delegated declarative selection, but adds CRDs |
| `loki.source.kubernetes_events` | Kubernetes Events | Events are API objects, not ordinary Pod logs |
| `loki.source.journal` | systemd journal | Needs host journal access |
| `loki.source.syslog` | Network devices and traditional servers | Requires a reachable listening port and syslog parsing |
| `loki.source.docker` | Docker Engine logs | Coupled to Docker API/runtime access |
| OTLP receiver components | Instrumented applications sending logs | Applications or SDKs must push OTLP |

After collection, `loki.relabel` and `loki.process` can:

- Add, remove, or normalize labels
- Parse JSON, logfmt, regex, or multiline data
- Drop noisy or unwanted entries
- Extract timestamps and severity
- Redact or transform fields
- Generate metrics from logs
- Route entries to one or more `loki.write` destinations

High-cardinality values such as request IDs should remain in the log body rather
than becoming Loki stream labels.

## Metrics, traces, and profiles

### Prometheus metrics

Alloy can:

- Discover Kubernetes, cloud, Consul, DNS, and static targets
- Run built-in exporters
- Scrape Prometheus/OpenMetrics endpoints
- Relabel and filter samples
- Apply clustering to distribute scrape targets
- Remote-write to Prometheus, Mimir, Grafana Cloud, or compatible systems

Remote write uses a WAL, so disk persistence and shutdown/replay behavior matter
more than they do in a simple traces-only receiver.

### OpenTelemetry

`otelcol.*` components can:

- Receive OTLP over gRPC or HTTP
- Batch, filter, sample, transform, and enrich telemetry
- Add Kubernetes resource attributes
- Route metrics, logs, and traces independently
- Export to OTLP-compatible backends

Stateful processors such as some tail-sampling designs need consistent routing
and careful horizontal scaling. Stateless transformations can sit behind a
normal load balancer more easily.

### Profiling and instrumentation

Alloy includes Pyroscope integrations and instrumentation components such as
Beyla/eBPF features. These can provide profiles or application telemetry without
manually instrumenting every code path, but kernel support and privileges must
be evaluated per environment.

## Configuration outside Helm

Alloy uses its own declarative language, normally stored in `.alloy` files.

### Files and directories

Run one file:

```bash
alloy run /etc/alloy/config.alloy
```

If given a directory, Alloy loads its top-level `*.alloy` files as one
configuration. This supports splitting sources, processors, and destinations
into understandable files.

### Reusable and remote modules

Custom components can be declared in modules and imported from:

- A local file
- A Git repository
- An HTTP endpoint
- A string

Remote modules execute with the privileges of the Alloy process. Pin and review
remote content just as carefully as container images or application code.

Grafana Fleet Management and Alloy remote-configuration features can manage
larger fleets centrally. That is useful when many servers or clusters would
otherwise drift, but it introduces a control-plane dependency.

### Environment and secrets

Alloy expressions can read environment variables and local files. On
Kubernetes, `envFrom`, Secret volume mounts, and external secret managers can
keep credentials out of the main configuration.

Avoid embedding backend tokens directly in a checked-in `values.yaml`.

### Reloading

Alloy can reload configuration after:

- `POST /-/reload`
- `SIGHUP`
- A Helm chart config-reloader detecting ConfigMap changes

The component controller adds, updates, and removes components to match the new
graph. Reloading is not a substitute for testing: a syntactically valid pipeline
can still have unreachable endpoints or insufficient permissions.

### Tooling

Useful commands include:

```bash
alloy fmt --test config.alloy
alloy validate config.alloy
```

`fmt` checks canonical formatting. `validate` checks syntax, component names,
required properties, and other graph-level configuration errors. Runtime
connectivity and credentials still require smoke tests and meta-monitoring.

Alloy can also convert several Prometheus, Promtail, OpenTelemetry Collector,
and older Grafana Agent configuration formats. Conversion should be reviewed;
equivalent syntax does not guarantee identical operational behavior.

## Installation mechanisms

| Method | Typical use |
| --- | --- |
| Kubernetes Helm chart | Cluster discovery, RBAC, Services, reload sidecar, and controller choice |
| Raw manifests or Kustomize | Direct ownership of Kubernetes objects and generated ConfigMaps |
| Linux package/systemd | Host metrics and journal logs on VMs or bare metal |
| Docker or Podman | Local labs, appliances, and container hosts |
| Standalone binary | Minimal hosts or custom supervision |
| macOS or Windows packages/services | Developer machines and Windows telemetry |
| OpenShift | Kubernetes environments with OpenShift security constraints |
| Ansible, Chef, or Puppet | Consistent installation across fleets of machines |

Alloy's pipeline language remains portable, but paths, service discovery,
credentials, capabilities, and network endpoints differ by platform.

## Example environment choices

| Environment | Reasonable starting point |
| --- | --- |
| Developer laptop | One Docker or local Alloy reading test files and sending to local backends |
| This homelab today | One Deployment using Kubernetes API Pod logs |
| Homelab with node journals | DaemonSet with journal and Pod-log host access |
| VM fleet | Alloy systemd service per host, managed with Ansible/Puppet/Chef |
| Application OTLP gateway | Deployment behind a Service, with batching and stateless processors |
| Central Prometheus collector | Clustered StatefulSet with persistent WAL volumes |
| Very large platform | Node agents plus clustered gateway tiers and centrally managed modules |
| Short-lived specialized task | Sidecar, only when locality or lifecycle coupling is genuinely required |

## What would trigger a homelab redesign?

Reconsider the current single Deployment if:

- Kubernetes API log tailing causes measurable kubelet/API pressure
- Host journal or file collection is added
- Losing collector positions during restart causes unacceptable gaps/duplicates
- One Alloy replica cannot keep up
- Alloy begins collecting Prometheus metrics with a meaningful WAL
- Applications start sending enough OTLP traffic to require a load-balanced
  gateway

For the planned host logs, the most natural next step is likely a DaemonSet with
node-local mounts. Kubernetes Events could instead remain a singleton or
cluster-sharded API source.

## Official references

- [Deploy Alloy and choose a topology](https://grafana.com/docs/alloy/latest/set-up/deploy/)
- [Alloy components](https://grafana.com/docs/alloy/latest/reference/components/)
- [Alloy configuration language](https://grafana.com/docs/alloy/latest/get-started/)
- [Alloy modules](https://grafana.com/docs/alloy/latest/get-started/modules/)
- [Alloy run and reload behavior](https://grafana.com/docs/alloy/latest/reference/cli/run/)
- [Alloy configuration validation](https://grafana.com/docs/alloy/latest/reference/cli/validate/)
- [Alloy installation methods](https://grafana.com/docs/alloy/latest/set-up/)
- [Kubernetes access and permissions](https://grafana.com/docs/alloy/latest/access_permissions/kubernetes/)
