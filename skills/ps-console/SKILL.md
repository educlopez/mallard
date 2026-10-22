---
name: ps-console
description: >
  Run PrestaShop maintenance and one-off scripts inside a project's local Lando
  container: clear the cache, run `bin/console` commands, and execute ad-hoc PHP
  that boots PrestaShop (read an Employee, check a Configuration value, list a
  hook, query a product through the classes). Use when the user says "limpia la
  caché", "cache clear", "bin/console", "ejecuta esto con PrestaShop cargado",
  "mira esta config de PS desde PHP", "haz un script rápido con la clase X".
  Local Lando only: never against PRE or PROD.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# ps-console

Short, safe recipes for the PHP-in-Lando commands that come up in almost every PrestaShop session.

## Rules

- Local only. Run from the project root where `.lando.yml` lives. If the user wants this on a server, stop: servers have their own deploy steps (OPcache, PHP-FPM reload) and other people use them.
- Scripts that write to the DB are fine locally, but say what they write. Throw-away scripts are never committed.
- Check Lando is up first (`lando info`); if not, use the `lando-check` skill.

## Cache

PrestaShop 8/9 (Symfony):

```bash
lando php bin/console cache:clear --env=prod 2>&1 | tail -3
lando php bin/console cache:clear --env=dev 2>&1 | tail -3     # when the project runs in dev mode
```

All versions, including 1.7 where `bin/console` is limited or missing, clear Smarty, XML and media caches by bootstrapping the shop:

```bash
lando php -r 'require "config/config.inc.php"; Tools::clearAllCache(); echo "cache cleared\n";'
```

If `bin/console` fails with a container/compile error, the Symfony cache is corrupt: remove `var/cache/<env>/` and run `cache:clear` again.

Front-office template changes (`.tpl`) need the Smarty cache cleared (or "Force compilation" in Performance); CSS/JS also need the theme asset cache cleared (`Media::clearCache()` is included in `clearAllCache`).

## One-off PHP with PrestaShop loaded

```bash
lando php -r 'require "config/config.inc.php"; /* your code */'
```

Notes that avoid the usual confusion:

- CLI has no `HTTP_HOST`. On multistore or multi-domain projects set it before the require so the right shop and URLs resolve: `$_SERVER["HTTP_HOST"]="<name>.lndo.site";`.
- Use the project classes, not raw SQL, when you only read data: `new Employee(13)`, `Configuration::get("PS_SHOP_NAME")`, `Product::getProductName(…)`. Use `Db::getInstance()->executeS(...)` for queries; the table prefix is `_DB_PREFIX_`.
- For anything longer than a line, write a throw-away script inside the project (Lando mounts only the project directory), e.g. `_tmp-check.php`, run it with `lando php _tmp-check.php`, and delete it afterwards. Never leave it tracked; check `git status` before any commit.
- Print results as plain text or `json_encode`, and keep them short. Do not echo secrets (`_COOKIE_KEY_`, DB password, API keys).

## Admin / BO checks without a browser

- Employee exists and is active: `lando php -r 'require "config/config.inc.php"; $e=new Employee(13); echo $e->email," ",(int)$e->active,"\n";'`
- A module is installed/enabled: `lando php -r 'require "config/config.inc.php"; var_dump(Module::isInstalled("<name>"), Module::isEnabled("<name>"));'`
- Hook registered modules: `lando php -r 'require "config/config.inc.php"; print_r(array_column(Hook::getHookModuleExecList("<hookName>")?:[], "module"));'`
