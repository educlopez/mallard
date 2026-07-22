---
name: task-context
description: >
  Turns an Intervals task into ready-to-work context. Fetches the task and its
  notes, decodes HTML entities, builds a compact brief, classifies the task type,
  creates the git branch, and routes to the right skill/agent. Handles time
  logging on finish with 15-minute banding. Intervals is the single source of
  truth; Google Workspace (via the `gws` CLI) is an opt-in enricher only —
  never searched automatically.
  Use this skill for the /task command, or whenever the user references an
  Intervals task by its numeric ID and wants to start or close work on it.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# task-context — Intervals task → ready-to-work

Single source of truth = **Intervals**. Google Workspace = **opt-in enricher only**. Never auto-search Gmail/Drive (wastes tokens on context that may not exist).

All Intervals calls use the `intervals` MCP tools. Gmail/Drive use the `gws` CLI via the helper script.

---

## Flow A — start a task: `/task <localid> [flags]`

### 1. Fetch

- `get_task(<localid>)` → core fields.
- `get_task_notes(taskId=<localid>)` → comments (real detail usually lives here).
- **Decode HTML entities** in `title`, `summary`, and note bodies — both named (`&oacute;`→ó, `&amp;`→&) and **numeric** (`&#10;`→newline, `&#39;`→'). Notes commonly use `&#10;&#10;` for paragraph breaks.

### 2. Assess completeness (label only — no search)

If `summary` is empty AND notes are thin (< ~200 chars / no actionable detail), mark the brief:

```
⚠ thin — no actionable detail in Intervals. Run  /task <localid> --enrich  to pull Gmail/Drive context.
```

Do **not** search anything here. Enrich is explicit only.

### 3. Enrich — ONLY if `--enrich` was passed

Never triggered automatically. When invoked:

```bash
~/.claude/skills/task-context/scripts/gmail_search.sh "<client>" "<pm-owner>" "<keywords>"
```

- `<client>` = task `client` field; `<pm-owner>` = task `owners` field (the PM); `<keywords>` = 2–4 salient words from the title.
- The script prints compact JSON (from / date / subject / snippet). Fold the relevant lines into the brief with a one-line source note.
- Optional Drive lookup for assets:
  ```bash
  gws drive files list --params '{"q":"name contains '\''<client>'\''","pageSize":10,"fields":"files(id,name,mimeType,webViewLink)"}' --format json
  ```
- Only the parsed summary enters context — never paste raw email bodies.

### 4. Classify type

From `module` + title/notes keywords, assign one type and its route:

| Type | Signals | Route |
|------|---------|-------|
| `translation` | i18n, "traducir", English strings | `ps-translate` skill (near-auto) |
| `image` | "imágenes", white zones, thumbnails, regen | `ps-image-regen` skill (near-auto) |
| `layout` | Disseny/design module, CSS, header/footer, product page | `layout-builder` agent (semi-auto) |
| `bug` | error, "no funciona", stacktrace | human-led, agent assists |
| `security` | CVE, module advisory, scan | `ps-security-audit` skill |
| `content` | text/CMS edits | semi-auto |
| `unknown` | fallthrough | present brief, ask |

If `--type <type>` was passed, use it instead.

### 5. Brief — print to chat AND save to engram

Print this, then `mem_save` it with `topic_key: task/<localid>/brief`, `project: cinetic`, `capture_prompt: false`. No repo file.

```
TASK <localid> · [<status> · P<priority> · <severity>]
<decoded title>
Client: <client>   Module: <module>   PM: <owners>
Time: <actual>h logged / est <estimate>   Due: <datedue or —>

WHAT'S ASKED
<decoded summary + distilled notes  |  ⚠ thin — ... (see step 2)>

CONTEXT (enriched)          ← only if --enrich ran
<gmail/drive snippets + links>

AMBIGUITIES / TO CHASE
- <missing info to ask the PM>

TYPE: <type>  →  route: <skill/agent>
BRANCH: <localid> (<created|switched|skipped>)
NEXT: <proposed first steps>
```

### 6. Git branch (unless `--no-branch`)

- Branch name = `<localid>`. If it exists or is stale → `<localid>-<slug>` (slug = kebab of a few title words).
- Create/switch in the **current repo** (cwd). No cross-repo logic.
- Never delete remote branches (team policy). Local hygiene only: `git fetch -p`.

### 7. Route

Name the matching skill/agent (step 4) and offer to hand off. For deterministic types (`translation`, `image`, `security`) you may proceed with the skill directly; for `layout`/`bug` draft and let the user review.

---

## Flow B — finish a task: `/task done <localid>`

**Time only. No git. Never commit or push.**

### Time model — 15-minute banding, human-confirmed

Billing is booked in **15-minute bands, minimum 15 min** (0.25h increments). Never auto-log from elapsed wall-clock — the session may have been paused, so elapsed is unreliable.

1. Optionally show a naive elapsed figure **labeled**: `hint (unreliable, includes idle): ~Xh`.
2. Round the hint **up** to the nearest 0.25h and present as a suggestion.
3. Ask the user for the **real** time (accept the suggestion or a typed value).
4. Ask **billable? (y/n)** — every time, no default.
5. Resolve the worktype id (see below), then:
   `add_time_entry(taskId=<localid>, worktypeid=<id>, date=<today YYYY-MM-DD>, time=<confirmed decimal hours>, billable=<bool>)`.

Never write a time entry without the user's confirmed number.

### Worktype resolution — always "Varis"/"Varios"

The Intervals API **cannot list worktypes** for this user group (`intervals://worktypes` → 403), and worktype ids are **project-scoped**. Resolve dynamically:

1. Check the cache `~/.config/mallard/worktypes.json` for this task's `projectid`.
2. If absent, call `get_time_entries(taskId=<localid>, datebegin=..., dateend=...)` (a wide range, e.g. the last year) and find an entry whose `worktype` name is `Varis` or `Varios`; reuse its `worktypeid`. Cache it under the `projectid`.
3. Seed value: `256467` for project `NEVEN - Tiendas Aka` (projectid `1380816`).
4. If no prior entry exists in the project, ask the user for the worktype id once, then cache it.

---

## Hard rules

- Intervals is the only source of truth. Slack/Trello/Gmail are not integrated; Gmail/Drive are opt-in enrichers via `gws` only.
- Never auto-enrich. Never auto-log time. Never touch remote branches.
- English-only content (mallard is public). No company names in code/paths.
