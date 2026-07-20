---
source_url: https://elementflow.io/third/posts/sign-in-widget-7185.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# Sign in widget

You can set 3 different actions for when the sign-in button is clicked:

1. Open a sidebar — a sign-in form can be placed inside the sidebar.
2. Link to a page — **leave the link field empty** to use the log-in page, and a **back** parameter for the current page will be added automatically.
3. Show a drop-down menu for more options.

Option #2 affects where the user is **redirected** after logging in. By default, the user is redirected to the page where they clicked the sign-in button. To always redirect to a specific page (e.g. order history), set the button's link to the Sign-in page itself (without the auto-added `back` parameter), then go to the **Login form** widget and set its **Fallback redirection** to the desired page.

### Customers

The following info can be shown for logged-in customers, with the same 3 possible click actions:

1. Open a sidebar.
2. Link to a page.
3. Show a drop-down menu for more options.
