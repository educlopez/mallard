---
name: mr-flow
description: >
  Branch, commit, push, open and merge a GitLab merge request for a project,
  after detecting which branch (main or master) that project actually uses as
  its base. Use when the user says "haz MR", "crea la MR", "sube esto", "mergea
  la MR", "abre una merge request", "empieza una rama para la tarea X", "haz
  commit y push", or finishes a change and wants it reviewed/merged. Works with
  `glab` on GitLab.com. Runs only the steps the user asked for.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# mr-flow

One consistent path from "I changed some files" to "merged into the base branch", without mixing up `main` and `master`.

## Rules

- Do only the steps the user asked for. "Haz commit" is not "haz MR"; "haz MR" is not "mergea". Each of commit, push, MR creation and merge needs its own request (or one message that clearly covers them).
- Never commit on the base branch or any protected branch. Never force-push. Never merge a failing or still-running pipeline, and never merge unless asked.
- Follow the CLAUDE.md of the repo and of the user for commit message language, attribution lines (e.g. co-author trailers) and PR/MR footers. Do not add any by default.
- Commits use Conventional Commits (`feat:`, `fix:`, `refactor:`, `chore:`, `docs:`, `style:`), with the scope when it helps.
- Stage explicit paths, not `git add -A`, so untracked scratch files do not end up in the commit.

## 1. Detect the base branch

Projects differ: some use `main`, others `master`. Never assume.

```bash
git fetch -q origin
git remote set-head origin -a >/dev/null 2>&1   # refresh origin/HEAD from the remote
git symbolic-ref --short refs/remotes/origin/HEAD   # e.g. origin/master
git branch -r | grep -E 'origin/(main|master)$'
```

- Exactly one of `origin/main` / `origin/master` exists and `origin/HEAD` points to it: that is the base. State it to the user in one line before continuing.
- Both exist, or `origin/HEAD` disagrees with what the task says (for example the task mentions a deploy branch), or the detection fails: stop and ask which one. A wrong base opens a huge, unrelated MR.
- Branch protection can tell you more: `glab api "projects/:fullpath/protected_branches" | head -c 600`. Use it only when in doubt.

## 2. Start the branch

```bash
git checkout -q -b <branch> origin/<base>
```

Branch name: follow the task/ticket convention of the project (`feat/<id>-short-slug`, `fix/<id>-short-slug`, or the id alone for tickets such as `42565`). If the working tree already has unrelated changes, say so before branching.

## 3. Commit and push

```bash
git add <explicit paths>
git commit -q -m "<type>(<scope>): <summary in English>

<why, in one or two sentences>"
git push -u origin <branch> 2>&1 | tail -3
```

Check `git status -sb` and `git diff --cached --stat` before committing and report what is in the commit.

## 4. Create the MR

```bash
glab mr create --source-branch <branch> --target-branch <base> \
  --title "<same Conventional Commit title>" \
  --description "<what changed, why, how it was tested>" --yes
```

Target branch is the detected base. Report the MR URL.

## 5. Wait for the pipeline (only if asked to merge or to check)

```bash
glab mr view <iid> -F json | python3 -c 'import json,sys; d=json.load(sys.stdin); p=d.get("head_pipeline") or {}; print(p.get("status","none"), d.get("detailed_merge_status"))'
```

Poll every ~15 s, at most a few minutes. Statuses: `success` → continue; `running`/`pending`/`created` → keep waiting; `failed`/`canceled` → stop and report the failing job (`glab ci status`, `glab ci trace <job>`); `none` → the repo has no pipeline, proceed only if the user confirms. If `detailed_merge_status` is not `mergeable`, report why (conflicts, discussions, approvals) instead of forcing it.

## 6. Merge (only when asked)

```bash
glab mr merge <iid> --yes --auto-merge=false
git fetch -q origin && git checkout -q <base> && git pull -q --ff-only origin <base>
git log --oneline -3 && git status -sb | head -1
```

Add `--squash` or `--remove-source-branch` only if the user or the repo convention says so. After merging, tell the user the base branch now contains the change and what remains on their side (pull on the server, reload, verification).

## Gotchas

- `glab` prints `[glab]` noise on stderr; pipe through `grep -v '^\[glab\]'` when parsing.
- On a repo where the default branch on GitLab is not the branch that gets deployed, ask which one to target. Do not infer it from `origin/HEAD` alone.
- A `fatal: ... --ff-only` on pull after merge means the base moved or the local base had commits; report it, do not reset.
