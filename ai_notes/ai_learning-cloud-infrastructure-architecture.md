# Learning Cloud Infrastructure Architecture

## Context and Motivation

This learning question grew out of designing observability across two home-lab
Kubernetes clusters:

- The Orange Pi (`opi`) cluster runs experimental applications and receives
  load tests.
- The HP (`hp`) cluster is intended to host the central observability stack.
- Metrics, traces, and logs must move from the workload cluster to the
  observability cluster without unnecessarily exposing every internal service.

The home-lab discussion produced an edge-collector architecture: retain Alloy
in the workload cluster, collect telemetry locally, and forward it through one
restricted ingestion endpoint to Prometheus, Tempo, and Loki in the monitoring
cluster.

The same problem was then considered in AWS. With production and observability
EKS clusters in separate VPCs, an internal Network Load Balancer combined with
AWS PrivateLink provides a narrowly scoped ingestion service without creating
general routing between the VPCs. VPC peering is also reasonable when broader
private connectivity is already required.

Those conclusions raised a broader question:

> How can someone develop the mental map needed to combine networking,
> Kubernetes, cloud services, security, reliability, and observability into a
> coherent architecture rather than learning each technology in isolation?

The desired knowledge is the layer above individual products: architecture
patterns, trust boundaries, data flows, failure modes, and tradeoff-driven
composition.

Related design notes:

- `hp-cluster/cross-cluster-observability.md`
- `hp-cluster/aws-production-observability-cluster-architecture.md`

## The Nature of the Skill

There is no single book that completely joins networking, Kubernetes, AWS,
security, reliability, and observability. This skill is broadly known as
solutions architecture.

It consists of three complementary forms of knowledge:

1. **Mechanisms**: understanding what routing, NAT, DNS, load balancers,
   Kubernetes Services, queues, databases, and identity systems actually do.
2. **Patterns**: recognizing reusable arrangements such as gateways, sidecars,
   asynchronous queues, bulkheads, circuit breakers, edge collectors, and
   hub-and-spoke networks.
3. **Judgment**: choosing among valid patterns based on security, reliability,
   cost, complexity, performance, organizational constraints, and expected
   growth.

Architectural judgment is not mysterious intuition. It can be developed
deliberately by studying patterns, understanding the underlying mechanisms,
and repeatedly making and reviewing design decisions.

## Primary Book Recommendation

### Fundamentals of Software Architecture, Second Edition

Mark Richards and Neal Ford's *Fundamentals of Software Architecture, Second
Edition* is the strongest single starting recommendation.

Despite the title's emphasis on software, it teaches the reasoning process
needed for infrastructure and cloud architecture:

- Everything in architecture involves tradeoffs.
- Start from required architectural characteristics.
- Compare alternatives rather than searching for a universally correct design.
- Identify coupling, boundaries, failure modes, and operational consequences.
- Record why a decision was made.
- Revisit decisions when their original constraints change.

It will not prescribe a specific solution such as "use PrivateLink for Alloy."
It teaches the process that leads to such a decision.

Reference:

- [Fundamentals of Software Architecture, Second Edition](https://www.oreilly.com/library/view/fundamentals-of-software/9781098175504/)

## Complementary Books

### Designing Distributed Systems, Second Edition

Brendan Burns develops a vocabulary of reusable distributed-system patterns,
including sidecars, ambassadors, adapters, replicated services, work queues,
event processing, and observability patterns.

It is relatively short and connects the pattern language directly to
containers and Kubernetes.

- [Designing Distributed Systems, Second Edition](https://www.oreilly.com/library/view/designing-distributed-systems/9781098156343/)

### Networking and Kubernetes

James Strong and Vallery Lancey explain the mechanics beneath Kubernetes and
cloud networking:

- Linux packet handling and routing;
- container and overlay networking;
- the Kubernetes network model;
- Services, `ClusterIP`, `NodePort`, and `LoadBalancer`;
- Ingress and service meshes;
- NetworkPolicy; and
- networking in AWS, Google Cloud, and Azure.

The book was published in 2021, so product-specific implementation details
should be checked against current vendor documentation. Its foundational
networking explanations remain useful.

- [Networking and Kubernetes](https://www.oreilly.com/library/view/networking-and-kubernetes/9781492081647/)

### Building Secure and Reliable Systems

Google's *Building Secure and Reliable Systems* treats reliability and
security as properties of the architecture rather than additions made after
deployment. It is valuable for studying:

- trust boundaries;
- least privilege;
- resilience and recovery;
- safe operational practices; and
- how large systems fail.

Google makes the book available to read online without charge.

- [Google SRE books](https://sre.google/books/)
- [Building Secure and Reliable Systems](https://google.github.io/building-secure-and-reliable-systems/)

### Kubernetes Patterns, Second Edition

This is a useful later resource for designing workloads and controllers inside
Kubernetes. It concentrates more on Kubernetes-native application patterns
than on networking between entire cloud environments.

- [Kubernetes Patterns, Second Edition](https://www.oreilly.com/library/view/kubernetes-patterns-2nd/9781098131678/)

## Free Architecture Pattern Catalog

The Microsoft Azure Architecture Center is one of the best organized catalogs
of cloud architecture patterns, even when Azure is not the target platform.

Its Cloud Design Patterns are intended to be broadly applicable to distributed
systems. Each pattern normally describes:

- the problem being addressed;
- the reusable pattern;
- situations in which it applies;
- tradeoffs and limitations; and
- a concrete cloud implementation.

The broader Architecture Center separates several kinds of knowledge that are
often mixed together elsewhere:

- architecture styles;
- cloud design patterns;
- technology decision guides;
- reference architectures; and
- performance antipatterns.

References:

- [Azure Cloud Design Patterns](https://learn.microsoft.com/en-us/azure/architecture/patterns/)
- [Azure Architecture Center](https://learn.microsoft.com/en-us/azure/architecture/)

## AWS-Specific Architecture Resources

### Amazon Builders' Library

Ordinary service documentation explains what a service does. The Amazon
Builders' Library more often explains why Amazon engineers make particular
design decisions and how production systems fail.

Important subjects include:

- timeouts and retries;
- backpressure and load shedding;
- dependency isolation;
- safe deployments;
- availability and redundancy; and
- operational simplicity.

- [Amazon Builders' Library](https://aws.amazon.com/builders-library/)

### AWS Architecture Center

Use the AWS Architecture Center for reference architectures, diagrams, and
technology decision guides. A reference architecture should be treated as a
worked example whose assumptions must be examined, not as a template to copy
without analysis.

- [AWS Architecture Center](https://aws.amazon.com/architecture/)

### AWS Architect Training

AWS currently provides an architect learning path that includes Architecting
on AWS, Well-Architected material, simulations, and hands-on learning.

The Solutions Architect Professional curriculum is valuable after the
fundamentals because its scenario questions require choosing between several
plausible designs under organizational, security, cost, and migration
constraints.

Certification study should be used as architecture practice. Passing an exam
is not by itself evidence of mature architectural judgment.

References:

- [AWS Solutions Architect training](https://aws.amazon.com/training/learn-about/architect/)
- [AWS Certified Solutions Architect Professional exam guide](https://docs.aws.amazon.com/aws-certification/latest/solutions-architect-professional-02/solutions-architect-professional-02.html)

## A Repeatable Architecture Method

The cross-cluster observability solution followed a sequence that can be used
for other infrastructure decisions.

### 1. State the desired outcome

In the motivating example, monitoring workloads needed to move away from the
cluster being load-tested so that the test would not impair the system used to
observe it.

### 2. Identify trust and failure boundaries

The OPI cluster produces telemetry. The HP cluster stores and visualizes it.
Load or failure in OPI should not make Grafana unusable or destroy central
monitoring visibility.

### 3. Draw data flows and their direction

Metrics, logs, and traces all originate in OPI. This suggests outbound
forwarding from OPI rather than granting HP inbound access to every OPI
workload.

### 4. Separate collection from storage

Alloy performs local discovery, collection, batching, filtering, and
forwarding. Prometheus, Loki, and Tempo provide central storage and querying.

### 5. Inspect the actual networking constraints

The two home-lab clusters use overlapping pod and Service networks. This
eliminates simple direct routing between pods as a viable design.

### 6. Minimize the exposed surface

One restricted ingestion gateway is preferable to independently exposing
Prometheus, Tempo, Loki, and every application metrics endpoint.

### 7. Compare patterns by tradeoff

- MetalLB with source restrictions is proportionate to the trusted home lab.
- PrivateLink provides a stronger service boundary for separated AWS VPCs.
- VPC peering is simpler when broader private connectivity is already needed.
- A full multicluster network could work but adds unjustified complexity to
  the current home-lab problem.

### 8. Define when to reconsider the decision

The chosen design might need to change if there are:

- many more source clusters;
- untrusted networks;
- substantially higher ingestion volume;
- regulatory or compliance requirements;
- strict tenant isolation requirements; or
- different recovery and availability objectives.

This process is more valuable than memorizing the final diagram.

## Questions to Ask During Infrastructure Design

### Requirements

- What outcome is the system meant to provide?
- Which requirements are functional and which are qualities such as security,
  availability, latency, scalability, or operability?
- What is explicitly outside the scope?

### Boundaries and data

- Where are the network, identity, administrative, and failure boundaries?
- What data crosses each boundary?
- Who initiates each connection?
- Is the flow pull-based, push-based, synchronous, or asynchronous?
- Does the data contain secrets or personally identifiable information?

### Failure and recovery

- What happens when each dependency is slow, unavailable, or returns a partial
  result?
- Can retries amplify a failure?
- Where is buffering performed and how much data can be lost?
- What is the blast radius of a failure or incorrect configuration?

### Security

- Is the endpoint merely private, or is the client authenticated?
- Can the allowed client access only this service or the entire network?
- Where are encryption, authorization, credential rotation, and audit controls
  applied?

### Operations

- How is the system deployed, upgraded, observed, backed up, and restored?
- Which team owns each component and boundary?
- Is the operational complexity proportionate to the problem?
- What evidence will show that the design is working?

### Cost and evolution

- Which resources incur fixed costs and which scale with traffic?
- What is the simplest design that meets today's requirements?
- Which future change would force a redesign?
- Is that future sufficiently likely to justify complexity now?

## Practical Architecture Exercise

For every meaningful home-lab project, create a one-page architecture note
with the following structure:

```text
Goal
Non-goals
Constraints
Trust boundaries
Data flows
Failure modes
Candidate designs
Tradeoffs
Decision
Conditions that would cause the decision to be revisited
```

Draw two small diagrams:

1. A deployment diagram showing networks, clusters, services, storage, and
   externally reachable endpoints.
2. A data-flow diagram showing who initiates connections and where trust
   boundaries are crossed.

For important decisions, deliberately produce three viable options:

- The simplest acceptable design.
- A more isolated or secure design.
- A more scalable design.

Choose one and explain why the other two are currently unjustified. This
exercise develops architectural judgment because it requires evaluating the
solution against real constraints rather than selecting technology by habit.

## Suggested Learning Sequence

1. Read *Fundamentals of Software Architecture, Second Edition* for the overall
   decision-making method.
2. Read the first six chapters of *Networking and Kubernetes* to build the
   packet, routing, Service, and load-balancer model.
3. Read *Designing Distributed Systems, Second Edition* to develop a reusable
   pattern vocabulary.
4. Read *Building Secure and Reliable Systems* alongside practical work.
5. Study one Azure cloud design pattern and one Amazon Builders' Library
   article each week.
6. Apply each important idea to the two-cluster home lab through small
   architecture decision records and diagrams.
7. Use AWS Solutions Architect scenario questions later as constrained design
   exercises rather than pure examination preparation.

The home lab is well suited to this learning process. It is already large
enough to expose real architectural forces—network boundaries, shared failure
domains, resource contention, telemetry transport, storage, security, and
operational complexity—without requiring a large production organization.

