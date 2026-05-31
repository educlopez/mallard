---
tags: [kb, panda, modules, index]
auto_generated: true
---

# Panda modules — Index

54 Panda-specific modules extracted from source (`panda-theme/panda/dependencies/modules/st*`, Panda v2.9.2).

> [!important] stthemeeditor es el hub de configuración
> La mayoría de estos módulos **NO tienen su propio BO screen**. Se configuran leyendo `Configuration::get('STSN_*')` desde `ps_configuration`, con las keys escritas por [[stthemeeditor-deep-dive|stthemeeditor]] (38 tabs, 529 keys). Sólo ~8 módulos exponen admin propio. Si Eduardo pregunta *"dónde configuro X"* → primero consulta stthemeeditor.

| Module | Category | Hooks | Configurable | Version | Public doc? |
|---|---|---:|:---:|---|:---:|
| [[stblog]] | `blog` | 5 | yes | 1.7.1 | — |
| [[stblogarchives]] | `blog` | 5 | no | 1.1.3 | — |
| [[stblogblockcategory]] | `blog` | 1 | yes | 1.2.0 | — |
| [[stblogcomments]] | `blog` | 8 | yes | 1.1 | — |
| [[stblogeditor]] | `blog` | 1 | yes | 1.0.0 | — |
| [[stblogfeaturedarticles]] | `blog` | 1 | yes | 1.0.2 | — |
| [[stbloglinknav]] | `blog` | 1 | no | 1.0 | — |
| [[stblogrecentarticles]] | `blog` | 1 | yes | 1.0.2 | — |
| [[stblogrelatedarticles]] | `blog` | 5 | yes | 1.0 | — |
| [[stblogsearch]] | `blog` | 3 | yes | 1.0 | — |
| [[stblogtags]] | `blog` | 2 | yes | 1.0 | — |
| [[stbestsellers]] | `product` | 2 | yes | 1.1.1 | — |
| [[stcompare]] | `product` | 5 | yes | 1.0.1 | — |
| [[stcountdown]] | `product` | 2 | yes | 1.0 | — |
| [[stfeaturedcategories]] | `product` | 5 | yes | 1.1.8 | — |
| [[stfeaturedslider]] | `product` | 1 | yes | 1.1.8 | — |
| [[sthomenew]] | `product` | 1 | yes | 1.7.9 | — |
| [[sthoverimage]] | `product` | 3 | yes | 1.0 | — |
| [[stlovedproduct]] | `product` | 10 | yes | 1.0.3 | [[docs/module-love-buttons\|yes]] |
| [[stpagebanner]] | `product` | 7 | yes | 1.7.0 | — |
| [[stproductcategoriesslider]] | `product` | 8 | yes | 1.7.3 | — |
| [[stproductcomments]] | `product` | 12 | yes | 1.0.5 | — |
| [[stproductlinknav]] | `product` | 2 | no | 1.0 | — |
| [[strelatedproducts]] | `product` | 1 | yes | 1.1.0 | — |
| [[stspecialslider]] | `product` | 2 | yes | 1.1.0 | — |
| [[stviewedproducts]] | `product` | 3 | yes | 1.2.4 | — |
| [[stwishlist]] | `product` | 7 | yes | 1.0.1 | [[docs/module-love-buttons\|yes]] |
| [[stmegamenu]] | `navigation` | 19 | yes | 2.0.8 | [[docs/module-megamenu\|yes]] |
| [[stmultilink]] | `navigation` | 12 | yes | 1.9.4 | — |
| [[stsearchbar]] | `navigation` | 7 | yes | 1.6.7 | — |
| [[stsidebar]] | `navigation` | 7 | yes | 1.1.0 | — |
| [[stswiper]] | `navigation` | 9 | yes | 2.2.2 | — |
| [[stbanner]] | `content` | 8 | yes | 1.9.0 | [[docs/module-advanced-banner\|yes]] |
| [[stbrandsslider]] | `content` | 4 | yes | 1.3.7 | — |
| [[steasycontent]] | `content` | 7 | yes | 1.7.0 | [[docs/module-advanced-custom-content\|yes]] |
| [[stfblikebox]] | `content` | 2 | yes | 1.0.7 | — |
| [[stinstagram]] | `content` | 2 | yes | 1.1.3 | [[docs/module-advanced-social-feed\|yes]] |
| [[stnewsletter]] | `content` | 10 | yes | 1.4.6 | — |
| [[stnotification]] | `content` | 3 | yes | 1.1 | [[docs/module-notification\|yes]] |
| [[stqrcode]] | `content` | 3 | yes | 1.0 | — |
| [[stsocial]] | `content` | 4 | yes | 1.0.3 | — |
| [[ststickers]] | `content` | 8 | yes | 1.1.3 | [[docs/module-stickers-block\|yes]] |
| [[sttags]] | `content` | 2 | yes | 1.2.2 | — |
| [[sttwitterembeddedtimelines]] | `content` | 2 | yes | 1.0 | — |
| [[stvideo]] | `content` | 6 | yes | 1.1.1 | — |
| [[stcurrencyselector]] | `helpers` | 3 | yes | 0.3.3 | — |
| [[stcustomersignin]] | `helpers` | 5 | yes | 0.3.5 | [[docs/module-customer-sign-in\|yes]] |
| [[stlanguageselector]] | `helpers` | 2 | yes | 1.3.4 | — |
| [[stoverride]] | `helpers` | 0 | no | 1.0.0 | — |
| [[stowlcarousel]] | `helpers` | 8 | yes | 2.2.2 | — |
| [[stshoppingcart]] | `helpers` | 4 | yes | 1.5.8 | — |
| [[stthemeeditor]] | `helpers` | 5 | yes | 2.9.2 | — |
| [[stupgrader]] | `helpers` | 0 | no | ? | — |
| [[stxlfgenerator]] | `helpers` | 0 | yes | 1.0.0 | — |

## Stats

- **Total modules:** 54
- **blog:** 11
- **product:** 16
- **navigation:** 5
- **content:** 13
- **helpers:** 9
- **Configurable (BO screen):** 49
- **Zero hooks (library/helper modules):** 3
- **Total LOC:** 166,967

## Top 5 most complex modules by LOC

- [[stthemeeditor]] — 25,271 LOC · 153 files · 5 hooks
- [[steasycontent]] — 16,048 LOC · 182 files · 7 hooks
- [[stblog]] — 12,796 LOC · 118 files · 5 hooks
- [[stproductcomments]] — 11,737 LOC · 103 files · 12 hooks
- [[stmegamenu]] — 6,879 LOC · 45 files · 19 hooks

## Top 5 by hook count

- [[stmegamenu]] — 19 hooks · 6,879 LOC
- [[stmultilink]] — 12 hooks · 3,354 LOC
- [[stproductcomments]] — 12 hooks · 11,737 LOC
- [[stlovedproduct]] — 10 hooks · 2,821 LOC
- [[stnewsletter]] — 10 hooks · 4,602 LOC
