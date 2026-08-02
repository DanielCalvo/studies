# OpenTelemetry Operator conceptual example

This file began as a reading-only comparison. The running version of the
gateway custom resource now lives in `gateway-collector.yaml`; the snippets in
this document remain conceptual rather than being applied directly.

## The additional layer

With the Helm chart, we provide values and explicitly ask Helm to generate and
install ordinary resources:

```text
values.yaml -> Helm command -> Deployment + ConfigMap + Service
```

With the Operator, we install a long-running Kubernetes controller and declare
OpenTelemetry-specific custom resources:

```text
OpenTelemetryCollector custom resource
                  |
                  | watched and reconciled continuously
                  v
        OpenTelemetry Operator
                  |
                  v
       Deployment + ConfigMap + Service
```

The Collector process and its pipeline do not fundamentally change. The new
piece is the controller that translates and continually reconciles the desired
state.

## Collector custom resource

The current official documentation uses `opentelemetry.io/v1beta1` for the
Collector custom resource:

```yaml
apiVersion: opentelemetry.io/v1beta1
kind: OpenTelemetryCollector
metadata:
  name: chapter9-gateway
spec:
  # This resembles the Helm chart's mode, but it belongs to the Operator CRD.
  mode: deployment
  replicas: 2

  # The pipeline is still ordinary Collector configuration.
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

Applying this custom resource would not directly create a Collector Pod.
Instead, it would ask the Operator to create and manage the necessary
Kubernetes resources.

## Instrumentation custom resource

Collector management and application auto-instrumentation are separate Operator
capabilities. An `Instrumentation` resource describes SDK-related settings:

```yaml
apiVersion: opentelemetry.io/v1alpha1
kind: Instrumentation
metadata:
  name: chapter9-instrumentation
spec:
  exporter:
    endpoint: http://chapter9-gateway-collector:4317
  propagators:
    - tracecontext
    - baggage
  sampler:
    type: parentbased_traceidratio
    argument: "1"
```

A workload must then opt in with a language-specific annotation, for example:

```yaml
metadata:
  annotations:
    instrumentation.opentelemetry.io/inject-java: "true"
```

The Operator's admission webhook uses the annotation and `Instrumentation`
resource when a Pod is created. This is different from merely deploying a
Collector: a Collector receives telemetry, while auto-instrumentation changes
the application Pod so that the application can produce and export telemetry.

Language-specific injection has different requirements and maturity. In
particular, current Go auto-instrumentation is not equivalent to the Java
example above and has additional executable-path and security requirements. We
will verify those constraints before attempting a later Go experiment.

## Helm versus Operator

| Question | Collector Helm chart | OpenTelemetry Operator |
| --- | --- | --- |
| What do we declare? | Chart values | OpenTelemetry custom resources |
| What runs afterward? | Collector resources; Helm itself is not a resident controller | Collector resources plus a resident Operator controller |
| Who translates the declaration? | Helm during install or upgrade | Operator reconciliation loop |
| Collector pipeline format | Normal Collector configuration | Normal Collector configuration inside the CR |
| Manages application auto-instrumentation? | No | Yes, through `Instrumentation` resources and admission-webhook injection |
| Requires OpenTelemetry CRDs? | No | Yes |

Installing the Operator itself can be done with Helm. That does not make Helm
and the Operator the same thing: Helm installs the controller and CRDs; the
running controller then manages OpenTelemetry custom resources.
