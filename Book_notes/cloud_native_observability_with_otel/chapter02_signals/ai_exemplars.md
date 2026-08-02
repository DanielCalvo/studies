Yes—**exemplars made it**. They're no longer just an idea in the OpenTelemetry spec; they're a real feature that's supported across much of the observability ecosystem. ([OpenTelemetry][1])

The catch is that they're not "metrics with trace information" in the sense of every metric carrying a trace. Instead, they're **sampled metric points** that contain a reference to the trace that produced them.

For example, suppose your service exports an HTTP latency histogram:

```
http.server.duration
```

One particular request takes 2.3 seconds. The SDK might attach an exemplar to that observation:

```
Bucket: >2s
Value: 2.3s
TraceID: 4bf92...
SpanID: 00f067...
Timestamp: ...
```

Then, when you're looking at a Grafana graph, you'll see a little marker (a star) on the graph. Clicking it jumps straight into the distributed trace for that request. It's an incredibly nice workflow:

1. "Why did latency spike at 14:32?"
2. Click the exemplar.
3. You're immediately looking at the exact trace that contributed to that spike. ([Grafana Labs][2])

For someone learning OpenTelemetry today, it's worth knowing a few practical points:

* OpenTelemetry SDKs support generating exemplars.
* Prometheus and OpenMetrics support exemplars, although **OpenMetrics** has better support than the classic Prometheus text exposition format. ([OpenTelemetry][1])
* Grafana has first-class support for visualizing them with Prometheus metrics and linking them to Tempo (or another tracing backend). ([Grafana Labs][2])

Since I know you're building a Kubernetes homelab with an LGTM stack, exemplars are actually one of the features I'd encourage you to enable once you have the basics working. They really show off the value of OpenTelemetry because you can go:

**Metric → Trace** with a single click, instead of manually hunting by timestamps.

One small caveat: because exemplars are only attached to **sampled observations**, not every point on a graph has one. That's intentional—they're designed to give you representative links without exploding storage requirements. ([opentelemetry.netlify.app][3])

[1]: https://opentelemetry.io/docs/specs/otel/compatibility/prometheus_and_openmetrics/?utm_source=chatgpt.com "Prometheus and OpenMetrics Compatibility | OpenTelemetry"
[2]: https://grafana.com/docs/grafana/latest/fundamentals/exemplars/?utm_source=chatgpt.com "Introduction to exemplars | Grafana documentation"
[3]: https://opentelemetry.netlify.app/docs/languages/dotnet/metrics/exemplars/?utm_source=chatgpt.com "Using exemplars | OpenTelemetry"
