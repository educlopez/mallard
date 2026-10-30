---
name: no-comments
description: >
  Review the current diff for comments that apologize for the code (workarounds,
  "hack", "TODO fix later", "do not touch") and fix the cause instead of
  explaining it. Use when the user says "limpia comentarios", "no-comments",
  "quita los hacks comentados", "hay muchos comentarios raros", or after a
  change that left workaround comments behind.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# no-comments

Adapted from pstack's `no-comments`. A comment that explains a workaround is a flag: the code surprised its author. Fix the surprise, then drop the comment. It does not strip useful comments, edit code outside the diff, or run a separate reviewer agent.

## Scope

The caller's files or the diff against the base branch (`main` or `master`), including the working tree.

## Steps

1. List comments in scope and sort each:
   - **Describes what the code does:** delete only if the code is already clear.
   - **Explains a workaround or a surprise:** a flag. Find the root cause.
   - **States a constraint** ("do not remove", "talk to X first", vendor quirk, PrestaShop core behavior): keep, unless a type, test or lint can enforce it.
   - **Lint or type suppression:** check it is still needed; correctness and safety suppressions stay.
2. For each flag, make the smallest root-cause fix in scope: delete the dead path, use the real API, drop the parameter. If the root cause is out of scope, make the smallest in-scope fix and report the rest as open.
3. For each kept constraint, offer the cheapest in-scope enforcement (test, lint, check) and wait for the user's OK before adding it. If approved, encode, then delete the comment.
4. Never delete a comment you were not asked to touch in files outside the diff. Never add a symptom guard.

## Report

Comments deleted, flags fixed (file, one line), constraints kept with reason, enforcement offers, open items.
