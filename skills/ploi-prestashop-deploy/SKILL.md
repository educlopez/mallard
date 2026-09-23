---
name: ploi-prestashop-deploy
description: >
  Deploys a PrestaShop 8/9 project from a local Lando setup onto a fresh
  Ploi-managed server (site creation, file upload, DB migration, and the
  PrestaShop-specific nginx rules Ploi's generic template doesn't know about).
  Use when the user says "monta esto en el servidor Ploi", "sube el proyecto a
  Ploi", "despliega en el server nuevo", "crea el site en Ploi para X", "pasa
  esto a producción/pre en Ploi", or is setting up a brand-new domain on a
  Ploi server for a PrestaShop project. Also use when debugging a
  freshly-deployed PrestaShop site on Ploi that shows a 404 on admin login,
  broken product/category images, a memory_limit fatal in the backoffice, or
  a redirect to the old local .lndo.site domain — these are almost always the
  same handful of known Ploi+PrestaShop gaps this skill documents, not a new
  bug in the project itself.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# ploi-prestashop-deploy

## Trigger

**New deploy** (fresh domain/site not yet on the server):
"monta esto en Ploi", "sube el proyecto a Ploi", "crea el site en Ploi para X", "despliega en el server nuevo"
→ Run all steps in order.

**Debugging an existing Ploi+PrestaShop deploy** (site already up, something's broken):
404 on `/adminXXX/login`, broken images, `Allowed memory size ... exhausted`, redirects to a
`.lndo.site` domain, admin panel unreachable
→ Skip straight to [Troubleshooting](#troubleshooting) — match the symptom to its cause there
before re-deriving it from scratch.

## Why this skill exists

Ploi is a generic server-management panel — its default nginx template is written for a
plain PHP script or a Laravel app, not PrestaShop. PrestaShop's admin panel and SEO-friendly
image URLs depend on rewrite rules that Apache normally provides via `.htaccess`, which nginx
does not read. Every symptom in [Troubleshooting](#troubleshooting) traces back to this one
gap, dressed up differently each time (a 404 that renders as the storefront's own theme, a
PHP memory fatal that looks unrelated, images that 404 despite the files existing on disk).
Once the PrestaShop-aware nginx config (Step 10) is in place, all of it resolves at once.

## Placeholders

- `{domain}` — the site's final domain (e.g. `shopname.pre-testing.com`)
- `{ssh_key}` — SSH key for the Ploi server (e.g. `id_rsa_cinetic`)
- `{server_ip}` — the Ploi server's IP
- `{local_project_dir}` — the local Lando project directory
- `{admin_dir}` — the project's admin folder name (random hash like `adminea8094707b02d4ca`,
  or a custom name like `baupanel` — check with `ls {local_project_dir} | grep -i admin`)
- `{db_name}` / `{db_user}` / `{db_pass}` — credentials Ploi generated for the site's database

## Step 1 — Confirm inputs

Ask the user (if not already given):
- Local project directory and its PrestaShop version (`grep -m1 VERSION src/Core/Version.php`
  or `grep _PS_VERSION_ config/defines.inc.php`)
- Target domain
- Server IP + SSH key
- Does the user have Ploi panel access themselves, or should you drive it? If they have
  access, most of Steps 2 and 10 are "give the user the panel steps/content", not something
  you do directly — Ploi's nginx `sites-available` files and `/etc/php/*/fpm/php.ini` are
  root-owned; without a passwordless sudo on the SSH account (the common case), you cannot
  edit them yourself and must go through the panel.

Test SSH connectivity before doing anything else:
```bash
ssh -i ~/.ssh/{ssh_key} -o ConnectTimeout=10 -o BatchMode=yes ploi@{server_ip} "whoami && hostname"
```

## Step 2 — Check PHP version compatibility BEFORE creating the site

Check PrestaShop's own compatibility table (search "PrestaShop PHP compatibility matrix") for
the exact minor version — it changes per release:
- PS 9.0 → PHP ≤ 8.4
- PS 9.1 → PHP ≤ 8.5

Check what's already installed on the server:
```bash
ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "ls /etc/php/"
```
If the needed version isn't listed, install it from Ploi panel → Server → PHP → Install new
version, before creating the site. Don't assume the server's *default* PHP version is right
for this specific PrestaShop minor version — check the table every time, PHP releases move
faster than PrestaShop's support window.

## Step 3 — Create the site in the Ploi panel

**Never `mkdir` the site folder over SSH.** Ploi's own database tracks sites, SSL certs, PHP-FPM
pool assignment, and deploy scripts — creating the folder by hand leaves all of that
disconnected from the panel (no SSL renewal, no PHP version management, nothing).

Site → New site → domain = `{domain}`, PHP version = the one confirmed in Step 2.
Confirms the convention: webroot ends up at `/home/ploi/{domain}/public`.

## Step 4 — Upload the codebase

Two paths: git-based deploy (Ploi's own deploy key + deploy script) or a direct rsync copy.
Ask the user which they want — rsync is faster to get a first working copy up when there's no
CI/CD pipeline wired yet; git-based deploy is right when the project is meant to auto-deploy
on push going forward. This skill covers rsync; if git-based, use Ploi's own "Repository" +
"Deploy script" site config panels instead of the steps below.

```bash
cd {local_project_dir} && rsync -az --progress -e "ssh -i ~/.ssh/{ssh_key}" \
  --exclude='.git/' \
  --exclude='.github/' \
  --exclude='.githooks/' \
  --exclude='.claude/' \
  --exclude='.codegraph/' \
  --exclude='.cursor/' \
  --exclude='.atl/' \
  --exclude='.lando/' \
  --exclude='.ui-craft/' \
  --exclude='.docker/' \
  --exclude='node_modules/' \
  --exclude='var/cache/' \
  --exclude='var/logs/' \
  ./ ploi@{server_ip}:/home/ploi/{domain}/public/
```

Notes:
- macOS ships an old rsync (2.6.9) — it doesn't understand `--info=progress2`, use `--progress`.
- `node_modules/` exclusion matches at any depth (no leading `/`), so it also catches nested
  ones (`themes/*/node_modules`, `adminXXX/themes/*/node_modules` — PS9 admin themes can have
  their own).
- Leave `vendor/` **in** the sync — including it skips a `composer install` on a server that
  may not have the exact same Composer/extension setup as local. It's usually 100-250MB, cheap
  compared to running composer resolution on a machine you don't fully control.
- Verify after: `ssh ... "ls {webroot}/index.php"` and check total size roughly matches local
  size minus the excluded dirs.

## Step 5 — Migrate the database

**`lando mysqldump` does not exist** — the command is `lando db-export`, and its target path
must be inside the project directory (the only thing bind-mounted into the container); an
absolute path like `/tmp/foo.sql` outside the project fails silently with "No such file or
directory" from inside the container, not from your shell.

```bash
cd {local_project_dir} && lando db-export dump.sql
```
(Ignore a stray `.gz`-suffixed output name from Lando's exporter — check with `file dump.sql*`
whether it actually gzipped; sometimes it names it `.gz` without compressing.)

Upload and import:
```bash
scp -i ~/.ssh/{ssh_key} dump.sql ploi@{server_ip}:/home/ploi/dump.sql
ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} \
  "mysql -h127.0.0.1 -u{db_user} -p'{db_pass}' {db_name} < /home/ploi/dump.sql"
```

**Use `-h127.0.0.1` (TCP), not a bare `mysql -u... -p...` (socket).** On Ploi's MySQL setup the
per-site DB user is only granted for TCP/`127.0.0.1`, not the local socket — the exact same
credentials that work with `-h127.0.0.1` fail with a generic "Access denied" over the socket,
which reads exactly like a wrong password. If you hit "Access denied" with credentials you're
sure are correct, try `-h127.0.0.1` before assuming the user made a typo.

**Collation gotcha.** Local Lando MySQL (8.0.x) or MariaDB dumps can use newer collations
(`utf8mb4_uca1400_ai_ci`, `utf8mb3_uca1400_ai_ci`) the server's MySQL version doesn't recognize:
```
ERROR 1273 (HY000) at line N: Unknown collation: 'utf8mb4_uca1400_ai_ci'
```
Fix by patching the dump, not by downgrading anything:
```bash
ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "
  sed -i 's/utf8mb4_uca1400_ai_ci/utf8mb4_unicode_ci/g; s/utf8mb3_uca1400_ai_ci/utf8mb3_general_ci/g' /home/ploi/dump.sql
"
```
Re-run the import after patching. Check for *both* variants — a dump can carry only one and
still fail again on the second a few hundred lines later if you only patch the first.

Clean up the dump file from both ends once the import is verified.

## Step 6 — Point parameters.php at the real DB

```bash
ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "
sed -i \"s/'database_host' => '[^']*',/'database_host' => '127.0.0.1',/\" {webroot}/app/config/parameters.php
sed -i \"s/'database_name' => '[^']*',/'database_name' => '{db_name}',/\" {webroot}/app/config/parameters.php
sed -i \"s/'database_user' => '[^']*',/'database_user' => '{db_user}',/\" {webroot}/app/config/parameters.php
sed -i \"s/'database_password' => '[^']*',/'database_password' => '{db_pass}',/\" {webroot}/app/config/parameters.php
"
```

## Step 7 — Fix the shop domain in the database

Without this, PrestaShop redirects every request back to the local `.lndo.site` domain baked
into the imported DB — the site loads then immediately bounces to a domain that doesn't
resolve outside your Mac.

Two separate places store the domain — both must be updated:
```sql
UPDATE ps_shop_url SET domain = '{domain}', domain_ssl = '{domain}' WHERE id_shop_url = 1;
UPDATE ps_configuration SET value = '{domain}' WHERE name IN ('PS_SHOP_DOMAIN', 'PS_SHOP_DOMAIN_SSL');
```
Run via `mysql -h127.0.0.1 -u{db_user} -p'{db_pass}' {db_name} -e "..."`. Clear
`{webroot}/var/cache/*` afterwards — PrestaShop caches resolved URLs.

## Step 8 — Strip absolute local-domain URLs from page-builder content

> **Theme-specific — check which theme/builder the project actually uses before assuming
> this applies.** Everything below (table names, column names, file paths) is exactly how
> **ElementFlow's `stsitebuilder` module** (an Elementor-based page builder) happens to store
> and export its content — that's the theme both reference projects (gaudi, llobet) used, not
> a general PrestaShop behavior. A different theme means different specifics:
> - **Panda (SunnyToo)** and its Easy Builder (`steasybuilder`) store layout differently —
>   check its own DB tables/config before assuming `ps_st_site_builder*` exists at all.
> - **Classic theme with no page builder** (plain Smarty/Bootstrap, hooks-only customization)
>   won't have this problem at all — a hardcoded local domain would only show up in actual CMS
>   page content (`ps_cms_lang.content`) or a module's own config, both of which the broad
>   text-column scan in Step 7's troubleshooting note already covers.
>
> The *general* lesson to carry forward regardless of theme is: **any visual/page builder that
> lets you paste absolute image or link URLs into a WYSIWYG editor can bake the local dev
> domain into stored content**, and that content might live in a DB column, a builder-specific
> cache, or (Step 8b) a statically-exported file. Confirm which theme/builder a *new* project
> uses first (`ls themes/`, check `composer.json`/`config.xml` for `stsitebuilder` vs
> `steasybuilder` vs neither), then adapt the table/path names below — don't blindly run
> ElementFlow-specific SQL against a Panda or Classic-theme project.

If the theme uses ElementFlow/`stsitebuilder`, its block content is JSON stored in
`ps_st_site_builder_postmeta.meta_value` (key `_elementor_data`), and it can contain
**absolute** URLs baked in at edit time (`https://project.lndo.site/img/cms/...`). These don't
get fixed by Step 7 — they're literal strings inside the JSON blob, not looked up via the shop
domain config.

Find them:
```sql
SELECT meta_id, post_id FROM ps_st_site_builder_postmeta WHERE meta_value LIKE '%lndo.site%';
```
Strip the absolute prefix so the URL becomes root-relative (works regardless of domain, which
is what you actually want for content that might move between staging/prod later):
```sql
UPDATE ps_st_site_builder_postmeta
SET meta_value = REPLACE(REPLACE(meta_value,
  'https://{local-lando-domain}', ''),
  'http://{local-lando-domain}', '')
WHERE meta_value LIKE '%{local-lando-domain}%';
```
Check both `http://` and `https://` variants. This is theme/builder-specific — if the project
uses a different page builder, check wherever it stores block content the same way (look for
`longtext`/JSON columns with a name suggesting builder data).

**If images/menu links still show the old domain after this AND after clearing `var/cache/*`,
the content isn't in the DB at all — check for statically-exported template files next
(Step 8b).** Don't assume a DB fix that returns 0 remaining rows means the job is done; a
`grep` against the live rendered page for the old domain is the only proof that actually
matters, and DB + cache being clean while the page still shows it is itself the signal to keep
looking rather than a contradiction to explain away.

## Step 8b — Check for statically-exported builder templates

> Also ElementFlow/`stsitebuilder`-specific (same caveat as Step 8) — the path
> `modules/stsitebuilder/views/templates/front/template/` won't exist on a Panda or
> Classic-theme project. If a different builder is in play, look for whatever its equivalent
> "compiled/exported template" location is (check its module's `views/templates/` tree for
> auto-generated `.tpl` files with builder-looking names, distinct from the theme's own
> hand-written templates) — the general pattern (a builder exporting rendered content to a
> real file instead of re-reading the DB every request) is common across Elementor-style
> builders even when the exact path differs.

`stsitebuilder` (and similar Elementor-style builders) can export a block's content to a
**real `.tpl` file on disk** the first time it's saved/rendered — not a cache, not
regenerated from the DB automatically, just a plain file checked in alongside the module,
under something like `modules/stsitebuilder/views/templates/front/template/*.tpl`. If that
export happened locally before the DB content was ever cleaned up, the exported file carries
the old absolute domain baked in **permanently**, independent of the DB row it was exported
from. Clearing `var/cache/*` (Symfony) or Smarty's own `cache/smarty/{cache,compile}/*`
(different directory, PS9's dev-mode Smarty output lives at `var/cache/dev/smarty/...` and
gets regenerated on every request in dev mode, so it always looks "fresh" without ever
becoming *correct*) does nothing for these, because they aren't cache — they're source.

Find them and check both the absolute (`https://olddomain/...`) and protocol-relative
(`//olddomain/...`) forms — builders commonly emit the latter for asset URLs:
```bash
ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "
  DIR={webroot}/modules/stsitebuilder/views/templates/front/template
  grep -rl '{local-lando-domain}' \"\$DIR\"
"
```
Fix in place (root-relative is what you want, same reasoning as Step 8):
```bash
ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "
  DIR={webroot}/modules/stsitebuilder/views/templates/front/template
  grep -rl '{local-lando-domain}' \"\$DIR\" | xargs sed -i \
    's#https\\?://{local-lando-domain}##g; s#//{local-lando-domain}#/#g'
  rm -rf {webroot}/var/cache/*
"
```
Then re-check the live page (`curl -s https://{domain}/ | grep -o '{local-lando-domain}' | wc -l`
should be `0`) — this is a case where fixing the DB and clearing every cache directory you can
think of will all report success while the actual page keeps showing the bug, because none of
that touches the real culprit.

## Step 9 — Fix the memory_limit fatal in the backoffice

Symptom in the admin panel:
```
Allowed memory size of 134217728 bytes exhausted (tried to allocate ...)
```
This is **not** a sign the project needs more resources long-term — it's Symfony's
Dependency Injection container compiling itself for the very first time (no cached container
exists yet after a raw file copy), which is a one-time heavy operation, hitting PHP-FPM's
default 128M `memory_limit`.

Two independent fixes, do both:

**1. Raise the web memory_limit** without needing root. `/etc/php/*/fpm/php.ini` is
root-owned and Ploi SSH users commonly don't have passwordless sudo — but PHP-FPM honors a
`.user.ini` file dropped at the site's webroot with no special privileges needed:
```bash
ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "echo 'memory_limit = 512M' > {webroot}/.user.ini"
```
Takes effect automatically within PHP's `user_ini.cache_ttl` (default 5 min), no service
restart needed.

**2. Pre-compile the container via CLI instead of a web request** — PHP-CLI has no
memory_limit restriction (`php -i | grep memory_limit` shows `-1`), so doing the compile here
is both faster and avoids ever hitting the fatal at all:
```bash
ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "cd {webroot} && php bin/console cache:clear --env=prod --no-warmup && php bin/console cache:warmup --env=prod"
```
Do this on every fresh deploy where `var/cache/` was excluded from the rsync (Step 4) or
wiped — don't rely on the `.user.ini` bump alone to survive the first request gracefully.

## Step 10 — Add the PrestaShop-aware nginx config (the big one)

This is the fix for the most confusing class of symptom this skill exists to document: a
`/adminXXX/login` request 404s and the response body is the **storefront theme's own 404
page** — looks like a completely unrelated front-end bug, but it's nginx never routing the
request to the admin app's `index.php` at all.

### How to tell it's this, not something else

Diagnostic that removes all doubt (browser caching, CDNs, and "did it actually run" guesses
all confound naive testing — this doesn't):
```bash
# 1. Note the current line count of the app's own log
ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "wc -l < {webroot}/var/logs/dev-*.log"
# 2. Hit the failing URL
curl -s -o /dev/null "https://{domain}/{admin_dir}/login?_token="
# 3. Check the line count again
ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "wc -l < {webroot}/var/logs/dev-*.log"
```
If the count didn't move, the request never reached PrestaShop's Symfony admin kernel at
all — it's purely an nginx routing gap, full stop, regardless of what the rendered page looks
like. Don't go chasing PHP version compatibility, cache staleness, or app-level bugs based on
the rendered HTML alone; confirm with this log-growth check first. (A bare `/{admin_dir}/`
request, with no further path, usually *does* reach the admin kernel already, because nginx's
`index index.php;` directive naturally serves the directory's own `index.php` for a literal
directory match — this is what makes the bug easy to half-diagnose and then get stuck: the
redirect to `/login` fires, then the *next* request silently falls through.)

### Where to add it

Per-site, via Ploi panel → Site → Manage → NGINX → **Server** section (right-hand column) →
**+** → new file, e.g. `prestashop.conf`. NOT the main site config file — the "Server" include
folder is meant for exactly this kind of addition and is included *before* the site's own
`location ~ \.php$` block, so its regex locations win the match for admin/image paths without
needing to touch or understand the rest of the generated file.

**Don't also hand-add a single `location ^~ /{admin_dir}/ { ... }` block directly into the
site's main config** as an alternative/shortcut — it's tempting since it looks simpler, but it
conflicts with the more specific regex-based routing below (which needs to distinguish
`{admin_dir}/index.php`, non-index PHP files under `{admin_dir}` like `filemanager/dialog.php`,
and clean asset/route paths from each other) and produces worse failure modes that are harder
to debug than not having it at all.

### The file

Adjust the `fastcgi_pass` socket to the site's actual PHP version (check
`{webroot}/.user.ini`'s sibling config or the site's PHP version in the panel).

```nginx
# === PrestaShop nginx rules (admin routing + SEO image URLs) ===
# Ploi's generic template has no idea PrestaShop needs any of this — see
# ploi-prestashop-deploy skill for why.

fastcgi_read_timeout 300;
fastcgi_buffers 32 32k;
fastcgi_buffer_size 64k;

add_header Referrer-Policy        "strict-origin-when-cross-origin" always;
add_header Strict-Transport-Security "max-age=31536000; includeSubDomains" always;
add_header Permissions-Policy    "geolocation=(), microphone=(), camera=()" always;
fastcgi_hide_header X-Powered-By;

# PS9 bug: some module action links render without a leading slash, producing
# a nested duplicate admin path segment. Collapse it back to the real index.php.
rewrite ^/(admin[-_\w]*)/[^/]+/[^/]+.*/index\.php$ /$1/index.php last;

# Product image SEO URL → real disk path (PS splits by digit: id 191 -> img/p/1/9/1/191-x.ext)
rewrite ^/([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$ /img/p/$1/$1$2.$3 last;
rewrite ^/([0-9])([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$ /img/p/$1/$2/$1$2$3.$4 last;
rewrite ^/([0-9])([0-9])([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$ /img/p/$1/$2/$3/$1$2$3$4.$5 last;
rewrite ^/([0-9])([0-9])([0-9])([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$ /img/p/$1/$2/$3/$4/$1$2$3$4$5.$6 last;
rewrite ^/([0-9])([0-9])([0-9])([0-9])([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$ /img/p/$1/$2/$3/$4/$5/$1$2$3$4$5$6.$7 last;
rewrite ^/([0-9])([0-9])([0-9])([0-9])([0-9])([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$ /img/p/$1/$2/$3/$4/$5/$6/$1$2$3$4$5$6$7.$8 last;

# Category images: /c/35-name[/seo-subpath].ext -> /img/c/35-name.ext
# The optional non-capturing group discards a PS9 SEO subfolder before the extension.
rewrite ^/c/([0-9]+\-[\w-]+)(?:/.+)?\.(jpg|jpeg|png|gif|webp)$ /img/c/$1.$2 last;

rewrite ^/(\w+)-sitemap\.xml$ /sitemap.php?lang=$1 last;

location ~* ^/(config|app|bin|src|var|vendor)(/|$) { deny all; }
location ~* /\.env { deny all; }
location ~* ^/(composer\.(json|lock)|package(-lock)?\.json|yarn\.lock|Makefile)$ { deny all; }
location ~* \.(bak|sql|log|twig)$ { deny all; }
location ~* ^/(upload|img)/.*\.php[0-9]?$ { deny all; }

# Admin: any .php file other than index.php (e.g. filemanager/dialog.php).
# Must come before the clean-URLs block below, or nginx serves the file as a
# static download via try_files instead of executing it.
location ~* ^/(admin[-_\w]*)/(?!index\.php).+\.php(/|$) {
    fastcgi_pass unix:/run/php/php{PHP_VERSION}-fpm.sock;
    fastcgi_split_path_info ^(.+\.php)(/.*)$;
    fastcgi_index index.php;
    fastcgi_param SCRIPT_FILENAME $realpath_root$fastcgi_script_name;
    fastcgi_param DOCUMENT_ROOT $realpath_root;
    include fastcgi_params;
}

# Admin clean URLs (Symfony routes): serve static assets directly, else route to index.php
location ~* ^/(admin[-_\w]*)/(?!index\.php)(.+)$ {
    try_files $uri $uri/ /$1/index.php/$2$is_args$args;
}

# Admin index.php (PS9 Symfony routing entry point)
location ~* ^/(admin[-_\w]*)/index\.php(/|$) {
    fastcgi_pass unix:/run/php/php{PHP_VERSION}-fpm.sock;
    fastcgi_split_path_info ^(.+\.php)(/.*)$;
    fastcgi_index index.php;
    fastcgi_param SCRIPT_FILENAME $realpath_root$fastcgi_script_name;
    fastcgi_param DOCUMENT_ROOT $realpath_root;
    include fastcgi_params;
}

location ~* \.(gif|jpe?g|png|ico|svg|webp|css|js|woff2?|ttf|eot|otf)$ {
    try_files $uri /index.php?$query_string;
    expires 1M;
    add_header Cache-Control "public, immutable";
    add_header X-Content-Type-Options "nosniff" always;
    access_log off;
    log_not_found off;
}
```

`admin[-_\w]*` matches any admin folder name Ploi/PrestaShop generates (random hash or a
custom rename like `baupanel`) automatically — you don't need to hardcode `{admin_dir}`
anywhere in this file. If the project uses a completely different naming scheme that this
regex misses, add it as an alternation: `(admin[-_\w]*|customname)`.

**Optional: rate limiting.** A reference version of this file may include
`limit_req zone=prestashop_admin burst=30 nodelay;` on the admin `index.php` block. That
requires a `limit_req_zone` defined server-wide, e.g. in `/etc/nginx/conf.d/rate-limits.conf`:
```nginx
limit_req_zone $binary_remote_addr zone=prestashop_admin:10m rate=120r/m;
```
This is root-owned config outside any single site's control — if the SSH account has no
sudo, either ask whoever manages the server to add it once (it benefits every PrestaShop site
on that box), or just drop the `limit_req` line from the site's `prestashop.conf`. Everything
else in the file works fine without it.

### robots.conf (staging/dev domains)

For any domain that isn't the real production one (a `pre-testing.com`, `.dev`, staging
subdomain, etc.), also add this as a separate file in the same Server section, so it doesn't
get indexed:
```nginx
add_header X-Robots-Tag "noindex, nofollow, nosnippet, noarchive";
```

### Verify the fix

Same log-growth check as above, now expecting the count to move:
```bash
LINES_BEFORE=$(ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "wc -l < {webroot}/var/logs/dev-*.log")
curl -s -o /dev/null "https://{domain}/{admin_dir}/login?_token="
LINES_AFTER=$(ssh -i ~/.ssh/{ssh_key} ploi@{server_ip} "wc -l < {webroot}/var/logs/dev-*.log")
```
`LINES_AFTER` > `LINES_BEFORE` confirms the admin kernel actually ran this time.

## Step 11 — DNS and SSL

DNS A record → `{server_ip}`, before requesting SSL (Let's Encrypt validates over HTTP/DNS —
if it doesn't resolve yet, the request just fails; wait for propagation rather than reaching
for "force request / skip DNS verification").

Ploi panel → Site → SSL → type = Let's Encrypt, domain = `{domain}`. No DNS provider option
needed unless using DNS-01 challenge for a wildcard.

## Step 12 — Final verification

```bash
curl -s -o /dev/null -w "front=%{http_code}\n" https://{domain}/
curl -s -L -o /dev/null -w "admin=%{http_code} -> %{url_effective}\n" https://{domain}/{admin_dir}/
curl -s -o /dev/null -w "product image=%{http_code}\n" "https://{domain}/{some-product-id}-home_default/whatever.jpg"
```
Expected: front `200`, admin final `200` at `.../login?_token=`, image `200`.

---

## Troubleshooting

Match the symptom, then jump to the linked step — don't re-derive the cause from scratch each
time, it's always one of these on a fresh Ploi+PrestaShop deploy.

| Symptom | Cause | Fix |
|---|---|---|
| Admin `/login` 404s, page body looks like the **storefront** theme | nginx has no PrestaShop-aware admin routing | [Step 10](#step-10--add-the-prestashop-aware-nginx-config-the-big-one) |
| Product/category images 404 despite the file existing on disk at `/img/p/.../...jpg` | nginx has no SEO-image-URL rewrite | [Step 10](#step-10--add-the-prestashop-aware-nginx-config-the-big-one) |
| `Allowed memory size of 134217728 bytes exhausted` in backoffice | first-time Symfony DI container compile hitting 128M default | [Step 9](#step-9--fix-the-memory_limit-fatal-in-the-backoffice) |
| Site loads then immediately redirects to a `.lndo.site` domain | `ps_shop_url`/`ps_configuration` still hold the local dev domain from the imported dump | [Step 7](#step-7--fix-the-shop-domain-in-the-database) |
| Some images/CMS blocks 404 even after Step 10, pointing at a `.lndo.site` URL in their `src` | absolute domain baked into page-builder JSON content | [Step 8](#step-8--strip-absolute-local-domain-urls-from-page-builder-content) |
| Menu links / images still show the old `.lndo.site` domain even though the DB check returns 0 rows and `var/cache/*` is empty | content was statically exported to a real `.tpl` file on disk, not read from the DB at request time | [Step 8b](#step-8b--check-for-statically-exported-builder-templates) |
| DB import fails with `Unknown collation: 'utf8mb...uca1400_ai_ci'` | local dump used a newer MySQL/MariaDB collation than the server supports | [Step 5](#step-5--migrate-the-database) |
| `Access denied for user` on DB import/queries despite correct credentials | tried the socket instead of TCP (`-h127.0.0.1`) | [Step 5](#step-5--migrate-the-database) |
| Not sure if a request even reached the app or died at nginx | — | the log-growth check in [Step 10](#how-to-tell-its-this-not-something-else) works for any "did this request even run" question, not just admin login |

### General debugging principle for this stack

When something on a fresh Ploi PrestaShop site behaves unexpectedly and the cause isn't
obvious, **compare against a known-working PrestaShop site's Server includes list** (Ploi
panel → that site → Manage → NGINX → right column shows "Before"/"Server"/"After" files) rather
than guessing at nginx internals from scratch. A working reference site's file list (e.g. it
has a `prestashop.conf` and `robots.conf` that a broken one doesn't) is faster and more
reliable evidence than reasoning about nginx `location` priority rules in the abstract — those
rules are easy to reason about incorrectly (prefix vs regex precedence, `error_page`
inheritance across locations, implicit directory-index behavior) even when each individual
rule is well understood.
