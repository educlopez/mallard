# Mallard

Claude Code, Codex & OpenCode toolkit — skills, commands, and subagents,
synced with one CLI. Built for the day-to-day work of PrestaShop
development teams.

**[Website & docs → mallard.educalvolopez.com](https://mallard.educalvolopez.com)**

## Related tools

| Tool | What |
|------|------|
| [ps-lando](http://ps-lando.educalvolopez.com/) | CLI to scaffold a Lando environment with PrestaShop + Panda theme |

> **Note**: the PrestaShop/Panda expert agents + knowledge base from
> [prestashop-experts](https://github.com/educlopez/prestashop-experts) are also bundled
> here — see the [Agents](#agents) section below. The standalone plugin stays available
> for external/freelance use; mallard is the team distribution.

## Install

### macOS / Linux (Homebrew)

```bash
brew install educlopez/tap/mallard
```

Upgrade later with `brew upgrade mallard`.

### Windows (Scoop)

```powershell
scoop bucket add educlopez https://github.com/educlopez/scoop-bucket
scoop install mallard
```

Upgrade later with `scoop update mallard`.

### macOS / Linux (curl installer)

```bash
curl -fsSL https://raw.githubusercontent.com/educlopez/mallard/main/install.sh | bash
```

Drops the binary into `~/.local/bin/mallard` (override with `MALLARD_INSTALL_DIR=/usr/local/bin`). Make sure `~/.local/bin` is on your `PATH`.

Pin a specific version:

```bash
MALLARD_VERSION=v0.2.0 curl -fsSL https://raw.githubusercontent.com/educlopez/mallard/main/install.sh | bash
```

### Windows (manual)

Grab the latest zip from [Releases](https://github.com/educlopez/mallard/releases) (`mallard_<version>_windows_amd64.zip` or `_arm64.zip`), extract it, and move `mallard.exe` somewhere on your `PATH`. Or use WSL / Git Bash with the curl installer above.

### After install

```bash
mallard update    # install Claude/Codex/OpenCode skills + commands
mallard doctor    # verify installation
mallard           # launch interactive TUI
```

### CLI

```bash
mallard                        # interactive TUI
mallard install                # same as above
mallard install --agent claude # install only to claude (non-interactive)
mallard install --all          # install to all detected agents (non-interactive)
mallard update                 # re-link skills/commands
mallard doctor                 # check symlink health per agent
mallard registry               # list installed skills/commands
mallard version                # print version
```

The binary auto-detects installed agents (claude, codex, opencode, agents,
gemini, cursor, windsurf) and symlinks the appropriate skills/commands
directories. Per-agent support varies: each agent only receives the kinds it
hosts (e.g. commands link to claude/opencode; sub-agents to claude/cursor).

### From source

```bash
git clone git@github.com:educlopez/mallard.git
cd mallard
go build -o mallard .
./mallard update
```

## Update

Re-run the curl-pipe installer to upgrade the binary, then re-link skills:

```bash
curl -fsSL https://raw.githubusercontent.com/educlopez/mallard/main/install.sh | bash
mallard update             # re-link skills/commands (backs up conflicts)
```

Windows: download the latest zip from [Releases](https://github.com/educlopez/mallard/releases) and replace `mallard.exe`.

`mallard update --list-backups` and `mallard update --restore <ts>` recover prior state if anything went sideways.

## Skills

| Skill | Trigger |
|-------|---------|
| `gitlab-security-setup` | Add Trivy + pnpm supply chain protection to GitLab projects (non-PS) |
| `lando-img-placeholder` | Static image placeholder for local Lando/PrestaShop dev |
| `ps-demo-user` | Create demo user in PrestaShop 8 (Lando) |
| `ps-security-audit` | Weekly CVE scan for PS projects: Friends of Presta module check + Trivy + core version |
| `ps-watch` | BrowserSync live-reload watcher for Panda child theme development in Lando |
| `panda-kb` | Knowledge base for the Panda theme by SunnyToo (st* modules, Easy Builder, demos) |
| `prestashop-kb` | Knowledge base for the PrestaShop 8/9 platform (Symfony, Twig BO, Smarty, hooks, migration) |

## Commands

| Command | What |
|---------|------|
| `/lando` | Lando environment helpers |
| `/ps-customer` | Create test customer in PrestaShop |
| `/ps-url` | PrestaShop URL utilities |
| `/panda` | Ask the `panda-expert` agent (Panda theme + Easy Builder + st* modules) |
| `/ps` | Ask the `prestashop-expert` agent (PS 8/9 core: Symfony, Twig, Smarty, hooks, migration) |

## Agents

Claude Code subagents, symlinked to `~/.claude/agents/` (Claude only). Both are CWD-first
(inspect project source before relying on KB snapshots) and consult the bundled `*-kb` skills.

| Agent | Domain |
|-------|--------|
| `panda-expert` | Panda theme by SunnyToo, `st*` modules, Easy Builder, SunnyToo demos |
| `prestashop-expert` | PrestaShop 8/9 core: themes, parent-child, Symfony BO, Twig, Smarty, hooks, modules, migration 8→9 |

The knowledge base lives in-repo under `skills/panda-kb/references/` and
`skills/prestashop-kb/references/` — that's the source of truth. Edit it directly; no
external vault or per-machine setup is required to build or ship the KB.

## Structure

```
main.go                       Entry point + CLI routing
cmd/
  install.go                  Install (TUI + non-interactive modes)
  update.go                   Update with backup, restore, list-backups
  doctor.go                   Doctor (thin delegate to internal/reports)
  registry.go                 Registry (thin delegate to internal/reports)
internal/
  agents/
    adapter.go                Multi-agent Adapter interface
    claude.go codex.go        Per-agent implementations
    opencode.go generic.go
    gemini.go cursor.go windsurf.go
    registry.go               All() / ByID() factory
  backup/backup.go            Snapshot/Restore + manifest, keep-latest-5 GC
  updater/updater.go          Pure Run(Options) → Report (CLI + TUI share it)
  reports/                    Writer-based doctor + registry renderers
  skillregistry/registry.go   Parse SKILL.md / command frontmatter
  skills/skills.go            Symlink helpers
  tui/
    model.go                  Bubbletea model + screen state machine
    welcome.go                Welcome screen calqued from gentle-ai
    logo.go                   Braille-art duck + yellow gradient
    styles.go                 Lipgloss palette (Rose Pine + duck accents)
skills/                       Claude Code skills (symlinked to ~/.claude/skills/)
claude/commands/              Slash commands (symlinked to ~/.claude/commands/)
claude/agents/                Claude subagents (symlinked to ~/.claude/agents/, Claude only)
.github/workflows/
  release.yml                 Tag push → goreleaser → binaries + brew + scoop
  ci.yml                      go vet + go build on push/PR
  pr-check.yml                400-line PR review-budget check
.goreleaser.yaml              Multi-platform build + brews/scoops publish
install.sh                    Curl-pipe installer (downloads release binary)
doctor.sh                     Bash dependency check (claude CLI, node, pnpm)
docs/                         Notes on adding skills and commands
```

## Adding a skill

See `docs/adding-skills.md`.
