---
name: ps-translate
description: >
  Automatiza la detección y traducción de strings sin traducir en una instalación
  PrestaShop, a CUALQUIER idioma. Detecta el child theme y los locales instalados
  (DB ps_lang con fallback a filesystem), escanea templates .tpl buscando llamadas
  {l s='...' d='...'} del sistema i18n de PS, compara contra los XLF existentes del
  locale objetivo y traduce los strings faltantes usando Claude de forma nativa —
  sin API key externa. Soporta varios idiomas en una sola corrida.
  Usa esta skill cuando el usuario mencione: "traducciones PS", "strings en inglés",
  "ps-translate", "traducir prestashop", "traducir a francés/alemán/español", "i18n
  pendiente", "hay textos sin traducir en la tienda", "¿qué strings faltan por
  traducir?" o cuando vea texto sin traducir en el storefront.
version: "2.0.1"
metadata:
  author: Eduardo Calvo
---

# ps-translate — Traducciones PrestaShop automatizadas (multi-idioma)

Detecta strings sin traducir y los traduce directamente usando Claude Code, sin
API key externa. Agnóstico de idioma: detecta los idiomas instalados o los pregunta.

## Flujo de trabajo

### Paso 0 — Detectar theme e idiomas (SIEMPRE primero)

```bash
python3 ~/.claude/skills/ps-translate/scripts/detect.py --base <ruta_ps>
```

Devuelve el child theme, su parent, y los locales instalados:

```json
{ "theme": "mystore", "parent": "classic",
  "locales": ["es-ES", "fr-FR", "de-DE"], "source": "db" }
```

- `source: "db"` → conexión directa vía `pymysql` (requiere el host de DB accesible
  desde donde corre el script — típico en instalaciones no-Lando o con DB expuesta).
- `source: "db (lando)"` → cuando `database_host` es el nombre del servicio Docker
  interno (no resoluble desde el host), se usa `lando mysql <db> -e "..."` en su
  lugar — esto SÍ ve los idiomas realmente activos en `ps_lang`, no adivines por
  `source: "filesystem"` que no hay forma de saberlo en un proyecto Lando.
- `source: "filesystem"` → último recurso: deducido de qué dirs `translations/<locale>/`
  ya existen. Solo dice qué locales YA tienen alguna traducción empezada, NO qué
  idiomas están realmente activos en la tienda — si el proyecto usa Lando y esto
  aparece, algo falló antes (revisa que `.lando.yml` exista y `lando` esté en el PATH).

**Luego decide el/los idioma(s) objetivo:**
- Si hay UN solo locale (aparte de `en-US`) → úsalo directo.
- Si hay VARIOS → **pregunta al usuario** a cuál(es) traducir (acepta varios).
- Si `locales` viene vacío o el usuario quiere otro → pídele el locale (formato
  `es-ES`, `fr-FR`, `de-DE`, `it-IT`, `pt-PT`…). No asumas es-CO ni ninguno por defecto.

### Paso 1 — Escanear strings faltantes (por cada locale objetivo)

```bash
python3 ~/.claude/skills/ps-translate/scripts/scan.py \
  --base <ruta_ps> --lang <locale> --output /tmp/ps_missing_<locale>.json
```

`--theme` se auto-detecta; pásalo solo para forzar otro. Argumentos:

| Flag | Default | Descripción |
|------|---------|-------------|
| `--base` | `.` (cwd) | Raíz de la instalación PS |
| `--lang` | auto si hay uno solo | Locale objetivo (ej. `fr-FR`) |
| `--theme` | auto-detectado | Child theme |
| `--theme-only` | false | Solo storefront (`Shop.*`), ignora módulos terceros — recomendado |
| `--domain` | todos | Filtrar por dominio (ej. `ShopThemeCatalog`) |
| `--output` | stdout | Ruta a fichero JSON de salida |
| `--include-admin` | false | Incluir dominios de backoffice |

> **Consejo:** empieza con `--theme-only`. Sin él, el scan incluye TODOS los strings de
> módulos terceros (mailchimp, paypal, feeds…) que suelen ser miles y muchos ya vienen
> traducidos por el módulo. `--theme-only` deja solo lo que el cliente ve en la tienda.

Salida: `{locale, theme, parent, total, domains: {"Shop.Theme.X": [strings...]}}`.

### Paso 2 — Traducir los strings (Claude lo hace aquí)

Con el JSON del scan, traduce cada dominio al **locale objetivo** (el que sea).

**Reglas universales (cualquier idioma):**
- Traduce a la variante regional del locale (`es-ES` España, `es-CO` Colombia,
  `pt-PT` vs `pt-BR`, etc.). Español de España = usa "vosotros"/formal según tono
  del sitio; español LATAM = "ustedes", sin vosotros.
- **Nunca traduzcas** placeholders: `%s`, `%d`, `{variable}`, `[1]...[/1]`, `%%param%%`.
- Strings cortos (botones, etiquetas) → 3-4 palabras máx.
- Mensajes de error → directo y claro.
- Terminología e-commerce coherente. Glosario según idioma:
  - **es**: cart→carrito, order→pedido, shipping→envío, checkout→pago, discount→descuento
  - **fr**: cart→panier, order→commande, shipping→livraison, checkout→paiement, discount→remise
  - **de**: cart→Warenkorb, order→Bestellung, shipping→Versand, checkout→Kasse, discount→Rabatt
  - **it**: cart→carrello, order→ordine, shipping→spedizione, checkout→pagamento, discount→sconto
  - **pt**: cart→carrinho, order→encomenda, shipping→envio, checkout→pagamento, discount→desconto
  - Otro idioma → aplica el equivalente estándar de e-commerce.
- Salida: JSON `{"Shop.Theme.X": {"source": "traducción", ...}}` por dominio.

### Paso 3 — Escribir los XLF (por cada locale)

```bash
python3 ~/.claude/skills/ps-translate/scripts/write_xlf.py \
  --base <ruta_ps> --lang <locale> --translations <json_traducciones>
```

Escribe en `themes/<theme>/translations/<locale>/<Domain>.<locale>.xlf`, sin
duplicar entradas existentes. `--theme` auto-detectado. `--dry-run` para preview.

### Paso 4 — Limpiar caché y verificar

```bash
rm -rf <base>/var/cache/prod/*
git add themes/<theme>/translations/
git commit -m "feat(i18n): auto-translate missing strings (<locale>) via ps-translate"
```

---

## Multi-idioma en una corrida

Para traducir a varios locales, repite pasos 1-3 por cada uno:

```bash
for LANG in es-ES fr-FR de-DE; do
  python3 ~/.claude/skills/ps-translate/scripts/scan.py \
    --base . --lang "$LANG" --theme-only --output "/tmp/ps_missing_$LANG.json"
done
```

Traduce cada `/tmp/ps_missing_<locale>.json` con el glosario del idioma
correspondiente, y corre `write_xlf.py` una vez por locale.

---

## Qué dominios ignorar (admin, no críticos)

Por defecto el scanner omite `Admin.*`, `psgdpr`, `ps_themecusto`, `steasybuilder`,
`ps_facebook`, `ps_accounts`. Pasa `--include-admin` para incluirlos.

---

## Dónde se guardan las traducciones

Siempre en el **child theme**:
`themes/<theme>/translations/<locale>/<DomainKey>.<locale>.xlf`

Nunca se modifican el core (`/translations/`) ni el tema padre.

---

## Notas / limitaciones

- Detección DB requiere `pymysql` (`pip install pymysql`). Sin él o sin DB
  accesible, cae a filesystem — funciona igual, solo lista los locales con dir de
  traducciones ya generado.
- Heurística "ya traducido": salta strings con diacríticos del idioma destino
  (á, ñ, ç, ü, ß…). Un string sin diacríticos ya escrito en destino (ej. francés
  "Sort by"→"Trier par" no, pero "Ajouter au panier" sí sin acento) puede
  re-listarse; es inofensivo (se traduce a sí mismo). El diff contra los XLF es el
  filtro principal.
