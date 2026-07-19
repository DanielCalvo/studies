# Shared Terraform Instructions

Use these instructions whenever editing, reviewing, generating, validating, or discussing Terraform in this repository. More specific `AGENTS.md` files may add local rules.

## Structure and Style

- Prefer small, targeted changes and preserve the existing file layout when possible: `versions.tf`, `provider.tf`, `variables.tf`, `locals.tf`, resource-scoped `.tf` files, and `outputs.tf`.
- Keep examples readable for humans first. Favor explicit values and straightforward relationships over clever abstractions.
- Reuse well-known community modules only when they reduce noise and keep the example easier to understand.
- Exception: when creating a VPC, it is fine to use the widely adopted Terraform AWS VPC module instead of defining every VPC resource individually. In real-world Terraform, most teams do not hand-roll every VPC subnet, route table, and gateway resource unless they have a specific reason.
- Do not put everything in `main.tf`. Keep variables, locals, data sources, and outputs in their conventional files, and group resources by domain, such as `vpcs.tf`, `s3.tf`, `iam.tf`, or `networking.tf`.

## Conventions

- Always declare and preserve explicit `required_version` and provider version constraints unless there is a clear reason to change them.
- When adopting a Terraform module for the first time in a folder, check the module's GitHub releases and use the latest available release tag unless the user asks for a different versioning strategy.
- When initializing a brand new provider in a folder, use the latest available provider version at that time. Do not re-check or bump provider versions repeatedly during unrelated edits in the same folder.
- Use typed variables and add descriptions for variables and outputs.
- For values that are very unlikely to change in a given example, such as fixed regions, availability zones, or similar near-constants, prefer explicit literal values over introducing `locals`.
- Use `locals` for genuinely shared computed values or repeated values that benefit from centralization, not for constants that are clearer when written directly.
- Use AWS provider `default_tags` for shared tagging instead of adding `tags` blocks to individual resources or modules unless there is a clear exception.
- Do not name resources, modules, data sources, or other Terraform objects `this`. Use the most intuitive descriptive name for the object instead to keep the configuration readable and avoid naming collisions.
- Keep names and tags consistent within a subfolder.
- When creating a Terraform folder for AI-generated examples, prefix the folder name with `ai_`, such as `ai_ec2_example`, so the source is clearly identified.
- Avoid introducing unnecessary complexity such as dynamic blocks, heavy meta-programming, or deeply nested conditionals unless the example clearly needs them.

## Safety Expectations

- Treat infrastructure changes as real changes even in lab code.
- Do not create, modify, or delete managed infrastructure through direct API calls, CLI commands, console actions, or other out-of-band mechanisms. Terraform must remain the source of truth.
- Read-only inspection of cloud state is allowed when needed for context, validation, or troubleshooting.
- Do not silently widen blast radius by adding high-cost, internet-facing, or destructive resources unless the task clearly requires them.
- Avoid hardcoding secrets, credentials, account IDs, or other sensitive values.
- Prefer inputs for values likely to vary by user, account, or region.
- Call out assumptions when a change depends on existing AWS or Terraform state that is not visible in the folder.

## Validation

- After edits, run `terraform fmt` on the affected subfolder when Terraform is available.
- Run `terraform validate` in the affected subfolder when practical and when initialization requirements are already satisfied.
- If validation cannot be completed, state that clearly and explain why.

## Scope Boundaries

- Do not create shared module coupling between sibling subfolders unless the user explicitly asks for it.
- Do not add CI, remote state, or deployment workflow machinery by default.
