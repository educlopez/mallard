---
source_url: https://elementflow.io/third/posts/tabs-megamenu-widget-7252.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# Tabs & Megamenu widget

This widget is a nested widget, which means all other widgets can be put inside it, even itself. But it has a performance issue when it's used to create a megamenu, because a megamenu usually contains quite a lot of content.

The performance issue has been improved in v2.0, but if too many widgets are added to it, it will still slow the editor quite a lot. The performance will be improved further in the future. Currently, using the shortcodes to avoid having the performance issue is the best solution. Follow these steps:

1. Create a menu using the Tabs & Megamenu widget.
2. Create every submenu as a shortcode.
3. Add shortcodes to the menu.

Screenshots: [1](https://elementflow.io/img/cms/document/1947.jpg), [2](https://elementflow.io/img/cms/document/1948.jpg)
