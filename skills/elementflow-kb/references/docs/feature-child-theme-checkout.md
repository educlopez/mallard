---
source_url: https://elementflow.io/third/posts/style-checkout-page-and-my-account-page-6987.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, feature, checkout, child-theme]
section: Features
---
# Style checkout page and my account page

Install the Element Flow child theme to style the checkout page and the my account page. The child theme is for the **Classic** theme, aiming to style the My Account page and the Checkout page, considering that massive **compatibility issues** may occur if those pages were **built with** Element Flow.

Here is a video to show how to style the checkout page and my account page, start at 3:52.

### 1. Download the child theme

**PrestaShop 9** allows each theme to have a set of image types. Installing a new theme won't erase existing image types anymore. Download the child theme for PrestaShop 9 below **(Right click on the links, select "Save link as")**.

[Child theme for the Classic theme (elementflow.zip)](http://elementflow.io/stupload/elementflow.zip)

[Child theme for the Hummingbird theme (elementflow_hb.zip)](http://elementflow.io/stupload/elementflow_hb.zip)

**For PrestaShop 1.7 and PrestaShop8**, go to BO > Element Flow > Settings > Child theme section to download the child theme. Why does the child theme have to be downloaded here? Because installing a theme erases all existing image types, which isn't desirable in this case. Downloading the child theme here will add the existing image types to the child theme, so your image types will be added right back after they are erased.

To show the checkout page in the canvas, please take these steps:

1. Add a product to the cart on the front end to make sure you can access the checkout page.
2. Go to **BO > Element Flow > Theme builder > Footer**, open a footer in the editor. If you don't have a footer record, then find a record that is on the checkout page, like a header or a sidebar.
3. See the image below to click the link there.
4. If you like a certain step to be opened in the canvas, for example the **Shipping** step, then go to the **Shipping** step on the front end, and then come back to reload the editor.

Open up the editor, go to Site settings > My account page > Navigation items > Create the navigation menu here.

Don't add the **dashboard** page to the navigation, because the dashboard page is empty. If the user gets redirected to the empty dashboard page after logining, then you can change the redirection, take a look at the **Where will users be redirected to after logging in?** section on this page [Login form](https://elementflow.io/third/en/posts/login-form-6961.html)

1. The code for the modal is located in the footer in most themes. So load the modal into the footer created by Element Flow by using the custom template file, the file is already built-in, its name is **terms-and-conditions-modal.tpl**.
2. If you are using the native cms page, then skip this step. If you have a **cms page template** in Element Flow, then edit this template to add a class name **page-cms** to the **Cms content** widget.

### 4. How to make the checkout page be in one column on mobile

The cart summary section is on the right side. Moving it to the top or bottom can make the checkout page appear in one column.
