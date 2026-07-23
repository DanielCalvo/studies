# Practical CKAD Topics Not Explicitly Covered by the CKA

The CKAD contains several workload-level topics that are useful in everyday Kubernetes but are not explicitly covered—or are covered much less directly—by the CKA.

This comparison uses the Kubernetes 1.35 curricula published by CNCF. The wording matters: the curricula define broad competencies, so “not in the CKA outline” does not necessarily mean that a topic could never appear incidentally in a CKA task.

## The Most Useful CKAD Additions

### 1. Probes and Health Checks

This remains the clearest example. CKAD explicitly requires implementing probes and health checks; CKA only refers more generally to “robust, self-healing application deployments.”

In practice, you should understand:

- Startup, readiness, and liveness probes
- HTTP, TCP, gRPC, and exec probes
- How a failing readiness probe differs from a failing liveness probe
- Appropriate delays, periods, thresholds, and timeouts
- Why an overly aggressive liveness probe can make an outage worse

### 2. Multi-Container Pod Patterns

CKAD explicitly covers sidecars, init containers, and other multi-container designs. CKA does not name them.

These appear frequently in real systems:

- Init containers performing setup or dependency checks
- Service-mesh and telemetry sidecars
- Shared volumes between containers
- Understanding that containers within a Pod share networking and lifecycle boundaries

### 3. Application Security Settings

CKAD explicitly includes:

- `securityContext`
- Linux capabilities
- Service accounts
- Application-level authentication, authorization, and admission concepts

CKA covers RBAC and admission, but it does not explicitly cover container security contexts, capabilities, or service accounts as workload configuration.

For everyday work, this is particularly important: running as non-root, setting `runAsNonRoot`, using read-only root filesystems, dropping capabilities, and avoiding unnecessary service-account tokens are all common production requirements.

### 4. Choosing the Appropriate Workload Controller

CKAD explicitly expects you to choose among resources such as:

- `Deployment`
- `DaemonSet`
- `Job`
- `CronJob`

CKA discusses application deployments and self-healing workloads, but does not explicitly enumerate this decision. Being able to distinguish a continuously reconciled service from a finite job or node-local daemon is fundamental day-to-day knowledge.

### 5. Deployment Strategies Beyond Ordinary Rolling Updates

Both exams cover rolling updates, but CKAD additionally names strategies such as:

- Blue/green deployments
- Canary deployments

Kubernetes does not provide a single native `CanaryDeployment` object; you implement these strategies using Deployments, labels, Services, replica ratios, routing, or external rollout tooling. Understanding the underlying mechanics is useful even when Argo Rollouts, Flagger, or a service mesh performs the higher-level orchestration.

### 6. API Deprecations

CKAD explicitly covers understanding API deprecations. CKA does not list this separately, although cluster lifecycle management naturally encounters them.

This matters whenever you:

- Upgrade a cluster
- Upgrade Helm charts or operators
- Maintain old manifests
- Discover that an API version has stopped being served

### 7. Container-Image Fundamentals

CKAD includes defining, building, and modifying container images; CKA does not.

This is useful for diagnosing:

- Incorrect entrypoints and arguments
- Missing files or dependencies
- Image architecture mismatches
- Image pull and registry problems
- Processes that do not handle termination signals correctly

Image building itself is not Kubernetes-specific, but understanding the boundary between the image and the Pod specification is essential for troubleshooting workloads.

### 8. Ephemeral Volumes

CKAD explicitly mentions both persistent and ephemeral volumes. CKA has much deeper persistent-storage coverage, including StorageClasses, access modes, reclaim policies, PVs, and PVCs, but does not explicitly emphasize application-oriented ephemeral volumes.

Useful examples include `emptyDir`, projected volumes, ConfigMap and Secret volumes, and ephemeral volumes used for scratch space or sharing data between containers.

## Areas That Are No Longer Meaningfully CKAD-Only

A lot of what might have looked CKAD-specific in 2020 is now present in both curricula:

- ConfigMaps and Secrets
- Requests and limits, at least to some degree
- Deployments, rolling updates, and rollbacks
- Services, Ingress, and NetworkPolicies
- Helm and Kustomize
- Monitoring application resource usage
- Container logs
- Application and networking troubleshooting
- CRDs and operators

In fact, the current CKA now goes beyond CKAD in some application-facing areas, including workload autoscaling, Gateway API, CoreDNS, and deeper network troubleshooting.

## Practical Conclusion

For someone who already has the CKA foundation, the genuinely valuable CKAD delta is approximately:

1. Probes and lifecycle behavior
2. Init containers and sidecars
3. Security contexts, capabilities, and service accounts
4. Jobs, CronJobs, and DaemonSets
5. Blue/green and canary mechanics
6. API deprecation handling
7. Container-image and process fundamentals

Studying those areas would provide most of the everyday operational value of CKAD without necessarily sitting another certification exam.

## Sources

- [Official CNCF curriculum repository](https://github.com/cncf/curriculum)
- [CKA 1.35 curriculum](https://github.com/cncf/curriculum/blob/master/CKA_Curriculum_v1.35.pdf)
- [CKAD 1.35 curriculum](https://github.com/cncf/curriculum/blob/master/CKAD_Curriculum_v1.35.pdf)
