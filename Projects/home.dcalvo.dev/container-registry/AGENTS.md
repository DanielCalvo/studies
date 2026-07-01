# Container Registry Context

This folder manages a local unauthenticated Docker Registry v2 instance for the home lab k3s cluster.

Key facts:
- Registry endpoint: `192.168.1.242:5000`
- Kubernetes namespace: `container-registry`
- Exposure: MetalLB `LoadBalancer` service
- Storage: `20Gi` PVC using the default `local-path` storage class
- Transport/auth: plain HTTP, no authentication; internal trusted network only
- Cluster nodes are arm64 Orange Pi machines:
  - `opi1`: `192.168.1.201`, k3s server
  - `opi2`: `192.168.1.202`, k3s agent

Files:
- `namespace.yaml`: namespace for the registry.
- `pvc.yaml`: persistent storage for `/var/lib/registry`.
- `deployment.yaml`: `registry:2` deployment with delete support enabled via `REGISTRY_STORAGE_DELETE_ENABLED=true`.
- `service.yaml`: MetalLB LoadBalancer pinned to `192.168.1.242`.
- `garbage-collect-job.yaml`: on-demand registry garbage collection job; run only while the registry deployment is scaled to zero.
- `push-and-test-busybox.sh`: pushes a `linux/arm64` BusyBox image with Docker Buildx and verifies Kubernetes can pull/run it.
- `notes.md`: concise user-facing deploy, client config, push/pull, and delete instructions.

Operational notes:
- Apply normal registry resources with `kubectl apply -f container-registry/`.
- The local Docker daemon must list `192.168.1.242:5000` in `insecure-registries`.
- Each k3s node must have `/etc/rancher/k3s/registries.yaml` pointing this registry to `http://192.168.1.242:5000`, followed by restarting `k3s` or `k3s-agent`.
- Use Buildx or another platform-aware tool when pushing images from an amd64 workstation for this arm64 cluster. Plain `docker tag && docker push` can accidentally push an amd64 image and cause `exec format error` in pods.
- A previous successful test pod logged:
  - `pulled 192.168.1.242:5000/busybox:latest`
  - `aarch64`

Be careful:
- Do not add authentication or TLS unless the user explicitly asks; this setup is intentionally simple for the internal network.
- Do not run garbage collection while the registry is accepting pushes.
- Keep manifests split by resource rather than recombining them.
