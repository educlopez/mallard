---
source_url: https://panda2.sunnytoo.com/doc/#customization
ps_version: [8, 9]
ingested: 2026-04-18
tags: [panda, docs, customization, css]
section: Theme customization
---
# Theme customization

#### Custom css code

If you want to use custom css code, there are two ways.

1. Put your custom css code to the "Custom css code" field in the Theme editor module, recommended.
2. If you prefer editing files, then edit the \\themes\\panda\\assets\\css\\custom.css file.

The same ways also work for custom js code.

---

#### Child theme

PrestaShop's child theme feature is not ready yet, but it will be available in future.

---

#### How to add a custom font to your site?

1. Make sure the font is a web font. A web font should have several files, like.eot,.woff,.ttf and.svg More info here: http://www.w3schools.com/css/css3\_fonts.asp
2. Upload font files to the /modules/stthemeeditor/views/css/ folder.
3. Put the font name into the "Custom fonts" field in the Theme editor moudle.
4. You need to put a piece of css code to the "Custom css code" field in the Theme editor moudle. Here is an example
	```markup
	@font-face {
	  font-family: 'My font';
	  src: url('my_font.eot');
	  src: url('my_font.eot#iefix') format('embedded-opentype'),
	       url('my_font.woff') format('woff'),
	       url('my_font.ttf') format('truetype'),
	       url('my_font.svg') format('svg');
	  font-weight: normal;
	  font-style: normal;
	}
	```

---

#### Short code feature

Short code feature currently avaiable in these 3 moduels: Advanced Banner, Swiper and Ins. The feature allows you to display blocks from those modules kinda like anywhere you want.

Let me use the Swiper module as an example to show you how to use the feature.

- Copy short codes.
	![Copy short codes](assets/img/copy_short_codes.jpg)
- Paste them into Sourth Code field.
	![Paste short codes](assets/img/paste_short_code.jpg)
	Don't put short codes in any other tags but <div></div>, like these ones are wrong, they might cause your site misplaced.
	```markup
	<h3><div>[stswiper id="23"]</div></h3>
	<h3>[stswiper id="23"]</h3>
	<p><div>[stswiper id="23"]</div></p>
	<p>[stswiper id="23"]</p>
	<span><div>[stswiper id="23"]</div></span>
	<a href="#"><div>[stswiper id="23"]</div></a>
	<div>It's okay to use nested div tags<p>[stswiper id="23"]<p></div>
	```
	These ones are correct.
	```markup
	<div>[stswiper id="23"]</div>
	<div>It's okay to use nested div tags<div>[stswiper id="23"]<div></div>
	```
