---
name: lando-mail-smtp
description: >
  Wires PrestaShop's outgoing email (BO Parámetros Avanzados > Correo
  electrónico) to Lando's mailhog service so local emails (order
  confirmations, module notifications, password resets, ...) are actually
  deliverable and inspectable instead of silently disappearing. Use when the
  user says "configura el mail en lando", "conecta mailhog", "los emails no
  llegan en local", "el correo no funciona en lando", "quiero ver los emails
  en local", "prueba el envío de correo", or after finishing a lando-setup
  full install if the project sends any transactional email (module
  notifications, order emails, password reset, contact form). Also trigger
  when debugging why a module's Mail::send() call produces no visible
  output locally.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# lando-mail-smtp

## Trigger

- "configura el mail en lando", "conecta mailhog", "los emails no llegan en local"
- "el correo no funciona en lando", "quiero ver los emails en local"
- "prueba el envío de correo" / "verifica que llega el email"
- Proactively, right after a `lando-setup` full install, if the project has any
  transactional email path (module notifications, order confirmation, password
  reset, contact form) — ask the user if they want mail wired up now.

## Why this is needed

Lando's `mail: type: mailhog` service alone does **not** make PrestaShop's
`Mail::send()` deliver anywhere. PrestaShop defaults to `PS_MAIL_METHOD=1`
(native PHP `mail()`), which mailhog never intercepts — emails just vanish
silently, no error, nothing in the logs. Mailhog only catches mail sent over
**SMTP**, so PrestaShop must be explicitly switched to SMTP mode and pointed
at the mailhog service.

## Prerequisite — mailhog service in `.lando.yml`

If not already present (the `lando-setup` skill's template doesn't add this
by default — add it separately), add to `services`:

```yaml
services:
  mail:
    type: mailhog
    hogfrom:
      - appserver
```

Run `lando rebuild -y` (or `lando start` if the project isn't running yet)
after adding it.

## Step 1 — Get the mailhog service's internal hostname

```bash
cd ~/developer/{workspace}/{project-name}
lando info
```

Look for the `mail` service block:

```
{ service: 'mail',
  urls: [ 'http://localhost:{PORT}' ],
  type: 'mailhog',
  internal_connection: { host: 'mail', port: '1025' },
  hostnames: [ 'mail.{lando-name}.internal' ],
  ... }
```

Two hostnames resolve to the mailhog container from inside `appserver`:
- `mail` (short alias)
- `mail.{lando-name}.internal` (project-scoped)

**Always use the project-scoped one** (`mail.{lando-name}.internal`), not the
bare `mail` alias. Lando's shared proxy network exposes the `mail`/`appserver`
short aliases to **every** running Lando project on the machine — with
several projects up at once, the short alias can resolve to a *different*
project's mailhog/appserver container (the exact collision this skill's
sibling `lando-setup` already warns about for `fastcgi_pass appserver:9000`
in nginx). The `.internal` form is scoped to this project only and can't
collide.

`urls` gives the mailhog **web UI** port (e.g. `http://localhost:32799`) —
that's where you'll read captured emails, separate from the SMTP port (always
`1025` internally).

## Step 2 — Configure PrestaShop's SMTP settings

### Via the BO (recommended — matches what a teammate would click through)

1. Log in to the BO.
2. **Parámetros Avanzados > Correo electrónico** (Advanced Parameters > Email).
3. Set:
   - **Servidor SMTP**: `mail.{lando-name}.internal` (from Step 1)
   - **Puerto SMTP**: `1025`
   - **Cifrado**: `Ninguno` / off — mailhog doesn't speak TLS
   - Usuario/contraseña: leave empty, mailhog needs none
4. Toggle **Usar SMTP** / activate SMTP mode.
5. Save.

### Via direct DB update (faster for scripted/agent setup, same effect)

```bash
lando ssh -s database -c "mysql -uroot {recipe} -e \"
  UPDATE ps_configuration SET value='2' WHERE name='PS_MAIL_METHOD';
  UPDATE ps_configuration SET value='mail.{lando-name}.internal' WHERE name='PS_MAIL_SERVER';
  UPDATE ps_configuration SET value='1025' WHERE name='PS_MAIL_SMTP_PORT';
  UPDATE ps_configuration SET value='off' WHERE name='PS_MAIL_SMTP_ENCRYPTION';
\""
```

Replace `ps_` with the project's actual table prefix and `{recipe}` with the
`.lando.yml` recipe name (also the default DB name/user/password — see
`lando-setup`). `PS_MAIL_METHOD`: `1` = native PHP `mail()` (default, broken
for mailhog), `2` = SMTP (what we want).

No `lando restart`/cache-clear needed for either method — `Configuration::get()`
reads live.

## Step 3 — Verify end-to-end

Trigger any email-sending action in the storefront or BO (password reset,
a module's notification button, an order confirmation, the passwordless
login code flow some PS9 projects use, ...), then check mailhog:

```bash
curl -s "http://localhost:{mailhog-port}/api/v2/messages?limit=1" | python3 -c "
import json,sys
d=json.load(sys.stdin)
m=d['items'][0]
print('To:', m['Content']['Headers'].get('To'))
print('Subject:', m['Content']['Headers'].get('Subject'))
"
```

Or just open `http://localhost:{mailhog-port}` in a browser — mailhog's own
UI lists every captured message with full HTML/text preview.

**`{mailhog-port}` changes across `lando restart`s** in some Lando versions —
don't hardcode it in scripts; re-run `lando info` (or grep its `mail` block)
each time before checking, or read the URL fresh via the browser tab if one's
already open.

## Gotchas

- **Confirm SMTP actually applied** — if emails still don't arrive after
  Step 2, re-check `PS_MAIL_METHOD` really is `2`. Some BO email settings
  screens silently fail to save if a required field (e.g. "De" sender email)
  is invalid; the DB update path skips that risk entirely.
- **`ps_shop_email` / `PS_SHOP_EMAIL` must be a syntactically valid address**
  for PrestaShop to attempt sending at all, even over SMTP — an empty or
  malformed sender address can cause `Mail::send()` to return `false` before
  ever touching SMTP.
- **A module that ignores `Mail::send()`'s return value never surfaces SMTP
  failures** — if verification (Step 3) shows nothing arriving, temporarily
  wrap the module's send call to `var_dump()`/log the boolean return, rather
  than assuming the module code itself is broken.
- **This is a per-project, per-environment setting** — `ps_configuration`
  lives in the DB, so a fresh DB import from server (`lando-setup`'s DB
  refresh mode) will overwrite `PS_MAIL_METHOD` back to whatever the server
  uses (almost always `1`, native `mail()`, since production servers don't
  run mailhog). **Re-run Step 2 after every DB refresh.**
