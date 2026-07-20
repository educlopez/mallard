---
source_url: https://elementflow.io/third/posts/shortcode-feature-6837.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, feature]
section: Features
---
# Shortcode feature

The shortcode feature is used to display the same content in several different places, for example, the left navigation on this page is added by using a shortcode, why? Because the same navigation is added to a sidebar as the mobile navigation.

These two widgets can be used to add shortcodes to a page:

1. Html widget
2. Custom template widget.

Always wrap your shortcodes in {literal}{/literal} in the HTML widget, like this {literal}{SSBC id=3108}{/literal}
To add a shortcode to a custom template file, no need to wrap it in literal tags; just put it as what it is. {SSBC id=3108}

[![](https://elementflow.io/img/cms/document/1737.jpg)](https://elementflow.io/img/cms/document/1737.jpg)

### 1. Every record can be displayed using the shortcode feature.

When you create a shortcode record, you will get a shortcode like this {SSBC id=3108}
3108 is the ID of the record. For example, if you have a custom page, you like to display the content of the custom page in another place, you can also use the shortcode feature, like this {SSBC id=the-custom-page-id}
You can always find the ID of a record in the address bar when the record is opened in the editor.

### 2. For your information

While you edit a shortcode record in the editor, some of your changes may not take effect on the canvas, which is probably caused by the same shortcode being added to the same page several times, when it happens, save your changes, and view them on the front office.
