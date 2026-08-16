# mallard — Backlog

gentle-ai-inspired borrows worth doing (mallard stays a focused skills/commands/agents
sync tool — the heavier orchestration layer (SDD, persona, MCP mgmt, model routing) is
deliberately out of scope).

## Done

- **Backup dedup** — SHA-256 content-addressed dedup in `internal/backup`; identical content stored once, hard-linked across sessions. Landed in commit 6ab84f9.
- **More agents** — extended beyond claude/codex/opencode/generic; additional adapters (gemini, cursor, windsurf, and others) added to the registry. Landed in commit 6ab84f9.
- **Persisted state (state.json)** — `~/.mallard/state.json` for user selections; graceful on missing/corrupt. Landed in commit 6ab84f9.

## Explicitly NOT doing (changes the product)
SDD workflow, Engram/persona injection, MCP server management, AI provider switcher /
per-phase model assignment. These belong to a "configurator" product, not mallard.

Commerce (see `docs/ecommerce-agent-team.md`) may steal gentle-ai *mental models*
only: organic routing (inline / delegate / ceremony), delegate-only orchestrator,
and a receipt+apply gate. It must not import SDD phases, Engram, GGA, persona,
or the RDD review lifecycle.
