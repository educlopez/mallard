---
name: lando-setup
description: >
  Sets up a PrestaShop 8/9 project locally with Lando, or refreshes its
  database from the server. Full setup configures Lando, syncs app/config, picks a
  free port, imports the DB and fixes shop URLs; DB refresh mode only re-dumps and
  re-imports the database. Use when the user says "monta el proyecto", "setup lando",
  "prepara en local", "levanta el lando de X", "configura X en local", "actualízame
  la DB", "refresca la DB de X", "bájame la DB de X", "sync DB", or "quiero la DB
  actualizada de X".
version: "0.3.0"
metadata:
  author: Eduardo Calvo
---

# lando-setup

## Trigger

**Full setup** (project not yet configured locally):
"monta el proyecto", "setup lando", "prepara en local", "levanta el lando de X", "configura X en local"
→ Run all steps in order.

**DB refresh only** (Lando already running, just need fresh DB):
"actualízame la DB", "refresca la DB de X", "bájame la DB de X", "sync DB", "quiero la DB actualizada de X"
→ Skip to [DB Refresh Mode](#db-refresh-mode) below.

## ⚠️ Invoke this skill WITHOUT arguments

The nginx template in Step 9 is full of `$1`…`$8` capture references. When this skill is
invoked with an `args` string, the harness expands those as positional parameters — the
image rewrite rules arrive mangled (`$1` becomes the first word of your args, `$5` becomes
a URL, etc.) and any vhost written from them will 404 every product image.

Invoke as `/lando-setup` with **no args** and state the project, server and options in the
conversation instead. If you already invoked it with args, do NOT copy the rewrite rules
from the loaded skill text — read `SKILL.md` from disk, or copy an existing working
`.lando/nginx-site.conf` from another PS9 project.

## Placeholders
- `{workspace}` — the directory under `~/developer/` where your projects live (e.g. your company or team name). Set it once; every path below uses `~/developer/{workspace}/{project-name}/`.
- `{ssh_key}` — the SSH private key configured for the servers (e.g. `id_rsa_company`). Used explicitly on every SSH/rsync command.

## Context
- Projects live at ~/developer/{workspace}/{project-name}/
- SSH hosts in ~/.ssh/config follow pattern: {alias}-prod, {alias}-pre
- All projects are PrestaShop 8 or 9
- DB credentials in app/config/parameters.php (on the SERVER)
- Existing port range in use: 33290–33409 (pick next free one)

## SSH Authentication

### SSH Keys
- `~/.ssh/{ssh_key}` — your primary company key (RSA). Replace `{ssh_key}` with your configured key name (e.g. `id_rsa_company`).
- `~/.ssh/{ssh_key_ed25519}` — Ed25519, your company SSH key — modern alternative.

Always use `-i ~/.ssh/{ssh_key}` explicitly on all SSH/rsync commands to avoid key ambiguity.

### Auth modes

**Key-based (most projects):** Your company public key is on the server. No password needed.

**Password-based (some projects):** Key not on server. Use `sshpass` via stdin to avoid password in process list:
```bash
export SSHPASS='{password}'
sshpass -e ssh -i ~/.ssh/{ssh_key} {user}@{ip} "command"
sshpass -e ssh -i ~/.ssh/{ssh_key} {user}@{ip} "mysqldump ..." > /tmp/dump.sql
unset SSHPASS
```
Install if missing: `brew install sshpass`
**Never put the password directly in the command string** — use `-e` (reads from `$SSHPASS` env var) instead of `-p`.

**Jump host / double hop (some internal servers):** the target is only reachable through a
bastion, AND the key that authenticates the second hop lives **on the bastion**, not on the
Mac. `ProxyJump` does NOT work here — it forwards the connection, not the bastion's keys.
Detect it: `ssh -J {bastion} {user}@{target} "echo OK"` fails with `Permission denied
(publickey)` while `ssh {bastion} "ssh {user}@{target} 'echo OK'"` succeeds.

Two ways forward — **ask the user which one, never assume**:
- Append their public key to `{user}@{target}:~/.ssh/authorized_keys` (via the bastion), then
  `ProxyJump` works and rsync/mysqldump run directly. Modifies the target server.
- Keep everything nested — no server change. Works fine, just more verbose:
```bash
# remote command
ssh {bastion} "ssh {user}@{target} 'command'"
# DB dump
ssh {bastion} "ssh {user}@{target} 'mysqldump --no-tablespaces --single-transaction {DB_NAME}'" > dump.sql
# directory transfer — tar over pipe, no staging on the bastion's disk (rsync can't traverse two hops)
ssh {bastion} "ssh {user}@{target} 'cd {REMOTE_ROOT} && tar czf - {dir}'" | tar xzf -
```
When nested, substitute this pattern for every `rsync`/`ssh` command in Steps 3–7 below.

**Adding new host to ~/.ssh/config** (do this when user gives raw `user@IP`):
```
Host {alias}-{env}
  HostName {IP}
  User {user}
  IdentityFile ~/.ssh/{ssh_key}
```
Ask user if they want to save it for future use. If yes, append to ~/.ssh/config.

## Step 1 — Confirm inputs

Ask user (if not already provided):
- Project folder name (e.g. `ps8-pincolor`)
- SSH connection: existing alias OR `user@IP`
- Environment: **prod** or **pre** — ALWAYS ask, never assume
- Auth method: key (default) or password? If password, ask for it now (will store in env var, not shell string).
- Web server: **nginx** (lemp) or **Apache** (lamp)?
  - Default: **nginx** for PS9 on Ploi servers; **Apache** for PS8 or legacy setups.
  - If unsure: check server — Ploi = nginx, cPanel/Plesk = Apache.

Derive:
- `PROJECT_DIR=~/developer/{workspace}/{project-name}`
- `SSH_TARGET={alias}-{env}` or `{user}@{ip}`
- `LANDO_NAME={project-name}` (same as folder)
- `LOCAL_URL={lando-name}.lndo.site`
- `LANDO_RECIPE=lemp` (nginx) or `lamp` (Apache)

**Test SSH connectivity before proceeding:**
```bash
ssh -i ~/.ssh/{ssh_key} -o ConnectTimeout=5 -o BatchMode=yes {SSH_TARGET} "echo OK" 2>&1
```
If output is not `OK`, stop and diagnose auth failure with user before continuing.

## Step 2 — Detect PrestaShop version

```bash
grep -A2 '"prestashop/prestashop"' ~/developer/{workspace}/{project-name}/composer.json
```

Or check:
```bash
grep "_PS_VERSION_" ~/developer/{workspace}/{project-name}/config/defines.inc.php 2>/dev/null | head -1
```

- Version 8.x → PS8, PHP 8.1
- Version 9.x → PS9, PHP 8.3

If version is not clearly readable from files, ask user: "¿PS8 o PS9?" — do not guess from folder name.

## Step 3 — Find remote webroot

```bash
ssh -i ~/.ssh/{ssh_key} {SSH_TARGET} "find /var/www -name 'parameters.php' -path '*/app/config/*' 2>/dev/null"
```

If multiple results appear, show them to user and ask which is the correct one.
Derive `REMOTE_ROOT` by stripping `/app/config/parameters.php` from the chosen path.
Example: `/var/www/html/app/config/parameters.php` → `REMOTE_ROOT=/var/www/html`

## Step 4 — Sync missing files from server

These files/folders are NOT in git and must be copied from the server via rsync.

### 4a — vendor/ (root, required)
```bash
rsync -avz --progress -e "ssh -i ~/.ssh/{ssh_key}" \
  {SSH_TARGET}:{REMOTE_ROOT}/vendor/ ~/developer/{workspace}/{project}/vendor/
```

### 4b — translations/ (required for PS9)
Without this folder, PS9 throws DirectoryNotFoundException on boot.
```bash
rsync -avz --progress -e "ssh -i ~/.ssh/{ssh_key}" \
  {SSH_TARGET}:{REMOTE_ROOT}/translations/ ~/developer/{workspace}/{project}/translations/
```

### 4b-bis — .env, localization/, mails/ (required for PS9)
Also missing from git and easy to overlook:
- `.env` — without it `bin/console` dies with
  `Symfony\Component\Dotenv\Exception\PathException: Unable to read the "/app/.env" environment file`,
  so Step 14 (cache clear) can never run.
- `localization/` — `LocalizationWarmer` warns `file_get_contents(/app/localization/default.xml): Failed to open stream`.
- `mails/` — transactional email templates.

```bash
rsync -avz --progress -e "ssh -i ~/.ssh/{ssh_key}" \
  {SSH_TARGET}:{REMOTE_ROOT}/.env ~/developer/{workspace}/{project}/.env
rsync -avz --progress -e "ssh -i ~/.ssh/{ssh_key}" \
  {SSH_TARGET}:{REMOTE_ROOT}/localization/ ~/developer/{workspace}/{project}/localization/
rsync -avz --progress -e "ssh -i ~/.ssh/{ssh_key}" \
  {SSH_TARGET}:{REMOTE_ROOT}/mails/ ~/developer/{workspace}/{project}/mails/
```

### 4c — app/config/ files (required)
If `~/developer/{workspace}/{project}/app/config/parameters.php` already exists locally, warn user:
> "app/config/ exists locally. Sync from server will overwrite parameters.php. Continue? (y/n)"

```bash
rsync -avz --progress -e "ssh -i ~/.ssh/{ssh_key}" \
  {SSH_TARGET}:{REMOTE_ROOT}/app/config/ ~/developer/{workspace}/{project}/app/config/
```

### 4d — modules/ (required)
Sync the full modules directory. Modules not in git (theme modules, standard PS modules) won't exist locally otherwise.
```bash
rsync -avz --progress -e "ssh -i ~/.ssh/{ssh_key}" \
  {SSH_TARGET}:{REMOTE_ROOT}/modules/ ~/developer/{workspace}/{project}/modules/
```
This is simpler and safer than syncing individual vendor/ subdirs — covers theme modules (e.g. Panda) and any PS module not committed to git.

### 4e — img/ folder (OPTIONAL — ask user first)
The img/ folder can be several GB. Ask user:
> "¿Descargo la carpeta img/ del servidor o usamos placeholder de imágenes? (recomendado: placeholder)"

- If placeholder → trigger `lando-img-placeholder` skill after setup (it handles both lamp and lemp automatically)
- If download → `rsync -avz --progress -e "ssh -i ~/.ssh/{ssh_key}" {SSH_TARGET}:{REMOTE_ROOT}/img/ ~/developer/{workspace}/{project}/img/`

**Note for nginx (lemp) projects**: placeholder rules are built into `.lando/nginx-site.conf` — no separate Apache file needed. The `lando-img-placeholder` skill detects the recipe and acts accordingly.

**Regardless of choice (placeholder or download), always create required subdirs:**
```bash
mkdir -p ~/developer/{workspace}/{project}/img/{c,p,m,st,cms,l,tmp,os,s,co,su}
```
Without these, `ps_mainmenu` crashes calling `scandir('/app/img/c/')` on boot.

## Step 5 — Update parameters.php with Lando credentials

After syncing app/config/ from server, parameters.php points to the remote DB → instant 500.
Update it with local Lando values.

**Capture the remote credentials FIRST** — the sed below destroys them, and Step 6/7 need them
to dump the server DB:
```bash
cp ~/developer/{workspace}/{project}/app/config/parameters.php /tmp/parameters.remote.php
grep -E "database_(name|user|password|prefix)" /tmp/parameters.remote.php
```

**The DB name/user/password is the RECIPE NAME, not always `lamp`:** recipe `lamp` → `lamp`,
recipe `lemp` → **`lemp`**. Getting this wrong gives `ERROR 1049 (42000): Unknown database 'lamp'`
and a 500 that looks like a broken install. Set `RECIPE` to whatever you chose in Step 1:

```bash
RECIPE=lemp   # or lamp
sed -i '' \
  -e "s/'database_host' => '[^']*'/'database_host' => 'database'/" \
  -e "s/'database_name' => '[^']*'/'database_name' => '${RECIPE}'/" \
  -e "s/'database_user' => '[^']*'/'database_user' => '${RECIPE}'/" \
  -e "s/'database_password' => '[^']*'/'database_password' => '${RECIPE}'/" \
  -e "s/'mailer_host' => '[^']*'/'mailer_host' => 'sendmailhog'/" \
  ~/developer/{workspace}/{project}/app/config/parameters.php
```

Lando DB credentials: host=`database`, name/user/password = recipe name, mailer=`sendmailhog`.
Verify the file was updated correctly before proceeding. Confirm against the running container:
```bash
cd ~/developer/{workspace}/{project} && lando ssh -s database -c "mysql -uroot -e 'SHOW DATABASES'"
```

## Step 6 — Get remote DB credentials

Read the **pre-rewrite backup** saved in Step 5 — the live `app/config/parameters.php` now holds
the Lando credentials, not the server's:
```bash
grep -E "database_(name|user|password|prefix)" /tmp/parameters.remote.php
```

Extract:
- `database_name` → `DB_NAME`
- `database_user` → `DB_USER`
- `database_password` → `DB_PASS`
- `database_prefix` → `DB_PREFIX` (default `ps_` but MUST read actual value)

## Step 7 — Dump DB from server

Key-based:
```bash
ssh -i ~/.ssh/{ssh_key} {SSH_TARGET} \
  "mysqldump -h 127.0.0.1 -u'${DB_USER}' -p'${DB_PASS}' '${DB_NAME}'" > ~/developer/{workspace}/{project-name}/dump.sql
```

Password-based (SSH password):
```bash
export SSHPASS='{ssh_password}'
sshpass -e ssh -i ~/.ssh/{ssh_key} {user}@{ip} \
  "mysqldump -h 127.0.0.1 -u'${DB_USER}' -p'${DB_PASS}' '${DB_NAME}'" > ~/developer/{workspace}/{project-name}/dump.sql
unset SSHPASS
```

Verify dump:
```bash
wc -c ~/developer/{workspace}/{project-name}/dump.sql
```
If file is under 100KB, likely failed or empty — check for errors before continuing.

### 7b — Match the local DB engine to the server's

Check what the server actually runs BEFORE writing `.lando.yml`:
```bash
ssh -i ~/.ssh/{ssh_key} {SSH_TARGET} "mysql -e 'SELECT VERSION()'"
```

A MariaDB 11.4+ server emits collations MySQL 8 cannot parse, and the import dies at line ~26:
```
ERROR 1273 (HY000) at line 26: Unknown collation: 'utf8mb4_uca1400_ai_ci'
```
Fix: set the local engine to the same family/version — e.g. `mariadb:11.8` (Lando 3.26.4 ships
MariaDB 10.11 and 11.0–11.8). Do NOT try to sed the collations out of the dump.

Changing the engine after a `lando start` needs `lando destroy -y` first — the existing volume
was initialised by the other engine.

## Step 8 — Pick available port

```bash
grep -r "portforward" ~/developer/{workspace}/*/.lando.yml 2>/dev/null
```

Pick the next free port above the highest found (start at 33290 if none found).

## Step 9 — Generate .lando.yml

Create `~/developer/{workspace}/{project-name}/.lando.yml`:

**If this project was previously set up with a different MySQL version**, the old Docker volume may be corrupt (error: "Upgrade is not supported after a crash or shutdown with innodb_fast_shutdown = 2"). Fix: run `lando destroy -y` before `lando start`. This deletes the local DB volume — safe since we're importing a fresh dump anyway.

First, create the PHP config override (prevents `max_execution_time` errors during theme install/parsing):
```bash
mkdir -p ~/developer/{workspace}/{project-name}/.lando
```

Create `~/developer/{workspace}/{project-name}/.lando/php.ini`:
```ini
max_execution_time = 300
max_input_time = 300
memory_limit = 512M
post_max_size = 64M
upload_max_filesize = 64M
```

**For PS8 (PHP 8.1):**
```yaml
name: {lando-name}
recipe: lamp
config:
  php: "8.1"
  webroot: ./
  database: mysql
  xdebug: false
services:
  appserver:
    config:
      php: .lando/php.ini
  database:
    portforward: {next-free-port}
    type: mysql
  mail:
    type: mailhog
    hogfrom:
      - appserver
```

**For PS9 with Apache (PHP 8.3):**
```yaml
name: {lando-name}
recipe: lamp
config:
  php: "8.3"
  webroot: ./
  database: mysql
  xdebug: false
services:
  appserver:
    config:
      php: .lando/php.ini
  database:
    portforward: {next-free-port}
    type: mysql
  mail:
    type: mailhog
    hogfrom:
      - appserver
```

**For PS9 with nginx (PHP 8.3) — default for Ploi servers:**

First create `.lando/nginx-site.conf`:
```nginx
server {
    listen 80;
    root /app;
    index index.php index.html;
    client_max_body_size 128M;

    charset utf-8;

    # Upstream PHP-FPM — SIEMPRE así, nunca "fastcgi_pass appserver:9000":
    # 1) El alias "appserver" existe en la red compartida del proxy de Lando para
    #    CADA app corriendo → Docker DNS puede devolver el appserver de OTRO
    #    proyecto (502 connection refused). "fpm" solo existe en la red privada
    #    del proyecto, no puede colisionar.
    # 2) Variable + resolver fuerzan resolución DNS por petición. Con hostname
    #    literal, nginx resuelve UNA vez al arrancar y cachea la IP — si el
    #    appserver arranca después o cambia de IP en un restart, nginx apunta
    #    a una IP stale (502 hasta reiniciar nginx).
    resolver 127.0.0.11 valid=10s ipv6=off;
    set $fpm_upstream fpm:9000;

    # Placeholder — friendly URLs: /{id}-{type}_default/...
    location ~* ^/[0-9]+-[a-z]+_default.*\.(jpg|jpeg|png|gif|webp)$ {
        try_files $uri /img/placeholder-dev.jpg;
    }

    # Placeholder — classic PS paths: /img/p/, /img/c/, /img/m/, /img/st/
    location ~* ^/img/(p|c|m|st)/.*\.(jpg|jpeg|png|gif|webp)$ {
        try_files $uri /img/placeholder-dev.jpg;
    }

    # PS9 double-admin-path fix (legacy module redirect bug with Symfony router)
    rewrite ^/(admin[-_\w]*)/.+/admin[-_\w]*/index\.php$ /$1/index.php last;

    # Product image SEO URL → real disk path
    rewrite ^/([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$           /img/p/$1/$1$2.$3 last;
    rewrite ^/([0-9])([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$    /img/p/$1/$2/$1$2$3.$4 last;
    rewrite ^/([0-9])([0-9])([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$ /img/p/$1/$2/$3/$1$2$3$4.$5 last;
    rewrite ^/([0-9])([0-9])([0-9])([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$ /img/p/$1/$2/$3/$4/$1$2$3$4$5.$6 last;
    rewrite ^/([0-9])([0-9])([0-9])([0-9])([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$ /img/p/$1/$2/$3/$4/$5/$1$2$3$4$5$6.$7 last;
    rewrite ^/([0-9])([0-9])([0-9])([0-9])([0-9])([0-9])(\-[\w-]+)?/.+\.(jpg|jpeg|png|gif|webp)$ /img/p/$1/$2/$3/$4/$5/$6/$1$2$3$4$5$6$7.$8 last;

    # Category images
    rewrite ^/c/([0-9]+\-[\w-]+)\.(jpg|jpeg|png|gif|webp)$ /img/c/$1.$2 last;

    # Sitemap
    rewrite ^/(\w+)-sitemap\.xml$ /sitemap.php?lang=$1 last;

    # Block sensitive directories
    location ~* ^/(config|app|bin|src|var|vendor)(/|$) { deny all; }
    location ~* /\.env { deny all; }
    location ~* ^/(composer\.(json|lock)|package(-lock)?\.json|yarn\.lock|Makefile)$ { deny all; }
    location ~* \.(bak|sql|log|twig)$ { deny all; }
    location ~* ^/(upload|img)/.*\.php[0-9]?$ { deny all; }

    # Admin: non-index PHP files (filemanager/dialog.php, etc.)
    location ~* ^/admin[-_\w]*/(?!index\.php).+\.php(/|$) {
        fastcgi_pass $fpm_upstream;
        fastcgi_split_path_info ^(.+\.php)(/.*)$;
        fastcgi_buffers 32 32k;
        fastcgi_buffer_size 32k;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param DOCUMENT_ROOT $document_root;
        fastcgi_read_timeout 300;
        include fastcgi_params;
    }

    # Admin: clean URLs (Symfony routes) → index.php
    location ~* ^/(admin[-_\w]*)/(?!index\.php)(.+)$ {
        try_files $uri $uri/ /$1/index.php/$2$is_args$args;
    }

    # Admin: index.php (PS9 Symfony routing entry point)
    location ~* ^/admin[-_\w]*/index\.php(/|$) {
        fastcgi_pass $fpm_upstream;
        fastcgi_split_path_info ^(.+\.php)(/.*)$;
        fastcgi_buffers 32 32k;
        fastcgi_buffer_size 32k;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param DOCUMENT_ROOT $document_root;
        fastcgi_read_timeout 300;
        include fastcgi_params;
    }

    # PS9 admin-api
    location /admin-api {
        try_files $uri $uri/ /admin-api/index.php$is_args$args;
    }

    # Static assets
    location ~* \.(gif|jpe?g|png|ico|svg|webp|css|js|woff2?|ttf|eot|otf)$ {
        try_files $uri /index.php?$query_string;
        expires 1d;
        access_log off;
        log_not_found off;
    }

    # Front → PS router
    location / {
        try_files $uri $uri/ /index.php?$query_string;
    }

    # PHP via FPM
    location ~ \.php$ {
        try_files $uri /index.php =404;
        fastcgi_pass $fpm_upstream;
        fastcgi_split_path_info ^(.+\.php)(/.+)$;
        fastcgi_buffers 32 32k;
        fastcgi_buffer_size 32k;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        fastcgi_param DOCUMENT_ROOT $document_root;
        fastcgi_param HTTP_PROXY "";
        fastcgi_read_timeout 300;
        include fastcgi_params;
    }
}
```

Then create `.lando.yml` (`database:` / `type:` must match the server engine found in Step 7b):
```yaml
name: {lando-name}
recipe: lemp
config:
  php: "8.3"
  webroot: ./
  database: mariadb:11.8   # or mysql:8.0 — match the server (Step 7b)
  xdebug: false
services:
  appserver:
    config:
      vhosts: .lando/nginx-site.conf
      php: .lando/php.ini
  database:
    portforward: {next-free-port}
    type: mariadb:11.8   # keep in sync with config.database above
  mail:
    type: mailhog
    hogfrom:
      - appserver
```

## Step 10 — Ensure OrbStack is running

OrbStack is the Docker runtime. If it's not running, Lando will fail or open Docker Desktop (not configured).

Check if OrbStack is running:
```bash
pgrep -x "OrbStack" > /dev/null && echo "RUNNING" || echo "STOPPED"
```

If STOPPED, launch it and wait for Docker to be ready:
```bash
open -a OrbStack
```
Then poll until Docker responds (up to 30s):
```bash
for i in $(seq 1 15); do
  docker info > /dev/null 2>&1 && echo "Docker ready" && break
  echo "Waiting for OrbStack... ($i)"
  sleep 2
done
```
If Docker is still not ready after 30s, stop and ask user to check OrbStack manually.

## Step 11 — Start Lando

```bash
cd ~/developer/{workspace}/{project-name} && lando start
```

After start, verify PHP version matches expected:
```bash
cd ~/developer/{workspace}/{project-name} && lando php -v
```
If version doesn't match (e.g. shows 8.1 but expected 8.3), stop and fix `.lando.yml` before importing DB.

## Step 12 — Import DB

```bash
cd ~/developer/{workspace}/{project-name} && lando db-import dump.sql
```

After import, verify tables exist.

**`lando mysql -e "..."` silently returns nothing** in recent Lando — it is a passthrough
wrapper, not a real client invocation. Run queries inside the database service instead, and
pass multi-statement SQL through a file in the project root (the container mounts it at `/app`;
`source` is a client command and cannot be combined with other statements in a single `-e`):

```bash
cd ~/developer/{workspace}/{project-name}
lando ssh -s database -c "mysql -uroot ${RECIPE} -e 'SELECT COUNT(*) FROM ${DB_PREFIX}configuration'"
```
If this returns 0 or errors, DO NOT delete the dump file — report failure to user and stop.

## Step 13 — Fix URLs in DB

Use the actual `DB_PREFIX` read from the Step 5 backup, not hardcoded `ps_`.

**Set SSL to 1, not 0**, when the vhost sends `fastcgi_param HTTPS on` (the nginx template in
Step 9 does): PS then sees every request as secure while the DB says SSL is off, and redirects
https→http forever. Symptom: `curl` returns `302` pointing at the same page, and following
redirects fails with curl exit 47 (too many redirects). Lando serves a valid https cert, so
SSL on is also the accurate setting.

```bash
cd ~/developer/{workspace}/{project-name}
cat > urls.sql <<EOF
UPDATE ${DB_PREFIX}configuration SET value = '{lando-name}.lndo.site' WHERE name = 'PS_SHOP_DOMAIN';
UPDATE ${DB_PREFIX}configuration SET value = '{lando-name}.lndo.site' WHERE name = 'PS_SHOP_DOMAIN_SSL';
UPDATE ${DB_PREFIX}shop_url SET domain = '{lando-name}.lndo.site', domain_ssl = '{lando-name}.lndo.site';
UPDATE ${DB_PREFIX}configuration SET value = 1 WHERE name = 'PS_SSL_ENABLED';
UPDATE ${DB_PREFIX}configuration SET value = 1 WHERE name = 'PS_SSL_ENABLED_EVERYWHERE';
EOF
lando ssh -s database -c "mysql -uroot ${RECIPE} -e 'source /app/urls.sql'"
rm urls.sql
```

## Step 14 — Clear cache

```bash
cd ~/developer/{workspace}/{project-name} && lando php bin/console cache:clear
```

If that fails, ask user before deleting:
> "cache:clear falló. ¿Borro var/cache/ manualmente? (y/n)"

If yes:
```bash
rm -rf ~/developer/{workspace}/{project-name}/var/cache/*
```

## Step 15 — Verify before declaring done

```bash
curl -sk -o /dev/null -w "front=%{http_code} -> %{redirect_url}\n" https://{lando-name}.lndo.site/
curl -sk -o /dev/null -w "admin=%{http_code} -> %{redirect_url}\n" https://{lando-name}.lndo.site/{admin-dir}/
```
Expected: front `200`, admin `302` to `/login`. A `302` on the front pointing back at itself is
the SSL loop from Step 13.

## Step 15b — Backoffice user (only if the user has no account in this DB)

Ask before creating one. Generate the hash with the project's own PHP, then insert:
```bash
cd ~/developer/{workspace}/{project-name}
lando php -r 'echo password_hash("{password}", PASSWORD_BCRYPT, ["cost"=>10]), PHP_EOL;'
```

Copy `id_profile` (1 = SuperAdmin), `id_lang` and `bo_theme` from an existing row, and insert
into both `{prefix}employee` and `{prefix}employee_shop`.

**`stats_date_from` / `stats_date_to` must hold real dates.** Left NULL, the Dashboard dies with
`PrestaShopException: Date must be a string` — `HelperCalendar::setDateFrom()` falls back to
`strtotime('-31 days')`, which returns an int, and the very next line demands `is_string()`.
`stats_compare_from` / `stats_compare_to` want `'0000-00-00'`, which needs `SET sql_mode='';`
at the top of the SQL file.

Verify the login actually works — do not just check the row exists:
```bash
TOKEN=$(curl -sk -c /tmp/c.txt "https://{lando-name}.lndo.site/{admin-dir}/login" \
  | grep -o 'name="_token"[^>]*value="[^"]*"' | head -1 | sed 's/.*value="//;s/"//')
curl -sk -b /tmp/c.txt -c /tmp/c.txt -o /dev/null -w "login=%{http_code} -> %{redirect_url}\n" \
  -X POST "https://{lando-name}.lndo.site/{admin-dir}/login" \
  --data-urlencode "_token=$TOKEN" --data-urlencode "email={email}" \
  --data-urlencode "passwd={password}" --data-urlencode "submitLogin=1"
```
Then follow the returned URL and confirm the Dashboard renders `200`.

Warn the user: this employee lives only in the local DB and disappears on the next DB refresh.

## Step 16 — Done

Clean dump only after successful import (verified in Step 12):
```bash
rm ~/developer/{workspace}/{project-name}/dump.sql
```

Report to user:
- Local URL: https://{lando-name}.lndo.site
- Admin URL: https://{lando-name}.lndo.site/admin{suffix} (check admin folder name in project root)
- DB port: {port} (for TablePlus/phpMyAdmin)

---

## DB Refresh Mode

Use when Lando is already set up and running. Only refreshes the database.

### Confirm inputs
- Project folder (e.g. `ps8-pincolor`)
- SSH host (check ~/.ssh/config — use existing alias if available)
- Environment: **prod** or **pre** — ALWAYS ask, never assume

### R0 — Read local credentials
Check parameters.php exists:
```bash
test -f ~/developer/{workspace}/{project}/app/config/parameters.php || echo "MISSING"
```
If MISSING: stop and tell user to run full setup or sync app/config from server first.

Read the **prefix** locally — but NOT the credentials. After a full setup the local
`parameters.php` holds the Lando values (`lemp`/`lemp`/`lemp`), which are useless for dumping
the server:
```bash
grep -E "database_prefix" ~/developer/{workspace}/{project}/app/config/parameters.php
```

Read the remote credentials from the SERVER's own copy:
```bash
ssh -i ~/.ssh/{ssh_key} {SSH_TARGET} "grep -E \"database_(name|user|password)\" {REMOTE_ROOT}/app/config/parameters.php"
```

### R1 — Test SSH + Dump from server

Test connectivity first:
```bash
ssh -i ~/.ssh/{ssh_key} -o ConnectTimeout=5 -o BatchMode=yes {SSH_TARGET} "echo OK"
```

Then dump (key-based):
```bash
ssh -i ~/.ssh/{ssh_key} {SSH_TARGET} \
  "mysqldump -h 127.0.0.1 -u'${DB_USER}' -p'${DB_PASS}' '${DB_NAME}'" > ~/developer/{workspace}/{project-name}/dump.sql
```

Verify: `wc -c ~/developer/{workspace}/{project-name}/dump.sql` — must be > 100KB.

### R2 — Import (drop + recreate)

```bash
cd ~/developer/{workspace}/{project-name} && lando db-import dump.sql
```

Verify tables after import (`lando mysql -e` returns nothing — query inside the service):
```bash
cd ~/developer/{workspace}/{project-name}
lando ssh -s database -c "mysql -uroot ${RECIPE} -e 'SELECT COUNT(*) FROM ${DB_PREFIX}configuration'"
```
If fails, keep dump and report to user.

### R3 — Fix URLs

Same SQL and the same SSL-must-be-1 rule as Step 13 — see there for why:
```bash
cd ~/developer/{workspace}/{project-name}
cat > urls.sql <<EOF
UPDATE ${DB_PREFIX}configuration SET value = '{lando-name}.lndo.site' WHERE name = 'PS_SHOP_DOMAIN';
UPDATE ${DB_PREFIX}configuration SET value = '{lando-name}.lndo.site' WHERE name = 'PS_SHOP_DOMAIN_SSL';
UPDATE ${DB_PREFIX}shop_url SET domain = '{lando-name}.lndo.site', domain_ssl = '{lando-name}.lndo.site';
UPDATE ${DB_PREFIX}configuration SET value = 1 WHERE name = 'PS_SSL_ENABLED';
UPDATE ${DB_PREFIX}configuration SET value = 1 WHERE name = 'PS_SSL_ENABLED_EVERYWHERE';
EOF
lando ssh -s database -c "mysql -uroot ${RECIPE} -e 'source /app/urls.sql'"
rm urls.sql
```

### R4 — Clear cache

```bash
cd ~/developer/{workspace}/{project-name} && lando php bin/console cache:clear
```

### R5 — Done
Delete dump (only if import verified): `rm ~/developer/{workspace}/{project-name}/dump.sql`. Report URL.

**Warn the user what the refresh just wiped:** any local-only backoffice employee (Step 15b) and
any builder/DB config that is not git-tracked (Elementor/stsitebuilder kit settings, per-widget
CSS). Offer to recreate the employee.

---

## Notes

- ALWAYS confirm prod vs pre with user before SSHing anywhere
- Enabling debug mode from the backoffice rewrites `config/defines.inc.php` (`_PS_MODE_DEV_` → `true`).
  That file IS git-tracked — flag it to the user so it never reaches a commit.
- The server sync leaves tracked files modified (module CSS caches, etc.) and new untracked dirs.
  List them at the end; never commit them without being asked.
- Never run destructive SQL (DROP, TRUNCATE) without explicit user confirmation
- Large DBs (>500MB) can take 5+ min to import — normal
- If shop shows wrong theme/images: check module status or trigger `lando-img-placeholder` skill
- If `lando db-import` hangs > 30min: check disk space with `lando exec appserver df -h`

### Troubleshooting — 502 en lemp con varios proyectos Lando a la vez

Síntoma: healthchecks de `lando start` fallan con 502 y el log de nginx muestra
`connect() failed (111: Connection refused) ... upstream: "fastcgi://<IP>:9000"`.

Diagnóstico: comparar esa IP contra
`docker network inspect landoproxyhyperion5000gandalfedition_edge -f '{{range .Containers}}{{.Name}} {{.IPv4Address}}{{"\n"}}{{end}}'`.
Si la IP pertenece al appserver de OTRO proyecto → colisión de alias DNS.

Causa: `fastcgi_pass appserver:9000` (o cualquier hostname literal). El alias
`appserver` existe en la red compartida del proxy para cada app Lando corriendo,
y nginx además cachea la IP al arrancar (resolución estática). Doble fallo:
puede resolver al contenedor equivocado, o quedarse con una IP stale tras un
restart.

Fix: el patrón `resolver 127.0.0.11` + `set $fpm_upstream fpm:9000` +
`fastcgi_pass $fpm_upstream` que ya incluye la plantilla de
`.lando/nginx-site.conf` de esta skill (Step 9). En proyectos antiguos con el
vhost viejo, aplicar ese mismo cambio y `lando restart`.
NO sirve usar el FQDN `appserver.{app}.internal` — ese alias puede no apuntar
al contenedor FPM y sigue siendo resolución estática.
- `lando db-import` only accepts paths inside the project directory — always dump to project root as `dump.sql`
