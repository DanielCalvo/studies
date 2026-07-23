# SRE Practices Beyond Incident Response

On-call, incident response, and postmortems are highly prominent SRE practices. However, several other practices are just as important for interviews and for operating production systems effectively.

## 1. SLOs, SLIs, and Error Budgets

This is arguably the defining SRE practice outside incident management.

You should be comfortable answering:

- What user behavior or outcome are we protecting?
- Which indicator actually measures that outcome?
- What reliability target is appropriate?
- Over what time window should it be measured?
- How much unreliability does the error budget permit?
- What should the team do when the budget is being consumed too quickly?
- Why is 100% reliability usually the wrong target?

### Practice exercise

Take a service, such as an API, web application, or Kubernetes platform, and define:

- Two or three user journeys.
- Corresponding SLIs.
- Proposed SLOs and measurement windows.
- Error-budget calculations.
- A simple policy describing what happens when the budget is healthy or exhausted.

For interviews, expect scenario questions rather than only definitions. For example: *The service meets its availability SLO, but users are complaining. What might be wrong?* A strong answer examines whether the SLI represents the user experience, including latency, correctness, freshness, and dependency failures.

## 2. Monitoring, Observability, and Alert Design

This is extremely prominent in real infrastructure work. The important skill is not merely knowing Prometheus, Grafana, or tracing. It is knowing:

- What should be measured.
- Which signals help diagnose a problem.
- Which conditions deserve to wake a human.
- How to avoid noisy, redundant, or unactionable alerts.
- The difference between symptom-based and cause-based alerts.
- How telemetry moves from application to storage, query, visualization, and notification.

### Practice exercise

Examine a small service and design:

- A dashboard for normal operation.
- Alerts tied to user impact or error-budget burn.
- A troubleshooting dashboard covering requests, errors, latency, and saturation.
- A runbook for each paging alert.

A particularly valuable topic is multi-window, multi-burn-rate alerting. It is not essential to memorize every threshold, but it is important to understand why SLO-based alerting is often better than a rule such as `CPU > 80%`.

## 3. Toil Identification and Reduction

Toil is another distinctly SRE idea, although interviews sometimes neglect it.

Toil is generally work that is:

- Manual.
- Repetitive.
- Automatable.
- Tactical rather than enduring.
- Growing proportionally with the service.

### Practice exercise

Keep a toil register for routine work. Record its frequency, time cost, risk, scalability, and possible automation. Then decide whether each task should be automated, eliminated, delegated, or intentionally retained.

Not every manual task deserves automation. Automation has development and maintenance costs, so frequency, risk, and return on investment matter.

## 4. Capacity Planning and Saturation Management

This appears frequently in actual jobs and system-design interviews:

- How much traffic can the system handle?
- Which resource becomes exhausted first?
- How much headroom is necessary?
- What happens during a traffic spike?
- How do growth, redundancy, and failure scenarios affect capacity requirements?
- When is autoscaling insufficient?

### Practice exercise

Estimate capacity using request rate, concurrency, latency, CPU, memory, storage growth, and network limits. Load-test a small service, identify its bottleneck, and predict what will happen before increasing the load.

Utilization is not the same as capacity. A system running at 50% average CPU might still be close to failure because of peaks, uneven distribution, queues, memory pressure, or a constrained dependency.

## 5. Safe Change and Release Engineering

Changes cause a large proportion of production failures, so SRE work pays close attention to:

- Progressive delivery.
- Canary releases.
- Automated validation.
- Rollback and roll-forward strategies.
- Feature flags.
- Backward-compatible database changes.
- Deployment health signals.
- Reducing change size and blast radius.

### Practice exercise

Design a deployment process that answers:

1. How is the change validated before production?
2. How is it exposed gradually?
3. Which metrics determine success?
4. When does deployment stop automatically?
5. How is the previous state restored?
6. What happens if rollback itself is unsafe?

## 6. Resilience and Failure-Mode Analysis

This is the proactive counterpart to incident response:

- Identifying single points of failure.
- Understanding dependency failure.
- Setting timeouts, retries, and retry budgets.
- Using circuit breakers and load shedding.
- Handling partial failure.
- Preventing retry storms and cascading failure.
- Designing graceful degradation.
- Testing backup and disaster-recovery assumptions.

### Practice exercise

Draw the dependency graph of a service and ask what happens when each component becomes slow, unavailable, inconsistent, or overloaded. Game days and controlled failure experiments are useful, but tabletop exercises also develop the underlying reasoning.

## 7. Automation and Production Engineering

Automation is broader than writing scripts. The SRE concern is creating repeatable, observable, and safe operational systems:

- Idempotent provisioning.
- Configuration management.
- Automated remediation with safeguards.
- CI/CD.
- Infrastructure as code.
- Kubernetes operators and controllers.
- Testing operational code.
- Avoiding automation that amplifies failure.

For interviews, be prepared to explain where automation should stop and require human approval. A powerful automated system also needs bounded scope, validation, auditability, and a safe failure mode.

## 8. Operational Readiness

Before accepting a service into production or on-call ownership, an SRE should examine:

- Ownership and escalation.
- SLOs.
- Dashboards and alerts.
- Runbooks.
- Capacity and scaling.
- Dependency risks.
- Deployment and rollback.
- Backups and restoration.
- Known failure modes.
- Security and access.
- Whether the service can realistically be supported.

### Practice exercise

Create and use a production-readiness checklist for one of your projects. This brings many otherwise abstract SRE topics together.

## Suggested Priority

For a compact curriculum with the greatest interview and workplace value:

1. SLIs, SLOs, and error budgets.
2. Monitoring, observability, and actionable alerting.
3. Capacity, overload, and scaling.
4. Safe releases and change management.
5. Resilience and dependency failure.
6. Toil management and automation.
7. Operational readiness and service ownership.

## A Practical Study Project

On-call and postmortems are highly visible because they happen around dramatic events. SLO design, alert engineering, capacity planning, safe delivery, and toil reduction are quieter practices, but they help prevent the on-call experience from becoming chaotic.

A strong way to study SRE is to maintain one small service and progressively give it:

- User-centered SLOs.
- SLO-based alerts.
- Useful dashboards.
- Capacity estimates and load tests.
- A safe deployment strategy.
- A dependency failure analysis.
- A production-readiness review.
- A toil register.

This exercise teaches considerably more than memorizing isolated definitions and provides concrete examples to discuss during interviews.
