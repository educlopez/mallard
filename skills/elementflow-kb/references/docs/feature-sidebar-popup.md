---
source_url: https://elementflow.io/third/posts/sidebar-6913.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, feature]
section: Features
---
# Sidebar & Popup

1. Go to BO > Improve > Element Flow > Theme builder > Sidebar.
2. Add a new record, and then open it in the editor.
3. You can add **anything** to the sidebar, add some links to make a side menu, add the product widget to show products in the shopping cart to make it a side cart, etc.
4. Click on the **gear icon** on the black top bar of the editor to adjust the width, change the location, set some display conditions, etc.
5. Save and view the sidebar on the front office, no, you can't see it yet (Note a popup can show out). Add a button somewhere to show the sidebar, use the **Sidebar button** widget to add the button.
6. When a sidebar is opened, clicking on its overlay can close the sidebar. Another way is to add a close button to the side by dragging in the **Action button** widget and setting the action to **Close a sidebar/popup**.

[![](https://elementflow.io/img/cms/document/1755.jpg)](https://elementflow.io/img/cms/document/1755.jpg)

[![](https://elementflow.io/img/cms/document/1756.jpg)](https://elementflow.io/img/cms/document/1756.jpg)

[![](https://elementflow.io/img/cms/document/1754.jpg)](https://elementflow.io/img/cms/document/1754.jpg)

### 1. How to open a side bar when a product is added to the cart?

![](https://elementflow.io/img/cms/document/1753.jpg)

Set the **Type** to **Side cart**, a side cart will automatically show out when a product is added to the cart.

Set a sidebar to locate in the center to make it become a popup.

1. A pop-up can show automatically, and you can set rules like, show one once, once a day, once a session, etc.
2. If the popup has a newsletter form, then you can set it to show for people who haven't subscribed yet.

[![](https://elementflow.io/img/cms/document/1757.jpg)](https://elementflow.io/img/cms/document/1757.jpg)

Besides using the sidebar button widget, another way is to use the dynamic tag feature. If anything has a link, then it can be used to open/close a sidebar.

[![](https://elementflow.io/img/cms/yogaclothes/ins1.jpg)](#elementor-action%3Aaction%3Dpopup_open%26settings%3DeyJpZCI6IkQ3VjhYUnU4In0%3D)

[![](https://elementflow.io/img/cms/document/1758.jpg)](https://elementflow.io/img/cms/document/1758.jpg)

Check #3 above, Element Flow already has two ways to open a side; this is the 3rd way by creating the link manually.

1. The first step is to find out the code of a sidebar, see the first pic below, that's the code of a sidebar.
2. ```
   encodeURIComponent(btoa('{"id":"24mMv3FT"}'))
   ```
   Run the JS code anywhere you can run JS code, see the second, it shows how to run the code in Chrome's Developer Tools. Don't forget to alter **24mMv3FT** with the code of your sidebar.
3. Copy the result. Add the result to the end of this text,
   ```
   #elementor-action%3Aaction%3Dpopup_open%26settings%3D
   ```
   which makes the URL open a side panel. Like this
   ```
   href="#elementor-action%3Aaction%3Dpopup_open%26settings%3DeyJpZCI6IjI0bU12M0ZUIn0%3D"
   ```

[![](https://elementflow.io/img/cms/document/2055.jpg)](https://elementflow.io/img/cms/document/2055.jpg)

[![](https://elementflow.io/img/cms/document/2056.jpg)](https://elementflow.io/img/cms/document/2056.jpg)

1. Click on the **gear icon** on the black top bar of the editor.
2. Set the **Position** to **Bottom**
3. Set a height, like **80vh**

![](https://elementflow.io/img/cms/document/2279.jpg)

![](https://elementflow.io/img/cms/document/2278.jpg)
