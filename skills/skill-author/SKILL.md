---
name: skill-author
description: >
  Guided way to create or edit a skill, command or agent in the mallard repo:
  frontmatter, trigger phrases in Spanish, a "what it does not do" section,
  version bump rule, writing pass, lint. Use when the user says "crea una skill",
  "añade una skill a mallard", "mejora la descripción de la skill", "esta skill no
  se dispara", "nueva skill para X", or edits anything under skills/ or claude/.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# skill-author

Adapted from pstack's skill-creation flow plus `technical-writing` and `unslop`. The release steps live in `mallard-release` and `docs/adding-skills.md`; this skill covers the writing.

## 1. Decide the shape

- Repeatable multi-step workflow with judgment: skill in `skills/<name>/SKILL.md`.
- Short entry point the user types (`/name`): command in `claude/commands/<name>.md`.
- Isolated expert with its own tools: agent in `claude/agents/<name>.md`.
- A rule that must always hold: not a skill. Encode it as a lint (`scripts/lint-skills.sh`) or a CI check, and list it in `docs/skill-rules.md`.

If a skill with the same job exists, edit it instead of adding another.

## 2. Frontmatter

```yaml
---
name: <dir-name>
description: >
  What it does, then "Use when the user says "..."" with 4-6 real phrases in
  Spanish (and English if the user mixes), plus the contexts and file names that
  should fire it.
version: "0.1.0"
metadata:
  author: <name>
---
```

The `description` is the only text that decides whether the skill loads. Put triggers there, not in the body.

## 3. Body

1. One sentence: what it does and the closest skill it is not (`does not set X up, use Y`).
2. Numbered steps with real commands. Imperative voice, one action per step, one term per thing.
3. A table only when comparing outcomes (symptom, cause, action).
4. A "what it does not do" line. Include any action that needs the user's OK.
5. Long reference material goes to `references/*.md`, linked from the body.

## 4. Writing pass

- Cut filler: "it is important to", "make sure to", opening summaries, closing recaps.
- Short sentences, active voice, no invented jargon. Keep commands and paths exact.
- Spanish tics to drop: "cabe destacar", "es importante señalar", "en definitiva", stacked adverbs in "-mente".

## 5. Version and lint

- New skill: `0.1.0`. Wording or small fix: patch. New step or flow: minor. Behavior change that breaks users: major.
- Add the skill to the README table, then run `scripts/lint-skills.sh`. Fix every error before reporting done.
- Test the trigger: read the description as the model would and check that three real phrasings would fire it and one unrelated phrasing would not.

Never commit, tag or release without the user asking.
