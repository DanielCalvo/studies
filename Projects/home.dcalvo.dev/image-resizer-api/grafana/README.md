# Grafana dashboards

This directory contains API payloads for dashboards published to the homelab
Grafana instance at `http://192.168.1.221`.

`image-resizer-overview.json` uses the Prometheus datasource with UID
`prometheus` and has a
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

`image-resizer-logs.json` uses the Loki datasource with UID `loki` and has the
stable dashboard UID `image-resizer-logs`. It parses the application's JSON at
query time and provides:

- Request totals, success percentage, client rejections, and server failures
- Request counts by HTTP status, error reason, and pod
- Log-derived successful-request latency and image payload sizes
- A recent unsuccessful-request log view
- A complete recent request-log view

Its default time range is one hour so recent overload and client-error
experiments remain visible. The published dashboard is available at:

```text
http://192.168.1.221/d/image-resizer-logs/image-resizer-logs
```

`image-resizer-tracing.json` connects the application's Prometheus metrics to
individual request traces in Tempo. It provides:

- Request rate by outcome and request latency percentiles
- Processing-stage p95 latency for decode, resize, and encode
- Recent clickable request traces
- Separate views of slow traces and rejected or failed traces

Click a trace ID to open its span waterfall. From a selected span, Grafana's
provisioned Tempo-to-Loki correlation can open the matching request log.

The time-series panels intentionally use the application's existing Prometheus
histograms. Tempo's metrics-generator is disabled in this small homelab, while
Tempo search supplies the trace tables.

The published dashboard is available at:

```text
http://192.168.1.221/d/image-resizer-tracing/image-resizer-tracing
```

Publish any version-controlled payload with:

```bash
curl --fail-with-body --silent --show-error \
  --user admin:admin \
  --header 'Content-Type: application/json' \
  --data-binary @image-resizer-overview.json \
  http://192.168.1.221/api/dashboards/db

curl --fail-with-body --silent --show-error \
  --user admin:admin \
  --header 'Content-Type: application/json' \
  --data-binary @image-resizer-logs.json \
  http://192.168.1.221/api/dashboards/db

curl --fail-with-body --silent --show-error \
  --user admin:admin \
  --header 'Content-Type: application/json' \
  --data-binary @image-resizer-tracing.json \
  http://192.168.1.221/api/dashboards/db
```
