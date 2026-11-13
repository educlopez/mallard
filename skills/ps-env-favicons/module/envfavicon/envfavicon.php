<?php
/**
 * Env Favicon · favicon distinto por entorno (local / PRE) sin tocar el tema.
 * -----------------------------------------------------------------------------
 * - Producción (y cualquier host que no case con las listas) NO se toca: sigue el
 *   favicon nativo de PrestaShop (Diseño › Tema y logotipo › Favicon).
 * - Local y PRE: se sustituyen los <link rel="icon"> de la <head> por una insignia
 *   de color con la marca del cliente. Un hook de salida (actionOutputHTMLBefore),
 *   sin plantillas, sin copiar head.tpl del tema padre.
 * - Los dominios de cada entorno se guardan en la base de datos (BO), no en ficheros
 *   públicos del tema.
 * - Se sube UN logo/marca (PNG/JPG/WEBP) y el módulo genera las insignias y el .ico de
 *   producción con GD, y puede cargarlo como favicon nativo de PrestaShop.
 *
 * Compatibilidad: PrestaShop 1.7.7 → 9 (PHP 7.2+, sin typed properties).
 *
 * @author Cinetic Digital
 */
if (!defined('_PS_VERSION_')) {
    exit;
}

class EnvFavicon extends Module
{
    const K_ENABLED = 'ENVFAV_ENABLED';
    const K_LOCAL = 'ENVFAV_LOCAL_PATTERNS';
    const K_PRE = 'ENVFAV_PRE_PATTERNS';
    const K_COLOR_LOCAL = 'ENVFAV_COLOR_LOCAL';
    const K_COLOR_PRE = 'ENVFAV_COLOR_PRE';
    const K_COLOR_PROD = 'ENVFAV_COLOR_PROD';
    const K_CROP = 'ENVFAV_CROP_MARK';
    const K_SOURCE = 'ENVFAV_SOURCE_FILE';
    const K_VERSION = 'ENVFAV_ASSET_VERSION';

    public function __construct()
    {
        $this->name = 'envfavicon';
        $this->tab = 'front_office_features';
        $this->version = '0.1.0';
        $this->author = 'Cinetic Digital';
        $this->need_instance = 0;
        $this->bootstrap = true;
        $this->ps_versions_compliancy = array('min' => '1.7.7.0', 'max' => _PS_VERSION_);

        parent::__construct();

        $this->displayName = $this->l('Env Favicon · favicon por entorno');
        $this->description = $this->l('Muestra un favicon distinto en local y PRE para distinguir las pestañas. Producción conserva el favicon nativo de PrestaShop.');
    }

    /* ============================================================
     * Instalación
     * ============================================================ */

    public function install()
    {
        if (!parent::install() || !$this->registerHook('actionOutputHTMLBefore')) {
            return false;
        }
        foreach ($this->defaults() as $k => $v) {
            if (Configuration::get($k) === false) {
                Configuration::updateValue($k, $v);
            }
        }

        return true;
    }

    public function uninstall()
    {
        foreach (array_keys($this->defaults()) as $k) {
            Configuration::deleteByName($k);
        }
        foreach (glob($this->imgDir() . '{source.*,favicon-*.png,favicon-production.ico}', GLOB_BRACE) ?: array() as $f) {
            @unlink($f);
        }

        return parent::uninstall();
    }

    private function defaults()
    {
        return array(
            self::K_ENABLED => 1,
            self::K_LOCAL => 'lndo.site, localhost, 127.0.0.1, .test, .local',
            self::K_PRE => 'pre-testing.com, ecommercloud',
            self::K_COLOR_LOCAL => '#1e8e4e',
            self::K_COLOR_PRE => '#996b2d',
            self::K_COLOR_PROD => '#0e1926',
            self::K_CROP => 1,
            self::K_SOURCE => '',
            self::K_VERSION => 1,
        );
    }

    private function imgDir()
    {
        return rtrim($this->getLocalPath(), '/') . '/views/img/';
    }

    /* ============================================================
     * Detección de entorno
     * ============================================================ */

    /**
     * 'local' | 'pre' | 'prod'. Solo compara el host con las listas del BO (stripos);
     * lo que no case es producción (fallo seguro: nunca una insignia de dev en prod).
     */
    public static function detectEnv($host = null)
    {
        $host = strtolower((string) ($host === null ? Tools::getHttpHost(false, false, true) : $host));
        if ($host === '') {
            return 'prod';
        }
        foreach (array('local' => self::K_LOCAL, 'pre' => self::K_PRE) as $env => $key) {
            foreach (self::parsePatterns((string) Configuration::get($key)) as $p) {
                if (strpos($host, $p) !== false) {
                    return $env;
                }
            }
        }

        return 'prod';
    }

    public static function parsePatterns($raw)
    {
        $out = array();
        foreach (preg_split('/[\s,;]+/', strtolower($raw)) as $p) {
            $p = trim($p);
            if ($p !== '' && strlen($p) <= 100 && preg_match('/^[a-z0-9.\-_:]+$/', $p)) {
                $out[] = $p;
            }
        }

        return array_values(array_unique($out));
    }

    /* ============================================================
     * Hook de salida: sustituir los iconos de la <head>
     * ============================================================ */

    public function hookActionOutputHTMLBefore($params)
    {
        if (!isset($params['html']) || !Configuration::get(self::K_ENABLED)) {
            return;
        }
        $env = self::detectEnv();
        if ($env === 'prod') {
            return;
        }
        $file = 'favicon-' . $env . '.png';
        if (!is_file($this->imgDir() . $file)) {
            return; // aún no se ha generado: se deja el favicon nativo
        }
        $html = (string) $params['html'];
        $end = stripos($html, '</head>');
        if ($end === false) {
            return;
        }
        $head = substr($html, 0, $end);
        $head = preg_replace('#<link\b[^>]*\brel=["\'](?:shortcut )?icon["\'][^>]*>\s*#i', '', $head);
        $url = $this->getPathUri() . 'views/img/' . $file . '?v=' . (int) Configuration::get(self::K_VERSION);
        $links = '<link rel="icon" type="image/png" sizes="64x64" href="' . $url . '">'
            . '<link rel="shortcut icon" type="image/png" href="' . $url . '">';
        $params['html'] = $head . $links . substr($html, $end);
    }

    /* ============================================================
     * Back office
     * ============================================================ */

    public function getContent()
    {
        $out = '';

        if (Tools::isSubmit('submitEnvFavicon')) {
            Configuration::updateValue(self::K_ENABLED, (int) Tools::getValue(self::K_ENABLED));
            Configuration::updateValue(self::K_CROP, (int) Tools::getValue(self::K_CROP));
            Configuration::updateValue(self::K_LOCAL, implode(', ', self::parsePatterns((string) Tools::getValue(self::K_LOCAL))));
            Configuration::updateValue(self::K_PRE, implode(', ', self::parsePatterns((string) Tools::getValue(self::K_PRE))));
            foreach (array(self::K_COLOR_LOCAL, self::K_COLOR_PRE, self::K_COLOR_PROD) as $k) {
                $c = (string) Tools::getValue($k);
                if (preg_match('/^#[0-9a-fA-F]{6}$/', $c)) {
                    Configuration::updateValue($k, $c);
                }
            }
            $out .= $this->displayConfirmation($this->l('Configuración guardada.'));

            $error = $this->handleUpload();
            if ($error) {
                $out .= $this->displayError($error);
            } elseif ($this->hasSource()) {
                $err = $this->generate();
                $out .= $err ? $this->displayError($err) : $this->displayConfirmation($this->l('Favicons regenerados.'));
            }
        }

        if (Tools::isSubmit('submitEnvFaviconNative')) {
            $err = $this->applyNative();
            $out .= $err ? $this->displayError($err) : $this->displayConfirmation($this->l('Favicon de producción cargado como favicon nativo de PrestaShop.'));
        }

        return $out . $this->renderForm() . $this->renderPreview();
    }

    private function hasSource()
    {
        $f = (string) Configuration::get(self::K_SOURCE);

        return $f !== '' && is_file($this->imgDir() . $f);
    }

    private function handleUpload()
    {
        if (empty($_FILES['ENVFAV_SOURCE']['name'])) {
            return '';
        }
        $file = $_FILES['ENVFAV_SOURCE'];
        if ($err = ImageManager::validateUpload($file, 3000000, array('png', 'jpg', 'jpeg', 'webp'))) {
            return $err;
        }
        $ext = strtolower(pathinfo($file['name'], PATHINFO_EXTENSION));
        $ext = $ext === 'jpeg' ? 'jpg' : $ext;
        foreach (glob($this->imgDir() . 'source.*') ?: array() as $old) {
            @unlink($old);
        }
        $dest = $this->imgDir() . 'source.' . $ext;
        if (!move_uploaded_file($file['tmp_name'], $dest)) {
            return $this->l('No se pudo guardar el archivo subido.');
        }
        Configuration::updateValue(self::K_SOURCE, 'source.' . $ext);

        return '';
    }

    private function renderForm()
    {
        $sw = array(
            array('id' => 'on', 'value' => 1, 'label' => $this->l('Sí')),
            array('id' => 'off', 'value' => 0, 'label' => $this->l('No')),
        );
        $env = self::detectEnv();
        $names = array('local' => $this->l('Local'), 'pre' => $this->l('PRE'), 'prod' => $this->l('Producción (favicon nativo)'));
        $host = Tools::safeOutput(Tools::getHttpHost(false, false, true));

        $fields = array(
            array('type' => 'free', 'label' => $this->l('Este host se detecta como'), 'name' => 'ENVFAV_CURRENT'),
            array('type' => 'switch', 'label' => $this->l('Activo'), 'name' => self::K_ENABLED, 'is_bool' => true, 'values' => $sw),
            array('type' => 'textarea', 'label' => $this->l('Dominios LOCAL'), 'name' => self::K_LOCAL, 'rows' => 2, 'desc' => $this->l('Fragmentos de host separados por coma. Si el host contiene alguno, es local.')),
            array('type' => 'textarea', 'label' => $this->l('Dominios PRE'), 'name' => self::K_PRE, 'rows' => 2, 'desc' => $this->l('Ej.: pre-testing.com, ecommercloud. Todo lo que no case con LOCAL ni PRE es producción y usa el favicon nativo.')),
            array('type' => 'color', 'label' => $this->l('Color LOCAL'), 'name' => self::K_COLOR_LOCAL),
            array('type' => 'color', 'label' => $this->l('Color PRE'), 'name' => self::K_COLOR_PRE),
            array('type' => 'color', 'label' => $this->l('Color PRODUCCIÓN (favicon .ico)'), 'name' => self::K_COLOR_PROD),
            array('type' => 'file', 'label' => $this->l('Logo o marca (PNG, JPG, WEBP)'), 'name' => 'ENVFAV_SOURCE', 'desc' => $this->l('Un solo archivo: se generan todas las insignias. Mejor con fondo transparente o marca oscura sobre blanco.')),
            array('type' => 'switch', 'label' => $this->l('Recortar solo la marca'), 'name' => self::K_CROP, 'is_bool' => true, 'values' => $sw, 'desc' => $this->l('Para logos "marca + texto": usa solo el primer bloque (a la izquierda) y quita el ® y las motas.')),
        );
        $form = array('form' => array(
            'legend' => array('title' => $this->l('Favicon por entorno'), 'icon' => 'icon-picture'),
            'input' => $fields,
            'submit' => array('title' => $this->l('Guardar y regenerar'), 'name' => 'submitEnvFavicon'),
            'buttons' => array(array(
                'title' => $this->l('Cargar el de producción como favicon nativo'),
                'name' => 'submitEnvFaviconNative',
                'type' => 'submit',
                'class' => 'btn btn-default pull-left',
                'icon' => 'process-icon-upload',
            )),
        ));

        $helper = new HelperForm();
        $helper->module = $this;
        $helper->name_controller = $this->name;
        $helper->token = Tools::getAdminTokenLite('AdminModules');
        $helper->currentIndex = AdminController::$currentIndex . '&configure=' . $this->name;
        $helper->submit_action = 'submitEnvFavicon';
        $helper->default_form_language = (int) Context::getContext()->language->id;

        $values = array();
        foreach ($this->defaults() as $k => $d) {
            $cur = Configuration::get($k);
            $values[$k] = ($cur === false) ? $d : $cur;
        }
        $values['ENVFAV_SOURCE'] = '';
        $values['ENVFAV_CURRENT'] = '<p class="form-control-static"><b>' . Tools::safeOutput($names[$env]) . '</b> &nbsp;<small class="text-muted">(' . $host . ')</small></p>';
        $helper->fields_value = $values;

        return $helper->generateForm(array($form));
    }

    private function renderPreview()
    {
        $base = $this->getPathUri() . 'views/img/';
        $v = (int) Configuration::get(self::K_VERSION);
        $items = array(
            'favicon-local.png' => $this->l('Local'),
            'favicon-pre.png' => $this->l('PRE'),
            'favicon-production.png' => $this->l('Producción'),
        );
        $html = '<div class="panel"><div class="panel-heading">' . $this->l('Vista previa') . '</div><div class="row">';
        $any = false;
        foreach ($items as $file => $label) {
            if (!is_file($this->imgDir() . $file)) {
                continue;
            }
            $any = true;
            $html .= '<div class="col-xs-4 text-center"><img src="' . Tools::safeOutput($base . $file . '?v=' . $v) . '" width="64" height="64" alt=""><br><small>' . $label . '</small></div>';
        }
        if (!$any) {
            $html .= '<div class="col-xs-12"><p class="text-muted">' . $this->l('Sube un logo y guarda para generar las insignias.') . '</p></div>';
        }

        return $html . '</div></div>';
    }

    /* ============================================================
     * Generación con GD
     * ============================================================ */

    /**
     * Genera favicon-local.png, favicon-pre.png (64 px), favicon-production.png (preview)
     * y favicon-production.ico (16–64). Devuelve '' si todo fue bien o el error.
     */
    public function generate()
    {
        if (!extension_loaded('gd')) {
            return $this->l('Hace falta la extensión GD de PHP.');
        }
        $src = $this->imgDir() . (string) Configuration::get(self::K_SOURCE);
        $mark = $this->extractMark($src, (bool) Configuration::get(self::K_CROP));
        if (!$mark) {
            return $this->l('No se pudo leer el logo o no se encontró la marca.');
        }
        $colors = array(
            'local' => (string) Configuration::get(self::K_COLOR_LOCAL),
            'pre' => (string) Configuration::get(self::K_COLOR_PRE),
            'production' => (string) Configuration::get(self::K_COLOR_PROD),
        );
        foreach ($colors as $env => $hex) {
            $big = $this->badge($mark, $hex, 256);
            imagepng($this->resize($big, 64), $this->imgDir() . 'favicon-' . $env . '.png');
            if ($env === 'production') {
                $png = array();
                foreach (array(16, 24, 32, 48, 64) as $s) {
                    ob_start();
                    imagepng($this->resize($big, $s));
                    $png[$s] = ob_get_clean();
                }
                file_put_contents($this->imgDir() . 'favicon-production.ico', $this->buildIco($png));
            }
            imagedestroy($big);
        }
        imagedestroy($mark);
        Configuration::updateValue(self::K_VERSION, (int) Configuration::get(self::K_VERSION) + 1);

        return '';
    }

    /**
     * Imagen GD (RGBA, marca blanca con el alfa de la tinta) con la marca recortada.
     * Tinta = alfa del PNG, o lo oscuro del píxel si no hay transparencia.
     */
    private function extractMark($path, $crop)
    {
        $info = @getimagesize($path);
        if (!$info) {
            return null;
        }
        $img = $info[2] === IMAGETYPE_PNG ? @imagecreatefrompng($path)
            : ($info[2] === IMAGETYPE_JPEG ? @imagecreatefromjpeg($path)
            : (function_exists('imagecreatefromwebp') && $info[2] === IMAGETYPE_WEBP ? @imagecreatefromwebp($path) : false));
        if (!$img) {
            return null;
        }
        $w0 = imagesx($img);
        $h0 = imagesy($img);
        $scale = min(1, 700 / max($w0, $h0));
        $w = max(1, (int) round($w0 * $scale));
        $h = max(1, (int) round($h0 * $scale));
        $small = imagecreatetruecolor($w, $h);
        imagealphablending($small, false);
        imagesavealpha($small, true);
        imagefill($small, 0, 0, imagecolorallocatealpha($small, 255, 255, 255, 127));
        imagecopyresampled($small, $img, 0, 0, 0, 0, $w, $h, $w0, $h0);
        imagedestroy($img);

        $hasAlpha = false;
        for ($x = 0; $x < $w && !$hasAlpha; $x += 5) {
            for ($y = 0; $y < $h; $y += 5) {
                if (((imagecolorat($small, $x, $y) >> 24) & 127) > 8) {
                    $hasAlpha = true;
                    break;
                }
            }
        }
        $ink = array();
        for ($x = 0; $x < $w; $x++) {
            for ($y = 0; $y < $h; $y++) {
                $c = imagecolorat($small, $x, $y);
                $a = 255 - (($c >> 24) & 127) * 2;
                if ($hasAlpha) {
                    $v = $a;
                } else {
                    $v = (int) round((255 - ((($c >> 16) & 255) + (($c >> 8) & 255) + ($c & 255)) / 3) * ($a / 255));
                }
                $ink[$x][$y] = $v;
            }
        }
        imagedestroy($small);

        $x0 = 0;
        $x1 = $w;
        if ($crop) {
            $runs = array();
            $start = null;
            for ($x = 0; $x <= $w; $x++) {
                $has = false;
                if ($x < $w) {
                    for ($y = 0; $y < $h; $y++) {
                        if ($ink[$x][$y] > 128) {
                            $has = true;
                            break;
                        }
                    }
                }
                if ($has && $start === null) {
                    $start = $x;
                }
                if (!$has && $start !== null) {
                    $runs[] = array($start, $x);
                    $start = null;
                }
            }
            $groups = array();
            foreach ($runs as $r) {
                $last = count($groups) - 1;
                if ($last >= 0 && $r[0] - $groups[$last][1] < $h * 0.2) {
                    $groups[$last][1] = $r[1];
                } else {
                    $groups[] = $r;
                }
            }
            foreach ($groups as $g) {
                if ($g[1] - $g[0] > $h * 0.5) {
                    list($x0, $x1) = $g;
                    break;
                }
            }
        }
        $keep = $this->keepLargeComponents($ink, $x0, $x1, $h);
        $minX = $w;
        $maxX = -1;
        $minY = $h;
        $maxY = -1;
        foreach ($keep as $k => $v) {
            list($x, $y) = explode(',', $k);
            $minX = min($minX, (int) $x);
            $maxX = max($maxX, (int) $x);
            $minY = min($minY, (int) $y);
            $maxY = max($maxY, (int) $y);
        }
        if ($maxX < 0) {
            return null;
        }
        $mw = $maxX - $minX + 1;
        $mh = $maxY - $minY + 1;
        $mark = imagecreatetruecolor($mw, $mh);
        imagealphablending($mark, false);
        imagesavealpha($mark, true);
        imagefill($mark, 0, 0, imagecolorallocatealpha($mark, 255, 255, 255, 127));
        foreach ($keep as $k => $v) {
            list($x, $y) = explode(',', $k);
            imagesetpixel($mark, (int) $x - $minX, (int) $y - $minY, imagecolorallocatealpha($mark, 255, 255, 255, 127 - (int) round($v / 2)));
        }

        return $mark;
    }

    /**
     * Píxeles (clave "x,y" => tinta) de los componentes conexos grandes dentro de [x0,x1),
     * con 2 px de margen alrededor para conservar el antialiasing; descarta el ® y las motas.
     */
    private function keepLargeComponents(array $ink, $x0, $x1, $h)
    {
        $seen = array();
        $comps = array();
        for ($sx = $x0; $sx < $x1; $sx++) {
            for ($sy = 0; $sy < $h; $sy++) {
                if ($ink[$sx][$sy] > 128 && !isset($seen[$sx . ',' . $sy])) {
                    $stack = array(array($sx, $sy));
                    $seen[$sx . ',' . $sy] = true;
                    $pts = array();
                    while ($stack) {
                        list($cx, $cy) = array_pop($stack);
                        $pts[] = array($cx, $cy);
                        foreach (array(array($cx + 1, $cy), array($cx - 1, $cy), array($cx, $cy + 1), array($cx, $cy - 1)) as $n) {
                            if ($n[0] >= $x0 && $n[0] < $x1 && $n[1] >= 0 && $n[1] < $h && !isset($seen[$n[0] . ',' . $n[1]]) && $ink[$n[0]][$n[1]] > 128) {
                                $seen[$n[0] . ',' . $n[1]] = true;
                                $stack[] = $n;
                            }
                        }
                    }
                    $comps[] = $pts;
                }
            }
        }
        $keep = array();
        if (!$comps) {
            return $keep;
        }
        $biggest = max(array_map('count', $comps));
        foreach ($comps as $pts) {
            if (count($pts) < $biggest * 0.2) {
                continue;
            }
            foreach ($pts as $p) {
                for ($dx = -2; $dx <= 2; $dx++) {
                    for ($dy = -2; $dy <= 2; $dy++) {
                        $x = $p[0] + $dx;
                        $y = $p[1] + $dy;
                        if ($x >= $x0 && $x < $x1 && $y >= 0 && $y < $h && !isset($keep[$x . ',' . $y]) && $ink[$x][$y] > 0) {
                            $keep[$x . ',' . $y] = $ink[$x][$y];
                        }
                    }
                }
            }
        }

        return $keep;
    }

    /** Cuadrado redondeado de color + marca blanca centrada (ajustada por ancho, ~74 %). */
    private function badge($mark, $hex, $side)
    {
        $ss = 4; // supermuestreo para bordes suaves
        $big = $side * $ss;
        $img = imagecreatetruecolor($big, $big);
        imagealphablending($img, false);
        imagesavealpha($img, true);
        imagefill($img, 0, 0, imagecolorallocatealpha($img, 0, 0, 0, 127));
        $col = imagecolorallocate($img, hexdec(substr($hex, 1, 2)), hexdec(substr($hex, 3, 2)), hexdec(substr($hex, 5, 2)));
        $r = (int) round($big * 0.22);
        imagefilledrectangle($img, $r, 0, $big - $r, $big, $col);
        imagefilledrectangle($img, 0, $r, $big, $big - $r, $col);
        foreach (array(array($r, $r), array($big - $r, $r), array($r, $big - $r), array($big - $r, $big - $r)) as $c) {
            imagefilledellipse($img, $c[0], $c[1], $r * 2, $r * 2, $col);
        }
        $mw = imagesx($mark);
        $mh = imagesy($mark);
        $gw = (int) round($big * 0.74);
        $gh = (int) round($mh * $gw / $mw);
        if ($gh > $big * 0.74) {
            $gh = (int) round($big * 0.74);
            $gw = (int) round($mw * $gh / $mh);
        }
        imagealphablending($img, true);
        imagecopyresampled($img, $mark, (int) (($big - $gw) / 2), (int) (($big - $gh) / 2), 0, 0, $gw, $gh, $mw, $mh);
        $out = $this->resize($img, $side);
        imagedestroy($img);

        return $out;
    }

    private function resize($img, $size)
    {
        $out = imagecreatetruecolor($size, $size);
        imagealphablending($out, false);
        imagesavealpha($out, true);
        imagefill($out, 0, 0, imagecolorallocatealpha($out, 0, 0, 0, 127));
        imagecopyresampled($out, $img, 0, 0, 0, 0, $size, $size, imagesx($img), imagesy($img));

        return $out;
    }

    /** .ico con imágenes PNG embebidas (válido en todos los navegadores actuales). */
    private function buildIco(array $pngBySize)
    {
        $n = count($pngBySize);
        $head = pack('vvv', 0, 1, $n);
        $dir = '';
        $data = '';
        $offset = 6 + 16 * $n;
        foreach ($pngBySize as $size => $png) {
            $dir .= pack('CCCCvvVV', $size, $size, 0, 0, 1, 32, strlen($png), $offset);
            $data .= $png;
            $offset += strlen($png);
        }

        return $head . $dir . $data;
    }

    /**
     * Carga favicon-production.ico como favicon nativo, como hace el BO
     * (LogoUploader::updateFavicon): img/favicon.ico + img/favicon-<id_shop>.ico + PS_FAVICON.
     */
    public function applyNative()
    {
        $ico = $this->imgDir() . 'favicon-production.ico';
        if (!is_file($ico)) {
            return $this->l('Primero sube un logo y guarda para generar el favicon de producción.');
        }
        $shopId = (int) Context::getContext()->shop->id;
        if (!@copy($ico, _PS_IMG_DIR_ . 'favicon.ico') || !@copy($ico, _PS_IMG_DIR_ . 'favicon-' . $shopId . '.ico')) {
            return $this->l('No se pudo escribir en el directorio img/ (permisos).');
        }
        Configuration::updateValue('PS_FAVICON', 'favicon-' . $shopId . '.ico');
        Configuration::updateGlobalValue('PS_FAVICON', 'favicon.ico');
        Configuration::updateValue('PS_IMG_UPDATE_TIME', time());

        return '';
    }
}
