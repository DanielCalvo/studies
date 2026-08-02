# Tempo storage backends

For a small monolithic Tempo instance, a local directory or Kubernetes
PersistentVolumeClaim can be reasonable. For production, Tempo is designed
primarily around object storage rather than a normal shared filesystem volume.

## Supported trace-storage backends

| `storage.trace.backend` | Storage system | Typical use |
| --- | --- | --- |
| `local` | Local filesystem, `emptyDir`, or PersistentVolumeClaim | Development, testing, homelabs, and small monolithic deployments |
| `s3` | Amazon S3 or an S3-compatible object store | Common for production and self-hosted installations |
| `gcs` | Google Cloud Storage | Production on Google Cloud |
| `azure` | Azure Blob Storage | Production on Azure |

Grafana recommends object storage for production. The local backend is intended
for development and testing and is only supported in monolithic mode.

## Configuration shape

A basic S3 configuration resembles:

```yaml
storage:
  trace:
    backend: s3
    s3:
      bucket: tempo-traces
      endpoint: s3.amazonaws.com
```

The equivalent backend selection for Google Cloud Storage is:

```yaml
storage:
  trace:
    backend: gcs
    gcs:
      bucket_name: tempo-traces
```

For Azure Blob Storage, it begins with:

```yaml
storage:
  trace:
    backend: azure
    azure:
      container_name: tempo-traces
```

Real configurations also need the appropriate region, endpoint,
authentication, encryption, and provider-specific settings.

## Common practical choices

There is no neutral adoption measurement that ranks these choices precisely.
The common practical pattern is:

1. S3 or an S3-compatible API is probably the most widespread choice.
2. GCS is a natural choice for Google Cloud environments.
3. Azure Blob is a natural choice for Azure environments.
4. S3-compatible storage is common in self-hosted and on-premises
   environments.
5. Local or PVC storage is common in development and small homelabs, but not
   for highly available distributed Tempo.

Cloud deployments commonly follow their cloud provider:

```text
AWS cluster   -> Amazon S3
GCP cluster   -> Google Cloud Storage
Azure cluster -> Azure Blob Storage
```

An on-premises deployment commonly follows this pattern:

```text
Tempo -> S3-compatible object-storage service -> disks
```

Examples of systems that expose an S3-compatible API include MinIO and Ceph
Object Gateway. Compatibility and production support should be validated for
the particular product and version.

Grafana's current Tempo documentation presents SeaweedFS, `rclone`, and
community MinIO instructions as local evaluation examples rather than
production recommendations. MinIO is still supported through the S3-compatible
API, but its open source repository and binary distribution changed
significantly, so its current operating and support model should be evaluated
before choosing it for a new production installation.

## Why object storage instead of one shared volume?

Object storage provides:

- access from many Tempo replicas
- independent scaling of readers and writers
- durable and replicated storage services
- relatively inexpensive retention of large data volumes
- storage lifecycle and retention controls
- no requirement to mount one writable filesystem into every Tempo component

Tempo stores traces as immutable Parquet blocks. A simplified bucket layout is:

```text
bucket/
`-- tenant-id/
    `-- block-id/
        |-- meta.json
        |-- data.parquet
        |-- bloom-0
        `-- index
```

Block builders write these objects, queriers read them, and backend workers
compact or expire them.

## Storage progression for this exercise

The learning environment begins with temporary storage:

```text
Tempo -> emptyDir
```

The next local persistence improvement is:

```text
Tempo -> Minikube PersistentVolumeClaim
```

A more realistic self-hosted environment could use:

```text
Tempo -> external S3-compatible storage
```

A production microservices deployment generally uses:

```text
Tempo microservices -> Kafka + managed or production-grade object storage
```

A PVC backed by a small-board disk or NAS can be adequate for a personal
monolithic Tempo instance. The tradeoff is that Tempo's availability and trace
retention become tied to that node, filesystem, and backup process.

Directly mounting a shared network filesystem does not turn the `local` backend
into the recommended distributed production design. Distributed Tempo is
designed around a shared object-storage API.

## Kafka is not the trace backend

In Tempo 3 microservices mode, Kafka provides immediate durability between
trace ingestion and block creation. It holds data while downstream components
consume it and produce blocks.

Kafka does not replace long-term trace storage:

```text
Kafka
  -> durable transport and buffering for the write path

Object storage
  -> long-term retained trace blocks
```

Object storage remains the durable trace backend.

## Current references

- [Tempo object-storage architecture](https://grafana.com/docs/tempo/latest/reference-tempo-architecture/object-storage/)
- [Tempo storage configuration](https://grafana.com/docs/tempo/latest/configuration/)
- [Amazon S3 and S3-compatible storage](https://grafana.com/docs/tempo/latest/configuration/hosted-storage/s3/)
- [Tempo deployment modes](https://grafana.com/docs/tempo/latest/reference-tempo-architecture/deployment-modes/)
