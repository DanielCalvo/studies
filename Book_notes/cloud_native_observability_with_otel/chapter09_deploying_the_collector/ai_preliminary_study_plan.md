# Chapter 9 preliminary study plan

This is an adaptable plan for studying Chapter 9, *Deploying the Collector*.
It is a guide rather than a fixed implementation checklist: we can split,
combine, revise, or omit steps as we learn from the running system.

Chapter 8 taught us how to configure a Collector pipeline. Chapter 9 changes
the main question from:

```text
How does a Collector process telemetry?
```

to:

```text
Where should Collectors run in a Kubernetes environment,
and what responsibility should each placement have?
```

The book introduces three primary Collector deployment patterns:

```text
Application Pod       Kubernetes Node          Cluster
┌───────────────┐    ┌──────────────────┐    ┌──────────────┐
│ Application   │───▶│ Agent Collector  │───▶│ Gateway      │
│ + Sidecar     │    │ (DaemonSet)      │    │ Collectors   │
└───────────────┘    └──────────────────┘    └──────┬───────┘
                                                    │
                                                    ▼
                                                 Backend
```

These components do not always need to be used together in production. We will
preserve the sidecar and agent deployments as separate runnable alternatives
so their responsibilities can be compared directly. Collector-to-Collector
forwarding can remain an optional additional scenario rather than requiring
every pattern to run in one chain.

We will use a local Minikube cluster. Docker, `kubectl`, Helm, and Minikube are
already available on this computer.

Each implemented step should be introduced separately, commented, tested, and
recorded in an `ai_incremental_steps_taken.md` file in the eventual example
directory.

## Step 1: Deploy the smallest ordinary application

Create a tiny application and deploy it to Minikube without an OpenTelemetry
Collector.

This gives us an ordinary Kubernetes baseline consisting of:

```text
Deployment -> Pod -> application container
```

We will inspect the Pod and its logs before introducing telemetry
infrastructure.

Expected result: the application Pod runs successfully and its ordinary output
is visible with `kubectl logs`.

Deliberately not introduced yet: a Collector, sidecar containers, Collector
configuration, Services, or telemetry forwarding.

## Step 2: Add a sidecar Collector

Add a Collector container to the same Pod as the application:

```text
Pod
├── application
└── Collector sidecar
```

The application will send OTLP telemetry to:

```text
localhost:4317
```

This demonstrates that containers in one Pod share the Pod's network namespace.
It also demonstrates the main reason for the sidecar pattern: the application
can hand telemetry to a nearby Collector through a stable, low-latency local
destination.

Expected result: the sidecar Collector's debug exporter prints telemetry
produced by the application.

## Step 3: Supply Collector configuration with a ConfigMap

Store the sidecar Collector configuration in a Kubernetes ConfigMap and mount
it into the Collector container.

This demonstrates:

- Separating configuration from the container image
- ConfigMap volumes and volume mounts
- Selecting the mounted configuration when the Collector starts
- Updating Collector behavior without rebuilding the application

Expected result: the sidecar starts with the mounted pipeline configuration and
continues to receive application telemetry.

## Step 4: Deploy an agent Collector as a DaemonSet

Deploy a Collector using a DaemonSet:

```text
one eligible Kubernetes node -> one agent Collector Pod
```

This demonstrates that an agent is associated with a node rather than with one
application Pod. It can provide a node-local aggregation point and collect
system-level information for that node.

Expected result: `kubectl get daemonsets` shows the desired and ready agent
counts matching the eligible nodes in the local cluster.

Initially, the agent will use a debug exporter so that its received telemetry
can be observed without installing a backend.

## Step 5: Send application telemetry directly to the agent

Make the application-only Deployment discover its node address and configure
the application's OTLP exporter to send directly to the node agent:

```text
application -> node agent Collector -> debug exporter
```

This demonstrates:

- An application reaching a Collector outside its own Pod
- Selecting the Collector running on the application's node
- Several application Pods being able to share one node agent
- The resource-saving trade-off of using an agent instead of one sidecar per
  application Pod

Expected result: application spans appear in the agent Collector's debug
output.

The archived sidecar manifests will remain unchanged and runnable for
comparison. If useful later, Collector-to-Collector forwarding can be added as
a separate `sidecar -> agent` scenario rather than overwriting either example.

## Step 6: Add Kubernetes resource metadata

Enrich telemetry with Kubernetes resource attributes such as:

```text
k8s.namespace.name
k8s.pod.name
k8s.deployment.name
k8s.node.name
```

This demonstrates why Kubernetes metadata is useful for querying and grouping
telemetry after it leaves an ephemeral Pod.

The current Collector commonly performs this enrichment with the Kubernetes
attributes processor. We will compare that with the book's simpler example,
which manually injects selected metadata from environment variables.

In the local Minikube environment, `hostPort` rewrites the connection source
address. The application therefore supplies only `k8s.pod.ip` as an association
hint; the agent uses that hint and read-only Kubernetes API access to discover
the remaining metadata.

Expected result: the agent's exported telemetry contains attributes identifying
the Kubernetes workload and node that produced it.

## Step 7A: Deploy a gateway Collector

Deploy a central Collector as a Kubernetes Deployment and expose it through a
Kubernetes Service:

```text
agent Collector
      |
      | OTLP
      v
otel-gateway:4317
      |
      v
gateway Collector
```

This demonstrates:

- A gateway as a cluster-level aggregation and processing point
- Kubernetes Service discovery
- Separating node-local collection from centralized processing
- Creating a controlled place for exporting telemetry outside the cluster

Expected result: one gateway Pod is ready behind a stable Kubernetes Service.
It does not receive application telemetry yet.

## Step 7B: Forward agent telemetry to the gateway

Replace the agent's debug output with an OTLP exporter targeting:

```text
chapter9-gateway:4317
```

Expected result: telemetry originating in the application travels through the
agent and is finally printed by the gateway's debug exporter. The sidecar
remains a separate alternative scenario.

No external backend is required for this step.

## Step 8: Scale the gateway

Increase the gateway Deployment's replica count and inspect the resulting Pods
and Service endpoints.

This demonstrates why the gateway pattern can scale independently of
applications and nodes.

It also introduces an important limitation: Kubernetes can distribute
connections among gateway replicas, but processors that need all spans from one
trace—such as tail sampling—may require trace-aware load balancing or consistent
routing.

Expected result: multiple gateway replicas run behind one stable Kubernetes
Service.

We may inspect Horizontal Pod Autoscaler configuration, but generating enough
load and installing supporting components should remain optional unless it
materially improves the lesson.

## Step 9: Examine the official Helm chart

After implementing the underlying Kubernetes resources ourselves, inspect how
the current OpenTelemetry Collector Helm chart represents:

- Deployment and DaemonSet modes
- ConfigMaps
- Services and ports
- Resource limits
- Presets, permissions, and mounted host resources

The objective is to understand what Helm generates for us, rather than treating
the chart as unexplained scaffolding.

Expected result: we can relate the rendered Helm resources to the raw resources
used in the preceding steps.

## Step 10: Examine the OpenTelemetry Operator

Briefly study how the OpenTelemetry Operator manages Collector resources and
supports application auto-instrumentation.

We do not need to install the Operator immediately. The first objective is to
understand the additional responsibility it assumes compared with applying
ordinary manifests or installing a Helm chart.

## Final concept map

At the end of the exercise, create a concise concept map covering:

- Sidecar, agent, and gateway responsibilities
- Pod-local, node-local, and cluster-level collection
- Collector-to-Collector OTLP forwarding
- Kubernetes metadata enrichment
- Services and Collector discovery
- Scaling and stateful processor considerations
- Raw manifests versus Helm versus the Operator
- How these patterns relate to a future Grafana Alloy deployment

## Compatibility note

The book's Collector images, Helm values, component names, and Kubernetes
examples are dated. We will preserve the concepts from Chapter 9 while checking
the syntax and behavior against the versions we actually use.

In particular, we should not assume that the book's old Helm values map
directly to the current official chart.
