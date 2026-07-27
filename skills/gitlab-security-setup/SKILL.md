---
name: gitlab-security-setup
description: >
  Sets up a full security stack on your company's projects hosted on GitLab.com
  (non-PrestaShop: Laravel, Astro, TanStack, etc.). Use ONLY when the project is
  a GitLab.com Free tier project. Triggers when the user asks
  to add dependency scanning, vulnerability alerts, security setup, Trivy, pnpm
  supply chain protection, or wants email reports of vulnerabilities. Do NOT use
  for GitHub-hosted projects, personal projects, or PrestaShop projects — use
  ps-security-audit skill instead for any PrestaShop project.
version: "0.2.0"
metadata:
  author: Eduardo Calvo
---

# GitLab Security Setup

Full security stack for your company's GitLab.com projects on the Free tier.
Covers: pnpm 11 supply chain, Trivy weekly scan, HTML email reports via Gmail.

> **Placeholder:** `{report_recipients}` is a comma-separated list of email
> addresses that receive the vulnerability reports (e.g. `you@example.com, teammate@example.com`).
> Replace it everywhere it appears below with your own recipient address(es) before running.

## What gets set up

1. **pnpm 11** with supply chain protection (`minimumReleaseAge`, overrides)
2. **Trivy** vulnerability + secret scanner via GitLab CI
3. **Weekly scheduled pipeline** (Monday 8am Madrid) with HTML email report
4. **Composer audit** for PHP/Laravel projects
5. **Gmail SMTP** delivery via GitLab CI/CD variables

---

## Step 1 — pnpm 11 Supply Chain

### `pnpm-workspace.yaml` (create or update)

```yaml
# WARNING: single-package repos do NOT need a `packages:` block on pnpm 11.
# BUT on pnpm 9 (Vercel default for older projects) the mere presence of this
# file REQUIRES a non-empty `packages:` or install dies with
# "packages field missing or empty". If targeting pnpm 9, add `packages: ['.']`.

minimumReleaseAge: 4320   # packages must be 72h old before install (minutes)

# Block transitive deps from git repos / raw tarball URLs (needs pnpm 10.26+,
# silently inert below). See supply-chain-security skill for the full checklist.
blockExoticSubdeps: true

# Allowlist for postinstall/build scripts. pnpm 10+ blocks ALL by default.
# List ONLY packages that genuinely need to compile. pnpm will prompt you to
# add new entries when you install a dep with a blocked build script.
allowBuilds:
  esbuild: true
  sharp: true
  # lightningcss-cli: true   # uncomment if using lightningcss

overrides:
  form-data: ">=4.0.6"
  axios: ">=1.15.2"
  lodash: ">=4.18.0"
  picomatch: ">=4.0.4"
  qs: ">=6.14.2"
  shell-quote: ">=1.8.4"
```

**Rules:**
- `minimumReleaseAge` is in **minutes** (4320 = 72h). Blocks supply chain attacks via typosquatting/fast-publish.
- `blockExoticSubdeps` needs pnpm 10.26+ — silently inert below. Check the pinned version.
- `allowBuilds` — pnpm 10+ blocks all postinstall scripts by default. Add ONLY packages that need to compile. pnpm 11 will tell you during install if a new dep needs adding here.
- `overrides` pins known vulnerable transitive deps. Add new entries as CVEs appear.
- Do NOT put `minimumReleaseAge` in `.npmrc` — pnpm 11 reads it from `pnpm-workspace.yaml` only.
- **Real test is CI**: `pnpm install --frozen-lockfile` on a clean machine enforces all policies; a warm local cache skips them.

### `package.json` additions

```json
{
  "packageManager": "pnpm@11.1.2",
  "private": true
}
```

Pin the **exact** version (not `11.x.x`) so CI/Vercel use the version you tested.
Remove any `overrides` or `pnpm.overrides` blocks from `package.json` — they belong in `pnpm-workspace.yaml` for pnpm 11.

### `publicar` deploy script (if project has one)

```bash
#!/bin/bash
php artisan migrate --force
pnpm build
```

Ensure it uses `pnpm`, not `npm run`.

### GitHub Actions lint workflow (if exists)

Replace `npm ci` / `npm install` / `npm run` with:
```yaml
- run: npm install -g pnpm
- run: pnpm install --frozen-lockfile
- run: pnpm run format
- run: pnpm run lint
```

---

## Step 2 — GitLab CI Trivy Scan

Create or update `.gitlab-ci.yml`:

```yaml
dependency-scan:
  image:
    name: aquasec/trivy:latest
    entrypoint: [""]
  before_script:
    - apk add --no-cache curl python3 py3-packaging
  script:
    # JSON report (structured data) — vuln + secret scan
    - trivy fs --exit-code 0 --scanners vuln,secret --format json -o trivy-report.json . 2>/dev/null

    # HTML report artifact
    - trivy fs --exit-code 0 --scanners vuln,secret --format template --template "@/contrib/html.tpl" -o trivy-report.html . 2>/dev/null || true

    # Parse JSON and build HTML email
    - |
      python3 << 'PYEOF'
      import json, os

      with open("trivy-report.json") as f:
          data = json.load(f)

      from packaging.version import Version, InvalidVersion

      def highest_fix(fixed_str):
          if not fixed_str or fixed_str == "N/A":
              return None
          parts = [p.strip() for p in fixed_str.split(",") if p.strip()]
          parsed = []
          for p in parts:
              try:
                  parsed.append((Version(p), p))
              except InvalidVersion:
                  parsed.append((Version("0"), p))
          return max(parsed, key=lambda x: x[0])[1] if parsed else None

      grouped = {}
      severity_order = {"CRITICAL": 0, "HIGH": 1, "MEDIUM": 2, "LOW": 3, "UNKNOWN": 4}

      for result in data.get("Results", []):
          for v in result.get("Vulnerabilities", []):
              key = (v.get("PkgName", ""), v.get("InstalledVersion", ""))
              fixed_raw = v.get("FixedVersion", "")
              sev = v.get("Severity", "UNKNOWN")
              cve = v.get("VulnerabilityID", "")
              fix = highest_fix(fixed_raw)
              if key not in grouped:
                  grouped[key] = {"pkg": key[0], "installed": key[1], "severity": sev, "cves": [], "fixes": []}
              entry = grouped[key]
              if severity_order.get(sev, 5) < severity_order.get(entry["severity"], 5):
                  entry["severity"] = sev
              if cve:
                  entry["cves"].append(cve)
              if fix:
                  entry["fixes"].append(fix)

      def max_version(versions):
          parsed = []
          for v in versions:
              try:
                  parsed.append((Version(v), v))
              except InvalidVersion:
                  pass
          return max(parsed, key=lambda x: x[0])[1] if parsed else None

      vulns = list(grouped.values())
      for entry in vulns:
          entry["best_fix"] = max_version(entry["fixes"])
      vulns.sort(key=lambda x: severity_order.get(x["severity"], 5))

      counts = {s: sum(1 for v in vulns if v["severity"] == s) for s in ["CRITICAL", "HIGH", "MEDIUM", "LOW"]}
      total = len(vulns)

      sev_colors = {
          "CRITICAL": ("#fff1f2", "#be123c", "#fecdd3"),
          "HIGH":     ("#fffbeb", "#b45309", "#fde68a"),
          "MEDIUM":   ("#eff6ff", "#1d4ed8", "#bfdbfe"),
          "LOW":      ("#f9fafb", "#374151", "#e5e7eb"),
      }

      project = os.environ.get("CI_PROJECT_NAME", "")
      branch = os.environ.get("CI_COMMIT_REF_NAME", "")
      sha = os.environ.get("CI_COMMIT_SHORT_SHA", "")
      pipeline_url = os.environ.get("CI_PIPELINE_URL", "#")
      gmail_user = os.environ.get("GMAIL_USER", "")
      # Comma-separated recipient list — set REPORT_RECIPIENTS as a CI/CD variable
      report_recipients = os.environ.get("REPORT_RECIPIENTS", "")
      pipeline_created_at = os.environ.get("CI_PIPELINE_CREATED_AT", "")
      scan_date = pipeline_created_at[:10] if pipeline_created_at else ""
      scan_ts = pipeline_created_at.replace("T", " ").replace("Z", " UTC") if pipeline_created_at else ""

      def cve_links(cves, limit=5):
          links = []
          for cve in cves[:limit]:
              if cve.startswith("CVE-"):
                  links.append(f"<a href='https://nvd.nist.gov/vuln/detail/{cve}' style='color:#f97316;text-decoration:none'>{cve}</a>")
              else:
                  links.append(cve)
          out = ", ".join(links)
          if len(cves) > limit:
              out += f"&nbsp;<span style='color:#98A2B3'>+{len(cves)-limit} more</span>"
          return out

      sev_pills = {
          "CRITICAL": ("background:#FFF1F3;color:#C01048;border:1px solid #FFC5D0", "C"),
          "HIGH":     ("background:#FFFAEB;color:#B54708;border:1px solid #FEDF89", "H"),
          "MEDIUM":   ("background:#EFF8FF;color:#175CD3;border:1px solid #B2DDFF", "M"),
          "LOW":      ("background:#F9FAFB;color:#344054;border:1px solid #D0D5DD", "L"),
      }

      dot = {"CRITICAL": "#F04438", "HIGH": "#F79009", "MEDIUM": "#2E90FA", "LOW": "#98A2B3"}

      rows = ""
      for i, v in enumerate(vulns):
          pill_style, _ = sev_pills.get(v["severity"], ("background:#F9FAFB;color:#344054;border:1px solid #D0D5DD", "?"))
          fixed = v["best_fix"] if v["best_fix"] else "<span style='color:#98A2B3'>No fix yet</span>"
          cves_html = cve_links(v["cves"])
          sep = "border-bottom:1px solid #EAECF0;" if i < len(vulns) - 1 else ""
          rows += f"""<tr style='background:#ffffff'>
            <td style='padding:8px 16px;{sep}font-size:13px;color:#101828;font-weight:500;white-space:nowrap'>{v['pkg']}</td>
            <td style='padding:8px 16px;{sep}font-family:ui-monospace,SFMono-Regular,Menlo,monospace;font-size:12px;color:#475467'>{v['installed']}</td>
            <td style='padding:8px 16px;{sep}font-family:ui-monospace,SFMono-Regular,Menlo,monospace;font-size:12px;color:#101828;font-weight:500'>{fixed}</td>
            <td style='padding:8px 16px;{sep}white-space:nowrap'>
              <span style='{pill_style};padding:2px 7px;border-radius:9999px;font-size:11px;font-weight:500;display:inline-block'>{v['severity'].capitalize()}</span>
            </td>
            <td style='padding:8px 16px;{sep}font-size:12px;color:#475467'>{cves_html}</td>
          </tr>"""

      badges = "".join([
          f"<span style='display:inline-block;padding:3px 10px;border-radius:9999px;font-size:12px;font-weight:500;margin-right:6px;{sev_pills[s][0]}'>{s.capitalize()} {counts[s]}</span>"
          for s in ["CRITICAL", "HIGH", "MEDIUM", "LOW"] if counts[s] > 0
      ])

      table_section = "" if not vulns else f"""
      <table border='0' width='100%' cellpadding='0' cellspacing='0' role='presentation' style='border-collapse:separate;border-spacing:0;border:1px solid #EAECF0;border-radius:8px;overflow:hidden;margin-top:20px'>
        <thead>
          <tr style='background:#F9FAFB'>
            <th style='padding:10px 16px;text-align:left;font-size:11px;font-weight:500;color:#475467;border-bottom:1px solid #EAECF0'>Package</th>
            <th style='padding:10px 16px;text-align:left;font-size:11px;font-weight:500;color:#475467;border-bottom:1px solid #EAECF0'>Installed</th>
            <th style='padding:10px 16px;text-align:left;font-size:11px;font-weight:500;color:#475467;border-bottom:1px solid #EAECF0'>Fix to</th>
            <th style='padding:10px 16px;text-align:left;font-size:11px;font-weight:500;color:#475467;border-bottom:1px solid #EAECF0'>Severity</th>
            <th style='padding:10px 16px;text-align:left;font-size:11px;font-weight:500;color:#475467;border-bottom:1px solid #EAECF0'>CVE</th>
          </tr>
        </thead>
        <tbody>{rows}</tbody>
      </table>"""

      no_vulns = """
      <div style='text-align:center;padding:48px 0'>
        <p style='font-size:18px;font-weight:600;color:#111827;margin:0 0 8px'>All clear!</p>
        <p style='font-size:14px;color:#6b7280;margin:0'>No vulnerabilities found in this scan.</p>
      </div>""" if not vulns else ""

      md_lines = [
          f"# Security Scan — {project}",
          f"Branch: {branch} | Commit: {sha}" + (f" | {scan_date}" if scan_date else ""),
          "", "## Vulnerabilities", "",
          "| Package | Installed | Fix version | Severity | CVEs |",
          "|---------|-----------|-------------|----------|------|",
      ]
      for v in vulns:
          fix = v["best_fix"] if v["best_fix"] else "No fix yet"
          md_lines.append(f"| {v['pkg']} | {v['installed']} | {fix} | {v['severity']} | {', '.join(v['cves'])} |")
      md_lines += ["", "## Task",
          "Review these vulnerabilities. For each package identify if it's a direct or transitive dependency and provide the exact command to update it."]
      md_content = "\n".join(md_lines)

      with open("trivy-report.md", "w") as f:
          f.write(md_content)

      md_escaped = md_content.replace("&", "&amp;").replace("<", "&lt;").replace(">", "&gt;")

      summary_text = f"{total} package{'s' if total != 1 else ''} with known vulnerabilities"
      if scan_ts:
          summary_text += f"&ensp;&middot;&ensp;{scan_ts}"

      meta_date = f'&ensp;<span style="color:rgb(163,163,163)">{scan_date}</span>' if scan_date else ""

      subject_suffix = f" · {scan_date}" if scan_date else ""
      subject = f"[{project}] Security Scan · {counts.get('CRITICAL',0)}C {counts.get('HIGH',0)}H{subject_suffix}"

      html = (
          f"From: {gmail_user}\r\nTo: {report_recipients}\r\nSubject: {subject}\r\nMIME-Version: 1.0\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n"
          f"<!DOCTYPE html><html dir='ltr' lang='en'><head>"
          f"<meta content='text/html; charset=UTF-8' http-equiv='Content-Type'>"
          f"<meta name='color-scheme' content='light'><meta name='supported-color-schemes' content='light'></head>"
          f"<body style='background-color:rgb(250,250,250);margin:0;padding:0'>"
          f"<table border='0' width='100%' cellpadding='0' cellspacing='0' role='presentation' align='center'><tbody><tr>"
          f"<td style='margin:0;background-color:rgb(250,250,250);padding:0;font-family:Inter,-apple-system,Segoe UI,system-ui,Roboto,Arial,sans-serif;-webkit-font-smoothing:antialiased'>"
          f"<table align='center' width='100%' border='0' cellpadding='0' cellspacing='0' role='presentation' style='max-width:640px;width:100%;background-color:rgb(250,250,250)'>"
          f"<tbody><tr style='width:100%'><td>"
          f"<table align='left' width='100%' border='0' cellpadding='0' cellspacing='0' role='presentation' style='max-width:100%;background-color:rgb(255,255,255);padding:24px'>"
          f"<tbody><tr style='width:100%'><td>"
          f"<table border='0' width='100%' cellpadding='0' cellspacing='0' role='presentation'><tbody><tr>"
          f"<td style='vertical-align:middle'>"
          f"<img alt='DepSheriff' src='https://ik.imagekit.io/16u211libb/security-logo.png' style='display:inline-block;outline:none;border:none;text-decoration:none;height:32px;width:auto;vertical-align:middle;margin-right:10px'>"
          f"<span style='font-size:15px;font-weight:600;color:rgb(23,23,23);vertical-align:middle'>DepSheriff</span></td>"
          f"<td style='text-align:right;vertical-align:middle'><span style='font-size:11px;color:rgb(163,163,163);font-family:ui-monospace,monospace'>{project}</span></td>"
          f"</tr></tbody></table></td></tr></tbody></table>"
          f"<table align='left' width='100%' border='0' cellpadding='0' cellspacing='0' role='presentation' style='max-width:100%;background-color:rgb(255,255,255);padding:32px 24px'>"
          f"<tbody><tr style='width:100%'><td>"
          f"<p style='font-size:24px;line-height:32px;margin:0;font-weight:600;color:rgb(23,23,23)'>Security Scan Report</p>"
          f"<table border='0' width='100%' cellpadding='0' cellspacing='0' role='presentation' style='margin-top:8px;margin-bottom:32px'>"
          f"<tbody><tr><td style='font-size:14px;color:rgb(82,82,82)'>"
          f"Branch <code style='background:#f3f4f6;padding:1px 5px;border-radius:4px;font-size:12px'>{branch}</code>"
          f"&ensp;Commit <code style='background:#f3f4f6;padding:1px 5px;border-radius:4px;font-size:12px'>{sha}</code>"
          f"{meta_date}</td></tr></tbody></table>"
          f"<div style='margin-bottom:6px'>{badges}</div>"
          f"<p style='font-size:13px;color:rgb(163,163,163);margin:0 0 8px;line-height:20px'>{summary_text}</p>"
          f"{no_vulns}{table_section}"
          f"<table border='0' width='100%' cellpadding='0' cellspacing='0' role='presentation' style='margin-top:28px'><tbody><tr><td>"
          f"<p style='font-size:14px;font-weight:600;color:#101828;margin:0 0 8px 0'>Paste to Agent</p>"
          f"<p style='font-size:14px;color:#475467;margin:0 0 12px 0'>Copy the block below into your AI agent for fix recommendations.</p>"
          f"<pre style='margin:0;background:#F9FAFB;padding:16px;border-radius:8px;font-size:11px;white-space:pre-wrap;border:1px solid #EAECF0;color:#344054;overflow-x:auto;line-height:1.6'>{md_escaped}</pre>"
          f"</td></tr></tbody></table>"
          f"<table align='center' border='0' width='100%' cellpadding='0' cellspacing='0' role='presentation' style='margin-top:32px'>"
          f"<tbody><tr style='width:100%'><td align='center'>"
          f"<a href='{pipeline_url}' style='line-height:24px;text-decoration:none;display:inline-block;border-radius:8px;background-color:#f97316;border:1px solid #ea6d10;color:#ffffff;padding:10px 20px;font-size:14px;font-weight:600'>"
          f"View Pipeline &amp; Download Artifacts</a></td></tr></tbody></table>"
          f"</td></tr></tbody></table>"
          f"<table align='left' width='100%' border='0' cellpadding='0' cellspacing='0' role='presentation' style='max-width:100%;background-color:rgb(255,255,255);padding:24px'>"
          f"<tbody><tr style='width:100%'><td>"
          f"<table align='center' width='100%' border='0' cellpadding='0' cellspacing='0' role='presentation' style='margin-bottom:24px'>"
          f"<tbody><tr><td><hr style='width:100%;border:none;border-top:1px solid rgb(229,229,229)'></td></tr></tbody></table>"
          f"<p style='font-size:13px;line-height:20px;margin:0;text-align:center;color:rgb(163,163,163)'>DepSheriff &mdash; because npm audit alone wasn&apos;t enough</p>"
          f"<p style='font-size:11px;margin:10px 0 0;text-align:center;color:rgb(163,163,163)'>Creado por <a href='https://www.linkedin.com/in/educlopez/' style='color:rgb(163,163,163);text-decoration:underline'>Edu Calvo</a></p>"
          f"</td></tr></tbody></table>"
          f"</td></tr></tbody></table></td></tr></tbody></table></body></html>"
      )

      with open("email_body.txt", "w") as f:
          f.write(html)

      print(f"Vulnerabilities found: {total} (CRITICAL: {counts['CRITICAL']}, HIGH: {counts['HIGH']})")
      PYEOF

    - |
      # REPORT_RECIPIENTS is a comma-separated CI/CD variable (e.g. "a@x.com,b@y.com").
      # Build one --mail-rcpt flag per address.
      RCPT_ARGS=""
      IFS=',' read -ra ADDRS <<< "$REPORT_RECIPIENTS"
      for addr in "${ADDRS[@]}"; do
        addr="$(echo "$addr" | xargs)"   # trim whitespace
        [ -n "$addr" ] && RCPT_ARGS="$RCPT_ARGS --mail-rcpt $addr"
      done
      curl --url "smtps://smtp.gmail.com:465" \
        --ssl-reqd \
        --mail-from "$GMAIL_USER" \
        $RCPT_ARGS \
        --user "$GMAIL_USER:$GMAIL_APP_PASS" \
        -T email_body.txt
  artifacts:
    paths:
      - trivy-report.html
      - trivy-report.json
      - trivy-report.md
    expire_in: 30 days
  rules:
    - if: $CI_PIPELINE_SOURCE == "schedule"
```

**Key implementation details:**
- `entrypoint: [""]` — mandatory; Trivy Docker image has no shell otherwise (exit code 127)
- `--exit-code 0` — never fail the pipeline; email even when clean
- `--scanners vuln,secret` — covers both dependency CVEs and leaked secrets
- `rules: schedule` — only runs on scheduled pipelines, not every push
- Vulnerabilities grouped by `(package, installed_version)` — one row per package, showing highest severity and best fix version
- `py3-packaging` via apk — not pip (pip is blocked in Alpine CI)
- **Recipients are parametrized** via the `REPORT_RECIPIENTS` CI/CD variable (comma-separated) in BOTH the Python `To:` header and the curl `--mail-rcpt` loop — never hardcode addresses in the job.
- **"DepSheriff" branded email** — modern Untitled-UI HTML (Inter font, rounded severity pills, orange accent, logo). Trivy's `contrib/html.tpl` artifact still ships too.
- **CVE → NVD links** via `cve_links()` (first 5 linked, rest collapsed to `+N more`).
- **Scan date/timestamp** from `CI_PIPELINE_CREATED_AT`, shown in the header and the subject (`[proj] Security Scan · 2C 5H · YYYY-MM-DD`).
- **"Paste to Agent"** block — a markdown table of findings ready to drop into any AI agent for fix recommendations.
- The logo (`ik.imagekit.io/.../security-logo.png`) and footer credit are the DepSheriff brand — keep or swap for your own.

---

## Step 3 — Gmail CI/CD Variables

In GitLab project: **Settings → CI/CD → Variables**

| Variable | Value | Protected | Masked |
|----------|-------|-----------|--------|
| `GMAIL_USER` | `your-account@gmail.com` | No | No |
| `GMAIL_APP_PASS` | App password from Google | No | Yes |
| `REPORT_RECIPIENTS` | `you@example.com,teammate@example.com` (comma-separated) | No | No |

> Set `REPORT_RECIPIENTS` to your own recipient address(es) — this is who receives the vulnerability report email.

**Getting Gmail App Password:**
1. Google Account → Security → 2-Step Verification (must be ON)
2. Search "App passwords" → Create → name it "GitLab CI"
3. Copy the 16-char password → paste as `GMAIL_APP_PASS`

---

## Step 4 — Weekly Scheduled Pipeline

Create via GitLab API (run once per project):

```bash
curl --request POST \
  --header "PRIVATE-TOKEN: <your-gitlab-token>" \
  "https://gitlab.com/api/v4/projects/<PROJECT_ID>/pipeline_schedules" \
  --form "description=Weekly security scan" \
  --form "ref=main" \
  --form "cron=0 7 * * 1" \
  --form "cron_timezone=Europe/Madrid"
```

- `0 7 * * 1` = Monday 07:00 UTC = 08:00/09:00 Madrid (winter/summer)
- `ref` = default branch (`main` or `develop`)
- `PROJECT_ID` = GitLab project → Settings → General

Or via UI: **CI/CD → Schedules → New schedule**

**Trigger manually to test:**

```bash
curl --request POST \
  --header "PRIVATE-TOKEN: <token>" \
  "https://gitlab.com/api/v4/projects/<PROJECT_ID>/pipeline_schedules/<SCHEDULE_ID>/play"
```

---

## Step 5 — PHP/Composer Projects (Laravel)

The Trivy `fs` scan above **already reads `composer.lock`** and reports PHP CVEs in
the email — so PHP deps are covered out of the box. `composer audit` adds the
**Packagist Security Advisories** feed on top (some advisories land there before a
CVE is assigned).

Do NOT add `composer audit` inline to the `dependency-scan` job — the `aquasec/trivy`
Alpine image has no PHP/Composer binary, so it would silently no-op. Add a **separate
job** with a `composer` image instead:

```yaml
# Packagist advisory audit — complements the Trivy composer.lock scan.
# Output goes to the job log (not the email). Non-blocking.
composer-audit:
  image: composer:2
  script:
    - composer audit --no-interaction --format=plain || true
  rules:
    - if: $CI_PIPELINE_SOURCE == "schedule"
  allow_failure: true
```

For fixing PHP vulnerabilities locally:

```bash
# Update specific packages
composer update "symfony/*" --with-all-dependencies

# Update all PHP deps (careful — test after)
composer update
```

Common PHP transitive dep CVEs — update these when flagged:
- `symfony/*` — update to latest patch on your major (e.g. `7.4.x`)
- `phpunit/phpunit` + `pestphp/pest` — must update together: `composer update phpunit/phpunit pestphp/pest --with-all-dependencies`
- `league/commonmark`, `psy/psysh` — `composer update <package>`

---

## Updating overrides for new CVEs

When the scan reports a fixable HIGH/CRITICAL on a transitive dep:

1. Check if it's in `pnpm-workspace.yaml` `overrides` already → bump version
2. If new package → add entry: `package-name: ">=fixed-version"`
3. Run `pnpm install` to regenerate lockfile
4. Commit + push → next weekly scan should show it resolved

**Skip alpha/RC fixes:** If the only fix is an alpha (e.g. `8.0.0-alpha.17`), skip — wait for stable release.

---

## Adapting for different project types

| Project type | Notes |
|---|---|
| **Laravel** | Include `composer audit` step; `publicar` = `php artisan migrate --force && pnpm build` |
| **Astro** | No composer step; `publicar` = `pnpm build` |
| **TanStack / pure frontend** | No composer step; check `pnpm-workspace.yaml` at root |
| **No Node** | Skip pnpm setup; Trivy still scans PHP deps |

---

## Checklist for new project

- [ ] `pnpm-workspace.yaml` created with `minimumReleaseAge: 4320` + `blockExoticSubdeps: true` + `overrides` + `allowBuilds` (add only what compiles)
- [ ] `package.json` has `"packageManager": "pnpm@11.x.x"` (exact version), `"private": true`, no `overrides` block
- [ ] Verified with clean `pnpm install --frozen-lockfile` (not just warm cache)
- [ ] `.gitlab-ci.yml` has `dependency-scan` job (Step 2) with recipients via `REPORT_RECIPIENTS` (never hardcoded)
- [ ] Laravel/PHP: separate `composer-audit` job added (Step 5) — NOT inline in the trivy job
- [ ] GitLab CI/CD variables set: `GMAIL_USER`, `GMAIL_APP_PASS`, `REPORT_RECIPIENTS`
- [ ] Pipeline schedule created (Monday 8am Madrid)
- [ ] Manual trigger test → email received by every address in `REPORT_RECIPIENTS`
- [ ] `pnpm install` runs clean locally
