# Loki deployment and configuration options

This is a reconnaissance guide to the flexibility available in Grafana Loki. It
is not a replacement for `values.yaml`, which remains the desired configuration
for this homelab.

The examples and Helm value names are based on the versions currently pinned in
this repository:

- Helm chart: `grafana/loki` `7.0.0`
- Loki: `3.6.7`

Chart interfaces change between releases. Before copying a value into a future
chart version, compare it with:

```bash
helm show values grafana/loki --version 7.0.0
```

## Mental model

Loki is both a single executable and a collection of cooperating components.
The same binary can run everything or only selected targets.

Its major responsibilities include:

- Accepting log streams from Alloy and other clients
- Validating, rate-limiting, and distributing writes
- Buffering and compressing log chunks
- Persisting chunks and their index
- Planning, splitting, and executing LogQL queries
- Applying retention and compacting stored indexes
- Optionally evaluating alerting and recording rules

Loki indexes stream labels, not every word or JSON field in a log line. Storage
design, label cardinality, ingestion volume, and query concurrency therefore
matter more than they would in a system that builds a full-text index.

## Helm chart value families

The chart turns `values.yaml` into Kubernetes workloads, Services, storage,
RBAC, and Loki's own YAML configuration. These are the most useful areas to
recognize.

### Deployment shape

| Value family | What it controls |
| --- | --- |
| `deploymentMode` | Selects `SingleBinary`, `SimpleScalable`, or `Distributed` in the pinned chart |
| `singleBinary` | Replica count, persistence, resources, scheduling, environment, and arguments for monolithic Loki |
| `read`, `write`, `backend` | The grouped workloads used by Simple Scalable mode |
| `distributor`, `ingester`, `querier`, `queryFrontend`, `queryScheduler`, `compactor`, `indexGateway`, `ruler` | Independently deployed components for Distributed mode |
| `bloomBuilder`, `bloomPlanner`, `bloomGateway` | Optional and comparatively advanced search-acceleration components |

Each workload family generally offers some combination of:

- `replicas` and autoscaling
- CPU and memory `resources`
- persistent volumes where state is involved
- image overrides
- extra CLI arguments and environment variables
- pod affinity, node selectors, tolerations, topology spreading, and priority
- update strategies and disruption budgets

For example, a distributed installation could give write-heavy ingesters more
memory while scaling queriers independently for a read-heavy dashboard workload.

### Loki's own configuration

The `loki` section produces Loki's `config.yaml`.

| Value | Purpose |
| --- | --- |
| `loki.auth_enabled` | Enables Loki multi-tenancy and the expectation of tenant IDs |
| `loki.commonConfig` | Shared paths, ring settings, replication factor, and common storage behavior |
| `loki.schemaConfig` | Selects index schema, index store, object store, prefixes, and schema start dates |
| `loki.storage` | Supplies filesystem, S3, GCS, Azure, Swift, or related storage settings exposed by this chart |
| `loki.limits_config` | Ingestion, stream, query, retention, and per-tenant default limits |
| `loki.compactor` | Index compaction, retention, and delete-request behavior |
| `loki.ingester` | Chunk encoding, flushing, WAL, lifecycler, and ingestion behavior |
| `loki.querier`, `loki.query_range`, `loki.frontend` | Query concurrency, splitting, caching, scheduling, and timeouts |
| `loki.rulerConfig` | LogQL alerting and recording-rule evaluation |
| `loki.runtimeConfig` | Configuration that Loki periodically reloads without a full restart |
| `loki.server` | HTTP/gRPC ports, timeouts, log settings, and server behavior |
| `loki.tracing`, `loki.analytics` | Loki's own tracing and usage reporting |

The chart offers three ways to supply the generated configuration:

1. Use its normal `loki.*` value sections. This is the easiest approach to
   maintain across chart upgrades.
2. Use `loki.structuredConfig`. This replaces the chart's generated structure,
   so the operator becomes responsible for providing a complete valid
   configuration.
3. Point the chart at an existing ConfigMap or Secret through its configuration
   object settings.

The first approach is preferable unless the chart does not expose a required
Loki option.

### Entry points and network exposure

The `gateway` section can deploy an NGINX reverse proxy in front of Loki. It can:

- Give clients one stable read/write endpoint
- Route requests to the correct component in scalable modes
- Expose a ClusterIP, LoadBalancer, NodePort, or Ingress
- Add basic authentication at the proxy layer
- Carry per-tenant headers
- Scale independently from Loki

The gateway is unnecessary in this homelab's one-process, trusted-network setup,
so Alloy and Grafana use the internal Loki Service directly.

Loki itself does not include an authentication layer. Internet-facing or
multi-user installations need an authenticating reverse proxy or another
trusted gateway in front of it.

### Caching and query acceleration

The chart can deploy:

- `chunksCache`: Memcached for recently accessed chunks
- `resultsCache`: Memcached for query results
- Optional L2 chunk caching
- Bloom-related components for advanced search acceleration

Caches can improve repeated queries and large installations, but their default
memory sizing can be inappropriate for a small ARM64 cluster. Both Memcached
caches are disabled here to keep the system understandable and lightweight.

### Storage helpers

The chart can optionally install MinIO, giving Loki an S3-compatible object
store inside Kubernetes. This is convenient for demonstrations or isolated
environments, but it adds another stateful system to operate.

For durable production use, an existing object store is normally a better
choice than filesystem storage or a single in-cluster MinIO instance.

### Rules and alerting

The `ruler` and `sidecar.rules` areas can load LogQL rules from ConfigMaps,
Secrets, local files, or supported object stores. Loki rules can:

- Alert on log-derived conditions
- Produce metrics from log queries with recording rules
- Send alerts to Alertmanager
- Remote-write recorded metrics to Prometheus or Mimir

The Ruler's WAL and rule files become operational state. A serious rule
deployment normally uses persistent storage and avoids frequent Ruler churn.

### Meta-monitoring and validation

The chart can create:

- A `ServiceMonitor` for Prometheus
- Loki recording and alerting rules
- Grafana dashboard ConfigMaps
- Loki Canary workloads that continuously write and query known log entries
- Helm test resources

The homelab enables the ServiceMonitor but disables the canary and chart tests to
keep the first workload small. A stricter environment could enable the canary
to detect failures across the complete ingest-to-query path.

### Kubernetes controls

The chart also exposes common platform features:

- ServiceAccount and RBAC creation
- NetworkPolicies
- Ingress and Service types
- Pod/container security contexts
- Image registries, tags, pull policies, and pull secrets
- Affinity, anti-affinity, topology spreading, and zone-aware replication
- Persistent volume class, size, access mode, and retention behavior
- Pod disruption budgets and rollout behavior

These values do not change LogQL or Loki's storage model, but they determine how
well a deployment survives node maintenance, zone failure, or security
constraints.

## Deployment modes

### Single binary / monolithic

All components run in one process with target `all`.

Good fit:

- Learning and development
- Small clusters
- Low log volume
- Environments where operational simplicity is more valuable than HA

Tradeoffs:

- One process is both the read and write path
- Filesystem storage ties the service to one disk/node
- Query and ingestion resources compete in the same process

This is the current homelab mode. The official documentation describes
monolithic mode as suitable for small volumes, roughly up to tens of gigabytes
per day.

Monolithic does not inherently require filesystem storage. Multiple monolithic
replicas can share an object store and ring state, although Distributed mode is
usually clearer once substantial horizontal scaling is required.

### Simple Scalable

Loki is split into `read`, `write`, and `backend` targets.

Benefits:

- Read and write capacity can scale separately
- Fewer workload types than fully distributed Loki
- Useful stepping stone for medium-sized systems

Costs:

- Requires object storage
- Needs a gateway to route API paths
- Adds stateful rings and more moving pieces

Simple Scalable mode is being deprecated upstream and is planned for removal in
Loki 4.0. It should not be selected for a new long-lived design without a
migration plan.

### Distributed / microservices

Distributors, ingesters, queriers, query frontends, schedulers, compactors,
index gateways, rulers, and optional components run as separate workloads.

Benefits:

- Fine-grained scaling and resource isolation
- High availability and zone-aware replication
- Better control for very large ingestion and query workloads

Costs:

- Most complex deployment and upgrade model
- More Services, rings, caches, stateful components, and alerts
- Requires object storage and careful capacity planning

This is mainly appropriate for large production environments or teams that need
precise scaling and failure-domain control.

## Storage choices

Modern Loki stores two related things:

- Compressed log chunks
- A TSDB index that maps label streams and time ranges to chunks

TSDB is the recommended current index store. A `schemaConfig` entry has a start
date so schema changes can be introduced without rewriting older data.

### Filesystem

Advantages:

- Simplest backend
- No separate object-store service
- Excellent for labs, local development, and small installations

Limitations:

- Durability is only as good as the underlying disk
- Difficult to make highly available
- Not appropriate for independently scaled components unless storage is shared
- A large number of chunk files eventually becomes a filesystem concern

The homelab uses a 10 GiB `local-path` PVC and seven-day Compactor retention.

### Object storage

The pinned chart exposes configurations for common backends including:

- Amazon S3 and S3-compatible services
- Google Cloud Storage
- Azure Blob Storage
- OpenStack Swift
- MinIO through its optional subchart

Current Loki documentation also describes other supported object-store
integrations. Check the exact Loki and chart versions before selecting one.

Object storage makes stateless query components and replicated ingestion much
easier. Cloud identity, credentials, bucket lifecycle rules, encryption,
latency, and request cost then become part of the design.

Do not use a blind bucket lifecycle policy to delete the whole Loki bucket.
Loki's Compactor understands indexes and retention; an object store lifecycle
rule does not.

### Retention

Retention is handled by the Compactor for current TSDB deployments:

- `limits_config.retention_period` supplies the default period
- Per-stream and per-tenant policies can keep different data for different
  durations
- `compactor.retention_enabled` activates deletion
- The Compactor needs durable marker/working storage

Loki does not automatically delete data merely because a disk is nearly full.
Disk-capacity monitoring remains necessary.

## Configuration without the Helm chart

Helm is only one packaging and lifecycle mechanism. Loki itself is configured by
a YAML file plus command-line flags.

### YAML and flags

Run Loki with a configuration file:

```bash
loki -config.file=/etc/loki/loki.yaml
```

CLI flags override YAML values. Useful diagnostic flags can print the effective
configuration, including defaults and overrides.

Environment placeholders such as `${S3_ENDPOINT}` can be used in YAML when Loki
is started with:

```bash
-config.expand-env=true
```

This works well with Kubernetes Secrets, systemd environment files, or container
environment variables.

### Runtime configuration

A separate runtime configuration file can be polled and reloaded without
restarting Loki. It is commonly used for:

- Per-tenant ingestion and query limits
- Per-tenant stream limits
- Temporary operational overrides
- KV-store migration controls

It does not make every Loki setting dynamically reloadable. Storage schemas and
most process topology changes still require careful rollout.

### Installation mechanisms

| Method | Typical use |
| --- | --- |
| Helm | Kubernetes lifecycle, Services, RBAC, persistence, and scalable modes |
| Raw Kubernetes manifests or Kustomize | Teams that want direct ownership of every Kubernetes object |
| Tanka/Jsonnet | Large or highly customized Kubernetes deployments |
| Docker or Docker Compose | Local labs, integration environments, and demonstrations |
| Standalone binary with systemd | A single server or VM |
| Build from source | Development, patches, or unusual platforms |
| Managed Loki/Grafana Cloud | Avoid operating the storage and query backend |

The Loki YAML concepts remain largely the same across these mechanisms; the
surrounding process supervision, networking, secrets, and storage attachment
change.

## Example environment choices

| Environment | Reasonable starting point |
| --- | --- |
| Laptop experiment | Docker Compose, monolithic Loki, filesystem storage |
| This two-node homelab | Helm, one single-binary StatefulSet, local PVC, short retention |
| Small production cluster | HA monolithic or Distributed mode with managed object storage |
| High ingest but light querying | More ingester/distributor capacity than querier capacity |
| Dashboard-heavy shared service | Scale query frontend, scheduler, queriers, and caches |
| Regulated multi-team environment | Multi-tenancy, authenticated gateway, object-store encryption, tenant limits, audit-conscious retention |
| Very large multi-zone platform | Distributed mode, object storage, zone-aware replication, caches, canary, and extensive meta-monitoring |

## What would trigger a homelab redesign?

The current configuration is intentionally simple. Reconsider it if:

- The Loki PVC grows too quickly even with retention
- Loss of one node's disk is no longer acceptable
- Queries noticeably interfere with ingestion
- More clusters need to write to the same backend
- Log volume grows into multiple gigabytes per day
- Authentication or tenant separation becomes necessary

The most likely next storage step would be an S3-compatible backend on separate
durable hardware. The most likely topology would remain monolithic until actual
resource or availability evidence justifies Distributed mode.

## Official references

- [Loki deployment modes](https://grafana.com/docs/loki/latest/get-started/deployment-modes/)
- [Loki Helm chart concepts](https://grafana.com/docs/loki/latest/setup/install/helm/concepts/)
- [Loki storage](https://grafana.com/docs/loki/latest/configure/storage/)
- [Loki configuration reference](https://grafana.com/docs/loki/latest/configure/)
- [Loki installation methods](https://grafana.com/docs/loki/latest/setup/install/)
- [Loki alerting and recording rules](https://grafana.com/docs/loki/latest/alert/)
- [Loki retention](https://grafana.com/docs/loki/latest/operations/storage/retention/)
