---
source_url: https://elementflow.io/third/posts/header-builder-8065.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, page-builders]
section: Page builders
---
# Header builder

### Why do some elements show a bit later?

The header is cached by default, but some widgets in it can't be cached, like the shopping cart and the user info widget. So those widgets get loaded separately to achieve a fast loading speed and accurate information, but there is a noticeable delay in appearance.

The best practice is to add some placeholders using the skeleton loader feature to prevent people from seeing the delay. [Check this demo](https://elementflow.io/second/en/pages/yoga-clothes-homepage-6202.html).

Another way is to not cache the header by disabling its Smarty cache setting, see the first picture below.

![](https://elementflow.io/img/cms/document/2130.jpg)
