# Local Path Provisioner

RKE2 does not bundle a local path provisioner by default, so it is installed as an add-on.

## AI notes on Installed configuration

- Provisioner: Rancher Local Path Provisioner `v0.0.36`
- Namespace: `local-path-storage`
- StorageClass: `local-path`
- Volume binding: `WaitForFirstConsumer`
- Reclaim policy: `Delete`
- Node storage root: `/opt/local-path-provisioner`

The StorageClass is deliberately not marked as the cluster default. Workloads must select it explicitly with `storageClassName: local-path`.

Local Path Provisioner does not enforce the requested volume capacity. Stateful applications must limit their own disk usage. Local volumes remain tied to
the node on which they were provisioned and are persistent, but not replicated.

## Persistence smoke test

Run the test from any directory:

```bash
./hp-cluster/local-path-provisioner/smoke-test.sh
```
