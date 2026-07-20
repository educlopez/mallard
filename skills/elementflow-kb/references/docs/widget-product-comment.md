---
source_url: https://elementflow.io/third/posts/product-comment-widget-7573.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# Product comment widget

### Why the product comment widget doesn't work?

It shows some info, but the comments aren't listed out.

1. Your site is probably under **maintenance**, and you didn't add your IP to the whitelist. **Adding your IP to the whitelist can fix the issue.**

The issue exists with quite a lot of Ajax requests. Although **Enable store for logged-in employees** is enabled, it's still recommended to add your IP to the Whitelist.

Screenshot: [1](https://elementflow.io/img/cms/document/1994.jpg)

2. Another possible reason is that **two** product comment widgets were added **to the same page**. It's common to display the product comment widget in a tab on the Desktop version and display it in an accordion on the Mobile version. Don't add the product comment widgets twice; the Tabs & Megamenu widget can switch to an accordion layout on the mobile version.

Screenshot: [2](https://elementflow.io/img/cms/document/1998.jpg)
