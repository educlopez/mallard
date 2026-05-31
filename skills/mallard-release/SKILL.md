---
name: mallard-release
description: >
  Workflow for adding skills/commands to mallard and shipping a new release.
  Use when the user asks how to release mallard, ship a new mallard version, add a
  skill or command to mallard, bump a mallard version, "publicar mallard", "subir
  mallard", "nueva versión de mallard", or anything involving editing this repo
  and getting the change to the team's machines.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# mallard release workflow

Reference for adding/updating skills, commands, and shipping releases for [mallard](https://github.com/educlopez/mallard). The tool lives at `/Users/eduardocalvolopez/Developer/Personal/mallard` (on Eduardo's machine).

## When to use this skill

- "How do I add a skill to mallard?"
- "Release a new version of mallard" / "ship mallard" / "tag mallard vX.Y.Z"
- "Update mallard" / "bump the mallard version"
- Anything that involves editing files under `skills/`, `claude/commands/`, `claude/agents/`, or shipping the change to other team members.

## TL;DR

```bash
# 1. Edit the skill/command file
# 2. Bump its `version:` in frontmatter (semver, conservative)
# 3. Test locally
go build -o mallard .
./mallard update                    # re-link with backup
./mallard registry                  # confirm new version + ok status

# 4. Commit using Conventional Commits, in English
git add skills/<name>/ claude/commands/<name>.md
git commit -m "feat(skills): add <name>"   # or fix:, refactor:, docs:, chore:

# 5. Push
git push

# 6. Tag a release (this is what propagates to the team)
git tag vX.Y.Z
git push origin vX.Y.Z

# GitHub Actions runs goreleaser → publishes binaries + auto-changelog
```

Team members pick it up with:

```bash
curl -fsSL https://raw.githubusercontent.com/educlopez/mallard/main/install.sh | bash
mallard update
```

## Adding a new skill

1. Create `skills/<skill-name>/SKILL.md` with frontmatter:
   ```yaml
   ---
   name: <skill-name>
   description: >
     When to trigger this skill. Be specific — Claude uses this to decide
     whether to invoke. Include keywords the user might say, project types,
     and contexts.
   version: "0.1.0"
   ---
   ```
2. Write the skill body in markdown. Sections, code blocks, tables — whatever helps Claude execute the skill. Reference existing skills in this repo for style.
3. Locally: `go build -o mallard . && ./mallard update` — the symlink will appear in `~/.claude/skills/<skill-name>/`.
4. Validate: `./mallard doctor` (managed count goes up) and `./mallard registry` (entry shows `v0.1.0 ok`).
5. Skip to **Shipping a release** below.

## Adding a new command

1. Create `claude/commands/<command-name>.md` with frontmatter:
   ```yaml
   ---
   name: <command-name>
   description: One-line summary of what the command does.
   version: "0.1.0"
   ---
   ```
2. Body: instructions for the slash command. Available in Claude Code as `/<command-name>` after install.
3. Locally: `./mallard update`.
4. Skip to **Shipping a release** below.

## Adding a new agent

Agents are Claude Code subagents. They symlink to `~/.claude/agents/` and are **Claude-only**
(codex/opencode/generic adapters return an empty agents dir and skip them).

1. Create `claude/agents/<agent-name>.md` with frontmatter:
   ```yaml
   ---
   name: <agent-name>
   description: When Claude should route to this agent.
   tools: Read, Grep, Glob, Bash
   model: sonnet
   color: orange
   version: "0.1.0"
   ---
   ```
2. Body: the agent's system prompt. Optionally add a `/<name>` command in `claude/commands/`
   to invoke it explicitly.
3. Locally: `./mallard update` — the symlink appears in `~/.claude/agents/<agent-name>.md`.
4. Skip to **Shipping a release** below.

## Bumping an existing skill, command, or agent

1. Edit the file.
2. Bump `version:` in the frontmatter. Semver guidance:
   - Patch (0.1.0 → 0.1.1): typo fixes, doc clarifications, no behavior change.
   - Minor (0.1.0 → 0.2.0): new functionality inside the skill, backward-compatible.
   - Major (0.1.0 → 1.0.0): breaking change in expected user-facing behavior (rare for a skill).
3. Locally: `./mallard update && ./mallard registry` — confirm new version shows `ok`, old `drift` if you also have an older symlink elsewhere.

## Shipping a release

```bash
git add <files>
git commit -m "feat(skills): add <name>"
# Conventional Commit prefixes: feat, fix, refactor, docs, chore, ci, test
git push
git tag vX.Y.Z
git push origin vX.Y.Z
```

Version the **mallard release**, not the skill: skills carry their own versions inside frontmatter, the tag versions the binary + the set of skills shipped together.

GitHub Actions (`.github/workflows/release.yml`) runs goreleaser, which:
- Builds for darwin/linux/windows × amd64/arm64
- Uploads tarballs + checksums to the GitHub release
- Writes a release changelog grouped by Features / Bug Fixes / Refactor / Others (filters out `docs:`, `test:`, `ci:`, `chore:`)

If the changelog auto-text needs richer notes (highlights, breaking changes, etc.), edit the release after publish:

```bash
gh release edit vX.Y.Z --notes "$(cat <<'EOF'
## Highlights
- ...

## Skills
- new: ...
- updated: ...
EOF
)"
```

## Team upgrade path

```bash
curl -fsSL https://raw.githubusercontent.com/educlopez/mallard/main/install.sh | bash
mallard update     # re-symlinks skills + commands, backs up any conflicts
mallard doctor     # verify
mallard registry   # see what's installed per agent + versions
```

Windows: download the latest zip from https://github.com/educlopez/mallard/releases and replace `mallard.exe` on the PATH.

## Gotchas

- **NEVER skip the `version:` field** on skill/command frontmatter. Without it, `mallard registry` reports `unversioned`.
- **Conventional Commits are mandatory** — they drive the changelog grouping. `chore:` commits get filtered out of release notes.
- **`mallard update` makes backups** before replacing files. If you blow something away, `mallard update --list-backups` and `--restore <ts>` will recover it.
- **The first run on a new agent dir** (Claude / Codex / OpenCode / generic) appears as `missing → installed` for every entry. That's expected, not drift.
- **Skills only ship via tags.** Pushing to `main` without tagging publishes nothing — GH Actions release workflow only fires on `v*` tags.

## Useful commands

| Command | What |
|---------|------|
| `./mallard` | Interactive TUI (Install / Update / Doctor / Registry / Quit) |
| `./mallard update --dry-run` | Show what update would change without touching disk |
| `./mallard update --restore <ts>` | Recover from a backup batch |
| `./mallard registry --source` | List what the repo ships (without scanning installed) |
| `./mallard registry --all` | Include foreign / unmanaged entries from other tooling |
| `./mallard version` | Print version (release tag injected at build time) |
