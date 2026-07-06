# CKA 2026 Study Refresh Guide

This guide is for refreshing 2020-era CKA notes against the current CKA curriculum. The old notes are still useful for Kubernetes fundamentals, but the exam emphasis has shifted and several topics now need explicit practice.

Current curriculum snapshot checked in July 2026:

- Kubernetes version: v1.35, per the Linux Foundation CKA exam page.
- CNCF curriculum domains:
  - 10% Storage
  - 15% Workloads and Scheduling
  - 20% Services and Networking
  - 25% Cluster Architecture, Installation and Configuration
  - 30% Troubleshooting

Primary sources:

- Linux Foundation CKA page: https://training.linuxfoundation.org/certification/certified-kubernetes-administrator-cka/
- CNCF curriculum repo: https://github.com/cncf/curriculum
- Kubernetes docs: https://kubernetes.io/docs/

## How To Use This Guide

Treat this as a bridge curriculum, not a full CKA course. For each topic:

1. Read or skim the current Kubernetes docs.
2. Do the labs from a clean cluster.
3. Repeat the troubleshooting labs without looking at notes.
4. Update the original section notes only after the hands-on work.

Suggested local setup:

- Use `kubeadm`, kind, or Killercoda/KodeKloud labs depending on the topic.
- Use a real multi-node kubeadm cluster for upgrade, CNI, kubelet, CRI, etcd, and node troubleshooting.
- Use kind for fast YAML, Gateway API, Helm, Kustomize, RBAC, and workload practice.

## 1. Gateway API

Why this matters:

The old notes cover classic Ingress. The current CKA curriculum explicitly lists Gateway API for managing ingress traffic. You should know how Gateway API relates to Ingress and how to inspect Gateway API resources.

Brush up on:

- `GatewayClass`
- `Gateway`
- `HTTPRoute`
- Route attachment to Gateways
- Hostname and path routing
- Status conditions on Gateway API resources
- Relationship between Gateway API and the controller implementation

Labs:

1. Install a Gateway API controller in a kind cluster, such as Envoy Gateway or NGINX Gateway Fabric.
2. Deploy two simple apps: `web` and `api`.
3. Create one `GatewayClass`, one `Gateway`, and two `HTTPRoute` resources:
   - `/` routes to `web`
   - `/api` routes to `api`
4. Break the route by using the wrong Service name. Use `kubectl describe httproute` and status conditions to find the issue.
5. Add hostname routing with two hostnames, then test traffic with `curl -H 'Host: ...'`.
6. Compare the equivalent classic Ingress manifest and write down the differences.

Completion check:

- You can explain what owns the actual dataplane.
- You can debug an unattached route.
- You can find routing errors from `status.conditions`.

Docs:

- https://kubernetes.io/docs/concepts/services-networking/gateway/
- https://gateway-api.sigs.k8s.io/

## 2. Helm

Why this matters:

The current CKA curriculum explicitly includes using Helm to install cluster components. Your old notes do not appear to cover Helm.

Brush up on:

- Repositories
- Charts
- Releases
- Values files
- `helm install`, `upgrade`, `rollback`, `uninstall`
- `helm template`
- Inspecting installed resources with `kubectl`

Labs:

1. Add a public Helm repo.
2. Install a small chart into a dedicated namespace.
3. Override values with `--set`.
4. Override values with a `values.yaml` file.
5. Use `helm list -A`, `helm status`, and `helm get values`.
6. Upgrade the release and then roll it back.
7. Render manifests with `helm template` and apply them manually with `kubectl apply`.
8. Intentionally set a bad image tag, observe the failed rollout, fix it with `helm upgrade`.

Completion check:

- You can install, inspect, upgrade, and remove a chart.
- You can tell which Kubernetes objects a chart created.
- You can recover from a bad Helm upgrade.

Docs:

- https://helm.sh/docs/

## 3. Kustomize

Why this matters:

Kustomize is now explicitly part of the CKA curriculum for installing and managing components. It is built into `kubectl`.

Brush up on:

- `kustomization.yaml`
- Resources
- Patches
- Name prefixes/suffixes
- Common labels
- ConfigMap and Secret generators
- Overlays
- `kubectl kustomize`
- `kubectl apply -k`

Labs:

1. Create a base Deployment and Service.
2. Create `dev` and `prod` overlays.
3. In `dev`, set replicas to 1 and use a test image tag.
4. In `prod`, set replicas to 3 and add resource requests/limits.
5. Generate a ConfigMap from literals.
6. Patch an environment variable into the Deployment.
7. Run `kubectl kustomize overlays/dev` and inspect the generated YAML.
8. Apply each overlay to separate namespaces.

Completion check:

- You can build and apply an overlay without editing the base.
- You can use patches instead of duplicating full manifests.
- You can explain when Kustomize is enough versus when Helm is better.

Docs:

- https://kubernetes.io/docs/tasks/manage-kubernetes-objects/kustomization/

## 4. CRDs And Operators

Why this matters:

The current curriculum explicitly mentions understanding CRDs and installing/configuring operators. In 2020 this was less central for CKA.

Brush up on:

- CustomResourceDefinition
- Custom resources
- Controllers and reconciliation
- Operators as packaged controllers
- `kubectl api-resources`
- `kubectl explain`
- CRD scope: namespaced vs cluster
- Finalizers and status fields at a high level

Labs:

1. Install a simple operator, such as cert-manager, Prometheus Operator, or another lightweight operator.
2. List new API resources with `kubectl api-resources`.
3. Inspect installed CRDs with `kubectl get crd`.
4. Create one custom resource managed by the operator.
5. Use `kubectl describe` to inspect events and status.
6. Delete the custom resource and observe cleanup behavior.
7. Break the custom resource spec and identify the failure through status/events.

Completion check:

- You can identify which CRDs an operator installed.
- You understand that custom resources are declarative inputs to a controller.
- You can inspect an operator-managed object without needing deep operator internals.

Docs:

- https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/
- https://kubernetes.io/docs/concepts/extend-kubernetes/operator/

## 5. Workload Autoscaling

Why this matters:

The current curriculum explicitly includes workload autoscaling. The old notes mention autoscale commands, but you should practice the full workflow.

Brush up on:

- HorizontalPodAutoscaler
- Metrics Server
- CPU and memory resource requests
- `kubectl autoscale`
- `autoscaling/v2`
- HPA status fields
- Why HPA does not work without metrics or requests

Labs:

1. Install Metrics Server in a test cluster.
2. Deploy a CPU-bound sample app.
3. Add CPU requests and limits.
4. Create an HPA with min/max replicas and CPU target.
5. Generate load and watch scaling with `kubectl get hpa -w`.
6. Remove CPU requests and observe why the HPA cannot calculate utilization.
7. Use `kubectl top pods` and `kubectl top nodes`.

Completion check:

- You can explain why resource requests matter for HPA.
- You can debug `unknown` HPA metrics.
- You can create HPA imperatively and declaratively.

Docs:

- https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/

## 6. Modern Storage: StorageClass, CSI, Dynamic Provisioning

Why this matters:

Your old notes cover PV/PVC basics. The current curriculum emphasizes StorageClasses, dynamic provisioning, reclaim policies, access modes, and CSI awareness.

Brush up on:

- `StorageClass`
- Dynamic provisioning
- Default StorageClass
- `volumeBindingMode`
- Reclaim policies: `Retain`, `Delete`
- Access modes: `ReadWriteOnce`, `ReadOnlyMany`, `ReadWriteMany`, `ReadWriteOncePod`
- CSI at a conceptual level

Labs:

1. Inspect the cluster's default StorageClass.
2. Create a PVC that dynamically provisions a PV.
3. Mount it into a Pod and write a file.
4. Delete the Pod and recreate it with the same PVC, then verify the file remains.
5. Change reclaim policy behavior in a test StorageClass if your environment supports it.
6. Create a PVC that cannot bind and troubleshoot why.
7. Practice `kubectl describe pvc`, `kubectl describe pv`, and events.

Completion check:

- You can explain when a PV is created automatically.
- You can debug a Pending PVC.
- You understand what CSI is responsible for without needing to implement a driver.

Docs:

- https://kubernetes.io/docs/concepts/storage/storage-classes/
- https://kubernetes.io/docs/concepts/storage/persistent-volumes/
- https://kubernetes.io/docs/concepts/storage/volumes/

## 7. Runtime Changes: CRI, containerd, And Dockershim Removal

Why this matters:

Old notes and old courses often assume Docker-specific behavior. Kubernetes removed dockershim in v1.24. Modern clusters commonly use containerd or CRI-O through the Container Runtime Interface.

Brush up on:

- CRI
- containerd
- CRI-O
- `crictl`
- kubelet runtime endpoint
- Difference between Docker as a developer tool and Kubernetes runtime integration
- Static Pod manifests under `/etc/kubernetes/manifests`

Labs:

1. On a kubeadm node, inspect the kubelet config and find the container runtime endpoint.
2. Use `crictl ps`, `crictl images`, `crictl logs`, and `crictl inspect`.
3. Find a control plane static Pod through `crictl`.
4. Stop or misconfigure a test workload and compare `kubectl logs` with `crictl logs`.
5. Restart containerd and observe node/pod behavior in a test cluster.

Completion check:

- You do not rely on `docker ps` for node-level Kubernetes troubleshooting.
- You can use `crictl` to inspect containers when the API server or kubelet path is impaired.
- You can identify the configured runtime endpoint.

Docs:

- https://kubernetes.io/docs/setup/production-environment/container-runtimes/
- https://kubernetes.io/blog/2022/02/17/dockershim-faq/

## 8. API Version And Manifest Modernization

Why this matters:

Some manifests valid in 2020 are invalid now. The most important exam habit is to generate or verify manifests against current APIs.

Brush up on:

- Ingress `networking.k8s.io/v1`
- Current Deployment, Service, Job, CronJob, RBAC, and NetworkPolicy APIs
- `kubectl explain`
- `kubectl api-resources`
- `kubectl create --dry-run=client -o yaml`
- Deprecated beta API removal patterns

Labs:

1. Find old manifests in your notes and run `kubectl apply --dry-run=server -f`.
2. Convert old Ingress manifests to `networking.k8s.io/v1`.
3. Use `kubectl explain ingress.spec.rules.http.paths.backend.service`.
4. Generate modern manifests imperatively where possible.
5. Create a short personal cheat sheet of current API versions for common exam resources.

Completion check:

- You can quickly detect an obsolete API version.
- You know how to use `kubectl explain` instead of guessing field names.
- Your old Ingress examples are updated to current syntax.

Docs:

- https://kubernetes.io/docs/reference/using-api/deprecation-guide/
- https://kubernetes.io/docs/reference/kubectl/generated/kubectl_explain/

## 9. Pod Security Admission

Why this matters:

PodSecurityPolicy is removed. Current Kubernetes uses Pod Security Admission for built-in namespace-level policy enforcement.

Brush up on:

- Pod Security Standards: privileged, baseline, restricted
- Namespace labels:
  - `pod-security.kubernetes.io/enforce`
  - `pod-security.kubernetes.io/audit`
  - `pod-security.kubernetes.io/warn`
- How securityContext interacts with restricted policies
- Difference between admission policy and RBAC

Labs:

1. Create a namespace with `restricted` enforcement.
2. Try to run a privileged Pod and observe rejection.
3. Add a compliant security context and get the Pod admitted.
4. Test `warn` mode before `enforce` mode.
5. Compare this with NetworkPolicy and RBAC to clarify what each security mechanism controls.

Completion check:

- You can explain why RBAC does not make a privileged Pod safe.
- You can apply namespace-level pod security labels.
- You can fix a Pod rejected by Pod Security Admission.

Docs:

- https://kubernetes.io/docs/concepts/security/pod-security-admission/
- https://kubernetes.io/docs/concepts/security/pod-security-standards/

## 10. Troubleshooting: Make This The Main Refresh Block

Why this matters:

Troubleshooting is now 30% of the curriculum. Your old notes already have troubleshooting sections, but this should become the center of the refresh.

Brush up on:

- Application failures
- Service and endpoint failures
- DNS failures
- NetworkPolicy failures
- Node NotReady
- kubelet failures
- CNI failures
- Control plane static Pod failures
- etcd snapshot and restore
- Certificate issues
- Resource pressure and scheduling failures

Labs:

1. Application failure:
   - Bad image
   - Wrong command
   - Failed readiness probe
   - Missing ConfigMap or Secret
   - Insufficient resource requests

2. Service failure:
   - Wrong selector
   - Wrong targetPort
   - No endpoints
   - App listening on the wrong port
   - Test with temporary debug Pods

3. DNS failure:
   - Break CoreDNS config in a lab cluster.
   - Restore it.
   - Practice `nslookup`, `dig`, and `/etc/resolv.conf` checks from a Pod.

4. NetworkPolicy failure:
   - Deny all ingress.
   - Allow only from a selected namespace and Pod label.
   - Debug why traffic is blocked.

5. Node failure:
   - Stop kubelet on one worker.
   - Inspect `kubectl get nodes`, `kubectl describe node`, and systemd logs.
   - Restart kubelet and confirm recovery.

6. Control plane failure:
   - Break one static Pod manifest in `/etc/kubernetes/manifests` in a disposable cluster.
   - Use container runtime logs and kubelet logs to recover.

7. etcd backup and restore:
   - Take a snapshot.
   - Delete a test namespace.
   - Restore the snapshot in a lab environment.
   - Repeat until the command sequence is memorized.

Completion check:

- You can move from symptom to root cause quickly.
- You can debug both with `kubectl` and on the node.
- You can recover from common broken-cluster scenarios without course hints.

Docs:

- https://kubernetes.io/docs/tasks/debug/
- https://kubernetes.io/docs/tasks/debug/debug-cluster/
- https://kubernetes.io/docs/tasks/debug/debug-application/

## 11. kubeadm And Cluster Lifecycle Refresh

Why this matters:

Cluster installation and lifecycle remain a major CKA area. The commands and package versions in the old notes are stale.

Brush up on:

- Current kubeadm installation flow
- Package repositories for current Kubernetes versions
- `kubeadm init`
- `kubeadm join`
- `kubeadm upgrade plan`
- `kubeadm upgrade apply`
- `kubeadm upgrade node`
- Static Pod control plane
- Certificate renewal basics
- Highly available control plane concepts

Labs:

1. Build a new kubeadm cluster from scratch.
2. Install a CNI plugin.
3. Join one worker node.
4. Deploy a test app and verify DNS, Service, and Pod networking.
5. Upgrade the cluster one minor version.
6. Renew or inspect certificates with `kubeadm certs`.
7. Create and restore an etcd snapshot.

Completion check:

- You can build a working cluster without relying on 2020 package commands.
- You can perform a minor version upgrade.
- You know where kubeadm places manifests, certs, and kubeconfigs.

Docs:

- https://kubernetes.io/docs/setup/production-environment/tools/kubeadm/
- https://kubernetes.io/docs/tasks/administer-cluster/kubeadm/kubeadm-upgrade/

## 12. Services, Endpoints, And Networking Modern Review

Why this matters:

Networking is 20% of the current curriculum. The old notes cover Services, DNS, NetworkPolicy, and Ingress, but you should practice them with current APIs and deeper troubleshooting.

Brush up on:

- ClusterIP
- NodePort
- LoadBalancer
- EndpointSlice
- CoreDNS
- CNI responsibilities
- NetworkPolicy ingress and egress
- Ingress and Gateway API

Labs:

1. Create a Deployment and expose it with ClusterIP.
2. Break the Service selector and debug missing endpoints.
3. Inspect EndpointSlices.
4. Convert the Service to NodePort and test node access.
5. Add a NetworkPolicy that blocks traffic, then selectively allow it.
6. Deploy Ingress and Gateway API examples for the same app.
7. Break DNS by using the wrong namespace/service name and debug from inside a Pod.

Completion check:

- You can debug Service routing from Pod labels to EndpointSlices.
- You understand the difference between Service, Ingress, and Gateway API.
- You can explain what Kubernetes does versus what the CNI/controller provides.

Docs:

- https://kubernetes.io/docs/concepts/services-networking/service/
- https://kubernetes.io/docs/concepts/services-networking/endpoint-slices/
- https://kubernetes.io/docs/concepts/services-networking/network-policies/
- https://kubernetes.io/docs/concepts/services-networking/ingress/

## Suggested Four-Week Refresh Plan

### Week 1: Modern APIs, Workloads, Helm, Kustomize

- Day 1: API version modernization and `kubectl explain`
- Day 2: Kustomize bases and overlays
- Day 3: Helm install, upgrade, rollback, template
- Day 4: HPA and Metrics Server
- Day 5: Pod Security Admission
- Day 6: Mixed practice
- Day 7: Rest or light review

Deliverable:

- Updated cheat sheet for current API versions.
- One working Helm lab.
- One working Kustomize lab.
- One HPA lab.

### Week 2: Networking Refresh

- Day 1: Services and EndpointSlices
- Day 2: CoreDNS troubleshooting
- Day 3: NetworkPolicy ingress and egress
- Day 4: Classic Ingress current API
- Day 5: Gateway API
- Day 6: Mixed network break/fix drills
- Day 7: Rest or light review

Deliverable:

- One Service troubleshooting checklist.
- One Gateway API routing lab.
- Updated Ingress notes using `networking.k8s.io/v1`.

### Week 3: Cluster Lifecycle, Runtime, Storage

- Day 1: kubeadm cluster build
- Day 2: kubeadm upgrade
- Day 3: containerd, CRI, `crictl`
- Day 4: StorageClass and dynamic provisioning
- Day 5: etcd backup and restore
- Day 6: CRDs and operators
- Day 7: Rest or light review

Deliverable:

- One fresh kubeadm cluster build log.
- One etcd restore runbook.
- One dynamic storage lab.
- One CRD/operator inspection lab.

### Week 4: Troubleshooting-Heavy Exam Practice

- Day 1: Application failure drills
- Day 2: Service and networking failure drills
- Day 3: Node and kubelet failure drills
- Day 4: Control plane failure drills
- Day 5: Full mock exam
- Day 6: Redo all failed areas
- Day 7: Final command review

Deliverable:

- A personal troubleshooting checklist.
- A list of commands you can type without notes.
- Redone weak areas from old mock exams: CSR/RBAC, Pod DNS, ServiceAccounts, ClusterRoles, JSONPath, NetworkPolicies.

## Personal Priority List Based On Existing Notes

Highest priority:

- Gateway API
- Helm
- Kustomize
- CRDs/operators
- HPA and metrics
- StorageClass and dynamic provisioning
- CRI/containerd/`crictl`
- API version modernization
- Pod Security Admission
- Troubleshooting drills

Medium priority:

- kubeadm install and upgrade using current docs
- EndpointSlice troubleshooting
- CoreDNS troubleshooting
- NetworkPolicy egress
- HA control plane concepts

Still useful from the old notes:

- Pods, ReplicaSets, Deployments
- Services
- Namespaces
- Scheduling, taints, tolerations, affinity
- ConfigMaps and Secrets
- Rolling updates and rollbacks
- RBAC
- NetworkPolicy basics
- PV/PVC basics
- etcd concepts
- kubeadm concepts
- Application, control plane, worker node, and networking troubleshooting

## Final Readiness Checks

Before booking or taking the exam, you should be able to do these without course hints:

- Create and debug Deployments, Services, ConfigMaps, Secrets, RBAC, NetworkPolicies, PVs, PVCs, StorageClasses, Ingress, Gateway API resources, and HPAs.
- Install and inspect a Helm chart.
- Apply Kustomize overlays.
- Identify and fix obsolete API versions.
- Use `kubectl explain`, `kubectl api-resources`, and `kubectl create --dry-run=client -o yaml` fluently.
- Troubleshoot a broken app, broken Service, broken DNS, broken node, and broken control plane.
- Use `crictl` on a node.
- Take and restore an etcd snapshot.
- Build and upgrade a kubeadm cluster using current docs.
