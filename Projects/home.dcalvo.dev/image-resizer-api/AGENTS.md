# Image Resizer API Instructions

The instructions in the parent `AGENTS.md` also apply to this project.

## Go toolchain

- Implement the application in Go.
- Go is installed at `/usr/local/go/bin/go`.
- `gofmt` is installed at `/usr/local/go/bin/gofmt`.
- The installed toolchain is currently Go 1.26.2 on Linux amd64.
- Codex command shells may be non-interactive and may not source the line in
  `/home/daniel/.bashrc` that appends `/usr/local/go/bin` to `PATH`.
- If `go` is reported as not found, do not assume that Go is uninstalled. Invoke
  `/usr/local/go/bin/go` directly. Invoke `/usr/local/go/bin/gofmt` directly for
  formatting when necessary.
- Before reporting a toolchain problem, verify it with
  `/usr/local/go/bin/go version`.

Keep the application and its dependencies simple. Preserve Kubernetes changes in
the existing manifests and deployment workflow.

## Shell script simplicity and ownership

Keep shell commands recognizable and close to the simplest command a person
would type manually. Include flags required for correctness, safety, or an
explicitly desired behavior, and briefly explain non-obvious flags. Prefer a
little repetition while there are only a few cases, and introduce abstractions
when repetition itself becomes an ownership burden. Do not add speculative
robustness or convenience options by default. This guidance applies only to
shell scripts.

Read `feature_list.md` for the chronological record of implemented features and
the reasoning behind their priority. Update it whenever a feature is selected
and implemented.

When the user asks to save an explanation or reusable reference note, create it
under `docs/` with an `ai_` filename prefix. Keep project operation and
feature-tracking files in their normal project locations.
