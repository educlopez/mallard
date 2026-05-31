# Adding a skill

1. Create `skills/<skill-name>/SKILL.md` in this repo
2. Frontmatter required:
   ```yaml
   ---
   name: skill-name
   description: >
     When to trigger this skill. Be specific — Claude uses this to decide
     whether to invoke the skill. Include contexts, project types, keywords.
   version: "0.1.0"
   ---
   ```
3. Locally (during development):
   ```bash
   go build -o mallard .
   ./mallard update           # re-link skills into ~/.claude/skills (with backup)
   ./mallard doctor           # confirm the new skill appears under "managed"
   ./mallard registry         # check version + drift status
   ```
4. Ship to the team:
   ```bash
   git add skills/<skill-name>/
   git commit -m "feat: add <skill-name> skill"
   git push
   git tag vX.Y.Z
   git push origin vX.Y.Z      # GH Actions builds + publishes Homebrew/Scoop/binaries
   ```
   Team members pick it up with `brew upgrade mallard` (or `scoop update mallard`) followed by `mallard update`.

## Adding a command

1. Create `claude/commands/<command-name>.md`
2. Frontmatter required:
   ```yaml
   ---
   name: command-name
   description: One-line summary of what the command does
   version: "0.1.0"
   ---
   ```
3. Available in Claude Code as `/<command-name>` after running `mallard update`.

## Adding an agent

Agents are Claude Code subagents, symlinked to `~/.claude/agents/` (Claude only —
codex/opencode/generic adapters skip them, since the format is Claude-specific).

1. Create `claude/agents/<agent-name>.md`
2. Frontmatter required:
   ```yaml
   ---
   name: agent-name
   description: When to invoke this agent (Claude routes to it based on this).
   tools: Read, Grep, Glob, Bash
   model: sonnet
   color: orange
   version: "0.1.0"
   ---
   ```
3. Available as a subagent after `mallard update`. A companion `/<name>` command in
   `claude/commands/` is the usual way to invoke it explicitly.

## Skill scope

| Scope | Where |
|-------|-------|
| Reusable across projects | `skills/` here |
| Specific to one project | `.claude/commands/` inside that project's repo |
| Personal preferences | `~/.claude/` (not tracked here) |

## Bumping a skill version

Just edit the `version:` field in the skill's frontmatter and ship a new release. `mallard registry` will surface drift between installed and source versions per agent.
