# Changelog

All notable changes to this project are documented in this file.

## [0.2.0] - 2026-11-13

### Added

- Back office page redesigned: logo drop zone with the detected mark, one card per
  environment with a live browser-tab preview, colour pickers and an "active" switch.
- Spanish translation (`translations/es.php`, generated from `scripts/i18n.json`).
- Module icon.

### Changed

- Author and licence (Eduardo Calvo, AFL-3.0); source strings are now English.

## [0.1.0] - 2026-11-13

### Added

- Output hook (`actionOutputHTMLBefore`) that replaces the `<link rel="icon">` tags on local
  and staging with a coloured badge; production keeps the native PrestaShop favicon.
- Domains per environment, stored in the database and editable in the back office.
- One uploaded logo generates the badges and the production `.ico` with GD, and the
  production favicon can be loaded as the native PrestaShop favicon.
