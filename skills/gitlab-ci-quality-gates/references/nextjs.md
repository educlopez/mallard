# Next.js — ci-templates wiring

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
      audit-level: high       # Next.js pulls in a large, fast-moving tree — moderate/low gets noisy fast
  - component: gitlab.com/educlopezcinetic/ci-templates/lint-node@v1.2.0
    inputs:
      package-manager: pnpm
      lint-command: 'pnpm run lint'
  - component: gitlab.com/educlopezcinetic/ci-templates/typecheck-node@v1.2.0
    inputs:
      typecheck-command: 'pnpm run types'   # Cinetic convention: "types" script running `tsc --noEmit`
  # Only if the project has real app logic worth testing — see below:
  - component: gitlab.com/educlopezcinetic/ci-templates/test-node@v1.2.0
    inputs:
      package-manager: pnpm
      test-command: 'pnpm run test'   # vitest run, typically
```

Validated against `educlopezcinetic/nextjs-demo` (deliberate lint/typecheck/
agent-guardrail/secret-scan/semantic-mr-title/**test** violations, one MR
each round). One real finding: `eslint-config-next`'s default ruleset has no
`no-console` rule at all, and treats `@typescript-eslint/no-unused-vars` as
a **warning**, not an error — so `lint` exits 0 (job passes) on a
`console.log` + unused var that would fail on Astro's or PrestaShop's eslint
configs. `agent-guardrail`'s own regex-based check still caught the
`console.log` independently — this is exactly why it's a separate gate from
`lint`, not redundant with it: it doesn't depend on the project's own ESLint
severity choices. Don't assume `lint` alone blocks `console.log`/unused vars
on a Next.js project unless you explicitly check (and if needed, tighten)
its own `eslint.config.mjs` rule severities.

## Test gate — depends on content-only vs real app logic

External research tells a clearer story here than for Astro — **every**
real-world Next.js project studied at any real scale (vercel/next.js core,
cal.com's production scheduling SaaS, formbricks' production SaaS) treats
Jest/Vitest unit tests AND Playwright e2e as a required merge gate, not
optional. Unlike Astro, there's no clean "pure content" exception for
Next.js in practice — most Next.js projects exist specifically because they
need app logic (API routes, forms, auth, server actions). Treat "no
test-node wired" as a real gap to close once a Next.js client project has
anything worth protecting, not a permanent stance; genuinely static
marketing sites with zero interactivity are the narrow exception.

**Gotcha found wiring this into `nextjs-demo`**: same as Astro (see
`references/astro.md`) — if `vitest` isn't a real `devDependency` yet and
you can't run `pnpm add -D vitest` locally, do NOT use `pnpm dlx vitest run`
as a lockfile-free shortcut. It breaks `typecheck` (`tsc --noEmit` can't
resolve the `vitest` module in typed test files). Clone over HTTPS with a
personal access token if SSH isn't set up for the account, run
`pnpm add -D vitest` for real, commit the updated `package.json` +
`pnpm-lock.yaml`.

If the project uses Playwright/Cypress for e2e instead of/alongside unit
tests, keep that as a SEPARATE repo-local job with its own `rules:` (with
`needs: []`, same as any repo-local addition — see SKILL.md) — e2e suites
are slow and flaky enough that gating every MR on them isn't worth it on a
minute-limited GitLab.com tier; consider `allow_failure: true` or a
scheduled-only trigger instead, same reasoning as `dependency-audit-*`'s
scheduled-pipeline design.
