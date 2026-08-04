# Testing Terminology

Tests can overlap in the behavior they exercise while verifying different layers
of a system.

## Unit tests

Unit tests exercise an isolated function or a small piece of behavior inside the
application process.

In the image-resizer project, directly testing the `scaledHeight` function is a
unit test.

## Component tests

Component tests exercise several parts of one application together without
requiring the complete deployed environment.

Sending a multipart request through the image-resizer HTTP handler with Go's
`httptest` package is closer to a component test than a pure unit test. It covers
HTTP parsing, validation, JPEG processing, response generation, logging, and
metrics in the same process, but it does not use a real network, container, or
Kubernetes cluster.

These are also sometimes called in-process integration tests.

## Integration tests

Integration tests verify that separately implemented components or services work
together. Examples include an application communicating with a database, message
broker, object store, or another service.

The exact boundary varies between teams, so the term should ideally be qualified
with what is being integrated.

## Smoke tests

Smoke tests are a small and fast set of critical checks against a built or
deployed application. They answer a basic question:

> Is the release sufficiently functional to continue testing or serving traffic?

The current checks in `build-and-deploy.sh` are post-deployment smoke tests:

```text
GET /livez
GET /readyz
GET /metrics
```

They verify that the release started, became ready, and exposes its essential
operational endpoints.

## End-to-end tests

End-to-end tests exercise a complete user journey through the real system and its
deployment boundaries.

For the deployed image resizer, uploading a known JPEG through the MetalLB address
and verifying the returned JPEG would exercise:

```text
workstation
  -> MetalLB
  -> Kubernetes Service
  -> selected pod
  -> JPEG decode and resize
  -> HTTP response
```

This verifies more than the equivalent in-process test:

- The ARM64 binary works.
- The container starts.
- Kubernetes can pull the image.
- The Service selector and port mapping are correct.
- MetalLB networking works.
- The real HTTP endpoint is reachable.
- The deployed application can resize a JPEG.

A small check of this journey after a deployment can be called an **end-to-end
smoke test** or **functional smoke test**.

## Acceptance and regression tests

Acceptance tests determine whether the deployed application satisfies its agreed
business or user requirements.

A broad deployed suite covering successful resizing, invalid inputs, resource
limits, error responses, concurrency, and supported formats could be described as
a functional acceptance suite or deployed end-to-end regression suite.

## Synthetic monitoring

Synthetic monitoring repeatedly performs selected user-like requests against a
running environment after deployment. Unlike a one-time smoke test, it continues
checking availability and behavior over time.

## Practical naming

For this project, the clearest names are:

| Test | Suggested name |
| --- | --- |
| Direct test of one Go function | Unit test |
| Request through the in-process Go HTTP handler | Component test or in-process integration test |
| Health and metrics checks immediately after deployment | Post-deployment smoke tests |
| Upload a real JPEG through MetalLB and verify the result | End-to-end smoke test |
| Comprehensive deployed behavior suite | End-to-end or functional acceptance tests |
| Repeated production-like checks over time | Synthetic monitoring |

It is normal for tests at different layers to cover the same behavior. A unit or
component test verifies the application logic quickly and precisely, while a
post-deployment test verifies that packaging, configuration, networking, and the
runtime environment allow that behavior to work in the real system.
