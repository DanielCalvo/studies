---
name: save-conversation
description: Save a recent conversation topic as a standalone Markdown reference note containing the user's reconstructed question and the substantive answer. Use when the user invokes `$save-conversation`, asks to save the current or last discussion for later, or specifies conversation boundaries or a topic to preserve from recent turns.
---

# Save Conversation

Turn a recent discussion into a durable, self-contained note in the Studies
repository. Preserve meaning and useful technical detail rather than producing
a raw chat transcript.

## Select the conversation scope

1. Follow explicit topic or turn boundaries from the user.
2. Otherwise select the most recent substantive user question and its answer.
   Ignore the current meta-request to save the conversation.
3. Expand backward or forward across adjacent turns only when they are clearly
   follow-up questions about the same underlying topic and are needed for a
   coherent note.
4. Stop at a genuine topic change. Do not combine unrelated discussions merely
   because they are recent.
5. When ambiguity would materially change the saved topic, ask one concise
   question. Otherwise choose the narrowest coherent scope and proceed.
6. If the requested discussion is no longer present in conversation context,
   ask the user to provide it. Do not reconstruct missing content from guesses.

## Reconstruct the question

- Include a `## Question` section even when the discussion covered several
  turns.
- Rewrite dictated text into a clear, coherent question while preserving the
  user's intent, assumptions, and level of detail.
- Correct likely voice-to-text errors using the technical context.
- Do not add requirements or claim the cleaned question is a verbatim quote.
- Combine closely related follow-ups into the question or add concise
  subquestions when that makes the scope clearer.

## Preserve the answer

- Include a `## Answer` section containing the substantive response.
- Make the answer understandable without access to the chat history.
- Preserve important reasoning, alternatives, tradeoffs, caveats, commands,
  configuration examples, diagrams, and source links from the discussion.
- Synthesize connected answers into one organized explanation and remove chat
  acknowledgements, repetition, and save-file logistics.
- Keep uncertainty and qualifications. Do not turn a tentative conclusion into
  a fact.
- Do not perform new research or broaden the subject unless the user asks.

Use additional sections such as `## Context`, `## Decision`, or
`## References` only when they materially improve the note. The required core
remains the reconstructed question and answer.

## Resolve the destination portably

Treat the directory three levels above this skill directory as the Studies
repository root:

```text
<studies-root>/.codex/skills/save-conversation/SKILL.md
<studies-root>/ai_notes/
```

Confirm the skill marker exists at that relationship. When the skill location
is unavailable, use the current Git worktree root only if it contains
`.codex/skills/save-conversation/SKILL.md`. Do not embed a machine-specific
absolute path in commands, notes, or skill instructions.

Read and follow any `AGENTS.md` files applicable to the destination. Create
`ai_notes/` if it is missing.

## Name and write the note

1. Choose a concise descriptive title.
2. Derive a lowercase snake-case filename with the required AI-generated
   prefix: `ai_<topic>.md`.
3. Inspect an existing destination before writing. Never overwrite an unrelated
   note. If the natural filename already exists, use a date or numeric suffix
   unless the user explicitly asked to update that file.
4. Write ordinary Markdown using only as much structure as the subject needs.
5. Verify that the saved file is readable and contains both `## Question` and
   `## Answer`.
6. Report the resulting path to the user. Mention an updated existing note
   explicitly when applicable.

Prefer a useful reference document over a turn-by-turn transcript. Preserve
the intellectual context: what prompted the question, why the answer follows,
and which constraints made the chosen design appropriate.
