# Image Resizer SLO alerting follow-up

This note captures the next observability exercise for the Image Resizer API.

## Current SLI/SLO state

The canonical dashboard is maintained in the HP cluster tree:

`hp-cluster/monitoring/grafana/dashboards/image-resizer/image-resizer-overview.json`

The dashboard currently includes:

- Availability SLI over the selected Grafana range
- Latency SLI for eligible requests completed within five seconds
- Composite SLO attainment over the selected Grafana range
- Error budget remaining over the selected Grafana range

The preliminary objective is:

> At least 99.9% of eligible requests should return the correct resized image
> within five seconds over a rolling 30-day window.

The current composite measurement treats a successful request completed within
five seconds as a good event. Rejected client input is excluded from the
eligible-request denominator. The measurement does not yet independently verify
that the returned image is correct.

## Next exercise: local alert sink

Build a small in-cluster `alert-sink` HTTP service for harmless alert testing.
Although the Image Resizer API runs on OPI, the Prometheus instance that stores
and evaluates its metrics runs on HP. Keep the sink, Alertmanager resources,
and alert rules in `hp-cluster/monitoring/` and target that cluster explicitly
with `--context hp`.
It should:

- Listen on a simple HTTP port.
- Accept `POST /alerts`.
- Log the received JSON payload to standard output.
- Return a successful HTTP response.
- Avoid external integrations such as Slack, email, PagerDuty, or Webhook.site.
- Run behind a private `ClusterIP` Service; it does not need a LoadBalancer.
- Use a small image that supports the HP cluster's `linux/amd64` architecture.

Use it as the receiver for both Prometheus Alertmanager and, if useful later,
Grafana Webhook contact points. The Alertmanager flow is:

```text
Prometheus rule -> Prometheus -> Alertmanager -> alert-sink -> pod logs
```

The Grafana flow is:

```text
Grafana alert rule -> Grafana contact point -> alert-sink -> pod logs
```

Keep the sink intentionally disposable and clearly labeled as a homelab test
component. A small Go service or a simple existing echo-server image is enough.

HP Prometheus does not yet have an active Alertmanager configuration: the
`alerting` section in `hp-cluster/monitoring/prometheus_operator/prometheus-instance.yaml`
is currently commented out. Before adding the demo rule, deploy or configure
Alertmanager on HP, configure Prometheus to discover it, and configure an
Alertmanager route to the sink.

## Deterministic dummy alert

Add a temporary Prometheus rule such as:

```yaml
groups:
  - name: image-resizer-demo
    rules:
      - alert: ImageResizerDemoAlert
        expr: vector(1)
        for: 1m
        labels:
          severity: info
          environment: homelab
        annotations:
          summary: Image Resizer demo alert
          description: This alert intentionally fires to test Prometheus and Alertmanager.
```

Route `environment: homelab` to the local sink and enable `send_resolved` for
that receiver. Verify the firing notification, inspect the sink logs, then make
the rule resolve by changing `vector(1)` to `vector(0)` or by removing the
temporary rule. Confirm that the resolved notification is also delivered.

Do not leave the always-firing rule enabled permanently. A better permanent
smoke test, if desired, is a manually controlled metric or a short-lived test
rule that is enabled only during alerting exercises. Keep this disposable demo
rule in a file separate from future Image Resizer SLO and burn-rate alert rules
so cleanup is unambiguous.

## Remaining SLI/SLO topics

1. **Independent correctness measurement**

   Add a synthetic checker with known valid JPEG inputs and expected output
   properties. Measure whether the returned image is valid and correctly resized
   rather than treating HTTP 200 alone as correctness.

2. **Finalize eligible-request semantics**

   Document exactly which requests count toward the SLO. Current intent is to
   exclude rejected client input and include valid requests that fail in the
   service or exceed the latency objective.

3. **Use a formal rolling 30-day SLO measurement**

   The dashboard panels intentionally follow the selected Grafana range. A
   30-day selection represents the stated objective, while shorter selections
   are useful diagnostic views. Later, consider recording a dedicated 30-day
   recording rule if an always-available formal SLO value is needed.

4. **Error-budget policy**

   The 99.9% objective provides a 0.1% error budget. Decide what actions should
   happen when the budget is partly or fully consumed, such as investigation,
   pausing feature work, or running a reliability exercise.

5. **Burn-rate alerting**

   Add Prometheus/Alertmanager rules for rapid budget consumption after the
   basic sink works. Start with informational homelab alerts and short windows;
   do not page anyone.

6. **No-traffic behavior**

   Decide whether SLI and error-budget panels should show `No data` when there
   are no eligible requests, rather than displaying a misleading percentage.

7. **Alert runbooks and annotations**

   Add concise summaries, descriptions, dashboard links, and a small runbook
   for investigating availability, latency, correctness, and error-budget
   alerts.

## Useful implementation constraints

- The canonical dashboard is in the `hp-cluster` tree and uses Grafana at
  `http://192.168.1.231`.
- Kubernetes commands must explicitly use `--context hp`.
- Verify the HP API endpoint and node identities before state-changing cluster
  commands.
- Keep the local sink private to the trusted home network.
- Preserve durable notes in `docs/` with the `ai_` filename prefix.
