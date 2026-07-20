---
source_url: https://elementflow.io/third/posts/integration-with-prestashop-modules-6158.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, feature, module-integration]
section: Features
---
# Integration with PrestaShop modules

### Element Flow is compatible with almost all PrestaShop modules, except other page builders.

- What Element Flow does mostly is to style the front office, it doesn't change the way how PrestaShop works, that's the reason it has great compatibility with other modules.
- **Most (let's say 85%) PrestaShop modules can work with Element Flow without any integration.** Element Flow has 3 widgets to load PrestaShop modules in different ways: the "PrestaShop modules" widget, "Hooks" widget and "Custom templates" widget. Here is a video of [how to add PrestaShop modules in Element Flow](https://www.youtube.com/watch?v=vBmBQdgIpak).
- For the rest 15%, integration will be built for popular modules. Some integration requires editing files, listed below.
- If the module you are going to use isn't in the list below, then it's most likely compatible; if it doesn't work fine, contact them — they provide a free integration service, no guarantee.

## Integration

### 0. Modules that add data to the buy form.

For modules that add data to the buy form on the product page, add a **form="add-to-cart-refresh"** attribute to every form field to make sure they get submitted when the Add to cart button is clicked. Form fields can be placed anywhere on the product page, as long as they have the **form** attribute they submit correctly. Examples:

```
<input type="text" name="message" form="add-to-cart-or-refresh">
```

```
<select name="message" form="add-to-cart-or-refresh">
  <option value="1">1</option>
  <option value="2">2</option>
</select>
```

```
<textarea name="message" form="add-to-cart-or-refresh"></textarea>
```

### 1. SEO Audit

Add this code to the hookDisplayOverrideTemplate function in the modules/ets_seo/ets_seo.php file.

```
if ((int)Tools::getValue('eb_id') && Tools::getValue('stsitebuilder')) {
    return;
}
```

### 2. Gift card

1. Add this code to the desired place for the gift card module to show out on the product page using the "HTML" widget.

```
<div id="gift_product"></div>
```

2. Add the code below to every field in the /modules/giftcard/views/templates/hook/v1_7/gift_radios.tpl file.

```
form="add-to-cart-or-refresh"
```

### 3. Additional Product Attributes / Custom Product Fields

Do two changes to the /modules/an_productfields/views/js/front.js file.

1. Add this code to the first line in the **chacgedPositionAnFields** function. If you have the Quick View feature, also add the code to the **changedPositionAnFieldsModal** function. The module uses these two functions to move fields to the desired position. With Element Flow, custom fields can be placed directly in the desired position.

```
return false;
```

```
$('.elementor-element .js-an-pf-fields-wrap').not('.js-an-pf-position').remove();
```

### 4. Amazzing filter v3.*

1. Add filters to the listing page using the "PrestaShop modules" widget. Module: amazzingfilter, Hook: displayAmazzingFilter

2. Add this code to the custom js code field in the Amazzing filter module.

```
if(typeof(customThemeActions)){
    customThemeActions.updateContentAfter = function(jsonData){
        if(typeof(jsonData['stsb'])=='undefined')
            return;
        elementorFrontend.updateStsb(jsonData['stsb']);
    }
}
```

3. Replace the code marked in the screenshot in the \modules\amazzingfilter\views\js\front.js file:

```
af.$dynamicContainer.animate({'opacity': 1}, 350);
```

```
exit(json_encode($this->module->prepareAjaxResponse($params)));
```

With this:

```
$res = $this->module->prepareAjaxResponse($params);
if (Tools::getValue('euid') && Module::isInstalled('stsitebuilder') && Module::isEnabled('stsitebuilder')) {
    $eb = Module::getInstanceByName('stsitebuilder');
    if($eb_res = $eb->displayAjaxRefresh(true, true))
        $res = array_merge($res, $eb_res);
}
exit(json_encode($res));
```

### 5. Easy filter

Two ways to add filters to the category page:

1. Use the **Custom template** widget to load the easyfilter_filters.tpl and easyfilter_active_filters.tpl files.
2. Use the **Product filters** widget, but color settings won't work; they are for the PrestaShop native faceted filter module. Use color settings in the easy filter module instead.

It's popular to show a **Filter button** on mobile, filters show out from the side when the filter button is clicked. How to do that:

1. Open up the category page template in the editor.
2. Drag and drop the **Custom template** widget to where you'd like the Filter button to be.
3. Select **easyfilter_filter_button.tpl** from the **Select a template tpl** drop-down menu.

It doesn't make sense for the filter button and the filters both to show at once, use the **Visibility** feature in the **Responsive** tab under the **Advanced** tab to hide one of them.

### 6. WebP module

All images you upload are in the /img/cms/ folder, so add the **/img/cms/** folder to the WebP module to generate .webp images.

Once .webp images are generated, you may find an image appearing twice in the file manager. There are menu02.jpg and menu02 — **select menu02, never select menu02.jpg**, because the file manager hides file extensions; the full name of menu02.jpg is menu02.jpg.webp. If you select menu02.jpg.webp for an image widget, the image won't show out on Safari.

### 7. PrestaShop Checkout

It can work fine on the **checkout page** without any integration.

This is about showing the **express checkout button** on the **product** page. The PrestaShop checkout module uses these selectors to find the main Add to cart button:

1. .product-container .product-add-to-cart button.add-to-cart
2. .product-add-to-cart .product-quantity

They should make the selectors configurable. Anyway, the required CSS classes can easily be added to the product page to make the PrestaShop Checkout module find the main Add to cart button.

You can add the CSS classes to existing elements. Prefer wrapping the **Product buttons** widget in two containers to control all button-related changes.

1. Add **product-container** to the outer container.
2. Add **product-add-to-cart** to the inner container.
3. Add **product-quantity** to the product buttons widget.
4. Add **add-to-cart** to the Add to cart button.
5. Add the **product-details-simplified.tpl** file anywhere on the product page via the **Custom template** widget.

### 9. All-in-one Rewards: loyalty referral affiliation review

The loyalty points info on the product page isn't updated after switching the product combination. The All-in-one Rewards module requires the product-details tab, so the solution is to use the **Custom template** widget to add the product-details div to the product page by loading the built-in **product-details-data-only.tpl** file.

For more info, check the **Product details** section on the Product page builder page.

### 10. Minimum and maximum purchase product quantity

The module adds a script to the product page using the displayProductAdditionalInfo hook, so the first step is to use the PrestaShop modules widget to add the script to the product page.

- Module: Custom module
- The name: **minpurchase**
- Hook: **displayProductAdditionalInfo**
- Default parameters: $product

Set the **Cache** of this widget to "No, refresh after changing product combinations on the product page."

Then edit the /modules/minpurchase/views/js/front.js file, replace this code:

```
$('#quantity_wanted').on('change', function(event) {
```

With:

```
$(document).on('change', 'body#product [data-elementor-type="template"] .elementor-widget-st-product-action-button .stsb_pro_quantity', function(event) {
```

And this code:

```
if (Number($('.product-actions .product-add-to-cart input[name='+fieldName+']').val()) < Number(minimum_quantity)) {
    $('.product-actions .product-add-to-cart input[name='+fieldName+']').val(minimum_quantity);
}
```

With:

```
if (Number($('.elementor-widget-st-product-action-button .stsb_pro_quantity').val()) < Number(minimum_quantity) || Number($('#add-to-cart-or-refresh #quantity_wanted').val()) < Number(minimum_quantity)) {
    $('.elementor-widget-st-product-action-button .stsb_pro_quantity, #add-to-cart-or-refresh #quantity_wanted').val(minimum_quantity);
}
```

### 11. Comparaison avancée des produits - cdproductcomparisonplus

Add it to the product page using the PrestaShop module widget.
Set the "Comparaison avancée des produits" to body#product #wrapper to make the "Add to compare" button work on the front page.

### 12. Notify Me Pro - nxtalnotify

Add it to the product page using the PrestaShop module widget.
Add a pb-center-column classname to the widget to make the button look nice on the front office.
