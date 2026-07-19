# Shared Kubernetes Instructions

Use these instructions whenever editing, reviewing, generating, applying, or installing Kubernetes resources in this repository. More specific `AGENTS.md` files may add local rules.

## Declarative Changes

- Every Kubernetes change applied to a cluster must be represented declaratively in version-controlled files.
- Prefer YAML manifests, Kustomize overlays, Helm values files, or other declarative configuration that can reproduce the resulting cluster state.
- Do not make cluster changes only with imperative `kubectl` commands. If `kubectl apply` is used, apply from a saved manifest file rather than inline YAML or ad hoc command flags.
- For Helm chart installs or upgrades, save the values in a file and install or upgrade using that file, such as `helm upgrade --install ... -f values.yaml`.
- If a one-off imperative command is unavoidable for investigation or cleanup, call it out explicitly and do not treat it as the primary implementation path.
