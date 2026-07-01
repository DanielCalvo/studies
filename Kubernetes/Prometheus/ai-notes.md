# Prometheus Operator Notes

## CRDs vs. the Operator

CRDs and the Prometheus Operator are related, but they are not the same thing.

**CRDs** install new Kubernetes API types into the cluster, for example:

- `Prometheus`
- `ServiceMonitor`
- `PodMonitor`
- `Alertmanager`
- `PrometheusRule`
- `Probe`
- `ScrapeConfig`
- `ThanosRuler`

Once the CRDs exist, the Kubernetes API server can accept objects of those kinds.

The **operator** is the controller deployment that watches those custom resources and turns them into real workload and configuration changes. For example, if you create a `Prometheus` object, the operator reconciles that desired state into Kubernetes resources such as StatefulSets, Secrets, ConfigMaps, and Services.

In short:

```text
CRDs = teach Kubernetes about new resource kinds
Operator = controller that acts on those resource kinds
Custom resources = desired monitoring objects
```

In production, this distinction matters because CRDs have different lifecycle rules than normal Kubernetes resources. Helm, for example, treats CRDs specially: it can install them from a chart's `crds/` directory, but CRD upgrades and deletions are intentionally limited. That is why many teams manage Prometheus Operator CRDs separately from the operator deployment.

A common declarative production pattern is:

```text
1. Install/update Prometheus Operator CRDs
2. Install/update the Prometheus Operator controller
3. Apply Prometheus, Alertmanager, ServiceMonitor, PrometheusRule, etc.
```

With GitOps, that often becomes separate apps or releases, ordered like:

```text
prometheus-operator-crds
prometheus-operator
monitoring-resources
```

For learning, installing everything together from the upstream bundle is okay. For production, separating CRDs from the operator is usually cleaner and safer.

## Mental Model for Placement

A good way to organize the Prometheus Operator pieces is:

```text
Cluster level:
  Prometheus Operator CRDs
    - Prometheus
    - Alertmanager
    - ServiceMonitor
    - PodMonitor
    - PrometheusRule
    - etc.

Dedicated namespace:
  prometheus-operator
    - Prometheus Operator Deployment
    - operator ServiceAccount/RBAC
    - maybe admission webhook resources

Monitoring namespaces:
  monitoring-prod
    - Prometheus instance
    - Alertmanager instance
    - ServiceMonitors
    - PrometheusRules

  monitoring-dev
    - another Prometheus instance
    - maybe separate rules/monitors
```

The CRDs define the API. The operator runs as the controller. Then you create custom resources, such as a `Prometheus` object, and the operator creates and manages the actual Prometheus workload for you.

For example:

```yaml
apiVersion: tmp_monitoring.coreos.com/v1
kind: Prometheus
metadata:
  name: main
  namespace: tmp_monitoring-prod
spec:
  serviceMonitorSelector: {}
  podMonitorSelector: {}
```

That means the normal production flow is:

```text
1. Install CRDs once.
2. Run the operator in a dedicated namespace.
3. Mostly leave those foundational pieces alone.
4. Use the CRDs to declaratively create Prometheus, Alertmanager, ServiceMonitor, PrometheusRule, and related resources.
```

One important production detail is namespace scoping. The operator can be configured to watch all namespaces or only selected namespaces. Each `Prometheus` instance can also be configured to select `ServiceMonitor`, `PodMonitor`, and `PrometheusRule` resources from specific namespaces. This is the part to design carefully when running multiple Prometheus instances or separating teams/environments.
