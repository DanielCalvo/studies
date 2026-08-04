# Production Monitoring and Observability Signals

The four golden signals are an excellent foundation, but they do not tell the
whole story. They primarily describe whether a request-oriented service appears
healthy. A production monitoring strategy should also cover user-visible
correctness, dependencies, resources, changes, and diagnostic evidence.

## An outside-in monitoring model

| Layer | Question | Useful signals |
| --- | --- | --- |
| User experience | Can users accomplish what they came to do? | Availability, correctness, latency, synthetic tests |
| Service behavior | How is the application processing work? | Rate, errors, duration, saturation |
| Resources | Is it approaching a physical or configured limit? | CPU, memory, queues, throttling, OOMs |
| Dependencies | Is something downstream causing the problem? | Dependency latency, errors, availability |
| Changes | What changed when behavior changed? | Deployments, configuration, feature flags, scaling |
| Diagnosis | Why did this particular operation fail? | Logs, traces, profiles, request IDs |

Google describes this as the distinction between monitoring **symptoms**—what
is broken—and **causes**—why it is broken. Both matter, but user-visible symptoms
should normally drive urgent alerts. See
[Monitoring Distributed Systems](https://sre.google/sre-book/monitoring-distributed-systems/).

## 1. Start with user-visible outcomes

Identify the small number of operations users actually care about, then measure:

- Did the operation complete successfully?
- Was the result correct?
- Did it finish within an acceptable time?
- Was the service reachable from where users access it?

For the Image Resizer API, the primary user journey is:

```text
Submit a valid supported JPEG -> receive a valid, correctly sized JPEG
```

Important signals include:

- Percentage of eligible requests that succeed
- Percentage that complete below a latency threshold
- Correct output dimensions and aspect ratio
- Valid JPEG output
- External endpoint availability

Correctness deserves special attention. An API could return `200 OK` quickly
while producing corrupt or incorrectly sized images. Ordinary HTTP metrics would
call that healthy. A synthetic test that periodically uploads a known image and
validates the result would catch it.

Google's SRE guidance treats generated golden data and comparison with expected
output as a useful way to monitor correctness when ordinary metrics cannot
establish it. See
[Data Processing](https://sre.google/workbook/data-processing/).

## 2. Use RED or the golden signals for service behavior

For request-oriented services, RED is an easily remembered subset:

- **Rate:** how much work is arriving?
- **Errors:** how much work is failing?
- **Duration:** how long does it take?

The golden signals add **saturation**: how close is the service to being unable
to accept more work? See
[The RED Method](https://grafana.com/blog/the-red-method-how-to-instrument-your-services/).

For the Image Resizer API, this includes:

- Requests per minute
- Success, rejection, and failure rates
- p50, p95, and p99 latency
- Decode, resize, and encode latency
- Active resize operations
- Requests in flight

Avoid relying on average latency. A service with nine fast requests and one
20-second request can have an acceptable-looking average while still providing a
terrible experience to some users. Histograms, percentiles, and the percentage
completed under a threshold are usually more informative.

Also separate different outcomes:

- `failed`: the service malfunctioned
- `rejected`: the service correctly refused invalid input
- `succeeded`: the requested operation completed

A corrupt JPEG correctly returning `415` should generally not consume an
availability error budget. It may still be operationally interesting if
rejection volume suddenly increases.

## 3. Monitor workload characteristics

Two periods with the same request rate can impose radically different loads.

For an image service, useful workload signals include:

- Encoded input and output bytes
- Input and output pixel counts
- Requested output dimensions
- Distribution of small, medium, and large inputs
- Compression or expansion trends
- Request mix by operation, if more operations are added later

This matters when interpreting latency. A p95 increase may indicate a regression,
or it may simply mean the service received more 25-megapixel images.

For a production latency SLI, we might eventually:

- Set one user-visible latency objective regardless of size, or
- Define supported workload classes with different expectations

That decision should be explicit. Otherwise, a changing traffic mix can make the
service look better or worse without any code change.

## 4. Apply USE to resources

The USE method recommends checking, for every relevant resource:

- **Utilization:** how much is being used?
- **Saturation:** is work waiting because the resource is full?
- **Errors:** is the resource failing?

It is particularly useful for diagnosing bottlenecks. See
[The USE Method](https://www.brendangregg.com/usemethod.html).

For the Image Resizer API:

- CPU usage and CPU throttling
- Memory usage relative to the container limit
- OOM kills and pod restarts
- Node CPU and memory pressure
- Requests waiting for processing capacity
- Active versus permitted resize operations
- Go heap size, garbage collection, and goroutines
- Network errors and bandwidth, if they become relevant

Saturation is often more useful than utilization. A CPU can appear only
moderately utilized over five minutes while short bursts create queues and poor
latency.

The current service does not have a concurrency limit or queue, so it lacks a
strong application-level saturation signal. `requests_in_flight` and
`resize_operations_in_flight` help, but they do not say how much work is waiting.
That becomes important if a concurrency semaphore or asynchronous queue is added.

## 5. Monitor dependencies separately

For every dependency, ask:

- Is it reachable?
- How often does it fail?
- How long do calls take?
- Is it saturated?
- Are retries hiding failures and adding latency?

Examples include databases, caches, object stores, queues, DNS, authentication
services, and third-party APIs.

The current Image Resizer API has few runtime dependencies because processing
happens locally in memory. Kubernetes, the node, networking, and the LoadBalancer
still form part of the delivery path.

## 6. Record changes as operational signals

Track changes such as:

- Application deployments and image versions
- Configuration changes
- Resource-limit changes
- Replica-count changes
- Feature-flag changes
- Node maintenance and pod movement

When latency rises at exactly the time a deployment occurred, that does not prove
causation, but it gives an investigator a strong place to begin. Useful
production dashboards normally show deployment annotations alongside traffic,
errors, and latency.

## 7. Use the right telemetry for each question

Metrics are not supposed to answer everything. A productive division is:

- **Metrics:** Is something happening at scale?
- **Logs:** What happened during a particular event?
- **Traces:** Where was time spent across stages or services?
- **Profiles:** Which code consumed CPU or allocated memory?
- **Synthetic tests:** Does the complete user journey actually work?
- **Deployment events:** What changed around that time?

Google's monitoring guidance includes metrics, structured logs, tracing, and
event introspection as different forms of monitoring data. See
[Monitoring Systems with Advanced Analytics](https://sre.google/workbook/monitoring/).

For the Image Resizer API, an investigation might proceed as follows:

1. A latency panel shows that requests became slow.
2. Processing-stage metrics show that resizing became slow rather than decoding.
3. A trace identifies the affected request and its stages.
4. A CPU profile identifies the functions consuming processing time.
5. Structured logs supply the request ID, dimensions, status, and error code.

Each signal narrows the investigation.

## 8. Turn important signals into SLIs and SLOs

A dashboard tells you what happened. An SLO tells you whether it matters enough
to require action.

Possible Image Resizer SLIs are:

```text
Availability =
successful eligible resize requests / eligible resize requests
```

```text
Latency =
eligible requests completed within 5 seconds / eligible requests
```

```text
Correctness =
synthetic resizes producing the expected valid output / synthetic attempts
```

After choosing objectives, alert on meaningful error-budget consumption rather
than every momentary threshold crossing. Google recommends burn-rate alerting
because it relates the alert to how quickly the service is consuming its
reliability allowance. See
[Alerting on SLOs](https://sre.google/workbook/alerting-on-slos/).

## 9. Separate alerts from dashboards

Not everything worth graphing is worth waking someone up for.

- **Page:** users are being affected now, or error budget is burning rapidly.
- **Ticket or message:** a trend needs attention but is not urgent.
- **Dashboard:** useful during investigation or planning.
- **Log only:** retained as diagnostic evidence.

Examples:

- High SLO burn rate: page
- Repeated OOM kills causing failed requests: page
- Memory growing gradually over several days: ticket
- CPU briefly reaching 80% while latency remains healthy: dashboard
- One invalid JPEG: structured log and rejection counter

CPU and memory are usually cause signals, not direct reasons to page. A
user-impact alert accompanied by high CPU is much more meaningful than a generic
"CPU above 80%" alert.

## Practical rule of thumb

For most production services, aim for:

1. One or two clearly defined user journeys.
2. Availability, correctness, and latency SLIs for those journeys.
3. RED or golden-signal metrics for every major service boundary.
4. USE signals for every resource that can become a bottleneck.
5. Health and latency signals for every important dependency.
6. Deployment and configuration-change annotations.
7. Structured logs with correlation identifiers.
8. Traces for important or distributed operations.
9. Synthetic tests from outside the service.
10. Alerts tied primarily to user impact and error-budget burn.
11. Monitoring of the monitoring system itself, including scrape failures,
    missing data, and notification delivery.

The key principle is:

> Monitor user-visible outcomes first, then collect enough internal evidence to
> explain those outcomes and enough capacity information to anticipate the next
> failure.
