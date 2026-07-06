# Datadog Features and Open Source Monitoring Equivalents

Last researched: 2026-07-06

This is a practical map from Datadog features to open source or self-hostable equivalents for this homelab monitoring stack.

The local baseline is:

- k3s on two arm64 Orange Pi nodes: `opi1` and `opi2`.
- MetalLB for LAN `LoadBalancer` services.
- Prometheus Operator installed from upstream bundle.
- Prometheus at `192.168.1.220`.
- Grafana at `192.168.1.221`.
- kube-state-metrics, node-exporter, Prometheus self-scrape, and Prometheus Operator scrape are already present.
- Main current gap for Kubernetes metrics: kubelet/cAdvisor scraping.

## Sources Used

Datadog feature inventory is based primarily on the current Datadog docs landing page and pricing/product listings:

- Datadog Docs product areas: https://docs.datadoghq.com/
- Datadog pricing/product list: https://www.datadoghq.com/pricing/list/
- Datadog billing docs, useful because they list many billable product families: https://docs.datadoghq.com/account_management/billing/pricing/

Primary open source references used:

- Prometheus: https://prometheus.io/docs/introduction/overview/
- OpenTelemetry: https://opentelemetry.io/docs/
- OpenTelemetry Collector: https://opentelemetry.io/docs/collector/
- Grafana Loki: https://grafana.com/docs/loki/latest/
- Grafana Tempo: https://github.com/grafana/tempo
- Grafana Mimir: https://grafana.com/docs/mimir/latest/
- Thanos: https://thanos.io/
- VictoriaMetrics: https://docs.victoriametrics.com/
- Prometheus blackbox_exporter: https://github.com/prometheus/blackbox_exporter
- Grafana k6: https://grafana.com/docs/k6/latest/
- Uptime Kuma: https://github.com/louislam/uptime-kuma
- Percona Monitoring and Management: https://docs.percona.com/pmm/
- Falco: https://falco.org/docs/
- OpenCost: https://opencost.io/docs/
- Backstage Software Catalog: https://backstage.io/docs/features/software-catalog/
- OpenStatus: https://github.com/openstatusHQ/openstatus
- Grafana OnCall OSS maintenance notice: https://grafana.com/blog/grafana-oncall-maintenance-mode/

## Executive Summary

Datadog is not one tool. It is a hosted observability, security, software delivery, and service management platform. Prometheus plus Grafana covers a real but narrow slice: metrics collection, metrics querying, dashboards, and alerting.

The nearest open source replacement is a stack, not a single product:

| Datadog area | Practical OSS replacement |
| --- | --- |
| Metrics, host dashboards, Kubernetes metrics | Prometheus, Prometheus Operator, kube-state-metrics, node-exporter, kubelet/cAdvisor, Grafana |
| Long-term metrics | VictoriaMetrics single-node first; Thanos or Mimir later |
| Logs | Loki plus Alloy/Promtail or Fluent Bit |
| Traces/APM | OpenTelemetry Collector plus Tempo or Jaeger |
| Synthetic HTTP/TCP/DNS checks | blackbox_exporter, Uptime Kuma |
| Browser/API journey tests | k6; Playwright plus Prometheus push/custom exporter |
| Database monitoring | postgres_exporter, mysqld_exporter, Percona PMM |
| Security runtime signals | Falco |
| Vulnerability scanning | Trivy, Grype, kube-bench, kube-score, kube-linter |
| Cost monitoring | OpenCost |
| Incident/on-call/status pages | Alertmanager, ntfy/Gotify, OpenStatus, maybe commercial paging for serious alerts |
| Service catalog | Backstage, or a simple Markdown catalog for a homelab |
| CI/test visibility | CI exporter scripts, OpenTelemetry traces, Allure Report, ReportPortal, Grafana dashboards |

For this cluster, the recommended learning path is:

1. Finish Kubernetes metrics: scrape kubelet/cAdvisor.
2. Add Alertmanager and a few concrete alerts.
3. Add Loki for pod logs.
4. Add blackbox_exporter for synthetic HTTP/TCP checks.
5. Add OpenTelemetry Collector and Tempo for one toy app.
6. Add VictoriaMetrics only when Prometheus retention becomes a real limitation.
7. Add Falco and Trivy after the observability basics are comfortable.

## Adoption Difficulty Scale

| Level | Meaning |
| --- | --- |
| Easy | One Helm chart or a few manifests; low operational burden. |
| Medium | Multiple components or careful instrumentation; still reasonable in a homelab. |
| Hard | Significant architecture, storage, scaling, or ongoing tuning. |
| Poor homelab fit | Possible, but the feature is mainly valuable in large organizations or cloud estates. |

## Current Local Coverage

| Capability | Current state |
| --- | --- |
| Node OS metrics | Covered by node-exporter. |
| Kubernetes object state | Covered by kube-state-metrics. |
| Prometheus health | Covered by Prometheus self-scrape. |
| Prometheus Operator health | Covered by ServiceMonitor. |
| Dashboards | Grafana with pinned kube-state-metrics, node-exporter, and Prometheus dashboards. |
| Pod/container CPU and memory | Not yet covered directly; add kubelet/cAdvisor. |
| Logs | Not covered. |
| Traces/APM | Not covered. |
| Synthetic checks | Not covered. |
| Alert routing | Not covered unless Alertmanager exists outside this repo. |
| Long-term metrics retention | Not covered; current Prometheus has ephemeral storage unless changed. |
| Database query monitoring | Not covered. |
| Security monitoring | Not covered. |

## Feature-by-Feature Map

### Infrastructure Monitoring

Datadog feature:

- Host inventory.
- Host map.
- CPU, memory, disk, filesystem, network, process, and uptime telemetry.
- Tags for environment, service, team, region, host, cluster, and other dimensions.
- Integrations for common host services.

OSS equivalents:

- Prometheus.
- node-exporter.
- Grafana dashboards.
- Process exporter if process-level visibility is needed: https://github.com/ncabatoff/process-exporter
- Netdata for a polished per-host live UI: https://github.com/netdata/netdata

Adoption difficulty: Easy.

Homelab notes:

- You already have node-exporter.
- The Datadog-style host dashboard workflow maps well to Grafana variables. Use a `node` or `instance` variable with multi-select enabled and PromQL selectors using `=~"$node"`.
- Datadog's host list is more polished out of the box. Grafana can match the visibility, but you build or tune the dashboard yourself.
- For two nodes, Prometheus plus node-exporter is the correct lightweight choice.

Recommended next work:

- Create a custom `Homelab Nodes` dashboard instead of heavily editing the imported `Node Exporter Full` dashboard.
- Add dashboard variables for `job`, `nodename`, and `instance`, with multi-select.

### Metrics Platform

Datadog feature:

- Metrics ingestion, storage, query, dashboards, monitors, tags, custom metrics, distributions, events, and derived metrics.

OSS equivalents:

- Prometheus for scrape, TSDB, PromQL, rules, and alerts.
- Alertmanager for notification routing.
- Grafana for dashboards and exploration.
- Recording rules for precomputed metrics.
- Pushgateway for short-lived batch jobs, with care.
- OpenTelemetry Collector for metrics pipeline use cases.

Adoption difficulty: Easy to Medium.

Caveats:

- Prometheus cardinality problems are your responsibility. Datadog also charges for custom metrics/cardinality, but it hides more operational details.
- Prometheus local storage is simple but not highly available.
- Pushgateway is often misused. It is best for service-level batch job outcomes, not for per-run high-cardinality data.
- Deleting accidentally ingested high-cardinality series is possible with Prometheus admin APIs only if enabled, but prevention is the real answer: relabeling, metric naming discipline, and cardinality review.

Homelab recommendation:

- Stay with Prometheus until you have a concrete retention or scale problem.
- Add recording rules once dashboards become slow or queries are repeated everywhere.

### Kubernetes and Container Monitoring

Datadog feature:

- Kubernetes cluster overview.
- Pod, deployment, daemonset, job, cronjob, service, and node views.
- Container CPU, memory, filesystem, network, restart, and image metadata.
- Kubernetes events and object state.

OSS equivalents:

- kube-state-metrics for Kubernetes desired/current object state.
- kubelet/cAdvisor for pod/container resource usage.
- node-exporter for node OS metrics.
- Kubernetes events exporter: https://github.com/resmoio/kubernetes-event-exporter
- Grafana Kubernetes dashboards.

Adoption difficulty: Easy to Medium.

Homelab notes:

- You already have kube-state-metrics and node-exporter.
- The biggest missing piece is kubelet/cAdvisor scraping. Without it, you can see whether Kubernetes objects are healthy and whether nodes are healthy, but not good per-pod/container resource usage.
- metrics-server is not a Prometheus replacement. It feeds `kubectl top` and autoscaling-style APIs; it is not meant for long-term querying or dashboards.

Recommended next work:

- Add a kubelet ServiceMonitor for `/metrics`, `/metrics/cadvisor`, and possibly `/metrics/resource`.
- Add a Kubernetes pods/resources dashboard after cAdvisor data is present.

### Kubernetes Autoscaling

Datadog feature:

- Autoscaling recommendations and control loops based on Datadog telemetry.
- Kubernetes workload scaling based on richer metrics than basic CPU/memory.

OSS equivalents:

- Kubernetes HorizontalPodAutoscaler.
- Kubernetes VerticalPodAutoscaler.
- KEDA for event-driven autoscaling: https://keda.sh/
- Prometheus Adapter for custom metrics autoscaling.

Adoption difficulty: Medium.

Caveats:

- Autoscaling is only useful when workloads actually have variable load and spare capacity.
- On a two-node Orange Pi cluster, autoscaling is more educational than operational.
- HPA needs good metrics and resource requests. Without those, it produces misleading behavior.

Homelab recommendation:

- Try HPA or KEDA with a toy app after kubelet/cAdvisor metrics are in place.
- Do not add autoscaling to core monitoring components until the cluster is more stable.

### Network Monitoring

Datadog feature:

- Cloud Network Monitoring.
- Network device monitoring.
- NetFlow monitoring.
- Network path monitoring.
- Service-to-service traffic visibility.

OSS equivalents:

- node-exporter network counters for host interfaces.
- Cilium Hubble if using Cilium as CNI.
- Pixie for Kubernetes observability through eBPF: https://px.dev/
- ntopng for network traffic analysis: https://www.ntop.org/products/traffic-analysis/ntop/
- SNMP exporter for network devices: https://github.com/prometheus/snmp_exporter
- blackbox_exporter for path-level reachability checks.

Adoption difficulty: Medium to Hard.

Caveats:

- Datadog's network product is much more integrated than the common OSS equivalents.
- Deep flow visibility usually means eBPF, CNI-specific tooling, or device-level flow export. That is a bigger operational jump than adding exporters.
- k3s often uses flannel by default. If this cluster is using flannel, you will not get Cilium Hubble without changing the networking stack.

Homelab recommendation:

- Start with simple checks: node network counters and blackbox probes.
- Do not change CNI just to get network dashboards unless network observability becomes the study topic.

### Synthetic Monitoring: HTTP, TCP, DNS, ICMP, gRPC

Datadog feature:

- API tests.
- HTTP tests.
- SSL/TLS checks.
- DNS, TCP, UDP, ICMP-style availability checks.
- Private locations.
- Alerting on uptime and latency.

OSS equivalents:

- Prometheus blackbox_exporter for HTTP, HTTPS, DNS, TCP, ICMP, and gRPC probes.
- Uptime Kuma for a friendly self-hosted uptime UI and notifications.
- OpenStatus for status pages plus uptime/API monitoring.

Adoption difficulty: Easy.

Caveats:

- A synthetic monitor inside the same cluster only proves cluster-local reachability. It does not prove the service works from the public internet.
- For a homelab, internal checks are still useful for learning and for detecting broken services.
- For real external uptime, run probes from outside the LAN or use a cheap external service.

Homelab recommendation:

- Install blackbox_exporter first because it integrates naturally with Prometheus and Alertmanager.
- Add Uptime Kuma if you want the Datadog/UptimeRobot-style friendly UI.

Good first checks:

- Prometheus: `http://192.168.1.220/-/ready`
- Grafana: `http://192.168.1.221/api/health`
- Registry: `http://192.168.1.242:5000/v2/`
- Traefik LAN IP: `192.168.1.222`

### Browser Synthetics and User Journeys

Datadog feature:

- Browser tests.
- Multi-step user journeys.
- Screenshot and failure artifacts.
- CI/CD synthetic test integration.

OSS equivalents:

- Grafana k6 browser API.
- Playwright.
- Selenium.
- Checkly is strong here but is not primarily an OSS self-hosted replacement.

Adoption difficulty: Medium.

Caveats:

- Browser tests are heavier than HTTP probes. On arm64 Orange Pi nodes, keep concurrency low.
- The hard part is not running a browser test once; it is storing results, screenshots, timings, and alert history cleanly.
- k6 is a better fit if you also want performance/load testing.
- Playwright is a better fit if you already write browser tests as part of app development.

Homelab recommendation:

- Use blackbox_exporter for basic uptime.
- Use k6 or Playwright for one or two critical flows only.
- Export pass/fail and duration into Prometheus using a small exporter or Pushgateway-style batch result.

### Mobile App Testing

Datadog feature:

- Mobile application journey testing and mobile-specific user flow validation.

OSS equivalents:

- Appium.
- Maestro.
- Detox for React Native.
- Android emulator/iOS simulator tests in CI.

Adoption difficulty: Medium to Hard.

Caveats:

- Mobile testing infrastructure is much heavier than HTTP checks.
- iOS testing requires Apple tooling and is awkward for a Linux homelab.
- Android emulator workloads may be too heavy for Orange Pi nodes.

Homelab recommendation:

- Poor fit unless you specifically want to study mobile testing.
- For a homelab, run mobile tests on a workstation or CI runner, then export summary metrics.

### Real User Monitoring

Datadog feature:

- Browser RUM.
- Mobile RUM.
- Page load timings.
- User actions.
- Frontend errors.
- Session correlation with traces and logs.

OSS equivalents:

- OpenTelemetry browser instrumentation, where suitable.
- Grafana Faro: https://github.com/grafana/faro
- OpenReplay for session replay and frontend monitoring: https://github.com/openreplay/openreplay
- Highlight.io for self-hostable product/session/error observability: https://github.com/highlight/highlight
- PostHog for product analytics and session replay: https://github.com/PostHog/posthog

Adoption difficulty: Medium to Hard.

Caveats:

- RUM can collect sensitive user data. Datadog gives you productized privacy controls; self-hosting means you own that problem.
- RUM is not very useful until you have a real browser app and enough sessions.
- Session replay storage can grow quickly.

Homelab recommendation:

- Skip until you have a web app worth observing.
- If you build a test app, try Grafana Faro first because it fits the Grafana ecosystem.

### Product Analytics

Datadog feature:

- Product usage analytics, funnels, cohorts, user behavior, and business/product events.

OSS equivalents:

- PostHog.
- Matomo.
- Plausible Community Edition.
- Umami.

Adoption difficulty: Medium.

Caveats:

- Product analytics is adjacent to observability but has different goals.
- It can introduce privacy and consent requirements.
- It is only useful when you have real users or a realistic application.

Homelab recommendation:

- Skip for infrastructure monitoring.
- If you build a real web app, PostHog is the most complete self-hosted product analytics option.

### Experiments

Datadog feature:

- Experimentation and A/B testing tied to product analytics and feature flags.

OSS equivalents:

- GrowthBook.
- PostHog experiments.
- Unleash strategy variants.

Adoption difficulty: Medium.

Caveats:

- Needs traffic volume and a product hypothesis.
- Not useful for cluster health monitoring.

Homelab recommendation:

- Skip unless you build an app where experimentation itself is the topic.

### Session Replay

Datadog feature:

- Visual replay of user sessions.
- Replay errors and frontend performance issues.
- Link sessions to RUM, logs, and traces.

OSS equivalents:

- OpenReplay.
- Highlight.io.
- PostHog.

Adoption difficulty: Medium to Hard.

Caveats:

- Privacy, masking, and retention matter.
- This is storage-heavy compared with metrics.
- It is overkill for infrastructure-only homelab monitoring.

Homelab recommendation:

- Treat this as an application observability experiment, not a core cluster monitoring component.

### APM and Distributed Tracing

Datadog feature:

- Distributed traces.
- Service pages.
- Request/error/latency metrics.
- Trace search and retention controls.
- Correlation with logs, RUM, synthetics, profiles, and database monitoring.

OSS equivalents:

- OpenTelemetry SDKs and auto-instrumentation.
- OpenTelemetry Collector.
- Grafana Tempo.
- Jaeger.
- Grafana dashboards and exemplars.

Adoption difficulty: Medium.

Caveats:

- Instrumentation is the main work. Installing Tempo alone gives you nothing if apps do not emit traces.
- Tempo is storage-efficient, but querying by arbitrary attributes is not the same experience as Datadog unless you plan the pipeline and indexing path carefully.
- Jaeger has a more traditional tracing UI. Tempo integrates better with Grafana metrics/logs.

Homelab recommendation:

- Use OpenTelemetry from the start. Do not use vendor-specific Datadog instrumentation for new experiments.
- Use Tempo if you want Grafana-native traces.
- Build one small demo app and trace HTTP requests through it before trying to instrument everything.

### Universal Service Monitoring

Datadog feature:

- Service discovery and basic service-level telemetry without code changes.

OSS equivalents:

- Pixie.
- Cilium Hubble, if using Cilium.
- Service mesh telemetry such as Istio, Linkerd, or Envoy metrics.
- eBPF-based tools.

Adoption difficulty: Hard.

Caveats:

- This is one of Datadog's more productized convenience areas.
- OSS options tend to be either CNI-specific, service-mesh-specific, or eBPF-heavy.
- In a small homelab, explicit instrumentation is usually more educational and less disruptive.

Homelab recommendation:

- Skip initially.
- Prefer OpenTelemetry instrumentation for apps you control.

### Continuous Profiler

Datadog feature:

- CPU, memory, allocation, lock, wall-time, and other profiles over time.
- Correlation with services, traces, and deployments.

OSS equivalents:

- Parca: https://www.parca.dev/
- Pyroscope: https://github.com/grafana/pyroscope
- Language-native profilers such as Go `pprof`, Java Flight Recorder, Python py-spy.

Adoption difficulty: Medium.

Caveats:

- Profiling is extremely useful when you have real app performance problems.
- It is not usually the first thing to install for infrastructure monitoring.
- Continuous profiling has overhead and storage implications.

Homelab recommendation:

- Try Pyroscope or Parca with one toy Go/Python app after traces are working.

### Dynamic Instrumentation

Datadog feature:

- Add logs, metrics, spans, and snapshots dynamically to running applications without redeploying.

OSS equivalents:

- Very limited direct OSS equivalent.
- Some language debuggers and eBPF tools can approximate pieces.
- OpenTelemetry helps with planned instrumentation, not arbitrary dynamic probes.

Adoption difficulty: Hard.

Caveats:

- This is a sophisticated commercial feature.
- The safe open source equivalent is usually better logging/tracing practices, not live production mutation.

Homelab recommendation:

- Skip.

### Error Tracking

Datadog feature:

- Group backend, browser, mobile, and log errors into issues.
- Regression detection.
- Suspected commits and ownership.
- Correlation with traces, sessions, and deployments.

OSS equivalents:

- GlitchTip: https://glitchtip.com/
- Sentry self-hosted, though operationally heavy.
- Highlight.io.
- OpenReplay for frontend-focused errors.
- Loki alerting for simpler log-based error spikes.

Adoption difficulty: Medium.

Caveats:

- Sentry self-hosted is not lightweight.
- Log alerts catch spikes but do not replace issue grouping, stack traces, releases, and suspect commits.

Homelab recommendation:

- For a lightweight app, start with structured logs in Loki plus a simple error-rate dashboard.
- Try GlitchTip if you want Sentry-style issue tracking without a large stack.

### Log Management

Datadog feature:

- Log ingestion.
- Pipelines and parsing.
- Search.
- Live tail.
- Indexing and retention controls.
- Log-based metrics.
- Archives.
- Correlation with traces and infrastructure.

OSS equivalents:

- Grafana Loki.
- Grafana Alloy or Promtail for collection.
- Fluent Bit as a lightweight collector.
- Vector as a stronger pipeline/router: https://vector.dev/
- OpenSearch for full-text search-heavy logging.
- ClickHouse-based stacks for high-volume logs, such as SigNoz or custom pipelines.

Adoption difficulty: Medium.

Caveats:

- Loki intentionally indexes labels, not full log text. This keeps it cheaper but changes how you model logs.
- Bad labels in Loki are as painful as bad labels in Prometheus. Do not use request IDs, user IDs, or high-cardinality fields as labels.
- OpenSearch gives more full-text search power but costs more CPU/memory.

Homelab recommendation:

- Use Loki, not OpenSearch, for this small arm64 cluster.
- Use Alloy or Promtail as a DaemonSet to collect pod logs.
- Keep retention short at first.

### BYOC / CloudPrem Log Management

Datadog feature:

- Datadog has product positioning around bringing log management closer to customer-controlled infrastructure in some offerings.
- The key capability is separating collection, pipeline, and control from where log data is stored and processed.

OSS equivalents:

- Loki self-hosted.
- OpenSearch self-hosted.
- ClickHouse-based log stores.
- Vector or OpenTelemetry Collector as routing layers.

Adoption difficulty: Medium to Hard.

Caveats:

- Self-hosting logs means you own disk, retention, upgrades, backup, query performance, and access control.
- This is basically the normal OSS model, not a special feature.

Homelab recommendation:

- Loki is the right version of this idea for the cluster.

### Observability Pipelines

Datadog feature:

- Collect, transform, redact, route, sample, and control telemetry before it reaches storage.

OSS equivalents:

- OpenTelemetry Collector.
- Vector.
- Fluent Bit.
- Logstash, though heavier.

Adoption difficulty: Medium.

Caveats:

- Pipelines are operationally powerful but add another thing that can silently drop data.
- Keep config in Git and add pipeline health dashboards.

Homelab recommendation:

- Use OpenTelemetry Collector for traces and maybe metrics.
- Use Fluent Bit or Alloy for logs.
- Avoid building an elaborate pipeline until you have logs and traces flowing.

### Sensitive Data Scanner

Datadog feature:

- Detect and redact PII, API keys, tokens, and credit-card-like data in telemetry.

OSS equivalents:

- Vector remap transforms.
- OpenTelemetry Collector processors, especially transform/filter/redaction-style processors.
- Fluent Bit filters.
- detect-secrets and gitleaks for source repositories.

Adoption difficulty: Medium.

Caveats:

- Datadog's scanner is productized across telemetry. OSS equivalents are pipeline-specific and require careful rules.
- Redaction should happen before data leaves the workload or node when possible.

Homelab recommendation:

- Do basic prevention first: structured logs, no secrets in logs, and short retention.
- Add pipeline redaction when log collection is installed.

### Audit Trail

Datadog feature:

- Audit who changed monitors, dashboards, roles, integrations, and org settings.

OSS equivalents:

- Kubernetes audit logs.
- Grafana audit logs require Enterprise for some features; OSS has normal server logs.
- Git history for GitOps-managed config.
- Loki/OpenSearch for storing audit events.

Adoption difficulty: Medium.

Caveats:

- Datadog audits Datadog itself. In OSS, each tool has its own audit story.
- GitOps is the simplest substitute for configuration auditability.

Homelab recommendation:

- Keep manifests and Helm values in Git.
- If you later add Loki, collect Kubernetes API audit logs only if you are ready for the volume and configuration complexity.

### Database Monitoring

Datadog feature:

- Database health metrics.
- Query performance.
- Query samples.
- Explain plans.
- Wait events and connection metrics.
- Correlation with traces.

OSS equivalents:

- Percona Monitoring and Management for MySQL, PostgreSQL, and MongoDB.
- postgres_exporter.
- mysqld_exporter.
- mongodb_exporter.
- pg_stat_statements plus Grafana dashboards.
- pgBadger for PostgreSQL log analysis.

Adoption difficulty: Medium.

Caveats:

- Basic DB metrics are easy. Query analytics is more invasive and database-specific.
- PMM is a strong self-hosted equivalent, but it is a bigger component than a simple exporter.
- Query text can contain sensitive data.

Homelab recommendation:

- If you run PostgreSQL, start with postgres_exporter and `pg_stat_statements`.
- Use PMM when you specifically want Datadog DBM-like query dashboards.

### Data Streams Monitoring

Datadog feature:

- Track Kafka and streaming pipeline latency, throughput, backlog, and service dependencies.

OSS equivalents:

- Kafka exporter.
- Burrow for Kafka consumer lag.
- OpenTelemetry traces/metrics around producers and consumers.
- Grafana dashboards.

Adoption difficulty: Medium.

Caveats:

- Useful only if you run Kafka, RabbitMQ, NATS, Redis streams, or similar.
- End-to-end stream lineage usually requires instrumentation, not just broker metrics.

Homelab recommendation:

- Skip unless you add Kafka/NATS/RabbitMQ as a study topic.

### Data Observability / Quality Monitoring

Datadog feature:

- Data quality, freshness, volume, schema, lineage, and anomaly monitoring for data pipelines.

OSS equivalents:

- Great Expectations.
- Soda Core.
- OpenLineage plus Marquez.
- dbt tests and exposures.

Adoption difficulty: Medium to Hard.

Caveats:

- This is data-platform monitoring, not Kubernetes monitoring.
- It becomes useful when you have scheduled ETL/ELT pipelines and downstream consumers.

Homelab recommendation:

- Skip for now.
- If you create a toy batch pipeline, use dbt tests or Great Expectations and export pass/fail metrics to Prometheus.

### Quality Monitoring

Datadog feature:

- Datadog's product list separates quality monitoring from broader data observability.
- In practice this overlaps with data quality, test quality, CI quality, and release quality signals.

OSS equivalents:

- Great Expectations or Soda Core for data quality checks.
- ReportPortal or Allure for test quality.
- Prometheus/Grafana for release health metrics.
- OpenTelemetry events/spans for pipeline quality signals.

Adoption difficulty: Medium.

Caveats:

- This is not one well-defined OSS category.
- You need to decide what "quality" means for the system: data freshness, failed tests, bad deploys, API errors, or user impact.

Homelab recommendation:

- Treat this as a future app/platform exercise, not a base monitoring component.

### Jobs Monitoring

Datadog feature:

- Monitor scheduled jobs, batch jobs, and workflow outcomes.
- Detect missing runs, failures, duration changes, and retries.

OSS equivalents:

- kube-state-metrics for Kubernetes Job/CronJob object state.
- Prometheus alerts on `kube_job_status_failed`, `kube_cronjob_next_schedule_time`, and related metrics.
- Pushgateway for final batch job outcome metrics.
- Healthchecks.io self-hosted alternative: https://github.com/healthchecks/healthchecks
- Cronicle or Airflow/Argo Workflows metrics for more advanced schedulers.

Adoption difficulty: Easy to Medium.

Caveats:

- Missing-run detection is subtle. A failed job emits state; a job that never ran may only be visible through expected schedule logic.
- Pushgateway metrics need lifecycle discipline so stale success metrics do not lie.

Homelab recommendation:

- This is a good study project because you already noted interest in batch job failure and absent metrics.
- Start with one CronJob and alert on failure plus missing successful completion in a time window.

### Serverless Monitoring

Datadog feature:

- AWS Lambda and other serverless runtime monitoring.
- Cold starts, invocations, errors, traces, and logs.

OSS equivalents:

- OpenTelemetry instrumentation.
- Cloud provider metrics exporters.
- OpenFaaS metrics if using OpenFaaS.
- Knative metrics if using Knative.

Adoption difficulty: Poor homelab fit.

Caveats:

- This cluster is k3s, not a serverless-heavy environment.

Homelab recommendation:

- Skip.

### Cloud Cost Management

Datadog feature:

- Cloud spend, allocation, showback/chargeback, and cost visibility correlated with observability tags.

OSS equivalents:

- OpenCost.
- Kubecost free tier, though not fully OSS in all editions.
- Cloud provider billing exports plus Grafana dashboards.

Adoption difficulty: Easy to Medium.

Caveats:

- In a home lab, cloud cost is mostly irrelevant unless you also monitor cloud resources.
- OpenCost can still teach Kubernetes allocation concepts using custom pricing.

Homelab recommendation:

- Optional. Useful if you want namespace/pod cost allocation practice, not because the Orange Pis need billing.

### Storage Management

Datadog feature:

- Cloud storage cost, usage, data freshness, and optimization.

OSS equivalents:

- Cloud provider exporters and billing data.
- Prometheus filesystem metrics.
- MinIO metrics if using MinIO.
- S3 inventory/metrics if using S3-compatible object storage.

Adoption difficulty: Medium.

Caveats:

- Datadog's feature is cloud-storage-oriented.
- For this homelab, node filesystem and PVC monitoring are more relevant.

Homelab recommendation:

- Use node-exporter filesystem metrics and kubelet volume metrics after adding kubelet scraping.
- Add MinIO metrics only if you deploy MinIO for object storage.

### Cloudcraft / Infrastructure Diagrams

Datadog feature:

- Cloud architecture visualization and diagrams.

OSS equivalents:

- Diagrams as code: https://diagrams.mingrammer.com/
- Structurizr Lite.
- Mermaid diagrams in Markdown.
- Backstage plugins for catalog visualization.

Adoption difficulty: Easy.

Caveats:

- No direct need for a small two-node cluster.

Homelab recommendation:

- Use Mermaid in Markdown if you want architecture notes.

### Cloud SIEM

Datadog feature:

- Security log ingestion.
- Detection rules.
- Signals.
- Investigation workflows.
- Cloud and on-prem telemetry correlation.

OSS equivalents:

- Wazuh.
- OpenSearch Security Analytics.
- Sigma rules plus a log backend.
- Falco events routed into Loki/OpenSearch.

Adoption difficulty: Hard.

Caveats:

- SIEM is a program, not just a tool. You need log sources, detections, triage process, retention, and tuning.
- Wazuh is comprehensive but heavier than the current cluster deserves.

Homelab recommendation:

- Start with Falco plus Loki before attempting SIEM.
- Use this as a later security study track.

### Cloud Security Management / Posture

Datadog feature:

- Misconfiguration detection.
- Cloud security posture management.
- Kubernetes security posture.
- Identity risk assessment.
- Vulnerability visibility.
- Compliance frameworks.

OSS equivalents:

- kube-bench.
- kube-score.
- kube-linter.
- Polaris.
- Trivy config scanning.
- Checkov.
- Prowler for AWS.
- ScoutSuite for cloud accounts.

Adoption difficulty: Medium.

Caveats:

- Posture tools produce lots of findings. You need a triage habit or they become noise.
- For a trusted LAN homelab, do not blindly apply enterprise hardening if it blocks experimentation.

Homelab recommendation:

- Run kube-bench and Trivy as periodic manual checks first.
- Keep findings as notes, not a wall of alerts.

### Workload Protection / Runtime Security

Datadog feature:

- Runtime threat detection on hosts, containers, and Kubernetes.
- Process, file, network, and syscall activity monitoring.

OSS equivalents:

- Falco.
- Tetragon, especially with Cilium/eBPF environments.
- Wazuh agent for host security monitoring.

Adoption difficulty: Medium.

Caveats:

- Runtime security can be noisy.
- Kernel/eBPF support on Orange Pi arm64 should be checked before assuming every feature works.
- Falco is the best first option because it is CNCF and Kubernetes-native.

Homelab recommendation:

- Try Falco after logs are in place, so alerts can be stored and explored.

### Vulnerability Management

Datadog feature:

- Host, container image, dependency, and application vulnerability visibility.

OSS equivalents:

- Trivy.
- Grype.
- Syft for SBOMs.
- Dependency-Track.
- Renovate for dependency update automation.

Adoption difficulty: Easy to Medium.

Caveats:

- Scanning images is easy. Turning CVE noise into meaningful remediation is harder.
- For arm64 images, make sure you scan the actual image variant you run.

Homelab recommendation:

- Add Trivy scans for locally built images pushed to `192.168.1.242:5000`.
- Do not alert on every CVE at first; report and review manually.

### App and API Protection

Datadog feature:

- Runtime app/API threat detection, attack detection, and sometimes blocking.

OSS equivalents:

- ModSecurity with OWASP CRS.
- Coraza WAF.
- NGINX/Traefik middleware plus logs and alerts.
- Falco for runtime signals.

Adoption difficulty: Medium to Hard.

Caveats:

- WAFs can break apps and require tuning.
- For internal homelab services, this is usually less important than basic network exposure control.

Homelab recommendation:

- Skip until you expose services publicly.

### Code Security, SAST, IAST, SCA, Secret Scanning, IaC Security

Datadog feature:

- Static code analysis.
- Runtime code analysis.
- Software composition analysis.
- Secret scanning.
- Infrastructure-as-code scanning.
- Pull request comments and gates.

OSS equivalents:

- Semgrep for SAST.
- Trivy for filesystem, image, dependency, and IaC scanning.
- Gitleaks for secrets.
- detect-secrets.
- Checkov/tfsec for IaC.
- OWASP Dependency-Check.
- Dependency-Track for dependency inventory.

Adoption difficulty: Easy to Medium.

Caveats:

- CI integration is straightforward; tuning false positives is the real work.
- IAST is less mature as a simple OSS drop-in than SAST/SCA/secret scanning.

Homelab recommendation:

- Use Trivy and Gitleaks locally or in CI first.
- Add Semgrep when there is app code worth scanning.

### LLM Observability

Datadog feature:

- Trace, monitor, evaluate, and secure LLM applications.

OSS equivalents:

- OpenTelemetry instrumentation for LLM calls.
- Langfuse.
- OpenLLMetry.
- Phoenix by Arize.

Adoption difficulty: Medium.

Caveats:

- Only useful if you are building LLM apps.
- Prompt/response logging can contain sensitive data.

Homelab recommendation:

- Skip unless an LLM app becomes part of the lab.

### Bits AI, MCP Server, and AI Assistants

Datadog feature:

- AI assistants for querying telemetry, summarizing incidents, supporting SRE/security workflows, and exposing observability context to tools through an MCP server.

OSS equivalents:

- Grafana has emerging assistant-style features, but full parity is not an OSS baseline assumption.
- Custom MCP servers against Prometheus, Loki, Tempo, Kubernetes, and Grafana APIs.
- RAG over runbooks and Markdown notes.
- LLM-assisted shell/runbook workflows outside the monitoring stack.

Adoption difficulty: Medium to Hard.

Caveats:

- The hard part is permissions, safety, and high-quality context, not just connecting an LLM to an API.
- AI summaries are useful only if the underlying telemetry and ownership data are good.

Homelab recommendation:

- Keep good Markdown runbooks and service notes first.
- Later, a small read-only helper for Prometheus and Kubernetes would be a good experiment.

### Watchdog / Anomaly Detection

Datadog feature:

- Automated anomaly, outlier, and insight surfacing across metrics/logs/APM.

OSS equivalents:

- Prometheus alerting rules.
- Grafana alerting.
- Grafana machine learning is mostly a cloud/enterprise-style area, not a simple OSS replacement.
- VictoriaMetrics anomaly detection tooling exists but is extra architecture.
- Custom PromQL anomaly patterns, such as z-score or week-over-week comparisons.

Adoption difficulty: Medium.

Caveats:

- Datadog's value here is productized correlation and automatic surfacing.
- Prometheus can express anomaly-like alerts, but you write and maintain them.

Homelab recommendation:

- Start with simple threshold and absent alerts.
- Add anomaly-style rules only after baseline metrics are stable.

### Dashboards

Datadog feature:

- Dashboards, template variables, widgets, notebooks, sharing, and team views.

OSS equivalents:

- Grafana.
- Grafana dashboard provisioning.
- Jsonnet/grafonnet for dashboard-as-code.

Adoption difficulty: Easy.

Caveats:

- Imported dashboards are useful but often too generic.
- The best Datadog-like experience comes from custom dashboards around your actual workflows.

Homelab recommendation:

- Build small custom dashboards:
  - `Homelab Nodes`
  - `Kubernetes Workloads`
  - `Prometheus Health`
  - `Synthetic Checks`
  - `Logs Overview` after Loki exists

### Monitors and Alerting

Datadog feature:

- Metric, log, trace, synthetic, RUM, process, integration, composite, anomaly, forecast, and SLO monitors.
- Downtimes, notification routing, and integrations.

OSS equivalents:

- Prometheus alerting rules.
- Alertmanager.
- Grafana Alerting.
- Loki ruler for log alerts.
- blackbox_exporter plus Prometheus for synthetic alerts.
- ntfy, Gotify, email, Slack/Discord webhooks, Matrix, or Telegram for notifications.

Adoption difficulty: Medium.

Caveats:

- Datadog makes monitor creation very approachable. OSS alerting is more configuration-oriented.
- Alertmanager is reliable, but inhibition, grouping, and routing need deliberate design.
- Grafana Alerting is approachable from the UI, but Prometheus rules in Git are easier to version and review.

Homelab recommendation:

- Add Alertmanager.
- Keep critical alerts few:
  - node down
  - Prometheus target down
  - disk almost full
  - pod crash looping
  - CronJob missing/failing
  - synthetic check failing

### SLOs

Datadog feature:

- Define and track service level objectives.
- Error budget views.
- Burn-rate alerts.

OSS equivalents:

- Prometheus recording and alerting rules.
- Sloth for generating SLO rules: https://github.com/slok/sloth
- Pyrra: https://github.com/pyrra-dev/pyrra
- Grafana SLO dashboards.

Adoption difficulty: Medium.

Caveats:

- SLOs require good SLIs. Without request metrics or synthetic checks, there is not much to calculate.
- For infrastructure-only monitoring, SLOs are less useful than service uptime and latency.

Homelab recommendation:

- Add SLOs after blackbox checks and at least one instrumented app exist.

### Incident Management

Datadog feature:

- Incident declaration, severity, timeline, responders, follow-ups, postmortems, Slack/Teams integration.

OSS equivalents:

- GitHub issues plus Markdown postmortems for simple use.
- FireHydrant has some open tooling but is mainly SaaS.
- Grafana OnCall OSS existed, but Grafana placed it into maintenance mode in 2025 and archived it on 2026-03-24.
- Alerta for alert console and deduplication: https://github.com/alerta/alerta
- OpenStatus for incidents/status pages.

Adoption difficulty: Medium.

Caveats:

- OSS incident management is weaker than Datadog/PagerDuty/Opsgenie-style products.
- For real paging, phone/SMS/push reliability matters. Self-hosted paging from the same cluster that is broken is a bad failure mode.

Homelab recommendation:

- Use Alertmanager plus ntfy/Gotify for experiments.
- For anything you genuinely rely on, use an external notification path.
- Keep postmortems as Markdown in the repo.

### On-Call

Datadog feature:

- Schedules, escalation policies, paging, notification rules.

OSS equivalents:

- Alertmanager routing for simple cases.
- Grafana OnCall OSS is no longer a good new choice because it was archived on 2026-03-24.
- Cabot and Alerta may cover parts, but with less polish.
- ntfy/Gotify for simple notifications.

Adoption difficulty: Medium to Hard.

Caveats:

- Reliable on-call is one area where SaaS is often worth paying for.
- Self-hosted on-call inside the monitored cluster can fail exactly when needed.

Homelab recommendation:

- Use ntfy/Gotify for learning.
- Do not spend too much time recreating PagerDuty unless that is the learning goal.

### Status Pages

Datadog feature:

- Public/private status pages and incident communication.

OSS equivalents:

- OpenStatus.
- Uptime Kuma status pages.
- Upptime.
- Cachet, though maintenance activity should be checked before adoption.

Adoption difficulty: Easy to Medium.

Caveats:

- A status page hosted inside the broken environment can be unavailable during incidents.
- Public status pages should live somewhere more reliable than the cluster they report on.

Homelab recommendation:

- Uptime Kuma is easiest.
- OpenStatus is more interesting if you want API monitoring as code and a more modern status-page workflow.

### Case Management

Datadog feature:

- Triage, track, assign, and remediate issues from a centralized case workflow.

OSS equivalents:

- GitHub Issues.
- GitLab Issues.
- Plane.
- Linear is not OSS, but common.

Adoption difficulty: Easy.

Caveats:

- Datadog's advantage is correlation with telemetry.
- For homelab, issue tracking is enough.

Homelab recommendation:

- Use Markdown notes or GitHub issues.

### Event Management

Datadog feature:

- Notable changes, alerts, deployments, incidents, and infrastructure events in one timeline.

OSS equivalents:

- Kubernetes events exporter into Prometheus/Loki.
- Grafana annotations.
- Prometheus Alertmanager history.
- GitOps/deployment annotations.

Adoption difficulty: Medium.

Caveats:

- Event correlation is one of the places where Datadog's unified platform feels much smoother.

Homelab recommendation:

- Add Kubernetes event export after Loki is present.
- Add Grafana annotations for deployments if you build a sample app.

### CI Visibility

Datadog feature:

- Pipeline, stage, job, and test visibility.
- CI duration and failure analytics.
- Committer-based views.
- Correlation with logs and tests.

OSS equivalents:

- CI-native metrics exporters.
- OpenTelemetry spans from CI jobs.
- GitHub Actions exporter, GitLab CI exporter, Jenkins Prometheus plugin.
- Buildkite Test Analytics is SaaS, not OSS.
- Allure Report for test reports.
- ReportPortal for test analytics.

Adoption difficulty: Medium.

Caveats:

- There is no universal OSS CI Visibility product as polished as Datadog's.
- Each CI system has different metadata and APIs.

Homelab recommendation:

- If using GitHub Actions or GitLab, start by exporting pipeline duration/failure metrics to Prometheus.
- Keep test details in the CI system until you have enough volume to justify ReportPortal.

### Test Optimization, Flaky Tests, Test Impact Analysis

Datadog feature:

- Flaky test detection.
- Test duration analytics.
- Impact analysis and test selection.

OSS equivalents:

- ReportPortal.
- Allure TestOps has OSS-adjacent tooling but not a simple full OSS match.
- Bazel test selection, pytest plugins, Jest reporters, and CI history scripts.

Adoption difficulty: Medium to Hard.

Caveats:

- Needs test history and consistent metadata.
- Low value unless you have a meaningful test suite.

Homelab recommendation:

- Skip unless you build a larger app repo.

### Continuous Testing

Datadog feature:

- Codeless API/browser tests integrated with CI/CD.

OSS equivalents:

- Playwright.
- k6.
- Cypress.
- Newman for Postman collections.

Adoption difficulty: Medium.

Caveats:

- The tests are easy. The managed scheduling, result retention, artifacts, and alerting are the productized parts.

Homelab recommendation:

- Use Playwright or k6 in CI, then export a small pass/fail metric.

### DORA Metrics

Datadog feature:

- Deployment frequency, lead time, change failure rate, mean time to restore.

OSS equivalents:

- Four Keys project: https://github.com/dora-team/fourkeys
- GitLab/GitHub APIs plus Grafana dashboards.
- Software Delivery metrics from CI/CD metadata.

Adoption difficulty: Medium.

Caveats:

- DORA metrics need clear definitions for deployment and incident.
- For a homelab, this is more useful as a learning exercise than an operational need.

Homelab recommendation:

- Skip until you have a real deployment pipeline.

### Feature Flags and Experiments

Datadog feature:

- Feature flags, gradual rollout, A/B experiments, product analytics linkage.

OSS equivalents:

- OpenFeature.
- Unleash.
- Flipt.
- GrowthBook for experimentation.

Adoption difficulty: Medium.

Caveats:

- Feature flags are application architecture, not monitoring infrastructure.
- Poorly managed flags become technical debt.

Homelab recommendation:

- Use Unleash or Flipt only if you build an app that needs flags.

### Code Coverage

Datadog feature:

- Coverage trends, PR gates, and repository views.

OSS equivalents:

- Coverage.py, JaCoCo, Istanbul/nyc, Go coverage.
- Codecov has free tiers but is SaaS.
- SonarQube Community for broader code quality.

Adoption difficulty: Easy.

Caveats:

- Usually belongs in app repos, not this monitoring repo.

Homelab recommendation:

- Skip for cluster monitoring.

### Internal Developer Portal / Software Catalog

Datadog feature:

- Service catalog / internal developer portal.
- Ownership, metadata, dependencies, docs, telemetry, and scorecards.

OSS equivalents:

- Backstage.
- Cortex/Port are commercial alternatives.
- Simple Markdown service catalog for small environments.

Adoption difficulty: Medium to Hard.

Caveats:

- Backstage is powerful but heavy for a two-node homelab.
- A simple `services.md` gives most of the value at this scale.

Homelab recommendation:

- Start with Markdown:
  - service name
  - namespace
  - owner
  - URLs
  - dashboards
  - alerts
  - runbook
- Consider Backstage only if developer portal work itself is the study goal.

### Teams, Access Control, Governance, API, and Marketplace

Datadog feature:

- Teams and ownership metadata.
- Role-based access control.
- Governance controls.
- API access.
- Marketplace/extensions.

OSS equivalents:

- Grafana teams, folders, and permissions.
- Kubernetes RBAC.
- Git repository permissions and CODEOWNERS.
- Backstage ownership metadata.
- Terraform/OpenTofu providers for configuration management.
- Tool-specific APIs.

Adoption difficulty: Medium.

Caveats:

- Datadog centralizes these controls. OSS spreads them across Grafana, Kubernetes, Git, CI, and each backend.
- For one user on a trusted LAN, complex RBAC is not worth much.

Homelab recommendation:

- Keep it simple.
- Use Git as the source of truth for manifests, values, notes, and runbooks.

### Workflow Automation and App Builder

Datadog feature:

- Automate workflows and build internal apps on top of telemetry.

OSS equivalents:

- n8n.
- StackStorm.
- Rundeck.
- Windmill.
- Grafana dashboard links and actions for simple workflows.

Adoption difficulty: Medium.

Caveats:

- Automation against a homelab can be useful, but it can also hide what is happening while you are learning.

Homelab recommendation:

- Use scripts and Makefiles first.
- Add n8n or Rundeck only if you want to study automation platforms.

### Fleet Automation

Datadog feature:

- Remotely configure, upgrade, and manage Datadog Agents.

OSS equivalents:

- Ansible.
- Flux or Argo CD for Kubernetes manifests.
- Renovate for dependency/chart updates.

Adoption difficulty: Medium.

Caveats:

- Datadog's feature manages Datadog agents specifically.
- For your stack, GitOps is the more general replacement.

Homelab recommendation:

- Keep Helm values and manifests in Git.
- Add Flux later if you want GitOps.

### Integrations Marketplace

Datadog feature:

- 1,000+ integrations for metrics, logs, traces, and dashboards.

OSS equivalents:

- Prometheus exporters.
- OpenTelemetry receivers/exporters.
- Grafana dashboards.
- Helm charts.
- Fluent Bit parsers.

Adoption difficulty: Varies.

Caveats:

- Datadog integrations are curated and consistent.
- OSS integrations vary heavily in quality, maintenance, and dashboard usefulness.

Homelab recommendation:

- Prefer official exporters or widely used community exporters.
- Pin Helm chart versions once something works.

## Recommended Homelab Architecture

### Phase 1: Finish Metrics

Goal: match the core Datadog host/Kubernetes metrics experience.

Install or configure:

- kubelet/cAdvisor ServiceMonitor.
- Alertmanager.
- A small set of alerting rules.
- Custom `Homelab Nodes` and `Kubernetes Workloads` dashboards.

Why first:

- Lowest overhead.
- Builds on what is already installed.
- Unlocks pod/container CPU, memory, network, and PVC metrics.

### Phase 2: Logs

Goal: add the most important Datadog feature missing from Prometheus/Grafana.

Install:

- Loki.
- Alloy/Promtail or Fluent Bit as DaemonSet.
- Grafana Loki datasource.
- Basic log dashboards and log alerts.

Important rules:

- Do not label logs with high-cardinality fields.
- Keep retention short initially.
- Start with pod logs only.

### Phase 3: Synthetic Monitoring

Goal: know whether important endpoints respond.

Install:

- blackbox_exporter.
- ServiceMonitor or scrape config.
- Alerts for probe failure and high latency.

Optional:

- Uptime Kuma for a friendly UI/status page.

### Phase 4: Traces

Goal: learn APM by instrumenting one service.

Install:

- OpenTelemetry Collector.
- Tempo.
- Grafana Tempo datasource.
- One toy app emitting traces.

Do not start by trying to trace every cluster component.

### Phase 5: Retention

Goal: make metrics durable beyond local Prometheus retention.

Options:

- VictoriaMetrics single-node: simplest practical option for this homelab.
- Thanos: good if you specifically want to learn object-storage-backed Prometheus architecture.
- Mimir: powerful but likely too much for two Orange Pis unless learning Mimir is the goal.

Recommendation:

- Use VictoriaMetrics first if retention becomes painful.

### Phase 6: Security

Goal: learn basic runtime and vulnerability monitoring.

Install:

- Trivy for image/config scans.
- Falco for runtime events.
- Possibly kube-bench for Kubernetes posture checks.

Caveat:

- Treat findings as learning material before turning them into noisy alerts.

## Rough Priority Matrix

| Priority | Capability | Why |
| --- | --- | --- |
| 1 | kubelet/cAdvisor | Biggest current metrics gap. |
| 2 | Alertmanager | Monitoring without alert routing is incomplete. |
| 3 | Loki | Logs are the next major Datadog-equivalent pillar. |
| 4 | blackbox_exporter | Easy synthetic monitoring with high practical value. |
| 5 | Custom Grafana dashboards | Imported dashboards are generic; Datadog's UX comes from good curation. |
| 6 | OpenTelemetry + Tempo | Adds real APM/tracing experience. |
| 7 | CronJob/job monitoring | Good learning project and directly matches your notes. |
| 8 | VictoriaMetrics/Thanos/Mimir | Only when retention matters. |
| 9 | Falco/Trivy | Useful, but after observability basics. |
| 10 | RUM/session replay | App-specific, not core cluster monitoring. |
| 11 | Backstage/service catalog | Overkill unless you want portal practice. |
| 12 | SIEM/security posture platform | Heavy; start with smaller security tools. |

## What Datadog Still Does Better

Datadog's main advantage is not that each individual feature is impossible to reproduce. It is that the features are integrated:

- Host, pod, log, trace, RUM, synthetic, deployment, incident, and service ownership data share a tag model.
- Dashboards, monitors, notebooks, service pages, and incident workflows link together.
- Many integrations work with minimal setup.
- Retention, scaling, auth, upgrades, and cross-product correlation are handled by the platform.
- Advanced features like Watchdog, dynamic instrumentation, universal service monitoring, and polished DBM are hard to match cleanly with OSS.

The OSS tradeoff is:

- Lower direct software cost.
- More control.
- Better learning value.
- More operational responsibility.
- More integration work.
- More opportunities to create noisy, fragile, or high-cardinality telemetry by accident.

## What Fits This Cluster Best

Good fit:

- Prometheus.
- Grafana.
- kube-state-metrics.
- node-exporter.
- kubelet/cAdvisor scraping.
- Alertmanager.
- Loki.
- blackbox_exporter.
- Uptime Kuma.
- OpenTelemetry Collector.
- Tempo for small experiments.
- Trivy.
- Falco, if arm64/kernel compatibility checks out.

Maybe later:

- VictoriaMetrics.
- PMM.
- OpenCost.
- k6 browser checks.
- Backstage.
- OpenStatus.

Probably not worth it here:

- Full SIEM.
- Mimir microservices mode.
- Large OpenSearch logging stack.
- Heavy session replay.
- Full developer portal.
- Complex workflow automation.
- Enterprise-style on-call management inside the same cluster.

## Concrete Next Tasks

1. Add kubelet/cAdvisor scraping.
2. Add Alertmanager with one simple notification target.
3. Add a custom multi-node Grafana dashboard.
4. Add blackbox_exporter and probes for Prometheus, Grafana, registry, and Traefik.
5. Add Loki with short retention.
6. Build a small sample app with:
   - Prometheus metrics.
   - structured logs.
   - OpenTelemetry traces.
   - one CronJob.
   - one synthetic check.
7. Use that app to learn dashboards, alerts, traces, logs, job monitoring, and incident notes end to end.
