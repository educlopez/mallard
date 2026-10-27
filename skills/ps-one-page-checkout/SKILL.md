---
name: ps-one-page-checkout
description: >
  How to enable, verify, style and translate the native PrestaShop 9.2 one-page
  checkout (module ps_onepagecheckout) on a project, with compatible versions
  and the gotchas found on Pies Sanos. Use when the user says "one page checkout",
  "checkout en una página", "activar el checkout de PrestaShop", "OPC",
  "ps_onepagecheckout", or asks to style/translate the PS9 checkout.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# PrestaShop 9.2 native one-page checkout (ps_onepagecheckout)

Source: PrestaShop announcement "Introducing One Page Checkout in PrestaShop 9.2"
(build.prestashop-project.org/news/2026/introducing-one-page-checkout/, 22 Jul 2026).
Dev docs: devdocs.prestashop-project.org/9/modules/checkout/ (module-developers, theme-developers).

## Compatible versions (checked on Pies Sanos, 2026-10)

| Piece | Version | Note |
|---|---|---|
| PrestaShop | 9.2.0 or newer | Module declares `ps_versions_compliancy` min 9.2.0. Not available on 8.x or 9.0/9.1. |
| Module `ps_onepagecheckout` | 0.6.7 (beta) | Ships inside PS 9.2; installed and active, but the layout is OFF by default. |
| Theme | Hummingbird (2.1.2) and children | Works out of the box: the module ships its own templates and takes over via `displayOverrideTemplate`. |
| Classic theme | Not supported by default | Needs template overrides of the module. |
| Core hook | `actionCheckoutBuildProcess` | Added in PS 9.2 (PR #41047). |

Announced as beta: re-check the module version and release notes before promising it to a client.

## Enable

Pick one. Config key: `PS_ONE_PAGE_CHECKOUT_ENABLED` (module installs it as 0).

1. Back office: Design > Checkout, choose "One-page checkout" layout, save. Switch back to the "Four-page checkout" the same way.
2. Console (servers, scripted):
   ```bash
   php bin/console prestashop:config set PS_ONE_PAGE_CHECKOUT_ENABLED --value=1
   php bin/console cache:clear
   php bin/console prestashop:config get PS_ONE_PAGE_CHECKOUT_ENABLED
   ```
   The `--value=` form is required (`set KEY 1` fails with the usage text). On Lando prefix with `lando`.
3. SQL (replayable, ignored by git unless `git add -f`): `UPDATE <prefix>configuration SET value='1', date_upd=NOW() WHERE name='PS_ONE_PAGE_CHECKOUT_ENABLED';` then clear cache. Keep it as `themes/<child>/sql/one-page-checkout.sql` so prod cut-over can replay it.

Multistore: the setting is per store (use `--shopId`), there is no global switch.
Check the module is enabled first: `lando mysql <db> -e "select name,active,version from <prefix>module where name='ps_onepagecheckout'"`.

## Verify

Add an in-stock product (`stock_available.quantity > 0`) to the cart, open `/carrito?action=show` (or `?controller=order`): one page with Contact (guest only asks for email), Delivery address, Delivery method, Payment, and the summary on the side. Shipping and payment only appear after the address is complete.

## Gotchas found

- **Spanish is missing.** The module ships en-US, fr-FR and pl-PL only; the repo-level `translations/es-ES` is git-ignored and has ~33 of 100 strings. Put a full `modules/ps_onepagecheckout/translations/es-ES/ModulesOnepagecheckoutShop.es-ES.xlf` (same ids as the en-US file, tracked in git) and clear the cache.
- **Default country "Francia"** on the address form: check `PS_COUNTRY_DEFAULT` and the active countries.
- **Styling hooks** (body id `checkout`): `.checkout-grid`, `.checkout-grid__aside-wrapper`, `.one-page-checkout__section` (one per step, title `.one-page-checkout__title`), `.one-page-checkout__submit`, `.opc-*` for dynamic parts. Number the steps with a CSS counter on `.tab-content`.
- **Reassurance blocks** in cart and checkout come from hook `displayReassurance` (`.blockreassurance_product` markup), so style that once for product, cart and checkout.
- **Social login** (iqitsociallogin hook `displayPersonalInformationTop`) renders outside the step cards; override its `checkout.tpl` in the child theme.
- **Theme CSS cache:** after editing the child `custom.css`, delete `assets/cache/theme-*.css` and clear the cache or the old CSS is served.
- **Payment modules:** with ps_checkout inactive no payment shows; the OPC keeps the classic hooks and adds JS events (`opcCarriersUpdated`, `opcPaymentMethodSelected`, `opcDeliveryAddressSelected`). Carrier/payment modules that inject markup must re-render after the AJAX refresh.
- Anything on the Pies Sanos repo: theme `pies-sanos-child`, MR !17.

## Rollback

Set `PS_ONE_PAGE_CHECKOUT_ENABLED` to 0 (or pick "Four-page checkout" in Design > Checkout) and clear the cache. No data migration involved.
