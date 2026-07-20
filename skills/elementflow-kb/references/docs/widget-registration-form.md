---
source_url: https://elementflow.io/third/posts/registration-widget-6727.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, widget]
section: Widgets
---
# Registration widget

Click on the "View all the required fields" link to see all required fields, and make sure all required fields are added to the registration form — otherwise the registration form can't be saved.

The registration form can be placed on any page and will work, except when using the "Official GDPR compliance" module (or other modules that add required fields to the form with restrictions to work on the registration page only).

The "Official GDPR compliance" module works on the **registration page only**, so it won't appear in the general module list. If you need it, add it to the registration form, since it adds a required checkbox. To make the Official GDPR compliance module work on every page (not just the dedicated registration page), you need to customize the `/modules/psgdpr/psgdpr.php` file.
