## 1. https://prometheus-operator.dev/docs/platform/platform-guide/
- Creates a cluster role and links it to prometheus. 
- Then creates a prometheus instance and an alertmanager instance through the CRDs!

## 2. https://prometheus-operator.dev/docs/platform/exposing-prometheus-and-alertmanager/
- Added `loadbalancer-services.yaml` with two `LoadBalancer` Services, one for Prometheus and one for Alertmanager.
- This exposes each UI directly on its own MetalLB IP without Ingress path prefixes.
- Pinned the MetalLB IPs:
  - Prometheus: `http://192.168.1.243:9090`
  - Alertmanager: `http://192.168.1.244:9093`

## 3. https://prometheus-operator.dev/docs/platform/exposing-prometheus-and-alertmanager/
- Did this with metallb (homelab special)

## 4. https://prometheus-operator.dev/docs/platform/webhook/
- This seems like a best practice for a production system but to just explore the operator i dont think we need this

## 5. https://prometheus-operator.dev/docs/platform/prometheus-agent/
- The prometheus agent is a feature/CRD you can use to forward collected data to a long term solution like thanos, cortex or more prometheus

## 6. https://prometheus-operator.dev/docs/platform/thanos/
- You can integrate the Prometheus operator with Thanos. 
- You can use a thanos sidecar with the prometheus definiton to use thanos as something you can query
- the ThanosRuler component evaluates prometheus alerting and recording rules with a query API... so you can query a Thanos data source and alert on it, apparently

## 7. https://prometheus-operator.dev/docs/platform/rbac/
- Describes permissions set up for the operator controller pod itself. Also describes RBAC for prometheus server pods so they can read resources for service discovery (you have an example of this in your current folder! otherwise you dont have permissions to scrape stuff)

## 8. https://prometheus-operator.dev/docs/platform/rbac-crd/
- Describes permission set up from to view, create, delete, etc, prometheus custom resources (like prometheus, alert manager, servicemonitor, podmonitor, etc)

## 9. https://prometheus-operator.dev/docs/platform/high-availability/
- Describes running Prometheus in high availability. One thing that caught my eye is, if you run two replicas of Prometheus with the operator, you run two Prometheus instances with separate local storage, and they both scrape and ingest data separately. This means that querying each instance might return slightly different values.

- Running multiple instances avoids having a single point of failure, but it doesn't help scaling out in case a single Prometheus instance can't handle all the targets and rules. This makes sense because you would be duplicating everything. In this case, the documentation recommends sharding.

- The docs then recommend using sticky sessions for dashboarding so you get consistent graphs.

- So it seems that by high availability here you have two copies of the same Prometheus instance with the same scrape settings, but different local storage. You're essentially duplicating the instance. It isn't like running two web services that query a single database.

## 10. https://prometheus-operator.dev/docs/platform/sharding/
- You can do some relabelling To change target distributions to perform sharding with Prometheus, so one instance only might read some things in a certain namespace and another another namespace for example

## 11. https://prometheus-operator.dev/docs/platform/storage/
So by the fault the operator stores data on empty deer volumes which aren't persisted when the pods are redeployed. to maintain data across deployments and versions you need to configure persistent storage for Prometheus alert manager and thanos ruler, wew

## 12. https://prometheus-operator.dev/docs/platform/strategic-merge-patch/
This just tells you how to patch certain resources on the install. seems situation specific.

## 13. https://prometheus-operator.dev/docs/platform/cli/
Interesting there's also an operator cli.  I-m not particularly interested in imperative confine for now but it-s good to know this exit

## 14. https://prometheus-operator.dev/docs/platform/troubleshooting/
Assorted troubleshooting info. We'll get here when we get here.