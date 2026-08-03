# PrestaShop — ci-templates wiring

## Components to include

```yaml
workflow:
  rules:
    - if: '$CI_PIPELINE_SOURCE == "merge_request_event"'
    - if: '$CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH'

include:
  - component: gitlab.com/educlopezcinetic/ci-templates/semantic-mr-title@v1.2.0
  - component: gitlab.com/educlopezcinetic/ci-templates/agent-guardrail@v1.2.0
  - component: gitlab.com/educlopezcinetic/ci-templates/secret-scan@v1.2.0
  - component: gitlab.com/educlopezcinetic/ci-templates/dependency-audit-php@v1.2.0
  # Only if the theme has its own package.json (PrestaShop theme JS/CSS build):
  - component: gitlab.com/educlopezcinetic/ci-templates/dependency-audit-node@v1.2.0
    inputs:
      package-manager: pnpm   # pnpm | npm | yarn — check the theme's lockfile
      working-directory: themes/<theme-name>
  - component: gitlab.com/educlopezcinetic/ci-templates/lint-node@v1.2.0
    inputs:
      package-manager: pnpm
      lint-command: 'pnpm run lint'
      working-directory: themes/<theme-name>
  - component: gitlab.com/educlopezcinetic/ci-templates/test-php@v1.2.0
    inputs:
      setup-command: 'composer install --no-interaction --no-progress'
      test-command: 'SYMFONY_DEPRECATIONS_HELPER=disabled php -d date.timezone=UTC vendor/phpunit/phpunit/phpunit -c tests/Unit/phpunit.xml'
```

`typecheck-node` is almost never relevant here — PrestaShop themes are plain
JS, not TS. Skip it unless the theme genuinely has a TS build.

If the project has no `package.json`/theme JS at all, skip
`lint-node`/`dependency-audit-node` entirely — don't add dead includes.

## Static analysis — PHPStan (repo-local, not a component)

Follows the official PrestaShop devdocs CI/CD pattern
(`devdocs.prestashop-project.org/9/modules/testing/ci-cd/`). Start
**non-blocking** (`allow_failure: true`) — a fresh PrestaShop core checkout
already has real upstream PHPStan findings a scoped `phpstan-baseline.neon`
subset won't cover, so gating merges on it from day one blocks everyone on
pre-existing noise, not anything the current MR introduced.

```yaml
phpstan:
  stage: test
  needs: []   # do NOT omit — see SKILL.md's needs:[] gotcha
  image: php:8.3-cli
  before_script:
    - apt-get update -qq && apt-get install -y -qq --no-install-recommends git unzip libicu-dev libzip-dev > /dev/null
    - docker-php-ext-install intl zip > /dev/null
    - curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer
  script:
    - composer install --no-interaction --no-progress
    - php -d memory_limit=-1 vendor/bin/phpstan analyze -c phpstan-scoped.neon --no-progress
  allow_failure: true
  rules:
    - if: '$CI_PIPELINE_SOURCE == "merge_request_event"'
    - if: '$CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH'
```

`phpstan-scoped.neon` should scope analysis to the project's own
module(s)/theme code, not the full PrestaShop core tree — core will always
have pre-existing findings that aren't this project's problem to fix.

## What to expect from `dependency-audit-php` on a fresh install

Don't be alarmed if it finds real CVEs immediately on a brand-new, untouched
PrestaShop checkout — normal, not a setup mistake. PrestaShop's
`composer.lock` pins third-party packages (`phpoffice/phpspreadsheet`,
`symfony/ux-icons`, `api-platform/core` are the common offenders) to
whatever was current at release time; new CVEs against those exact pinned
versions surface constantly and PrestaShop doesn't retroactively patch old
releases. Standard for any framework bundling third-party deps, not
PrestaShop-specific negligence.

**Before hand-patching any finding**, compare the installed version
(`composer.lock`) and the project's own constraint (`composer.json`) against
the advisory's fixed-version range:

- Fix is a **patch/minor bump within the existing constraint** (e.g.
  `phpoffice/phpspreadsheet: ^1.19` installed at `1.30.4`, fix at `1.30.6` —
  still `<2.0.0`) → safe, low-risk, `composer update <package>` alone.
- Fix requires **crossing a major version boundary outside the constraint**
  (e.g. `api-platform/core: ^3.4` but the fix needs `4.1.29+`) → do NOT
  hand-patch. Breaking-change jump the core code wasn't tested against; wait
  for an official PrestaShop release that bumps it, or treat it as an
  accepted risk in the interim and say so explicitly rather than silently
  leaving it or silently forcing a risky bump.
