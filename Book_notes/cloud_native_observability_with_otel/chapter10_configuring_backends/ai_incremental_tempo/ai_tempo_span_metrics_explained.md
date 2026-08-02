# Tempo span metrics explained

Tempo's span-metrics processor provides a predefined family of metrics, while
its configuration controls how those metrics are grouped, filtered, and
aggregated.

The most important distinction is that Tempo processes individual spans. It
does not treat an entire trace as one indivisible execution.

## What Tempo creates automatically

Enabling the complete span-metrics processor looks like this:

```yaml
overrides:
  defaults:
    metrics_generator:
      processors:
        - span-metrics
```

Tempo then examines every received span and produces three predefined metrics:

| Metric | Type | Meaning |
|---|---|---|
| `traces_spanmetrics_calls_total` | Counter | Number of matching spans |
| `traces_spanmetrics_latency` | Histogram | Distribution of matching span durations |
| `traces_spanmetrics_size_total` | Counter | Total bytes occupied by matching spans |

These metric names and fundamental meanings are defined by Tempo. The user does
not invent them.

There is no separate predefined `errors_total` metric. Errors can be selected
from the calls counter through the `status_code="STATUS_CODE_ERROR"` label.

See the official
[Tempo span-metrics documentation](https://grafana.com/docs/tempo/latest/metrics-from-traces/span-metrics/span-metrics-metrics-generator/)
for the current configuration and metric definitions.

## A histogram accumulates many span executions

Suppose a service produces the same operation span three times:

```text
service:   names-api
span_name: GET /names
span_kind: SERVER
status:    OK

durations:
10 ms
70 ms
300 ms
```

Because all three spans have the same metric dimensions, Tempo groups them into
the same time series. The counter is conceptually:

```text
traces_spanmetrics_calls_total{
  service="names-api",
  span_name="GET /names",
  span_kind="SPAN_KIND_SERVER",
  status_code="STATUS_CODE_OK"
} 3
```

The histogram receives three observations:

```text
0.010 seconds
0.070 seconds
0.300 seconds
```

Its associated count and sum become:

```text
latency_count = 3
latency_sum   = 0.380
```

Example cumulative buckets would contain:

```text
le="0.064"  -> 1 span
le="0.128"  -> 2 spans
le="0.512"  -> 3 spans
```

Tempo therefore does not create a new histogram for every trace. It
continually adds span-duration observations to histogram time series.

## What happens when a trace has several spans

Consider a trace with this structure:

```text
HTTP request span             400 ms
├── database query span        80 ms
└── outbound HTTP span        200 ms
```

Each span can contribute its own observation:

```text
HTTP span      -> HTTP operation's counter and histogram
Database span  -> database operation's counter and histogram
Outbound span  -> client operation's counter and histogram
```

Tempo does not create a special histogram observation for the complete trace's
400 ms execution time. The root HTTP span may represent that end-to-end request
duration, but that is because the root span itself lasted 400 ms.

## Default dimensions

By default, Tempo groups span metrics using four dimensions:

```text
service
span_name
span_kind
status_code
```

For example, these operations become separate time series:

```text
service="names-api", span_name="GET /names"
service="names-api", span_name="POST /names"
service="database",  span_name="SELECT names"
```

Successful and failed spans also become separate series because their status
codes differ:

```text
status_code="STATUS_CODE_OK"
status_code="STATUS_CODE_ERROR"
```

The rate of error spans can then be selected with PromQL:

```promql
rate(
  traces_spanmetrics_calls_total{
    status_code="STATUS_CODE_ERROR"
  }[5m]
)
```

An error percentage can be calculated by dividing the rate of failed calls by
the rate of all calls.

## What the current exercise configures

The exercise currently enables:

```yaml
overrides:
  defaults:
    metrics_generator:
      processors:
        - span-metrics
        - service-graphs
```

Enabling the complete `span-metrics` processor generates all three predefined
span metrics:

- count
- latency
- size

The exercise has not customized their dimensions, filtering, or histogram
buckets. It is currently demonstrating Tempo's defaults.

## What can be configured

### Select individual metric types

Instead of the complete processor, individual subprocessors can be selected:

```yaml
overrides:
  defaults:
    metrics_generator:
      processors:
        - span-metrics-latency
        - span-metrics-count
        # span-metrics-size is deliberately omitted
```

This example generates the latency histogram and calls counter but omits the
span-size counter.

### Change histogram buckets

Tempo supplies default duration buckets, but applications with different
latency expectations can provide their own:

```yaml
metrics_generator:
  processor:
    span_metrics:
      histogram_buckets:
        - 0.01
        - 0.05
        - 0.1
        - 0.25
        - 0.5
        - 1.0
```

These boundaries are expressed in seconds.

### Add span or resource attributes as labels

Suppose spans contain:

```text
http.request.method="GET"
deployment.environment.name="production"
```

They can be added as metric dimensions:

```yaml
metrics_generator:
  processor:
    span_metrics:
      dimensions:
        - http.request.method
        - deployment.environment.name
```

The resulting Prometheus labels are sanitized:

```text
http_request_method="GET"
deployment_environment_name="production"
```

Additional dimensions allow more detailed querying, but each dimension can
increase metric cardinality.

### Disable default dimensions

Individual intrinsic dimensions can be disabled:

```yaml
metrics_generator:
  processor:
    span_metrics:
      intrinsic_dimensions:
        service: true
        span_name: true
        span_kind: false
        status_code: true
```

In this example, `span_kind` no longer separates the generated series.

### Filter which spans generate metrics

Include and exclude policies can restrict the spans considered by the
processor. A common example is excluding health-check spans so `/health`
traffic does not distort application request metrics.

Span metrics normally apply to all received spans. They can be filtered by
useful span or resource information, but they are not normally configured for
one particular trace ID.

## What span metrics cannot define

The span-metrics processor is not a general-purpose metric programming system.
It cannot be configured to invent an arbitrary business metric such as:

```text
names_deleted_total
```

That metric would normally be created explicitly in the application using an
OpenTelemetry counter. Alternatively, TraceQL metrics can calculate suitable
measurements at query time when the required information already exists in the
traces.

The overall model is:

```text
Enable span-metrics
       |
       v
Receive predefined count, latency, and size metrics
       |
       v
Customize dimensions, buckets, filtering, and enabled subsets
```

The current exercise deliberately begins with all three predefined metrics,
Tempo's default labels, and Tempo's default histogram buckets.
