# Karma Alertmanager dashboard

Karma is a read-only web dashboard for the HP Alertmanager API. It displays
active alerts and provides the Alertmanager silence workflow; it does not
evaluate Prometheus or Grafana rules itself.

The deployment is pinned to the official `ghcr.io/prymitive/karma:v0.131`
linux/amd64 image and reads the in-cluster HP Alertmanager service. It is
exposed privately on the trusted LAN at:

```text
http://192.168.1.234
```

On 2026-08-06 Karma successfully connected to Alertmanager `v0.33.1` and
collected its alerts and silences. It reported zero alert groups because the
temporary demo rule is currently inactive.

Apply the manifests with the explicit HP context:

```bash
kubectl --context hp apply -f deployment.yaml -f service.yaml
```
