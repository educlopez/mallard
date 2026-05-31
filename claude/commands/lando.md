---
name: lando
description: Crear configuración Lando para PrestaShop
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# Crear configuración Lando para PrestaShop

Genera `.lando.yml`, `.lando/php.ini`, y (si nginx) `.lando/nginx-site.conf` para el proyecto PrestaShop actual.

## Argumento

`$ARGUMENTS` — Puerto MySQL para portforward (ej: 33405, 33410...)

## Instrucciones

### 1. Detectar nombre del proyecto
Último segmento del path del directorio actual.

### 2. Detectar versión PrestaShop y PHP

Leer `app/AppKernel.php` (busca `MAJOR_VERSION`) o `composer.json` (`prestashop/prestashop`):

| PrestaShop | PHP  |
|------------|------|
| 1.7.x      | 7.4  |
| 8.x        | 8.1  |
| 9.x        | 8.3  |

### 3. Preguntar web server

Preguntar al usuario: **¿nginx (lemp) o Apache (lamp)?**
- Default: **nginx** para PS9 (servidores Ploi usan nginx)
- Default: **Apache** para PS8 o versiones anteriores

### 4. Crear `.lando/php.ini`

```ini
max_execution_time = 300
max_input_time = 300
memory_limit = 512M
post_max_size = 64M
upload_max_filesize = 64M
```

### 5a. Si Apache (lamp) — crear `.lando.yml`

```yaml
name: NOMBRE_PROYECTO
recipe: lamp
config:
  php: "VERSION_PHP"
  webroot: ./
  database: mysql
  xdebug: false
services:
  appserver:
    config:
      php: .lando/php.ini
  database:
    portforward: PUERTO_ARGUMENTO
    type: mysql
  mail:
    type: mailhog
    hogfrom:
      - appserver
```

### 5b. Si nginx (lemp) — crear `.lando/nginx-site.conf` + `.lando.yml`

Crear primero `.lando/nginx-site.conf`:

```nginx
server {
    listen 80;
    root /app;
    index index.php index.html;
    client_max_body_size 128M;

    charset utf-8;

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
        fastcgi_pass fpm:9000;
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
        fastcgi_pass fpm:9000;
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
        fastcgi_pass fpm:9000;
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

Luego crear `.lando.yml`:

```yaml
name: NOMBRE_PROYECTO
recipe: lemp
config:
  php: "VERSION_PHP"
  webroot: ./
  database: mysql
  xdebug: false
services:
  appserver:
    config:
      vhosts: .lando/nginx-site.conf
      php: .lando/php.ini
  database:
    portforward: PUERTO_ARGUMENTO
    type: mysql
  mail:
    type: mailhog
    hogfrom:
      - appserver
```

### 6. Actualizar `app/config/parameters.php`

Cambiar SOLO los valores de conexión, conservando todo lo demás:

```
database_host     → 'database'
database_port     → ''
database_name     → 'lamp'
database_user     → 'lamp'
database_password → 'lamp'
mailer_transport  → 'smtp'
mailer_host       → 'sendmailhog'
mailer_user       → NULL
mailer_password   → NULL
```

### 7. Mostrar resumen

- Archivos creados
- URL: `https://NOMBRE_PROYECTO.lndo.site/`
- Puerto MySQL externo: `127.0.0.1:PUERTO`
- Credenciales DB: user=`lamp` pass=`lamp` db=`lamp` host=`database`
- Mailhog SMTP: `sendmailhog:1025`
- Comandos: `lando start` para levantar, `lando rebuild -y` si se cambia `.lando.yml`
