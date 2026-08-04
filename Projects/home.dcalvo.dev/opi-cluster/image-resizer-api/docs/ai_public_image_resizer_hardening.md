# Hardening the Image Resizer for Public Internet Access

A public image-resizing endpoint has an unusually convenient denial-of-service
primitive: an attacker can send a relatively small compressed file that requires
significant memory and CPU to decode and transform.

The main risks are resource exhaustion, parser vulnerabilities, abuse, and
accidentally exposing the rest of the homelab—not only someone exploiting the
JPEG endpoint in a traditional sense.

This is a theoretical hardening reference, not a record of implemented features.

## Existing protections

The current application already has several useful controls:

- JPEG content is decoded and validated instead of trusting the filename or
  supplied `Content-Type`.
- Uploads are limited to 10 MiB.
- Decoded images are limited to 40 million pixels.
- Output width is limited.
- Images are processed in memory and are not stored.
- Unknown multipart fields are discarded.
- Error responses do not expose internal errors.
- HTTP read, write, header, and idle timeouts are configured.
- The container runs as a numeric non-root user.
- Its root filesystem is read-only.
- Linux capabilities are removed.
- Privilege escalation is disabled.
- The default seccomp profile is enabled.
- The pod does not receive a Kubernetes service-account token.

These controls align with OWASP recommendations to allowlist formats, validate
actual content, restrict upload size, keep processing libraries updated, and
rewrite images instead of serving uploaded bytes directly. See the
[OWASP File Upload Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/File_Upload_Cheat_Sheet.html).

## The most important missing control: concurrency

One valid request can consume substantial memory across:

- Uploaded compressed bytes
- Decoded source pixels
- Resized RGBA pixels
- Encoded output buffer
- Temporary allocations and garbage collection

The current 40-megapixel limit bounds an individual request, but several requests
can execute simultaneously. A small burst could exceed the pod's 512 MiB memory
limit and cause an OOM kill.

Before public exposure, add:

- A strict per-pod concurrency limit around decode, resize, and encode
- An explicit overload response when no processing slot is available
- A global or edge request-rate limit
- Metrics for admitted and overload-rejected requests
- Load tests to select the concurrency limit empirically

For this homelab, an initial limit might be only one or two concurrent resizes per
pod. That should be established through measurement rather than accepted as a
final number without testing.

OWASP specifically identifies image resizing as a resource-exhaustion risk and
recommends upload limits, request limits, rate limiting, and careful treatment of
resource-intensive operations. See the
[OWASP Denial-of-Service Cheat Sheet](https://cheatsheetseries.owasp.org/cheatsheets/Denial_of_Service_Cheat_Sheet.html).

## Reconsider the public input limits

The existing limits were selected for experimentation. They should be
reconsidered as a public service contract.

Questions to answer:

- Do users genuinely need 40-megapixel inputs?
- Is 10 MiB required, or would 5 MiB be sufficient?
- Is a 4,000-pixel output necessary?
- What is the maximum acceptable processing time?
- How much memory can one request consume under worst-case dimensions?

Reducing these limits is one of the easiest and strongest protections. Supporting
smaller images well is better than nominally accepting larger images and becoming
unstable under concurrency.

Test worst-case files, not only ordinary photographs:

- Maximum encoded bytes
- Maximum pixel count
- Extremely wide or tall aspect ratios
- Truncated and malformed JPEG structures
- Progressive JPEGs
- Images designed to compress unusually well
- Client disconnects during upload and processing
- Repeated maximum-sized requests

Go fuzz tests around multipart parsing, dimension handling, and malformed JPEG
input would also be useful.

## Understand the timeout limitation

The Go server already sets:

- `ReadHeaderTimeout`
- `ReadTimeout`
- `WriteTimeout`
- `IdleTimeout`

These protect connections and slow clients. Go exposes these controls through
[`http.Server`](https://pkg.go.dev/net/http#Server).

However, a connection or write timeout does not necessarily interrupt CPU-bound
image decoding or resizing immediately. The JPEG decoder and scaling operation
are not context-aware. Once expensive processing begins, the handler may continue
consuming CPU until the current stage completes.

Timeouts therefore complement but do not replace:

- Input limits
- Pixel limits
- Concurrency limits
- Process or container isolation
- Resource limits

If strict termination of individual jobs became essential, processing images in
isolated worker processes would provide a harder boundary than a request-context
deadline.

## Put an edge layer in front

Do not expose the current MetalLB Service directly to the internet.

Use a dedicated public entry point that provides:

- DNS and a real hostname
- HTTPS with an automatically renewed certificate
- Hostname and path routing
- Request-body limits
- Connection and request-rate limits
- A maximum request duration
- Basic access logging
- Optional bot or abuse protection
- An emergency way to disable traffic

Kubernetes Ingress or Gateway can provide HTTP routing and TLS termination,
although the exact protections depend on the selected controller. See the
[Kubernetes Ingress documentation](https://kubernetes.io/docs/concepts/services-networking/ingress/).

For a home connection, an edge proxy or outbound tunnel service has an additional
advantage: the origin IP does not need to be exposed and the router does not need
a broad inbound route to the cluster. A public copy hosted on an inexpensive VPS
or managed container service would be safer than exposing the home network.

No application-level measure can protect a residential connection from a
volumetric network attack. That protection must occur upstream.

## Expose only the public operation

The public hostname should route only the intended endpoint:

```text
POST /v1/resize
```

Do not expose publicly:

- `/metrics`
- Prometheus
- Grafana
- The container registry
- Kubernetes API
- Node ports
- Kubelet endpoints
- Internal readiness or liveness endpoints unless the edge explicitly needs them

Grafana currently uses simple administrator credentials, and the registry and
Prometheus are intentionally lightweight homelab services. They must remain
unreachable from the public route.

Use an exact hostname and explicit path rules. A catch-all route into the
homelab would substantially increase the consequences of a configuration mistake.

## Isolate the service from the rest of the homelab

Treat the public application as potentially compromised.

At minimum:

- Allow ingress to its pods only from the selected ingress or gateway.
- Deny application pod egress by default because the resizer needs no outbound
  network access.
- Prevent access from the public namespace to monitoring, the registry,
  Kubernetes APIs, and other homelab workloads.
- Confirm that the k3s CNI actually enforces `NetworkPolicy`.
- Keep the public entry point separated from management interfaces.
- Do not expose SSH, Grafana, Prometheus, the registry, kubelet, etcd, or the
  Kubernetes API.

Kubernetes recommends allowlist-oriented ingress and egress policies and warns
against exposing control-plane interfaces publicly. See the
[Kubernetes Security Checklist](https://kubernetes.io/docs/concepts/security/security-checklist/)
and [NetworkPolicy documentation](https://kubernetes.io/docs/concepts/services-networking/network-policies/).

For stronger isolation, run publicly reachable workloads on a separate VLAN,
separate nodes, or a separate cluster or VPS. A container boundary is useful, but
it should not be the only boundary between an anonymous internet user and the
home network.

## Decide whether the service is truly anonymous

A link shared in a Slack group can escape that group quickly.

Possible models are:

1. **Fully anonymous**

   Simplest for users, but requires strict edge rate limits, global concurrency
   limits, and acceptance that per-IP controls can be bypassed.

2. **Shared invite token**

   Simple but spreads easily and provides poor individual accountability.

3. **Individual API keys**

   Allows per-user quotas and revocation without requiring a complete identity
   platform.

4. **OAuth or an access gateway**

   Stronger identity and easier revocation, but adds more infrastructure.

For a friendly experiment, individual low-privilege API keys or an access
gateway would provide a reasonable balance. If anonymous access is itself part
of the experiment, make the capacity limit extremely conservative.

CORS does not prevent abuse. It only controls what browser JavaScript may read;
anyone can call the endpoint directly with `curl`.

## Response hardening

For successful images, consider explicitly returning:

```text
Content-Type: image/jpeg
X-Content-Type-Options: nosniff
Content-Disposition: attachment
Cache-Control: no-store
```

Whether `no-store` is appropriate depends on whether browsers or intermediaries
should cache results. Because uploaded images may contain personal material,
avoiding unintended caching is a sensible default.

Continue re-encoding the image instead of returning the original uploaded bytes.
Re-encoding also avoids preserving arbitrary embedded file data and metadata
unless that capability is intentionally added later.

## Kubernetes and container controls

The current pod security context is already a strong start. Retain it and add or
verify:

- An intentionally selected CPU limit or another method of preventing starvation
  of other homelab services
- Memory limits tested against worst-case concurrency
- Network policies
- A supported Pod Security standard for the namespace
- Pinned base and application dependencies
- Image scanning
- An SBOM
- Deployment by immutable image digest
- Prompt rebuilding when Go or image-processing dependencies receive security
  fixes

Parser libraries are part of the attack surface because every uploaded file
reaches them.

## Monitoring and incident response

Before inviting users, add alerts for:

- No healthy replicas
- Service failures and rapidly increasing latency
- Overload rejection rate
- OOM kills and pod restarts
- CPU throttling or node pressure associated with user impact
- Sudden traffic or upload-volume increases
- Missing Prometheus targets
- TLS certificate expiry
- External synthetic resize failures

Prepare a small operational response:

- One command to disable public routing
- One command to reduce allowed traffic
- A known-good rollback
- A way to block a token or abusive source
- A way to see current concurrency, errors, pods, and resource usage
- A clear maximum amount of time to support the experiment

Avoid logging image content, filenames, authorization tokens, or complete request
bodies. Think deliberately about IP-address retention because public access logs
can contain personal information.

## Practical minimum before sharing it

For this experiment, a reasonable minimum release gate is:

1. Put HTTPS and an edge proxy or tunnel in front.
2. Route only `/v1/resize` to the application.
3. Keep every management and monitoring service private.
4. Add a strict concurrency limit and explicit overload response.
5. Add edge rate limiting.
6. Revisit and probably reduce image byte, pixel, and output limits.
7. Add a `NetworkPolicy` with no application egress.
8. Test worst-case concurrent images against the 512 MiB pod limit.
9. Add an external correctness check and a few critical alerts.
10. Prepare a kill switch and rollback procedure.
11. Keep Go, container, and image-processing dependencies patched.
12. Prefer deploying the public copy away from the trusted home LAN.

The service already has a better baseline than many small experimental APIs. The
largest gaps are not ordinary input validation; they are bounded concurrency,
edge protection, network isolation, and ensuring that only this endpoint—not the
rest of the homelab—becomes public.
