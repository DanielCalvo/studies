This folder contains Kubernetes configuration and notes for my k3s homelab.

Read `ai_cluster_context.md` for reusable cluster facts such as node IPs, architecture, MetalLB usage, local registry details, and home-lab constraints.

Repo areas:
- `container-registry/`: local Docker Registry v2 manifests and notes for images used in the homelab.
- `metallb/`: MetalLB configuration and validation examples.

Guidance:
- This runs on a trusted local home network, so prefer simple solutions over production-grade TLS/auth when that helps experiments move faster.
- The cluster is small and arm64, so prefer lightweight components and verify image architecture when relevant.
- Keep durable operational notes in normal Markdown files, not only in `AGENTS.md`.
