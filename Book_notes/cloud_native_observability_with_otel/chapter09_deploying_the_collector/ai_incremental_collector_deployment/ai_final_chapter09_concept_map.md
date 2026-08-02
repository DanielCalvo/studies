# Chapter 9 final concept map

Chapter 9 is primarily about two independent decisions:

```text
1. Where should each Collector run?
2. How should those Collectors be installed and managed?
```

Choosing a gateway does not require choosing Helm or the Operator. Choosing the
Operator does not require a gateway. Placement and management can be combined
in different ways.

## 1. Collector placement

```text
Application Pod                 Kubernetes node              Cluster
┌────────────────────┐         ┌────────────────────┐        ┌────────────────────┐
│ Application        │         │ Application Pods   │        │ Gateway replicas   │
│ Collector sidecar  │         │ Agent DaemonSet    │        │ behind a Service   │
└────────────────────┘         └────────────────────┘        └────────────────────┘
       Pod-local                      Node-local                   Cluster-level
```

| Placement | Kubernetes shape | Scope | Main benefit | Main cost or limitation |
| --- | --- | --- | --- | --- |
| Sidecar | Extra container in each application Pod | One Pod | `localhost` networking and lifecycle alignment | One Collector per Pod and no independent scaling |
| Agent | DaemonSet, normally one Pod per eligible node | One node | Shared node-local collection and access to host resources | Replica count follows nodes rather than telemetry load |
| Gateway | Deployment or StatefulSet behind a Service | Cluster or environment | Central processing and independent scaling | Adds a network hop and must handle shared load |

A real design may use one placement or combine them:

```text
application -> sidecar -> backend
application -> agent -> backend
application -> gateway -> backend
application -> agent -> gateway -> backend
```

Our final experiment uses the last path.

## 2. The running telemetry path

```text
Greeting application
    |
    | OTLP/gRPC to the current node IP:4317
    v
Agent Collector DaemonSet
    |
    | k8s_attributes processor
    | adds namespace, Pod, Deployment, and node identity
    |
    | OTLP/gRPC to chapter9-gateway-collector:4317
    v
Operator-managed gateway Deployment
    |
    | memory_limiter -> batch
    v
Debug exporter
```

Collector-to-Collector forwarding is still ordinary OTLP. The agent is an OTLP
server to the application and an OTLP client to the gateway:

```text
application --OTLP--> [agent receiver -> processor -> exporter]
                                      --OTLP--> [gateway receiver -> processors -> exporter]
```

The metadata added by the agent remains attached when the span reaches the
gateway.

## 3. Kubernetes metadata enrichment

Pods are ephemeral, so a Pod IP alone is a poor long-term workload identity.
The Kubernetes attributes processor turns runtime identity into queryable
resource attributes:

```text
k8s.pod.ip association hint
          |
          v
Kubernetes API metadata cache
          |
          v
k8s.namespace.name
k8s.pod.name
k8s.pod.uid
k8s.deployment.name
k8s.node.name
```

The agent requires read-only Kubernetes API permissions for this lookup.
Authorization to read metadata is separate from configuring the processor in
the trace pipeline; both are necessary.

## 4. Discovery and Services

A Kubernetes Service gives senders a stable destination while Collector Pods
come and go:

```text
chapter9-gateway-collector:4317
                  |
                  |-- gateway Pod 1
                  `-- gateway Pod 2
```

Important distinction:

```text
Kubernetes Service
    Network discovery and routing to Pods.

Collector service.pipelines
    Activates telemetry receivers, processors, and exporters.
```

They share the word “service” but solve unrelated problems.

The Operator generated three Services for the gateway:

| Service | Purpose |
| --- | --- |
| `chapter9-gateway-collector` | Stable OTLP client destination |
| `chapter9-gateway-collector-headless` | Direct Pod-address discovery |
| `chapter9-gateway-collector-monitoring` | Collector internal telemetry |

The monitoring Service does not carry the application's OTLP spans.

## 5. Gateway scaling

Gateway replicas scale independently from application replicas and node count:

```text
Application replicas: 1
Agent replicas:       1 per node
Gateway replicas:     2, chosen independently
```

A Service balances connections, not individual spans. OTLP/gRPC normally uses
a persistent connection:

```text
two Service endpoints
          does not imply
one established connection alternates spans between both Pods
```

This matters for stateful processing. A processor such as tail sampling needs
related spans from the same trace to reach the same processing instance.
Scaling such pipelines may require trace-aware routing. Stateless processing is
generally easier to distribute.

## 6. Management method

The same Collector role and pipeline can be managed in different ways:

| Method | Input | Who creates resources? | Continuous OpenTelemetry-specific reconciliation? |
| --- | --- | --- | --- |
| Raw manifests | Deployment, ConfigMap, Service YAML | `kubectl` submits them directly | No |
| Collector Helm chart | Chart values | Helm renders and submits ordinary resources | No resident Helm controller |
| OpenTelemetry Operator | `OpenTelemetryCollector` custom resource | The Operator creates owned child resources | Yes |

The pipeline configuration remains recognizable in every method:

```text
receivers -> processors -> exporters
```

### Raw manifests

Raw YAML makes every Kubernetes object explicit. It is excellent for learning
and offers complete control, but we must maintain all repetitive wiring.

### Collector Helm chart

```text
deployment-values.yaml + chart templates
                    |
                    | helm install/upgrade
                    v
Deployment + ConfigMap + Service + ServiceAccount
```

Helm reduces scaffolding and packages defaults. It recalculates the release
when we invoke Helm; it does not continuously watch the values file.
Kubernetes' built-in controllers still maintain the resulting Deployments and
Pods.

### OpenTelemetry Operator

```text
Helm release
    |
    `-- installs Operator Deployment + CRDs + webhook
            |
            `-- watches OpenTelemetryCollector CR
                    |
                    `-- reconciles Collector child resources
```

The Operator installation and an Operator-managed Collector are separate
layers:

```text
Helm owns the Operator installation.
Operator owns the Collector resources.
```

We observed this directly by scaling the generated gateway Deployment from two
replicas to one. Because the `OpenTelemetryCollector` CR still declared two,
the Operator immediately restored the Deployment to two.

## 7. Gateway role versus Deployment mode

These terms answer different questions:

```text
Gateway
    Architectural role: centralized telemetry aggregation and processing.

mode: deployment
    Kubernetes placement: Pods managed by a Deployment.
```

Our `OpenTelemetryCollector` therefore describes a gateway by configuring:

```yaml
spec:
  mode: deployment
  replicas: 2
  config:
    # The gateway's Collector pipeline
```

`gateway` is not a special Operator mode. The same Deployment mode could run a
Collector with a different responsibility.

## 8. Relationship to Grafana Alloy

The placement lessons transfer to a future Alloy deployment:

```text
OpenTelemetry agent or gateway concepts
                    |
                    v
Alloy host daemon or centralized collection topology
```

For the intended Grafana environment, a plausible later design is:

```text
Alloy DaemonSet
    Collect node-level metrics and Pod logs.

Alloy Deployment or StatefulSet
    Receive and centrally process application telemetry.

Tempo
    Store traces.

Loki
    Store logs.

Prometheus-compatible backend
    Store metrics.
```

Alloy can receive OTLP and connect receiver, processor, and exporter components,
so the pipeline concepts remain useful even though Alloy uses its component
configuration language. Its deployment topology still depends on collection
scope: node-level collection favors a DaemonSet, while centralized traces can
use a Deployment when persistent storage is unnecessary.

Alloy is a Collector, not a telemetry backend. Tempo, Loki, and a
Prometheus-compatible database remain responsible for storage and querying.

A more detailed comparison of Alloy's Kubernetes controller types, component
graph, and clustering behavior is available in:

```text
../ai_opentelemetry_vs_alloy_deployment_and_pipelines.md
```

## 9. Decision checklist

```text
Need localhost isolation for one workload?
    Consider a sidecar.

Need node files, host metrics, or one shared node receiver?
    Consider an agent DaemonSet.

Need centralized processing or independent scaling?
    Consider a gateway.

Need transparent, explicit Kubernetes objects?
    Start with raw manifests.

Want packaged Kubernetes scaffolding?
    Consider the Collector Helm chart.

Want OpenTelemetry CRDs, reconciliation, or auto-instrumentation management?
    Consider the OpenTelemetry Operator.

Using Grafana's telemetry stack and collecting several signal types?
    Evaluate which responsibilities belong in Alloy DaemonSets and which
    belong in centralized Alloy instances.
```

## 10. Deliberately outside this exercise

- A real Tempo, Loki, or metrics backend
- TLS and authentication between Collectors
- Persistent queues and failure recovery
- Horizontal Pod Autoscaling
- Trace-aware load balancing and tail sampling
- Operator-managed application auto-instrumentation
- A production Grafana Alloy deployment

## Current references

- [OpenTelemetry Collector Helm chart](https://opentelemetry.io/docs/platforms/kubernetes/helm/collector/)
- [OpenTelemetry Operator for Kubernetes](https://opentelemetry.io/docs/platforms/kubernetes/operator/)
- [Grafana Alloy deployment topologies](https://grafana.com/docs/alloy/latest/set-up/deploy/)
- [Receiving and forwarding OpenTelemetry data with Alloy](https://grafana.com/docs/alloy/latest/collect/opentelemetry-data/)
