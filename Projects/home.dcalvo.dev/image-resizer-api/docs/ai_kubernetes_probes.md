# Kubernetes Probes

Kubernetes supports three probe purposes.

## Liveness probe

Asks: **Is this container still functioning?**

Repeated failure causes Kubernetes to restart the container.

## Readiness probe

Asks: **Should this pod receive traffic?**

Failure removes the pod from matching Service endpoints, but does not restart the
container.

## Startup probe

Asks: **Has this application finished starting?**

When a startup probe is configured, Kubernetes holds back liveness and readiness
checks until the startup probe succeeds. Repeated startup-probe failure eventually
causes the container to be restarted.

The short conceptual distinction is:

```text
Startup:   Has it started yet?
Readiness: Should it receive traffic?
Liveness:  Should it be restarted?
```

## Probe mechanisms

Each probe can perform its check using:

- An HTTP request
- A TCP connection
- A gRPC health check
- A command executed inside the container

Common timing and threshold settings include:

```yaml
initialDelaySeconds:
periodSeconds:
timeoutSeconds:
failureThreshold:
successThreshold:
```

## Image Resizer API

The likely configuration for the image-resizer service is:

```text
livenessProbe  -> GET /livez
readinessProbe -> GET /readyz
```

A startup probe is probably unnecessary while the application starts quickly and
performs no lengthy initialization. It becomes useful for slow-starting
applications that might otherwise be killed prematurely by their liveness probe.
