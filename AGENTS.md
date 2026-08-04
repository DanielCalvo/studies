## Here's some context on how we do things around these here parts:
1. If the user asks you to create a Markdown file with the steps that were followed to accomplish something, or to write down notes on something we have investigated together, prefix the file with `ai_`. The exception to this is if a certain AI skill saves files with a defined format or name. In this case, follow the skill and ignore this rule (examples could be a timestamp, a log file under some folder, or any other specific path described in the skill)

2. In the same vein, any file that has the `ai_` prefix on it indicates that it is AI-generated. You can assume the user explored the topic, but do not treat the contents of these files as proof of expertise. On the other hand, if a file does not start with `ai_`, it was user-generated, either dictated or typed

3. Human-generated Markdown files are the most valuable artifacts in this repository and you are not to change them unless explicitly asked to. Any Markdown file that does not start with `ai_` is human-generated

4. The user frequently uses Talon voice-to-text for prompting, which frequently misunderstands words and abbreviations. When strange words appear in a prompt or text file, assume them to be misheard words or abbreviations in the context of DevOps/SRE/infra/cloud. If you ever correct notes, please keep the wording and sentence structure of the author. Only correct issues around punctuation, spacing, grammar, and possible voice-to-text misunderstandings

5. When investigating something (e.g., why did this Helm chart fail to install, or why is this pod not initializing), do not take action (edit files or run commands) on your own. Only take action and change remote systems or files when requested to. Remember our purpose here is to study and learn, not rush into fixing things

6. When troubleshooting or fixing something requires changing a config file, leave a comment at the top of the introduced/changed config briefly explaining why it is in place

- If asked to work with Kubernetes, read [.agents/kubernetes.md](.agents/kubernetes.md)
- If asked to work with Terraform, read [.agents/terraform.md](.agents/terraform.md)
- If asked to work with Go, read [.agents/go.md](.agents/go.md)

All Codex skills are located under [.codex/skills](.codex/skills).
