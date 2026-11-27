# Task workflow

How the skills chain on one tracker task. Each step is a skill the user starts, and each skill ends with a `Next:` line that points to the following one.

Works on any stack mallard covers: PrestaShop, Laravel, plain PHP or JS. Stack-specific steps live inside each skill.

```
/task <id>              brief + branch + type          (task-context)
  |
  |-- brief thin or unclear?      grill                 ask until no decision is open
  |-- client-only questions?      client-questionnaire  (planned) draft, never sent
  |-- fact to look up?            research-brief        (planned) cited .md
  v
route by type           bug -> diagnose-bug
                        translation / image / layout / security -> the skill or agent /task names
  v
verify                  verify-ps-project (PrestaShop), the project's own tests (Laravel)
                        blast-radius when shared code changed
  v
task-review             diff vs brief + AGENTS.md, read-only
  v
mr-flow                 branch, commit, push, MR (only when the user asks)
  v
/task done <id>         log time
  v
reflect                 turn the session's repeated mistakes into skill edits
```

## Rules for the chain

- One owner per step. A skill that needs another loads it by name ("Follow `lando-setup`") and does not copy its text.
- Every external action waits for the user: commit, push, MR, tracker writes, client messages.
- A new skill that fits the chain gets a `Next:` line and a row here. Skills outside the chain (`ps-*` tools, bootstrap skills) are not listed.
