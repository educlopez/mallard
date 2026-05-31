---
source_url: https://devdocs.prestashop-project.org/9/modules/concepts/controllers/
ps_version: [8, 9]
ingested: 2026-04-18
section: Modules > Controllers (admin + front)
topic: modules
tags: [prestashop, modules, controllers, admin, front]
---

[Improve this page](https://github.com/PrestaShop/docs/edit/9.x/modules/concepts/controllers/_index.md "Improve this page") [How](/9/contribute/documentation/how)

In the previous chapters, we saw how to add content in the existing pages of the front and back office with hooks & widgets.

As soon as a module needs to implement more than a configuration page, building controllers will offer a dedicated space for your features. In a MVC architecture, a Controller manages the synchronization events between the View and the Model, and keeps them up to date. It receives all the user events and triggers the actions to perform.

If an action needs data to be changed, the Controller will “ask” the Model to change the data, and in turn the Model will notify the View that the data has been changed, so that the View can update itself. Module controllers will behave like the core ones thanks the class inheritance:

![](img/module-controllers.png)

Controller classes inheritance

The core classes can be found in the `/classes/controller` folder.

[Next: Front controllers ›](https://devdocs.prestashop-project.org/9/modules/concepts/controllers/front-controllers/ "Front controllers")

## Related in KB

- [[module-class|Module class reference]]
- [[../version-migration/architecture-modern|Modern architecture (Symfony)]]
- [[../version-migration/architecture-legacy|Legacy architecture]]
