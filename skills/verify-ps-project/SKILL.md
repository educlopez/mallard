---
name: verify-ps-project
description: >
  Prove a change works in the real local PrestaShop storefront: bring Lando up,
  run a read-only doctor, drive home, category, product page, cart and checkout
  in a browser, capture evidence, clean up. Also keeps the project's feature map
  honest (maintenance pass). Use when the user says "verifica este cambio en
  local", "prueba el checkout", "demuestra que funciona", "haz una pasada de
  verificación", "revisa que el mapa de features sigue vigente", or before
  declaring done a theme, module or checkout change.
version: "0.1.1"
metadata:
  author: Eduardo Calvo
---

# verify-ps-project

Drives the real storefront the way a customer does and leaves proof. Adapted from pstack's `create-verification-skill` / `maintain-verification-skill`. It does not set Lando up (`lando-setup`), fix Lando (`lando-check`) or edit product code.

Write everything for an agent that reads it cold: real selectors and URLs from this project, no placeholders.

## 1. Launch

1. Run `lando-check`. Ready means the shop answers `200` on its real local URL, not that a container exists.
2. Confirm the DB is usable: `lando ssh -c "php bin/console doctrine:query:sql 'SELECT 1'"` (PS 8/9) or open the BO login.
3. Pick the browser tool already available: `agent-browser` skill, Chrome DevTools MCP or Claude in Chrome. Do not install Playwright just for this.

## 2. Doctor (read-only, run first whenever something looks off)

- Right project: `.lando.yml` name matches the repo directory.
- Right URL: the shop does not redirect to the production domain (`/ps-url` if it does).
- Right theme: the active child theme is the one under test.
- Port and containers are ours: never drive an instance you did not start or confirm.
- Child theme CSS is built: the compiled file (usually `themes/<child>/assets/css/custom.css`) exists and is not empty. A 404 or 0-byte file here means the CSS build was never run (`ps-css-build`).

## 3. Drive

Cover the flows a customer touches, in order. Prefer stable handles (`data-*`, ids, route paths, ARIA labels) over coordinates.

| Flow | Check |
|---|---|
| Home | Loads, no console errors, no 404 assets |
| Category | Listing renders, sort and filter change the list |
| Product | Gallery, price, combination switch, add to cart |
| Cart | Line appears, quantity change updates totals |
| Checkout | Reaches each step; demo user via `/ps-customer` or `ps-demo-user`; stop before real payment |

After clicking "Añadir al carrito", wait until the cart counter changes (or the modal appears) before navigating: the add is AJAX, and navigating early leaves an empty cart that looks like a failure.

Side effects count: check the cart row, order record or email in Mailhog alongside what is visible. No test-only endpoints.

## 4. Evidence

Save to `.verify/evidence/<YYYY-MM-DD-HHMM>/` in the project: one screenshot per step plus `console.txt` with errors. Some browser MCPs only write files inside their configured workspace roots and refuse the project path; then take the screenshot without `filePath`, use the path the tool returns, and copy that file into `.verify/` with `cp`. Keep `.verify/` out of git via `.git/info/exclude`, not the shared `.gitignore`. Report one line per flow: pass, fail or not driven, with the evidence path.

## 5. Cleanup

Stop only what this run started. Never kill by process name. Remove scratch data created for the run (demo customer, test cart). Evidence stays.

## 6. Feature map and maintenance

Keep `.verify/features/<feature>.md` per user-facing feature (3-5 to start): what it is, how to reach it, how to drive it, what proves it works. Seed it from the routes and menus you actually see.

Maintenance pass (when asked, not on a schedule):

1. Read the map index; fix missing, duplicate or dead entries.
2. For each feature, read the source entry points, then drive it once live.
3. Classify each difference: wrong description is doc drift (fix the map); working behavior the flow cannot drive is a harness gap (fix it); broken behavior is a product bug (report it, do not hide it in docs).
4. Show the proposed corrections and wait for the user's OK before editing. Never edit product code in this pass.
