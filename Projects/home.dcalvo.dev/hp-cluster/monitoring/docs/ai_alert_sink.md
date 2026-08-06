# HP Alertmanager alert-sink exercise

Implemented and verified on 2026-08-05:

- `monitoring/alert-sink/` contains a disposable Go HTTP receiver. It accepts
  JSON `POST /alerts`, logs the payload to stdout, and exposes `/healthz`.
- The image is built for `linux/amd64` and pushed to the OPI plain-HTTP local
  registry at `192.168.1.225:5000/alert-sink`. HP nodes must list that registry
  as insecure in their RKE2 registry settings before pulling it.
- `monitoring/alertmanager/` creates one operator-managed Alertmanager and
  routes alerts labeled `environment="homelab"` to the private ClusterIP sink.
  Webhook delivery has `send_resolved: true`.
- `prometheus-instance.yaml` selects `PrometheusRule` objects labeled
  `prometheus: hp` and sends notifications to `alertmanager-operated`.
- `image-resizer-demo-rule.yaml` is intentionally inactive (`vector(0) > 0`).
  For a test, temporarily change it to `vector(1)`, apply it, wait for its
  one-minute `for` period, inspect sink logs, then restore `vector(0) > 0`.

The test delivered both firing and resolved Alertmanager webhook payloads to
the sink. HP Alloy now tails only this pod's Kubernetes logs into Loki, and the
Grafana dashboard is available at
`http://192.168.1.231/d/alert-sink-logs/alert-sink-logs`. The sink is private
and has no external integrations.

The sink also serves a minimal in-memory dashboard at `/` and JSON state at
`/api/alerts`. Rows are keyed by Alertmanager fingerprint; resolved webhooks
change the matching row from `firing` to `resolved`. No PVC or other persistent
storage is used, so restarting the pod clears the table.

Karma `v0.131` reads the same Alertmanager API at
`http://192.168.1.234`. It is currently showing zero alert groups because the
temporary demo rule is inactive.
