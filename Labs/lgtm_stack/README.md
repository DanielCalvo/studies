
minikube start \
--profile=lgtm-lab \
--nodes=2 \
--memory=4096 \
--driver=docker

Grafana LGTM: Loki, Grafana, Tempo, Mimir


## loki
https://grafana.com/docs/loki/latest/get-started/
log aggregation system

## alloy
https://grafana.com/docs/alloy/latest/
looks like this is the thing that ships the logs, but also ships telemetry data (metrics, logs, traces, and profiles)

## mimir
https://grafana.com/docs/mimir/latest/?pg=oss-mimir&plcmt=resources
provides metrics storage for prom and otel

## grafana
famous graphing software you used many times: https://grafana.com/docs/grafana/latest/
I wonder how you use it with logs though!