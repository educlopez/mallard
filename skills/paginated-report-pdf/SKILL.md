---
name: paginated-report-pdf
description: >
  Generates a client-facing, document-style paginated PDF report (cover page,
  numbered sections, one page per topic/site/item, tables, callout boxes,
  numbered step lists) from an HTML source, using headless Chrome — no
  external PDF library needed. Use this whenever the user asks for a "PDF
  report", "informe en PDF", "documento para el cliente", "PDF con nuestra
  marca/logo", or wants an existing HTML dashboard/artifact turned into a
  proper paginated document instead of a browser screenshot or a raw print of
  the dashboard itself. Also use when asked to add a persistent logo/header to
  every page of a PDF, or to regenerate/update a PDF that was built this way
  before. Works for ANY brand/client — always run the brand intake step first,
  never assume a specific client's colors or logo.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# Paginated report PDF (HTML → Chrome headless → PDF)

Produces a real, linear document (cover, numbered sections, per-item pages)
rendered by headless Chrome from a self-contained HTML file — not a
screenshot, not a raw print of an interactive dashboard. Brand-agnostic: the
skill asks for the brand's visual identity before writing a single line of
HTML, so the same instructions produce a correctly-branded PDF for whichever
client/project you're in.

## When to use this vs. other options

- **Use this** when the deliverable is a real *document*: cover page,
  sections, tables, a page per item, meant to be read start-to-finish or
  archived/emailed as a PDF.
- **Don't use this** to just "export" an existing interactive dashboard (tabs,
  sidebar nav) as-is — that produces a broken, unreadable PDF (nav elements
  repeated, tab content hidden). Write a *separate*, purpose-built paginated
  HTML file instead. It's fine — even expected — for this file to reuse
  content/wording from the dashboard, restructured into linear document flow.
- Never reach for a heavier toolchain (wkhtmltopdf, Puppeteer, a PDF library)
  unless headless Chrome is confirmed unavailable — it almost always already
  is (see prerequisite check below). If the task is PDF *manipulation*
  (merge/split/extract/forms/OCR) rather than *authoring a new document*,
  that's a different job — don't force this skill onto it.

## Step 0 — brand intake (always do this first)

Before writing any HTML, establish the brand's visual identity. Resolve each
of these, in order of preference:

1. **Already evident from the current project** — an existing design system,
   `theme.css`/tokens file, a previous PDF built with this skill in the same
   repo, or a logo file already in the project (`assets/`, `public/`,
   `static/`). If found, reuse it and say so — don't re-ask for what's already
   on disk.
2. **Given directly by the user in this conversation** — a color, a font
   name, a path/URL to a logo. Use it as stated.
3. **Ask.** If neither of the above resolves it, ask (a single
   `AskUserQuestion` covering all of the below, or a plain question if that
   tool isn't available):
   - **Logo**: file path, URL, or "no logo for this one".
   - **Primary color** (hex) — used for callout borders, pills, accent text.
     Default to a neutral dark gray/near-black if the user has no preference
     and there's truly no brand to match.
   - **Font family** — a real font name if the brand has one; otherwise the
     system-font stack in Step 1 is a fine default, say so explicitly instead
     of silently picking something.
   - **Document title / client name** — goes in the cover page and the
     per-page header.

Never hardcode a specific client's brand (colors, logo, name) as a default —
every value in the HTML below is a placeholder resolved from this step:
`{brand_name}`, `{primary_color}`, `{font_family}`, `{logo_data_uri}`.

## Prerequisite check

```bash
ls "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome"
```

If missing, look for `chromium`/`google-chrome` on PATH, or ask the user.
Don't install a new dependency for this.

## Step 1 — write the paginated HTML source

One `.html` file, self-contained (no external CSS/font/image URLs —
everything inline or base64, since the file may be printed with no network).
Structure:

```html
<!DOCTYPE html>
<html lang="es">
<head>
<meta charset="utf-8">
<title>{brand_name} — {document_title}</title>
<style>
  :root { --brand-primary: {primary_color}; }
  @page { size: A4; margin: 22mm 18mm 20mm 18mm; }
  * { box-sizing: border-box; }
  body {
    font-family: {font_family}, -apple-system, "Helvetica Neue", Helvetica, Arial, sans-serif;
    -webkit-print-color-adjust: exact; print-color-adjust: exact;
  }
  .page { page-break-after: always; }
  .page:last-child { page-break-after: auto; }
</style>
</head>
<body>

<section class="page"> ... cover ... </section>
<section class="page"> ... section 2 ... </section>
<!-- one <section class="page"> per logical page; content that overflows a page
     just flows onto the next physical page automatically, no extra markup needed -->

</body>
</html>
```

Key CSS pieces worth building (reference tokens like `var(--brand-primary)`,
never a hardcoded hex, so the same structure repaints for any brand):
- `.callout` / `.callout.strong` — colored left-border box for a key finding
  or warning, border color = `var(--brand-primary)`.
- `.stat-grid` / `.stat .num` — big-number summary stats row.
- `.pill.good` / `.pill.warn` — small colored status badges inside tables.
- `.steps li` + `.step-circle` — numbered/checkmarked circular steps (a plan
  or a checklist), circle background = `var(--brand-primary)`.
- `.kv-table` — a label/value table (first column bold+uppercase+narrow) for a
  spec sheet.
- `.item-badge`, `.item-head`, `.meta-pills` — a per-item detail-page header
  pattern (name, URL/ref, status badge, spec chips) — reuse this whenever the
  report has "one page per X" (per site, per store, per module, per
  incident).

Always include both `-webkit-print-color-adjust: exact;` and
`print-color-adjust: exact;` on `body` — without them Chrome's print engine
can wash out background colors on callout boxes.

## Step 2 — embed the logo (or any image) as base64, once

```bash
curl -sL -o /tmp/logo.png "<logo-path-or-url-from-step-0>"
base64 -i /tmp/logo.png | tr -d '\n' > /tmp/logo.b64
```

Then inline it: `<img src="data:image/png;base64,$(cat /tmp/logo.b64)" alt="{brand_name}">`.
If the user said "no logo", skip this step entirely and drop the `<img>` from
the header template in Step 3 — don't fabricate a placeholder logo.

Once embedded in one file, extract it back out for reuse elsewhere with:

```bash
grep -o 'data:image/png;base64,[A-Za-z0-9+/=]*' source.html | head -1
```

— no need to re-fetch/re-encode the image for every file that needs it, or
for a later PDF for the same brand.

## Step 3 — logo/header on every single page (not just the cover)

Chrome's CLI (`--headless --print-to-pdf`) does **not** support a custom
per-page header/footer template (that's a Puppeteer-only API,
`page.pdf({headerTemplate})`) — `--print-header-footer` only gives Chrome's
own plain URL/date/title footer, not a custom image. The reliable fix: put
the logo **in the HTML content of every page section**, not in a
browser-managed header:

```css
.page-header {
  display: flex; align-items: center; justify-content: space-between;
  margin-bottom: 18px; padding-bottom: 10px; border-bottom: 1px solid #eeeef1;
}
.page-header img { height: 14px; width: auto; }
```

```html
<section class="page">
  <div class="page-header">
    <img src="data:image/png;base64,..." alt="{brand_name}">
    <span class="doc-title">{document_title}</span>
  </div>
  ... rest of page content ...
</section>
```

Add this `<div class="page-header">` as the first child of every
`<section class="page">` except the cover (which usually has its own
full-size logo treatment already). If there's no logo (Step 0/2 said so),
keep just the `<span class="doc-title">`. If there are many pages, script the
insertion rather than hand-editing each one — see the one-liner pattern
below.

## Step 4 — generate the PDF

```bash
"/Applications/Google Chrome.app/Contents/MacOS/Google Chrome" \
  --headless --disable-gpu --no-pdf-header-footer \
  --print-to-pdf="/absolute/path/to/output.pdf" \
  "file:///absolute/path/to/source.html"
```

- `--no-pdf-header-footer` suppresses Chrome's own plain-text header/footer
  (URL, date, page number) — always pass it, since the HTML already carries
  its own header per Step 3.
- The `file://` URL must be an **absolute** path.
- Chrome prints a lot of unrelated `updater`/`crashpad` log noise to stderr —
  harmless, ignore it; check the actual output file size/existence instead of
  parsing stderr for success.
- This runs Chrome fully headless, no window, no user interaction — safe to
  run repeatedly while iterating.

## Step 5 — verify before calling it done

Read a few pages of the generated PDF back (the `Read` tool renders PDF pages
as images) — check: logo present per page (if one was requested), tables
didn't get cut mid-row across a page break in an ugly way, colored boxes
rendered with the brand color (not washed out or defaulted to a generic
color), no leftover placeholder text (`{brand_name}` etc. never survives into
the actual output — if you see it, Step 0 was skipped or a substitution was
missed). Iterate on the HTML and regenerate — this loop is fast (a few
seconds per Chrome invocation).

## Iterating on structure (adding/removing pages, renumbering sections)

When inserting a new page in the middle of a numbered sequence (e.g. a
"2 · Section name" eyebrow scheme), grep for all `section-eyebrow` occurrences
first and fix the numbering in one pass — it's easy to insert a new page and
forget every later page still says the old number.

When batch-inserting the same snippet (like the per-page header in Step 3)
into many `<section class="page">` blocks at once, do it with a small Python
script that splits on `<section class="page` and inserts right after each
opening tag's `>` — faster and less error-prone than editing each section by
hand:

```python
import re
with open(path, encoding='utf-8') as f:
    c = f.read()
m = re.search(r'data:image/png;base64,[A-Za-z0-9+/=]+', c)
data_uri = m.group(0) if m else None  # None if this brand has no logo
header_html = (
    f'\n  <div class="page-header"><img src="{data_uri}" alt="{{brand_name}}" /><span class="doc-title">{{document_title}}</span></div>'
    if data_uri else
    f'\n  <div class="page-header"><span class="doc-title">{{document_title}}</span></div>'
)
parts = c.split('<section class="page')
new_parts = [parts[0], parts[1]]  # keep head + cover page untouched
for seg in parts[2:]:
    idx = seg.index('>') + 1
    new_parts.append(seg[:idx] + header_html + seg[idx:])
c = '<section class="page'.join(new_parts)
with open(path, 'w', encoding='utf-8') as f:
    f.write(c)
```

## Keeping the PDF in sync with other artifacts

If this PDF is a paginated version of an existing interactive HTML dashboard,
the two will drift if only one gets edited. When a fact changes: update the
dashboard's per-item section AND the matching PDF-source page AND regenerate
the PDF in the same pass — don't defer the PDF regen "for later", it's cheap
(Step 4 is one command) and skipping it is exactly how documents go stale.

## Prior art

This skill's technique (HTML → Chrome headless → PDF, per-page HTML header
since Chrome CLI has no custom header template) was worked out end-to-end on
a real client deliverable — a security-audit report with a cover page,
per-site incident pages, callout boxes, and a persistent logo header. That
file is a legitimate reference for the CSS patterns in Step 1 if you want to
see them fully worked, but treat every brand value in it as
project-specific — never copy its actual colors/logo into a different
client's report.
