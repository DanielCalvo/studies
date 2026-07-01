# Pushgateway App

Go version of the Chapter 4 Pushgateway batch-job example from
`Prometheus: Up & Running`.

It pushes:

- `my_job_duration_seconds`
- `my_job_last_success_seconds` only when the simulated job succeeds

Run locally while the Compose stack is up:

```bash
go run .
```

Run through Compose:

```bash
cd ../prometheus_config
docker compose run --rm --no-deps pushgateway_app
```

Or from this folder:

```bash
./run-job.sh
```

This assumes the main Compose stack is already running, including Pushgateway and
Prometheus.

Useful PromQL:

```promql
my_job_duration_seconds
time() - my_job_last_success_seconds
push_time_seconds{job="batch"}
```

Configuration:

- `PUSHGATEWAY_URL`, default `http://localhost:9092`
- `JOB_NAME`, default `batch`
- `WORK_DURATION_SECONDS`, default `1`
- `FAIL_JOB=true` to simulate a failed job
