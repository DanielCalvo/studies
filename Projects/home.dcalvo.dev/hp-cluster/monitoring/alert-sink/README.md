# Alert sink

This disposable homelab service accepts Alertmanager-compatible webhook
requests at `POST /alerts`, writes their JSON bodies to standard output, and
returns HTTP 200. It is private behind a `ClusterIP` Service. The image is
published to the OPI plain-HTTP registry, so HP RKE2 nodes must be configured
to allow `192.168.1.225:5000` as an insecure registry before deployment.

Open the in-memory alert table through a temporary port-forward:

```bash
kubectl --context hp -n monitoring port-forward service/alert-sink 8080:8080
```

Then visit `http://127.0.0.1:8080/`. The table is keyed by Alertmanager
fingerprint, so a resolved webhook updates the existing firing row. The
`/api/alerts` endpoint exposes the same state as JSON. State is intentionally
ephemeral and is lost when the pod restarts.

Build and deploy the `linux/amd64` image to the local registry with:

```bash
./build-and-deploy.sh
```

The script always verifies and uses Kubernetes context `hp`. Inspect received
notifications with:

```bash
kubectl --context hp -n monitoring logs deployment/alert-sink
```
