---
title: Recovery — cliente editó el parent theme directamente
source: extrapolation from devdocs + dogfood Q5 (2026-05-19)
ps_version: [8, 9]
tags: [themes, child-theme, recovery, parent-theme, override, smarty]
last-updated: 2026-05-20
---

# Recovery: Cliente editó el parent theme directamente

> [!warning] Anti-patrón
> Editar `themes/<parent>/templates/...`, `themes/<parent>/assets/...` o cualquier cosa dentro del parent theme **rompe el contrato de updates**. La siguiente actualización del theme sobrescribe el trabajo del cliente sin aviso. Este documento es el protocolo para **recuperar** cambios mal aplicados y dejarlos donde corresponde: el child theme.

## Cuándo aplica este protocolo

Aparece típicamente al recibir un proyecto cliente donde:
- No existe `themes/<parent>_child/`, o existe pero no es el activo.
- El cliente (o un dev anterior) tocó `themes/<parent>/templates/`, `themes/<parent>/modules/<mod>/views/templates/hook/`, o assets del parent.
- Hay riesgo inminente de update del theme (Panda, Hummingbird, etc.) que pisaría los cambios.

Aplica igual para **Panda**, **Hummingbird**, **Classic** o cualquier otro theme PS.

---

## Protocolo de recuperación (orden estricto)

### 1. Capturar el diff antes de tocar nada

```bash
# Con git en el proyecto
git -C <project> log --all --diff-filter=M -- themes/<parent>/templates/catalog/product.tpl

# Sin git: diff contra distribución original
diff -u <original-theme-zip>/templates/catalog/product.tpl \
        themes/<parent>/templates/catalog/product.tpl \
        > /tmp/product.patch
```

> [!warning] Bloqueante sin baseline
> Si no hay git **y** no tienes la distribución original, no puedes separar "edits del cliente" de "tpl original". Para Panda: descarga el zip del theme desde la zona cliente SunnyToo o pide el archivo a quien lo instaló. **No avances sin baseline**.

### 2. Verificar / crear child theme

```bash
ls themes/                # debe mostrar <parent>/ y <parent>_child/
```

Si falta el child, créalo mínimo:

> [!info] El nombre del child NO tiene que seguir la convención `<parent>_child`
> El folder puede llamarse `panda-v2_child`, `mi-tema-custom`, etc. La relación se declara en `config/theme.yml` mediante el campo **`parent:`** (no `parent_theme:`).

```yaml
# themes/<child-name>/config/theme.yml
name: <child-name>
display_name: <Child display name>
version: 1.0.0
parent: <parent>                # <- campo CORRECTO (no `parent_theme:`)
use_parent_assets: true         # carga CSS/JS del parent además del child
author:
  name: <tu-empresa>
  email: dev@tu-empresa.com

meta:
  compatibility:
      from: 1.7.1.0
      to: ~

assets:
  css:
    all:
      - id: theme-custom-css
        path: assets/css/custom.css
  js:
    all:
      - id: theme-custom-js
        path: assets/js/custom.js
```

### 3. Confirmar que el shop usa el child, no el parent

```sql
SELECT theme_name FROM ps_shop WHERE id_shop = 1;
```

Si devuelve `<parent>` (no `<parent>_child`):
- Cambia desde BO → **Design → Theme & Logo** al child **antes** de continuar.
- Si no lo cambias, los overrides del child **no se cargarán** y el parent seguirá sirviéndose.

### 4. Mover el override al child (mismo path relativo)

Crea `themes/<parent>_child/templates/catalog/product.tpl` con el contenido **editado** del parent.

**Regla**: mismo path relativo = PS lo prioriza automáticamente sobre el del parent (cascada nativa).

> [!info] Override por bloque vs tpl entero
> Si el tpl original empieza con `{extends file='...'}` y tiene `{block}` parciales, considera **override por bloque** en vez de copiar el tpl entero:
>
> ```smarty
> {extends file='parent:catalog/product.tpl'}
> {block name='page_content'}
>   {* tu markup custom *}
> {/block}
> ```
>
> Más resiliente a updates del parent. Si Panda añade un bloque nuevo al parent, lo heredas gratis.

### 5. Restaurar el parent a su estado original

```bash
# Con git
git -C <project> checkout HEAD -- themes/<parent>/templates/catalog/product.tpl

# Sin git
cp <original-theme-zip>/templates/catalog/product.tpl \
   themes/<parent>/templates/catalog/product.tpl

# Verificar
diff themes/<parent>/templates/catalog/product.tpl <original>
# Debe ser vacío
```

### 6. Invalidar caché y verificar render

```bash
# Smarty compiled templates
rm -rf var/cache/*

# Si usas Docker / PHP-FPM
docker exec <project>_php_1 service php-fpm reload
# o
docker restart <project>_php_1
```

Recargar la PDP en navegador. **El render debe ser idéntico al pre-recovery**. Si difiere:

- Verifica `theme_name = <parent>_child` en `ps_shop` (paso 3).
- Verifica path exacto del override (case-sensitive en Linux).
- Comprueba que Symfony container/asset versioning no esté pinneado.

### 7. Auditar el resto del parent

```bash
git -C <project> diff --stat HEAD -- themes/<parent>/
```

Suele haber más archivos editados que el cliente no mencionó. **Repite el protocolo para cada uno**.

---

## Casos especiales

### Assets editados (CSS/JS/SCSS)

Semántica distinta a templates:
- Templates: child **sobrescribe** parent (mismo path = solo se sirve el child).
- CSS/JS del child: se cargan **además** del parent (no en vez de).

**Recovery para assets**:
- Cliente tocó `themes/<parent>/assets/css/style.css` directamente → mueve el delta a `themes/<parent>_child/assets/css/custom.css`.
- No copies el `style.css` entero del parent al child — sería un fork. Solo el delta.
- Restaura `themes/<parent>/assets/css/style.css` al original.

### Templates de módulos editados

Path: `themes/<parent>/modules/<mod>/views/templates/hook/X.tpl`.

Recovery → mover a:
```
themes/<parent>_child/modules/<mod>/views/templates/hook/X.tpl
```

Mismo principio: PS prioriza el child path.

### Multi-shop

Si hay multi-tienda, cada `id_shop` puede tener theme distinto:

```sql
SELECT id_shop, theme_name FROM ps_shop;
```

Verifica **todas** las tiendas y aplica el protocolo por cada una.

### `theme.yml` editado en el parent

El cliente puede haber editado `themes/<parent>/config/theme.yml` (layouts, hooks, dependencies). El child theme puede **sobreescribir** secciones del theme.yml desde `themes/<parent>_child/config/theme.yml` — pero con cautela: la sintaxis de override de theme.yml es menos documentada que la de templates.

**Estrategia**: identifica el delta, replica el delta en el child theme.yml, restaura parent al original, prueba.

---

## Riesgos a vigilar

- **Caché OPCache de PHP** además de Smarty. Si tras `rm -rf var/cache/*` sigue sirviendo el viejo → reiniciar PHP-FPM.
- **Symfony container** cacheado en PS 9: a veces requiere `php bin/console cache:clear --env=prod` (si la pregunta tocaba pantallas BO Symfony).
- **Asset versioning** (cache busters): un asset cacheado en CDN o reverse proxy puede seguir sirviéndose. Confirma el hash en el HTML servido.
- **Permisos**: tras los `cp` desde un zip, verifica que el user del web server pueda leer (`chown -R www-data:www-data themes/`).

---

## Checklist final (post-recovery)

- [ ] `theme_name` en DB apunta al child.
- [ ] `themes/<parent>/` matchea byte-a-byte la distribución original (`diff` vacío).
- [ ] Caché Smarty + OPCache invalidados.
- [ ] Render visual idéntico al pre-recovery (golden path).
- [ ] Update simulado del parent (descomprimir nueva versión del zip encima) NO rompe los overrides del child.
- [ ] Documentado en el README del proyecto que los overrides viven en el child y nunca en el parent.

---

## Referencias

- [[parent-child-feature]] — mecánica de child theme en PrestaShop.
- [[template-inheritance]] — `{extends}` y `{block}` en Smarty/PS.
- [[asset-management]] — gestión de assets en child themes.
- [[overriding-modules]] — overrides de templates de módulos.
- [[theme-yml]] — declarando layouts, hooks y assets en theme.yml.
