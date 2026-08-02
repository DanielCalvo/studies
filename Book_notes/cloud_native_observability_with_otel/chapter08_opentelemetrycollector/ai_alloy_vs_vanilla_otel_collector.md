# Grafana Alloy versus a vanilla OpenTelemetry Collector

Grafana Alloy is not merely a renamed OpenTelemetry Collector, although
Grafana's commercial interests naturally influence which integrations and
features it prioritizes.

The first important point is that the OpenTelemetry Collector is both a
framework and a collection of components. Organizations can assemble their own
Collector **distributions**. A distribution can add or remove components,
provide different defaults and packaging, perform additional testing, or add
vendor-specific integrations.

Alloy is Grafana's expanded distribution of the Collector model. It combines
OpenTelemetry collection with functionality developed for the Prometheus and
Grafana ecosystems.

## Alloy's main advantages

### 1. It combines several observability ecosystems

A vanilla Collector is principally OpenTelemetry-oriented. Alloy brings
together several ecosystems in one process:

```text
OpenTelemetry -> traces, metrics, and logs
Prometheus    -> scraping, discovery, relabeling, and remote write
Loki          -> log collection and processing
Pyroscope     -> continuous profiles
Grafana Beyla -> eBPF-based instrumentation
```

One Alloy deployment can therefore collect:

- OTLP traces from instrumented applications
- Prometheus metrics from Kubernetes targets
- Container logs for Loki
- Profiles for Pyroscope

This can avoid separately operating an OpenTelemetry Collector, a
Prometheus-oriented agent, Promtail, and a profiling agent.

For someone already using Kubernetes, Tempo, Loki, and Grafana, this is probably
Alloy's most relevant advantage.

### 2. It has first-class Grafana-stack integration

Alloy has particularly direct integration paths for Grafana's LGTM stack:

```text
metrics  -> Mimir or Prometheus
logs     -> Loki
traces   -> Tempo
profiles -> Pyroscope
```

A vanilla Collector can also send data to these systems, often using OTLP or
components from the Collector Contrib distribution. Alloy makes the
combinations Grafana expects more prominent, integrated, and tested.

This is one place where Grafana's commercial interests are visible, but the
integration is also genuinely useful when this is the stack being operated.

### 3. It provides a programmable component graph

A conventional OpenTelemetry Collector configuration normally describes
receivers, processors, exporters, and signal pipelines in YAML:

```yaml
receivers:
processors:
exporters:
service:
  pipelines:
```

Alloy connects named components into an explicit graph:

```text
discovery -> collection -> processing -> output
```

Its configuration language supports references between components, expressions,
functions, conditional logic, and reusable modules. These features can be
helpful for complex pipelines and for platform teams that want to share standard
collection modules across services.

The corresponding trade-off is that Alloy configuration is specific to Alloy
and is less directly portable than standard Collector YAML.

### 4. It includes a troubleshooting UI

Alloy's built-in web interface can help inspect:

- The component graph
- Component health
- Component configuration
- Discovered targets
- Some component outputs and debug information
- Alloy clustering state

A vanilla Collector exposes internal telemetry and offers troubleshooting
extensions, but Alloy provides a more unified visual view of its component
pipeline.

### 5. It provides built-in clustering for supported pull workloads

Suppose three collectors discover the same 3,000 Prometheus targets. If they are
not coordinated, all three might scrape every target, duplicating collection.

Alloy clustering can distribute supported work approximately like this:

```text
3,000 discovered targets

Alloy A -> approximately 1,000
Alloy B -> approximately 1,000
Alloy C -> approximately 1,000
```

If one instance disappears, the remaining instances can redistribute its
targets. Alloy uses cluster membership and consistent hashing for this.

A vanilla Collector can also be scaled, but scaling pull-based receivers such
as the Prometheus receiver may require a separate sharding strategy or the
OpenTelemetry Operator's Target Allocator.

### 6. It carries forward Grafana Agent functionality

Alloy is the successor to Grafana Agent. Grafana had already developed
functionality involving:

- Prometheus-compatible metrics
- Loki log collection
- Kubernetes discovery
- Prometheus remote write
- Efficient telemetry forwarding
- Grafana Cloud integrations

Alloy consolidates that work around the OpenTelemetry Collector architecture
instead of continuing to maintain Grafana Agent as a separate collection
system.

## When a vanilla Collector may be preferable

A vanilla OpenTelemetry Collector is attractive when:

- Standard upstream Collector YAML and portability are priorities.
- Only a straightforward OTLP pipeline is required.
- A specific upstream component or release is needed immediately.
- A smaller and more narrowly selected distribution is desirable.
- The organization already standardizes on the OpenTelemetry Operator and
  Collector tooling.
- Alloy's Prometheus, Loki, profiling, clustering, or module functionality is
  not needed.

For example, the following pipeline does not inherently require Alloy:

```text
applications -> OTLP receiver -> batch processor -> OTLP backend
```

A vanilla Collector can handle that perfectly well.

## Does Alloy create vendor lock-in?

Not inherently at the telemetry protocol level. Alloy is open source, can send
OTLP to non-Grafana destinations, and can send Prometheus remote-write data to
compatible systems.

There can nevertheless be **soft operational lock-in**:

- Alloy's component configuration syntax is Grafana-specific.
- Some components are designed particularly around Grafana products.
- Moving to another distribution may require translating configuration and
  operational procedures.

The telemetry remains portable when open protocols and formats such as OTLP and
Prometheus remote write are used. The less portable part is mainly the
collector's configuration and the team's operational knowledge.

## Practical conclusion

Alloy exists partly because Grafana wants an excellent collection layer for its
products, but it is not merely cosmetic branding. It combines OpenTelemetry
Collector functionality with the work Grafana previously did in Grafana Agent,
as well as Prometheus, Loki, Kubernetes, profiling, clustering, modules, and
debugging capabilities.

Given an environment already using Kubernetes, Alloy, Loki, Tempo, and Grafana,
Alloy is a logical choice because it can consolidate several collection
pipelines into one system.

If the only requirement were to receive OTLP traces and forward them to a
single generic backend, Alloy's additional functionality might provide little
advantage over a vanilla Collector.

## Current references

- [OpenTelemetry: Collector distributions](https://opentelemetry.io/docs/concepts/distributions/)
- [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/)
- [Grafana: Why Alloy](https://grafana.com/docs/alloy/latest/introduction/why-alloy/)
- [Grafana: How Alloy works](https://grafana.com/docs/alloy/latest/introduction/how-alloy-works/)
- [Grafana: Alloy clustering](https://grafana.com/docs/alloy/latest/get-started/clustering/)

