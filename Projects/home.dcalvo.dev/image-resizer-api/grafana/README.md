# Grafana dashboards

`image-resizer-overview.json` is the API payload for the dashboard published to
the homelab Grafana instance at `http://192.168.1.221`.

The dashboard uses the Prometheus datasource with UID `prometheus` and has a
stable dashboard UID, `image-resizer-overview`. Publishing the file again updates
the existing dashboard instead of creating a duplicate.

The top row shows healthy replicas, total, successful, and failed request counts,
plus the successful-request percentage for the time range selected in Grafana.
A requests-per-minute graph shows how traffic changes over time. A failed request
means an application `failed` outcome; rejected client input is reported
separately.

The published dashboard is available at:

```text
http://192.168.1.221/d/image-resizer-overview/image-resizer-overview
```
