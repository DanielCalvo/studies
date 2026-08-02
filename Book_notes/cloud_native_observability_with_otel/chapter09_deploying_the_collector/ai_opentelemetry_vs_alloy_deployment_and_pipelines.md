# OpenTelemetry Collector versus Grafana Alloy: deployment and pipelines

OpenTelemetry Collector and Grafana Alloy have closely related deployment
topologies and telemetry flows, but they expose them through different
configuration models.

The shortest summary is:

```text
Deployment topology:
    Very similar Kubernetes choices.

Telemetry processing:
    Same broad receive -> process -> export idea.

Configuration:
    OpenTelemetry Collector uses named pipelines.
    Alloy connects typed components into a graph.

Alloy clustering:
    An optional coordination feature for supported component workloads,
    not another Kubernetes deployment type.
```

## First terminology correction: gateway is a role

The upstream Collector does not have a Kubernetes workload type named
`gateway`.

```text
Gateway
    Architectural role: centralized telemetry receiver and processor.

Deployment, DaemonSet, StatefulSet, sidecar
    Kubernetes placement and lifecycle choices.
```

For example, our gateway is an `OpenTelemetryCollector` with:

```yaml
spec:
  mode: deployment
```

It is a gateway because agents send cluster telemetry to it, not because
`gateway` appears as a mode.

The current upstream Collector Helm chart supports these `mode` values:

```text
deployment
daemonset
statefulset
```

The OpenTelemetry Operator additionally supports `sidecar` mode. Raw
Kubernetes manifests can, of course, place a Collector container in a Pod
without using either tool.

## Deployment-topology comparison

| Architectural role | OpenTelemetry Collector | Grafana Alloy | Typical use |
| --- | --- | --- | --- |
| Pod-local sidecar | Collector container in the application Pod; the OpenTelemetry Operator can inject a sidecar-mode Collector | Alloy container placed in the application Pod | Specialized or short-lived workloads needing Pod-local collection |
| Node agent or host daemon | DaemonSet | DaemonSet | Pod logs, host metrics, and one node-local receiver |
| Central gateway | Deployment, or sometimes StatefulSet | Deployment for stateless/traces-only processing; StatefulSet when stable identity or persistent data such as a metrics WAL matters | Shared application telemetry ingestion and centralized processing |

The standard Alloy Helm chart expresses the Kubernetes workload through:

```yaml
controller:
  # daemonset, deployment, or statefulset
  type: daemonset
  replicas: 1
```

`replicas` is ignored for a DaemonSet because node eligibility determines its
Pod count. The chart currently defaults to `daemonset`. This was verified
against Alloy chart `1.11.0`, which packages Alloy `v1.18.0`.

Therefore the conceptual mapping is:

```text
OpenTelemetry Helm mode: daemonset
    ~= Alloy controller.type: daemonset

OpenTelemetry Helm mode: deployment
    ~= Alloy controller.type: deployment

OpenTelemetry Helm mode: statefulset
    ~= Alloy controller.type: statefulset

OpenTelemetry Operator mode: sidecar
    ~= manually placing Alloy as a Pod sidecar
```

The configurations are analogous, not interchangeable: their Helm values and
templates use different schemas.

## What should run where?

The decision is driven by data locality, not by which Collector distribution
is used:

```text
Need files or resources from every node?
    Use a DaemonSet.

Need a central OTLP endpoint that scales independently?
    Use a Deployment or StatefulSet behind a Service.

Need Pod-local networking or isolation?
    Consider a sidecar.
```

Grafana recommends a host-daemon topology for node-level data such as Pod logs
and cAdvisor metrics. It describes a centralized service for application
telemetry. For centralized Prometheus metric collection, a StatefulSet is
commonly useful because persistent Pod identity can be paired with the
write-ahead log. A traces-only centralized pipeline can use a Deployment when
persistent storage is unnecessary.

Sidecars are possible with both systems, but Grafana recommends Alloy sidecars
mainly for short-lived applications or specialized cases rather than as the
default for long-running services.

## OpenTelemetry Collector pipeline model

The upstream Collector configuration declares components and then explicitly
activates an ordered signal pipeline under `service.pipelines`:

```yaml
receivers:
  otlp:
    protocols:
      grpc:
        endpoint: 0.0.0.0:4317

processors:
  batch: {}

exporters:
  otlp/tempo:
    endpoint: tempo:4317
    tls:
      insecure: true

service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [batch]
      exporters: [otlp/tempo]
```

Conceptually:

```text
OTLP receiver -> batch processor -> OTLP exporter
```

Declaring a receiver, processor, or exporter is not enough. Referencing it from
an active `service.pipelines` entry is what makes it part of that pipeline.

## Alloy component-graph model

Alloy follows the same broad data-flow model, but it does not build the default
configuration around one central `service.pipelines` section. Each labeled
component exposes values or typed inputs, and other components refer to them.

An equivalent simplified Alloy trace flow looks like:

```alloy
otelcol.receiver.otlp "applications" {
  grpc {
    endpoint = "0.0.0.0:4317"
  }

  output {
    traces = [otelcol.processor.batch.default.input]
  }
}

otelcol.processor.batch "default" {
  output {
    traces = [otelcol.exporter.otlp.tempo.input]
  }
}

otelcol.exporter.otlp "tempo" {
  client {
    endpoint = "tempo:4317"

    tls {
      insecure = true
    }
  }
}
```

Conceptually:

```text
receiver output
    -> processor input

processor output
    -> exporter input
```

The references form a directed component graph:

```text
otelcol.receiver.otlp.applications
                  |
                  v
otelcol.processor.batch.default
                  |
                  v
otelcol.exporter.otlp.tempo
```

Alloy configuration is declarative. Block order does not determine execution
order; references express dependencies and data flow.

## Inputs, outputs, filters, and types in Alloy

Your “inputs and outputs” description is correct. An Alloy component generally:

1. Accepts arguments describing its behavior.
2. May consume the exported input of another component.
3. Exports values or receiver-like inputs that another component can use.
4. May route metrics, logs, and traces to different next components.

For OpenTelemetry data, Alloy connects compatible `otelcol.Consumer` values.
The configuration can keep signals together or route them separately:

```alloy
output {
  metrics = [some_metrics_component.input]
  logs    = [some_logs_component.input]
  traces  = [some_traces_component.input]
}
```

A filter is another component in the graph:

```text
receiver
    -> memory limiter
    -> filter
    -> batch
    -> exporter
```

Depending on the signal and ecosystem, Alloy can use component families such
as:

```text
otelcol.*      OpenTelemetry receivers, processors, connectors, exporters
prometheus.*   Prometheus discovery, scraping, relabeling, remote write
loki.*         Log sources, processing, relabeling, Loki writing
pyroscope.*    Profile collection and writing
discovery.*    Target discovery shared by other components
```

This is broader than an ordinary OpenTelemetry-only pipeline. A discovery
component can export targets to a Prometheus scrape component, while an OTLP
receiver can simultaneously feed OpenTelemetry processors.

Alloy checks component compatibility through the types they export and consume.
This is why the graph is more than arbitrary text references: a target list,
Prometheus metrics receiver, Loki logs receiver, and OpenTelemetry consumer are
different kinds of values.

## Pipeline versus component graph

| Question | OpenTelemetry Collector | Grafana Alloy |
| --- | --- | --- |
| Where are components declared? | `receivers`, `processors`, `exporters`, `connectors`, and `extensions` sections | Individually labeled component blocks such as `otelcol.receiver.otlp "apps"` |
| Where is order expressed? | Ordered lists inside `service.pipelines` | References from one component's output to another component's input |
| How is a signal pipeline activated? | It must appear under `service.pipelines` | The connected component graph defines the flow |
| Can signals branch? | Yes, through multiple exporters, connectors, and pipelines | Yes, by forwarding an output to multiple component inputs |
| Can it mix Prometheus, Loki, and OpenTelemetry-native processing? | Some integrations exist in Collector distributions | This cross-ecosystem component graph is a central Alloy capability |
| Is configuration portable to another upstream Collector? | Standard Collector YAML is comparatively portable | Alloy component syntax requires translation |

The mental models are therefore compatible:

```text
OpenTelemetry:
    build an ordered pipeline from registered components.

Alloy:
    build a typed graph by connecting component outputs and inputs.
```

## What Alloy clustering actually does

Clustering is independent of whether Alloy runs as a Deployment, StatefulSet,
or DaemonSet:

```text
Kubernetes controller type
    Controls Pod placement and lifecycle.

Alloy clustering
    Lets Alloy processes discover peers and coordinate supported work.
```

The Helm chart has a separate setting:

```yaml
alloy:
  clustering:
    enabled: true
```

At runtime, Alloy peers use a gossip-based, eventually consistent membership
model. Participating instances are expected to have equivalent configuration.
Clustering-aware components can use consistent hashing to divide work.

The classic example is Prometheus scraping:

```text
Alloy A, B, and C all discover 3,000 targets
                    |
                    | clustering enabled on prometheus.scrape
                    v
A scrapes one subset
B scrapes one subset
C scrapes one subset
```

When membership changes, target ownership is recalculated. Remaining peers can
take over targets from a failed peer.

Clustering must also be enabled on the particular component when that component
supports it:

```alloy
prometheus.scrape "pods" {
  targets = discovery.kubernetes.pods.targets

  clustering {
    enabled = true
  }

  forward_to = [prometheus.remote_write.default.receiver]
}
```

Enabling Alloy clustering does not automatically shard every declared
component.

## Comparison with upstream Collector coordination

Ordinary upstream Collector replicas do not form a general-purpose peer cluster
equivalent to Alloy clustering. They can still scale, but coordination depends
on the workload:

```text
Pushed OTLP data
    Kubernetes Service or another load balancer

Prometheus pull targets managed through the Operator
    Optional OpenTelemetry Target Allocator

Stateful trace routing
    Load-balancing exporter or another trace-aware routing layer
```

Alloy clustering brings peer membership and consistent-hash work distribution
into Alloy itself for components that implement clustering support. It still
does not remove the need for the pushed-data and stateful-processing strategies
described above.

## What clustering does not automatically solve

### It is not Kubernetes scheduling

Clustering does not replace a Deployment, StatefulSet, DaemonSet, Service, or
load balancer. Kubernetes still starts the Pods and routes network traffic.

### It does not automatically balance pushed OTLP spans

For an OTLP receiver, applications push data into a network endpoint. A
Kubernetes Service or another load balancer distributes client connections.
As we observed with the upstream Collector, one persistent gRPC connection can
remain attached to one replica.

```text
Alloy cluster membership
    does not imply
every incoming span is redistributed among Alloy peers
```

### It does not make stateful trace processing trivial

Components such as tail sampling, service graphs, and span metrics need related
spans to meet at an appropriate instance. Those pipelines may require
trace-aware or attribute-aware routing, such as a load-balancing exporter,
rather than ordinary connection balancing.

### It may be pointless for node-local log collection

Each DaemonSet Pod normally reads logs from its own mounted node. There is no
shared global target set to redistribute, so enabling clustering can add
complexity without useful work sharing.

### It is not a storage cluster

Alloy remains a telemetry collector and processor:

```text
Alloy -> Tempo for trace storage
Alloy -> Loki for log storage
Alloy -> Prometheus/Mimir-compatible storage for metrics
```

Alloy peer clustering does not turn Alloy into any of those backends.

## Practical topology for the intended Grafana environment

A reasonable design to investigate later is:

```text
Alloy DaemonSet
    - Pod log collection
    - Node-local infrastructure metrics
    - Possibly a node-local OTLP endpoint
              |
              v
Central Alloy Deployment or StatefulSet
    - Shared OTLP ingestion
    - Batching, filtering, enrichment, routing
    - Optional clustering for components that support useful work sharing
              |
              |-- traces  -> Tempo
              |-- logs    -> Loki
              `-- metrics -> Prometheus-compatible backend
```

This is analogous to the upstream agent-to-gateway architecture:

```text
OpenTelemetry Agent DaemonSet -> OpenTelemetry gateway

Alloy host-daemon DaemonSet   -> centralized Alloy service
```

It is not mandatory to use both layers. If one Alloy DaemonSet can collect and
export everything reliably, a central Alloy tier may be unnecessary. If a
central tier provides useful shared processing, isolation, or scaling, then the
extra hop may be justified.

## Final comparison

```text
OpenTelemetry Collector
    Deployment choices:
        Deployment, DaemonSet, StatefulSet, sidecar
    Processing model:
        receivers/processors/exporters activated by service.pipelines
    Scaling:
        Kubernetes replicas plus explicit routing/sharding strategies

Grafana Alloy
    Deployment choices:
        Deployment, DaemonSet, StatefulSet, sidecar topology
    Processing model:
        typed, labeled components connected as a declarative graph
    Scaling:
        Kubernetes replicas, network load balancing, and optional
        clustering for components that explicitly support it
```

The concepts learned from the upstream Collector remain directly useful:
receiving, processing, exporting, batching, filtering, OTLP forwarding,
resource metadata, Services, persistent gRPC connections, and stateful
processing all still matter. Alloy changes the configuration language and adds
a broader component ecosystem plus optional coordination for supported work.

## Current references

- [OpenTelemetry Collector Helm chart modes](https://opentelemetry.io/docs/platforms/kubernetes/helm/collector/)
- [OpenTelemetry Operator deployment modes](https://opentelemetry.io/docs/platforms/kubernetes/operator/)
- [Grafana Alloy deployment topologies](https://grafana.com/docs/alloy/latest/set-up/deploy/)
- [Grafana Alloy clustering](https://grafana.com/docs/alloy/latest/get-started/clustering/)
- [How Grafana Alloy works](https://grafana.com/docs/alloy/latest/introduction/how-alloy-works/)
- [Grafana Alloy compatible components](https://grafana.com/docs/alloy/latest/reference/compatibility/)
- [Collect and forward OpenTelemetry data with Alloy](https://grafana.com/docs/alloy/latest/collect/opentelemetry-data/)
