# mallard — Backlog

gentle-ai-inspired borrows worth doing.

## Done

- **Backup dedup** — SHA-256 content-addressed dedup in `internal/backup`; identical content stored once, hard-linked across sessions. Landed in commit 6ab84f9.
- **More agents** — extended beyond claude/codex/opencode/generic; additional adapters (gemini, cursor, windsurf, and others) added to the registry. Landed in commit 6ab84f9.
- **Persisted state (state.json)** — `~/.mallard/state.json` for user selections; graceful on missing/corrupt. Landed in commit 6ab84f9.

## Open

- **`ci-templates` agent-guardrail: exclude the CI file by default.** The component flags the `.gitlab-ci.yml` that lists its own `forbidden-patterns`, and any doc that quotes them, so every project that sets the input has to add `.gitlab-ci.yml` to `exclude-paths` by hand (seen twice in one week). Fix in `educlopezcinetic/ci-templates`, then drop the workaround from `gitlab-ci-quality-gates`. Owner: the `ci-templates` repo, not mallard.
