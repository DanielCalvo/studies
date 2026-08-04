These sub folders in here contain Kubernetes configuration and notes for clusters in my home lab. Each sub folder is a separate cluster with its own services and settings

Read `ai_cluster_context.md` for reusable cluster facts such as node IPs, architecture, MetalLB usage, local registry details, home-lab constraints and other useful contextual information.

Guidance for all clusters in subfolders:
- These clusters run on a trusted local home network, so prefer simple solutions over production-grade TLS/auth when that helps experiments move faster
- Keep durable operational notes in normal Markdown files, not only in `AGENTS.md`
- Make sure to update `ai_cluster_context.md` with cluster related context as you perform changes to the cluster. This is for agents to read so they have context about the cluster that persists across sessions
- Only SSH into machines when necessary and once logged in, only perform read only operations by default. Only make changes if explicitly requested by the user

## Kubernetes cluster targeting safety

Treat the repository folder as the authoritative cluster selector:

| Repository folder | Merged kubeconfig | Context | Source kubeconfig | Expected API server | Expected nodes |
| --- | --- | --- | --- | --- | --- |
| `opi-cluster/` | `/home/daniel/.kube/config` | `opi` | `/home/daniel/.kube/config-opi` | `https://192.168.1.201:6443` | `opi1` (`192.168.1.201`), `opi2` (`192.168.1.202`) |
| `hp-cluster/` | `/home/daniel/.kube/config` | `hp` | `/home/daniel/.kube/config-hp` | `https://192.168.1.211:6443` | `hp1` (`192.168.1.211`), `hp2` (`192.168.1.212`) |

- Never use bare `kubectl`, `helm`, or another Kubernetes client command that
  relies on the global current context. The merged configuration is at the
  default location, `/home/daniel/.kube/config`, so the kubeconfig argument may
  be omitted, but always select the named context explicitly, including for
  read-only inspection.
- For `kubectl`, use `kubectl --context <context> ...`.
- For Helm, use `helm --kube-context <context> ...`.
- For other Kubernetes tools, use their named-context option. Do not perform a
  cluster operation with a tool that cannot explicitly select `opi` or `hp`.
- Before any command that changes cluster state, verify both the API endpoint
  and node identities with the selected context:

  ```bash
  kubectl --context <context> \
    config view --minify \
    -o jsonpath='{.clusters[0].cluster.server}{"\n"}'
  kubectl --context <context> get nodes -o wide
  ```

- Stop and ask the user which cluster to use if the requested work does not
  clearly identify `opi-cluster/` or `hp-cluster/`. Do not infer the target from
  the global current context.
- If one task intentionally touches both clusters, use the explicit matching
  context separately for every command and verify each cluster before its first
  state-changing command.
- `/home/daniel/.kube/config` is the normalized merged configuration and has
  unique `opi` and `hp` cluster, user, and context names. The two source files
  retain their original internal `default` names, so do not concatenate or
  directly merge those source files without renaming all three record types.
- A human may use `kubectl config use-context opi`,
  `kubectl config use-context hp`, or `kubectx` for interactive convenience.
  Agents and scripts must still pass the explicit context on every command.
