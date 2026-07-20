---
source_url: https://elementflow.io/third/posts/how-to-translate-6883.html
ps_version: [8, 9]
ingested: 2026-07-20
tags: [elementflow, docs, howto, translation]
section: Features
---
# How to translate

### 1. Translate the front office.

**Most texts** on the front end **can be translated in the Live Editor**. There is a language switcher along every text field for you to do translation. Only a few texts can't be translated in the editor, so download translations for them on this [translation tool](https://elementflow.io/module/sttranslation/list?tpn=Sitebuilder).

### 2. Translate the back office.

### How to use the Translation tool?

[Click here to use this tool to translate the Element Flow module.](https://elementflow.io/module/sttranslation/list?tpn=Sitebuilder)

There, you can translate the front office and the back office separately. **The font office translation** has **fewer than 100** expressions and has already been almost fully translated into **20+** languages.

How to use the downloaded translation files? Let us use the es-ES.zip as an example.

**For PrestaShop 9 and 8.1.3+.**

1. Unzip the es-ES.zip, and you will get a file named ModulesSitebuilderShop.es-ES.xlf
2. **Copy** the .xlf file into the **/translations/es-ES/** folder.
3. **Rename** the .xlf to **ModulesSitebuilderShop.xlf**, and copy it to the **/translations/default/** folder. If the file already exists, which means you did this step before, you don't have to do this step again and again; just make sure there is one ModulesSitebuilderShop.xlf existing in the /translations/default/ folder.
4. Remove these two folders to clear the translation cache.
   **\var\cache\dev\translations**
   **\var\cache\prod\translations**

**For PrestaShop v1.7 and PrestaShop v8 - 8.1.2.**

It's almost the same as the above tutorial; the only difference is the folder where the translation files got copied into.

1. Unzip the es-ES.zip, and you will get a file named ModulesSitebuilderShop.es-ES.xlf
2. **Copy** the .xlf file into the **/app/Resources/translations/es-ES/** folder.
3. **Rename** the .xlf to **ModulesSitebuilderShop.xlf**, and copy it to the /app/Resources/translations/default/ folder.
4. Remove these two folders to clear the translation cache.
   **\var\cache\dev\translations**
   **\var\cache\prod\translations**

### What about using the PrsetasShop's native translation system?

Follow the tutorial above to add .xlf translation files to your site first. If you need to change translations, you can choose either to use the tool or use the PrestaShop translation system.

To use the PrestaShop translation system, select **Front Office translation** from the **Type of translation** dropdown menu.
