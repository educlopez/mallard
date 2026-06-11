---
name: ps-image-regen
description: >
  Arregla zonas blancas (padding) en imágenes de productos PrestaShop y regenera
  thumbnails de forma segura sin bloquear Apache. Úsala cuando el usuario diga:
  "imágenes con blanco", "white zones en imágenes PS", "thumbnails con relleno
  blanco", "regenerar miniaturas PS", "regenerar thumbnails prestashop", "el admin
  de PS se queda colgado regenerando", "imágenes de productos con bordes blancos",
  "imágenes no llenan el contenedor", o cuando vea padding blanco en el listing
  de productos de una tienda PrestaShop. También aplica cuando hay que sincronizar
  comportamiento de generación de imágenes entre local Lando y servidor de staging/prod.
version: "1.0.0"
metadata:
  author: Eduardo Calvo
---

# ps-image-regen — Fix zonas blancas + regeneración segura de thumbnails PrestaShop

## El problema

Por defecto (`PS_IMAGE_GENERATION_METHOD = 0`), PrestaShop genera thumbnails cuadrados
(ej. 320×320) haciendo "fit" proporcional y rellenando el resto con blanco. Una imagen
landscape 1920×1080 en un canvas 320×320 → bandas blancas arriba y abajo.

El método está en `classes/ImageManager.php` línea ~280:
```php
$psImageGenerationMethod = Configuration::get('PS_IMAGE_GENERATION_METHOD');
// 0 = fit con padding blanco (default)
// 1 = escala por anchura, output proporcional, sin padding
// 2 = escala por altura, output proporcional, sin padding
```

## Paso 1 — Cambiar método de generación

### Opción A: Admin PS
Admin → Disseny / Diseño → Configuració d'imatges → **"Generar imatges basades en un costat de la imatge d'origen"** → seleccionar **"Amplada"** (método 1) → Desar.

### Opción B: DB directamente (local o servidor)
```sql
INSERT INTO {prefix}configuration (name, value, date_add, date_upd)
VALUES ('PS_IMAGE_GENERATION_METHOD', 1, NOW(), NOW())
ON DUPLICATE KEY UPDATE value=1, date_upd=NOW();
```

Prefijo de la instalación: ver `app/config/parameters.php` → `database_prefix`.

En Lando local (prefijo `amics_`):
```bash
lando mysql -uroot prestashop -e "INSERT INTO amics_configuration (name, value, date_add, date_upd) VALUES ('PS_IMAGE_GENERATION_METHOD', 1, NOW(), NOW()) ON DUPLICATE KEY UPDATE value=1, date_upd=NOW();"
```

## Paso 2 — Regenerar thumbnails sin colgarse

El regenerador nativo del admin PS usa AJAX con timeout — se cuelga con muchas imágenes
(>500). Usar este script PHP en su lugar:

```php
<?php
// regen_thumbs.php — colocar en la raíz de PS, ejecutar con lando php regen_thumbs.php
require_once __DIR__ . '/config/config.inc.php';
require_once __DIR__ . '/init.php';

$imageTypes = ImageType::getImagesTypes('products');
$images = Db::getInstance()->executeS('SELECT id_image FROM ' . _DB_PREFIX_ . 'image');

$ok = 0; $skip = 0; $errors = 0;

foreach ($images as $row) {
    $id = $row['id_image'];
    $path = __DIR__ . '/img/p/' . implode('/', str_split((string)$id)) . '/' . $id . '.jpg';

    if (!file_exists($path)) { $skip++; continue; }

    foreach ($imageTypes as $type) {
        $dest = __DIR__ . '/img/p/' . implode('/', str_split((string)$id)) . '/' . $id . '-' . $type['name'] . '.jpg';
        $result = ImageManager::resize($path, $dest, (int)$type['width'], (int)$type['height'], 'jpg');
        $result ? $ok++ : ($errors++ || print("ERROR: image $id, type {$type['name']}\n"));
    }
}
echo "Done. Generated: $ok | Skipped (no source): $skip | Errors: $errors\n";
```

Ejecutar:
```bash
# Local Lando
lando php regen_thumbs.php

# Servidor remoto
scp -i ~/.ssh/id_rsa_cinetic regen_thumbs.php user@host:/path/to/ps/regen_thumbs.php
ssh -i ~/.ssh/id_rsa_cinetic user@host "cd /path/to/ps && php regen_thumbs.php"
ssh -i ~/.ssh/id_rsa_cinetic user@host "rm /path/to/ps/regen_thumbs.php"
```

El script usa `ImageManager::resize()` — exactamente lo mismo que el admin, pero:
- Salta imágenes sin fuente local (no falla)
- Corre en CLI, sin timeout de Apache
- Muestra conteo final: Generated / Skipped / Errors

## Paso 3 — Sincronizar imágenes fuente desde servidor (opcional)

Si en local faltan imágenes fuente (img/p/123.jpg), el script las saltará. Para descargar
solo las fuentes (sin thumbnails, que se regeneran):

```bash
rsync -avz --progress \
  -e "ssh -i ~/.ssh/id_rsa_cinetic" \
  --filter='+ */' \
  --filter='+ [0-9][0-9][0-9][0-9][0-9].jpg' \
  --filter='+ [0-9][0-9][0-9][0-9].jpg' \
  --filter='+ [0-9][0-9][0-9].jpg' \
  --filter='- *' \
  user@host:/path/to/ps/img/p/ \
  ./img/p/
```

Los filtros incluyen solo archivos como `290.jpg` (solo dígitos), excluyendo
`290-home_default.jpg` (thumbnails).

## Paso 4 — CSS para grid uniforme (post-regen)

Con método 1/2, las imágenes generadas son proporcionales (no cuadradas). El grid
quedará desigual si no se fuerza ratio fijo. Añadir en el child theme:

```css
/* themes/amicstheme/assets/css/custom.css */
a.product_img_link {
    display: block;
    position: relative;
    aspect-ratio: 1 / 1;
    overflow: hidden;
}
a.product_img_link picture.front_image_pic {
    display: block;
    width: 100%;
    height: 100%;
}
a.product_img_link picture.front_image_pic img {
    width: 100%;
    height: 100%;
    object-fit: cover;
}
```

Ajustar `aspect-ratio` según el diseño (ej. `4 / 3`, `16 / 9`).

> **Selector del tema Panda**: el contenedor de imagen en listings es
> `a.product_img_link > picture.front_image_pic > img`.
> En otros temas PS puede ser `.thumbnail-container img` o `.product-thumbnail img`.

## Notas

- `PS_IMAGE_GENERATION_METHOD = 0` es el default de PS — nunca se inserta en DB
  hasta que se cambia desde el admin. Por eso `SELECT` puede devolver vacío.
- El script no toca las imágenes fuente (`290.jpg`), solo sobreescribe thumbnails.
- Si el proceso Apache se cuelga regenerando desde el admin, matar con:
  `lando ssh -s appserver -c "kill <PID>"` (buscar proceso con >50% CPU en `ps aux`).
- Método 1 (Amplada) es el más seguro: escala por anchura, la altura es proporcional.
  Método 2 (Alçada) escala por altura. Ambos eliminan el padding blanco.
