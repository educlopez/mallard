---
source_url: https://panda2.sunnytoo.com/doc/#translation
ps_version: [8, 9]
ingested: 2026-04-18
tags: [panda, docs, translation, i18n]
section: Theme translation
---
# Theme translation

You can translate This theme using the PrestaShop's new translation system, check it out [here](http://build.prestashop.com/news/new-translation-system-prestashop-17/ "New translation system"). Although PrestaShop team said the new translation system currently did not work for 3rd modules, we figured out a way to make the theme work with the new translation system, cause it's really great.

---

#### How to add languages?

Use "Spanish" as an example.

1. Go to BO > "International" > "Translations" page > "Modify translations" section > Check if "Spanish" is in the "Select your language" drop down list, if it's in there, which means it already added, you can skip the next step. If it's not in there, do the next step to add it.
	![Check if a language exists](assets/img/check_if_a_language_exists.jpg)
2. On the same page you can find a "Add / Update a language" section which can be used to add new languages, but it's not recommended, the best way is adding new languages from the "International" > "Localizatoin" page.
	![Add a new language](assets/img/add_a_new_language.jpg)

---

#### How to translate the theme?

There are two ways to translate the theme. (1) Translate form the back office. (2) Translate on the ST translation tool.

##### Translate from the back office

Also use "Spanish" as an example.

1. Make sure you have Spanish language added on your site.
2. To translate the theme, you need to generate translation files first. Use the "Translation files generator" module to generate translation files. If you can not find the module on the "Installed module" list, then go to BO > Modules > You would find the module there, just install it. Go the the module's configuration page, you will find two options, it is recommended to generate translation files for the front office only, because generally we do not need to translate the back office, another reason is that the back office has 2000+ phrases, way better more than the front office's 300+ phrases, those 2000+ phrases would slow the translation page down. Don't worry about losing your translations when regenrating translatoin files, your translations are stored in database, they are safe.
3. Translation fiels are stored in the /themes/panda/translations/ folder. If you have problem of generating files, please check if the folder has write permission.
4. After generating translation files, go to BO > "International" > "Translations" page > "Modify translations" section > Select "Theme translations" + "Panda" + "Spanish", we use Spansih as an example, so we select it here.
5. To translate the front office, locate to "Shop" > "Theme" > "Panda". Here you can translate all phrases displaying on the font office which includes phrases from modules and phrases from files in the /themes/panda/ folder.
	![Translate front office](assets/panda/translate_front_office.jpg)
6. To translate the back office, locate to "Admin" > "Theme" > "Panda" and "Modules" > "St". Why translate the back office from two different places, because of phrases on the "Admin" > "Theme" > "Panda" are common phrases, they are used by more than one module, you just need to translate them for once here.
	![Translate front office](assets/img/translate_st.jpg)
	![Translate front office](assets/panda/translate_admin.jpg)
7. Don't modify translation files manually, your modifications will be overridden when regenerating files.
8. Your translations saved in database, so don't worry about losing them when regenerating files.
9. Clear the Smarty cache to see the result.

---

##### Translate on the ST translation tool(Recommended)

ST translation tool is used to translate the front office and back office of Panda 2, the advantage of using it is that you can share your translations with others and use others translations.

Also use "Spanish" as an example.

1. Make sure you have Spanish language added on your site.
2. [http://panda2.sunnytoo.com/en/module/sttranslation/list?tpn=Panda2](http://panda2.sunnytoo.com/en/module/sttranslation/list?tpn=Panda2 "ST translation tool")
3. After translating or the translation you are looking for has already translated, then download the translation file.
4. Upload.xlf file(s) in the downloaded.zip file to your site, we are using Spanish as an example, so check if there is a folder named "es-ES" in /themes/panda/translations/ folder, create one if there is not. Then upload.xlf file(s) into the "es-ES" folder. You can find a list of ISO language code here [http://www.lingoes.net/en/translator/langcode.htm](http://www.lingoes.net/en/translator/langcode.htm "ISO language code")
5. Clear the Smarty cache.
6. That's all. If you want to change some translations, you can still use the back office to make changes. You won't lose your changes once you update the file.
