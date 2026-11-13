{**
 * Env Favicon - back office page.
 *
 * @author    Eduardo Calvo
 * @copyright Eduardo Calvo
 * @license   https://opensource.org/licenses/AFL-3.0 Academic Free License 3.0
 *}

<div class="ef-wrap" data-current="{$ef_env|escape:'html':'UTF-8'}">

  <header class="ef-hero">
    <img class="ef-hero__logo" src="{$ef_module_logo|escape:'html':'UTF-8'}" width="64" height="64" alt="">
    <div class="ef-hero__text">
      <h1>{l s='Env Favicon' mod='envfavicon'}</h1>
      <p>{l s='One logo, a different favicon on every environment. Production always keeps the native PrestaShop favicon.' mod='envfavicon'}</p>
    </div>
    <div class="ef-here" title="{l s='How this host is detected right now' mod='envfavicon'}">
      <span class="ef-here__label">{l s='This host' mod='envfavicon'}</span>
      <code>{$ef_host|escape:'html':'UTF-8'}</code>
      <span class="ef-chip ef-chip--{$ef_env|escape:'html':'UTF-8'}">{if $ef_env == 'local'}{l s='Local' mod='envfavicon'}{elseif $ef_env == 'pre'}{l s='Staging' mod='envfavicon'}{else}{l s='Production' mod='envfavicon'}{/if}</span>
    </div>
  </header>

  <form class="ef-form" method="post" action="{$ef_action|escape:'html':'UTF-8'}" enctype="multipart/form-data">

    {* 1 · Logo *}
    <section class="ef-card">
      <div class="ef-card__head">
        <span class="ef-step">1</span>
        <div>
          <h2>{l s='Logo' mod='envfavicon'}</h2>
          <p>{l s='Upload it once: every badge and the production .ico are generated from it.' mod='envfavicon'}</p>
        </div>
      </div>

      <div class="ef-logo">
        <label class="ef-drop" for="ef-file" id="ef-drop">
          <input class="ef-sr" type="file" id="ef-file" name="ENVFAV_SOURCE" accept="image/png,image/jpeg,image/webp">
          <svg class="ef-drop__icon" viewBox="0 0 24 24" width="30" height="30" aria-hidden="true"><path fill="none" stroke="currentColor" stroke-width="1.7" stroke-linecap="round" stroke-linejoin="round" d="M12 16V5m0 0-4 4m4-4 4 4M5 15v3a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-3"/></svg>
          <strong>{l s='Drop your logo here' mod='envfavicon'}</strong>
          <span>{l s='or click to choose a file' mod='envfavicon'}</span>
          <small>{l s='PNG, JPG or WEBP, up to 3 MB. A transparent background (or a dark mark on white) works best.' mod='envfavicon'}</small>
          <span class="ef-drop__file" id="ef-file-name" hidden></span>
        </label>

        <div class="ef-mark">
          <div class="ef-mark__stage" id="ef-mark-stage">
            {if $ef_has_mark}
              <img src="{$ef_mark_url|escape:'html':'UTF-8'}" alt="{l s='Detected mark' mod='envfavicon'}" id="ef-mark-img">
            {else}
              <span class="ef-mark__empty" id="ef-mark-empty">{l s='No logo yet' mod='envfavicon'}</span>
            {/if}
          </div>
          <div class="ef-mark__raw" id="ef-raw" hidden><img alt="" id="ef-raw-img"></div>
          <p class="ef-mark__cap">{if $ef_has_mark}{l s='Detected mark (white on dark)' mod='envfavicon'}{else}{l s='The detected mark will appear here' mod='envfavicon'}{/if}</p>
        </div>
      </div>

      <label class="ef-switch">
        <input type="checkbox" name="ENVFAV_CROP_MARK" value="1"{if $ef_v.ENVFAV_CROP_MARK} checked{/if}>
        <span class="ef-switch__track" aria-hidden="true"></span>
        <span class="ef-switch__text">
          <strong>{l s='Crop only the mark' mod='envfavicon'}</strong>
          <small>{l s='For "mark + text" logos: keeps the first block on the left and removes the (R) sign and specks.' mod='envfavicon'}</small>
        </span>
      </label>
    </section>

    {* 2 · Environments *}
    <section class="ef-card">
      <div class="ef-card__head">
        <span class="ef-step">2</span>
        <div>
          <h2>{l s='Environments' mod='envfavicon'}</h2>
          <p>{l s='A host that contains any of the fragments belongs to that environment. Anything that matches nothing is production.' mod='envfavicon'}</p>
        </div>
      </div>

      <div class="ef-envs">

        <article class="ef-env{if $ef_env == 'local'} is-current{/if}" data-env="local" style="--c: {$ef_v.ENVFAV_COLOR_LOCAL|escape:'html':'UTF-8'}">
          <div class="ef-env__top">
            <h3>{l s='Local' mod='envfavicon'}</h3>
            {if $ef_env == 'local'}<span class="ef-pill">{l s='You are here' mod='envfavicon'}</span>{/if}
          </div>
          <div class="ef-tabbar"><div class="ef-tab"><span class="ef-badge">{if $ef_has_mark}<img src="{$ef_mark_url|escape:'html':'UTF-8'}" alt="">{/if}</span><span class="ef-tab__title">{$ef_shop_name|escape:'html':'UTF-8'}</span><span class="ef-tab__x" aria-hidden="true">&times;</span></div></div>
          <label class="ef-label" for="ef-local">{l s='Domains' mod='envfavicon'}</label>
          <textarea class="ef-input" id="ef-local" name="ENVFAV_LOCAL_PATTERNS" rows="3" spellcheck="false">{$ef_v.ENVFAV_LOCAL_PATTERNS|escape:'html':'UTF-8'}</textarea>
          <p class="ef-hint">{l s='Comma separated, e.g. lndo.site, localhost, .test' mod='envfavicon'}</p>
          <label class="ef-label" for="ef-c-local">{l s='Badge colour' mod='envfavicon'}</label>
          <div class="ef-color"><input type="color" id="ef-c-local" name="ENVFAV_COLOR_LOCAL" value="{$ef_v.ENVFAV_COLOR_LOCAL|escape:'html':'UTF-8'}"><input class="ef-hex" type="text" value="{$ef_v.ENVFAV_COLOR_LOCAL|escape:'html':'UTF-8'}" maxlength="7" spellcheck="false" aria-label="{l s='Hex colour' mod='envfavicon'}"></div>
        </article>

        <article class="ef-env{if $ef_env == 'pre'} is-current{/if}" data-env="pre" style="--c: {$ef_v.ENVFAV_COLOR_PRE|escape:'html':'UTF-8'}">
          <div class="ef-env__top">
            <h3>{l s='Staging' mod='envfavicon'}</h3>
            {if $ef_env == 'pre'}<span class="ef-pill">{l s='You are here' mod='envfavicon'}</span>{/if}
          </div>
          <div class="ef-tabbar"><div class="ef-tab"><span class="ef-badge">{if $ef_has_mark}<img src="{$ef_mark_url|escape:'html':'UTF-8'}" alt="">{/if}</span><span class="ef-tab__title">{$ef_shop_name|escape:'html':'UTF-8'}</span><span class="ef-tab__x" aria-hidden="true">&times;</span></div></div>
          <label class="ef-label" for="ef-pre">{l s='Domains' mod='envfavicon'}</label>
          <textarea class="ef-input" id="ef-pre" name="ENVFAV_PRE_PATTERNS" rows="3" spellcheck="false">{$ef_v.ENVFAV_PRE_PATTERNS|escape:'html':'UTF-8'}</textarea>
          <p class="ef-hint">{l s='Comma separated, e.g. pre-testing.com, ecommercloud' mod='envfavicon'}</p>
          <label class="ef-label" for="ef-c-pre">{l s='Badge colour' mod='envfavicon'}</label>
          <div class="ef-color"><input type="color" id="ef-c-pre" name="ENVFAV_COLOR_PRE" value="{$ef_v.ENVFAV_COLOR_PRE|escape:'html':'UTF-8'}"><input class="ef-hex" type="text" value="{$ef_v.ENVFAV_COLOR_PRE|escape:'html':'UTF-8'}" maxlength="7" spellcheck="false" aria-label="{l s='Hex colour' mod='envfavicon'}"></div>
        </article>

        <article class="ef-env{if $ef_env == 'prod'} is-current{/if}" data-env="prod" style="--c: {$ef_v.ENVFAV_COLOR_PROD|escape:'html':'UTF-8'}">
          <div class="ef-env__top">
            <h3>{l s='Production' mod='envfavicon'}</h3>
            {if $ef_env == 'prod'}<span class="ef-pill">{l s='You are here' mod='envfavicon'}</span>{/if}
          </div>
          <div class="ef-tabbar"><div class="ef-tab"><span class="ef-badge">{if $ef_has_mark}<img src="{$ef_mark_url|escape:'html':'UTF-8'}" alt="">{/if}</span><span class="ef-tab__title">{$ef_shop_name|escape:'html':'UTF-8'}</span><span class="ef-tab__x" aria-hidden="true">&times;</span></div></div>
          <div class="ef-note">
            <strong>{l s='Everything else' mod='envfavicon'}</strong>
            <span>{l s='Production uses the native PrestaShop favicon (Design > Theme & Logo). Load the generated one below, or upload your own there.' mod='envfavicon'}</span>
          </div>
          <label class="ef-label" for="ef-c-prod">{l s='Favicon colour' mod='envfavicon'}</label>
          <div class="ef-color"><input type="color" id="ef-c-prod" name="ENVFAV_COLOR_PROD" value="{$ef_v.ENVFAV_COLOR_PROD|escape:'html':'UTF-8'}"><input class="ef-hex" type="text" value="{$ef_v.ENVFAV_COLOR_PROD|escape:'html':'UTF-8'}" maxlength="7" spellcheck="false" aria-label="{l s='Hex colour' mod='envfavicon'}"></div>
        </article>

      </div>

      <label class="ef-switch">
        <input type="checkbox" name="ENVFAV_ENABLED" value="1"{if $ef_v.ENVFAV_ENABLED} checked{/if}>
        <span class="ef-switch__track" aria-hidden="true"></span>
        <span class="ef-switch__text">
          <strong>{l s='Module active' mod='envfavicon'}</strong>
          <small>{l s='Off: every environment shows the native PrestaShop favicon.' mod='envfavicon'}</small>
        </span>
      </label>
    </section>

    <div class="ef-actions">
      <button type="submit" name="submitEnvFavicon" class="ef-btn ef-btn--primary">{l s='Save and regenerate' mod='envfavicon'}</button>
      <button type="submit" name="submitEnvFaviconNative" class="ef-btn ef-btn--ghost"{if !$ef_has_ico} disabled title="{l s='Upload a logo and save first' mod='envfavicon'}"{/if}>{l s='Load the production favicon as the PrestaShop favicon' mod='envfavicon'}</button>
    </div>
  </form>

</div>
