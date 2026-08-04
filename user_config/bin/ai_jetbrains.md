# JetBrains IDE external-tool integration map
Use this note when adding, changing, or troubleshooting an external tool in a JetBrains IDE on this machine, including GoLand. It describes the local structure and the repeatable workflow rather than any one tool's behavior.

## Where things live
- GoLand installation: `/home/daniel/IDEs/GoLand-2026.2/`
- GoLand user configuration: `/home/daniel/.config/JetBrains/GoLand2026.2/`
- GoLand external-tool definitions: `/home/daniel/.config/JetBrains/GoLand2026.2/tools/External Tools.xml`
- Version-controlled user scripts: `/home/daniel/Projects/studies/user_config/bin/`
- Commands exposed on the user's `PATH`: `/home/daniel/bin/`

External-tool definitions are stored in the user configuration directory, not in the IDE installation directory. The XML is not currently version controlled in this repository.

Scripts intended for interactive use should live in this repository directory, be executable, and be linked into `/home/daniel/bin/`. The repository copy is the source of truth; the home-bin entry should be a symbolic link to it.

## Adding a new script-backed tool
1. Create or update the script in this directory. Give it a clear command-line contract: required arguments, optional arguments, defaults, supported working directories, and errors for invalid input
2. Make the repository script executable and create or update its symbolic link in `/home/daniel/bin/`
3. Test the script directly in a terminal before integrating it with GoLand. Use safe flags such as `--help` when they are available
4. Decide which GoLand context variable supplies each input. Common choices are:

   | Need | GoLand macro |
   | --- | --- |
   | Current file path | `$FilePath$` |
   | Current file directory | `$FileDir$` |
   | Current file name | `$FileName$` |
   | Project root | `$ProjectFileDir$` |
   | User home directory | `$USER_HOME$` |

5. Add the tool through **Settings | Tools | External Tools**, or add a `<tool>` entry to `External Tools.xml` manually, pointing `COMMAND` at the script or at a terminal application that then executes the script
6. Apply the settings when using the GoLand UI. When editing the XML manually, restart GoLand. Then invoke the tool against a representative file or project

## Editing `External Tools.xml` safely
The following workflow applies when editing the file directly instead of managing tools through the GoLand settings UI:

1. Fully close GoLand before editing the file. It may otherwise rewrite settings when it exits
2. Copy the live XML to a temporary file and make all edits there first
3. Keep each tool definition inside the root `<toolSet name="External Tools">` element
4. Validate the temporary XML before installing it:

   ```bash
   python3 -c 'import xml.etree.ElementTree as ET; ET.parse("/path/to/External Tools.xml")'
   ```

   `xmllint` is not installed on this machine; Python's built-in XML parser is the available validation method.
5. Back up the live XML with a dated `External Tools.xml.before-...` filename beside it
6. Replace the live XML with the validated temporary copy, preserving its file permissions
7. Restart GoLand to load the definitions

Use `&quot;` for double quotes and `&amp;` for ampersands inside XML attribute values.

## Generic external-tool shape
For clarity, give each definition a unique `name`. The basic definition has an `<exec>` section containing the command, its arguments, and the working directory:

```xml
<tool name="Descriptive Tool Name" description="What this tool does" showInMainMenu="false" showInEditor="false" showInProject="false" showInSearchPopup="false" disabled="false" useConsole="true" showConsoleOnStdOut="false" showConsoleOnStdErr="false" synchronizeAfterRun="true">
  <exec>
    <option name="COMMAND" value="$USER_HOME$/bin/example-tool" />
    <option name="PARAMETERS" value="$FilePath$" />
    <option name="WORKING_DIRECTORY" value="$FileDir$" />
  </exec>
</tool>
```

For a tool that needs its own visible terminal, use `/usr/bin/xfce4-terminal` as `COMMAND`. Provide `--default-working-directory=&quot;$FileDir$&quot;` in `PARAMETERS`, then use `--execute` followed by the home-bin script and its arguments. This preserves the GoLand-selected directory while letting the user interact with the command.

## Existing examples
`External Tools.xml` already contains examples of these patterns:

- `Fix Markdown` calls a home-bin script directly and passes `$FilePath$`
- `Open terminal here` starts XFCE Terminal in `$FileDir$`
- The `Open Codex` tools start XFCE Terminal in `$FileDir$`; the Terra and Luna variants pass model-selection arguments to the home-bin script, while the default tool relies on the script's Sol default

Read their current XML entries as implementation examples, but do not modify unrelated tools when adding or changing another one.

## Verification checklist
- Script syntax or lint checks pass
- The repository script is executable
- The `/home/daniel/bin/` entry resolves to the repository script
- The edited XML parses successfully
- The new definition has a clear name that does not unintentionally duplicate another tool
- The tool receives the intended macro values and starts in the intended directory
- GoLand is restarted before UI verification when the XML was edited manually
