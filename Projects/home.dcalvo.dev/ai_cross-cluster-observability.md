# Cross-Cluster Observability Design

## Purpose

Move the central monitoring and observability stack from the Orange Pi (`opi`)
cluster to the HP (`hp`) cluster while continuing to collect metrics, traces,
and logs from workloads running in `opi`.

The intended central stack is:

- Prometheus for metrics
- Tempo for traces
- Loki for logs
- Grafana for visualization
- Grafana Alloy for telemetry collection and forwarding

The design should remain simple enough for a trusted home-lab network and
should not require a full VPN or elaborate multicluster networking system.

## Recommended Architecture

Keep an Alloy collector in every environment that produces telemetry, and
send telemetry outward to a restricted Alloy gateway in the HP cluster.

```text
OPI applications
  |-- /metrics -------+
  |-- OTLP traces ----+--> OPI Alloy (ClusterIP, local only)
  `-- pod logs -------+              |
                                     | outbound LAN traffic
                                     v
                            HP Alloy gateway (MetalLB)
                            restricted to OPI node IPs
                              |-- metrics --> Prometheus
                              |-- traces  --> Tempo
                              `-- logs    --> Loki.

                            Grafana --> internal ClusterIP backends
```

In this arrangement:

- Applications in `opi` send traces to the local Alloy `ClusterIP`, as they do
  now.
- OPI Alloy discovers and scrapes OPI applications locally.
- OPI Alloy reads OPI pod logs locally.
- OPI Alloy forwards all three signals to the HP cluster.
- Prometheus, Loki, Tempo, and Grafana live in the HP cluster.
- Prometheus in HP does not need direct access to OPI pods.
- Tempo and Loki can remain private `ClusterIP` services in HP.

This is an edge-collector-to-central-observability pattern. Alloy supports
Prometheus scraping and remote write for metrics, `loki.write` for logs, and
OTLP forwarding for traces.

Terminology reminder: Prometheus stores metrics and Tempo stores traces. Alloy
can receive both but sends them to different backends.

## Why Direct Pod and ClusterIP Routing Is Not Suitable

Both clusters' live internal networks were inspected on 2026-08-03 using their
explicit kubeconfig contexts. They overlap exactly:

| Network | OPI | HP |
| --- | --- | --- |
| Node pod CIDRs | `10.42.0.0/24`, `10.42.1.0/24` | `10.42.0.0/24`, `10.42.1.0/24` |
| Kubernetes Service network evidence | `kubernetes` Service at `10.43.0.1` | `kubernetes` Service at `10.43.0.1` |

An HP pod therefore cannot route directly to an OPI pod or `ClusterIP`. An
address such as `10.42.0.15` could independently exist in both clusters.
Adding static routes would not resolve the ambiguity.

A service mesh, Submariner, or routed VPN would require dealing with these
overlapping networks and would be excessive for the current use case.

## MetalLB Exposure and the Home LAN

The MetalLB ranges are private LAN addresses:

- OPI: `192.168.1.220-192.168.1.229`
- HP: `192.168.1.230-192.168.1.239`

MetalLB advertises these addresses only on the home Layer 2 network. They are
not Internet-accessible unless the router has a port-forward or another rule
that deliberately publishes them.

This is reasonably analogous to an internal cloud load balancer: it is
reachable from the surrounding private network but not inherently from the
public Internet.

Other devices on the home LAN could nevertheless attempt to connect to the
VIP. Create a dedicated HP `LoadBalancer` Service for telemetry ingestion and
restrict its accepted source addresses:

```yaml
spec:
  type: LoadBalancer
  loadBalancerSourceRanges:
    - 192.168.1.201/32 # opi1
    - 192.168.1.202/32 # opi2
    - 192.168.1.X/32   # development workstation, when needed
```

Source-address filtering depends on the Kubernetes Service traffic policy and
networking implementation. Test the finished Service from both an allowed and
a denied machine.

`externalTrafficPolicy: Local` can preserve the original client address, but it
also requires an Alloy gateway endpoint on the node announcing the MetalLB
address. This behavior and the deployment topology should be validated before
relying on source filtering as the security boundary.

Additional optional safeguards include:

- Restrict the telemetry ports in the HP nodes' firewall.
- Put a small reverse proxy in front of Alloy and require a shared bearer token
  or basic authentication.
- Add TLS if the LAN is no longer considered trusted.

An Ingress is not required. A dedicated Layer 4 MetalLB Service exposing only
the ingestion ports is simpler.

## Metrics Flow

There are two reasonable edge-scraping choices:

1. Run Alloy on OPI as the scraper and remote-write into HP Prometheus.
2. Run Prometheus in agent mode on OPI and remote-write into HP.

Alloy is preferred because it is already required for traces and logs. It can
perform Kubernetes discovery and may be able to reuse the existing
ServiceMonitor resources, including the Image Resizer ServiceMonitor at
`opi-cluster/image-resizer-api/k8s/service-monitor.yaml`.

The central Prometheus must enable its remote-write receiver. Its write
endpoint is `/api/v1/write`.

Prometheus documents that its built-in receiver is not intended to replace a
high-volume general-purpose ingestion backend. It should nevertheless be
reasonable for this home-lab volume. If ingestion grows substantially,
VictoriaMetrics or Grafana Mimir would be more natural central receivers.

## Trace Flow

The Image Resizer currently sends OTLP to the OPI Alloy `ClusterIP`, configured
in `opi-cluster/monitoring/alloy/values.yaml`. Keep that application-to-Alloy
relationship local and change Alloy's exporter destination:

```text
Current:
OPI application -> OPI Alloy -> OPI Tempo

After migration:
OPI application -> OPI Alloy -> restricted HP gateway -> HP Tempo
```

Applications do not need to know that Tempo moved to another cluster. Tempo
itself can remain a `ClusterIP` service in HP.

## Log Flow

Keep Alloy in OPI so it can discover OPI pods and read their logs through the
OPI Kubernetes API. Replace its current local Loki destination with the
restricted HP ingestion endpoint. The HP gateway then forwards the logs to the
HP Loki `ClusterIP` service.

## Development Workstation

The same central HP gateway can later accept telemetry from the development
workstation by adding the workstation's stable LAN address to the source
allowlist.

A locally running Alloy collector is preferable to having every development
application communicate with the HP backends independently. Applications can
send OTLP to localhost, while local Alloy handles batching, buffering, and the
single connection to the HP gateway.

## Security Boundary

If "not exposed to the LAN" means that LAN devices may see the IP but their
connections are denied, MetalLB combined with source allowlisting and optional
node firewall rules is sufficient.

If it means that the endpoint must not exist on the home LAN at all, Kubernetes
cannot provide that isolation by itself. A separate network boundary would be
required, such as:

- a VLAN;
- a secondary physical network;
- WireGuard or another VPN; or
- a multicluster networking system.

For this environment, a restricted HP Alloy gateway plus an Alloy edge
collector in OPI provides the best balance: no elaborate VPN, no direct pod
routing, and only one controlled cross-cluster ingestion surface.

## References

- [Grafana Alloy: Collect OpenTelemetry data and forward to the LGTM stack](https://grafana.com/docs/alloy/latest/collect/opentelemetry-to-lgtm-stack/)
- [Grafana Alloy as a proxy or aggregation layer](https://grafana.com/docs/alloy/latest/configure/proxy/)
- [Prometheus remote-write storage support](https://prometheus.io/docs/prometheus/latest/storage/)
- [Prometheus Agent mode](https://prometheus.io/docs/prometheus/latest/prometheus_agent/)
- [Kubernetes: Using source IP](https://kubernetes.io/docs/tutorials/services/source-ip/)
- [MetalLB advanced Layer 2 configuration](https://metallb.io/configuration/_advanced_l2_configuration/)

