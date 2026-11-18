---
name: diagnose-bug
description: >
  Diagnose a hard bug or regression with a repeatable feedback loop before
  touching the fix. Use when the user says "diagnostica", "depura esto", "por
  qué falla", "error 500", "no funciona", or when `/task` classifies a task as
  `bug`. Works on PrestaShop, Laravel and other PHP or JS projects.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# diagnose-bug

Adapted from Matt Pocock's `diagnosing-bugs`. It finds the cause and the fix. It does not commit, deploy or write to Intervals.

## 1. Pick the environment

Work locally: Lando for PrestaShop (`lando-check`, `ps-console`), Sail, Herd or `php artisan serve` for Laravel. PRE and PROD are read-only: look at logs and config, and ask before running anything there. Never add instrumentation to PRE or PROD without the user's OK. Replace secrets and customer data in anything you print with `<REDACTED>`.

## 2. Build the feedback loop

One command that is deterministic, fast and turns red for this bug. Pick the cheapest that works:

| Stack | Loop |
|---|---|
| Laravel | failing Pest/PHPUnit test, `php artisan tinker` script, `php artisan <command>`, `curl` on the local route, `php artisan pail` for the log |
| PrestaShop | script run through `ps-console`, `curl` on the local URL, `var/logs`, `PrestaShopLogger` |
| Front end | `verify-ps-project` or the browser tools |

Done when the loop fails for the reported reason three times in a row. If you cannot build one, stop and tell the user what is missing: data, access, steps.

## 3. Reproduce and shrink

Cut the input until removing anything more makes the bug disappear. Done when you can state the minimal trigger in one sentence.

## 4. Write 3 to 5 hypotheses

Falsifiable, ranked by likelihood, each with the check that kills it. Show the list to the user before testing the first one.

## 5. Test one hypothesis at a time

One change, one loop run. Tag every debug line `[DEBUG-<4 chars>]` (`error_log`, `Log::debug`, `dump`). Done when exactly one hypothesis survives.

## 6. Fix with a regression test

If the project has a runner (Pest, PHPUnit) and a place to hook the test, write the test first and watch it fail, then fix. If not, say that as a finding and do not invent a test harness. Done when the loop is green and the neighbouring tests still pass.

## 7. Clean up

`grep -rn "DEBUG-<tag>" .` returns nothing and scratch files are gone. Give the root cause in one line for the MR body and save it to engram when it is non-obvious.

Next: `blast-radius` if the fix touches shared code, then `task-review`.

It does not commit, push or touch PRE and PROD without the user's OK.
