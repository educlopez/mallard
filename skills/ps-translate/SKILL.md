---
name: ps-translate
description: >
  Automatiza la detección y traducción de strings sin traducir en una instalación
  PrestaShop, a CUALQUIER idioma. Detecta el child theme y los locales instalados
  (DB ps_lang con fallback a filesystem/lando mysql), escanea templates .tpl
  (Smarty {l s='...' d='...'}) Y, con --include-emails, plantillas de email Twig
  (mails/themes/**/*.twig), incluyendo también las facturas/albaranes PDF de la
  raíz /pdf/**/*.tpl que antes se pasaban por alto. Compara contra los XLF
  existentes del locale objetivo y traduce los strings faltantes usando Claude
  de forma nativa — sin API key externa. Soporta varios idiomas en una sola corrida.
  Usa esta skill cuando el usuario mencione: "traducciones PS", "strings en inglés",
  "ps-translate", "traducir prestashop", "traducir a francés/alemán/español", "i18n
  pendiente", "hay textos sin traducir en la tienda", "¿qué strings faltan por
  traducir?", "traducir facturas/PDF", "traducir emails/correos de PrestaShop",
  o cuando vea texto sin traducir en el storefront, en una factura PDF, o en un
  email transaccional.
version: "2.2.0"
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
| `--include-emails` | false | Incluir plantillas de email Twig (`mails/themes/**/*.twig`, dominio `Emails.Body`) — ver sección "Facturas PDF y emails" |

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

## Facturas PDF y emails transaccionales

Estos dos NO son un mecanismo aparte que necesite un sub-skill propio — son el
MISMO sistema de dominios/XLF que todo lo demás, solo que las plantillas viven
fuera de `themes/**` y `modules/**` (donde el scanner mira por defecto) y, en el
caso de emails, usan una sintaxis de traducción distinta (Twig en vez de Smarty).
Por eso está resuelto ampliando `scan.py`, no creando skills nuevas.

**Facturas PDF** — ya cubierto automáticamente, sin flag extra:
- Viven en la raíz `/pdf/**/*.tpl` (fuera de `themes/`/`modules/`, por eso antes
  el scanner las pasaba por alto por completo — un `--domain Shop.Pdf` sin este
  fix solo encontraba strings de módulos de terceros con ese mismo dominio, ej.
  `psgdpr`, NO la factura real).
- Usan `{l s='...' d='Shop.Pdf'}` igual que el resto — un scan normal (con o sin
  `--theme-only`, ya que `Shop.Pdf` empieza por `Shop.`) ya las incluye.
- Overrides de PDF a nivel child-theme (`themes/<theme>/pdf/**`) ya estaban
  cubiertos por el glob de `themes/**/*.tpl` de siempre.

**Emails transaccionales** — requiere `--include-emails`, y ANTES de nada
comprueba qué motor de email usa el proyecto:

```sql
SELECT value FROM ps_configuration WHERE name = 'PS_MAIL_THEME';
```

- `modern` / `classic` → el proyecto usa el sistema Twig de PS 8/9
  (`mails/themes/<theme>/**/*.twig`), **agnóstico de idioma** (un solo fichero
  sirve para todos los locales; `locale` es una variable Twig). Los ficheros
  REALES por idioma que `Mail::send()` usa (`mails/<lang>/*.html`+`.txt`) no
  existen hasta que se GENERAN desde esos layouts — ver "Generación de plantillas"
  más abajo, es un paso aparte, no basta con tener el XLF traducido.
- Si el proyecto tiene una carpeta `mails/<lang>/*.html` + `.txt` POBLADA (ficheros
  reales, no solo `index.php`), puede estar en el sistema LEGACY de plantillas
  duplicadas por idioma — en ese caso, antes de asumir que hace falta traducir esos
  ficheros, compara los nombres contra `mails/themes/<theme activo>/core/*.html.twig`:
  si CADA fichero legacy tiene su equivalente Twig, el legacy está muerto/sin usar
  (el motor Twig activo tiene prioridad) y no hace falta tocarlo. Si encuentras
  ficheros legacy SIN equivalente Twig (típico de módulos que no migraron), esos sí
  se traducen duplicando el fichero completo por idioma — mecanismo manual, fuera
  del alcance de scan.py/write_xlf.py (no son strings sueltos, son ficheros enteros).
- Las líneas de ASUNTO (`Emails.Subject`) se fijan en PHP core
  (`$translator->trans('...', [], 'Emails.Subject')` dentro de clases como
  `SendProcessOrderEmailHandler`), no en las plantillas — no son escaneables por
  `scan.py`. Vienen ya traducidas en el paquete de idioma OFICIAL de PrestaShop
  para ese locale; si ese paquete no está instalado (comprueba si existe
  `translations/<locale>/` en la raíz del core, no solo en el theme), ni esto ni
  nada de este skill lo puede completar — es instalar el paquete oficial desde
  Admin → Internacional → Traducciones, no un gap de contenido que traducir a mano.

### El dominio de emails va en CORE, no en el theme — excepción confirmada empíricamente

`Emails.Body` (y cualquier otro dominio referenciado desde una plantilla de email,
ver más abajo) se resuelve, durante `prestashop:mail:generate`, por un traductor
Symfony que **NO lee `themes/<theme>/translations/` en absoluto** — probado
directamente: ni con el nombre de fichero con puntos (`Emails.Body.fr-FR.xlf`) ni
sin puntos (`EmailsBody.fr-FR.xlf`), colocado en la carpeta del theme, cambió nada;
solo funcionó al ponerlo en la raíz `translations/<locale>/`. Esto contradice la
norma general de este skill ("nunca tocar core") — es una excepción deliberada y
limitada SOLO a los dominios de email, confirmada comparando contra el paquete
oficial es-ES (`translations/es-ES/EmailsBody.es-ES.xlf` — sin puntos, exactamente
la convención `MailsBodyProviderDefinition::FILENAME_FILTERS_REGEX = ['#EmailsBody*#']`
del código fuente de PS9).

**Un email puede referenciar MÁS de un dominio.** Algunos layouts arman un `%label%`
traducido con un `|trans()` ANIDADO usando OTRO dominio (ej. `order_conf.html.twig`
construye `%order_history_label%` desde `'Order history and details'|trans({}, 'Shop.Theme.Customeraccount', locale)`
dentro de los parámetros del `|trans()` exterior). Ese dominio interior TAMBIÉN
necesita su copia en core, aunque sea un dominio normalmente theme-scoped que ya
funciona bien en el storefront — el traductor de generación de emails no distingue,
trata todo lo referenciado desde dentro de una plantilla de email igual. Para
encontrar TODOS los dominios que hacen falta en core para un proyecto dado:

```bash
python3 ~/.claude/skills/ps-translate/scripts/scan.py --base <ruta_ps> --list-mail-domains
# → Emails.Body,Shop.Theme.Customeraccount   (lista completa, separada por comas)
```

### Flujo completo de emails (en orden)

```bash
# 1. Dominios necesarios en core para ESTE proyecto (varía por mail theme/módulos instalados)
DOMAINS=$(python3 ~/.claude/skills/ps-translate/scripts/scan.py --base <ruta_ps> --list-mail-domains)

# 2. Escanea + traduce + escribe cada dominio de email — con --core-domains "$DOMAINS"
#    (si un dominio como Shop.Theme.Customeraccount YA está traducido en el theme por otro
#    motivo, basta con copiar ese XLF ya existente a translations/<locale>/ con el nombre
#    sin puntos — no hace falta re-traducirlo desde cero)
python3 ~/.claude/skills/ps-translate/scripts/scan.py --base <ruta_ps> --lang <locale> \
  --include-emails --output /tmp/ps_emails_<locale>.json
# … traducir (Paso 2) …
python3 ~/.claude/skills/ps-translate/scripts/write_xlf.py --base <ruta_ps> --lang <locale> \
  --translations <json_traducido> --core-domains "$DOMAINS"

# 3. Limpiar caché — rm -rf var/cache/* NO es suficiente, hace falta el comando
php bin/console cache:clear

# 4. Generar las plantillas reales por idioma (esto NO pasa solo)
php bin/console prestashop:mail:generate <mail-theme> <locale> --overwrite
# <mail-theme> = valor de ps_configuration.PS_MAIL_THEME (normalmente "modern")
```

**Por qué el paso 4 es obligatorio y no opcional**: los layouts Twig son la FUENTE,
no lo que `Mail::send()` usa en producción — ese método lee los ficheros estáticos
generados en `mails/<locale>/`. Si ese idioma nunca se generó (típico cuando el
idioma se activó por API/DB en vez del flujo normal de Admin → Idiomas → Añadir),
esa carpeta ni existe, y ningún XLF por sí solo la crea.

**Verificación final** — vuelve a correr el scan tras generar; debe dar 0:

```bash
python3 ~/.claude/skills/ps-translate/scripts/scan.py --base <ruta_ps> --lang <locale> \
  --include-emails --output /dev/null   # total: 0 en stderr = todo cubierto
```

Y opcionalmente, barrido de texto plano sobre los HTML generados buscando frases
en inglés que se sepa deberían estar traducidas (grep simple, sin `<tags>`), como
capa extra de confianza antes de dar el email por completo.

---

## Qué dominios ignorar (admin, no críticos)

Por defecto el scanner omite `Admin.*`, `psgdpr`, `ps_themecusto`, `steasybuilder`,
`ps_facebook`, `ps_accounts`. Pasa `--include-admin` para incluirlos.

---

## Dónde se guardan las traducciones

Por defecto, siempre en el **child theme**:
`themes/<theme>/translations/<locale>/<DomainKey>.<locale>.xlf`

Nunca se modifican el core (`/translations/`) ni el tema padre — **excepto** los
dominios de email (`write_xlf.py --core-domains "..."`, ver sección de arriba),
que van a `translations/<locale>/<DomainKeySinPuntos>.<locale>.xlf` en la raíz del
core porque el traductor de generación de emails no lee el theme en absoluto. Es
la única excepción deliberada a esta norma, y está acotada a esos dominios
concretos — todo lo demás sigue yendo siempre al theme.

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
- Dos bugs de comparación ya corregidos (no deberían reaparecer, pero documentados
  por si algo similar vuelve a pasar): (1) un `|trans()` anidado dentro de los
  parámetros de otro (típico en emails, ver arriba) hacía que una regex simple
  capturase el dominio INTERIOR en vez del exterior — se resolvió balanceando
  llaves manualmente (`find_twig_trans_calls`) en vez de una regex de un solo paso.
  (2) el contenido de `<source>` en un XLF ya escrito está XML-escapado en disco,
  pero los strings detectados en plantillas (`used`) no lo están — comparar ambos
  tal cual hacía que CUALQUIER string con `<`, `>`, `"` o `&` (todo el HTML de los
  emails) se reportara como "sigue faltando" incluso después de traducirlo
  correctamente. Se corrigió desescapando en `load_translations()` antes de comparar.
