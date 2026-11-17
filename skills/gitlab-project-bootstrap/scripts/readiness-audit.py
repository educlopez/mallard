#!/usr/bin/env python3
"""Read-only agentic-readiness audit of GitLab projects (see SKILL.md, section 8).

Usage:
  readiness-audit.py [--namespace PREFIX] [--project group/name ...] [--json]

Without arguments it audits every project the user is a member of. It never
writes anything. Requests are retried with a back-off on HTTP 429.

Per project it reports: default branch protected (and how), whether protection
is hollow (all human members Maintainer or above), AGENTS.md / CLAUDE.md on the
default branch, a CI file and whether a pipeline was ever created, and which
common paths .gitignore covers. "None" means the value could not be read.
"""
import argparse, json, subprocess, sys, time, urllib.parse as up
from concurrent.futures import ThreadPoolExecutor


def glab(path, raw=False):
    cmd = ["glab", "api", path] + (["--paginate"] if "per_page" in path else [])
    for attempt in range(6):
        r = subprocess.run(cmd, capture_output=True, text=True)
        if "429" in r.stderr or "Rate limit" in r.stdout:
            time.sleep(15 * (attempt + 1))
            continue
        break
    if r.returncode:
        return "404" if ("404" in r.stderr or "Not Found" in r.stdout) else None
    if raw:
        return r.stdout
    try:
        return json.loads(r.stdout)
    except ValueError:
        out, dec, i, s = [], json.JSONDecoder(), 0, r.stdout.strip()
        while i < len(s):
            obj, i = dec.raw_decode(s, i)
            out += obj if isinstance(obj, list) else [obj]
            while i < len(s) and s[i].isspace():
                i += 1
        return out


def ok(v):
    return v not in (None, "404")


def audit(p):
    pid, branch = p["id"], p.get("default_branch") or "main"
    q = up.quote(branch)

    def fetch(name):
        r = glab(f"projects/{pid}/repository/files/{up.quote(name, safe='')}/raw?ref={q}", raw=True)
        return None if r is None else ("" if r == "404" else r)

    prot = glab(f"projects/{pid}/protected_branches")
    members = glab(f"projects/{pid}/members/all?per_page=100")
    pipes = glab(f"projects/{pid}/pipelines?per_page=1")
    agents, claude, ci, gi = (fetch(f) for f in ("AGENTS.md", "CLAUDE.md", ".gitlab-ci.yml", ".gitignore"))

    protected = push = None
    if ok(prot):
        match = [b for b in prot if b["name"] == branch]
        protected = bool(match)
        push = [x["access_level_description"] for x in match[0]["push_access_levels"]] if match else []
    humans = [m for m in members if m.get("access_level", 0) < 50] if ok(members) else None
    hollow = None
    if humans is not None and protected:
        hollow = bool(humans) and all(m["access_level"] >= 40 for m in humans) and "No one" not in (push or [])
    lines = [l.strip() for l in (gi or "").splitlines() if l.strip() and not l.startswith("#")]
    covers = {k: any(any(x in l for x in v) for l in lines) for k, v in {
        "parameters.php": ("parameters.php",), "img": ("img/", "/img"), "upload": ("upload",),
        "download": ("download",), "var/cache": ("var/cache", "var/"), "node_modules": ("node_modules",)}.items()} if gi is not None else None
    return {
        "project": p["path_with_namespace"], "default_branch": branch, "protected": protected, "push": push,
        "hollow_protection": hollow, "agents_md": None if agents is None else bool(agents),
        "claude_md": None if claude is None else ("bridge" if claude.strip() == "@AGENTS.md" else bool(claude)),
        "ci_file": None if ci is None else bool(ci),
        "pipeline_ever_created": None if not ok(pipes) else bool(pipes),
        "gitignore_covers": covers,
    }


def main():
    ap = argparse.ArgumentParser()
    ap.add_argument("--namespace", default="")
    ap.add_argument("--project", action="append", default=[])
    ap.add_argument("--json", action="store_true")
    a = ap.parse_args()
    if a.project:
        projects = [glab(f"projects/{up.quote(x, safe='')}") for x in a.project]
    else:
        projects = glab("projects?membership=true&archived=false&per_page=100") or []
        projects = [p for p in projects if p["path_with_namespace"].startswith(a.namespace)]
    with ThreadPoolExecutor(2) as ex:
        rows = list(ex.map(audit, [p for p in projects if isinstance(p, dict)]))
    if a.json:
        json.dump(rows, sys.stdout, indent=1)
        return
    for r in rows:
        flags = [k for k in ("agents_md", "ci_file", "pipeline_ever_created") if r[k] is False]
        if r["protected"] is False:
            flags.append("default branch unprotected")
        if r["hollow_protection"]:
            flags.append("hollow protection")
        if r["claude_md"] is True:
            flags.append("CLAUDE.md has its own content")
        print(f'{r["project"]:<50} {", ".join(flags) or "ok"}')


if __name__ == "__main__":
    main()
