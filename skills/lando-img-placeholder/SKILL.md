---
name: lando-img-placeholder
description: Sets up a static image placeholder for local Lando PrestaShop development so broken/missing product images show a nice placeholder instead of broken icons. Use this skill when the user says "placeholder de imagenes en lando", "lando image placeholder", "imagenes rotas en lando", "broken images in lando", or wants to avoid downloading hundreds of product images for local development. Trigger proactively whenever the user is working on a Lando PrestaShop project and mentions broken images or missing product photos.
version: "0.2.0"
metadata:
  author: Eduardo Calvo
---

# Lando Image Placeholder for PrestaShop

Intercepts missing product image requests and serves a static placeholder JPEG. Works for both Apache (lamp) and nginx (lemp) Lando recipes.

## Step 0 — Detect recipe

Read `.lando.yml` and check `recipe:`:
- `recipe: lamp` → follow **Apache path**
- `recipe: lemp` → follow **nginx path**

---

## Apache path (recipe: lamp)

### 1. Create `.lando/apache-placeholder.conf`

Replace `{project}` with the `name:` field from `.lando.yml`:

```apache
<VirtualHost *:80>
  ServerName {project}.lndo.site
  DocumentRoot /app
  <Directory /app>
    Options Indexes FollowSymLinks
    AllowOverride All
    Require all granted
  </Directory>
  RewriteEngine On
  # PS friendly image URLs: /{id}-{type}_default/...
  RewriteCond /app%{REQUEST_URI} !-f
  RewriteCond %{REQUEST_URI} ^/[0-9]+-[a-z]+_default
  RewriteCond %{REQUEST_URI} \.(jpg|jpeg|png|gif|webp)$
  RewriteRule ^ /img/placeholder-dev.jpg [L]
  # PS traditional image paths: /img/p/, /img/c/, /img/m/, /img/st/
  RewriteCond /app%{REQUEST_URI} !-f
  RewriteCond %{REQUEST_URI} ^/img/(p|c|m|st)/
  RewriteCond %{REQUEST_URI} \.(jpg|jpeg|png|gif|webp)$
  RewriteRule ^ /img/placeholder-dev.jpg [L]
  # NOTE: /stupload/ excluded — slider/module images live there.
</VirtualHost>
```

The `RewriteCond /app%{REQUEST_URI} !-f` checks the real filesystem path in the container — only fires for truly missing files.

### 2. Get the placeholder image

See [Choosing the placeholder image](#choosing-the-placeholder-image) below — prefer a real
product photo from the shop over a generic stock image.

### 3. Update `.lando.yml`

Add `vhosts` key under `services.appserver.config` (merge, don't overwrite):

```yaml
services:
  appserver:
    config:
      vhosts: .lando/apache-placeholder.conf
```

### 4. Rebuild

```bash
lando rebuild -y
```

---

## nginx path (recipe: lemp)

Placeholder rules go **directly into `.lando/nginx-site.conf`** — no separate file needed. If `nginx-site.conf` already exists (standard PS9 setup), just add the location blocks. If it doesn't exist, create it.

### 1. Add/verify placeholder blocks in `.lando/nginx-site.conf`

These two `location` blocks must appear **before** the main `location /` block:

```nginx
# Placeholder — friendly URLs: /{id}-{type}_default/...
location ~* ^/[0-9]+-[a-z]+_default.*\.(jpg|jpeg|png|gif|webp)$ {
    try_files $uri /img/placeholder-dev.jpg;
}

# Placeholder — classic PS paths: /img/p/, /img/c/, /img/m/, /img/st/
location ~* ^/img/(p|c|m|st)/.*\.(jpg|jpeg|png|gif|webp)$ {
    try_files $uri /img/placeholder-dev.jpg;
}
```

If `nginx-site.conf` doesn't exist yet, use the full PS9 template from lando-setup skill Step 9 — it already includes these blocks.

### 2. Get the placeholder image

See [Choosing the placeholder image](#choosing-the-placeholder-image) below — prefer a real
product photo from the shop over a generic stock image.

### 3. `.lando.yml` must reference the nginx config

```yaml
services:
  appserver:
    config:
      vhosts: .lando/nginx-site.conf
```

If already set, no change needed.

### 4. Rebuild

```bash
lando rebuild -y
```

---

## Choosing the placeholder image

Use a **real product photo from this shop**, not a generic stock image. A landscape or abstract
photo repeated across a product grid reads as "everything is broken" and makes it impossible to
judge layout, aspect ratio or crop while working. A real product photo makes the grid look like
the real store.

Try these in order and stop at the first that works.

### A — a product image already on disk

Works when `img/` was synced (fully or partially) from the server.

```bash
SRC=$(find img/p -name "*-home_default.jpg" -size +15k 2>/dev/null | head -1)
[ -z "$SRC" ] && SRC=$(find img/p -name "*.jpg" -size +20k 2>/dev/null | head -1)
[ -n "$SRC" ] && cp "$SRC" img/placeholder-dev.jpg && echo "placeholder <- $SRC"
```

The `-size` filter skips the tiny 1×1 and empty-thumb files PS leaves around, which would
otherwise be picked and give an invisible placeholder.

### B — pull exactly ONE image from the server

Works when `img/` was skipped entirely (the usual case for multi-GB shops). Costs one file, not
the whole folder.

```bash
REMOTE_SRC=$(ssh -i ~/.ssh/{ssh_key} {SSH_TARGET} \
  "find {REMOTE_ROOT}/img/p -name '*-home_default.jpg' -size +15k 2>/dev/null | head -1")
rsync -az -e "ssh -i ~/.ssh/{ssh_key}" {SSH_TARGET}:"$REMOTE_SRC" img/placeholder-dev.jpg
```

Also worth grabbing the shop logo while you are there — a missing logo in the header is just as
misleading as missing products:

```bash
rsync -az -e "ssh -i ~/.ssh/{ssh_key}" {SSH_TARGET}:{REMOTE_ROOT}/img/logo.jpg img/logo.jpg
```

### C — generic stock photo (last resort)

Only when there is no server access at all and `img/` is empty.

```bash
curl -L "https://images.unsplash.com/photo-1611572789411-6240f6cea970?q=80&w=500&h=500&fit=crop" -o img/placeholder-dev.jpg
```

### Sizing

Whatever the source, a square-ish image around 500–800px works best: PS requests the same URL for
every thumb size, and a square scales acceptably into both `home_default` and `large_default`
containers. Do not use a wide landscape — it letterboxes in portrait product cards.

## Verification (both paths)

```bash
curl -I "http://{project}.lndo.site/img/p/1/test.jpg"
# Should return HTTP/1.1 200 OK
```

## Notes

- `/stupload/` is excluded in both paths — sliders/modules store images there.
- `try_files $uri /img/placeholder-dev.jpg` in nginx only fires if the file doesn't exist (nginx `try_files` checks real disk).
- `.lando.yml` is typically gitignored in PS projects — safe to modify locally.
