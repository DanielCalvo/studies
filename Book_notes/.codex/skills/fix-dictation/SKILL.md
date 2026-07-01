---
name: fix-dictation
description: Clean up voice-to-text dictated notes while preserving the author's own wording, structure, tone, and note-taking style. Use when editing Markdown or plain-text notes that contain dictation artifacts, misheard technical terms, punctuation/spacing problems, filler backtracking such as "sorry, I mean...", or domain-specific misunderstandings in DevOps, cloud platforms, Kubernetes, observability, SRE, infrastructure, software engineering, and related book-study notes.
---

# Fix Dictation

## Goal

Edit dictated notes so they read like the author wrote them clearly the first time. Preserve the author's voice. Do not turn the notes into polished AI prose, a summary, or a rewritten article.

## Workflow

1. Read the target file before editing.
2. Infer the subject matter from the file, surrounding repository, filenames, headings, and nearby code/config when useful.
3. Correct obvious dictation artifacts in place:
   - Misheard technical terms, product names, acronyms, and commands.
   - Broken punctuation, capitalization, spacing, Markdown headings, and code spans.
   - Missing or duplicated words caused by voice-to-text.
   - Cut-off sentence fragments when the intended phrase is clear from context.
   - Backtracking phrases where the author clearly corrected themselves, such as "sorry, I mean X" or "no wait, X".
4. Keep edits narrow when the author's intent is ambiguous.
5. Review the diff for over-polishing before finishing.

## Preserve The Author's Voice

Keep the author's sentence structure and casual note-taking rhythm where possible. Preserve first-person thoughts, uncertainty, reactions, and transitions such as "Right", "So", "Well", "Nice", "I think", "which makes sense", and similar phrasing.

Do not:

- Reframe the note into a tutorial or documentation.
- Replace casual wording with formal prose unless it fixes a real transcription error.
- Remove opinions, uncertainty, or asides just because they are informal.
- Summarize, reorganize, or add new explanations.
- Add facts that are not already implied by the note or needed to repair a misheard term.

## DevOps And Cloud Dictation Fixes

Prefer domain-correct terms when context supports them, for example:

- `Prometheus`, `PromQL`, `Alertmanager`, `Pushgateway`, `Node Exporter`, `promtool`
- `Kubernetes`, `kubectl`, `Helm`, `Ingress`, `Service`, `Pod`, `Deployment`
- `AWS`, `EC2`, `S3`, `IAM`, `GCP`, `Azure`, `Consul`
- `SRE`, `SLI`, `SLO`, `SLA`, `RED`, `USE`
- `cache hit`, `cache miss`, `backend`, `database`, `worker pool`
- `counter`, `gauge`, `histogram`, `summary`, `bucket`, `time series`, `cardinality`

Use code formatting for commands, metric names, label matchers, filenames, and short query examples when doing so clarifies the note without changing its style.

## Backtracking Cleanup

When the author clearly corrects themselves, keep the final intended phrase and remove the false start. For example:

- "push gateway sorry I mean Pushgateway" -> "Pushgateway"
- "five golden signals no wait four golden signals" -> "four golden signals"
- "cash hit sorry cache hit" -> "cache hit"

If the correction is not obvious, keep the wording and only fix punctuation.

## Final Check

Before responding, check for:

- Remaining obvious dictation artifacts such as malformed contractions, repeated spaces, or broken technical names.
- Accidental tone changes where a casual note now sounds like generated documentation.
- Markdown formatting issues, especially headings, list spacing, inline code, and code blocks.

When reporting back, briefly state that the file was cleaned while preserving the author's wording and tone.
