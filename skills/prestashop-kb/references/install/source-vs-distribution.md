---
source_url: https://build.prestashop-project.org/news/2026/prestashop-9-1-0-available/
ps_version: [8, 9]
ingested: 2026-04-19
section: PS source vs distribution
topic: install
tags: [prestashop, install, source, distribution, ps9, composer]
---

# PS source (git clone) vs distribution (zip release)

> [!warning] Por qué importa
> Desde PS 9.0, PrestaShop **ya no publica zips en GitHub Releases** — sólo el source tarball. La única forma de obtener un PS 9 "listo para instalar" es (a) descargar desde prestashop.com/versions/ el zip oficial, (b) clonar y construir desde source con composer. Las dos rutas dejan el disco en estados DIFERENTES que hay que conocer.

## Dos rutas, dos layouts

### Ruta A — Distribution zip (`prestashop_X.Y.Z.zip`)

Origen: `https://prestashop.com/versions/` (PS 9) o `github.com/PrestaShop/PrestaShop/releases` (PS 8 y anteriores).

El zip es un **"release artifact"** generado por el build tool `tools/build/CreateRelease.php`. Trae el filesystem listo para que el installer CLI corra directamente:

```
<root>/
├── admin/                   ← folder ya renombrado (random suffix a veces)
├── install/                 ← installer presente
├── vendor/                  ← composer deps compiladas (bundled)
├── themes/
│   ├── classic/             ← en PS 8
│   └── hummingbird/         ← en PS 9 (opcional según build config)
├── modules/                 ← stock modules incluidos
├── composer.json
├── autoload.php
├── init.php
├── index.php
└── ...
```

**Install**:
```bash
php install/index_cli.php --language=es --domain=... [otros flags]
```

No requiere composer, node, ni build steps. 2-3 minutos end-to-end.

### Ruta B — Source clone (`git clone --branch 9.1.0`)

Origen: `https://github.com/PrestaShop/PrestaShop.git` a un tag específico.

El repo source tiene el **código fuente de desarrollo**:

```
<root>/
├── admin-dev/               ← NO renombrado (sufijo -dev)
├── install-dev/             ← NO renombrado
├── (sin vendor/)            ← hay que generar con composer install
├── themes/
│   ├── classic/
│   └── hummingbird/         ← en main branch, PS 9 tags
├── modules/
├── admin-api/               ← presente en PS 9 source
├── tools/build/             ← CreateRelease.php + helpers
├── docker-compose.yml       ← PS 9 trae docker setup built-in
├── composer.json
├── composer.lock
└── ...
```

**Steps extra vs distribution**:

1. Renombrar `install-dev/` → `install/` y `admin-dev/` → `admin/` (o usar `tools/build/CreateRelease.php` que lo hace automáticamente + produce un zip).
2. `composer install --no-dev --optimize-autoloader` (2-4 min).
3. (Opcional) `npm ci && npm run build` en `themes/hummingbird/` si quieres activar Hummingbird.
4. `php install/index_cli.php ...` como en la ruta A.

Total: 5-7 minutos end-to-end.

## Diferencias que importan

| | Distribution zip | Source clone |
|---|---|---|
| Tamaño descarga | ~80 MB | ~100-200 MB (con .git + deps metadata) |
| Tiempo total install | 2-3 min | 5-7 min (por composer) |
| Disponibilidad | PS 8: GitHub ✅ / PS 9: **solo prestashop.com** | PS 9 git ✅ (reproducible por tag) |
| Folder installer | `install/` | `install-dev/` |
| Folder admin | `admin/` (a veces con sufijo random) | `admin-dev/` |
| `vendor/` (composer deps) | Incluido | Hay que generarlo |
| Assets front-end | Compilados (CSS/JS listos) | Hay que correr `npm run build` |
| Theme Hummingbird | Incluido según config del build | Siempre presente en PS 9 source |
| Reproducibilidad | Por URL + size check | Por git commit hash / tag + composer.lock |
| CI/CD friendly | ✅ Sencillo | ✅ Pero requiere más tooling |
| Legal/licensing | Igual (AFL/OSL) | Igual |

## Lo que hace el CLI `ps-lando`

Resuelve ambas rutas transparentemente:

- **PS 8.x** → descarga zip de GitHub releases (ruta A).
- **PS 9.x sin `--ps-zip`** → git clone + rename install-dev→install + admin-dev→admin + composer install (ruta B automatizada).
- **PS 9.x con `--ps-zip <path>`** → usa el zip local como ruta A.

> [!info] Caso de uso típico en equipo
> - El desarrollador A descarga una vez el zip PS 9 de prestashop.com/versions/ → lo comparte con el equipo vía cloud storage o cada uno lo baja con su licencia → todos usan `--ps-zip`.
> - Alternativa zero-infra: cada dev deja que el CLI clone + composer install cada vez. Más lento pero 100% reproducible sin coordinación.

## Caveats descubiertos en validación real

Probado 2026-04-19 con `ps-lando v0.1.3` en `panda-v2`:

- **Rename automático** funciona — `install-dev/` → `install/` + `admin-dev/` → `admin/` sin intervención manual.
- **Composer install** toma ~3-4 min en Apple Silicon con OrbStack/Docker Desktop, dependiendo de caché.
- **Hummingbird 2.0** aparece en `themes/hummingbird/` en un PS 9 clonado por tag 9.1.0. Si el usuario no activa el theme (queda en classic o se activa Panda), Hummingbird sigue ahí como opción latente. El CLI no lo borra — correcto (usuario puede querer probarlo).
- **Assets de Hummingbird NO se compilan** por el CLI — falta `npm ci && npm run build` en `themes/hummingbird/`. Si alguien activa Hummingbird sin build, se ve roto. Panda no tiene este problema (assets compilados vienen en el zip SunnyToo).
- **`_PS_VERSION_` detection**: en source clone no hay fichero que lo exponga (ni `install_version.php` ni `defines.inc.php` pre-install). Post-install, Symfony lo sabe por `composer.json` / `AppKernel`. Si alguien clonó source al tag `9.1.0`, la única forma confiable de saber la versión es mirando el tag git (`.git/HEAD` o similar).

## Implicaciones operativas

**Para el agente `panda-expert`**:
- Si al inspeccionar un proyecto ves `install-dev/` (no `install/`) → es un source clone sin rename aplicado. Útil para deducir cómo se creó el sandbox.
- Si falta `vendor/` → composer install no corrió. Install no funcionará hasta completar.
- Si ves `admin-dev/` → misma historia. URL `/admin/` no responderá.
- Si hay `themes/hummingbird/` y el assets falta compilar → el cliente podría tener problemas visuales si cambia a Hummingbird sin build step. Flaggear.

**Para Eduardo (dev)**:
- Si clonas repo PS 9 a mano: corre `tools/build/CreateRelease.php` para generar un zip equivalente al distribution, o renombra install-dev/admin-dev + composer install manualmente.
- Si vas a migrar sandboxes entre máquinas: prefiere distribution zip — menos partes móviles. O deja al CLI recrear desde source usando el mismo tag.

## Related in KB

- [[../themes/hummingbird-v2]] — theme default PS 9, presente en source y distribution
- [[../themes/parent-child-feature]] — child themes aplica igual en ambas rutas
- [[../../easybuilder/00-overview]] — Easy Builder: distribución zip propia de SunnyToo, no afectado por esta dicotomía
- [[../../_agent/panda-expert]] — el agente debería aplicar las reglas "workflow paso 0" considerando esta dicotomía
