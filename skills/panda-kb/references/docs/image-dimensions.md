---
source_url: https://panda2.sunnytoo.com/doc/#dimensions
ps_version: [8, 9]
ingested: 2026-04-18
tags: [panda, docs, images, specs]
section: Image dimensions
---
# Image dimensions

The theme does not require image types to be certain sizes. I will use the main demo store as an example to expain how to set image dimensions to make images look sharp. Here are the dimensions of images we used on our main demo store.

![Image dimensions](assets/img/dimensions.jpg)

Image dimesions

1. To know the ratio of your product images. Product images we prepared for the main demo store are 1200px in wide, 1372px in height, so they have a 1200:1372 ratio. It's highly recommended to make all product images have a certain ratio, otherwise PrestaShop will automaticly fill white color around them.
2. The theme allows product images in differnt blocks having different sizes. Let me use the "Featured product slider" module as an example, you can find a recommended width for the module on the back end.
	![Recommended image size](assets/img/recommended_image_size.jpg)
3. Use the recommended width (280) and the ratio (1200:1372) to get the hight. 280\*1372/1200=320.133, round it to get the height 320. If you get 320.5, then the height would be 321.
4. Please check if you already have an image type which is 280x320, if yes, choose it. If no and all product images on your site are in the same size, then you can change the size of an existed image type, generally home\_default. If you have several differnt sizes, then create a new image type for products, name it whatever you want, just be sure it ends up with \_default, if you enable the retina option, then create another image type with the same name, but ends up with \_default\_2x.

Generally you don't need to chanage the width of " cart\_default " and " small\_default ".

PS no way to make images on fullscreen blocks look sharp, cause the width fullscreen block are different on different devices, it could be any number, like 1660, 1920, 2100...
