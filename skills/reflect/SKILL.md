---
name: reflect
description: >
  Mine the current session for durable learnings and route each to a concrete
  edit of an existing skill, command or agent, waiting for approval before
  applying anything. Use when the user says "reflect", "reflexiona", "qué hemos
  aprendido", "mejora las skills con lo de hoy", "esto debería ir a una skill", or
  closes a session that exposed a repeated mistake or a missing step.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# reflect

Adapted from pstack's `reflect`. One pass, no parallel reviewers, no model switching. It does not save preferences or session state (that is engram) and does not touch the wiki (that is `second-brain`).

## 1. Skip when

The session was trivial, off-topic, or the skill that applied was followed correctly. One-offs are not learnings.

## 2. Collect

Use what is already in the conversation. If more is needed, read only this session's transcript under `~/.claude/projects/<slug>/` where `<slug>` is the working directory with each `/` turned into `-`. Never read other projects' transcripts.

## 3. Find candidates

A candidate needs evidence from the session:

- The user corrected the same thing twice.
- A skill step was wrong, stale or missing and you worked around it.
- A skill did not fire when it should have (description problem).
- A command or path failed in a way the skill did not warn about.

## 4. Route each candidate

| Kind | Route |
|---|---|
| A rule a script could enforce | Backlog item: lint or check, see `docs/skill-rules.md` |
| Wrong or stale fact in a skill | Trivial edit of that skill |
| New step, pattern or section | Substantive edit, use `skill-author` |
| Skill did not trigger | Tune its `description` |
| Nothing fits and it will recur | New skill via `skill-author` |

Prefer the structural route over adding more prose.

## 5. Approval

Show a table: candidate, evidence, destination file, proposed change. Reject weak ones with a one-line reason. Wait for the user's OK, then apply only the approved rows. Never edit installed copies under `~/.claude/skills/` (they are symlinks into this repo; edit the repo). Never commit, push or open a PR.

## 6. Report

Edits applied (file, one line), backlog items, dropped candidates with reason.
