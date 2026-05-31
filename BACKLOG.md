# mallard — Backlog

gentle-ai-inspired borrows worth doing (mallard stays a focused skills/commands/agents
sync tool — the heavier orchestration layer (SDD, persona, MCP mgmt, model routing) is
deliberately out of scope).

## 1. Backup dedup (low effort)
Add content-addressed dedup to `internal/backup`: SHA-256 each snapshotted file; skip
re-storing identical content across backup sessions (store once, reference). Completes the
backup system (we already have tar.gz + pin + keep-5 GC). Ref: gentle-ai `internal/backup`.
- Acceptance: identical configs across runs don't duplicate on disk; restore still byte-identical; pin/GC unaffected; tests for dedup hit/miss.

## 2. More agents (low–med effort)
Broaden beyond claude/codex/opencode/generic. Candidates: gemini, cursor, windsurf,
(others gentle-ai supports). Each = a new `agents.Adapter` impl with the right global +
workspace dirs (`SkillsDir`/`CommandsDir`/`AgentsDir` + `*For(scope,ws)`); add to the
registry. Skip agents that can't host our markdown skills/commands meaningfully.
- Acceptance: detected if installed; install/update/uninstall/doctor honor them in both scopes; agents-dir only where the agent supports the claude-style format; tests per adapter.

## 3. Persisted state (state.json) (med effort — only if the TUI grows)
A small `~/.mallard/state.json` for user selections (which agents/skills last chosen,
preferences). Graceful on missing/corrupt. Ref: gentle-ai `internal/state`. Low priority
until the TUI does more than one-shot install — revisit then.
- Acceptance: survives across runs; never blocks a run if absent; covered by tests.

## Explicitly NOT doing (changes the product)
SDD workflow, Engram/persona injection, MCP server management, AI provider switcher /
per-phase model assignment. These belong to a "configurator" product, not mallard.
