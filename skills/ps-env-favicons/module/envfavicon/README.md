# Env Favicon

A different favicon on **local** and **staging** so you can tell the tabs apart at a glance.
**Production never changes**: it keeps the native PrestaShop favicon (Design > Theme & Logo > Favicon).

![Env Favicon](logo.png)

**Author:** Eduardo Calvo · **License:** [AFL-3.0](LICENSE.md) · **Compatible with:** PrestaShop 1.7.7 → 9.x · **PHP:** 7.2+ with GD

## Features

- **One logo in, everything out.** Upload your logo once; the module crops the mark (it can drop the ®
  and a text next to it) and builds a coloured badge per environment plus the production `.ico` (16–64 px).
- **Domains in the back office.** A host that contains any fragment of the *Local* or *Staging* list gets that
  badge; anything else is production. Staging domains are never hard-coded (`pre-testing.com`, `ecommercloud`, …).
- **Theme independent.** A single output hook swaps the `<link rel="icon">` tags of the `<head>`. No template
  overrides, no copy of the parent theme's `head.tpl`, no JS, no extra requests besides the 64 px badge.
- **Native favicon, one click.** "Load the production favicon as the PrestaShop favicon" writes
  `img/favicon.ico` and `img/favicon-<shop>.ico` and sets `PS_FAVICON`, exactly like the back office upload.
- **Safe by default.** Unknown hosts are production. The host is only compared, never printed. Uploads go through
  `ImageManager::validateUpload` and are stored under a fixed name.

## Install

1. `scripts/build-zip.sh v0.2.0` and upload `build/envfavicon-v0.2.0.zip` in *Modules > Module Manager*, or copy
   the folder to `modules/envfavicon` and install it.
2. Open the module, upload the logo, review the domains and colours, **Save and regenerate**.
3. On the server that will be production, press **Load the production favicon as the PrestaShop favicon**.

## How it decides

| Host contains… | Environment | Favicon |
|---|---|---|
| any *Local* fragment (default `lndo.site, localhost, 127.0.0.1, .test, .local`) | Local | green badge |
| any *Staging* fragment (default `pre-testing.com, ecommercloud`) | Staging | gold badge |
| nothing | Production | native PrestaShop favicon |

## Translations

Source strings are English. `php scripts/gen-translations.php` rebuilds `translations/es.php` from
`scripts/i18n.json` and fails if a string is missing from the dictionary.

## Notes

- The generated files (`views/img/source.*`, `favicon-*`, `mark-white.png`) are not versioned: each environment
  builds its own from the back office.
- Browsers cache favicons aggressively. Use a hard reload or a private window; the module appends `?v=<n>`
  and bumps it every time it regenerates.
