# Image Resizer API Feature Sequence

This file records the order in which features are implemented and why that order
was chosen. The purpose is to make the reasoning and dependencies between steps
visible, not merely to maintain a checklist.

`ai_feature_ideas.md` contains the broader collection of possible exercises. This
file is the chronological record of what we actually build.

## How to maintain this file

- Add an entry only when a feature is selected for implementation.
- Give every entry a sequence number and a status.
- Record why the feature belongs at that point in the sequence.
- Identify which earlier capability it depends on and what later work it enables.
- Preserve completed entries so that the project's evolution remains visible.
- Keep unselected ideas in the candidate section rather than presenting them as
  an agreed implementation order.

## Implemented sequence

### 1. Minimal JPEG resizing service

**Status:** Implemented

The first version establishes a complete, usable path through the application:

```text
HTTP upload -> validate -> decode JPEG -> resize -> encode JPEG -> HTTP response
```

It includes:

- A Go HTTP server
- `POST /v1/resize?width=<pixels>`
- One JPEG uploaded in the multipart field named `image`
- Aspect-ratio-preserving resizing to a smaller width
- JPEG output at a fixed quality
- In-memory, synchronous processing
- Rejection of upscaling
- Limits on upload size, decoded pixel count, and output width
- JSON error responses
- Tests for successful resizing and important rejection cases
- Local run and request instructions

#### Why this came first

Observability and reliability practices need real application behavior to
observe. Before adding metrics, traces, service-level indicators, or operational
objectives, the project needs a small but complete user-visible operation with
meaningful success, failure, latency, CPU, and memory behavior.

This vertical slice also prevents us from designing operational machinery around
an application that does not yet work. Its narrow scope gives later measurements
a stable meaning: a successful request is a valid JPEG resize, and failures have
specific causes.

#### What this enables

Later features can now measure or improve concrete behavior, including:

- Request rate, errors, and duration
- JPEG decoding, resizing, and encoding time
- Input and output sizes
- Resource saturation and overload behavior
- User-visible success and latency
- Diagnosis of slow or failed requests

### 2. Structured request and application logs

**Status:** Implemented

The service now emits JSON logs to standard output using Go's `log/slog`
package. It records the active address and safety limits at startup and writes one
completion event for every resize request.

Each request receives a generated request ID in the `X-Request-ID` response
header. The corresponding completion event contains the fields that are known at
the point where processing ends:

- Request ID, method, route, HTTP status, and total duration
- Target width
- Input width, height, and encoded byte size
- Output height and encoded byte size
- A stable application error code for rejected or failed requests

Image contents, multipart bodies, filenames, complete query strings, and
arbitrary request headers are deliberately excluded.

#### Why this came second

Once the minimal user operation worked, the next immediate need was to understand
individual requests without attaching a debugger. Structured logs make success,
rejection, latency, and the available image characteristics visible while the
application is still small enough to establish consistent event names and fields.

This came before aggregate metrics and reliability objectives because those
features answer different questions. Logs first provide evidence about individual
events and help verify which outcomes the later metrics should count.

#### What this enables

- Following one request by its identifier
- Diagnosing rejected and failed requests from stable error codes
- Comparing request latency with image dimensions and encoded sizes
- Reusing the request ID as correlation context for future traces
- Designing metrics from already-defined user-visible outcomes

### 3. Prometheus application and domain metrics

**Status:** Implemented

The service now exposes a Prometheus-compatible `GET /metrics` endpoint using an
application-owned registry. It includes:

- HTTP request totals labeled by bounded method, route, status, and outcome
- Request-duration histograms and in-flight request gauges
- Decode, resize, and encode duration histograms
- A gauge for active image-processing operations
- Input and output encoded-size histograms
- Input and output pixel-count histograms
- Rejection totals labeled by stable application error code
- Standard Go runtime and process collectors

Request IDs, filenames, exact dimensions, complete URLs, and raw error messages
are deliberately excluded from metric labels to prevent unbounded cardinality.

#### Why this came third

The working service first established meaningful behavior, and structured logs
then made individual requests understandable. Metrics are the next layer because
they aggregate those outcomes across many requests and over time. They let us see
rates, latency distributions, workload sizes, resource behavior, and saturation
instead of inspecting one event at a time.

Instrumentation also comes before selecting protection thresholds or reliability
targets. We can now measure the application's actual behavior under load rather
than choosing concurrency limits and objectives without evidence.

#### What this enables

- Establishing a performance and resource baseline through load testing
- Comparing latency with input size and pixel count
- Identifying which processing stage dominates request duration
- Recognizing concurrency and resource saturation
- Defining a measurable success and latency SLI
- Choosing an initial SLO and later evaluating error-budget consumption

### 4. Kubernetes runtime lifecycle contract

**Status:** Implemented

The service now exposes two lightweight lifecycle endpoints:

- `GET /livez` reports whether the application process is alive.
- `GET /readyz` reports whether the application is willing to receive traffic.

During normal operation both return `200 OK`. When the process receives
`SIGTERM` or `SIGINT`, it first changes readiness to `503 Service Unavailable`,
stops accepting new connections, and gives in-flight requests up to 30 seconds to
finish before shutdown. Liveness remains conceptually separate from readiness
during that transition.

The readiness state is concurrency-safe, and tests cover both the endpoint state
transition and completion of the bounded shutdown path.

#### Why this came fourth

The next planned vertical slice is a containerized Kubernetes deployment.
Readiness, liveness, and termination behavior form the runtime contract between
the application and Kubernetes, so implementing them before writing probe and
Deployment configuration avoids deploying placeholders with unclear semantics.

This also preserves the earlier sequencing principle: application behavior is
defined and tested before the platform is configured to depend on it.

#### What this enables

- Meaningful Kubernetes readiness and liveness probes
- Removal from service traffic when termination begins
- Graceful completion of in-flight resize requests during rollout or eviction
- A bounded pod termination grace period
- Safer rolling deployments and node maintenance

### 5. ARM64 container build and first Kubernetes deployment

**Status:** Implemented and deployed

The application is now packaged with a multi-stage Docker build. Go compiles a
static `linux/arm64` binary on the amd64 workstation, and the final `scratch`
image contains only that binary and runs as a numeric non-root user.

The `build-and-deploy.sh` workflow:

1. Checks Docker, Buildx, `kubectl`, and registry access.
2. Generates a local-time tag in `vYYYY-MM-DD-HH-MM-SS` format.
3. Builds and pushes the ARM64 image to the homelab registry.
4. Replaces `REPLACE_WITH_TAG` while rendering a temporary Deployment manifest.
5. Applies the saved namespace and Service plus the rendered Deployment.
6. Waits for the rollout and requested MetalLB address.
7. Verifies the health and metrics endpoints through the external address.

The initial Kubernetes deployment uses:

- Timestamped images under `192.168.1.225:5000/image-resizer-api`
- Namespace `image-resizer`
- Two replicas with topology spreading across cluster nodes
- HTTP readiness and liveness probes
- Resource requests and limits
- A 35-second pod termination grace period
- A restricted non-root, read-only container security context
- A LoadBalancer Service at `192.168.1.222`

The first rollout placed one healthy replica on each Orange Pi. An external
end-to-end request resized a `94x94` JPEG to `60x60`; the request appeared in the
structured logs and updated the application and processing-stage metrics.

#### Why this came fifth

The service first needed useful behavior, individual-request visibility,
aggregate metrics, and a clear runtime lifecycle contract. With those foundations
in place, deploying it to Kubernetes produces an environment we can understand
and measure instead of introducing several opaque platform concerns at once.

The shell script is intentionally a small homelab release mechanism rather than a
production delivery system. Kubernetes resource structure remains represented by
declarative manifests, while the script renders the release-specific image and
removes repetitive build, push, apply, wait, and validation steps.

#### What this enables

- Prometheus discovery with a `ServiceMonitor`
- Grafana dashboards based on real cluster traffic
- Load testing on the ARM64 hardware
- Measurement-driven concurrency and resource tuning
- Rolling-update and graceful-termination experiments
- Comparing traffic distribution across replicas

### 6. Post-deployment functional smoke tests

**Status:** Implemented and passing against the homelab deployment

The project now contains a self-contained smoke-test runner at
`smoke-tests/post-deployment-smoke-test.sh`. It targets the MetalLB endpoint by
default and accepts another base URL for testing a different environment.

The suite generates all fixtures under `smoke-tests/test-data/` and verifies:

- Liveness, readiness, and metrics endpoint availability
- Successful resizing of a `120x80` JPEG to `60x40`
- A valid JPEG response and an `X-Request-ID` header
- Rejection of PNG, corrupt JPEG, text, and MP4 inputs
- Rejection of uploads larger than 10 MiB
- Rejection of JPEG dimensions exceeding the pixel limit
- Rejection of upscaling and zero-width requests
- Stable HTTP statuses and application error codes for every rejection
- Continued liveness and readiness after the invalid requests

The current suite contains 14 checks and passes against
`http://192.168.1.222`.

#### Why this came sixth

In-process tests already verify application behavior quickly, but they cannot
prove that the ARM64 image, registry, Kubernetes configuration, Service routing,
MetalLB exposure, and deployed process work together. Once the first deployment
was operational, a small functional smoke suite became the fastest way to catch a
broken release or an incorrect runtime configuration.

The suite deliberately covers only representative success and rejection paths.
It provides rapid release confidence without becoming a comprehensive deployed
regression suite.

#### What this enables

- Repeating consistent validation after every ad hoc deployment
- Separating application test failures from packaging or platform failures
- Expanding selected scenarios into a broader API regression suite later
- Reusing critical journeys as future synthetic monitoring checks

### 7. Prometheus discovery and collection

**Status:** Implemented and actively scraped

A `ServiceMonitor` in the `monitoring` namespace now selects the Image Resizer
Service in the `image-resizer` namespace. It asks the homelab Prometheus instance
to scrape the named `http` port at `/metrics` every 15 seconds.

Prometheus discovered both application replicas through the Service and reports
both targets as healthy. Queries against Prometheus also return the application's
request counter and image-processing stage histograms, confirming that the
metrics are being ingested rather than merely exposed by the application.

#### Why this came seventh

The application already emitted useful metrics, but those metrics only became a
durable measurement source after the service was deployed and Prometheus was
configured to collect them. The post-deployment smoke suite first provided a
repeatable way to generate and validate representative traffic; Prometheus can
now retain and aggregate the resulting observations across both replicas.

This belongs before dashboards, alerts, SLIs, and SLOs because each of those
depends on verified time-series collection. Defining them before confirming
discovery and ingestion would build operational decisions on an untested data
path.

#### What this enables

- Querying application behavior across both replicas with PromQL
- Building Grafana dashboards from real application traffic
- Defining user-visible success and latency SLIs
- Establishing an initial SLO and measuring error-budget consumption
- Alerting and load-testing exercises using retained measurements

### 8. Low-rate traffic generator

**Status:** Implemented

The project now includes `traffic-gen/run.sh`, a deliberately simple background
traffic source for the deployed API. It repeatedly cycles through ten generated
JPEGs ranging from `640x480` to `5000x5000`, resizing each to half its original
dimensions. The images include square, landscape, and portrait aspect ratios.

The deterministic fixtures live under `traffic-gen/test-data/` and are generated
locally when the script starts. The JPEG list is built and shuffled once per run,
then reused for continuous sequential traffic. The script prints which resize it
is sending, and `curl` stops the script if the service returns an HTTP error.

#### Why this came eighth

Prometheus is now collecting the application metrics, but dashboards and PromQL
experiments are easier to understand when the service receives ongoing traffic
with several repeatable image sizes. A low-rate sequential generator supplies
that activity without prematurely turning the exercise into capacity testing or
introducing concurrency controls and test-run statistics.

#### What this enables

- Populating request, processing-stage, image-size, and runtime metrics
- Observing a steady stream of structured request logs
- Creating and validating initial Grafana panels and PromQL expressions
- Comparing latency across ten repeatable image dimensions
- Evolving the utility into a real concurrent load test later if useful

### 9. Image Resizer Grafana overview

**Status:** Implemented, published, and validated with live traffic

The `Image Resizer Overview` dashboard is published to the homelab Grafana with
the stable UID `image-resizer-overview`. Its version-controlled API payload lives
at `grafana/image-resizer-overview.json` and contains twelve panels:

- Healthy Prometheus targets
- Total, successful, and failed request counts for the selected time range
- Requests per minute over time
- Request rate grouped by application outcome
- Successful-request percentage
- Successful-request p50, p95, and p99 latency
- Processing-stage p95 latency grouped by decode, resize, and encode
- Rejected-request rate grouped by stable reason
- Process CPU usage grouped by pod
- Resident process memory grouped by pod

The dashboard was read back from Grafana after publication to confirm all twelve
panels use the expected Prometheus datasource. Traffic-generator and smoke-test
requests then confirmed that every metric family returns the intended live
series, including both replicas, all three processing stages, and five rejection
reasons.

#### Why this came ninth

Application metrics first needed to be defined, exposed, collected by
Prometheus, and populated with repeatable traffic. With that complete, a small
overview dashboard turns the raw time series into a practical operational view
without introducing alert thresholds or reliability objectives prematurely.

The initial dashboard stays deliberately focused on twelve panels. This gives us
an opportunity to learn which views are genuinely useful before adding more
dashboard structure and maintenance surface.

#### What this enables

- Observing health, traffic, latency, failures, and runtime cost in one place
- Comparing application behavior across the two replicas
- Relating total latency to decode, resize, and encode time
- Selecting a user-visible success and latency SLI from real measurements
- Choosing evidence-based alert thresholds and an initial SLO later

### 10. Rate-configurable Go traffic generator

**Status:** Implemented

The original shell traffic generator has been replaced by a small Go program.
It preserves the generated JPEG fixtures, half-width resize requests, sequential
request behavior, and default homelab endpoint. Its request rate is configurable
in requests per minute, and it independently selects a random fixture for every
request so the observed image-size frequencies vary over time.

The program reads each JPEG's real dimensions rather than relying on its
filename, supports a configurable image directory and base URL, stops cleanly on
`Ctrl+C`, and treats HTTP error responses as failures. Focused tests verify rate
conversion, image discovery, multipart upload behavior, target resize width, and
HTTP error handling.

#### Why this came tenth

The initial generator successfully populated the Grafana overview, which made
its fixed 0.1-second pause and once-per-run shuffle the next practical
limitations. A requests-per-minute setting makes slow demonstrations and denser
dashboard experiments equally easy, while per-request random selection produces
more natural variation without turning the utility into a concurrent capacity
test.

#### What this enables

- Changing dashboard traffic density without editing source code
- Observing naturally uneven image-size distributions
- Pointing the same utility at local or alternate deployments
- Adding bounded concurrency and run statistics later if load testing becomes
  useful

### 11. Open-loop traffic scheduling

**Status:** Implemented

The Go traffic generator now launches each request in its own goroutine on the
configured schedule. Request completion no longer delays the next launch: at 60
requests per minute, a request starts every second even if earlier requests are
still in flight.

The shared HTTP client safely reuses connections across concurrent requests.
Each request retains the existing 30-second timeout. On `Ctrl+C` or the first
request failure, the scheduler cancels all in-flight requests and waits for their
goroutines before exiting.

#### Why this came eleventh

The rate-configurable sequential generator was useful for steady dashboard
traffic, but its achieved rate remained limited by single-request latency.
Open-loop scheduling separates the requested arrival rate from response time,
making the requests-per-minute setting useful for deliberately applying load.

#### What this enables

- Maintaining a requested arrival rate while requests overlap
- Observing queueing, latency, and saturation when the service cannot keep up
- Comparing offered traffic with completed-request throughput
- Adding an optional maximum in-flight limit later if a safety cap is useful

### 12. In-flight request visibility

**Status:** Implemented

Every resize-start line now includes the current number of in-flight requests.
Request completions remain silent, keeping the continuous output to one line per
launched request. The counter uses atomic operations so concurrent request
completions can update the value safely before the next request-start line.

#### Why this came twelfth

Open-loop scheduling intentionally allows requests to overlap. Once completion
was decoupled from launch timing, the generator needed a direct way to show how
much concurrent work its configured arrival rate was creating.

#### What this enables

- Seeing concurrency accumulate when the service falls behind
- Relating the offered request rate to observed service latency
- Choosing a future maximum in-flight safety limit from observed behavior

### 13. Continue traffic after request failures

**Status:** Implemented

Individual request failures no longer stop the open-loop scheduler. EOFs,
timeouts, connection failures, and HTTP error responses are logged with the
current in-flight count, while later requests continue to launch at the
configured rate. Startup and configuration failures remain fatal because the
generator cannot begin useful traffic in those cases.

#### Why this came thirteenth

Applying enough traffic to reach a service or platform throughput limit can
produce transient connection failures. Stopping on the first such failure ended
the experiment precisely when the system began showing interesting overload
behavior, so request errors now remain observations rather than generator-level
failures.

#### What this enables

- Sustaining offered traffic through transient service failures
- Observing recovery while the same traffic rate continues
- Comparing error frequency with concurrency and service saturation
- Running longer overload experiments without manually restarting the generator

### 14. Bounded resize concurrency and overload rejection

**Status:** Implemented

The API now bounds active resize requests before reading their multipart bodies.
The default concurrency is `runtime.GOMAXPROCS(0)`, matching the Go runtime's
available CPU parallelism, and a positive `MAX_CONCURRENT_RESIZES` value can
override it for experiments.

When all processing slots are occupied, another request is not queued and its
body is not read. The API immediately returns `503 Service Unavailable` with
`Retry-After: 1`, closes the connection, emits a structured completion log with
the stable `server_overloaded` error code, and increments
`image_resizer_overload_rejections_total`. Slots are released after every
successful or failed admitted request.

The HTTP server also bounds slow clients with header, whole-request read, write,
and idle timeouts plus a one-mebibyte header limit. The Kubernetes deployment
now requests one CPU and 512 MiB of memory per replica, permits up to 1536 MiB,
and sets `GOMEMLIMIT=1280MiB` so garbage collection has headroom before the
container boundary. A commented `MAX_CONCURRENT_RESIZES` environment-variable
example in the Deployment documents how to override the default without enabling
the override.

#### Why this came fourteenth

Open-loop load testing exposed the first measured capacity failure: requests
accumulated faster than the two ARM64 replicas could process them, concurrent
decoded images exhausted the original 512 MiB limit, and both pods repeatedly
terminated with `OOMKilled`. That lost in-flight request logs and temporarily
removed every Service endpoint.

Admission control was therefore prioritized over defining reliability targets.
It changes overload from an uncontrolled pod failure into an explicit,
observable API outcome while keeping health endpoints available.

#### What this enables

- Sustaining overload experiments without unbounded application concurrency
- Measuring explicit overload responses instead of inferring lost requests from
  pod restarts
- Tuning concurrency independently from image-size and memory limits
- Comparing offered load, completed throughput, and overload rejection rate
- Establishing capacity-aware alerts, SLIs, and SLOs from stable behavior

### 15. Image Resizer Loki log dashboard

**Status:** Implemented, published, and validated with live logs

The `Image Resizer Logs` dashboard is published to the homelab Grafana with the
stable UID `image-resizer-logs`. Its version-controlled API payload lives at
`grafana/image-resizer-logs.json` and contains twelve Loki-backed panels:

- Total and successful request counts for the selected time range
- Separate client rejection (4xx) and server failure (5xx) counts
- Log-derived successful-request percentage
- Request counts grouped by HTTP status, application error, and pod
- Successful-request p50 and p95 latency grouped by pod
- Successful input and output payload p95 sizes grouped by pod
- Recent unsuccessful request logs with status, error, duration, request ID,
  and pod
- Recent logs for every completed request

The dashboard defaults to the last hour so load experiments and their
`server_overloaded` responses remain visible after the traffic generator stops.
All request fields are parsed from the JSON log body at query time; high-cardinality
fields such as request IDs remain unindexed.

The dashboard queries were validated directly against Loki before publication.
Live log data included successful resize traffic on both replicas, client-side
`method_not_allowed` responses, and service-side `server_overloaded` failures.

#### Why this came fifteenth

Loki and Alloy made the application logs durable and queryable, while overload
testing created enough successful and failed traffic to decide which log views
were operationally useful. A dedicated dashboard now makes those logs usable
without reconstructing LogQL queries during every investigation.

Keeping this separate from the Prometheus overview also makes the distinction
between metrics and logs explicit: the overview is the efficient aggregate
health view, while the log dashboard provides request-level evidence and IDs for
drill-down.

#### What this enables

- Seeing recent failed requests without manually writing LogQL
- Distinguishing client rejection from service failure
- Finding the error reason, request ID, and serving pod for an incident
- Comparing log-derived latency and traffic distribution across replicas
- Correlating the Loki dashboard with the existing Prometheus overview
- Building future alerts from error patterns that prove useful

### 16. OpenTelemetry request tracing

**Status:** Implemented, deployed, and validated with Tempo and Loki

The API now creates an OpenTelemetry HTTP server span for every
`POST /v1/resize` request and four child spans for the application-specific
processing stages:

```text
POST /v1/resize
├── image.read_upload
├── image.decode
├── image.resize
└── image.encode
```

Health and Prometheus endpoints are excluded from tracing. The server span
records a bounded outcome and stable error code plus the available image
dimensions, target width, and encoded byte sizes. Child spans expose where an
individual request spent its time. The service accepts W3C Trace Context,
samples new traces in this low-volume learning environment, batches completed
spans, and exports them over OTLP gRPC to Alloy.

The existing structured request-completion log now includes the active trace ID,
span ID, and sampled flag. These high-cardinality identifiers remain fields in
the JSON body rather than Loki labels. The request ID remains an independent
application identifier.

Focused tests use an in-memory exporter to verify incoming context propagation,
the processing-span hierarchy, success and rejection attributes, trace/log
identifier correlation, and the absence of traces for liveness checks.

The ARM64 image was deployed to both Orange Pis and the complete post-deployment
smoke suite passed. Its nine resize requests produced twenty spans. The
successful `120x80` to `60x40` resize produced the expected server span and four
child spans; rejected requests produced shorter traces with their stable error
codes. Loki stored the matching completion logs with trace and span IDs, and
Grafana's provisioned Tempo and Loki data sources expose bidirectional
trace-to-log navigation. Alloy accepted and exported every application span
without refusing one.

#### Why this came sixteenth

Metrics already explain aggregate rates, failures, latency, and resource usage,
while Loki makes individual completion events searchable. Tempo and Alloy were
then installed and verified independently with synthetic spans. Instrumenting
the existing request path now adds the missing per-request timing structure
without combining application changes with an untested tracing backend.

The image-processing stages are already meaningful units of CPU and latency, so
they provide a useful first trace without inventing another service or tracing
every helper function.

#### What this enables

- Seeing which processing stage dominates one slow resize request
- Finding an individual Tempo trace from its correlated Loki completion log
- Comparing successful, rejected, and overloaded request traces
- Propagating the same trace through the traffic generator later
- Navigating bidirectionally between Tempo traces and Loki completion logs
- Learning sampling, span metrics, and service graphs from real application data

### 17. Metrics-to-traces Grafana dashboard

**Status:** Implemented, published, and validated with generated traffic

The `Image Resizer Tracing` dashboard provides a guided transition from
aggregate Prometheus signals to individual Tempo traces. Its metric panels show
request totals, success percentage, unsuccessful requests, request rate by
outcome, request latency percentiles, and per-stage p95 latency. Its Tempo
tables show recent resize traces, traces exceeding a selectable slow threshold,
and rejected or failed traces.

Each Tempo result retains Grafana's data-source-provided Trace ID link. Opening
one shows the complete span waterfall and exposes the existing trace-to-Loki
navigation. Links in the dashboard header also connect to the dedicated metrics
overview and request-log dashboards.

The dashboard uses the application's existing Prometheus histograms for trends
and Tempo search for individual traces. It does not require Tempo's
metrics-generator, which remains disabled to keep the homelab stack lightweight
and avoid duplicating metrics already emitted by the application.

Live traffic validation returned fifty recent traces, thirty traces slower than
the default `500ms` threshold, and ten rejected or failed traces in the selected
hour. During validation, Prometheus observed about two successful requests per
second and returned processing-stage latency series for decode, resize, and
encode.

#### Why this came seventeenth

Instrumentation and trace/log correlation were working, but using only Explore
still required knowing what to query and where to start. The dashboard turns
the existing telemetry into a repeatable investigation path without adding new
collection components.

Placing aggregate metrics above trace results also makes the different roles
clear: metrics reveal a pattern across many requests, while a trace explains
the work performed by one request.

#### What this enables

- Learning the relationship between metrics, traces, and logs in one workflow
- Moving from a latency spike to examples of slow request traces
- Comparing trace waterfalls with aggregate processing-stage latency
- Finding rejected and failed traces without writing TraceQL manually
- Adjusting the definition of a slow trace from a dashboard variable
- Reusing the version-controlled dashboard after Grafana reinstalls

## Candidate next steps

These are possible follow-up features, not yet an agreed sequence. When one is
selected, it should become the next numbered entry with an explanation of why it
was prioritized over the other candidates.

- Service-level indicators based on user-visible success and latency
- An initial service-level objective
- Alerts and error-budget exercises
- Asynchronous resizing with a queue and worker
- Security and software supply-chain exercises

The sequence should be decided one step at a time. A common dependency pattern is
that an SLI requires a reliable measurement source, and an SLO requires a clearly
defined SLI. That does not by itself determine whether metrics, tracing,
deployment work, or another operational concern should be the immediate next
feature.
