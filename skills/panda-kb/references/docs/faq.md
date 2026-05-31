---
source_url: https://panda2.sunnytoo.com/doc/faq.php
ps_version: [8, 9]
ingested: 2026-04-18
tags: [panda, docs, faq]
section: FAQ
---
# FAQ

## Frequently asked questions

If you have a question and couldn't find it here, just send us an [email](mailto:support@shamsoft.net).

  

- ###### Where do I put the Tracking code I got from Google?
	1. Turn off the "Use HTMLPurifier Library" setting on the Preferences > General page.
	2. Put the code to the "Tracking code" in the Theme editor moudle.
- ###### Why my site shows nothing, why my site is down? How to debug?
	There are variety reasons, may be due to server settings or the hosting environment, maybe there are bugs. When this happens, the first thing you need to do is enabling the Prestashop debug to see if there are any error messages, you will either be able to troubleshoot further or send the error messages to us to fix the problem.
	Debug Mode can be enabled from the back office CONFIGURE menu under Advanced Parameters > Performance
- ###### Why I can not put a video to the product description?
	Turn off the "Allow iframes on HTML fields" setting on the Preferences > General page.
- ###### How to make Youtube or Vimeo videos be responsive?
	Please refer to [here](http://v4-alpha.getbootstrap.com/utilities/responsive-helpers/#responsive-embeds).
	Videos will be full width disregarding the width and height you set. If that's not what you want, you can use the Bootstrap's grid system to set a proper width to videos.
	```markup
	<div class="row">
	  <div class="col-12 col-md-8 col-lg-6">
	    <div class="embed-responsive embed-responsive-16by9">
	      <iframe class="embed-responsive-item" src="//www.youtube.com/embed/zpOULjyy-n8?rel=0" allowfullscreen>
	      </iframe>
	    </div>
	  </div>
	  <div class="col-12 col-md-4 col-lg-6">
	    Put any thing here.
	  </div>
	</div>
	```
	![Responsive videos](assets/img/responsive_videos.jpg)
- ###### Where can I find a list of all available icons
	/assets/fontello.zip in the main.zip file.
- ###### Does not have quick view feature on the product page
	Currently can not have quick view feature on the product page, due to a problem of Prestashop.
- ###### How to find out what version of the theme you are using?
	Go to BO > Design > Theme & Logo page.
- ###### How to find out what version of Prestashop you are running?
	Go to BO > Advanced Parameters > Information page.
- ###### Why got "This connection is not secure. Logins entered here could be compromised"?
	Solution is enabling SSL. More info here: https://support.mozilla.org/en-US/kb/insecure-password-warning-firefox?as=u&utm\_source=inproduct
- ###### How to set a bottom margin to a layer sldier?
	Please refer to the "Modules" > "Layer sldier" section.
- ###### How to adjust the gap between header and content.
	Please refer to the "Modules" > "Breadcrumbs" section.
- ###### Why I keep losing my page layout settings after clear the Smarty cache.
	Please refer to the "Prestashop" > "Tips" section.
- ###### "BO > module and service" does not work after installing the theme.
	Please check if your hosting meets PrestaShop System requirements. [https://www.prestashop.com/en/system-requirements](https://www.prestashop.com/en/system-requirements) Try increasing memory\_limit to "128M" or "256M", if that does not work, increasing max\_execution\_time to "60". If you don't know how to do that, contact your hosting company for help. If you hosting company is not helpful, then contact us.
- ###### Why I got an error when installing a 3rd party module.
	In most case, the problem is a problem of network, try re-installing it for one or two more times. Another solution is to manually upload the module via FTP of cPanel to avoid having network problems. ![Responsive videos](assets/img/install_module_error.jpg)
- ###### Be careful when you edit files.
	If you find the file, which you are going to editing, having empty lines after every line of code, then don't edit it or upload it back to your site. ![Responsive videos](assets/img/problems_of_file_transfer.jpg) The problem is probably due to your FTP clients or your server settings, please Google for solutions. An alternative way is using cPanel to zip files before downloading or uploading them to avoid having problems during file transfers.
- ###### How to increase max alowed character of product names from 128 to 255?
	1\. Change database, go to ps\_product\_lang table and change length of "name" and "link\_rewrite" fields.  
	2\. Open \\classes\\Product.php, change the size attribute of "name" and "link\_rewrite" fields.  
	3\. Open \\src\\PrestaShopBundle\\Form\\Admin\\Product\\ProductInformation.php, seach for 128, change it to 255.
- ###### How to add media servers?
	Refer to the "How to speed your site up" section on [here](/doc/index.php#prestashop_tips "How to speed your site up").
- ###### Icons disappeared after using CDN or media servers.
	Go to BO > Modules > Theme editor module > “Custom codes” tab > Put this code in the “Custom CSS code” field.
	```markup
	@font-face {
	  font-family: 'fontello';
	  src: url('//your-domain/themes/panda/font-fontello/font/fontello.eot?81926201');
	  src: url('//your-domain/themes/panda/font-fontello/font/fontello.eot?81926201#iefix') format('embedded-opentype'),
	       url('//your-domain/themes/panda/font-fontello/font/fontello.woff2?81926201') format('woff2'),
	       url('//your-domain/themes/panda/font-fontello/font/fontello.woff?81926201') format('woff'),
	       url('//your-domain/themes/panda/font-fontello/font/fontello.ttf?81926201') format('truetype'),
	       url('//your-domain/themes/panda/font-fontello/font/fontello.svg?81926201#fontello') format('svg');
	  font-weight: normal;
	  font-style: normal;
	}
	```
- ###### Got a white page when you upload images in the Revolustion sldier module.
	1\. Check if mbstring is installed in PHP.  
	2\. If the above step did not work, then enable error reporting to see if there are any error messages.
- ###### How to get the latest version of the theme.
	Re-download the theme from the Download page on sunnytoo.com to get the latest version.
- ###### Why my site went down, I got a white page with logo only.
	Go to Advanced Parameters > Performance page, check if these settings "Disable non PrestaShop modules" and "Disable all overrides" are disabled, if they are not, disable them and save.
	![An important step](assets/img/debug_section.jpg)
- ###### Revolution slider' "Margin bottom" setting does not work for fullscreen sliders
	If you set a top margin to the block which is right below the fullscreen revolution slider to set them apart. Every module of the theme has options to set margins.

This page will be updated regularly.
