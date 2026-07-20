---
source_url: https://elementflow.io/third/posts/from-dev-to-production-6757.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, getting-started]
section: Getting started
---
# From dev to production

There are 3 common cases when moving a site from dev/staging to production:

1. Change the domain only.
2. Transfer the site to a new server.
3. Transfer data created by Element Flow only.

Note on internal links: don't copy the URL of an internal page from the address bar into a link field. It's not recommended — the URL won't change automatically when the domain or Friendly URL changes, requiring extra work to fix later. Instead, select the internal page directly, or fill the link field using the **dynamic tag** feature, so the URL of an internal page stays correct automatically.

### #1 - Change the domain only

When you take your site live and need to change the domain (e.g. from `staging.elementflow.com` to `elementflow.com`), go to BO > Configure > Traffic & SEO > SEO & URLs to change the domain. That's it.

### #2 - Transfer the site to a new server

Transfer all files and the database to a new server:

1. (Old server) Change `database_name`, `database_user`, and `database_password` in `/app/config/parameters.php`. Generally no need to change `database_host` — keep it as `localhost` or `127.0.0.1`.
2. (Old server) Zip all files of the staging site. Recommended to **not zip the `/var/` folder** to keep the zip smaller.
3. Upload the `.zip` file to your server via FTP, then unzip via cPanel.
4. (Old server) Export the database to a `.sql` file, generally via phpMyAdmin.
5. Open the `.sql` file in a text editor (EditPlus, Sublime, etc.) and replace all old URLs with the new ones — e.g. replace `staging.element.com` with `element.com`.
6. (New server) Import the modified `.sql` file on the live site using the tool provided by your hosting provider, generally phpMyAdmin.
7. (New server) Log in to the site back office to clear the Smarty cache.

### #3 - Transfer data created by Element Flow only

1. (Old server) Go to BO > Improve > Element Flow > Settings > Import & export > **Export API**, to find the **API URL** and **Token**.
2. (Old server) Click the **Export** button to select data to export.
3. (New server) Go to BO > Improve > Element Flow > Settings > Import & export > **Import from other sites**, paste the **API URL** and **Token** from the old server, and save.
4. Click the **Import** button in the top right corner.
5. Select the old server to import its data.
6. For more info about the import & export feature, see the [Import demo data](https://elementflow.io/third/posts/import-demo-data-6756.html) page.

### Move a site from a subfolder up to the root folder

It's common to build a staging site in a subfolder and move it to the root folder when going live. In this case, **images** uploaded via the image uploader may stop showing — one more step is needed to fix it.

When in a subfolder, image paths saved in the database look like `/a-subfolder/img/cms/an-image.jpg`. When the site moves to the root folder, paths need to change to `/img/cms/an-image.jpg`. Use the **Path replacement tool** on the BO > Improve > Element Flow > Settings page:

- Fill `/a-subfolder/img/cms/` into the **Old path** field, and `/img/cms/` into the **New path** field. Paths must contain `/img/cms/` and must start and end with a `/`.
- If instead copying a live site into a subfolder, do the reverse: fill `/img/cms/` into **Old path**, and `/a-subfolder/img/cms/` into **New path**.
