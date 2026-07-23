## Here's some context on how we do things around these here parts:

1. If the user asks you to create a markdown file with the steps that were followed to accomplish something, or to write down notes on something we have investigated together, prefix the file with "ai_". The exception to this is if a certain AI skill saves files with a defined format or name. In this case, follow the skill and ignore this rule. (examples could be a timestamp, a log file under some folder, or any other specific path described in the skill). 

2. In the same vein, any file that has the "ai_" prefix on it indicates that it is AI generated. You can assume the user explored the topic, but do not treat the contents of these files as proof of expertise. On the other hand, if a file does not start with "ai_" it was user generated, either dictated or typed.

3. The user frequently uses talon voice to text for prompting, which frequently misunderstands words and abbreviations. When strange words appear in a prompt or text file, assume them to be misheard words or abbreviations in the context of devops/sre/infra/cloud. If you ever correct notes, please keep the wording and sentenced structure of the author. only correct issues around punctuation, spacing, grammar and possible voice to text the misunderstandings.

4. When investigating something (ex: why did this helm chart fail to install, or why is this pod not initializing) do not take action (edit files or run commands) on your own. Only take action and change remote systems or files when requested to. Remember our purpose here is to study and learn, not rush into fixing things.

5. When troubleshooting and having to change a config file to make things work, troubleshoot or fix something, leave a comment on top of the introduced/changed config briefly explaining why this is in place

- If asked to work with Kubernetes, read [.agents/kubernetes.md](.agents/kubernetes.md)
- If asked to work with Terraform, read [.agents/terraform.md](.agents/terraform.md)
- If asked to work with Go, read [.agents/go.md](.agents/go.md)

All codex skills are located under [.codex/skills](.codex/skills)
