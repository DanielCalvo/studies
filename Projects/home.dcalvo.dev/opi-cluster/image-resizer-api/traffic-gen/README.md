# Traffic generator

This small Go program creates steady application traffic for exploring logs,
Prometheus metrics, and dashboards. It is not a capacity or load test.

From this directory, run it against the homelab deployment at the default rate
of 60 requests per minute:

```bash
/usr/local/go/bin/go run .
```

Set a different request rate with `-rpm`:

```bash
/usr/local/go/bin/go run . -rpm 30
```

The program creates any missing fixtures under `test-data/`, then sends one
randomly selected JPEG per request. Each image is resized to half its original
width:

- `640x480`, resized to `320x240`
- `800x1200`, resized to `400x600`
- `1200x1200`, resized to `600x600`
- `1600x1200`, resized to `800x600`
- `2000x1500`, resized to `1000x750`
- `2400x1800`, resized to `1200x900`
- `2000x3000`, resized to `1000x1500`
- `3200x2400`, resized to `1600x1200`
- `4000x3000`, resized to `2000x1500`
- `5000x5000`, resized to `2500x2500`

Random selection is independent for every request, so some images will naturally
be sent more often than others. The configured rate controls the interval between
request starts. Each request runs independently, so a slow response does not
delay the next request. At 60 requests per minute, the program launches one
request every second even when earlier requests are still in flight.

This is an open-loop traffic generator: if the service cannot keep up, concurrent
in-flight requests can accumulate until they finish or reach the 30-second HTTP
timeout. Individual request failures are logged without stopping the scheduler,
so temporary EOFs, timeouts, connection failures, and HTTP error responses do not
interrupt traffic generation. On `Ctrl+C`, the program cancels its in-flight
requests and waits for their goroutines before exiting.

Every resize-start line includes the current number of requests in flight:

```text
[in_flight=1] Starting resize of 640x480.jpg from 640x480 to 320x240
```

Request completions are not logged, keeping the continuous output to one line per
launched request. A failed request produces an additional line because the
scheduler is continuing:

```text
[in_flight=2] Resize of 3200x2400.jpg failed; continuing: send resize request: Post "http://192.168.1.222/v1/resize?width=1600": EOF
```

Stop the program with `Ctrl+C`. Other useful flags are:

```text
-base-url http://192.168.1.222
-images test-data
```

For example, to send 30 requests per minute to another endpoint:

```bash
/usr/local/go/bin/go run . -rpm 30 -base-url http://127.0.0.1:8080
```
