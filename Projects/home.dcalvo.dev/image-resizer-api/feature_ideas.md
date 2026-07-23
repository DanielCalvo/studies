# Image Resizer API Feature Ideas

This project explores modern application, observability, Kubernetes, and SRE
practices through a small image-resizing service. The business operation is easy
to understand while still producing meaningful latency, CPU, memory, failure,
and workload signals.

This file describes the current state and the possible work ahead. It is not a
commitment to implement everything listed here. `feature_list.md` remains the
chronological record of features actually selected and implemented.

## Current state

The project is now a useful deployed observability lab, although it is not a
production-ready service.

### Application

- A synchronous Go endpoint accepts one JPEG and returns a smaller JPEG.
- Resizing preserves aspect ratio and does not support upscaling.
- Upload bytes, decoded pixels, and requested output width are bounded.
- Invalid input produces stable JSON error responses.
- Unit tests cover successful resizing and important rejection paths.

### Operational behavior

- Structured JSON request logs include request IDs, outcomes, duration, image
  dimensions, byte sizes, and stable error codes.
- Prometheus metrics cover requests, outcomes, latency, in-flight work,
  processing stages, image sizes, pixel counts, rejections, and Go runtime
  behavior.
- Liveness, readiness, graceful shutdown, and bounded request draining provide a
  Kubernetes lifecycle contract.

### Deployment and validation

- A static ARM64 binary runs in a small non-root container.
- Kubernetes runs two replicas across the Orange Pi nodes behind a MetalLB
  address.
- The ad hoc build-and-deploy script creates timestamped images, pushes them to
  the homelab registry, applies the manifests, and checks the deployed endpoints.
- Post-deployment smoke tests cover health, successful resizing, malformed and
  unsupported input, resource limits, and continued health after bad requests.

### Monitoring and experimentation

- A `ServiceMonitor` makes both replicas active Prometheus scrape targets.
- A traffic generator continuously cycles through ten generated JPEG dimensions
  in a shuffled order.
- The Grafana overview shows health, request totals and rate, outcomes, success
  percentage, latency percentiles, processing stages, rejection reasons, and
  per-pod CPU and memory.

## Most important production gaps

The following capabilities would generally matter more for a serious production
service than adding additional product features.

### 1. Define service-level indicators

Define user-visible measurements before choosing alert thresholds.

Candidate availability SLI:

```text
successful eligible resize requests / eligible resize requests
```

Candidate latency SLI:

```text
eligible resize requests completed within the chosen threshold
---------------------------------------------------------------
                    eligible resize requests
```

Candidate correctness SLI:

```text
synthetic resizes producing a valid result with expected dimensions
-------------------------------------------------------------------
                         synthetic attempts
```

Decisions still required:

- Which requests are eligible for the availability and latency SLIs?
- Should correctly rejected client input be excluded from the denominator?
- Should one latency threshold cover all supported image sizes?
- Would defined workload classes provide a fairer latency contract?
- What rolling measurement window should be used?

### 2. Establish an initial SLO

Once the SLIs are precise, choose modest initial objectives based on observed
behavior rather than an arbitrary industry number.

An example shape, not yet an agreed target, is:

> Over a rolling measurement window, X% of eligible JPEG resize requests succeed
> and Y% complete within Z seconds.

Then calculate error-budget consumption and use the results to decide when
reliability work should take priority over feature work.

### 3. Add concurrency limits and overload behavior

The service currently performs expensive work synchronously without an explicit
application-level concurrency budget.

Possible exercise:

- Add a small semaphore around decode, resize, and encode work.
- Select the initial limit from measurements on the Orange Pis.
- Return a deliberate `429` or `503` when capacity is unavailable.
- Define whether callers should retry and, if so, expose `Retry-After`.
- Add bounded metrics for admitted, active, waiting, and rejected work.
- Confirm that readiness remains meaningful during heavy traffic.

The goal is predictable degradation instead of allowing memory or CPU exhaustion
to decide the failure mode.

### 4. Perform controlled concurrent load testing

The existing traffic generator produces background activity; it is not a load
test.

A real load-testing exercise should:

- Generate configurable concurrency and request rates.
- Use a documented distribution of image sizes.
- Distinguish offered load from completed throughput.
- Record success, rejection, failure, and latency distributions.
- Increase load until latency, errors, or resource saturation become visible.
- Compare one replica with two replicas.
- Repeat tests after concurrency and resource changes.

This provides evidence for concurrency limits, Kubernetes requests and limits,
and replica counts.

### 5. Add alerts and synthetic monitoring

The service is observable, but no operator is notified when behavior degrades.

Potential alerts:

- Rapid availability or latency error-budget burn
- No healthy Image Resizer scrape targets
- Repeated pod restarts or OOM kills affecting requests
- Sustained overload rejection
- Missing application metrics

Avoid paging directly on generic CPU or memory thresholds unless they represent
imminent user impact. Resource signals are often better used to explain a
user-visible symptom.

The existing smoke-test journey can inspire a scheduled black-box check that:

1. Uploads a known valid JPEG.
2. Expects a successful response within a time limit.
3. Verifies the output is a valid JPEG with the expected dimensions.
4. Exposes success, correctness, and duration as metrics.

### 6. Improve production delivery

The current shell workflow is appropriate for this homelab but not a complete
production delivery process.

Possible increments:

- Run unit tests before building.
- Validate manifests before applying them.
- Record deployments as Grafana annotations.
- Add a staging or canary step.
- Run smoke tests automatically after rollout.
- Define rollback criteria and a rollback procedure.
- Promote immutable artifacts rather than rebuilding per environment.
- Trace deployed images back to source and build metadata.

### 7. Add security and supply-chain controls

Existing protections include strict input limits, malformed-input tests, a
non-root container, a read-only filesystem, and reduced container privileges.

Further exercises:

- Add authentication and authorization if the service leaves the trusted LAN.
- Add TLS at an ingress or gateway.
- Add rate limiting and abuse controls.
- Expand malformed-image and decompression-bomb testing.
- Pin and scan application and container dependencies.
- Generate an SBOM.
- Record build provenance.
- Sign and verify images.

## Diagnostic and observability extensions

These would deepen investigation capabilities but are not substitutes for SLIs,
overload protection, or alerting.

### Centralized logs

The application produces useful structured logs, but nothing currently collects
or stores them.

Possible exercise:

- Install a lightweight log collector and Loki.
- Preserve the existing structured fields.
- Query logs by request ID, outcome, error code, dimensions, and pod.
- Link Grafana metric panels to relevant logs.
- Define a sensible homelab retention period.

### OpenTelemetry tracing

Create a request span and child spans around the application stages:

```text
POST /v1/resize
|-- parse and validate
|-- decode JPEG
|-- resize pixels
`-- encode JPEG
```

Useful bounded span attributes could include formats, dimensions, processing
outcome, and stable error code. Image content should never be recorded.

The first trace remains local to one process, but it teaches trace structure and
log correlation. Cross-process context propagation becomes more meaningful if a
queue and worker are introduced later.

### Profiling

When metrics show CPU or memory pressure, use profiling to determine which code
is responsible.

Possible exercises include:

- Capture CPU profiles during small- and large-image traffic.
- Compare decode, scaling, and encoding costs.
- Inspect heap allocation during concurrent work.
- Compare profiling evidence with stage-duration metrics.

## Product and architectural extensions

These are useful learning exercises but would substantially expand the service's
operational surface.

### Additional image formats and operations

- Accept PNG and WebP input.
- Select JPEG, PNG, or WebP output.
- Support height, fit, crop, and quality options.
- Define clear behavior for animated or metadata-heavy images.
- Preserve or remove metadata deliberately.

Each format and option increases the validation, security, test, and performance
matrix.

### Asynchronous processing

For jobs that should not hold an HTTP connection open:

```text
API -> job queue -> resize worker -> object storage
```

This enables study of:

- Trace propagation across processes
- Idempotency keys and job deduplication
- Bounded retries with backoff and jitter
- Queue depth, oldest-job age, and processing metrics
- Backpressure and load shedding
- Dead-letter handling
- Graceful worker shutdown
- Partial failure and recovery

The synchronous service should be understood and protected from overload before
adding this architecture.

## Revised production-oriented sequence

If production readiness were the main objective, a reasonable next sequence from
the current state would be:

1. Define availability, latency, and correctness SLIs.
2. Establish a modest initial SLO.
3. Add concurrency limits and explicit overload behavior.
4. Perform controlled concurrent load testing.
5. Tune concurrency, Kubernetes resources, and replica count from measurements.
6. Add SLO and critical platform alerts.
7. Schedule an external synthetic resize check.
8. Add deployment annotations and centralized logs.
9. Add tracing or profiling when a concrete diagnostic question justifies it.
10. Add automated delivery and supply-chain controls.
11. Consider asynchronous processing only if synchronous behavior proves
    insufficient.

This ordering is not mandatory. It makes the dependency chain visible: reliable
alerts require meaningful objectives; overload limits benefit from measurements;
and distributed architecture should be introduced only when its additional
operational cost solves a demonstrated problem.

## Learning-oriented alternatives

Because this remains a study project, the next feature may deliberately differ
from the production sequence. Productive alternatives include:

- Add tracing to learn OpenTelemetry before defining an SLO.
- Centralize logs to explore metric-to-log correlation.
- Use profiling to understand why large images are expensive.
- Add one new image format to study validation and metric evolution.
- Introduce a queue specifically to study asynchronous observability.

When choosing one of these, record that it was selected for learning value rather
than presenting it as the highest production priority.
