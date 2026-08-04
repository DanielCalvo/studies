# AWS Production and Observability Cluster Architecture

## Purpose

This note explores how to separate application workloads and observability
backends when running two Kubernetes clusters on Amazon Web Services:

- A **production EKS cluster** serves application traffic.
- An **observability EKS cluster** runs Prometheus, Tempo, Loki, Grafana, and
  related telemetry services.

The goal is to transfer metrics, traces, and logs from the production cluster
to the observability cluster with a reasonable level of security, isolation,
reliability, and operational simplicity.

Although this discussion is based on Amazon EKS and Amazon VPC networking, the
central design principle is portable: collect telemetry close to its source
and forward it through a small, controlled ingestion boundary.

## Recommended Data Flow

Run Grafana Alloy or another OpenTelemetry-compatible collector inside the
production cluster. Applications communicate only with this local collector.
The collector then pushes telemetry to a gateway in the observability cluster.

```text
Production EKS cluster

  Application workloads
    |-- Prometheus /metrics --+
    |-- OTLP traces ----------+--> Production Alloy
    `-- application logs -----+      edge collector
                                      |
                                      | metrics: Prometheus remote write
                                      | traces: OTLP
                                      | logs: Loki push or OTLP
                                      v
                             Private ingestion endpoint
                                      |
                                      v
Observability EKS cluster

  Alloy or OpenTelemetry gateway
    |-- metrics --> Prometheus or compatible backend
    |-- traces  --> Tempo
    `-- logs    --> Loki

  Grafana reads from the internal observability backends.
```

This design avoids giving the observability cluster direct access to every
application pod in production.

## Why Collection Should Remain in the Production Cluster

Central Prometheus could theoretically scrape production workloads directly
over cross-VPC networking. That creates several disadvantages:

- Production pod addresses and Kubernetes discovery must be reachable from the
  observability VPC.
- The observability cluster needs credentials and network access into the
  production cluster.
- Security rules must cover many changing scrape targets.
- Network failures become difficult to distinguish from failed applications.
- Logs still require a production-side collector, so direct metric scraping
  does not eliminate the need for an edge agent.

Keeping Alloy in production gives the following properties:

- Applications use a local `ClusterIP` or node-local endpoint.
- Kubernetes discovery stays inside the production trust boundary.
- Only one controlled destination must be reachable across networks.
- The collector can batch, queue, retry, filter, relabel, and enrich telemetry.
- The production cluster initiates the connection; the observability cluster
  does not need to initiate connections into production.
- The same pattern handles metrics, traces, and logs.

The signal-specific paths are:

```text
Metrics:
production targets <- local scrape - production Alloy
production Alloy - Prometheus remote write -> observability gateway

Traces:
production applications - OTLP -> production Alloy
production Alloy - OTLP -> observability gateway

Logs:
production Alloy reads production pod logs
production Alloy - Loki push or OTLP -> observability gateway
```

## Internal AWS Load Balancers

An AWS load balancer can use one of two schemes:

- `internet-facing`: its nodes have public addresses and can receive Internet
  traffic when the surrounding routing and security rules permit it.
- `internal`: its nodes have only private addresses and clients require private
  network connectivity to its VPC.

An internal load balancer in private subnets is therefore an appropriate front
door for observability ingestion. It is not, however, a complete authorization
boundary by itself.

The distinction is:

- The **internal load-balancer scheme** determines that the endpoint is not
  publicly addressed.
- **VPC connectivity and routing** determine which private networks can reach
  it.
- **Security groups, TLS, and application authentication** determine which
  clients should be allowed to use it.

For EKS, the AWS Load Balancer Controller or EKS Auto Mode can create an
internal Network Load Balancer from a Kubernetes `Service` of type
`LoadBalancer`. IP target mode can send traffic directly to eligible pod IPs.

## Scenario 1: Both Clusters in One VPC

If the production and observability clusters share one VPC, an internal NLB in
private subnets is normally sufficient as the network entry point.

The deployment should also include:

- Security groups allowing only the production nodes, production pod security
  groups, or explicitly identified client ranges.
- Listeners only for the required telemetry protocols and ports.
- TLS encryption.
- Authentication between production Alloy and the observability gateway.
- Kubernetes NetworkPolicy around the gateway and backend services.
- Multiple gateway replicas across Availability Zones.

This is the simplest AWS arrangement, but both clusters share a relatively
broad network trust domain. Security groups and Kubernetes policy carry more
of the isolation responsibility.

## Scenario 2: Separate VPCs with Peering

For two VPCs, a VPC peering connection can provide private routed
connectivity. The architecture becomes:

```text
Production Alloy
    |
    v
VPC peering connection
    |
    v
Internal NLB in observability VPC
    |
    v
Observability gateway
```

This requires:

- Non-overlapping VPC CIDR ranges.
- Routes in the relevant production and observability subnet route tables.
- Security-group rules allowing only the necessary source and destination.
- DNS configuration if private load-balancer names must resolve across the
  peering connection.

Peering is a reasonable choice when the two VPCs already need broader private
communication. It is simple for a small number of VPCs, but it establishes a
routed relationship between them. The security groups must ensure that this
does not accidentally become broader access than intended.

## Scenario 3: Separate VPCs with Transit Gateway

AWS Transit Gateway is appropriate when many VPCs or accounts need controlled
connectivity through a central network hub.

It can place production and observability VPCs into different routing domains
and permit only intended paths. It is usually unnecessary for only two VPCs,
but becomes useful in a larger organization with multiple application clusters
feeding a central observability platform.

Transit Gateway still provides routed network connectivity. Route tables,
security groups, and network architecture must enforce the intended
segmentation.

## Scenario 4: Separate VPCs with AWS PrivateLink

AWS PrivateLink is the most narrowly scoped option for this use case.

The observability VPC publishes its ingestion NLB as a VPC endpoint service.
The production VPC creates an interface VPC endpoint for that service.
Production Alloy connects to private interface addresses located inside the
production VPC.

```text
Production EKS cluster
  Production Alloy
          |
          v
  Interface VPC endpoint
  private IPs in production subnets
          |
          v
      AWS PrivateLink
          |
          v
  Endpoint service backed by NLB

Observability EKS cluster
  Alloy/OpenTelemetry ingestion gateway
          |
          +--> Prometheus
          +--> Tempo
          `--> Loki
```

PrivateLink provides several useful properties:

- Traffic does not traverse the public Internet.
- No NAT gateway is needed for the telemetry path.
- The VPCs do not require a general routing relationship.
- Only the consumer initiates connections to the endpoint service.
- The endpoint service can authorize specific AWS principals or accounts.
- Production sees private endpoint interfaces inside its own VPC.
- The producer and consumer VPCs can have overlapping CIDR ranges.

PrivateLink endpoint services require a Network Load Balancer or Gateway Load
Balancer. An NLB is the natural choice for OTLP gRPC, OTLP HTTP, Prometheus
remote write, and Loki ingestion.

The tradeoffs are additional hourly and data-processing charges, endpoint
management per VPC and Availability Zone, and somewhat more infrastructure
than straightforward peering.

## Choosing Between Peering and PrivateLink

| Requirement | VPC peering | AWS PrivateLink |
| --- | --- | --- |
| General bidirectional VPC connectivity | Good fit | Not intended for this |
| One narrowly exposed ingestion service | Possible | Best fit |
| Overlapping VPC CIDRs | Not supported | Supported |
| Production initiates connections only | Must be enforced with rules | Natural model |
| Simple setup for two trusted VPCs | Usually simpler | More components |
| Strong service-level isolation | Relies heavily on routing and security groups | Stronger default boundary |
| Cost | Usually lower | Endpoint and data charges |

For two trusted VPCs that already need to communicate, peering plus an
internal NLB is adequate and may be the simplest design.

For a production VPC that should communicate with only a dedicated
observability ingestion service, PrivateLink is the cleaner security model.

## Ingestion Gateway

The public-facing component in this context is not actually public. It is a
privately reachable ingestion gateway behind the internal NLB.

This gateway can be Grafana Alloy, an OpenTelemetry Collector, or a small
reverse-proxy layer in front of those collectors. It should expose only the
required receiver protocols, for example:

- OTLP gRPC for traces on TCP `4317`.
- OTLP HTTP for telemetry on TCP `4318`.
- A dedicated HTTP endpoint for Prometheus remote write.
- A Loki-compatible push endpoint if logs are not transported through OTLP.

The gateway then talks to Prometheus, Tempo, and Loki through private
Kubernetes `ClusterIP` services. The storage backends do not need their own
cross-VPC load balancers.

Using one gateway provides:

- A small, auditable network surface.
- Central authentication and TLS handling.
- Consistent tenant or cluster labels.
- Rate limiting and request-size controls.
- A stable destination while backend services change.
- Separation between remote clients and storage services.

## Metrics Backend Considerations

Production Alloy can scrape Prometheus-format endpoints locally and use
Prometheus remote write for transfer.

A single Prometheus server can enable its built-in remote-write receiver at
`/api/v1/write`. Prometheus documents this as appropriate for selected
lower-volume use cases rather than as a high-scale general-purpose ingestion
backend.

For a larger production system, a backend designed for distributed remote
write ingestion may be preferable, such as:

- Grafana Mimir;
- Amazon Managed Service for Prometheus; or
- another Prometheus-compatible scalable metrics store.

The edge-collector and private-ingestion architecture remains the same when
the metrics backend changes.

## Security Controls

An internal load balancer answers whether an endpoint is privately addressed.
It does not, on its own, authenticate telemetry producers.

A reasonable AWS security design combines the following layers.

### Network exposure

- Use an internal NLB, never an Internet-facing load balancer, for ingestion.
- Place load-balancer nodes and EKS workloads in appropriate private subnets.
- Prefer PrivateLink when the VPCs should not have general routed access.
- Restrict security groups to the required clients and ports.
- Use Kubernetes NetworkPolicy to restrict access after traffic enters EKS.

### Encryption and authentication

- Use TLS for remote write, OTLP, and log ingestion.
- Authenticate each production collector with a client certificate, bearer
  token, or another supported credential.
- Store secrets in AWS Secrets Manager, Kubernetes Secrets with appropriate
  controls, or another managed secret system.
- Rotate credentials rather than treating private connectivity as sufficient
  authentication.

PrivateLink authorizes which principals may create endpoint connections, but
that does not replace application-level authentication at the telemetry
gateway.

### Availability and failure handling

- Run gateway replicas across at least two Availability Zones.
- Create NLB targets in the corresponding zones.
- Give production collectors persistent queues or write-ahead logs where
  supported.
- Apply batching, bounded retries, memory limits, and backpressure controls.
- Ensure telemetry failures cannot exhaust production application resources.

### Data separation

- Add an immutable cluster or environment label such as
  `cluster=production-eu-west-1` at the trusted collector or gateway.
- Prevent clients from overriding authoritative tenant or environment labels.
- Limit high-cardinality labels and sensitive log or trace attributes.
- Apply appropriate retention policies to each signal.

## Recommended AWS Architecture

For strongly separated production and observability VPCs, the preferred design
is:

```text
Production applications
    |
    v
Production Alloy collector
    |
    v
PrivateLink interface endpoint in production VPC
    |
    v
Internal NLB endpoint service in observability VPC
    |
    v
Alloy/OpenTelemetry ingestion gateway
    |-- Prometheus-compatible metrics backend
    |-- Tempo
    `-- Loki
```

Use VPC peering and an internal NLB instead when the two VPCs already have a
legitimate need for broader private connectivity and their CIDR ranges do not
overlap.

In either case, retain the production-side collector. The important safety
property is that production pushes telemetry through one authenticated,
encrypted, tightly restricted endpoint; the observability cluster does not
receive broad access back into production.

## References

- [AWS: How Elastic Load Balancing works](https://docs.aws.amazon.com/elasticloadbalancing/latest/userguide/how-elastic-load-balancing-works.html)
- [Amazon EKS: Route TCP and UDP traffic with Network Load Balancers](https://docs.aws.amazon.com/eks/latest/userguide/network-load-balancing.html)
- [Amazon EKS: Load-balancing best practices](https://docs.aws.amazon.com/eks/latest/best-practices/load-balancing.html)
- [AWS: How VPC peering connections work](https://docs.aws.amazon.com/vpc/latest/peering/vpc-peering-basics.html)
- [AWS: Create a service powered by AWS PrivateLink](https://docs.aws.amazon.com/vpc/latest/privatelink/create-endpoint-service.html)
- [AWS: Share services through AWS PrivateLink](https://docs.aws.amazon.com/vpc/latest/privatelink/privatelink-share-your-services.html)
- [AWS: PrivateLink in a multi-VPC architecture](https://docs.aws.amazon.com/whitepapers/latest/building-scalable-secure-multi-vpc-network-infrastructure/aws-privatelink.html)
- [Grafana Alloy: Collect OpenTelemetry data and forward to the LGTM stack](https://grafana.com/docs/alloy/latest/collect/opentelemetry-to-lgtm-stack/)
- [Prometheus remote-write storage support](https://prometheus.io/docs/prometheus/latest/storage/)

