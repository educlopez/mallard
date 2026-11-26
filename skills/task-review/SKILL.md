---
name: task-review
description: >
  Review the branch diff against the task brief and the repo's
  AGENTS.md before opening the MR. Use when the user says "revisa mi diff",
  "task review", "revisa antes de la MR", "¿cumple la tarea?", or after
  finishing the work on a `/task` branch.
version: "0.1.1"
metadata:
  author: Eduardo Calvo
---

# task-review

Adapted from Matt Pocock's `code-review`: two axes, reported separately. Its spec axis is the task brief, not an issue tracker. Read-only: it never edits code, comments on the MR or touches the tracker.

## 1. Fix the diff

Base = the branch `mr-flow` detects (`main` or `master`). Run `git diff <base>...HEAD`. Done when the diff is not empty; if it is, stop and say so.

## 2. Find the two references

- Spec: engram `task/<localid>/brief`, where `<localid>` is the digits at the start of the branch name. No brief: ask the user to paste the task, or skip the spec axis and say you skipped it.
- Standards: `AGENTS.md`, `CLAUDE.md`, `CONTRIBUTING.md`, and the lint or static-analysis config the repo has (Pint, PHPStan, ESLint, Prettier). Run the check commands that `AGENTS.md` lists and report failures without fixing them.

## 3. Review the two axes

Run them as parallel subagents when available, otherwise one after the other. Do not merge their reports.

- Spec: for each thing the brief asks, mark `done`, `partial` or `missing` with a `file:line`. List diff changes the brief did not ask for.
- Standards: violations of the repo's documented rules first, then judgment smells: duplicated logic, dead code, long functions, a comment that explains a workaround (see `no-comments`).

## 4. Check the blast radius

If the diff touches shared code, run `blast-radius`: a PrestaShop override or hook, a theme partial used on many pages, a Laravel migration, middleware, policy, config or queued job.

## 5. Report

One line per finding: `file:line: problem. fix.` Then one verdict line: `ready for mr-flow` or `N blockers`. Done when both axes have a section, even if empty.

Next: `mr-flow` once the user agrees, then `/task done <localid>`.

It does not commit, push, open the MR or log time.
