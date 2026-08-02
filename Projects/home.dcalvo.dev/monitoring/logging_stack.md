# Kubernetes application logging stack

The first logging pipeline collects only the structured JSON logs written to
stdout by `image-resizer/image-resizer-api`.

```text
image-resizer pods -> Alloy -> Loki -> Grafana
```

All configuration is checked into this repository. Do not create data sources,
RBAC, or collector configuration manually in the cluster.

## Prerequisites

- A working k3s cluster and `kubectl` context
- Helm 3
- The `monitoring` namespace from `monitoring/namespace.yaml`
- The Grafana Helm repository:

  ```bash
  helm repo add grafana https://grafana.github.io/helm-charts
  helm repo update grafana
  ```

The selected Loki, Alloy, and Alloy config-reloader images all publish
`linux/arm64` variants.

## Recreate or reconcile

Apply the namespace, then reconcile the three pinned releases in dependency
order:

```bash
kubectl apply -f monitoring/namespace.yaml

helm upgrade --install loki grafana/loki \
  --version 7.0.0 \
  --namespace monitoring \
  --values monitoring/loki/values.yaml \
  --wait \
  --timeout 10m

helm upgrade --install alloy grafana/alloy \
  --version 1.10.0 \
  --namespace monitoring \
  --values monitoring/alloy/values.yaml \
  --wait \
  --timeout 10m

helm upgrade --install grafana grafana/grafana \
  --version 10.5.15 \
  --namespace monitoring \
  --values monitoring/grafana/values.yaml \
  --wait \
  --timeout 10m
```

The commands are idempotent: the same sequence installs a new stack or
reconciles an existing one.

## Explore the image-resizer logs

Open Grafana at `http://192.168.1.221`, select **Explore**, and choose the
provisioned Loki data source.

All application logs:

```logql
{namespace="image-resizer", app="image-resizer-api"}
```

Parse the JSON fields:

```logql
{namespace="image-resizer", app="image-resizer-api"} | json
```

Application errors:

```logql
{namespace="image-resizer", app="image-resizer-api"}
  | json
  | level="ERROR"
```

Requests slower than 500 milliseconds:

```logql
{namespace="image-resizer", app="image-resizer-api"}
  | json
  | duration_ms > 500
```

A particular request without indexing its high-cardinality ID:

```logql
{namespace="image-resizer", app="image-resizer-api"}
  | json
  | request_id="<request-id>"
```

## Current boundaries

- Loki is a single instance using a node-local PVC, not an HA service.
- Logs are retained for seven days.
- Loki and Alloy are internal `ClusterIP` services.
- Alloy currently collects only image-resizer application logs.
- Kubernetes Events and node journal logs are intentionally deferred.
