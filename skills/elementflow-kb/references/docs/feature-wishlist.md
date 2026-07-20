---
source_url: https://elementflow.io/third/posts/wishlist-feature-6811.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, feature]
section: Features
---
# Wishlist feature

It's PrestaShop's native wishlist block module (blockwishlist), a great module, I have several things to tell you about it.

The module is transplanted to the **displayFooter** hook by default, removing the module from the displayFooter stops the module from work.
If the **Add to Wishlist** button doesn't work on your site, then go to BO > Modules > Positions > displayFooter > Check this the wishlist module is here.

Clicking on an "Add to wishlist" button opens a popup window where you can choose which wishlist the product will be added to, click on a wishlist to add the product to the clicked wishlist, and a message will soon show to let you know the result.

Got several requests about showing products that have been added to the wishlist in a sidebar. One can have **several wishlists**, then which wishlist should be displayed in the sidebar, the last clicked one? Does it make sense to you? Let me know. Some people suggested showing recently added products, but that doesn't make much sense to me.
