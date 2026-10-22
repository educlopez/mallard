---
name: lando-check
description: >
  Check that a project's local Lando environment is up and serving, start it if
  not, find the real local URL, and diagnose the usual failures (Docker not
  running, port already allocated, stale containers, shop redirecting to the
  production domain). Use when the user says "levanta lando", "está arriba el
  local?", "verifica en local", "comprueba el sitio local", "no carga el
  lndo.site", "lando no arranca", "puerto ocupado", or before verifying a change
  in the browser.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# lando-check

Gets from "is the local site working?" to a verified answer in one pass. It does not set a project up from scratch (use `lando-setup` for that) and it does not change code.

## 1. Docker engine

`lando` needs a running Docker engine (OrbStack or Docker Desktop).

```bash
docker info >/dev/null 2>&1 && echo "docker ok" || echo "docker down"
```

If it is down, open the app the user uses (`open -a OrbStack` or `open -a Docker`), wait until `docker info` succeeds, then continue. Do not run `lando` commands before that.

## 2. Is this project running?

From the project root (the directory with `.lando.yml`):

```bash
lando info 2>&1 | grep -E 'urls|https?://|service:' | head -20
lando list 2>&1 | grep -i "$(basename "$PWD")" | head
```

- No `.lando.yml` in the directory: stop and say so. Offer `lando-setup` if the project was never set up locally.
- Containers `running` and URLs listed: go to step 3.
- Not running: `lando start 2>&1 | tail -25`, then back to step 3.

## 3. Health check

Use the URL from `lando info` (usually `https://<name>.lndo.site`; multistore projects have several domains, check each one the user cares about):

```bash
curl -sk -o /dev/null -w "%{http_code} %{redirect_url}\n" "https://<name>.lndo.site/"
```

| Result | Meaning | Action |
|---|---|---|
| `200` | Serving | Report OK. |
| `301/302` to the **production domain** | `ps_shop_url` / `.htaccess` still point to prod (usual after a DB import) | Fix with the `/ps-url` command or the DB refresh step of `lando-setup`. Do not edit prod config. |
| `301/302` to `http://` or another `lndo.site` | Wrong scheme or domain mismatch | Check `ps_shop_url` and the domain in `.lando.yml` proxy. |
| `500` / `502` | PHP or nginx error | `lando logs -s appserver 2>&1 | tail -30` and report the real error. |
| connection refused / timeout | Proxy or containers down | Step 4. |

## 4. Common failures

- **`port is already allocated` / `address already in use`** (usually the MySQL `portforward` in `.lando.yml`, or 80/443 held by another Lando app or a local server):
  ```bash
  lsof -nP -iTCP:<port> -sTCP:LISTEN
  lando list 2>&1 | grep -E 'name|running|port'
  grep -rn portforward ~/Developer/*/*/.lando.yml 2>/dev/null | head
  ```
  Tell the user which process/app holds the port. Do not kill processes or stop other projects without asking. If the clash is another project's `portforward`, pick the next free port in this project's `.lando.yml` and ask before editing.
- **Stale or half-started app** (`lando start` hangs or errors on the first try): `lando restart`; if it persists, `lando rebuild -y` (asks before it, because it recreates containers; the DB volume is kept).
- **Browser says certificate error**: Lando's CA is not trusted on this machine; the `curl -k` check still works. Report it, do not change system trust settings.

## 5. Report

One short block: project, Docker state, containers, URL(s) and HTTP status, and anything the user must fix. If you started Lando, say so.
