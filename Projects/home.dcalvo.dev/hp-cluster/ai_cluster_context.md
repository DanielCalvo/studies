# HP Homelab Cluster Context

This is the x86_64 home-lab Kubernetes cluster used for study, experiments,
and the central observability stack. Live state was last verified on
2026-08-04.

## Cluster

- Distribution: RKE2 `v1.35.6+rke2r1`
- API server: `https://192.168.1.211:6443`
- Architecture: `linux/amd64`
- Network: trusted local home LAN
- CNI: RKE2 Canal (Calico and Flannel)
- Ingress: RKE2 ingress-nginx using host ports 80 and 443
- Merged kubeconfig: `/home/daniel/.kube/config`
- Source kubeconfig: `/home/daniel/.kube/config-hp` (its internal records are
  named `default`; the merged file renames the cluster, user, and context)
- Required context: `hp`

Agents must still pass the kubeconfig and `--context hp` (or Helm's
`--kube-context hp`) explicitly on every cluster command.

## Nodes

- `hp1`: `192.168.1.211`, control-plane and worker
- `hp2`: `192.168.1.212`, worker
- The cluster was reinstalled on 2026-08-04. Both nodes were then Ready on
  Ubuntu 26.04 LTS with kernel `7.0.0-28-generic` and containerd
  `2.2.5-k3s2`.
- Each node is an HP ProDesk with an Intel Core i7-4790S, 16 GiB RAM, and a
  250 GB SSD.
- Non-root SSH access is available as `daniel@192.168.1.211` and
  `daniel@192.168.1.212`. SSH inspection is read-only unless the user
  explicitly authorizes a node change.

## Networking

- MetalLB is installed from Helm chart `0.16.1`.
- Its layer-2 pool is `192.168.1.230-192.168.1.239`.
- RKE2 ingress-nginx currently uses `192.168.1.239`.
- Monitoring reserves `192.168.1.231` for Grafana and `192.168.1.233` for
  Prometheus.
- These services use plain HTTP on the trusted LAN. Do not expose them through
  Internet port forwarding without a separate authentication and TLS design.

## Storage

- Rancher Local Path Provisioner `v0.0.36` is installed in the
  `local-path-storage` namespace.
- Its `local-path` StorageClass uses `WaitForFirstConsumer`, a `Delete` reclaim
  policy, and `/opt/local-path-provisioner` as the node storage root.
- The StorageClass is not the default; workloads must set
  `storageClassName: local-path` explicitly.
- Local Path Provisioner storage is node-local and not replicated. A workload
  using a bound volume cannot fail over to the other node while retaining that
  data.
- Requested PVC capacity is not enforced by Local Path Provisioner. Workloads
  such as Prometheus should configure their own retention size to protect the
  nodes' root disks.
