---
source_url: https://elementflow.io/third/posts/product-page-builder-6944.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, page-builders]
section: Page builders
---
# Product page builder

The layout of the product page is unlimited, create one suit for your business, oh, not one, as many as you want, you can create a product page for every product.

### 1. Add 3rd party PrestaShop modules to the product page.

Use the **PrestaShop module** widget to add 3rd-party PrestaShop modules to the product page. Check the Widgets > [PrestShop module widget](https://elementflow.io/third/en/posts/prestashop-module-widget-6857.html) for more info.

If a module adds fields to the buy from, then it may need to be integrated, check #0 on the [module integration page](https://elementflow.io/second/en/pages/integration-with-prestashop-modules-6158.html).

If a module **needs to be updated** when the combination is changed, then go to the **Advanced** tab of the PrestaShop module widget to set the **Cache** to **No, refresh after changing product combinations on the product page.**

![](https://elementflow.io/img/cms/document/1776.jpg)

### 2. How to add a thumbnail slider for the main slider.

1. Add **two** Product gallery widgets to the product page.
2. Set the **Type** to **Thumbs gallery - main** for one widget, set the **Type** to **Thumbs gallery - thumbs** for another.

![](https://elementflow.io/img/cms/document/1775.jpg)

### 3. How to add the product description to the tab.

1. The tab is created by the **Tab & Megamenu** widget, so click to set the Tab widget.
2. Change a tab item's title to **Description**
3. Add the Product description to the tab's container.

![](https://elementflow.io/img/cms/document/1773.jpg)

### 4. How to display custom tabs out?

Use the **Tab & Megamenu** widget, enable the **Display product extra tabs** to show all tabs added by PrestaShop modules.

The **Product extra content** widget can also be used to show extra content, the widget can be used inside and outside the tab.

![](https://elementflow.io/img/cms/document/1768.jpg)

![](https://elementflow.io/img/cms/document/1771.jpg)

### 5. How to add custom tabs?

- Go to BO > Improve > Element Flow > Content Builder > Product.
- Add a new record, set the **Show on** to **displayProductExtraContent(Product tab)**. Via the **Apply to** setting, you can set where the tab will be applied to; you have two options:

1. Add it to some selected products.
2. Add it to all the products of a selected category.

![](https://elementflow.io/img/cms/document/1767.jpg)

### 6. How to remember the viewed products

Add the **Viewed products block** module (ps_viewedproduct) to the product page by using the **PrestaShop Modules** widget to remember the viewed products.

Use the **Product** widget to show viewed products.

If the **Viewed products block** module is added before the **Product** widget, then the current product will show in the **Product** widget; otherwise, the current product won't show right away, it will show when another page is opened.

### 7. Product details

PrestaShop puts basic **product info** in a div, the ID of the div is **product-details**. Some modules retrieve product info from the div, so if you have one of those modules, then use the **Custom template** widget to add the product-details div to the product page by loading the built-in **product-details-data-only.tpl** file.

### 8. Gallery images don't show out?

It often happens when the product is imported/copied. Go check the **Image type** setting under the **Layout** section in the Product gallery widget, it's probably empty, then set a value to it. The source of the issue is that your site doesn't have the image type that is used on the site where the gallery is imported/copied from.

### 9. The page template being used isn't the correct one.

The page template is wrong when editing the product description. In this case, you must have multiple product page templates, and some of them are disabled.

To fix the issue, set the priority for the **disabled ones** to 0 (This step is important), and set proper priorities to the other ones.

### 10. Add a loader while changing combination

Add a loader to indicate that the product info is getting updated, and also block clicks.

Check the demos:

[https://elementflow.io/second/109-886-solid-rib-yoga-crop-top.html#/3-size-l/10-color-red](https://elementflow.io/second/109-886-solid-rib-yoga-crop-top.html#/3-size-l/10-color-red)

[https://elementflow.io/fourth/901-1086-ethereal-blossom.html](https://elementflow.io/fourth/901-1086-ethereal-blossom.html)

Use the JS events feature to achieve that.

1. Add a loader; it can be anything. In the above demos, I added a container and put one loading icon inside.
2. Go to the container's Advanced tab, find the JS events section. Set the **Type** to be **Listener**, set the **Listener** to **Show the widget when the product info is getting updated**.
3. The loader needs to be hidden; it only shows out when the product is getting updated. So add a **stsb_display_none** classname to the container.

![](https://elementflow.io/img/cms/document/2331.jpg)

![](https://elementflow.io/img/cms/document/2332.jpg)
