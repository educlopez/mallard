---
title: Custom module migration — PrestaShop 8 → 9
source: extrapolation from devdocs.prestashop-project.org + dogfood Q4 (2026-05-19)
ps_version: [8, 9]
tags: [migration, modules, symfony, php-8, custom]
last-updated: 2026-05-20
---

# Custom Module Migration: PrestaShop 8 → 9

> [!info] Scope
> Walkthrough específico para migrar un **módulo de tercero / custom** de PS 8 a PS 9. La doc oficial de PrestaShop cubre el patrón core BO migration (Symfony) pero no este caso concreto — esta nota lo materializa.

## TL;DR

- **PHP 8.1+ obligatorio** en PS 9.
- **Hooks y `Module` base class compatibles** — sin cambios obligatorios.
- **FO ModuleFrontController**: déjalo legacy. PS 9 no migra el FO a Symfony.
- **BO ModuleAdminController**: sigue funcionando legacy. Migra a Symfony **solo si vale la pena**.
- **Composer + PSR-4** recomendado si declaras services o controllers Symfony nuevos.

---

## Qué cambia obligatorio en PS 9

| Área | PS 8 | PS 9 obligatorio | Por qué |
|---|---|---|---|
| **PHP** | 7.4+ | **8.1+** | Runtime. Tipos nullable, readonly, enums disponibles. |
| **Composer / autoload** | opcional | `composer.json` con PSR-4 recomendado | Symfony 6 lo asume; sin esto los services nuevos no se cargan. |
| **`config.xml`** | igual | igual | Sin cambios estructurales. |
| **`Module::install()` hooks** | `registerHook('displayX')` | igual | Compatible. No cambia. |
| **Module class base** | `Module` / `PaymentModule` / `CarrierModule` | igual | Compatible. |
| **Symfony DI (`config/services.yml`)** | opcional | **obligatorio si declaras servicios o controllers Symfony nuevos** | Symfony 6 + autowiring estricto. Revisa `bind`, `arguments` explícitos. |

---

## Controllers — decisión por tipo

### FO `ModuleFrontController` (`controllers/front/*.php`)

✅ **Déjalo como legacy.** PS 9 NO migra el FO a Symfony. Sigue siendo Smarty + `FrontController`. **Cero cambios obligatorios**.

Solo toca si tu código usa APIs deprecadas. Revisa:
- `$this->context->smarty->assign()` → sigue OK.
- `Tools::getValue()` → sigue OK.
- Tipos nullable implícitos → ver sección PHP 8.1 más abajo.

### BO `ModuleAdminController` (`controllers/admin/*.php`)

⚠️ **Sigue funcionando** en PS 9 — el patrón legacy convive con Symfony. Ver `architecture-introduction.md`:

> BO controllers can belong to either the Legacy subsystem or to the PrestaShop Bundle.

**Decisión por escenario:**

| Escenario | Recomendación |
|---|---|
| Form de config simple, listado básico, CRUD con `HelperForm`/`HelperList` | ✅ Déjalo legacy |
| Grid moderno con filtros, sort, bulk actions | 🔄 Migra a Symfony |
| Twig + Symfony Forms (validación rica) | 🔄 Migra a Symfony |
| Integración con BO API (VueJS pages) | 🔄 Migra a Symfony |

### Si migras un BO controller a Symfony

Checklist (de `migration-guide-bo.md` adaptado):

1. Crear controller bajo `src/Controller/Admin/` en tu módulo.
2. Declarar routing en `config/routes.yml` del módulo (**no** en `PrestaShopBundle/Resources/config/routing/` — eso es core).
3. Symfony Form Types por cada form.
4. Templates Twig en `views/templates/admin/`.
5. Re-enganchar hooks que llamabas desde la pantalla legacy.
6. Borrar el viejo `controllers/admin/AdminTuController.php` y su tpl Smarty.

> [!info] PS 9: `PrestaShopAdminController`
> En PS 9, la base recomendada para BO Symfony es `PrestaShopAdminController` (sustituye a `FrameworkBundleAdminController` de PS 8). API similar — adapta los `extends` al hacer la migración.

---

## Services (`config/services.yml`)

| Aspecto | PS 8 | PS 9 |
|---|---|---|
| **Obligatoriedad** | Opcional. Muchos módulos PS 8 instanciaban todo en el `Module` class. | Si declaras services o controllers Symfony nuevos → **obligatorio**. |
| **Autowire** | `autowire: false` por defecto | `autowire: true` + `autoconfigure: true` recomendados |
| **Legacy bridge** | `@=service('prestashop.adapter.legacy.context')` | Idem — Adapter pattern sigue activo |
| **Renombres** | — | ⚠️ Algunos services core fueron renombrados entre 8 y 9. Grep `services.yml` core en PS 9 antes de asumir nombres viejos. |

---

## Hooks

✅ **Compatibles entre PS 8 y PS 9** — la API de hooks **no rompe**. `registerHook()`, `hookDisplayX()`, `hookActionX()` siguen igual.

### Lo que SÍ vigilar

- **Hooks BO de pantallas migradas a Symfony**: pasan a dispararse desde Twig con `{{ renderhook('hookName') }}` en vez de Smarty `{hook h='hookName'}`. Si tu módulo se engancha a un hook de pantalla migrada, **sigue funcionando** porque `Hook::exec()` es agnóstico — pero confirma que el hook sigue existiendo en PS 9.
- **Hooks renombrados o eliminados**: algunos hooks legacy se renombraron en PS 9. Verifica tras instalar:
  ```sql
  SELECT name FROM ps_hook WHERE name LIKE '%TuHook%';
  ```
- **Display hooks en FO**: cero cambios.

---

## Qué puedes dejar como legacy (resumen)

- ✅ Front controllers Smarty (`controllers/front/*`).
- ✅ Tpl Smarty front (`views/templates/front/`, `views/templates/hook/`).
- ✅ BO controllers legacy con `HelperForm`/`HelperList` si la pantalla es modesta.
- ✅ `ObjectModel` custom (`classes/MiEntidad.php`).
- ✅ Hooks registrados (lifecycle igual).
- ✅ Overrides de clases legacy (sistema sigue activo, aunque desaconsejado por upgrade-fragility).

---

## Lo que SÍ debes tocar mínimo

1. **`composer.json`**: bump PHP a `^8.1`, revisa deps no compatibles.
2. **`config.xml`**: `<version>` y `<ps_versions_compliancy><min>9.0.0</min><max>9.99.99</max></ps_versions_compliancy>` si restringes a PS 9.
3. **Deprecaciones PHP 8.1**:
   - Nullable params implícitos: `function f(int $x = null)` → `?int $x = null`.
   - `each()` eliminado.
   - `strftime()` deprecated → usa `IntlDateFormatter`.
   - Tipos de propiedades en clases (PHP 7.4+ ya, pero 8.1 endurece).
4. **Tests humo tras instalar en PS 9**:
   - `install` / `uninstall` sin warnings.
   - Hooks que registras se disparan.
   - Página BO si la tiene.
   - Página FO si la tiene.
   - Multi-idioma si aplica.

---

## Trade-off final

Migrar el BO controller a Symfony **solo si vale la pena**:
- Pantalla de config simple → **déjalo `ModuleAdminController`**.
- Vas a tocar grids, validación rica, o integración con BO API → **vale la inversión**.

---

## Workflow recomendado

```
1. Setup PS 9 local con tu módulo PS 8 instalado.
2. Ajustar composer.json y config.xml (PHP, ps_versions_compliancy).
3. Smoke test: install/uninstall + flujos críticos.
4. Fix PHP 8.1 deprecations (nullable params, each, etc.).
5. Verificar hooks (SQL query en ps_hook).
6. Decidir caso a caso si BO controllers se migran a Symfony.
7. Si se migran: checklist Symfony BO arriba.
8. Tests humo finales en PS 9 — visual y funcional.
```

---

## Referencias

- [[architecture-introduction]] — modelo de capas Legacy/Core/Adapter/Symfony Bundle.
- [[architecture-legacy]] — el subsystem legacy en PS 9.
- [[architecture-modern]] — Symfony Bundle subsystem.
- [[migration-guide-bo]] — checklist oficial BO core migration (extrapolable a módulos).
- [[modules/controllers]] — controllers de módulos.
- [[upgrade-faq]] — FAQ general de upgrade.
