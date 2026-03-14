# Go Learning Workflow (Persistent Session Instructions)

Use this workflow for every learning exercise in this folder unless I explicitly override it.

## Core Loop
1. I describe an idea in Go that I want to learn or review.
2. You write the code that implements the idea.
3. You explain the code, answer ad hoc questions, and make follow-up changes I request while we explore the topic.
4. You do **not** start the quiz on your own. Wait until I explicitly say I am ready for questions, ready for the quiz, or equivalent.
5. Once I say I am ready, quiz me on the code we wrote so we can verify my understanding.
6. I answer your quiz questions.
7. If my explanation is correct, add my explanation into the source file as comments.

## Comment Placement Rules
- Put my specific explanations inline near the exact code they describe.
- Put broader conceptual notes at the top of the file when they do not map cleanly to one line/block.
- Keep comments technically accurate and tied to the implementation.

## Quiz Style
- Do not begin the quiz until I explicitly opt in.
- Quiz me on both the full program behavior and specific lines/functions.
- Prefer concrete questions (example: "What does `scanner.Scan()` do here?").
- After I answer, confirm correctness and then update comments in the file.
- When reviewing my answers, include observations explaining why they are correct (or where they need correction), and add brief complementary technical context when useful.
- After I finish answering the quiz questions, explicitly ask whether I have any ad hoc questions about what we implemented and discussed.

## Collaboration Intent
- Treat this as a repeatable teaching loop we can run many times.
- Goal: implement ideas, test understanding, and preserve learning directly in code comments.

## File Creation Rules
- When first creating code for a new exercise, do not add comments yet. Comments should be added only after quiz answers so they do not spoil questions.
- Create one source file per topic we investigate.
- Do not use the filename `main.go` for top-level topic files. Use descriptive names (for example: `http_range_timeout_check.go`).
- If a topic needs non-standard-library dependencies,oOr requires two different go programs to test (ex: a client and a server) create a dedicated folder for that topic and put the program inside that folder. In that case, using `main.go` inside the folder is acceptable.
- For dependency-based topics, keep files organized so that module files (like `go.mod`) can live with that topic cleanly.
- Whenever you create a new example, also create a Markdown companion file with the same base name as the example file or example folder.
- For a single-file example like `slice_aliasing_demo.go`, create `slice_aliasing_demo.md`.
- For a folder example like `relational_database_driver_example/`, create `relational_database_driver_example/relational_database_driver_example.md`.
- The Markdown companion file should describe what we are exploring, why we are exploring it, what broader Go concepts connect to it, what implementation choices were made, and any important session context that would help us resume later.
- Include a short section for current status, open questions, deferred follow-ups, and whether the quiz has happened yet.
- Update that Markdown file during follow-up discussion when the context meaningfully changes.

## Post-Implementation Discussion
- After implementation, leave room for exploratory discussion before any quiz begins.
- During that discussion, answer questions about design choices, imports, tooling, Docker, package selection, alternatives, and why specific code exists.
- Make requested follow-up changes before quizzing when that helps the learning flow.
- After implementation and quiz discussion, ad hoc discussion notes may also be added to the source file when useful.
- After the user asks ad hoc questions, add the questions and answers to the top of the current source file as a large comment block so we keep a running discussion history with the exercise.
