# Metrics App

Go version of the Chapter 3 Hello World instrumentation example from
`Prometheus: Up & Running`.

The app serves:

- `http://localhost:8001/` for the Hello World endpoint
- `http://localhost:8000/metrics` for Prometheus metrics

Metrics added from the chapter examples:

- `hello_worlds_total`
- `hello_world_exceptions_total`
- `hello_world_sales_euro_total`
- `hello_worlds_inprogress`
- `hello_world_last_time_seconds`
- `time_seconds`
- `hello_world_latency_seconds`

Run locally:

```bash
go run .
```

Useful PromQL examples:

```promql
rate(hello_worlds_total[1m])
rate(hello_world_exceptions_total[1m]) / rate(hello_worlds_total[1m])
rate(hello_world_sales_euro_total[1m])
time() - hello_world_last_time_seconds
histogram_quantile(0.95, rate(hello_world_latency_seconds_bucket[1m]))
```

Set `ERROR_RATE=0` to disable the random simulated failures, or request
`http://localhost:8001/?fail=true` to force one failure.

