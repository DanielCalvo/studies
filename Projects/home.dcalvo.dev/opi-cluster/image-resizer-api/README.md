(this document is AI generated)

# Image Resizer API

A small Go HTTP service that accepts a JPEG, resizes it to a smaller width while
preserving its aspect ratio, and returns a JPEG.

## Run locally

```bash
/usr/local/go/bin/go run .
```

The service listens on port `8080` by default. Set `PORT` to use another port.

## Resize a JPEG

```bash
curl --fail-with-body \
  --output resized.jpg \
  --form image=@easy-clutch-plus-_1_.jpg \
  'http://192.168.1.222:8080/v1/resize?width=200'
```

The initial implementation accepts JPEG input only, does not upscale images,
limits uploads to 10 MiB, limits decoded images to 40 megapixels, and limits the
requested output width to 4,000 pixels.

## Logs

The service writes structured JSON logs to standard output. Each request receives
an `X-Request-ID` response header and produces one completion event containing
the request method, route, status, duration, available image dimensions and byte
sizes, and a stable error code when the request is rejected.

Image contents, multipart bodies, uploaded filenames, and arbitrary request
headers are not logged.

## Metrics

Prometheus metrics are exposed at:

```text
GET /metrics
```

The endpoint includes:

- Request totals, duration histograms, and in-flight requests
- Bounded request outcomes and rejection reasons
- Decode, resize, and encode duration histograms
- Active resize operations
- Input and output byte-size and pixel-count histograms
- Standard Go runtime and process metrics

Request IDs, filenames, exact dimensions, and raw error messages are not used as
metric labels.

## Traces

OpenTelemetry traces are exported with OTLP gRPC. In Kubernetes the application
sends them to Alloy at `alloy.monitoring.svc.cluster.local:4317`; Alloy batches
them and forwards them to Tempo.

Only `POST /v1/resize` is traced. Health and metrics requests are deliberately
excluded. A successful resize contains one HTTP server span and four
application-specific child spans:

```text
POST /v1/resize
├── image.read_upload
├── image.decode
├── image.resize
└── image.encode
```

The server span records the request outcome, stable application error code when
present, available image dimensions, target width, and input/output byte sizes.
Image contents, filenames, multipart bodies, and arbitrary headers are not
recorded.

Request-completion logs include `trace_id`, `span_id`, and `trace_sampled` when
there is an active trace. These remain JSON fields rather than Loki labels. The
existing request ID remains available independently in the log and
`X-Request-ID` response header.

The service uses parent-based, always-on sampling for this low-volume learning
environment and accepts standard W3C `traceparent` headers. Pending spans are
flushed for up to five seconds after the HTTP server finishes graceful shutdown.

To export local traces to a collector:

```bash
OTEL_EXPORTER_OTLP_ENDPOINT=http://127.0.0.1:4317 \
OTEL_EXPORTER_OTLP_INSECURE=true \
/usr/local/go/bin/go run .
```

In Grafana Explore, select the Tempo data source and find deployed application
traces with:

```traceql
{ resource.service.name = "image-resizer-api" }
```

Grafana is provisioned to navigate from an image-resizer span to the matching
Loki completion log and from a log's `trace_id` field back to its complete Tempo
trace.

## Health and shutdown

The service exposes separate process health and traffic-readiness endpoints:

```text
GET /livez
GET /readyz
```

Both return `200 OK` during normal operation. When the process receives
`SIGTERM` or `SIGINT`, readiness changes to `503 Service Unavailable` and the HTTP
server gets up to 30 seconds to finish in-flight requests before shutdown.

## Overload protection

The service admits at most one active resize per available Go execution thread.
By default, the limit comes from `runtime.GOMAXPROCS(0)`, which follows the Go
runtime's available CPU setting. Override it with a positive integer when
running an experiment:

```bash
MAX_CONCURRENT_RESIZES=2 /usr/local/go/bin/go run .
```

The slot is acquired before the multipart body is read. When every slot is
occupied, the service does not queue the request or read its image. It returns
`503 Service Unavailable` with `Retry-After: 1`, closes the connection, and logs
the request with the stable error code `server_overloaded`.

Slow clients are bounded by a 5-second header timeout and a 15-second whole
request read timeout. The existing 10 MiB upload limit and 40-megapixel decoded
image limit remain in effect.

Prometheus exposes overload rejections separately:

```text
image_resizer_overload_rejections_total
```

An invalid, zero, or negative `MAX_CONCURRENT_RESIZES` value prevents startup
instead of silently disabling admission control.

## Build and deploy to the homelab

The deployment workflow targets the ARM64 k3s cluster and pushes each build to
the local registry with a timestamp tag based on the workstation's local time:

```text
192.168.1.225:5000/image-resizer-api:vYYYY-MM-DD-HH-MM-SS
```

Run the complete build, push, apply, rollout, and endpoint validation workflow
with:

```bash
./build-and-deploy.sh
```

The checked-in Deployment is a template whose image ends in
`REPLACE_WITH_TAG`. The script generates a new image tag, renders a temporary
Deployment manifest using one explicit `sed` replacement, builds and pushes the
image for `linux/arm64`, and applies the rendered manifest. It then waits for the
rollout and MetalLB assignment and checks `/livez`, `/readyz`, and `/metrics`
through the external address.

The Kubernetes resources use:

- Namespace `image-resizer`
- Two replicas, spread across nodes when possible
- Readiness and liveness HTTP probes
- A 35-second pod termination grace period
- CPU and memory requests and limits
- CPU-aware resize concurrency with immediate overload rejection
- A Go soft memory limit below the container memory limit
- A non-root user, read-only root filesystem, no Linux capabilities, and no
  mounted service-account token
- A LoadBalancer Service at `192.168.1.222`
- A cross-namespace `ServiceMonitor` that scrapes both replicas every 15 seconds
- OTLP gRPC trace export to the internal Alloy Service

Every script run produces a unique pod-template image reference, which triggers a
Deployment rollout and works safely with `imagePullPolicy: IfNotPresent`. The
checked-in template keeps its placeholder; inspect the live Deployment to see the
currently deployed tag. Do not apply `k8s/deployment.yaml` directly.

Useful inspection commands are:

```bash
kubectl -n image-resizer get deployment,pods,service -o wide
kubectl -n image-resizer get deployment image-resizer-api -o jsonpath='{.spec.template.spec.containers[0].image}{"\n"}'
kubectl -n image-resizer logs -l app.kubernetes.io/name=image-resizer-api --prefix=true
curl http://192.168.1.222/livez
curl http://192.168.1.222/metrics
```

Prometheus is available at `http://192.168.1.223`. The ServiceMonitor is created
in the `monitoring` namespace, carries the `prometheus: opi` discovery label,
and selects the `image-resizer-api` Service in the `image-resizer` namespace.
The deployment script applies it along with the application resources.

Example PromQL queries:

```promql
up{namespace="image-resizer", service="image-resizer-api"}
sum by (outcome) (rate(image_resizer_http_requests_total[5m]))
```

## Post-deployment smoke tests

Run the functional smoke suite against the homelab deployment with:

```bash
./smoke-tests/post-deployment-smoke-test.sh
```

It generates its fixtures under `smoke-tests/test-data/` and checks the health and
metrics endpoints, a successful JPEG resize, request-ID behavior, unsupported
formats, corrupt input, upload and pixel limits, upscaling, and invalid width
handling. A different environment URL can be supplied as the first argument.

## Generate background traffic

To continuously send one resize request every five seconds for metrics, logs,
and dashboard experiments, run:

```bash
./traffic-gen/run.sh
```

The traffic generator cycles through ten generated JPEGs ranging from `640x480`
to `5000x5000` and resizes each to half its original dimensions. It randomizes
the traversal order once per run. It is an activity generator, not a capacity or
load test. See `traffic-gen/README.md` for details.

## Grafana dashboard

The version-controlled dashboard definition is stored at
`hp-cluster/monitoring/grafana/dashboards/image-resizer/image-resizer-overview.json`
and published to the HP Grafana at
`http://192.168.1.231/d/image-resizer-overview/image-resizer-overview`. It
presents service health, time-range request totals, request outcomes and latency,
processing-stage latency, rejection reasons, and per-pod process CPU and memory.

The Loki-backed request-log dashboard is stored at
`hp-cluster/monitoring/grafana/dashboards/image-resizer/image-resizer-logs.json`
and published at
`http://192.168.1.231/d/image-resizer-logs/image-resizer-logs`. It presents
request counts, 4xx and 5xx outcomes, error reasons, log-derived latency and
payload sizes, per-pod traffic, recent unsuccessful requests, and the complete
request log stream.
