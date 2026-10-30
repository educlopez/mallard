# Skill rules and what enforces them

A rule nothing enforces will be broken again. Each row pairs a rule with its check; when a rule is only prose, say so. When you catch a repeated mistake, fix it at the highest level that works (structure, then types or lint, then tests, then docs), add a row, and drop rows whose mistake can no longer happen. Idea adapted from pstack's `correct` and `encode-lessons-in-structure`.

| Rule | Enforced by |
|---|---|
| Every skill has `SKILL.md` with frontmatter | `scripts/lint-skills.sh` (CI job `skills`) |
| `name` equals the directory (skills, agents) or file name | `scripts/lint-skills.sh` |
| `description` is present and not empty | `scripts/lint-skills.sh` |
| `version` is semver `X.Y.Z` | `scripts/lint-skills.sh` |
| `references/...` paths in a skill exist | `scripts/lint-skills.sh` |
| Every skill, command and agent appears in `README.md` | `scripts/lint-skills.sh` |
| PR stays within 400 changed lines | `.github/workflows/pr-check.yml` |
| Bump `version` when behavior changes | Prose only. Candidate: compare against the last tag in CI. |
| Trigger phrases in the `description` | Prose only |

## Review checklist for operations that can be interrupted

For `update`, `restore`, backup and install steps (adapted from pstack's idempotency principle):

- Re-running after a crash or partial run converges to the same end state.
- A half-created symlink or half-copied directory is detected and repaired, not skipped.
- Nothing is deleted before its replacement exists.
- Backups are written before links change, and `--restore` works from any interrupted point.
