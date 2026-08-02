# Why use a Collector before Tempo?

Tempo exposes OTLP endpoints, so an instrumented application can send traces
directly to it:

```text
Application -> Tempo
```

An OpenTelemetry Collector or Grafana Alloy is therefore not a strict
requirement. It becomes valuable because it provides a shared processing and
delivery layer between applications and the tracing backend:

```text
Applications -> Collector or Alloy -> Tempo
```

Application SDKs can already provide features such as batching and limited
retries. The advantage of the Collector is that these policies can be managed
consistently outside every individual application.

## Batching

The Collector can combine many small OTLP exports into fewer, larger requests
before sending them to Tempo. This reduces connection and request overhead and
usually makes the telemetry pipeline more efficient. Applications may also
batch spans inside their SDKs; Collector-side batching is an additional batch
at the shared pipeline boundary.

## Retries

If Tempo temporarily refuses a request or becomes unreachable, the Collector's
exporter can retry delivery with configurable delays and backoff. This keeps
retry policy out of each application and avoids requiring every service team to
configure the same backend-specific behavior. Retries are finite and cannot
guarantee delivery through an indefinitely unavailable backend.

## Queuing and buffering

The Collector can queue telemetry before exporting it, absorbing short
differences between the application's production rate and Tempo's ingestion
capacity. A sending queue can normally use memory and can optionally be paired
with persistent storage, depending on the Collector distribution and
configuration. Queues are bounded: when they fill, backpressure or dropped data
must still be handled.

## Filtering

The Collector can discard spans that are not useful or that are too expensive
to retain. For example, it might remove routine health-check traces. Filtering
reduces network and storage volume, but dropping individual spans carelessly
can produce incomplete traces, so filter policy needs to be intentional.

## Attribute enrichment

The Collector can add or normalize resource and span attributes before traces
reach Tempo. A Kubernetes Collector might attach cluster, namespace, Pod, node,
region, or environment information. Central enrichment provides consistent
metadata even when applications are written in different languages or use
different instrumentation libraries.

## Tail sampling

Tail sampling waits until enough of a trace has arrived to make a decision
using the completed trace's behavior. It can retain errors and unusually slow
traces while sampling ordinary successful traces more aggressively. This is
more informed than deciding at the beginning of a trace, but it is stateful,
uses memory, and requires spans from the same trace to be routed to the same
sampling decision point.

## Sensitive-data removal

The Collector can delete, hash, redact, or transform sensitive attributes
before telemetry reaches its long-term backend. This provides a centralized
privacy boundary for credentials, tokens, personal data, or prohibited request
content. Preventing sensitive data from being instrumented in the first place
is still preferable; Collector-side scrubbing is an additional safeguard.

## Routing and export fan-out

The Collector can route different telemetry to different destinations based on
signal, environment, tenant, service, or policy. It can also export the same
telemetry to more than one destination during a migration or investigation.
Applications then send to one OTLP endpoint instead of knowing the addresses
and protocols of every backend.

## Centralized configuration

Operational teams can change batching, filtering, enrichment, sampling,
routing, and export behavior in the Collector configuration rather than
rebuilding every application. This creates a consistent policy layer across
many services. Collector configuration still needs version control, testing,
safe rollout procedures, and ownership because one bad shared policy can
affect many applications.

## Decoupling applications from Tempo

Applications can send to a stable Collector or Alloy endpoint without knowing
where Tempo runs or how it is authenticated. Tempo can move, scale, change
credentials, or be replaced while the application-facing endpoint remains
stable. This does not make applications immune to telemetry pipeline failures:
the Collector itself must be deployed with suitable capacity and availability.

## Practical conclusion

Direct export is a useful option for learning and small installations:

```text
Application -> Tempo
```

For production, a shared collection layer commonly provides a safer and more
manageable boundary:

```text
Application
    -> nearby Collector or Alloy
    -> optional gateway Collector or Alloy
    -> Tempo
```

The Collector transports and processes traces. Tempo stores and queries them.

## Current references

- [OpenTelemetry Collector](https://opentelemetry.io/docs/collector/)
- [OpenTelemetry Collector resiliency](https://opentelemetry.io/docs/collector/resiliency/)
- [Tempo: set up a collector](https://grafana.com/docs/tempo/latest/set-up-for-tracing/instrument-send/set-up-collector/)
