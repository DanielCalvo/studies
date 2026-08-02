# Mapping the learning deployment to production Tempo

This document maps the small Chapter 10 Minikube deployment to a plausible
self-managed production Tempo architecture. It is a decision map, not a
drop-in production configuration: actual replica counts, resources, limits,
retention, and storage settings must come from measured workload and business
requirements.

The exercise currently proves the data flow with the smallest understandable
components:

```text
applications
    |
    v
one Alloy Pod
    |
    v
one Tempo Pod, target: all
    |
    +-> one local PVC for trace blocks
    |
    +-> Prometheus for generated metrics

Grafana -> the same Tempo Pod for queries
```

This is useful because every responsibility is visible. It is not highly
available, independently scalable, secured, or sized for sustained production
traffic.

## The likely production topology

Tempo 3 supports monolithic and microservices deployment modes. Microservices
mode is the normal choice when production requires high availability,
independent scaling, or high trace volume. It requires a Kafka-compatible
system. A small production environment with modest volume and no HA requirement
could retain monolithic mode for operational simplicity, but it would retain
the single failure domain and shared resource pool.

### Write path

```text
instrumented applications
          |
          v
Alloy or OpenTelemetry Collector tier
  - batch
  - retry and queue
  - enrich
  - filter or redact
  - sample
          |
          v
authenticated OTLP gateway / load balancer
          |
          v
multiple Tempo distributors
  - receive OTLP
  - validate
  - enforce ingestion limits
  - shard by trace ID
          |
          v
highly available Kafka-compatible cluster
       /          |                \
      v           v                 v
live-stores   block-builders   metrics-generators
  |               |                 |
  |               v                 v
  |         object storage       Prometheus or
  |                              Grafana Mimir
  |
  +-> recent-trace query path
```

In Tempo 3 microservices mode, Kafka is not an optional Collector-style buffer.
It is part of Tempo's architecture. Distributors acknowledge a successful write
after Kafka acknowledges it. Live-stores, block-builders, and
metrics-generators consume the trace stream independently.

### Read path

```text
Grafana or Tempo API client
          |
          v
authenticated query gateway / load balancer
          |
          v
multiple query-frontends
  - accept trace-ID and TraceQL requests
  - split work into jobs
  - queue work fairly
  - merge results
          |
          v
multiple queriers
       /        \
      v          v
live-stores    object storage
recent data    historical blocks
```

The query frontend is the public Tempo query endpoint. Queriers are workers,
not endpoints that Grafana should address directly.

### Backend maintenance path

```text
singleton backend-scheduler
          |
          v
pool of backend-workers
          |
          +-> compact small object-storage blocks
          +-> apply retention
          +-> perform redaction jobs
          +-> maintain the blocklist
```

Tempo 3 replaced the old compactor target with the backend scheduler and
backend workers. The scheduler coordinates work; workers can scale horizontally
when compaction falls behind.

## Current choice compared with production

| Area | Learning deployment | Production direction | Why it changes |
|---|---|---|---|
| Deployment mode | One `target: all` process | Microservices mode when HA or independent scaling is required | Separates write, read, recent-data, block-building, and maintenance failure domains |
| Kubernetes packaging | Handwritten Deployment, Service, and ConfigMap | Usually the official `tempo-distributed` Helm chart for microservices mode; `tempo` chart for monolithic mode | Encodes the many component workloads and Services consistently |
| Trace entry point | Alloy sends directly to `tempo:4317` | Alloy sends through a secured gateway/load balancer to distributor replicas | Stable routing, authentication, TLS, and horizontal distribution |
| Query entry point | Grafana sends directly to `tempo:3200` | Grafana sends through a secured gateway to query-frontends | Keeps internal component APIs private and secures reads |
| Durable ingest | In-process call; no Kafka | HA Kafka-compatible system in microservices mode | Durable decoupling of distributors from downstream consumers |
| Long-term trace storage | `backend: local` on one 1 GiB PVC | S3, GCS, Azure Blob, or an S3-compatible on-premises object store | Shared, durable storage accessible by all relevant components |
| Recent data | One process and temporary runtime directory | Replicated, zone-aware live-stores consuming Kafka | Recent traces remain queryable despite a Pod or zone failure |
| Block creation | Performed inside monolithic Tempo | Independently scaled block-builders | Block construction scales with write volume |
| Query execution | Same process as ingestion | Query-frontends and independently scaled queriers | Expensive queries do not compete directly with ingestion |
| Compaction | In-process backend scheduler and worker | One scheduler plus a scalable worker pool | Controls retention and block count without sharing the ingest Pod |
| Generated metrics | In-process generator to one Prometheus | Metrics-generator replicas to durable Prometheus-compatible storage such as Mimir or managed Prometheus | Generated series need their own HA, retention, and capacity plan |
| Authentication | None | Authenticating reverse proxy or gateway; TLS or mTLS | Tempo does not include an authentication layer |
| Tenancy | Single tenant | Deliberate choice: remain single tenant or enforce trusted `X-Scope-OrgID` assignment | Prevents accidental cross-team data access |
| Configuration secrets | None in the example | Secret manager, workload identity, or Kubernetes Secrets | Object-store and Kafka credentials must not live in a public ConfigMap |
| Availability | One replica and one node | Replicas, topology spread, Pod disruption budgets, multi-zone dependencies | Removes individual Pod and node failure as an outage |
| Sizing | Tiny fixed requests and limits | Measurement-driven component sizing and autoscaling | Each production component has a different resource driver |
| Monitoring | Manual readiness and metric checks | Prometheus scraping, Tempo mixin dashboards, alerts, and synthetic tests | Detects loss, lag, saturation, and query failure before users do |

## Change 1: Choose deployment mode from requirements

Do not choose microservices mode merely because it sounds more
"production-like." It adds Kafka, many workloads, networking, and operational
dependencies.

Keep monolithic mode when all of these are acceptable:

- one Tempo process is an acceptable failure domain
- vertical scaling is sufficient
- query load and ingestion can share CPU and memory
- the expected trace volume is modest
- operational simplicity is more valuable than independent scaling

Move to microservices mode when any of these is a requirement:

- high availability
- isolation between ingestion and expensive queries
- independent scaling of write, query, block-building, or recent-data capacity
- multi-zone recent-trace availability
- traffic beyond a safely tested monolithic capacity

Running multiple replicas of our existing `target: all` Deployment is not a
substitute for microservices mode. The current backend scheduler is a singleton,
and a monolithic replica still combines all resource and failure domains.

## Change 2: Replace local trace storage with object storage

Our configuration currently says:

```yaml
storage:
  trace:
    backend: local
    wal:
      path: /data/tempo/wal
    local:
      path: /data/tempo/blocks
```

All trace blocks belong to one PVC and one Tempo instance. Production
microservices need shared object storage so block-builders, queriers, and
backend workers see the same trace blocks.

An illustrative S3 configuration shape is:

```yaml
storage:
  trace:
    backend: s3
    s3:
      bucket: company-tempo-traces
      endpoint: s3.example.internal
      region: eu-west-1
```

Production decisions include:

- AWS S3, Google Cloud Storage, Azure Blob, or S3-compatible on-premises storage
- bucket availability and durability
- encryption at rest and in transit
- workload identity or secret delivery
- least-privilege read, write, list, and delete permissions
- object-store request rates, latency, and cost
- whether bucket versioning, replication, or backups are required
- coordination between Tempo retention and any bucket lifecycle rules

Local volumes may still exist for component working data, WALs, or caches, but
they no longer form the shared long-term trace database.

## Change 3: Add and operate Kafka

In microservices mode the configuration gains an ingest dependency shaped like:

```yaml
ingest:
  kafka:
    address: tempo-kafka:9092
    topic: tempo-traces
```

The real production design must also decide:

- Kafka, Redpanda, or another supported Kafka-compatible service
- broker and availability-zone count
- topic partition count
- replication and minimum in-sync replicas
- retention long enough for consumers to recover from an outage
- TLS and SASL authentication
- capacity for peak trace bytes, not merely average throughput
- monitoring and alerting on consumer lag
- procedures for partition expansion and Tempo component scaling

Kafka protects the ingest handoff, but it does not replace long-term object
storage. Object storage holds completed trace blocks; Kafka holds the durable
stream that Tempo consumers process.

## Change 4: Separate and scale Tempo components

| Component | Primary responsibility | Main scaling signal |
|---|---|---|
| Distributor | Receive, validate, limit, and write spans to Kafka | Ingest requests, bytes, CPU, rejected spans |
| Live-store | Serve recent traces from the Kafka stream | Kafka partitions, recent-data volume, memory, lag |
| Block-builder | Build Parquet blocks and upload them | Kafka lag and block-building throughput |
| Metrics-generator | Derive span metrics and service graphs | Span rate, active metric series, remote-write throughput |
| Query-frontend | Split, queue, retry, and merge query jobs | Request rate, queue length, latency |
| Querier | Read live-stores and object-storage blocks | Query concurrency, inspected bytes, CPU and memory |
| Backend-scheduler | Coordinate compaction, retention, and redaction | Must remain a singleton; monitor job backlog |
| Backend-worker | Execute backend maintenance jobs | Outstanding blocks, job duration, object-store throughput |

Kubernetes production controls normally include:

- resource requests and limits per component
- at least two replicas where the component supports horizontal availability
- topology spread or anti-affinity across nodes and zones
- Pod disruption budgets
- readiness probes and graceful termination
- autoscaling only from meaningful component-specific signals
- NetworkPolicies restricting internal and external paths
- pinned chart and image versions

Replica counts should not be copied from an example. They follow the actual
ingest rate, Kafka partitioning, object-store performance, query workload, and
availability objective.

## Change 5: Introduce explicit security boundaries

Tempo itself does not provide an authentication layer. A production deployment
should put an authenticating reverse proxy or gateway in front of both:

- the OTLP write path to distributors
- the query path to query-frontends

The gateway is responsible for decisions such as:

- client identity
- authorization
- TLS termination or mTLS
- trusted tenant-header assignment
- request-size and rate controls
- audit logging
- routing write and read APIs to the correct internal Services

Internal Tempo Services should not all be publicly reachable. Use Kubernetes
Services and NetworkPolicies so applications can reach only the intended Alloy
or gateway endpoints and users can query only through the approved query path.

## Change 6: Decide whether multi-tenancy is needed

Tempo enables tenant isolation with:

```yaml
multitenancy_enabled: true
```

Writes and reads must then carry:

```text
X-Scope-OrgID: <trusted-tenant-id>
```

This header is a routing and isolation key, not proof of identity. Clients
should not be trusted to choose arbitrary tenant IDs. The authenticated gateway
should derive or validate the tenant and set the header. Alloy must provide it
on writes, and Grafana must provide it on queries.

Multi-tenancy introduces:

- tenant naming and ownership
- per-tenant limits and retention
- tenant-aware dashboards and alerts
- access-control policy
- incident procedures for a noisy tenant
- potentially federated queries

A single organization can deliberately remain single tenant if its access and
isolation requirements do not justify this complexity.

## Change 7: Define retention and compaction

Retention is a business and cost decision, not merely a Tempo default. Determine
it from:

- incident investigation windows
- regulatory and privacy requirements
- object-storage cost
- trace volume and sampling strategy
- deletion or redaction obligations

An illustrative global setting is:

```yaml
compaction:
  block_retention: 168h
```

This example means seven days; it is not a recommendation. Current Tempo's
default is 336 hours, or 14 days, unless configured otherwise. Per-tenant
retention can be supplied through overrides.

Backend workers compact small immutable blocks into fewer larger blocks and
delete blocks that exceed retention. Monitor at least:

- outstanding blocks
- blocklist length and polling failures
- compaction failures and duration
- retention deletion failures
- object-storage errors

Bucket lifecycle rules must not delete objects earlier than Tempo expects.
Compacted-block retention and backend polling intervals must also remain
consistent so queriers do not reference blocks that disappeared too soon.

## Change 8: Establish ingestion and query limits

The study deployment demonstrates only:

```yaml
max_attribute_bytes: 64
```

A production policy normally evaluates:

- ingestion bytes per second and burst allowance
- maximum trace size
- maximum attribute key and value sizes
- maximum search/query duration and inspected bytes
- outstanding queries per tenant
- metrics-generator active-series and label cardinality limits
- accepted clock skew or late spans

Limits protect availability, but every dropped or truncated span must be
observable. Alert on metrics such as received spans, discarded spans by reason,
and generated-series demand.

Sampling usually remains upstream in Alloy or another collection tier. Tempo
stores what it receives; a well-designed sampling strategy controls cost while
preserving errors, slow requests, and representative normal traffic.

## Change 9: Size from measured traffic

Gather these measurements before choosing production resources:

- average and peak spans per second
- average and peak trace bytes per second
- average and maximum spans per trace
- trace duration and late-span behavior
- expected retention
- number and cardinality of span-metric dimensions
- query requests per second
- typical and worst query time ranges
- bytes inspected by TraceQL queries
- availability-zone and recovery objectives

Then load-test representative traffic and query patterns. Average ingestion
alone is insufficient: burst behavior, large traces, expensive searches, and
metrics cardinality often determine the required headroom.

Capacity planning also covers external dependencies:

- Kafka partitions, storage, network, and recovery window
- object-store throughput and request rate
- Prometheus-compatible metric storage
- caches where query analysis proves they are beneficial
- cluster network bandwidth

## Change 10: Monitor Tempo as a production service

Tempo exposes readiness and Prometheus metrics, and the Tempo mixin supplies
dashboards and alert rules. Production monitoring should cover:

### Write path

- spans and bytes received
- distributor request errors and rejected spans
- Kafka write errors and latency
- consumer lag for live-stores, block-builders, and metrics-generators
- block upload failures

### Read path

- query rate, errors, and latency
- query-frontend queue length
- inspected bytes
- querier concurrency and memory
- live-store readiness and recent-data availability
- object-store read errors and latency

### Backend maintenance

- outstanding blocks
- compaction and retention failures
- blocklist age and polling failures
- backend scheduler and worker job backlog

### Metrics generation

- active and demanded time series
- spans discarded from metric generation
- remote-write failed or dropped samples
- failed or dropped exemplars

Also run a synthetic end-to-end check:

```text
emit known trace -> confirm ingest -> retrieve by trace ID -> alert if absent
```

Readiness alone proves that a process considers itself ready; it does not prove
that an actual trace can traverse Kafka, object storage, and the query path.

## Change 11: Protect trace data

Traces can contain customer identifiers, URLs, database statements, request
metadata, or accidentally captured secrets. Production controls should include:

- an attribute allowlist or redaction policy in Alloy
- avoidance of sensitive values during application instrumentation
- encryption in transit and at rest
- least-privilege storage credentials
- tenant isolation where required
- documented retention and deletion behavior
- restricted access to Grafana and Tempo APIs
- auditability for administrative and query access

Filtering before Tempo is preferable when data must never reach storage.
Backend redaction is a remediation capability, not a substitute for preventing
sensitive collection.

## Change 12: Adopt a managed deployment workflow

For Kubernetes, Grafana provides:

- the `tempo` Helm chart for monolithic mode
- the `tempo-distributed` Helm chart for microservices mode
- the Tempo Operator as another lifecycle-management option

The distributed chart is generally preferable to manually maintaining every
Tempo component Deployment, Service, role, and configuration mapping.

A production workflow should still:

- pin the chart and Tempo image versions
- keep values in version control
- render and review manifests before applying them
- validate configuration in CI
- stage upgrades against representative data
- read release and migration notes
- define rollback limits and procedures
- back up configuration and protect storage credentials

Helm reduces manifest repetition; it does not choose retention, security,
capacity, tenancy, availability, or limits for the operator.

## An illustrative production configuration skeleton

This deliberately incomplete fragment shows where major Tempo choices live. It
is not sufficient to deploy a production cluster:

```yaml
# Each microservice workload receives its own target through chart values or
# command-line arguments; one shared config can describe their connections.

multitenancy_enabled: true

distributor:
  receivers:
    otlp:
      protocols:
        grpc: {}
        http: {}

ingest:
  kafka:
    address: tempo-kafka:9092
    topic: tempo-traces

storage:
  trace:
    backend: s3
    s3:
      bucket: company-tempo-traces
      endpoint: s3.example.internal
      region: eu-west-1

compaction:
  # Example only: choose from actual retention requirements.
  block_retention: 168h

metrics_generator:
  storage:
    remote_write:
      - url: https://metrics.example.internal/api/v1/write
        send_exemplars: true

overrides:
  defaults:
    ingestion:
      max_attribute_bytes: 2048
    metrics_generator:
      processors:
        - span-metrics
        - service-graphs
```

Authentication, authorization, TLS, credentials, replicas, resources, Kafka
durability, object-store policy, and most limits intentionally remain outside
this fragment or use placeholders.

## A practical migration sequence

Moving from this exercise toward production should happen in stages:

1. Measure representative span volume, size, burstiness, query load, and
   required retention.
2. Decide whether monolithic mode satisfies availability and scaling
   requirements. Choose microservices mode when it does not.
3. Provision and validate object storage.
4. For microservices mode, provision and validate the Kafka-compatible system.
5. Define security, tenant, TLS, and network boundaries before accepting
   production traffic.
6. Render a pinned official Helm chart with explicit resources, replicas,
   topology, disruption budgets, and configuration.
7. Deploy Tempo's own monitoring, dashboards, alerts, and synthetic checks.
8. Send a controlled traffic subset and compare accepted, discarded, stored,
   and queryable traces.
9. Load-test ingestion, queries, backend maintenance, and dependency failure.
10. Increase traffic gradually with documented rollback criteria.

## Production readiness questions

Before calling Tempo production-ready, the owning team should be able to answer:

1. What are peak spans and bytes per second?
2. What is the retention period, and who approved it?
3. Which data is prohibited from entering traces?
4. Who authenticates writes and reads?
5. Who assigns and validates tenant IDs?
6. What happens if a distributor, live-store, querier, node, or zone fails?
7. How long can Kafka retain data while consumers recover?
8. Can object storage sustain peak write, query, and compaction traffic?
9. Which limits protect Tempo, and which alerts reveal discarded data?
10. How are Kafka lag, object-store failures, query saturation, and compaction
    backlog detected?
11. How is an upgrade tested and rolled back?
12. Can a synthetic trace be emitted and retrieved end to end?

## Recommended reading

- [Tempo deployment modes](https://grafana.com/docs/tempo/latest/reference-tempo-architecture/deployment-modes/)
- [Tempo architecture](https://grafana.com/docs/tempo/latest/introduction/architecture/)
- [Tempo components](https://grafana.com/docs/tempo/latest/reference-tempo-architecture/components/)
- [Tempo object storage](https://grafana.com/docs/tempo/latest/reference-tempo-architecture/object-storage/)
- [Tempo authentication](https://grafana.com/docs/tempo/latest/operations/authentication/)
- [Tempo multi-tenancy](https://grafana.com/docs/tempo/latest/operations/manage-advanced-systems/multitenancy/)
- [Tempo compaction](https://grafana.com/docs/tempo/latest/reference-tempo-architecture/components/compaction/)
- [Monitor Tempo](https://grafana.com/docs/tempo/latest/operations/monitor/)
- [Deploy Tempo with Helm](https://grafana.com/docs/tempo/latest/set-up-for-tracing/setup-tempo/deploy/kubernetes/helm-chart/)
