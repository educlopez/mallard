---
name: gitlab-ci-quality-gates
description: >
  Wires Cinetic's reusable GitLab CI/CD Components
  (gitlab.com/educlopezcinetic/ci-templates — lint, typecheck, test, semantic
  MR titles, agent guardrails, secret scanning, dependency audits) into a
  project's .gitlab-ci.yml, adds the stack's static-analysis job (PHPStan for
  PrestaShop, Larastan for Laravel — not yet a ci-templates component, wired
  repo-local), and configures branch protection + merge checks to match how
  the project is actually worked. Use when the user asks to "set up CI
  quality gates", "montar CI en este proyecto", "añadir ci-templates", wants
  PHPStan/Larastan added, or wants MR-based branch protection on a GitLab
  project. Covers PrestaShop, Laravel, Astro, and Next.js — add a new
  reference file the same shape when another stack needs this, don't build
  ahead of need. Distinct from
  gitlab-security-setup (scheduled Trivy/pnpm supply-chain email report) and
  ps-security-audit — this is about MR-time quality/security GATES, not
  scheduled reports. Do NOT use for GitHub-hosted projects.
version: "0.4.0"
metadata:
  author: Eduardo Calvo
---

# GitLab CI Quality Gates (ci-templates)

Wires the already-built, versioned components from
[`educlopezcinetic/ci-templates`](https://gitlab.com/educlopezcinetic/ci-templates)
into a project, plus the one piece that repo doesn't cover yet (PHP static
analysis), and sets branch protection to match how the project is really
worked. Validated end-to-end against `educlopezcinetic/prestashop-demo` and
`laravel-demo` (real deliberate-failure smoke tests for every gate, not just
"pipeline ran").

The stack-specific include blocks and static-analysis job live in
`references/` — this file is the shared decision tree and gotchas, identical
regardless of stack. Load the matching reference file for Step 1 and Step 2:

- `references/prestashop.md`
- `references/laravel.md`
- `references/astro.md`
- `references/nextjs.md`
- (No other stack references exist yet — add one following the same shape
  when a new stack actually needs this, don't build ahead of need.)

## Step 0 — ask, don't assume: blocking or informational?

This is the one decision that changes everything else. Ask the user:

> "¿Este proyecto lo trabajás solo con hábito de PR, o el equipo mergea ramas
> directo a `test` sin MR?"

- **Solo / personal workflow, real PR habit** → gates are **blocking**.
  Enable `only_allow_merge_if_pipeline_succeeds`, protect `test`/`main` with
  push access "No one" (merge via MR only). This is the Eduardo-personal-project
  case — see Step 3.
- **Team client project** (coworkers merge branches straight to `test`, no
  MR habit yet) → gates are **informational only**. Do **NOT** enable
  `only_allow_merge_if_pipeline_succeeds` and do **NOT** restrict push access
  on `test`/`main` — that would silently block a coworker's normal workflow
  the moment they push, with no warning, for a policy they never agreed to.
  The pipeline still shows red/green on every push as a signal; it just never
  blocks anything. Skip Step 3 entirely for this case.

Never flip a real client project to blocking without the user explicitly
confirming the whole team knows and is ready for it.

## Step 1 — wire the ci-templates components

Open the reference file matching the project's stack (`references/prestashop.md`
or `references/laravel.md`) for the exact `include:` block and `workflow:`
rules to use.

Always verify the current tag before wiring — do not assume the version in
the reference file is still latest:

```bash
glab api "projects/educlopezcinetic%2Fci-templates/repository/tags"
```

Check the `ci-templates` repo's own README for the current, authoritative
input list per component before wiring — it's the source of truth; this
skill and its references just point you at it and add the missing
static-analysis piece.

## Step 2 — static analysis (repo-local, not a component yet)

`ci-templates` doesn't have a PHPStan/Larastan component yet — add it as a
repo-local job per the stack's reference file. **Critical gotcha**: every
`ci-templates` component job declares `needs: []` on purpose (see the repo's
own README) so one gate failing doesn't silently skip the others via classic
stage-based blocking. Any repo-local job you add alongside them — like this
one — MUST also declare `needs: []`, or it will get marked `skipped` (not
run, not failed, just invisible) the first time an earlier-stage gate fails.
Confirmed as a real bug in `prestashop-demo` before this skill existed —
don't repeat it.

## Step 3 — branch protection + merge checks (blocking case ONLY)

Only do this after Step 0 confirmed the project is solo/personal with a real
PR habit. Skip entirely for team client projects.

```bash
enc=$(python3 -c "import urllib.parse,sys; print(urllib.parse.quote(sys.argv[1], safe=''))" "<namespace>/<project>")

# Create + protect a `test` branch (pre-prod), merge-only:
glab api -X POST "projects/$enc/repository/branches" --field "branch=test" --field "ref=main"
glab api -X POST "projects/$enc/protected_branches" --field "name=test" --field "push_access_level=0" --field "merge_access_level=40"

# Block merge on red pipeline:
glab api -X PUT "projects/$enc" --field "only_allow_merge_if_pipeline_succeeds=true"
```

`push_access_level=0` ("No one") is deliberate — it forces even the repo
owner through a Merge Request, which is the whole point if the goal is
building the habit. If that's too strict for how the user actually wants to
work, use `40` (Maintainer) instead and say so explicitly — don't silently
soften it.

If the project follows Cinetic's Intervals-task branching model (task branch
off `main`/`master`, merged to `test` first for pre-prod validation, then
**the same task branch** — never `test` itself — merged to `main` once the
client approves), the `workflow:` rules already cover both merge targets:
`merge_request_event` fires regardless of which branch the MR targets, so
gates run identically whether the MR goes into `test` or `main`. No extra
config needed for this branching model specifically.

## Gotchas (stack-agnostic)

- **`needs: []` on every repo-local job** you add next to ci-templates
  components — omitting it causes silent `skipped` (not `failed`) status the
  first time an earlier gate fails. See Step 2.
- **Pin the exact `ci-templates` tag** (`@vX.Y.Z`), never `@main`. If you're
  the one maintaining `ci-templates` itself, also protect its tags
  (Settings → Repository → Protected tags, pattern `v*`, Maintainer-only) —
  GitLab tags are mutable by default, and an unprotected tag is a single
  point of failure across every project that includes it.
- **No zizmor-equivalent exists for GitLab CI.** The closest tool is
  `gitlab-ci-verify` (ShellCheck over `script:` blocks) — small, low-adoption.
  The real defenses are GitLab-native: protected CI/CD variables are never
  exposed to non-protected-branch or fork-originated pipelines by default;
  the actual gotcha is a maintainer manually running a fork MR's pipeline
  *in the parent project* — review the `.gitlab-ci.yml` diff first if you
  ever do that.
- **GitLab Free tier is enough for all of this.** `only_allow_merge_if_pipeline_succeeds`
  is a plain repository setting, not a Premium feature. What Free tier
  *doesn't* have is GitHub-style named required-status-checks or MR approval
  rules — the whole pipeline is the unit; if it's green, every gate in it
  passed.
- **CI minutes are limited** on GitLab.com's standard tier (shared runners).
  All the ci-templates gates are already designed around this: secret-scan
  and dependency-audit are MR-diff/lockfile-scoped (not full-repo), and
  `dependency-audit-*` also runs on scheduled pipelines rather than every
  push where it makes sense. Don't add a full test/build matrix across
  multiple PHP/framework versions on a client project's real CI — that's an
  OSS-scale practice (see PrestaShop core/Laravel core's own CI), not
  appropriate for a small team's per-minute budget. One job pinned to the
  actual production PHP/framework version is enough.
- **Dependency-audit findings on a fresh/clean install are normal**, not a
  setup mistake — see the stack reference file for the specific "safe to
  hand-patch vs. wait for upstream" version-range logic before touching
  anything a scan finds.

## Known gaps — researched, not built yet

External research across large OSS PrestaShop/Laravel/Astro/Next.js
projects (astro, starlight, next.js, cal.com, formbricks, laravel/framework,
spatie/laravel-permission, filamentphp/filament) surfaced practices this
skill doesn't cover yet. Don't build these speculatively — add them when a
real project's needs justify the CI-minute cost, same reasoning as
everything else here:

- **Dependabot/Renovate-equivalent** — near-universal in every repo studied,
  currently absent here. GitLab's own Dependency Scanning / a self-hosted
  Renovate bot both fit; hasn't been evaluated yet for this team's scale.
- **Severity-threshold hard-fail on dependency audits** — `cal.com` runs
  `audit --severity critical` as an actual blocking gate rather than a
  passive report. `dependency-audit-node`/`dependency-audit-php` here
  already support an `audit-level` input (see references) — using it as a
  hard gate rather than leaving it at a lenient default is a cheap upgrade
  worth considering per-project.
- **Free-tier SAST (CodeQL/Semgrep-equivalent)** — used at scale by
  astro/formbricks/cloudflare-docs for any project touching auth/user data.
  No direct GitLab-native equivalent evaluated yet — flagged as a candidate,
  not scoped.
