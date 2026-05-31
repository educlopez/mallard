---
source_url: https://panda2.sunnytoo.com/doc/#theme_update
ps_version: [8, 9]
ingested: 2026-04-18
tags: [panda, docs, install, update]
section: Theme update
---
# Theme update

Important this is for update Panda V2.A to V2.B, not update Panda V1.X to Panda V1.X

The version of the theme is the version of the "Theme editor" module, not the version showing on the "BO>Design>Theme" page.

1. Backup PrestaShop files and the store database. This is very important.
2. If you have changed some theme files, you have to manually merge your changes into the update files. A free software called WinMerge might help.
3. Unzip the update2.x.x.zip file, and then upload folders in zip to the root folder of your site via ftp.
	![Upload update folders](assets/panda/upload_update_folders.jpg)
4. Go to BO > Module > Installed modules page, check through all modules of the theme, upgrade ones with a "Upgrade" button. Like this one:
	![Upgrade modules](assets/panda/upgrade_modules.jpg)
5. Clear the Smarty cache.
6. Congratulations! You have successfully upgraded the theme!
