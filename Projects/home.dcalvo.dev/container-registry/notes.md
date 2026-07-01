# Local container registry

Unauthenticated HTTP Registry v2 for the home lab.

- Endpoint: `192.168.1.242:5000`
- Namespace: `container-registry`
- Service: MetalLB `LoadBalancer`
- Storage: `20Gi` PVC, default storage class

Use only on the trusted internal network.

## Deploy

```sh
kubectl apply -f .
kubectl -n container-registry rollout status deployment/registry
kubectl -n container-registry get svc registry
```

## Insecure Registry Config

Docker client, `/etc/docker/daemon.json`:

```json
{"insecure-registries":["192.168.1.242:5000"]}
```

```sh
sudo systemctl restart docker
```

k3s nodes, `/etc/rancher/k3s/registries.yaml`:

```yaml
mirrors:
  "192.168.1.242:5000":
    endpoint:
      - "http://192.168.1.242:5000"
```

```sh
sudo systemctl restart k3s        # server node
sudo systemctl restart k3s-agent  # worker nodes
```

## Push And Pull

Normal Docker push:

```sh
docker pull busybox:latest
docker tag busybox:latest 192.168.1.242:5000/busybox:latest
docker push 192.168.1.242:5000/busybox:latest
curl http://192.168.1.242:5000/v2/busybox/tags/list
```

For this arm64 k3s cluster, use Buildx to avoid pushing an amd64-only image from an amd64 workstation:

```sh
./container-registry/push-and-test-busybox.sh
kubectl logs busybox-registry-test
kubectl delete pod busybox-registry-test
```

Use registry images in manifests with the full endpoint:

```yaml
image: 192.168.1.242:5000/busybox:latest
```

## Delete Images

Delete a tag's manifest:

```sh
DIGEST=$(curl -sI \
  -H 'Accept: application/vnd.docker.distribution.manifest.v2+json' \
  http://192.168.1.242:5000/v2/busybox/manifests/latest \
  | awk -F': ' '/Docker-Content-Digest/ {print $2}' \
  | tr -d '\r')

curl -X DELETE "http://192.168.1.242:5000/v2/busybox/manifests/${DIGEST}"
curl http://192.168.1.242:5000/v2/busybox/tags/list
```

Reclaim disk space with garbage collection while the registry is stopped:

```sh
kubectl -n container-registry scale deployment/registry --replicas=0
kubectl -n container-registry wait --for=delete pod -l app=registry --timeout=120s
kubectl -n container-registry delete job registry-garbage-collect --ignore-not-found
kubectl apply -f container-registry/garbage-collect-job.yaml
kubectl -n container-registry wait --for=condition=complete job/registry-garbage-collect --timeout=300s
kubectl -n container-registry logs job/registry-garbage-collect
kubectl -n container-registry scale deployment/registry --replicas=1
kubectl -n container-registry rollout status deployment/registry
```
