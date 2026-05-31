---
source_url: https://panda2.sunnytoo.com/doc/#prestashop
ps_version: [8, 9]
ingested: 2026-04-18
tags: [panda, docs, prestashop]
section: PrestaShop 1.7 & 8
---
# PrestaShop 1.7 & 8

#### PrestaShop update

You can use the "1-Click Upgrade" module to upgrade Prestashop. After you successfully upgraded your site, please go to BO > Advanced Parameters > Performance page, check if these settings "Disable non PrestaShop modules" and "Disable all overrides" are disabled, if they are not, disable them and save.

---

#### CSV import

Sample CSV files provided by PrestaShop are stored in the /docs/csv\_import folder of your installation of PrestaShop, but 3 of them are not correct missing several columns, We've corrected them and put them in the /csv\_import/ folder in the main.zip file.

---

#### Tips of using PrestaShop

###### 1\. How to configure modules in new tabs?

Currently if you are unable to configure modules in new tabs, it is not convenient if you want to configure several modules at the same time. Here is a trick, you need to edit the \\src\\PrestaShopBundle\\Resources\\views\\Admin\\Module\\Includes\\action\_menu.html.twig file for a little bit. Repalce all `<form` to `<form target="_blank" `

---

###### 2\. One thing you need to do after choosing layouts for pages.

Back up the \\app\\cache\\prod\\themes\\panda\\shop1.json file or the \\app\\cache\\dev\\themes\\panda\\shop1.json file after you choosing "Full width" or "Two columns, small left column" or "Two columns, small right column" for your pages, because of you may loose your settings afeter clearing the Smarty cache.

---

###### 3\. How to speed your site up?

- Go to BO > Advanced Parameters > Performance > Set "Template compilation" to "Never recompile template files" and Enable cache.
- On the same page, disable "Debug mode", enable all settings in the "CCC" section.
- Adding at least one media server to your site would dramatically speed up your site.
	![How to add media servers](assets/img/how_to_add_media_servers.jpg)
	You may ask what's that "w-static.sunnytoo.com", how can I get my own one?
	"w-static.sunnytoo.com" is just a sub-domain, once you bought a domain, you can create as many sub-domain as you want. Generally a sub-domain would take some time to work after creating, to check if it works, you can upload an image(named as 123.jpg) to the root folder of your site, try if you can reach the file via http://the-sub-domain-you-just-created/123.jpg, if yes, then it's a correct one.
	Don't add media servers to your site if it's under maintenance.
- Disable (not uninstall) modules you do not use. Here are a list of Presatshop's native modules we disabled.
	1. Shopping cart
	2. Newsletter subscription
	3. Search bar
	4. Currency block
	5. Language selector block
	6. Link List
	7. Contact information
	8. Customer reassurance
	9. Main menu
	10. Featured products
	11. Social media share buttons
	12. Custom text blocks
	13. Social media follow links
	14. Image slider
	15. My Account block
	16. Banner
	17. Customer "Sign in" link
	18. Customer accounts
	19. Newsletter
	20. Browsers and operating systems
	21. Best suppliers
	22. Pages not found - This module is useful when your site is not stable.
	23. Registered customer information
	24. Welcome
	25. Best vouchers
	26. Best customers
	27. Best categories
- The theme comes with 50+ modules, you do not need all of them, so disable those ones you do not need, here are a list of modules you can not disable.
	1. Theme editor
	2. Overriding default behaviors
	3. If you use blog, then "Blog module" and "Blog editor" are also can not be disabled. If you do not use blog, you can disable all blog modules including those two.

---

###### 4\. Maintenance.

If you are going to put your site in maintenance mode, then do not forget to add your IP address to the "Maintenance IP" field, otherwise some features on back office can not work fine.
