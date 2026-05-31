---
source_url: https://panda2.sunnytoo.com/doc/#theme_install
ps_version: [8, 9]
ingested: 2026-04-18
tags: [panda, docs, install]
section: Theme install
---
# Theme install

#### Prepare

- The theme is for PrestaShop 1.7.1+ & Prestashop 8, Go to BO > Advanced Parameters > Information page to check if the version of your PrestaShop is larger than 1.7.1.0.
	![Find out PrestaShop version](assets/img/findoutprestashopversion.jpg)
- Backup PrestaShop files and the store database. This is very important.
- If you are going to install the theme on a live site, then you need ot enable maintenance mode to disable your site during the installation process.

---

#### Install

If this is your first time to isntall the theme, then install the "/Prestashop1.7/.zip" file, no need to applly the update2.x.x.zip file.

If you have installed the theme and the version of your "Theme editor" module is lower than the "/Prestashop1.7/update2.x.x.zip" file, then follow [this guide](#theme_update) to upgrade your theme to the latest version.

Here is a video guide: [https://youtu.be/\_QXtPY50OPw](https://youtu.be/_QXtPY50OPw "Panda 2 installation")

- Login to your back-office and navigate to the Design > Theme & Logo.
- Click on the Add new theme button on the top toolbar.
	![Add new theme](assets/img/addnewtheme.jpg)
- In the "Import from your computer" section, click on the "Add file" button to select panda.zip and upload. Don't upload the main.zip file you downloaded from www.sunnytoo.com, unzip the main.zip file first, and then upload panda.zip.
	![Improt from computer](assets/panda/importfromcomputer.jpg)
- If you have problems uploading the panda.zip file, there are two solutions: (1) Ask your hosting company for help to incrase the 'upload\_max\_filesize' to at least 10 MB. (2) Upload the panda.zip file to the /themes/ folder using a FTP client or cPanel, and then install from the "Import from FTP" section.
	![Improt from FTP](assets/panda/importfromftp.jpg)
- Hover over the Panda theme on the themes list, and then click on the "Use this theme" button.
	![Use this theme](assets/panda/usethisthem.jpg)
- After a while you will get a page like this, which means the theme has been successfully installed.
	![Success](assets/panda/install_success.jpg)
- This step is important, go to Advanced Parameters > Performance page, check if these settings "Disable non PrestaShop modules" and "Disable all overrides" are disabled, if they are not, disable them and save.
	![An important step](assets/img/debug_section.jpg)
- This step is important, go to BO > Modules > Installed modules > check if the "Overriding default behaviors" module is installed and actived, if it's not try reinstalling/activating it. If you can not get it installed/activated, please take these steps:
	1\. Check if the \\override\\controllers\\front\\CategoryController.php file exists. If it exists, make a backup of it, and then remove it.
	2\. Check if the \\modules\\stoverride\\override\\controllers\\front\\CategoryController.php file exists. Remove it if it exists.
	3\. Try activating the Override default behavior module again.
	If it still can't be actived, take a look on this tutorial [How to manually installed the override default behaviors module](https://www.sunnytoo.com/12920/manually-installed-override-default-behaviors-module) or contact us.
- Clear the Smarty cache.
	![Clear the smarty cache](assets/img/clear_cache.jpg)
- Congratulations! You have successfully installed the theme!
