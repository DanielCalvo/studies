# Preliminary Tempo Study Plan

## Why Tempo belongs in Chapter 10

Chapter 10 of *Cloud Native Observability with OpenTelemetry* is about telemetry
backends: systems that store telemetry and make it available for querying and
visualization.

The book was published in 2021 and uses Zipkin and Jaeger as its principal
examples of tracing backends. This exercise goes slightly off the book's script
by using Grafana Tempo as a modern tracing-backend case study. The underlying
chapter concepts remain the same:

- exporting telemetry to a backend
- storing and retrieving telemetry
- querying and visualizing telemetry
- understanding production concerns such as retention, scaling, availability,
  limits, and privacy

This is a preliminary plan. We can adapt it as our understanding and interests
develop.

## The central distinction

Tempo is not another general-purpose OpenTelemetry Collector.

An OpenTelemetry Collector normally moves telemetry through a pipeline:

```text
receivers -> processors -> exporters
```

Grafana Alloy represents its work as a graph of connected components. It
includes many OpenTelemetry Collector components as well as components from the
Prometheus, Loki, and other ecosystems.

Tempo is a tracing database. Its primary responsibilities are:

- receiving and validating spans
- retaining recent trace data
- writing trace data into persistent blocks
- finding traces by trace ID or a TraceQL query
- returning trace data to a client such as Grafana

A simplified Tempo data path is:

```text
                           +-- recent trace data
OTLP -> distributor ------+
                           +-- persistent trace blocks
                                        |
Grafana -> query frontend -> querier ---+
```

Processing such as filtering, enrichment, retries, batching, and sampling
usually belongs in the Collector or Alloy in front of Tempo.

## Scope and learning objective

The objective is not to study every Tempo option. It is to become comfortable
enough with Tempo's major configuration areas that we can:

- read a small Tempo configuration and explain each major section
- understand the write, storage, and read paths
- distinguish Tempo's job from the jobs of Alloy, the Collector, and Grafana
- operate a useful local Tempo instance
- identify the major changes required for a production deployment

We will begin with monolithic Tempo and temporary local storage in Minikube. We
will not begin with Kafka, object storage, multi-tenancy, or a distributed Tempo
deployment.

## Study environment: Minikube

The exercise will run inside the existing Minikube cluster rather than through
Docker Compose. This adds practical Kubernetes experience while preserving the
incremental Tempo learning path.

All Chapter 10 resources will live in a dedicated namespace:

```text
chapter10-tempo
```

We will initially write small Kubernetes manifests ourselves instead of hiding
Tempo's configuration inside a Helm chart:

- a ConfigMap will contain `tempo.yaml`
- a Deployment will run one monolithic Tempo replica
- a Service will expose Tempo's query and OTLP ingestion ports inside the
  cluster
- `kubectl port-forward` will give us local access when needed
- a trace producer will run as a Pod or Job inside the cluster

The initial topology will be:

```text
Trace producer Pod
        |
        | OTLP
        v
Tempo Service
        |
        v
Monolithic Tempo Pod
```

As the exercise develops, it will become:

```text
Application -> Alloy -> Tempo -> Grafana
                         |
                         +-> persistent volume
```

Kubernetes controls how the Tempo process is deployed and reached. The
`tempo.yaml` file still controls how Tempo receives, stores, and serves trace
data. Keeping these two configuration layers distinct is part of the exercise.

Before starting the first step, we will inspect Minikube's available CPU and
memory and the workloads left running from Chapter 9. We will add only the
component needed by the current step so that Tempo, Grafana, Alloy, and
Prometheus do not consume resources before they are needed.

## Step 1 - Run the smallest possible Tempo

Create:

- a dedicated `chapter10-tempo` namespace
- one ConfigMap containing a short, annotated `tempo.yaml`
- one Deployment containing a single Tempo Pod
- one Service exposing Tempo inside the cluster
- monolithic mode
- temporary local filesystem storage through an `emptyDir` volume
- an OTLP receiver

Initially, we will only start Tempo and inspect its rollout, Pod logs, and
readiness endpoint. We can reach the HTTP endpoint with
`kubectl port-forward`.

Concepts:

- the purpose of `target: all`
- Tempo's HTTP server compared with its OTLP receiver
- `distributor.receivers`
- `storage.trace`
- the distinction between the Tempo ConfigMap, Deployment, and Service
- why local storage is useful for study but generally not the production choice

Expected observation:

- the Tempo Pod becomes ready and the Tempo readiness endpoint responds
  successfully.

Deliberately excluded:

- Grafana
- Alloy or an OpenTelemetry Collector
- a trace-producing application
- Prometheus
- persistent storage
- object storage

## Step 2 - Ingest one trace directly

Run a small trace-producing program or telemetry generator in a Kubernetes Pod
or Job. It will send one trace directly to the Tempo Service over OTLP.

Concepts:

- OTLP is an ingestion protocol
- the distributor is Tempo's write entry point
- Kubernetes service discovery lets the producer address Tempo by Service name
- spans with the same trace ID form a trace
- accepting a trace and conveniently querying it are separate capabilities

Expected observation:

- Tempo accepts a trace, and we retrieve it using its trace ID.

We will deliberately bypass Alloy at this point so that Tempo's responsibility
remains visible.

## Step 3 - Add Grafana and query Tempo

Add a single-replica Grafana Deployment and Service, then provision Tempo as a
Grafana data source using its Kubernetes Service address. We will access Grafana
locally through `kubectl port-forward`.

Concepts:

- Tempo stores and queries traces
- Grafana renders and explores them
- Grafana reaches Tempo over the cluster network
- querying by trace ID compared with searching through TraceQL
- resource and span attributes as searchable information

Expected observation:

- we can inspect the trace in Grafana and run one small TraceQL query.

## Step 4 - Understand trace persistence

Replace Tempo's temporary `emptyDir` storage with a
PersistentVolumeClaim and examine the related Kubernetes and Tempo storage
configuration.

Concepts:

- recent trace data compared with persisted trace blocks
- Tempo's write-ahead and block-storage lifecycle
- `emptyDir` lifetime compared with PersistentVolumeClaim lifetime
- persistence across restarts
- block creation and compaction
- retention

Expected observation:

- after replacing the Tempo Pod, a previously stored trace remains queryable.

We will discuss object storage here but retain Minikube's local persistent
storage for the exercise.

## Step 5 - Put Alloy in front of Tempo

Change the telemetry path to:

```text
application -> Alloy -> Tempo
```

Alloy will run as a small Kubernetes workload with its own ConfigMap and
Service. The producer will be changed to send OTLP to Alloy rather than directly
to Tempo.

Concepts:

- why applications commonly send to a collection layer
- where batching, filtering, enrichment, retries, and sampling belong
- why Tempo should concentrate on trace storage and querying
- why using OTLP on both connections does not make Alloy and Tempo equivalent

Expected observation:

- the same trace reaches Tempo through the Alloy Service.

Because Alloy is relevant to the intended environment, we will use it here and
briefly compare its configuration with the equivalent vanilla Collector
pipeline.

## Step 6 - Add one attribute-size limit

Configure one understandable default ingestion override and intentionally exceed
it with one oversized span attribute.

Concepts:

- global defaults
- protecting Tempo from unexpectedly large attributes
- accepted spans compared with modified telemetry
- truncation as a different behavior from rejecting or discarding spans
- why operational limits belong in the backend

Expected observation:

- Tempo accepts the trace but truncates the deliberately oversized attribute,
  and the result is visible in the stored trace, logs, and a dedicated metric.

We do not need to explore every available limit. Per-tenant runtime overrides,
rate limiting, and outright span rejection remain later topics.

## Step 7 - Generate metrics from traces

Enable Tempo's metrics generator and add a small Prometheus deployment to
receive the generated metrics. Grafana will then use both Tempo and Prometheus
as data sources.

Concepts:

- span metrics
- RED measurements: rate, errors, and duration
- service graphs
- exemplars and metrics-to-trace navigation
- why derived metrics are stored in Prometheus rather than Tempo

Expected observation:

- traces appear in Tempo while metrics derived from those traces appear in
  Prometheus and Grafana.

This is a larger change, but all of its pieces are required to demonstrate one
coherent feature.

## Step 8 - Build the production configuration map

Finish with an annotated conceptual comparison instead of deploying a complex
production cluster.

Topics:

- monolithic versus microservices mode
- local filesystem storage versus object storage
- handwritten Kubernetes manifests versus the official Tempo Helm charts
- authentication at the Tempo boundary
- multi-tenancy and `X-Scope-OrgID`
- high availability
- retention and compaction
- Tempo's own readiness and metrics endpoints
- which Tempo components can scale independently
- Kafka's role in the current microservices architecture

Current Tempo has monolithic and microservices modes. Monolithic mode runs the
necessary components in one process and does not require Kafka. Tempo 3
microservices mode separates the components for independent scaling and
requires a Kafka-compatible system. The older scalable-monolithic mode has been
removed.

We will not deploy Kafka, distributed object storage, and every Tempo
microservice merely to learn their names. Instead, we will relate the production
components to behavior already observed in the monolithic Minikube exercise.
We can render or inspect the official Helm chart at this point without replacing
the understandable manifests used to build the initial mental model.

## Expected final understanding

At the end of the exercise, we should be able to answer:

1. Where does Tempo receive traces?
2. Which endpoint is used for ingestion and which is used for querying?
3. Where are traces stored?
4. What trace data survives a restart?
5. How long is trace data retained?
6. How does Grafana find and display traces?
7. Where should trace filtering and processing happen?
8. What protects Tempo from excessive ingestion?
9. How are service graphs and span metrics produced?
10. When is monolithic Tempo no longer appropriate?
11. What major components and dependencies appear in a production deployment?
12. Which concerns belong in Kubernetes manifests and which belong in
    `tempo.yaml`?

## Current references

- [Grafana Tempo documentation](https://grafana.com/docs/tempo/latest/)
- [Tempo architecture](https://grafana.com/docs/tempo/latest/introduction/architecture/)
- [Tempo deployment modes](https://grafana.com/docs/tempo/latest/reference-tempo-architecture/deployment-modes/)
- [Deploy Tempo](https://grafana.com/docs/tempo/latest/set-up-for-tracing/setup-tempo/deploy/)
