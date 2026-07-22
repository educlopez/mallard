---
name: new-project
description: Scaffold the day-1 setup for a new PrestaShop client project — child theme (panda or elementflow), CSS build, and CLAUDE.md — by composing existing skills. Does not create the repo, server, or install PrestaShop core.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# new-project — new client day-1 scaffold

Run inside an already-cloned PrestaShop repo. Does the repetitive setup that's still manual after `lando-setup`: **child theme → CSS build → CLAUDE.md**.

## Arguments

```
--track <panda|elementflow>   (required) theme track
--name <child-theme-name>      default: <repo-slug>-child
--no-lando                     skip lando-setup handoff
--no-css                       skip ps-css-build
--no-doc                       skip ps-project-doc (CLAUDE.md)
```

## Instructions

Load and follow the **`ps-new-project`** skill (`~/.claude/skills/ps-new-project/SKILL.md`):

1. Verify cwd is a PS repo; detect Panda vs ElementFlow installed; warn on track mismatch.
2. Scaffold the child theme via `scripts/scaffold-theme.sh <track> <name> "<display>" .`.
3. Wire the CSS build via the `ps-css-build` skill (both tracks use the same lightningcss pipeline) unless `--no-css`.
4. Generate CLAUDE.md via `ps-project-doc` non-interactively (gitignored, never committed) unless `--no-doc`.
5. Optionally hand off to `lando-setup` unless `--no-lando`.
6. Report what was created/skipped + remaining manual steps (activate theme in BO; repo/server stay manual).

Never create the repo/server, never install PS core, never deploy, never clobber an existing child theme.
