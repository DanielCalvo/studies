---
name: fix-markdown
description: Clean newly added Markdown lines that were written through mixed typing and Talon voice dictation. Use when the user asks to fix, clean, polish, normalize formatting, or review dictated Markdown notes in a tracked file, especially when only uncommitted added lines should be edited and existing, deleted, or merely changed content should be left alone.
---

# Fix Markdown

## Goal

Clean only the newly added Markdown text in a target file. Treat the file as personal study notes, not formal documentation: preserve the user's voice, wording, intent, order, headings, bullets, code blocks, commands, links, and technical meaning while correcting dictation artifacts and applying the user's Markdown formatting preferences.

## Required Input

Require a target file path from the user. If the user invokes the skill without a file path, ask for the file before editing.

## Workflow

1. Identify the repository root with `git rev-parse --show-toplevel`.
2. Inspect the target file's local diff against `HEAD` with `git diff -- <path>`.
3. Work only on added lines from the diff: lines beginning with `+` that are not diff metadata (`+++`, hunk headers, etc.).
4. Ignore deleted lines and unchanged context lines.
5. For modified lines, edit only the newly added replacement line when it clearly contains dictation cleanup issues. Do not try to reconstruct deleted text.
6. Apply focused edits to the file itself, keeping surrounding untouched content stable.
7. Re-run `git diff -- <path>` and verify the final diff only changes intended added-line cleanup.

When the user explicitly asks to review before applying changes, show proposed edits or a concise summary and do not modify the target file.

## What To Fix

Correct obvious Talon or voice-to-text artifacts:

- Misheard Kubernetes and tooling terms, such as `cube ADM` -> `kubeadm`, `cube proxy` -> `kube-proxy`, `containerD` -> `containerd`, `redis` -> `Redis`, `docker` -> `Docker` when referring to the product, and `apis` -> `APIs`.
- The user's name when it is misheard in notes headings or text: use `Dani`, not `Danny` or `Daniel`, when the note clearly refers to the user's own notes such as `Notes from Dani`.
- Dictation punctuation errors, such as `doesn-t`, `it-s`, `here-s`, and missing spaces after commas, periods, colons, or question marks.
- Obvious word substitutions when context is clear, such as `rights` -> `writes`, `right` -> `write`, `the bug` -> `debug`, `speck` -> `spec`, `notes` -> `nodes` when discussing Kubernetes nodes, and `comfort map` -> `ConfigMap`.
- Extra spaces within lines, leading spaces in normal prose, and accidental double spaces.
- Too many blank lines in a row. Prefer a single blank line between paragraphs unless a code block requires exact spacing.
- Missing Markdown heading spaces, such as `####Heading` -> `#### Heading`.
- Sentence starts and proper nouns when the fix is mechanical and low risk.
- Spoken version numbers or numeric values when the intended number is clear from context, such as `Kubernetes one point thirty three` -> `Kubernetes 1.33`, `version one point twenty nine` -> `version 1.29`, and `fifty percent` -> `50%`.
- Voice-to-text backtracks where the user dictated a correction, such as `thing A, no wait sorry I mean thing B`; keep only the corrected meaning (`thing B`) when the intended final wording is obvious.

## Formatting Preferences

Apply these preferences to added Markdown lines:

- Do not leave a blank line between a Markdown heading and the first content line under it. A heading on line 1 should be followed immediately by text, a bullet, or other content on line 2.
- Use four spaces for nested list indentation. Top-level bullets use `- `; subitems use exactly four leading spaces before `- `. Normalize one, two, three, or other odd indentation widths to four spaces when the item is clearly nested.
- List items should not end with a period. Remove a final period from bullet items unless the line is a code command, URL, abbreviation, or literal text where the period is meaningful.
- Keep question marks and exclamation marks at the end of list items when they are intentional.

## What Not To Change

Do not broadly rewrite the notes. Avoid changing:

- The user's explanatory style, informal phrasing, or first-person notes unless they are hard to understand.
- The user's wording or sequence of words merely to make the note prettier, smoother, more formal, or more "AI-written."
- The word `Gepeto`; this is the user's intentional affectionate name for ChatGPT and should be preserved.
- Technical meaning, commands, YAML, code blocks, inline code, URLs, file paths, API names, resource names, flags, or literal command output unless the error is unmistakable.
- Existing committed content outside added diff lines.
- Deleted lines from the diff.
- Changed lines only for style if there is no dictation or correctness issue.

If a word looks wrong but the intended correction is uncertain, leave it unchanged or ask instead of guessing.
If a pass over the file finds incoherent voice-to-text text, unknown abbreviations, or likely misheard words that cannot be confidently reconstructed, leave those terms unchanged and ask the user about them at the end of the pass. Include enough surrounding context for each question that the user can identify the original note.

## Editing Guidelines

Prefer small, local edits. Keep the corrected line recognizably the same note.

Examples:

- `With these spec in mind anyone can build a container runtime` -> `With these specs in mind, anyone can build a container runtime`
- `crictl is what you would use on a Kubernetes node to the bug containers` -> `crictl is what you would use on a Kubernetes node to debug containers`
- `if either node dies, there is no majority` should stay as-is because it is already clear.
- `Gepeto says` should stay as `Gepeto says`.
- `- keep rollout history.` -> `- keep rollout history`
- `  - nested item` -> `    - nested item`
- `### Topic\n\nFirst sentence` -> `### Topic\nFirst sentence`
- `I mean the scheduler, no wait sorry I mean the kubelet creates the pod` -> `the kubelet creates the pod`

## Verification

Before finishing, report:

- The target file reviewed.
- Whether edits were applied or only proposed.
- The validation command used, usually `git diff -- <path>`.
- Any uncertain terms left unchanged.
