# On-premises object storage for Tempo

The central distinction is that an on-premises Tempo deployment should use an
object store rather than expose block storage directly to Tempo.

The physical implementation may still use disks or block devices underneath:

```text
Physical disks or block devices
            |
            v
Distributed object-storage system
            |
            | S3-compatible API
            v
          Tempo
```

Tempo sees an S3-compatible endpoint, a bucket, and credentials. The
object-storage system handles disk placement, replication, erasure coding, node
failures, and capacity.

## Common on-premises patterns

### 1. Use the company's existing object-storage platform

Reusing an existing production storage platform is generally the preferred
option.

A company with hundreds of servers may already operate something such as:

- Ceph Object Gateway
- Dell ECS or ObjectScale
- NetApp StorageGRID
- Cloudian
- Scality
- another enterprise system exposing a compatible S3 API

The storage team would provide information such as:

```text
S3 endpoint: https://objects.internal.example
Bucket:      tempo-traces
Credentials: Tempo-specific access key
```

Tempo would be configured approximately as follows:

```yaml
storage:
  trace:
    backend: s3
    s3:
      endpoint: objects.internal.example
      bucket: tempo-traces
      access_key: ${S3_ACCESS_KEY}
      secret_key: ${S3_SECRET_KEY}
      insecure: false
```

The credentials should come from a Kubernetes Secret or another secrets system
rather than being written directly into a ConfigMap.

### 2. Operate Ceph with an S3 gateway

For an open source, Kubernetes-oriented environment, Rook-managed Ceph is a
significant architectural option.

Ceph can expose three different types of storage interface:

```text
Ceph
|-- RBD     -> block volumes and Kubernetes PVCs
|-- CephFS  -> shared filesystem
`-- RGW     -> S3-compatible object storage
```

Tempo needs the third interface: Ceph Object Gateway, commonly called RGW.

If a company already uses Ceph RBD for Kubernetes PVCs, that does not
automatically mean Tempo can use those block volumes as distributed trace
storage. The team would enable Ceph RGW and provide Tempo with an S3 bucket.

Rook can deploy that object gateway inside Kubernetes:

```text
Tempo Pods
    |
    | S3 API
    v
Rook Ceph RGW Service
    |
    v
Ceph storage cluster
```

A Rook `CephObjectStore` provides an S3 API backed by replicated or
erasure-coded Ceph pools across storage nodes.

### 3. Use a commercial S3-compatible product

Some organizations prefer a supported storage appliance or software product
instead of operating Ceph themselves.

This can be operationally attractive because the storage vendor supports:

- storage-node failures
- disk replacement
- upgrades
- replication and erasure coding
- capacity expansion
- multi-site replication
- monitoring and incident response

From Tempo's perspective, this still appears as an S3-compatible endpoint.

### 4. Use MinIO or another self-hosted implementation

MinIO has historically been a common answer for self-hosted S3-compatible
storage. Its community distribution and product model have changed, however,
and Grafana's current Tempo documentation describes the community MinIO setup
as an evaluation example rather than a production recommendation.

MinIO may still be viable through a supported offering, but it should not be
selected automatically for a new large installation without evaluating its
current licensing, release, upgrade, and support arrangements.

SeaweedFS and `rclone serve s3` also appear in Tempo's local testing
documentation. Grafana does not recommend those documented examples as
production Tempo storage.

## A likely large on-premises architecture

One possible architecture is:

```text
Applications
    |
    v
Alloy or OpenTelemetry Collectors
    |
    v
Tempo microservices
    |-- Kafka cluster
    `-- internal S3-compatible object store
            `-- replicated across storage nodes
```

The object store is normally operated as an infrastructure service of its own.
It might run:

- inside the observability Kubernetes cluster
- in a dedicated Kubernetes storage cluster
- on separate storage servers
- as an existing enterprise storage appliance

Operating storage separately reduces correlated failures. If the observability
Kubernetes cluster fails completely, its historical traces can remain
available in the independent storage system.

## Storage sizing is not based only on server count

Having hundreds of servers does not directly determine the required trace
storage. The important inputs include:

- spans produced per second
- average encoded span size
- sampling percentage
- retention period
- replication or erasure-coding overhead
- compaction efficiency
- query and object-request volume

A simplified estimate is:

```text
storage required
approximately equals
ingest bytes per second
times retention seconds
times storage overhead
```

A company with 500 lightly traced servers might produce less trace data than a
company with 30 extremely busy microservices.

## Practical conclusion

The usual on-premises design is to provide Tempo with a bucket on an existing,
production-grade S3-compatible object-storage platform.

If the organization does not already have one, Ceph RGW, often managed through
Rook in Kubernetes, is an important open source option. Operating Ceph is a
substantial infrastructure responsibility, however, and should be treated as
its own production storage service rather than as a minor Tempo add-on.

## Current references

- [Tempo object-storage architecture](https://grafana.com/docs/tempo/latest/reference-tempo-architecture/object-storage/)
- [Tempo S3 and S3-compatible storage](https://grafana.com/docs/tempo/latest/configuration/hosted-storage/s3/)
- [Rook Ceph object storage](https://rook.io/docs/rook/latest-release/Storage-Configuration/Object-Storage-RGW/object-storage/)
- [Rook Ceph storage architecture](https://rook.io/docs/rook/latest/Getting-Started/storage-architecture/)
