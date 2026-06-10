---
name: ps-css-build
description: >
  Sets up a modern CSS build pipeline for a PrestaShop child theme: source
  partials in _dev/css/, lightningcss bundle+minify to a single generated
  assets/css/custom.css, pnpm 11 with supply-chain hardening, pre-commit hook
  with zero-dependency Node fallback, and agent/team documentation. Replaces
  runtime @import chains and manual ?v= cache-bust versioning. Use when
  scaffolding CSS architecture for a new PS child theme, when a theme has
  multiple @import'ed CSS partials loaded at runtime, when the user mentions
  manual CSS version bumps, CSS cache busting problems, "el CSS no se
  actualiza", or wants the Milagros-style CSS build on another project.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# PS CSS Build — pipeline para child themes

Montado y probado en `milagros-colombia-b2b-v9.1` (PS 9.1, child de Panda).
Sustituye: `@import` runtime (N requests encadenados) + bumps manuales `?v=N`
+ hacks tipo `_bust={$smarty.now}` (que anulan caché por completo).

## Resultado final

```
themes/{child}/
├── _dev/
│   ├── README.md            # doc del sistema (plantilla abajo)
│   └── css/
│       ├── custom.css       # entry point: @import de los parciales (SIN ?v=)
│       ├── tokens.css       # :root design tokens
│       ├── base.css
│       ├── components.css
│       ├── utilities.css
│       └── pages/*.css
├── scripts/bundle-css.js    # fallback cero-deps para el hook
├── assets/css/custom.css    # GENERADO — único CSS servido
├── package.json             # pnpm + scripts css/watch
├── pnpm-workspace.yaml      # supply-chain hardening
├── pnpm-lock.yaml
└── .gitignore               # node_modules/
.githooks/pre-commit         # raíz del repo
.gitattributes               # dist marcado generated
```

PS core auto-registra `assets/css/custom.css` como `theme-custom` — mantener
ese path/nombre exacto = cero cambios PHP, tema padre intacto.

## Paso 1 — Mover fuentes

```bash
cd themes/{child}
mkdir -p _dev/css
git mv assets/css/*.css _dev/css/        # parciales + custom.css
git mv assets/css/pages _dev/css/pages   # si existe
```

Limpiar el entry point `_dev/css/custom.css`:
- Quitar todos los `?v=N` de los imports: `sed -i '' -E "s/\.css\?v=[0-9]+'/.css'/g" _dev/css/custom.css`
- **CRÍTICO**: buscar `@import` con URL remota (Google Fonts) en TODOS los
  parciales: `grep -rn "@import url('http" _dev/css/`. lightningcss NO bundlea
  remotas — panic con error críptico `ResolverError NotFound`. Moverlas como
  `<link rel="preconnect">` + `<link rel="stylesheet">` al override
  `templates/_partials/stylesheets.tpl` (mejor perf además: descarga paralela).
- Revisar `url()` relativos en parciales (`grep -rn "url(" _dev/css/ | grep -v @import`).
  Rutas absolutas (`/img/...`) OK; relativas pueden romper al cambiar de nivel.

## Paso 2 — package.json + pnpm hardened

`themes/{child}/package.json`:

```json
{
  "name": "{child}-theme",
  "private": true,
  "packageManager": "pnpm@11.1.2",
  "scripts": {
    "css": "lightningcss --bundle --minify --targets 'defaults' _dev/css/custom.css -o assets/css/custom.css",
    "watch": "chokidar '_dev/css/**/*.css' -c 'pnpm run --silent css' --initial"
  },
  "devDependencies": {
    "chokidar-cli": "^3.0.0",
    "lightningcss-cli": "^1.30.2"
  }
}
```

Pin de versión pnpm **exacta** (la instalada: `pnpm --version`), no `11.x.x`.

`themes/{child}/pnpm-workspace.yaml` (ver skills `supply-chain-security` y
`gitlab-security-setup` para el detalle):

```yaml
allowBuilds:
  lightningcss-cli: true   # único postinstall permitido

minimumReleaseAge: 4320 # 72h quarantine (estándar Cinetic)

blockExoticSubdeps: true # bloquea transitivas git/tarball (pnpm 10.26+)

overrides:
  form-data: ">=4.0.4"
  axios: ">=1.15.2"
  lodash: ">=4.18.0"
  picomatch: ">=4.0.4"
  qs: ">=6.14.2"
```

`themes/{child}/.gitignore`: `node_modules/`

```bash
pnpm install --dir themes/{child}
pnpm --dir themes/{child} run css     # genera el dist
```

**Gotcha**: pnpm 11 bloqueará el postinstall de lightningcss-cli en el primer
install y escribirá un placeholder `allowBuilds: lightningcss-cli: set this to
true or false` en pnpm-workspace.yaml — editarlo a `true` y re-instalar.

**Test real**: `rm -rf node_modules && pnpm install --frozen-lockfile` —
la caché caliente NO aplica las políticas.

## Flujo de desarrollo diario (documentar al usuario)

Mientras se edita CSS, el watcher debe estar corriendo en una terminal:

```bash
pnpm --dir themes/{child} run watch
```

- Guardas un parcial en `_dev/css/` → chokidar dispara lightningcss → dist
  regenerado en ~10ms → recarga del browser y listo.
- `--initial` hace un build al arrancar, así que el dist nunca queda stale
  al empezar la sesión.
- Sin watch corriendo, los cambios en `_dev/css/` NO se ven en el browser
  (el browser carga el dist generado, no los fuentes). Síntoma típico:
  "cambié el CSS y no sale" → watch no estaba corriendo → build puntual:
  `pnpm --dir themes/{child} run css`.
- Combina con la skill `ps-watch` (BrowserSync) para recarga automática del
  browser además del rebuild: chokidar regenera el dist y BrowserSync detecta
  el cambio en `assets/css/` y recarga. Watch + BrowserSync = guardar y ver.

## Paso 3 — Limpiar el tpl

Revisar `templates/_partials/stylesheets.tpl` del child: si tiene hacks de
cache-bust (`?_bust={$smarty.now}`, timestamps), quitarlos — anulan caché en
cada pageview. Si el override solo existía para eso, dejarlo igual que el
padre + el bloque de fonts:

```smarty
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link rel="stylesheet" href="https://fonts.googleapis.com/css2?family={...}&display=swap">

{foreach $stylesheets.external as $stylesheet}
  <link rel="stylesheet" href="{$stylesheet.uri}" media="{$stylesheet.media}">
{/foreach}
...resto igual que el padre (Panda añade $sttheme.custom_css al final)...
```

Cache-busting de producción: el dist commiteado cambia de contenido → ETag
nuevo → revalidación. Si nginx sirve `.css` con `expires` largo, opciones:
`Cache-Control: no-cache` para el dir del tema (304s baratos), o activar
CCC "Smart cache CSS" en BO Rendimiento (ahora funciona porque el CSS es
plano — con @imports runtime se rompía).

## Paso 4 — Hook pre-commit con fallback cero-deps

`.githooks/pre-commit` (raíz, ejecutable `chmod +x`):

```sh
#!/bin/sh
# Rebuild theme CSS bundle when staged changes touch _dev/css sources.
# Works for everyone: pnpm if installed, zero-dep Node.js bundler otherwise.

THEME="themes/{child}"

if git diff --cached --quiet -- "$THEME/_dev/css"; then
  exit 0
fi

echo "pre-commit: _dev/css changed — rebuilding custom.css..."

if [ -d "$THEME/node_modules" ]; then
  pnpm --dir "$THEME" run --silent css || exit 1
  echo "pre-commit: rebuilt OK (minified)"
else
  node "$THEME/scripts/bundle-css.js" || exit 1
  echo "pre-commit: rebuilt OK (unminified — run 'pnpm install --dir $THEME' for minified)"
fi

git add "$THEME/assets/css/custom.css"
```

`themes/{child}/scripts/bundle-css.js` — bundler Node sin dependencias
(resuelve @imports recursivo, salta remotas). Copiarlo de
`milagros-colombia-b2b-v9.1/themes/milagros/scripts/bundle-css.js` ajustando
nada (usa rutas relativas a `__dirname`).

Activación (por clone, git no auto-activa hooks por diseño):

```bash
git config core.hooksPath .githooks
```

`.gitattributes` (raíz):

```
themes/{child}/assets/css/custom.css linguist-generated=true -diff
```

## Paso 5 — Documentación (obligatorio, es la red de seguridad real)

El hook NO se activa solo en clones nuevos. La protección para compañeros/
agentes que no conocen el sistema es documentación:

1. **CLAUDE.md del proyecto** — sección "CSS del tema hijo — REGLA OBLIGATORIA":
   dist generado (jamás editarlo), fuentes en `_dev/css/`, comando rebuild,
   activación del hook, sin `?v=`, remotas como `<link>`, no tocar hardening
   pnpm. Copiar la sección de `milagros-colombia-b2b-v9.1/CLAUDE.md`.
2. **`themes/{child}/_dev/README.md`** — doc completa: diagrama del flujo,
   setup por clone, watch, cómo funciona el hook (2 caminos), supply chain,
   tabla troubleshooting. Copiar de
   `milagros-colombia-b2b-v9.1/themes/milagros/_dev/README.md`.

## Paso 6 — Verificar

1. `pnpm --dir themes/{child} run css` → dist generado, comparar tamaño razonable
2. Cascade: `tail -c 300 assets/css/custom.css` — los overrides del entry deben
   estar al FINAL (lightningcss respeta orden de imports)
3. Hook ambos caminos: stage un parcial → `.githooks/pre-commit` → exit 0 y
   dist staged; repetir con `node_modules` renombrado (camino fallback)
4. En Lando con agent-browser: página renderiza, `<link>` custom.css sin
   query hacks, tokens aplicados (`getComputedStyle` de una variable), fuente OK
5. `rm -rf node_modules && pnpm install --frozen-lockfile && pnpm run css`

## Limitación conocida (aceptada)

Si alguien clona, NO activa el hook y commitea cambios de `_dev/css/`, el dist
queda desactualizado y "su cambio no sale". Mitigación: documentación (Paso 5)
+ el que integre con agente lo detecta por CLAUDE.md. Opciones descartadas en
Milagros por complejidad/deps: Husky, package.json raíz con prepare, build en
deploy, job CI de sincronía (válida si el proyecto ya tiene runners — ver
`gitlab-security-setup`).
