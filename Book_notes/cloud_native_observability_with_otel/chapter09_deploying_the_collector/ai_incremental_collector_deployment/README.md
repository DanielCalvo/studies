# Incremental Collector deployment

This exercise studies the Kubernetes deployment patterns from Chapter 9 one
small change at a time.

The current implementation is **Step 10A**: the OpenTelemetry Operator manages
the two-replica gateway Collector. The direct Collector Helm release from Step
9 has been uninstalled, while its values and rendered resources remain saved
for comparison.

## Current architecture

```text
Application Deployment        Agent DaemonSet
        |                            |
        | OTLP to node IP             |
        `---------------------------->|
                                     v
                         Kubernetes attributes
                                processor
                                     |
                                     v
                              OTLP exporter ------> Operator-managed gateway
                                                   debug exporter
```

The application supplies its Pod IP as an association hint. The agent matches
that hint to a Pod in its Kubernetes API cache, then adds Pod, namespace,
Deployment, and node identity to the span's resource.

The Operator-managed gateway is independently scalable and reachable inside
the cluster at:

```text
chapter9-gateway-collector:4317
```

## Reproduce Step 9: Collector Helm chart

This section preserves the previous Helm-based gateway experiment. It is not
the currently running scenario. To recreate it fully, change the agent exporter
back to `chapter9-gateway:4317`; the active agent configuration now targets the
Operator-generated `chapter9-gateway-collector:4317` Service.

Start Minikube if it is not already running:

```bash
minikube start
```

Build the application image directly into Minikube's image store if it is not
already present:

```bash
minikube image build \
  -t chapter9-greeting:step6 \
  ./application
```

Apply the application-only Deployment and agent DaemonSet:

```bash
# Create the ServiceAccount and its read-only permissions first.
kubectl apply -f k8s/agent/rbac.yaml

kubectl apply \
  -f k8s/agent/configmap.yaml \
  -f k8s/agent/daemonset.yaml \
  -f k8s/agent/deployment.yaml

kubectl rollout status deployment/chapter9-greeting
kubectl rollout status daemonset/chapter9-agent-collector
```

Add the official chart repository:

```bash
helm repo add \
  open-telemetry \
  https://open-telemetry.github.io/opentelemetry-helm-charts
helm repo update open-telemetry
```

Render the pinned chart locally before installing it:

```bash
helm template \
  chapter9-gateway \
  open-telemetry/opentelemetry-collector \
  --version 0.165.0 \
  --namespace default \
  --values k8s/helm/deployment-values.yaml \
  --output-dir k8s/helm/rendered
```

This command does not contact the Kubernetes API or create runtime resources.
It turns the chart templates and our values into ordinary YAML files:

```text
k8s/helm/rendered/opentelemetry-collector/templates/
|-- configmap.yaml
|-- deployment.yaml
|-- service.yaml
`-- serviceaccount.yaml
```

Install that same pinned chart and values file:

```bash
helm upgrade --install \
  chapter9-gateway \
  open-telemetry/opentelemetry-collector \
  --version 0.165.0 \
  --namespace default \
  --values k8s/helm/deployment-values.yaml

kubectl rollout status deployment/chapter9-gateway
```

Verify the Helm release, both gateway Pods, and Service endpoints:

```bash
helm list --filter '^chapter9-gateway$'
kubectl get deployment chapter9-gateway
kubectl get pods -l app.kubernetes.io/instance=chapter9-gateway -o wide
kubectl get service chapter9-gateway
kubectl get endpointslice -l kubernetes.io/service-name=chapter9-gateway
```

Expected result:

```text
Helm release status:                deployed
Gateway Deployment ready replicas: 2
Gateway Service endpoints:         2
```

After applying a changed agent ConfigMap, restart the DaemonSet because the
configuration file is mounted through `subPath`:

```bash
kubectl rollout restart daemonset/chapter9-agent-collector
kubectl rollout status daemonset/chapter9-agent-collector
```

Compare their placement and container counts:

```bash
kubectl get pods -l app=chapter9-greeting
kubectl get daemonset chapter9-agent-collector
kubectl get pods -l app=chapter9-agent-collector -o wide
```

Temporarily forward local port `8080` to the application:

```bash
kubectl port-forward deployment/chapter9-greeting 8080:8080
```

In another terminal, create one traced request:

```bash
curl "http://localhost:8080/?name=Agent"
```

The application responds normally:

```text
Hello, Agent!
```

After the application's batch span processor exports the completed span,
inspect both Collectors:

```bash
kubectl logs daemonset/chapter9-agent-collector -c collector
kubectl logs deployment/chapter9-gateway -c opentelemetry-collector
```

The restarted agent should show its receiver and processor starting but should
not print the span. The gateway debug output should contain the `GET /` server
span with `service.name=chapter9-greeting` and resource attributes resembling:

```text
k8s.namespace.name: default
k8s.pod.name: chapter9-greeting-...
k8s.pod.uid: ...
k8s.deployment.name: chapter9-greeting
k8s.node.name: minikube
```

The application startup log also shows the expanded endpoint:

```bash
kubectl logs deployment/chapter9-greeting -c greeting
```

```text
msg="configuring OTLP trace exporter" endpoint=http://192.168.49.2:4317
```

## Archived sidecar scenario

The complete Step 3 sidecar manifests remain available:

```text
k8s/sidecar/configmap.yaml
k8s/sidecar/deployment.yaml
```

They still describe:

```text
Pod
|-- greeting application
`-- Collector sidecar
```

Do not apply the sidecar and application-only Deployment simultaneously because
both manifests intentionally describe a Deployment named `chapter9-greeting`.
They are two alternative definitions of the same workload.

To return to the sidecar version:

```bash
kubectl apply \
  -f k8s/sidecar/configmap.yaml \
  -f k8s/sidecar/deployment.yaml
```

Both scenarios use the same application image. The sidecar manifest sets:

```text
OTEL_EXPORTER_OTLP_ENDPOINT=http://localhost:4317
```

whereas the agent manifest constructs the endpoint from the Pod's node IP.

## Why the agent has Kubernetes API permissions

The agent uses a dedicated ServiceAccount and read-only RBAC from
`k8s/agent/rbac.yaml`. The processor must list and watch Pods to maintain the
IP-to-Pod cache used for association.

The pipeline order is:

```text
OTLP receiver -> Kubernetes attributes processor -> OTLP gateway exporter
```

Connection-based association must run before processors such as batch or tail
sampling that can discard the receiver's connection context.

The processor attempts association in this order:

```text
k8s.pod.ip resource attribute -> connection source address
```

The explicit hint is necessary in this local Minikube setup because its
`hostPort` network path rewrites the connection source address. Environments
that preserve the source Pod IP can use connection association alone.

## Hand-written gateway reference

The raw resources from Steps 7 and 8 remain available for comparison:

```text
k8s/gateway/
|-- configmap.yaml
|-- deployment.yaml
`-- service.yaml
```

They are not currently applied. The Operator now owns equivalent live resources
derived from the `chapter9-gateway` custom resource. Avoid applying the raw
gateway simultaneously because it would create another gateway and make the
comparison ambiguous.

Unlike an agent DaemonSet, a gateway Deployment does not create one Pod per
node. Unlike a sidecar, it does not belong to an application Pod. It can be
scaled independently behind the stable `chapter9-gateway` Service.

The Kubernetes Service and the Collector's `service.pipelines` configuration
are unrelated concepts:

- The Kubernetes Service provides network discovery and routing to Pods.
- `service.pipelines` activates Collector receivers, processors, and exporters.

## Active Collector chain

The complete running trace path is now:

```text
application
    |
    | OTLP to node IP:4317
    v
agent
    |
    | Kubernetes metadata enrichment
    | OTLP to chapter9-gateway-collector:4317
    v
Operator-managed gateway
    |
    v
debug exporter
```

The application knows only its node agent. The agent knows only the stable
gateway Service name. Neither component needs to know a gateway Pod name or IP.

## Observe the long-lived OTLP/gRPC connection

Send several application requests without restarting the agent, then inspect
each gateway Pod separately:

```bash
kubectl get pods -l app.kubernetes.io/instance=default.chapter9-gateway
kubectl logs <first-gateway-pod> -c otc-container
kubectl logs <second-gateway-pod> -c otc-container
```

The Service has two endpoints, but the agent's existing gRPC connection will
normally continue sending spans to the gateway Pod it already selected:

```text
two healthy Service endpoints
              does not imply
one connection alternates between both endpoints
```

Scaling gives Kubernetes more gateway destinations for multiple connections
and senders. It does not redistribute every telemetry item independently.

## Step 10A: Operator-managed gateway

The Operator configuration and readable comparison are in:

```text
k8s/operator/operator-values.yaml
k8s/operator/gateway-collector.yaml
k8s/operator/conceptual_example.md
```

The central distinction is:

```text
Helm
    Render and install resources when we run a Helm command.

Operator
    Run a Kubernetes controller that watches OpenTelemetry custom resources
    and continuously reconciles their underlying resources.
```

The Operator is installed by Helm because Helm is still useful for installing
the controller itself:

```bash
helm upgrade --install \
  chapter9-operator \
  open-telemetry/opentelemetry-operator \
  --version 0.120.0 \
  --namespace opentelemetry-operator-system \
  --create-namespace \
  --values k8s/operator/operator-values.yaml
```

The gateway is not a second Helm release. It is an
`OpenTelemetryCollector` custom resource:

```bash
kubectl apply -f k8s/operator/gateway-collector.yaml
kubectl get opentelemetrycollector chapter9-gateway
kubectl rollout status deployment/chapter9-gateway-collector
```

The current relationship is:

```text
Helm release chapter9-operator
    |
    `-- Operator Deployment and OpenTelemetry CRDs
            |
            `-- watches OpenTelemetryCollector/chapter9-gateway
                    |
                    `-- gateway Deployment, ConfigMap, Services,
                        ServiceAccount, and Pods
```

The `Instrumentation` CRD and admission webhook are installed as part of the
Operator, but no application auto-instrumentation resource or annotation has
been applied.

## Final Chapter 9 concept map

The completed placement, forwarding, scaling, management, and future Alloy
relationships are summarized in:

```text
ai_final_chapter09_concept_map.md
```

This is the concise reference for the chapter. The chronological implementation
details and observed outputs remain in `ai_incremental_steps_taken.md`.
