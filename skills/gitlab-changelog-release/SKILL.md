---
name: gitlab-changelog-release
description: >
  Manually generate a CHANGELOG.md and cut a GitLab release, using either
  git-cliff (Conventional Commits) or GitLab's native Changelog API (Git
  trailers) depending on which convention the repo already uses. Use when
  the user asks to "generar changelog", "cortar un release", "sacar un tag",
  "hacer release notes", or wants a CHANGELOG.md set up/updated for a project
  on GitLab.com. This is the ONLY supported flow — run it by hand every time
  you cut a release. Do NOT wire this into GitLab CI/CD tag-triggered jobs;
  see "Why manual, not CI" below.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# GitLab Changelog + Release (manual flow)

Two valid approaches — **pick based on what the repo's commit history
already uses**, don't mix them:

| Repo's commit convention | Use |
|---|---|
| Conventional Commits (`feat:`, `fix:`, `docs:`...) — most Cinetic repos | **Path A — git-cliff** |
| `Changelog: added/fixed/changed/...` Git trailers | **Path B — GitLab native Changelog API** |

**They are not interchangeable.** GitLab's native Changelog API
(`/repository/changelog`) and `glab changelog generate` ONLY read `Changelog:`
trailers — they do NOT parse `feat:`/`fix:` prefixes at all. Point either one
at a repo that only has Conventional Commits and it silently finds nothing
(tested live: returned `"No changes."` on a repo with 187 `feat:`/`fix:`
commits and zero trailers). Verified against
[GitLab's own Changelogs docs](https://docs.gitlab.com/user/project/changelogs/) —
no config exists to make the native API match commit-title prefixes instead
of trailers. Check `git log --oneline -20` before picking a path: if you see
`feat(...)`/`fix(...)` etc, use Path A; if commits end with a `Changelog:
added` line, use Path B.

## Why manual, not CI

GitLab.com tag-push pipeline creation was tested exhaustively on a real
Cinetic project and found completely non-functional — not a config bug on
our side. Proof: a job with **no `rules:` at all** (unconditional, should
always run) still produced **zero pipeline objects** on a plain push to the
default branch. That rules out YAML, `rules:`, job-token permissions, and
any known GitLab CI_JOB_TOKEN-push restriction (that one only applies to
pushes authenticated by a job token, not normal SSH/HTTPS pushes). Also
learned: `glab api projects/<id>/ci/lint --dry-run` is unreliable — it
reports "resulting pipeline would have been empty" even for an unconditional
job, so don't trust it as a signal either way. The remaining explanation
requires Owner-level access to Usage Quotas / Audit Events to confirm
(compute minutes, namespace restriction, etc) — until someone with that
access resolves it, do not build release automation on GitLab CI tag
triggers. Do the steps below by hand every time instead.

---

# Path A — git-cliff (Conventional Commits)

## Prerequisites

- Commits in the repo follow **Conventional Commits** (`feat:`, `fix:`,
  `docs:`, `chore:`, `refactor:`, etc). git-cliff groups by these prefixes —
  non-conventional commits get silently skipped from the changelog.
- `git-cliff` installed locally: `brew install git-cliff` (Rust binary, zero
  npm dependencies/postinstall scripts — deliberately NOT
  `conventional-changelog-cli`, which is deprecated with 54 transitive deps).
- `glab` CLI authenticated (`glab auth status`) — used to create the GitLab
  Release without touching the UI.

## Step 1 — `cliff.toml` config (once per project)

Create at repo root:

```toml
[changelog]
header = """
# Changelog\n
Todos los cambios notables de este proyecto se documentan aquí.\n
"""
body = """
{% if version %}\
    ## [{{ version | trim_start_matches(pat="v") }}]{% if timestamp %} - {{ timestamp | date(format="%Y-%m-%d") }}{% endif %}
{% else %}\
    ## [Unreleased]
{% endif %}\
{% for group, commits in commits | group_by(attribute="group") %}
    ### {{ group | striptags | trim | upper_first }}
    {% for commit in commits %}
        - {{ commit.message | upper_first }} ([{{ commit.id | truncate(length=7, end="") }}](../../commit/{{ commit.id }}))\
    {% endfor %}
{% endfor %}\n
"""
trim = true

[git]
conventional_commits = true
filter_unconventional = true
split_commits = false
commit_parsers = [
    { message = "^feat", group = "Features" },
    { message = "^fix", group = "Bug Fixes" },
    { message = "^docs", group = "Documentation" },
    { message = "^style", group = "Styling" },
    { message = "^refactor", group = "Refactor" },
    { message = "^perf", group = "Performance" },
    { message = "^test", group = "Testing" },
    { message = "^chore\\(deps\\)", group = "Dependencies" },
    { message = "^chore", group = "Miscellaneous" },
    { message = "^ci", group = "CI/CD" },
    { message = "^build", group = "Build" },
    { message = "^revert", group = "Reverts" },
]
filter_commits = true
tag_pattern = "v[0-9]*"
ignore_tags = ""
topo_order = false
sort_commits = "oldest"
```

Adjust `commit_parsers` groups/order to taste — this is the full default set.

## Step 2 — first CHANGELOG.md (once, covers all history)

```bash
git-cliff -c cliff.toml -o CHANGELOG.md
```

If the repo has **no tags yet**, every commit lands under `## [Unreleased]`.
That's expected and fine — the next tag you cut will split it out properly.
git-cliff prints a warning like `N commit(s) were skipped due to parse
error(s)` for non-conventional commits — that's normal, not a bug.

Commit it:

```bash
git add cliff.toml CHANGELOG.md
git commit -m "feat(changelog): add changelog generation with git-cliff"
git push origin <default-branch>
```

## Step 3 — cutting a release (repeat every release)

1. Decide the version (semver: `vMAJOR.MINOR.PATCH`). Bump MAJOR for breaking
   changes, MINOR for features, PATCH for fixes only — same rule regardless
   of whether commits used exact Conventional Commit types.

2. Regenerate the full changelog and extract just this version's notes:

   ```bash
   git-cliff -c cliff.toml -o CHANGELOG.md
   git-cliff -c cliff.toml --unreleased --tag vX.Y.Z --strip header -o release_notes.md
   ```

   `--unreleased --tag vX.Y.Z` tells git-cliff "treat everything since the
   last tag as if it were tagged vX.Y.Z" — this works BEFORE the tag exists,
   so you get the right notes without a chicken-and-egg problem.

3. Commit the updated CHANGELOG.md:

   ```bash
   git add CHANGELOG.md
   git commit -m "docs(changelog): update for vX.Y.Z"
   git push origin <default-branch>
   ```

4. Tag and push:

   ```bash
   git tag -a vX.Y.Z -m "vX.Y.Z"
   git push origin vX.Y.Z
   ```

5. Create the GitLab Release with the extracted notes:

   ```bash
   glab release create vX.Y.Z -F release_notes.md --repo <group/subgroup/project>
   ```

   Or without `glab`, via API:

   ```bash
   glab api -X POST "projects/<PROJECT_ID>/releases" \
     -f "tag_name=vX.Y.Z" \
     -f "description=$(cat release_notes.md)"
   ```

6. Clean up the local scratch file:

   ```bash
   rm release_notes.md
   ```

---

# Path B — GitLab native Changelog API (Git trailers)

Use only if the repo's commits already end with a `Changelog: <type>` trailer
(`added`, `fixed`, `changed`, `deprecated`, `removed`, `security`, `other`) —
not Conventional Commit prefixes. Official GitLab tutorial:
[Automate releases and release notes with GitLab](https://about.gitlab.com/blog/tutorial-automated-release-and-release-notes-with-gitlab/).

No local tool needed — everything runs through GitLab's own API, including
the CHANGELOG.md commit itself (no `git clone`/push, no job-token setup).

## Setup (once per project)

1. Create a **Project Access Token** with the `api` scope: repo → Settings →
   Access Tokens.
2. Store it as an env var / password manager entry (this is a manual flow,
   not a CI/CD variable — no CI job reads it).

## Cutting a release (repeat every release)

1. Commits destined for the changelog must have a trailer, e.g.:
   ```
   Add ChatBot

   Changelog: added
   ```
   For squash-merged MRs, put the trailer on the MR's squash/merge commit
   message (edit it in the "Ready to merge" box before merging).

2. Tag the release: `git tag -a vX.Y.Z -m "vX.Y.Z" && git push origin vX.Y.Z`

3. Generate release notes from the API (writes markdown, does NOT touch the
   repo):
   ```bash
   curl -H "PRIVATE-TOKEN: $TOKEN" \
     "https://gitlab.com/api/v4/projects/<PROJECT_ID>/repository/changelog?version=vX.Y.Z" \
     | jq -r .notes > release_notes.md
   ```

4. Create the GitLab Release:
   ```bash
   glab release create vX.Y.Z -F release_notes.md --repo <group/subgroup/project>
   ```

5. **Optional** — persist the same notes into a `CHANGELOG.md` file in the
   repo. Same call with `-X POST` instead of GET — it commits directly to the
   default branch via the API:
   ```bash
   curl -H "PRIVATE-TOKEN: $TOKEN" -X POST \
     "https://gitlab.com/api/v4/projects/<PROJECT_ID>/repository/changelog?version=vX.Y.Z"
   ```

## Customizing categories/template

Add `.gitlab/changelog_config.yml` to rename categories or change the
template — see the [Changelogs docs](https://docs.gitlab.com/user/project/changelogs/#customize-the-changelog-output)
for the full templating language. Also lets you override `tag_regex` if tags
don't follow `vMAJOR.MINOR.PATCH`.

---

## Checklist

**Path A (git-cliff):**
- [ ] `cliff.toml` present at repo root, `commit_parsers` matches this
      project's Conventional Commit types in use
- [ ] `CHANGELOG.md` generated and committed
- [ ] Version decided (semver, based on change content not commit count)
- [ ] `release_notes.md` generated with `--unreleased --tag vX.Y.Z`
- [ ] Tag created + pushed
- [ ] GitLab Release created via `glab release create` with the notes
- [ ] Scratch `release_notes.md` removed

**Path B (native API):**
- [ ] Project access token (scope `api`) available (env var / password
      manager, not a CI/CD variable)
- [ ] Commits destined for the changelog carry a `Changelog: <type>` trailer
- [ ] Tag created + pushed
- [ ] `release_notes.md` fetched via GET, GitLab Release created from it
- [ ] (optional) CHANGELOG.md persisted via the `-X POST` variant

**Both paths — do NOT:**
- [ ] Wire `changelog`/`release` jobs into `.gitlab-ci.yml` gated on
      `rules: if: $CI_COMMIT_TAG` and assume they'll fire. Confirmed
      non-functional on at least one GitLab.com project even with an
      unconditional (no-rules) job. Revisit only if a group Owner confirms
      the root cause (suspected compute-minutes quota) and fixes it.
