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

## Health and shutdown

The service exposes separate process health and traffic-readiness endpoints:

```text
GET /livez
GET /readyz
```

Both return `200 OK` during normal operation. When the process receives
`SIGTERM` or `SIGINT`, readiness changes to `503 Service Unavailable` and the HTTP
server gets up to 30 seconds to finish in-flight requests before shutdown.

## Build and deploy to the homelab

The deployment workflow targets the ARM64 k3s cluster and pushes each build to
the local registry with a timestamp tag based on the workstation's local time:

```text
192.168.1.242:5000/image-resizer-api:vYYYY-MM-DD-HH-MM-SS
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
- A non-root user, read-only root filesystem, no Linux capabilities, and no
  mounted service-account token
- A LoadBalancer Service at `192.168.1.222`
- A cross-namespace `ServiceMonitor` that scrapes both replicas every 15 seconds

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
in the `monitoring` namespace, carries the `prometheus: homelab` discovery label,
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
`grafana/image-resizer-overview.json` and published to the homelab Grafana at
`http://192.168.1.221/d/image-resizer-overview/image-resizer-overview`. It
presents service health, time-range request totals, request outcomes and latency,
processing-stage latency, rejection reasons, and per-pod process CPU and memory.
