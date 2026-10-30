---
name: blast-radius
description: >
  Before shipping a change, find what it could break outside the diff and prove
  the one fact it is safe because of by running real code. Use when the user
  says "blast radius", "qué puede romper esto", "revisa el impacto", "cambio de
  riesgo", or before touching shared code: a PrestaShop override or hook, a theme
  partial used on many pages, a module config, or mallard internals such as the
  skill registry, backups or agent adapters.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# blast-radius

Adapted from pstack's `blast-radius`. Listing callers is not the job, grep does that. The job is the breakage grep does not show. Read-only: it never edits the change. It does not launch parallel model reviews.

## 1. Read the change

The diff, the symbols added, changed or deleted, and what now behaves differently, including what the diff does not spell out.

## 2. Find the one fact it is safe because of

Most risky-looking changes are safe because of a single fact ("this hook only runs in the BO", "this symlink target is recreated on update"). Find it. Spend the time here, not on a long list of maybes.

## 3. Look where grep stops

- PrestaShop: overrides and the class index cache, hooks registered by other modules, Smarty partials included from other templates, `theme.yml` and `config.xml` values read elsewhere, multistore and language differences, the DB column or config key another module reads.
- Theme CSS: selectors that now match elsewhere, the compiled output versus its source.
- mallard (Go): other agents' adapters, symlink and backup paths, Windows path handling, state read from `~/.mallard/state.json`.

## 4. Rate how sure you are

For each fact the change depends on, get as far down this list as is cheap and say where you stopped:

1. You said so (worthless alone).
2. You pointed at a real `file:line`.
3. You walked the failure step by step and it does not reach.
4. You ran a script or test against the real code.
5. You reproduced it in the running app (`verify-ps-project`).

## 5. Report

- **What it does.** Including the non-obvious part.
- **The one safety fact.** Proven (state the step) or `unproven`.
- **Risks.** How it breaks, `file:line`, likelihood, cost, how to check.
- **Cleared.** What you checked and why it is fine.
- **Before merge.** The cheapest check that catches the real bug.

Never invent a caller or an API. A search that finds nothing is still an answer.
