<?php
/**
 * PrestaShop Invoice PDF Viewer (local dev only)
 * Access: {LANDO_URL}/invoice-preview/
 *
 * Lista pedidos y renderiza el PDF de factura (classes/pdf/HTMLTemplateInvoice.php
 * + pdf/*.tpl + themes/{child}/pdf/*.tpl override, resuelto por PS mismo) directo
 * en el navegador, sin pasar por el backoffice. Útil para validar cambios de
 * maquetación rápido. No requiere ningún placeholder: todo se auto-detecta en
 * runtime (admin dir, URL, nombre de tienda, credenciales DB).
 */

define('ROOT', dirname(__DIR__));
define('LANDO_URL', (isset($_SERVER['HTTPS']) && $_SERVER['HTTPS'] !== 'off' ? 'https' : 'http') . '://' . $_SERVER['HTTP_HOST']);

// ── Protección: solo empleados logueados en el backoffice ─────────────────────

function denyAccess(string $title, string $msg, int $code = 403): never {
    http_response_code($code);
    echo '<!DOCTYPE html><html lang="es"><head><meta charset="UTF-8"><link href="https://fonts.googleapis.com/css2?family=Geist:wght@400;500;600&display=swap" rel="stylesheet"><style>*{box-sizing:border-box;margin:0;padding:0}body{font-family:Geist,-apple-system,sans-serif;background:#fafafa;display:flex;align-items:center;justify-content:center;min-height:100vh}.card{background:#fff;border:1px solid #eaeaea;border-radius:12px;padding:40px 48px;text-align:center;max-width:360px;width:100%}h1{font-size:15px;font-weight:600;color:#000;margin-bottom:8px}p{font-size:13px;color:#666;line-height:1.6}</style></head><body><div class="card"><h1>' . htmlspecialchars($title) . '</h1><p>' . htmlspecialchars($msg) . '</p></div></body></html>';
    exit;
}

// Detecta la carpeta del admin (a menudo renombrada por seguridad): busca la
// primera carpeta de primer nivel que tenga init.php + bootstrap.php.
function findAdminDir(): string {
    foreach (glob(ROOT . '/*', GLOB_ONLYDIR) as $d) {
        if (file_exists($d . '/init.php') && file_exists($d . '/bootstrap.php')) {
            return basename($d);
        }
    }
    return 'admin';
}

function requirePsAdminAuth(): array {
    $psConfig = ROOT . '/config/config.inc.php';
    if (!file_exists($psConfig)) denyAccess('No disponible', 'PrestaShop no encontrado en esta ruta.', 503);
    try {
        ob_start();
        @require_once $psConfig;
        ob_end_clean();
        if (!class_exists('Cookie')) denyAccess('No disponible', 'El sistema no pudo iniciarse. Comprueba que la BD esté activa.', 503);
        $cookieName = defined('_COOKIE_ADMIN_') ? _COOKIE_ADMIN_ : 'psAdmin';
        $cookie = new Cookie($cookieName);
        if (empty($cookie->id_employee)) {
            $loginUrl = LANDO_URL . '/' . findAdminDir() . '/index.php';
            http_response_code(403);
            echo '<!DOCTYPE html><html lang="es"><head><meta charset="UTF-8"><title>Acceso restringido</title>'
               . '<link href="https://fonts.googleapis.com/css2?family=Geist:wght@400;500;600&display=swap" rel="stylesheet">'
               . '<style>*{box-sizing:border-box;margin:0;padding:0}body{font-family:Geist,-apple-system,sans-serif;background:#fafafa;display:flex;align-items:center;justify-content:center;min-height:100vh}.card{background:#fff;border:1px solid #eaeaea;border-radius:12px;padding:40px 48px;text-align:center;max-width:360px;width:100%}h1{font-size:16px;font-weight:600;color:#000;margin-bottom:8px}p{font-size:13px;color:#888;line-height:1.6;margin-bottom:24px}a{display:inline-flex;align-items:center;gap:6px;height:36px;padding:0 18px;background:#000;color:#fff;border-radius:7px;font-size:13px;font-weight:500;text-decoration:none}a:hover{background:#333}</style></head><body>'
               . '<div class="card"><h1>Acceso restringido</h1><p>Necesitas estar logueado en el backoffice de PrestaShop para ver el visor de facturas.</p>'
               . '<a href="' . htmlspecialchars($loginUrl) . '">Ir al backoffice →</a></div></body></html>';
            exit;
        }
        $params = @include ROOT . '/app/config/parameters.php';
        if ($params) {
            $p = $params['parameters'];
            $pdo = new PDO(
                'mysql:host=' . $p['database_host'] . ';dbname=' . $p['database_name'] . ';charset=utf8',
                $p['database_user'], $p['database_password'], [PDO::ATTR_ERRMODE => PDO::ERRMODE_SILENT]
            );
            $stmt = $pdo->prepare("SELECT id_employee, firstname, lastname, email FROM `{$p['database_prefix']}employee` WHERE id_employee = ? AND active = 1 LIMIT 1");
            $stmt->execute([(int) $cookie->id_employee]);
            $emp = $stmt->fetch(PDO::FETCH_ASSOC);
            if ($emp) return ['employee' => $emp, 'pdo' => $pdo, 'prefix' => $p['database_prefix']];
        }
    } catch (Throwable $e) {
        ob_end_clean();
        denyAccess('No disponible', 'El sistema no pudo iniciarse. Comprueba que la BD esté activa.', 503);
    }
    denyAccess('No disponible', 'No se pudo verificar la sesión de empleado.', 503);
}

$auth = requirePsAdminAuth();
$pdo = $auth['pdo'];
$prefix = $auth['prefix'];

// Nunca cachear nada de este visor (ni el shell HTML ni el PDF): el
// contenido se re-renderiza en cada carga a partir de la BD.
header('Cache-Control: no-store, no-cache, must-revalidate, max-age=0');
header('Pragma: no-cache');
header('Expires: 0');

// Nombre de tienda leído de la config PS (sin hardcodear nada del proyecto).
$shopName = $pdo->query("SELECT value FROM `{$prefix}configuration` WHERE name = 'PS_SHOP_NAME' LIMIT 1")->fetchColumn() ?: 'PrestaShop';

// ── Bootstrap del contexto PS (container legacy 'front' para formateo de precios) ─
// PS8 no tiene AdminKernel; el contenedor legacy 'front' (el mismo que usa
// FrontController) expone prestashop.core.localization.locale.repository,
// que es lo que Tools/Locale necesitan para formatear precios en los .tpl del PDF.
// (En proyectos PS9 con AdminKernel disponible, se puede usar ese en su lugar.)

function bootPsContext(): void {
    global $auth;
    $context = Context::getContext();
    $context->container = PrestaShop\PrestaShop\Adapter\ContainerBuilder::getContainer('front', _PS_MODE_DEV_);
    $context->employee = new Employee((int) $auth['employee']['id_employee']);
    $context->currency = Currency::getDefaultCurrency();
    $context->country = new Country((int) Configuration::get('PS_COUNTRY_DEFAULT'));
    $context->language = new Language((int) Configuration::get('PS_LANG_DEFAULT'));
}

// ── Acción: generar factura si el pedido aún no tiene ──────────────────────────

if (isset($_POST['generate_invoice']) && ctype_digit($_POST['id_order'])) {
    bootPsContext();
    $order = new Order((int) $_POST['id_order']);
    if (Validate::isLoadedObject($order)) {
        $order->setInvoice(true);
    }
    header('Location: ' . $_SERVER['PHP_SELF']);
    exit;
}

// ── Acción: renderizar el PDF de un pedido ─────────────────────────────────────

if (isset($_GET['pdf']) && isset($_GET['id_order']) && ctype_digit($_GET['id_order'])) {
    bootPsContext();
    $order = new Order((int) $_GET['id_order']);
    if (!Validate::isLoadedObject($order) || !$order->invoice_number) {
        denyAccess('Sin factura', 'Este pedido no tiene número de factura generado todavía.', 404);
    }
    $invoices = $order->getInvoicesCollection();
    $pdf = new PDF($invoices, PDF::TEMPLATE_INVOICE, Context::getContext()->smarty);
    $content = $pdf->render(false);
    header('Content-Type: application/pdf');
    header('Content-Disposition: inline; filename="invoice-' . (int) $order->id . '.pdf"');
    header('Content-Length: ' . strlen($content));
    header('Cache-Control: no-store, no-cache, must-revalidate, max-age=0');
    header('Pragma: no-cache');
    echo $content;
    exit;
}

// ── Listado ─────────────────────────────────────────────────────────────────

$search = trim($_GET['q'] ?? '');
$where = '';
$params = [];
if ($search !== '') {
    $where = "WHERE o.reference LIKE ? OR c.firstname LIKE ? OR c.lastname LIKE ? OR o.id_order = ?";
    $like = '%' . $search . '%';
    $params = [$like, $like, $like, ctype_digit($search) ? (int) $search : 0];
}
$stmt = $pdo->prepare("
    SELECT o.id_order, o.reference, o.invoice_number, o.date_add, o.total_paid,
           c.firstname, c.lastname,
           osl.name AS state_name
    FROM `{$prefix}orders` o
    JOIN `{$prefix}customer` c ON c.id_customer = o.id_customer
    LEFT JOIN `{$prefix}order_state_lang` osl ON osl.id_order_state = o.current_state AND osl.id_lang = (SELECT id_lang FROM `{$prefix}lang` WHERE active = 1 LIMIT 1)
    $where
    ORDER BY o.id_order DESC
    LIMIT 50
");
$stmt->execute($params);
$orders = $stmt->fetchAll(PDO::FETCH_ASSOC);

?><!DOCTYPE html>
<html lang="es">
<head>
<meta charset="UTF-8">
<title>Invoice Preview — <?= htmlspecialchars($shopName) ?></title>
<link href="https://fonts.googleapis.com/css2?family=Geist:wght@400;500;600;700&display=swap" rel="stylesheet">
<style>
* { box-sizing: border-box; margin: 0; padding: 0; }
body { font-family: 'Geist', -apple-system, sans-serif; background: #fafafa; color: #111; }
.layout { display: flex; height: 100vh; }
.sidebar { width: 360px; flex-shrink: 0; background: #fff; border-right: 1px solid #eaeaea; display: flex; flex-direction: column; }
.sidebar-header { padding: 16px; border-bottom: 1px solid #eaeaea; }
.sidebar-header h1 { font-size: 14px; font-weight: 600; margin-bottom: 10px; }
.sidebar-header input { width: 100%; height: 34px; padding: 0 10px; border: 1px solid #ddd; border-radius: 7px; font-size: 13px; font-family: inherit; }
.order-list { flex: 1; overflow-y: auto; }
.order-row { display: block; padding: 12px 16px; border-bottom: 1px solid #f2f2f2; text-decoration: none; color: inherit; }
.order-row:hover { background: #f7f7f7; }
.order-row.active { background: #eef4ff; }
.order-row .ref { font-size: 13px; font-weight: 600; }
.order-row .meta { font-size: 11px; color: #888; margin-top: 2px; }
.order-row .state { display: inline-block; font-size: 10px; padding: 2px 6px; border-radius: 4px; background: #f0f0f0; color: #666; margin-top: 4px; }
.order-row form { margin-top: 6px; }
.btn-mini { font-size: 11px; padding: 4px 8px; border-radius: 5px; border: 1px solid #ddd; background: #fff; cursor: pointer; font-family: inherit; }
.btn-mini:hover { background: #f0f0f0; }
.viewer { flex: 1; display: flex; flex-direction: column; }
.viewer-empty { flex: 1; display: flex; align-items: center; justify-content: center; color: #999; font-size: 13px; }
.viewer iframe { flex: 1; border: none; }
</style>
</head>
<body>
<div class="layout">
	<div class="sidebar">
		<div class="sidebar-header">
			<h1>Facturas — <?= htmlspecialchars($shopName) ?></h1>
			<form method="get"><input type="text" name="q" placeholder="Buscar por ref, cliente, id..." value="<?= htmlspecialchars($search) ?>"></form>
		</div>
		<div class="order-list">
			<?php foreach ($orders as $o): ?>
				<div class="order-row <?= (isset($_GET['id_order']) && (int) $_GET['id_order'] === (int) $o['id_order']) ? 'active' : '' ?>">
					<div class="ref"><?= htmlspecialchars($o['reference']) ?> — #<?= (int) $o['id_order'] ?></div>
					<div class="meta"><?= htmlspecialchars($o['firstname'] . ' ' . $o['lastname']) ?> · <?= htmlspecialchars($o['date_add']) ?></div>
					<span class="state"><?= htmlspecialchars($o['state_name'] ?? '—') ?></span>
					<?php if ($o['invoice_number']): ?>
						<div style="margin-top:6px;"><a class="btn-mini" href="?id_order=<?= (int) $o['id_order'] ?><?= $search !== '' ? '&q=' . urlencode($search) : '' ?>" style="text-decoration:none;display:inline-block;">Ver factura #<?= (int) $o['invoice_number'] ?></a></div>
					<?php else: ?>
						<form method="post">
							<input type="hidden" name="id_order" value="<?= (int) $o['id_order'] ?>">
							<button class="btn-mini" type="submit" name="generate_invoice" value="1">Generar factura</button>
						</form>
					<?php endif; ?>
				</div>
			<?php endforeach; ?>
			<?php if (!$orders): ?>
				<div style="padding:16px;font-size:13px;color:#999;">Sin resultados.</div>
			<?php endif; ?>
		</div>
	</div>
	<div class="viewer">
		<?php if (isset($_GET['id_order']) && ctype_digit($_GET['id_order'])): ?>
			<iframe src="?id_order=<?= (int) $_GET['id_order'] ?>&pdf=1&_=<?= time() ?>"></iframe>
		<?php else: ?>
			<div class="viewer-empty">Elegí un pedido con factura generada para ver el PDF.</div>
		<?php endif; ?>
	</div>
</div>
</body>
</html>
