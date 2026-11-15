<?php
/**
 * Env Favicon - a different favicon per environment (local / staging), no theme edits.
 *
 * - Production (and any host that matches no list) is NEVER touched: it keeps the
 *   native PrestaShop favicon (Design > Theme & Logo > Favicon).
 * - Local and staging: the <link rel="icon"> tags of the <head> are replaced by a
 *   coloured badge with the client's brand mark, through an output hook
 *   (actionOutputHTMLBefore): no templates, no copy of the parent theme's head.tpl.
 * - The domains of each environment live in the database (back office), not in
 *   public theme files.
 * - ONE logo is uploaded; the module builds the badges and the production .ico with
 *   GD, and can load it as the native PrestaShop favicon.
 *
 * Compatibility: PrestaShop 1.7.7 -> 9 (PHP 7.2+, no typed properties).
 *
 * NOTICE OF LICENSE
 *
 * This source file is subject to the Academic Free License 3.0 (AFL-3.0)
 * that is bundled with this package in the file LICENSE.md.
 *
 * @author    Eduardo Calvo
 * @copyright Eduardo Calvo
 * @license   https://opensource.org/licenses/AFL-3.0 Academic Free License 3.0
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
        $this->version = '0.2.0';
        $this->author = 'Eduardo Calvo';
        $this->need_instance = 0;
        $this->bootstrap = true;
        $this->ps_versions_compliancy = array('min' => '1.7.7.0', 'max' => _PS_VERSION_);

        parent::__construct();

        $this->displayName = $this->l('Env Favicon');
        $this->description = $this->l('Shows a different favicon on local and staging so you can tell the tabs apart. Production keeps the native PrestaShop favicon.');
    }

    /* ============================================================
     * Install
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
        $files = glob($this->imgDir() . '{source.*,favicon-*.png,favicon-production.ico,mark-white.png}', GLOB_BRACE);
        foreach ($files ? $files : array() as $f) {
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
     * Environment detection
     * ============================================================ */

    /**
     * 'local' | 'pre' | 'prod'. The host is only compared with the lists saved in the
     * back office (strpos); anything that matches nothing is production (fail safe:
     * a dev badge is never shown on prod).
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
     * Output hook: replace the icons of the <head>
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
            return; // not generated yet: the native favicon stays
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

            $error = $this->handleUpload();
            if ($error) {
                $out .= $this->displayError($error);
            } elseif ($this->hasSource()) {
                $err = $this->generate();
                $out .= $err ? $this->displayError($err) : $this->displayConfirmation($this->l('Saved. The favicons have been regenerated.'));
            } else {
                $out .= $this->displayConfirmation($this->l('Saved. Upload a logo to generate the favicons.'));
            }
        }

        if (Tools::isSubmit('submitEnvFaviconNative')) {
            $err = $this->applyNative();
            $out .= $err ? $this->displayError($err) : $this->displayConfirmation($this->l('The production favicon is now the native PrestaShop favicon.'));
        }

        return $out . $this->renderPage();
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
        $error = ImageManager::validateUpload($file, 3000000, array('png', 'jpg', 'jpeg', 'webp'));
        if ($error) {
            return $error;
        }
        $ext = strtolower(pathinfo($file['name'], PATHINFO_EXTENSION));
        $ext = $ext === 'jpeg' ? 'jpg' : $ext;
        $old = glob($this->imgDir() . 'source.*');
        foreach ($old ? $old : array() as $o) {
            @unlink($o);
        }
        if (!move_uploaded_file($file['tmp_name'], $this->imgDir() . 'source.' . $ext)) {
            return $this->l('The uploaded file could not be saved.');
        }
        Configuration::updateValue(self::K_SOURCE, 'source.' . $ext);

        return '';
    }

    private function renderPage()
    {
        $ctx = $this->context;
        $ctx->controller->addCSS($this->_path . 'views/css/admin.css?v=' . $this->version);
        $ctx->controller->addJS($this->_path . 'views/js/admin.js?v=' . $this->version);

        $base = $this->getPathUri() . 'views/img/';
        $v = (int) Configuration::get(self::K_VERSION);
        $val = array();
        foreach ($this->defaults() as $k => $d) {
            $cur = Configuration::get($k);
            $val[$k] = ($cur === false) ? $d : $cur;
        }
        $hasMark = is_file($this->imgDir() . 'mark-white.png');

        $ctx->smarty->assign(array(
            'ef_action' => AdminController::$currentIndex . '&configure=' . $this->name . '&token=' . Tools::getAdminTokenLite('AdminModules'),
            'ef_host' => Tools::getHttpHost(false, false, true),
            'ef_env' => self::detectEnv(),
            'ef_v' => $val,
            'ef_has_mark' => $hasMark,
            'ef_mark_url' => $hasMark ? $base . 'mark-white.png?v=' . $v : '',
            'ef_has_source' => $this->hasSource(),
            'ef_has_ico' => is_file($this->imgDir() . 'favicon-production.ico'),
            'ef_module_logo' => $this->getPathUri() . 'logo.png',
            'ef_shop_name' => Configuration::get('PS_SHOP_NAME'),
        ));

        return $this->display(__FILE__, 'views/templates/admin/configure.tpl');
    }

    /* ============================================================
     * Generation with GD
     * ============================================================ */

    /**
     * Writes favicon-local.png, favicon-pre.png (64 px), favicon-production.png (preview),
     * favicon-production.ico (16-64) and mark-white.png (live preview in the back office).
     * Returns '' on success or the error message.
     */
    public function generate()
    {
        if (!extension_loaded('gd')) {
            return $this->l('The PHP GD extension is required.');
        }
        $src = $this->imgDir() . (string) Configuration::get(self::K_SOURCE);
        $mark = $this->extractMark($src, (bool) Configuration::get(self::K_CROP));
        if (!$mark) {
            return $this->l('The logo could not be read or no mark was found in it.');
        }
        imagepng($mark, $this->imgDir() . 'mark-white.png');
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
     * GD image (RGBA, white mark with the ink as alpha) of the cropped mark.
     * Ink = PNG alpha, or the darkness of the pixel when there is no transparency.
     */
    private function extractMark($path, $crop)
    {
        $info = @getimagesize($path);
        if (!$info) {
            return null;
        }
        $img = false;
        if ($info[2] === IMAGETYPE_PNG) {
            $img = @imagecreatefrompng($path);
        } elseif ($info[2] === IMAGETYPE_JPEG) {
            $img = @imagecreatefromjpeg($path);
        } elseif ($info[2] === IMAGETYPE_WEBP && function_exists('imagecreatefromwebp')) {
            $img = @imagecreatefromwebp($path);
        }
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
     * Pixels (key "x,y" => ink) of the large connected components inside [x0,x1), with a
     * 2 px margin to keep the anti-aliasing; drops the (R) sign and specks.
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

    /** Rounded colour square + white mark centred (fitted by width, ~74 %). */
    private function badge($mark, $hex, $side)
    {
        $ss = 4; // supersampling for smooth edges
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

    /** .ico with embedded PNG images (valid in every current browser). */
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
     * Loads favicon-production.ico as the native favicon, as the back office does
     * (LogoUploader::updateFavicon): img/favicon.ico + img/favicon-<id_shop>.ico + PS_FAVICON.
     */
    public function applyNative()
    {
        $ico = $this->imgDir() . 'favicon-production.ico';
        if (!is_file($ico)) {
            return $this->l('Upload a logo and save first to generate the production favicon.');
        }
        $shopId = (int) Context::getContext()->shop->id;
        if (!@copy($ico, _PS_IMG_DIR_ . 'favicon.ico') || !@copy($ico, _PS_IMG_DIR_ . 'favicon-' . $shopId . '.ico')) {
            return $this->l('Could not write to the img/ directory (permissions).');
        }
        Configuration::updateValue('PS_FAVICON', 'favicon-' . $shopId . '.ico');
        Configuration::updateGlobalValue('PS_FAVICON', 'favicon.ico');
        Configuration::updateValue('PS_IMG_UPDATE_TIME', time());

        return '';
    }
}
