---
source_url: https://elementflow.io/third/posts/product-miniature-6924.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, page-builders]
section: Page builders
---
# Product miniature

It's also called Loop, one of the most used features in Element Flow. A product miniature or a loop is a block that gets used over and over to display a list of products.

### How to create a button that combines the Add to cart text and the price?

There are two ways:

1. Set a button's text to **Add to cart** + **Product price**
2. Another way is to add a text widget and a product widget widget to a container, set the link of the contianer to **Add to cart**

You can find exmaple of them on the [product miniature demo page](https://elementflow.io/second/en/pages/product-miniature-builder-1387.html). If you like them, copy them to your site.

### How to have a hover image?

Another image shows out when the miniature is on hover. Here are key points of doing it.

1. Stack to two images by setting one of the image's position to absolute to make it on top of another.
2. Set the top image's opacity to 0 to make it invisible.
3. Set the top image's opacity back to 1 on hover.

![](https://elementflow.io/img/cms/document/1761.jpg)

![](https://elementflow.io/img/cms/document/1762.jpg)

### Troubleshoot

Product name and product price are in the same line, if the miniature is small and the name is long, the name may push the price out of the miniature.

The solution is to set the **Size** to **1 1 0**

![](https://elementflow.io/img/cms/document/1763.jpg)

![](https://elementflow.io/img/cms/document/1764.jpg)

### Why does the product miniature look different in the front office?

A miniature template will be updated when its content is changed, but the templates that include the newly updated miniature template won't be updated automatically. At the time, if you check the miniature on the front end, the miniature doesn't look the same as it does in the editor. Clearing the Smarty cache won't work; the first step is to clear the template cache, and then the Smarty cache.

Check the **Clear cache** section on the [FAQ page](https://elementflow.io/third/posts/faqs-2805.html) for more info.
