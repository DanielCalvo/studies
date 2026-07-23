# Shared Go Instructions

Use these instructions whenever editing, reviewing, generating, testing, or discussing Go code in this repository. More specific `AGENTS.md` files may add local rules.

## Toolchain and validation

- Go is installed under `/usr/local/go/bin`. Codex command shells may not include that directory in `PATH`, so if `go` or `gofmt` is reported as not found, use `/usr/local/go/bin/go` or `/usr/local/go/bin/gofmt` directly rather than assuming the toolchain is missing.
- Format changed Go files with `gofmt` and run the relevant tests, normally `go test ./...`, after code changes. If either check cannot be run, state that clearly.

## Standard-library-inspired style

Prefer idiomatic, standard-library-inspired Go: focused packages, small APIs,
cohesive functions, direct control flow, explicit errors, and restrained
abstraction. Treat this as a readability preference, not a constraint. Use
dependencies, interfaces, constructors, concurrency, and additional application
layers when they provide a clear practical benefit.
