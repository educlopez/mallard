# Astro — ci-templates wiring

No PHP, so no static-analysis gap to fill like PrestaShop/Laravel —
`lint-node` + `typecheck-node` already cover everything. `test-node` (added
v1.2.0) covers the test gate — see below for when to actually wire it.

## Components to include

```yaml
workflow:
  rules:
    - if: '$CI_PIPELINE_SOURCE == "merge_request_event"'
    - if: '$CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH'

include:
  - component: gitlab.com/educlopezcinetic/ci-templates/semantic-mr-title@v1.2.0
  - component: gitlab.com/educlopezcinetic/ci-templates/agent-guardrail@v1.2.0
  - component: gitlab.com/educlopezcinetic/ci-templates/secret-scan@v1.2.0
  - component: gitlab.com/educlopezcinetic/ci-templates/dependency-audit-node@v1.2.0
    inputs:
      package-manager: pnpm   # pnpm | npm | yarn — check the lockfile
  - component: gitlab.com/educlopezcinetic/ci-templates/lint-node@v1.2.0
    inputs:
      package-manager: pnpm
      lint-command: 'pnpm run lint'
  - component: gitlab.com/educlopezcinetic/ci-templates/typecheck-node@v1.2.0
    inputs:
      typecheck-command: 'pnpm run types'   # Astro convention: "types" runs `astro check`, not tsc directly
  # Only if the project has real app logic worth testing — see below:
  - component: gitlab.com/educlopezcinetic/ci-templates/test-node@v1.2.0
    inputs:
      package-manager: pnpm
      test-command: 'pnpm run test'   # vitest run, typically
```

Validated against `educlopezcinetic/astro-demo` (deliberate lint/typecheck/
agent-guardrail/secret-scan/semantic-mr-title/**test** violations, one MR
each round, all fired correctly).

## Test gate — depends on content-only vs real app logic

External research (withastro/astro core, withastro/starlight, a large
community Astro/Starlight docs site) shows this only holds as "no test
needed" for **pure-content sites** (docs, marketing, blog — no forms, no
auth, no real business logic): those keep testing to lint+typecheck+build.
The moment a project has real app logic, every comparable real-world Astro
project studied has Vitest (unit) and often Playwright (e2e) as a
**required** merge gate, not optional — don't treat "no tests" as a
permanent default once a client project crosses that line (a form with
validation, an auth flow, a dashboard with state).

Don't wire `test-node` speculatively on a genuinely content-only project
with nothing to test — an empty test job is worse than no job, it just adds
CI minutes for nothing. Do wire it the moment the project has real logic
worth protecting.

**Gotcha found wiring this into `astro-demo`**: if `vitest` isn't already a
real `devDependency` and you can't run `pnpm add -D vitest` locally to
regenerate the lockfile properly (no SSH access to whichever account hosts
the repo), do NOT reach for `pnpm dlx vitest run` as a lockfile-free
workaround — it breaks `typecheck` (`astro check` / `tsc --noEmit` can't
resolve the `vitest` module for typed test files, e.g.
`Cannot find module 'vitest' or its corresponding type declarations`),
turning one fix into a new regression. Confirmed by cloning over HTTPS with
a personal access token (`git clone https://oauth2:<token>@gitlab.com/...`)
when SSH isn't configured for the account, running `pnpm add -D vitest` for
real, and committing the updated `package.json` + `pnpm-lock.yaml`.
