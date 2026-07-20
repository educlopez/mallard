---
source_url: https://elementflow.io/third/posts/display-condition-feature-6836.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, feature]
section: Features
---
# Display Condition Feature

The display condition feature is mighty, it's one of the key features of Element Flow.

Here are some conditions:

**General**

- Customer groups
- Languages
- Currencies
- The shopping cart is empty or not
- Day & time
- Date Range
- Pages

**Blog**

- If a post has a video
- If a category has any post.

**Product**

- If a product has combinations
- Product stock
- Combination stock
- Behavior when out of stock
- If a product is in the selected categories.
- If a product is in the selected products.
- If a product is in the selected brands.
- If a product has images
- If a product has related products.
- If a product is on sale.
- If a product is new.
- etc.

**Category**

- If a category is empty.
- If a category has any subcategories.
- etc.

I will show you some live examples of using the display condition feature.

### #1 - Display different content by custom group

You aren't logged in yet, so you can see this message, [click here](https://elementflow.io/third/login?back=https%3A%2F%2Felementflow.io%2Fthird%2Fposts%2Fdisplay-condition-feature-6836.html) to log in, and then come back you will see a different message.

[![](https://elementflow.io/img/cms/document/1779.jpg)](https://elementflow.io/img/cms/document/1779.jpg)

### #2 - Display rich content when the shopping cart is empty.

Instead of showing a message saying the shopping cart is empty, displaying some feature products would be a better option when the user enters an empty shopping cart page.

[![](https://elementflow.io/img/cms/document/1778.jpg)](https://elementflow.io/img/cms/document/1778.jpg)

### #3 - Display a gray button when a product is out of stock.

[Check this demo](https://elementflow.io/fourth/en/perfume/901-1086-ethereal-blossom.html), it has the very last item in stock, add it to the cart, an gray out-of-stock button would show out.

### Nested display conditions

You might need to show a message for a certain group in a selected category. How would you do that?

Nest two containers, set one condition to the out container, and set another condition to the inner container.
