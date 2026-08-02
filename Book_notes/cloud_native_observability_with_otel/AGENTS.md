These are my study notes from reading Cloud Native Observability with OpenTelemetry.

I'm mostly focusing on:
- Concepts and terminology the book explains and the fundamentals of OpenTelemetry so I can understand and use the technology as a site reliability engineer
- Things that the book describes that I don't know and should investigate (like exemplars)

The book is available in this folder for reference as needed at [Cloud-NativeObservabilitywithOpenTelemetry.pdf](Cloud-NativeObservabilitywithOpenTelemetry.pdf)

## Incremental study workflow

When helping me study a chapter or implement an example, use an incremental,
dialogue-driven approach.

1. Begin with the smallest ordinary program or example that I can understand
   without the technology being studied.

2. Introduce one new concept at a time. Prefer one small code change per step.
   If multiple components must be introduced together to make a coherent step,
   explain why they belong together.

3. Before changing code, explain:
   - The single concept being introduced
   - Why it is the next logical step
   - The minimal code change required
   - What observable difference I should expect

4. Wait for me to confirm that I want the step implemented unless I have already
   explicitly requested implementation.

5. After implementing a step:
   - Add comments near the new code explaining the concept
   - Preserve my existing comments and edits
   - Format and compile/test the code
   - Briefly report what changed and what was verified

6. Maintain an `ai_incremental_steps_taken.md` file in the example's directory.
   After every implemented step, append a numbered section describing:
   - What was changed
   - The concept it demonstrates
   - A small relevant code snippet
   - The expected output or behavior
   - Important distinctions from related concepts
   - What has deliberately not been introduced yet

7. When I paraphrase my understanding, verify it precisely. Correct only the
   necessary details and explain important terminology distinctions without
   jumping ahead to another implementation step.

8. When I ask "what's next?", propose one next step rather than presenting an
   entire implementation plan. Mention later topics only briefly for context.

9. Do not introduce production complexity—such as Collectors, backends,
   distributed deployments, or advanced configuration—until the simpler local
   concept is understood.

10. Prefer examples where each addition can be observed directly in program
    output, generated telemetry, or a small behavioral change.

11. Use the book for the concepts being studied. Because its implementation
    examples may be dated, verify current library APIs against the installed
    dependency version or current official documentation when necessary.

The objective is understanding, not reaching the final implementation as quickly
as possible. A partially implemented example that clearly demonstrates the
current concept is preferable to introducing several concepts simultaneously.
