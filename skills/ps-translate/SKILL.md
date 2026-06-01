---
name: ps-translate
description: >
  Automatiza la detección y traducción de strings sin traducir en una instalación
  PrestaShop. Escanea todos los templates .tpl y archivos PHP buscando llamadas
  {l s='...' d='...'} del sistema i18n de PS, compara contra los XLF existentes
  del locale activo (es-CO por defecto) y traduce los strings faltantes usando
  Claude de forma nativa — sin API key externa.
  Usa esta skill cuando el usuario mencione: "traducciones PS", "strings en inglés",
  "ps-translate", "traducir prestashop", "i18n pendiente", "hay textos en inglés
  en la tienda", "¿qué strings faltan por traducir?" o cuando vea texto en inglés
  en el storefront que debería estar en español.
version: "1.0.0"
metadata:
  author: Eduardo Calvo
---

# ps-translate — Traducciones PrestaShop automatizadas

Detecta strings sin traducir y los traduce directamente usando Claude Code, sin necesitar API key externa.

## Flujo de trabajo

### Paso 1 — Escanear strings faltantes

Corre el script de escaneado para obtener qué strings faltan:

```bash
python3 ~/.claude/skills/ps-translate/scripts/scan.py \
  --base <ruta_instalacion_ps> \
  --theme <nombre_child_theme> \
  --lang <locale>
```

**Argumentos:**

| Flag | Default | Descripción |
|------|---------|-------------|
| `--base` | `.` (cwd) | Raíz de la instalación PS |
| `--theme` | `milagros` | Nombre del child theme |
| `--lang` | `es-CO` | Locale objetivo |
| `--domain` | todos | Filtrar por dominio (ej. `ShopThemePanda`) |
| `--output` | stdout | Ruta a fichero JSON de salida |
| `--include-admin` | false | Incluir dominios de backoffice |

El script devuelve un JSON con este formato:

```json
{
  "locale": "es-CO",
  "theme": "milagros",
  "domains": {
    "ShopThemePanda": ["Filter", "Sort by", "No products were found."],
    "ShopThemeActions": ["Buy now", "Show all"]
  },
  "total": 44
}
```

### Paso 2 — Traducir los strings (Claude lo hace aquí)

Con el JSON del paso anterior, traduce cada dominio al locale objetivo.

**Instrucciones para Claude:**
- Traduce al **español colombiano (es-CO)** — natural, de UI, sin vosotros
- Mantén coherencia de terminología:
  - cart → carrito · order → pedido · shipping → envío
  - discount → descuento · checkout → pago/compra · product → producto
- **Nunca traduzcas** placeholders: `%s`, `%d`, `{variable}`, `[1]...[/1]`, `%%param%%`
- Strings cortos (botones, etiquetas) → máximo 3-4 palabras
- Mensajes de error → directo y claro, sin "Por favor, asegúrese de que..."
- Salida: JSON `{source: translation}` por dominio

### Paso 3 — Escribir los XLF

Una vez tienes las traducciones, escribe los nuevos `<trans-unit>` al child theme:

```bash
python3 ~/.claude/skills/ps-translate/scripts/write_xlf.py \
  --base <ruta_ps> \
  --theme <child_theme> \
  --lang <locale> \
  --translations <ruta_al_json_de_traducciones>
```

El script escribe en `themes/<theme>/translations/<lang>/<Domain>.<lang>.xlf`, sin duplicar entradas existentes.

### Paso 4 — Limpiar caché y verificar

```bash
# Limpiar caché PS
rm -rf <base>/var/cache/prod/*

# Verificar en el storefront
# Luego commitear
git add themes/<theme>/translations/
git commit -m "feat(i18n): auto-translate missing strings via ps-translate"
```

---

## Uso rápido (todo en uno)

Cuando el usuario pida traducir strings, ejecuta directamente:

```bash
# 1. Escanear (guarda JSON en /tmp/ps_missing.json)
python3 ~/.claude/skills/ps-translate/scripts/scan.py \
  --base . --output /tmp/ps_missing.json

# 2. Revisar qué hay
cat /tmp/ps_missing.json | python3 -c "import sys,json; d=json.load(sys.stdin); [print(f'[{k}] {len(v)} strings') for k,v in d['domains'].items()]"
```

Luego traduce el contenido del JSON y ejecuta `write_xlf.py` con las traducciones.

---

## Qué dominios ignorar (admin, no críticos)

Por defecto el scanner omite:
- `Admin.*` — backoffice, no lo ve el cliente
- `psgdpr`, `ps_themecusto`, `steasybuilder` — módulos admin/configuración
- `ps_facebook`, `ps_accounts` — módulos de integración externa

Pasa `--include-admin` si quieres incluirlos.

---

## Dónde se guardan las traducciones

Siempre en el **child theme** del proyecto:
```
themes/<theme>/translations/<lang>/<DomainKey>.<lang>.xlf
```

Nunca se modifican:
- `/translations/<lang>/` (core PS)
- `themes/panda/translations/` (tema padre)
