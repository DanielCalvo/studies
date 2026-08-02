# Incremental steps taken

This file records each implemented step in the Chapter 9 Collector deployment
exercise.

## Step 1: Deploy the smallest ordinary application

### What changed

We created:

- A tiny Go HTTP greeting application
- A Dockerfile that packages the application
- A Kubernetes Deployment containing one application container

The application writes ordinary structured logs to standard output. It does not
use OpenTelemetry.

The Deployment's essential structure is:

```yaml
spec:
  replicas: 1
  template:
    spec:
      containers:
        - name: greeting
          image: chapter9-greeting:step1
```

### Concept demonstrated

A Deployment manages the desired application replicas. For this exercise, it
creates one Pod containing one application container:

```text
Deployment -> Pod -> greeting container
```

This is our baseline. Later steps will modify the same Pod template and make
the telemetry path visible.

### Expected output and behavior

After temporarily forwarding port `8080`, this request:

```bash
curl "http://localhost:8080/?name=Daniel"
```

returns:

```text
Hello, Daniel!
```

The application's container logs include:

```text
msg="greeting application started" address=:8080
msg="greeting requested" name=Daniel
```

These messages come directly from the application process through the
container's standard output. No telemetry pipeline processes them.

### Important distinctions

- A Deployment is a Kubernetes workload controller.
- A Pod is the unit Kubernetes schedules.
- A container runs the application process inside the Pod.
- `kubectl logs` reads the container's standard output and standard error; it
  does not require OpenTelemetry.
- `kubectl port-forward` is temporary local access to the workload. It is not a
  Kubernetes Service.

### Deliberately not introduced yet

- OpenTelemetry APIs or SDKs
- OTLP
- An OpenTelemetry Collector
- A sidecar container
- ConfigMaps
- Kubernetes Services
- Agent or gateway Collectors

## Step 2: Add a sidecar Collector

### What changed

We added basic HTTP trace instrumentation to the greeting application. Its OTLP
exporter sends completed spans to:

```go
otlptracegrpc.WithEndpoint("localhost:4317")
```

We also added a Collector as the second container in the Deployment's Pod
template:

```yaml
containers:
  - name: greeting
    image: chapter9-greeting:step2
  - name: collector
    image: chapter9-sidecar-collector:step2
```

The Collector has a minimal trace pipeline:

```text
OTLP/gRPC receiver -> traces pipeline -> debug exporter
```

### Concept demonstrated

A sidecar is a supporting container deployed in the same Pod as the
application container:

```text
Pod
|-- greeting application
`-- Collector sidecar
```

All containers in a Pod share one network namespace and Pod IP. Therefore,
`localhost:4317` from the greeting container reaches the Collector listening on
port `4317` in that same Pod.

This is different from containers in separate Pods: they do not share
`localhost` and normally communicate through Pod addresses or Kubernetes
Services.

### Expected output and behavior

An HTTP request still produces the ordinary application response and log:

```text
Hello, Daniel!
msg="greeting requested" name=Daniel
```

It now also creates an HTTP server span. After the application's batch span
processor exports it over OTLP, the Collector container's debug output includes:

```text
Name: GET /
Kind: Server
service.name: Str(chapter9-greeting)
```

The two outputs are inspected separately because each container has its own
standard-output stream:

```bash
kubectl logs deployment/chapter9-greeting -c greeting
kubectl logs deployment/chapter9-greeting -c collector
```

### Important distinctions

- A sidecar is a deployment pattern, not a special Kubernetes container type.
- Both containers belong to one Pod but run separate processes and have
  separate logs.
- Sharing a network namespace does not mean sharing a filesystem.
- The application log remains an ordinary `slog` record. Only traces are sent
  over OTLP in this step.
- `EXPOSE` and `containerPort` document intended ports; neither creates a
  Kubernetes Service.

### Deliberately not introduced yet

- A ConfigMap: the Collector configuration is temporarily packaged in its
  image so this step can focus on the sidecar relationship.
- OpenTelemetry logs or metrics
- A Kubernetes Service
- An agent Collector or DaemonSet
- A gateway Collector

## Step 3: Supply Collector configuration with a ConfigMap

### What changed

We moved the same minimal Collector pipeline out of the custom Collector image
and into a Kubernetes ConfigMap:

```yaml
kind: ConfigMap
data:
  collector-config.yaml: |
    receivers:
      otlp:
        protocols:
          grpc:
            endpoint: 0.0.0.0:4317
    exporters:
      debug:
        verbosity: detailed
    service:
      pipelines:
        traces:
          receivers: [otlp]
          exporters: [debug]
```

The Deployment now uses the official Collector image and mounts the ConfigMap
key as its configuration file:

```yaml
volumeMounts:
  - name: collector-config
    mountPath: /etc/otelcol-contrib/config.yaml
    subPath: collector-config.yaml
volumes:
  - name: collector-config
    configMap:
      name: chapter9-sidecar-collector
```

### Concept demonstrated

A ConfigMap stores non-secret configuration as a Kubernetes object. A
ConfigMap-backed volume can present each data key to a container as a file.

The relationships in this example are:

```text
ConfigMap data key
        |
        v
ConfigMap-backed Pod volume
        |
        v
Collector container volume mount
        |
        v
/etc/otelcol-contrib/config.yaml
```

This separates the Collector's configuration from its executable image. We can
therefore use the pinned official Collector image without maintaining a custom
image containing only a configuration file.

### Expected output and behavior

The telemetry behavior deliberately remains unchanged:

```text
application -> OTLP over localhost -> sidecar -> debug exporter
```

An application request still creates a `GET /` server span with
`service.name=chapter9-greeting` in the Collector container's logs.

The Kubernetes-level difference is observable with:

```bash
kubectl get configmap chapter9-sidecar-collector
kubectl describe pod -l app=chapter9-greeting
```

The first command shows the separate ConfigMap object. The second shows the
ConfigMap volume and the Collector container's mount.

### Important distinctions

- A ConfigMap stores configuration; it does not run the Collector.
- A volume makes the ConfigMap data available to the Pod.
- A volume mount makes that volume visible inside a particular container.
- `subPath` mounts only the selected file at the requested path rather than
  replacing the entire `/etc/otelcol-contrib` directory.
- Updating a ConfigMap does not necessarily make a running Collector reload its
  configuration. The Collector process must use the new content, commonly
  after a controlled Pod restart or rollout.
- ConfigMaps are intended for non-secret data. Kubernetes Secrets are the
  corresponding API type for sensitive values.

### Deliberately not introduced yet

- OpenTelemetry logs or metrics
- A Kubernetes Service
- An agent Collector or DaemonSet
- Kubernetes metadata enrichment
- A gateway Collector

## Step 4: Deploy an agent Collector as a DaemonSet

### What changed

We preserved the completed sidecar example by renaming its manifest files:

```text
k8s/sidecar/configmap.yaml
k8s/sidecar/deployment.yaml
```

We then created an alternative application-only Deployment:

```yaml
kind: Deployment
spec:
  template:
    spec:
      containers:
        - name: greeting
          image: chapter9-greeting:step2
```

The new agent Collector consists of:

```text
k8s/agent/configmap.yaml
k8s/agent/daemonset.yaml
```

Its OTLP receiver listens on container port `4317`, which is also exposed as
node port `4317` through `hostPort`.

### Concept demonstrated

A DaemonSet ensures that a matching Pod runs on every eligible Kubernetes node:

```text
eligible node 1 -> agent Pod 1
eligible node 2 -> agent Pod 2
eligible node 3 -> agent Pod 3
```

Minikube currently has one node, so the expected DaemonSet state is:

```text
DESIRED=1  CURRENT=1  READY=1
```

This placement differs from a sidecar:

```text
Sidecar Collector
    One Collector container per application Pod
    Scales with application replicas

Agent Collector
    One Collector Pod per eligible node
    Scales with cluster nodes
```

### Expected output and behavior

There are now two separately managed Pods:

```text
Deployment -> application-only Pod -> greeting container

DaemonSet  -> node agent Pod        -> collector container
```

The agent Collector starts its OTLP receiver and reports that it is ready. It
does not receive application spans yet because we have deliberately not changed
the application's exporter destination.

The placement can be observed with:

```bash
kubectl get daemonset chapter9-agent-collector
kubectl get pods -l app=chapter9-agent-collector -o wide
```

### Important distinctions

- A DaemonSet is a Kubernetes workload controller, not an OpenTelemetry
  component.
- The agent is still an ordinary OpenTelemetry Collector; its DaemonSet
  placement is what gives it the agent role.
- A sidecar is another container in an application Pod. This agent is a
  separate Pod owned by a DaemonSet.
- `hostPort: 4317` makes the agent receiver available on its node's IP at port
  `4317`. It does not create a Kubernetes Service.
- Because one agent reserves node port `4317`, Kubernetes cannot schedule a
  second Pod requesting that same host port onto the same node.
- The archived sidecar and current application-only manifests intentionally
  use the same Deployment name. They are alternatives and should not be
  applied as two independent workloads.

### Deliberately not introduced yet

- Forwarding application traces to the agent
- A Kubernetes Service
- System-level host metrics
- Kubernetes metadata enrichment
- A gateway Collector

### Organization refinement

After completing this step, we grouped the manifests by deployment method:

```text
k8s/
|-- sidecar/
|   |-- configmap.yaml
|   `-- deployment.yaml
`-- agent/
    |-- configmap.yaml
    |-- daemonset.yaml
    `-- deployment.yaml
```

This does not change any Kubernetes object or telemetry behavior. It keeps each
deployment method self-contained so it can be read, applied, and compared
without relying on filename suffixes.

## Step 5: Send application telemetry directly to the agent

### What changed

The application no longer hard-codes `localhost:4317` in Go. The exporter now
reads the standard OpenTelemetry environment configuration:

```go
exporter, err := otlptracegrpc.New(ctx)
```

The agent-style application Deployment uses Kubernetes's Downward API to
discover the node on which its Pod is running:

```yaml
env:
  - name: NODE_IP
    valueFrom:
      fieldRef:
        fieldPath: status.hostIP
  - name: OTEL_EXPORTER_OTLP_ENDPOINT
    value: "http://$(NODE_IP):4317"
```

The archived sidecar Deployment sets the same standard variable to
`http://localhost:4317`, allowing both deployment methods to use the same
application image.

### Concept demonstrated

The Downward API exposes selected information about a Pod to its containers.
Here, `status.hostIP` supplies the node address:

```text
Pod scheduled on node
        |
        v
status.hostIP
        |
        v
NODE_IP
        |
        v
OTEL_EXPORTER_OTLP_ENDPOINT
        |
        v
node agent hostPort 4317
```

The active telemetry path is now:

```text
application Pod
      |
      | OTLP/gRPC to node IP:4317
      v
agent Collector Pod
      |
      v
debug exporter
```

### Expected output and behavior

The application logs the resolved endpoint when it starts:

```text
msg="configuring OTLP trace exporter" endpoint=http://192.168.49.2:4317
```

After an HTTP request, the agent Collector's logs contain:

```text
Name: GET /
Kind: Server
service.name: Str(chapter9-greeting)
```

This proves that the span crossed Pod boundaries. It was created in the
application Pod and printed by the separate agent Pod.

### Important distinctions

- `status.hostIP` is the node's address; it is not the application Pod's IP.
- The Downward API supplies Kubernetes metadata. It is not an OpenTelemetry
  feature.
- `OTEL_EXPORTER_OTLP_ENDPOINT` is an OpenTelemetry standard environment
  variable interpreted by the Go OTLP exporter.
- The `http://` scheme selects an insecure, non-TLS connection for this local
  exercise. Production transport security requirements may differ.
- The application talks directly to the agent. This is not
  Collector-to-Collector forwarding.
- Several application Pods on one node can send to the same agent, while a
  sidecar deployment has one Collector container per application Pod.

### Deliberately not introduced yet

- A Kubernetes Service for Collector discovery
- System-level host metrics
- Kubernetes metadata enrichment
- An agent-to-gateway connection
- A telemetry backend

## Step 6: Add Kubernetes resource metadata

### What changed

We added the Kubernetes attributes processor to the agent's trace pipeline:

```yaml
processors:
  k8s_attributes:
    auth_type: serviceAccount
    filter:
      node_from_env_var: KUBE_NODE_NAME
    extract:
      metadata:
        - k8s.namespace.name
        - k8s.pod.name
        - k8s.pod.uid
        - k8s.deployment.name
        - k8s.node.name
    pod_association:
      - sources:
          - from: resource_attribute
            name: k8s.pod.ip
      - sources:
          - from: connection
```

The agent DaemonSet receives its own node name through the Downward API:

```yaml
- name: KUBE_NODE_NAME
  valueFrom:
    fieldRef:
      fieldPath: spec.nodeName
```

We also added a dedicated ServiceAccount, ClusterRole, and ClusterRoleBinding.
They allow the processor to read, list, and watch Pods and namespaces without
granting permission to modify them.

The first local test showed that Minikube's `hostPort` path rewrites the
connection source address. We therefore used the Downward API to supply the
application Pod IP as a standard OpenTelemetry resource hint:

```yaml
- name: POD_IP
  valueFrom:
    fieldRef:
      fieldPath: status.podIP
- name: OTEL_RESOURCE_ATTRIBUTES
  value: "k8s.pod.ip=$(POD_IP)"
```

The application SDK reads this through `resource.WithFromEnv()`.

### Concept demonstrated

The processor maintains a cache of Kubernetes Pods and their IP addresses.
When telemetry arrives, it first tries the supplied `k8s.pod.ip` resource hint.
It retains connection-source association as a fallback:

```text
incoming k8s.pod.ip hint
          |
          v
cached Kubernetes Pod IP
          |
          v
matching Pod metadata
          |
          v
OpenTelemetry resource attributes
```

The node filter ensures that each DaemonSet agent watches only Pods running on
its own node. This matters when the same configuration runs on every node of a
large cluster.

### Expected output and behavior

The `GET /` span still has `service.name=chapter9-greeting`. The agent's debug
output now also shows resource attributes similar to:

```text
k8s.namespace.name: Str(default)
k8s.pod.name: Str(chapter9-greeting-...)
k8s.pod.uid: Str(...)
k8s.deployment.name: Str(chapter9-greeting)
k8s.node.name: Str(minikube)
```

The application sent only `k8s.pod.ip` as an association hint. The agent
obtained the Pod name, UID, namespace, Deployment, and node from the Kubernetes
API after matching that hint.

### Important distinctions

- These are resource attributes because they describe the workload that
  produced the telemetry, not one individual HTTP operation.
- The Downward API gives the agent its own node name for discovery filtering.
- Pod association identifies the application that sent the telemetry. This
  setup prefers the `k8s.pod.ip` resource hint and retains connection-based
  association as a fallback.
- Source-IP preservation depends on the Kubernetes network path. Minikube's
  local `hostPort` path rewrites it, which is why the explicit Pod IP hint is
  needed here.
- The Kubernetes attributes processor is an OpenTelemetry Collector component.
- ServiceAccounts, ClusterRoles, and ClusterRoleBindings are Kubernetes access
  control objects.
- RBAC grants permission to inspect Kubernetes metadata; it does not add
  attributes by itself.
- The processor uses the current component name `k8s_attributes`. Older
  examples commonly use its deprecated alias, `k8sattributes`.
- Connection association must run before processors that remove connection
  context, including batch and tail sampling.

### Deliberately not introduced yet

- Extracting arbitrary Kubernetes labels or annotations
- System-level host metrics
- An agent-to-gateway connection
- Gateway scaling
- A telemetry backend

## Step 7A: Deploy a gateway Collector

### What changed

We created a separate gateway deployment method:

```text
k8s/gateway/
|-- configmap.yaml
|-- deployment.yaml
`-- service.yaml
```

The ConfigMap contains another minimal Collector trace pipeline:

```text
OTLP/gRPC receiver -> traces pipeline -> debug exporter
```

The Deployment starts one gateway Collector Pod. The Kubernetes Service selects
that Pod and provides the stable in-cluster destination:

```text
chapter9-gateway:4317
```

### Concept demonstrated

A gateway is a centralized Collector service whose replicas are managed
independently of applications and nodes:

```text
Sidecar
    Scales with application Pods

Agent DaemonSet
    Scales with eligible nodes

Gateway Deployment
    Scales according to gateway capacity requirements
```

The Kubernetes Service gives senders a stable name even though gateway Pod
names and IP addresses may change:

```text
chapter9-gateway:4317
          |
          v
Service selector
          |
          v
gateway Pod endpoint
```

### Expected output and behavior

The gateway Deployment has one ready replica, and its Service has one selected
endpoint. The Collector startup log shows its OTLP/gRPC receiver listening on
port `4317`.

The gateway debug exporter does not print application spans yet. The current
telemetry path remains:

```text
application -> node agent -> agent debug exporter
```

The agent-to-gateway connection is deliberately a separate change.

### Important distinctions

- The gateway role comes from the Collector's centralized Deployment and
  routing position; it is not a different Collector binary.
- A Deployment maintains the desired gateway replica count.
- A Kubernetes Service provides stable discovery and distributes connections
  across matching gateway Pods.
- A Kubernetes Service is not the same as the Collector configuration's
  `service` section.
- A ClusterIP Service is reachable inside the cluster. It does not expose OTLP
  publicly.
- Declaring an OTLP receiver in the gateway does not cause agents to discover
  or use it automatically.

### Deliberately not introduced yet

- Agent-to-gateway OTLP forwarding
- Multiple gateway replicas
- A Horizontal Pod Autoscaler
- Stateful gateway processors such as tail sampling
- A telemetry backend

## Step 7B: Forward agent telemetry to the gateway

### What changed

We replaced the agent's debug exporter with an OTLP exporter:

```yaml
exporters:
  otlp_grpc/gateway:
    endpoint: chapter9-gateway:4317
    tls:
      insecure: true
```

The agent trace pipeline now ends at that exporter:

```yaml
service:
  pipelines:
    traces:
      receivers: [otlp]
      processors: [k8s_attributes]
      exporters: [otlp_grpc/gateway]
```

The gateway configuration did not need to change because its OTLP receiver and
debug exporter were already active.

### Concept demonstrated

Collectors can send telemetry to other Collectors through OTLP:

```text
application
    |
    v
agent Collector
    | enrich Kubernetes resource metadata
    |
    | OTLP/gRPC
    v
chapter9-gateway Kubernetes Service
    |
    v
gateway Collector
    |
    v
debug exporter
```

The Kubernetes Service decouples the agent from individual gateway Pods. The
agent uses `chapter9-gateway:4317`, while Kubernetes maintains the current set
of matching Pod endpoints.

### Expected output and behavior

After an application request:

- The agent receives and enriches the span.
- The agent does not print the span because it no longer has a debug exporter
  in its active pipeline.
- The gateway receives the enriched span and prints it.
- Kubernetes resource attributes added by the agent remain present at the
  gateway.

This proves that processors run before exporters and that OTLP preserves the
enriched telemetry between Collectors.

### Important distinctions

- Agent-to-gateway communication is Collector-to-Collector OTLP forwarding.
- The gateway Service provides discovery; it does not process telemetry.
- The gateway Collector Pod processes telemetry after the Service routes the
  network connection to it.
- `otlp_grpc/gateway` is a named instance of the OTLP/gRPC exporter. The
  `/gateway` suffix distinguishes this configured instance; it is not a
  different exporter type.
- Collector `0.153.0` calls this exporter `otlp_grpc`. Its older `otlp` alias
  still works but produces a deprecation warning.
- `tls.insecure: true` is acceptable only for this local study environment.
- The gateway is still using a debug exporter rather than a real backend.

### Deliberately not introduced yet

- Multiple gateway replicas
- A Horizontal Pod Autoscaler
- Load-balancing considerations for stateful processors
- A real telemetry backend

## Step 8: Scale the gateway

### What changed

We increased the gateway Deployment's desired replica count:

```yaml
spec:
  replicas: 2
```

The gateway Service configuration did not change. Its label selector
automatically includes both gateway Pods as endpoints.

### Concept demonstrated

The gateway scales independently from application replicas and node count:

```text
Application Deployment replicas: 1
Eligible Kubernetes nodes:       1
Agent DaemonSet Pods:             1
Gateway Deployment replicas:      2
```

Kubernetes maintains one stable Service address while its endpoint set changes:

```text
chapter9-gateway:4317
          |
          |-- gateway Pod endpoint 1
          `-- gateway Pod endpoint 2
```

### Expected output and behavior

The gateway Deployment reports two ready replicas, and its EndpointSlice
contains two Pod IP addresses.

Sending several spans through the existing agent connection demonstrates an
important gRPC behavior: the agent typically keeps its long-lived connection
to one selected gateway Pod. The individual spans are therefore not expected
to alternate between gateway replicas.

This means:

```text
Service endpoint availability
    Kubernetes can route new connections to either healthy Pod.

Per-span distribution
    A persistent gRPC connection continues using its selected Pod.
```

### Observed local result

The gateway EndpointSlice contained two ready Pod IP addresses. We then sent
five requests without restarting the agent:

```text
requests sent:                  5
spans printed by gateway Pod 1: 5
spans printed by gateway Pod 2: 0
```

The agent's existing OTLP/gRPC connection remained attached to the original
gateway Pod. This confirms that the Service made both replicas available for
connections, but did not redistribute individual spans from an established
connection.

### Important distinctions

- Horizontal scaling creates additional capacity; it does not guarantee equal
  traffic distribution at every moment.
- The Kubernetes Service balances network connections, not individual spans.
- Multiple agents and reconnecting clients can be distributed across gateway
  replicas even when one agent remains connected to one replica.
- Stateless gateway pipelines can generally process telemetry on any replica.
- Stateful processors may require related telemetry to reach the same replica.
  Tail sampling is an important example because it needs all spans from a trace
  before making a decision.
- This step uses a fixed replica count. It does not automatically respond to
  CPU, memory, queue length, or traffic.

### Deliberately not introduced yet

- A Horizontal Pod Autoscaler
- Load generation for automatic scaling
- Trace-aware load balancing
- Tail sampling configuration
- A real telemetry backend

## Step 9: Render and install the official Helm chart

### What changed

We removed the three live resources previously created from
`k8s/gateway/`:

```text
Deployment/chapter9-gateway-collector
ConfigMap/chapter9-gateway-collector
Service/chapter9-gateway
```

The YAML files themselves remain in the repository as the hand-written
reference. We then added a chart values file:

```text
k8s/helm/deployment-values.yaml
```

Its central settings are:

```yaml
mode: deployment
fullnameOverride: chapter9-gateway
replicaCount: 2
```

We rendered and installed the official Collector chart at the pinned chart
version `0.165.0`, whose Collector app version is `0.156.0`.

The rendered output is saved under:

```text
k8s/helm/rendered/opentelemetry-collector/templates/
```

### Concept demonstrated

Helm values and chart templates generate normal Kubernetes resources:

```text
deployment-values.yaml + chart templates
                    |
                    | helm template
                    v
 ConfigMap + Deployment + Service + ServiceAccount
                    |
                    | helm install
                    v
          Kubernetes API resources
```

`helm template` performs only the rendering step. It lets us inspect the result
without changing the cluster. `helm install` submits the rendered release and
stores release state so Helm can later upgrade or uninstall it.

The chart does not introduce a fourth Collector placement pattern. The selected
`mode: deployment` still creates a normal Deployment-based gateway.

### Small relevant code snippet

The values file removes unwanted default chart components with `null` and keeps
only our OTLP/gRPC trace pipeline:

```yaml
config:
  receivers:
    jaeger: null
    prometheus: null
    zipkin: null
    otlp:
      protocols:
        grpc:
          endpoint: ${env:MY_POD_IP}:4317
        http: null

  service:
    pipelines:
      logs: null
      metrics: null
      traces:
        receivers: [otlp]
        processors: [memory_limiter, batch]
        exporters: [debug]
```

The chart's values are merged with its defaults. In a top-level installation,
setting a default component or pipeline to `null` removes it from that merged
configuration.

### Expected output and behavior

The installed release should report:

```text
release:             chapter9-gateway
status:              deployed
chart version:       0.165.0
Collector version:   0.156.0
ready gateway Pods:  2
Service endpoints:   2
```

The existing agent still exports to:

```text
chapter9-gateway:4317
```

Because `fullnameOverride` preserves that Service name, the agent does not need
configuration changes. After the new Service becomes ready, the existing gRPC
client resolves its new ClusterIP and reconnects. An application request again
appears in the Helm-managed gateway's detailed debug output with the Kubernetes
resource attributes previously added by the agent.

### Important distinctions

- `deployment-values.yaml` is input; files under `rendered/` are generated
  output.
- `helm template` renders locally and does not install anything.
- `helm install` creates a release and marks its resources as managed by Helm.
- The Helm release name, Kubernetes resource names, and Collector service name
  are related but distinct concepts.
- `fullnameOverride` controls generated Kubernetes names; it does not change
  the telemetry resource attribute `service.name`.
- The chart generated a ServiceAccount even though this gateway does not query
  the Kubernetes API. Creating an identity is not the same as granting it
  additional ClusterRole permissions.
- The chart adds health probes, a configuration checksum annotation,
  downward-API environment variables, and `GOMEMLIMIT` wiring that our raw
  gateway omitted.
- The chart's health-check extension supports the generated liveness and
  readiness probes; it is not part of the trace pipeline.
- The `memory_limiter` receives a percentage of the configured memory limit,
  while Kubernetes independently enforces the container memory limit.
- Removing and recreating the Service changed its ClusterIP. Its stable DNS
  name allowed the agent to discover and reconnect to the replacement.

### Observed local result

Helm installed revision 1 successfully. Kubernetes reported two ready gateway
Pods and two endpoints for `chapter9-gateway`. A request returned:

```text
Hello, Helm!
```

One Helm-managed gateway Pod printed the `GET /` span, including
`service.name=chapter9-greeting` and the Pod, namespace, Deployment, and node
resource attributes added by the agent.

### Deliberately not introduced yet

- Rendering the chart in `daemonset` mode
- Collector chart presets such as host metrics or Kubernetes attributes
- Chart-generated ClusterRoles and host volume mounts
- Helm upgrades or rollbacks
- The OpenTelemetry Operator
- A real telemetry backend

## Step 10: Understand the OpenTelemetry Operator

### What changed

We added a readable, non-applied conceptual example:

```text
k8s/operator/conceptual_example.md
```

It compares Helm with the Operator and shows the two important custom resource
types without installing them:

```text
OpenTelemetryCollector
Instrumentation
```

The running Minikube resources did not change. The gateway remains managed by
the Helm release from Step 9.

### Concept demonstrated

The Operator is a long-running Kubernetes controller:

```text
OpenTelemetry custom resource
             |
             | watch
             v
OpenTelemetry Operator reconciliation loop
             |
             | create/update/delete managed children
             v
Deployment + ConfigMap + Service + Pods
```

The Operator extends the Kubernetes API with CRDs. A Custom Resource (CR) is an
instance of one of those new API types and expresses desired state. The
Operator watches the CR and reconciles the actual child resources toward that
desired state.

Helm also creates Kubernetes resources, but the Helm command is not a resident
controller continuously watching its values file. After Helm creates a
Deployment, Kubernetes' built-in Deployment controller keeps the requested Pods
running; Helm only recalculates the chart when we explicitly run an operation
such as `helm upgrade`.

### Small relevant code snippet

An Operator-managed Collector begins with an OpenTelemetry-specific resource
rather than a direct Deployment:

```yaml
apiVersion: opentelemetry.io/v1beta1
kind: OpenTelemetryCollector
metadata:
  name: chapter9-gateway
spec:
  mode: deployment
  replicas: 2
  config:
    receivers:
      otlp:
        protocols:
          grpc:
            endpoint: 0.0.0.0:4317
    exporters:
      debug: {}
    service:
      pipelines:
        traces:
          receivers: [otlp]
          exporters: [debug]
```

The Collector configuration remains familiar. The important difference is that
the Operator reads this CR and produces the underlying Kubernetes resources.

### Expected output and behavior

There is deliberately no new runtime output in this conceptual step. If the
Operator and its CRDs were installed later, applying the example
`OpenTelemetryCollector` would cause the controller to create Collector
resources. Commands such as these would then expose both layers:

```text
kubectl get opentelemetrycollectors
kubectl get deployments
```

The first would show the desired OpenTelemetry resource; the second would show
one of the Kubernetes resources reconciled from it.

### Important distinctions

- A CRD defines a new Kubernetes API type; a CR is one object of that type.
- An Operator is the controller that gives those custom resources behavior.
- `OpenTelemetryCollector` manages Collector infrastructure.
- `Instrumentation` describes SDK and auto-instrumentation settings.
- Deploying a Collector does not automatically instrument applications.
- Auto-instrumentation also requires workload opt-in, normally through
  language-specific Pod annotations processed by an admission webhook.
- Helm can install the Operator itself. Helm would manage the Operator
  installation, while the running Operator would manage its OpenTelemetry CRs.
- The Collector pipeline is still made of receivers, processors, and exporters.
- Operator reconciliation does not eliminate the need to choose deployment
  placement, pipeline configuration, resources, or security settings.

### Deliberately not introduced yet

- Installing the Operator or its CRDs
- Admission webhooks and their certificates
- Applying an `OpenTelemetryCollector` CR
- Applying an `Instrumentation` CR
- Mutating or restarting an application workload
- Go auto-instrumentation and its additional security requirements
- The Target Allocator
- A real telemetry backend

## Step 10A: Replace the Collector chart with an Operator-managed gateway

### What changed

We uninstalled the direct Collector Helm release:

```bash
helm uninstall chapter9-gateway
```

The values and rendered chart files from Step 9 remain saved. We then installed
the Operator chart `0.120.0` into its own namespace:

```text
Helm release:    chapter9-operator
Namespace:       opentelemetry-operator-system
Operator:        0.156.0
```

Minikube does not have cert-manager, so `operator-values.yaml` selects the
chart's Helm-generated self-signed certificate for the admission webhook.

Finally, we applied:

```text
k8s/operator/gateway-collector.yaml
```

This created `OpenTelemetryCollector/chapter9-gateway` with
`mode: deployment` and two replicas.

### Concept demonstrated

The Operator itself and the Collector it manages are separate deployments with
different responsibilities:

```text
Operator Deployment
    |
    | watches and reconciles
    v
OpenTelemetryCollector/chapter9-gateway
    |
    | desired mode: deployment
    v
Gateway Collector Deployment with two Pods
```

“Gateway” describes the Collector's role in the telemetry architecture.
`mode: deployment` describes its Kubernetes placement. A gateway can run as a
Deployment, but `gateway` is not an Operator deployment-mode value.

### Small relevant code snippet

The custom resource combines Kubernetes placement with ordinary Collector
configuration:

```yaml
apiVersion: opentelemetry.io/v1beta1
kind: OpenTelemetryCollector
metadata:
  name: chapter9-gateway
spec:
  mode: deployment
  replicas: 2
  config:
    receivers:
      otlp:
        protocols:
          grpc:
            endpoint: 0.0.0.0:4317
    processors:
      memory_limiter:
        check_interval: 5s
        limit_percentage: 80
        spike_limit_percentage: 25
      batch: {}
    exporters:
      debug:
        verbosity: detailed
    service:
      pipelines:
        traces:
          receivers: [otlp]
          processors: [memory_limiter, batch]
          exporters: [debug]
```

The literal receiver endpoint also lets the Operator discover port `4317` and
generate Services for it.

### Expected output and behavior

The parent custom resource reports its mode, Collector version, and readiness:

```text
NAME               MODE         VERSION   READY
chapter9-gateway   deployment   0.156.0   2/2
```

The Operator generated these child resources:

```text
Deployment  chapter9-gateway-collector
ConfigMap   chapter9-gateway-collector-<configuration-hash>
Service     chapter9-gateway-collector
Service     chapter9-gateway-collector-headless
Service     chapter9-gateway-collector-monitoring
ServiceAccount chapter9-gateway-collector
```

The agent endpoint changed to the Operator-generated OTLP Service:

```yaml
endpoint: chapter9-gateway-collector:4317
```

### Helm-chart versus Operator result

The direct Collector chart from Step 9 generated resources during a Helm
operation and labeled them as managed by Helm:

```text
values -> Helm -> Deployment + ConfigMap + Service + ServiceAccount
```

The Operator setup has an additional persistent management layer:

```text
Operator chart -> Operator Deployment + CRDs + webhooks
OpenTelemetryCollector CR -> Operator -> managed Collector resources
```

The generated resources are still normal Kubernetes objects, but their
`ownerReferences` point to `OpenTelemetryCollector/chapter9-gateway` and their
`app.kubernetes.io/managed-by` label identifies
`opentelemetry-operator`.

The Operator additionally generated a headless Service, a monitoring Service,
and a content-hashed ConfigMap for this Collector. Those details are
implementation conveniences, not changes to the trace pipeline.

### Observed local result

The Operator CR and generated Deployment both reported `2/2` ready replicas. A
request returned:

```text
Hello, Operator!
```

One Operator-managed gateway Pod printed the `GET /` span with
`service.name=chapter9-greeting` and the Kubernetes metadata added by the
agent. This verified:

```text
application -> agent -> Operator-managed gateway -> debug exporter
```

We also manually changed the generated Deployment to one replica:

```bash
kubectl scale deployment/chapter9-gateway-collector --replicas=1
```

The `OpenTelemetryCollector` CR still declared two replicas. The Operator
immediately reconciled the generated Deployment back to two, and the CR again
reported `2/2` ready. This is the clearest observed difference from a direct
Collector chart installation: changing a Helm-generated Deployment does not
cause a resident Helm controller to reread `values.yaml` and correct it.

### Important distinctions

- Installing the Operator with Helm does not mean Helm directly manages each
  Collector created afterward.
- Helm owns the Operator release; the Operator owns the Collector child
  resources.
- The `OpenTelemetryCollector` CR is desired state, not the Collector Pod.
- Deleting the CR would ask the Operator to remove its owned child resources.
- The Operator-generated hashed ConfigMap name can change when the Collector
  configuration changes.
- The regular Service provides stable client discovery; the headless Service
  exposes individual Pod addresses through DNS.
- The monitoring Service exposes the Collector's internal telemetry port; it
  does not carry the application's OTLP spans.
- Auto-instrumentation capability is installed but has not been enabled for
  the greeting application.

### Deliberately not introduced yet

- An `Instrumentation` custom resource
- Application admission-webhook injection
- Go auto-instrumentation
- The Target Allocator
- A backend other than the debug exporter

## Step 11: Create the final Chapter 9 concept map

### What changed

We added:

```text
ai_final_chapter09_concept_map.md
```

It consolidates the placement patterns, running telemetry path, metadata
enrichment, discovery, scaling behavior, management methods, Operator
ownership, and the relationship to a future Grafana Alloy deployment.

No application or Kubernetes resource changed in this documentation-only step.

### Concept demonstrated

Collector placement and Collector management are independent design axes:

```text
Placement:
    sidecar | agent | gateway

Management:
    raw manifests | Helm chart | OpenTelemetry Operator
```

A gateway can be managed with raw YAML, Helm, or the Operator. Likewise, using
the Operator does not automatically determine whether a Collector should be a
sidecar, agent, or gateway.

### Small relevant snippet

The final running path is:

```text
application
    -> node agent
    -> Kubernetes metadata enrichment
    -> Operator-managed gateway
    -> debug exporter
```

The final gateway is architecturally a gateway and is placed by the Operator
using:

```yaml
spec:
  mode: deployment
  replicas: 2
```

### Expected output and behavior

The concept map is documentation rather than executable configuration. It
should let a reader answer:

- Which placement matches Pod-local, node-local, or cluster-level collection?
- Why does an agent forward through OTLP to a gateway?
- Where are Kubernetes resource attributes added?
- Why can two gateway replicas still receive uneven span traffic?
- What changes between raw YAML, Helm, and the Operator?
- How do these lessons transfer to Grafana Alloy?

### Important distinctions

- A Collector placement describes architectural responsibility and locality.
- A Kubernetes workload kind describes process scheduling and lifecycle.
- A management method describes who creates and reconciles resources.
- A receiver, processor, or exporter describes telemetry pipeline behavior.
- Alloy can perform collection and processing roles but is not Tempo, Loki, or
  a Prometheus-compatible storage backend.

These categories interact, but none of them is a substitute for another.

### Deliberately not introduced yet

- A real telemetry backend
- Production Grafana Alloy manifests
- TLS and authentication
- Persistent queues
- Autoscaling
- Tail sampling and trace-aware routing
- Operator-managed application auto-instrumentation

## Step 11A: Compare OpenTelemetry deployment and pipelines with Alloy

### What changed

We added a focused comparison:

```text
../ai_opentelemetry_vs_alloy_deployment_and_pipelines.md
```

It covers Kubernetes placement, the OpenTelemetry pipeline model, Alloy's
component graph, and the limits and purpose of Alloy clustering.

No program or Kubernetes resource changed.

### Concept demonstrated

OpenTelemetry Collector and Alloy support similar placement topologies:

```text
Pod-local:      sidecar
Node-local:     DaemonSet
Centralized:    Deployment or StatefulSet
```

Alloy's standard Helm chart calls the workload setting `controller.type` and
supports `daemonset`, `deployment`, and `statefulset`. Sidecar is a topology
rather than a `controller.type` value.

Their processing models express similar data flow differently:

```text
OpenTelemetry Collector:
    service.pipelines orders receivers, processors, and exporters.

Alloy:
    component references connect typed outputs and inputs into a graph.
```

### Small relevant snippet

An Alloy component forwards traces by referring to the next component's input:

```alloy
otelcol.receiver.otlp "applications" {
  output {
    traces = [otelcol.processor.batch.default.input]
  }
}

otelcol.processor.batch "default" {
  output {
    traces = [otelcol.exporter.otlp.tempo.input]
  }
}
```

This graph plays the same broad role as:

```text
OTLP receiver -> batch processor -> OTLP exporter
```

### Expected output and behavior

This documentation should make it possible to choose an Alloy topology without
mistaking:

- Gateway role for a Kubernetes mode
- Kubernetes replica management for Alloy clustering
- Alloy cluster membership for per-span load balancing
- An Alloy component graph for a telemetry storage backend

### Important distinctions

- Alloy clustering is independent of Kubernetes `controller.type`.
- Only components that support clustering and explicitly enable it distribute
  their work through the Alloy cluster.
- Pull-based scrape targets are the clearest clustering use case.
- Pushed OTLP traffic still needs network load balancing.
- Persistent gRPC connections can still concentrate spans on one replica.
- Stateful trace processing can still need trace-aware routing.
- Alloy's typed graph is broader than OpenTelemetry-only processing because it
  can connect Prometheus, Loki, Pyroscope, discovery, and `otelcol` components.

### Deliberately not introduced yet

- Deploying Alloy in this Minikube exercise
- Configuring an Alloy cluster
- Tempo, Loki, or Prometheus destinations
- A side-by-side load or failure test
