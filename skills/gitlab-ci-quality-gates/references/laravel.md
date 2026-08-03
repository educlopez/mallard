# Laravel — ci-templates wiring

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
  - component: gitlab.com/educlopezcinetic/ci-templates/dependency-audit-node@v1.2.0
    inputs:
      package-manager: pnpm   # pnpm | npm | yarn — check the lockfile
  - component: gitlab.com/educlopezcinetic/ci-templates/lint-node@v1.2.0
    inputs:
      lint-command: 'pnpm run lint'
  # Only if there's a real TS frontend (Inertia+React+TS, etc) — a plain
  # Blade/Livewire app has nothing to typecheck, skip this include entirely:
  - component: gitlab.com/educlopezcinetic/ci-templates/typecheck-node@v1.2.0
  - component: gitlab.com/educlopezcinetic/ci-templates/test-php@v1.2.0
    inputs:
      setup-command: 'composer install --no-interaction --prefer-dist --no-progress; cp .env.example .env; php artisan key:generate'
      test-command: 'php artisan test'
      # php-image defaults to php:8.3-cli — check composer.lock BEFORE trusting the
      # default. A fresh Laravel 12/13 skeleton's symfony/* v8.x transitive deps
      # already require PHP >=8.4.1; if the lockfile has them, `composer install`
      # fails with "lock file does not contain a compatible set of packages" on
      # 8.3, unrelated to anything the MR actually changed. Confirmed as a real
      # bug in laravel-demo before this skill existed — set explicitly:
      php-image: 'php:8.4-cli'
```

## Static analysis — Larastan (repo-local, not a component)

Match `php-image` to whatever you set on `test-php` above — same PHP-version
mismatch risk applies here (composer install runs in this job too).

```yaml
larastan:
  stage: test
  needs: []   # do NOT omit — see SKILL.md's needs:[] gotcha
  image: php:8.4-cli   # match test-php's php-image — see note above
  before_script:
    - curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer
  script:
    - composer install --no-interaction --no-progress
    - cp .env.example .env
    - php artisan key:generate
    - vendor/bin/phpstan analyze --memory-limit=-1
  allow_failure: true
  rules:
    - if: '$CI_PIPELINE_SOURCE == "merge_request_event"'
    - if: '$CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH'
```

If `larastan/larastan` isn't in `composer.json` yet and you can't run
`composer require --dev` locally to regenerate the lockfile properly (e.g.
no SSH access to the account hosting the repo), install it fresh at job
runtime instead of assuming it's pinned — slower per-run, but self-contained
and doesn't require touching the committed lockfile:

```yaml
    - composer install --no-interaction --no-progress
    - composer require --dev larastan/larastan --no-interaction --no-progress
```

Also needs a `phpstan.neon`/`phpstan.neon.dist` at the repo root:

```neon
includes:
    - vendor/larastan/larastan/extension.neon

parameters:
    paths:
        - app
    level: 0
```

Requires `larastan/larastan` in `composer.json` (`require-dev`) and a
`phpstan.neon`/`phpstan.neon.dist` with
`includes: [vendor/larastan/larastan/extension.neon]`. If the project has
neither, add them first — ask before setting `level` above `1` on an
existing app; a higher level on a codebase that's never run static analysis
usually surfaces a wall of pre-existing findings unrelated to the current
MR. Start low (`level: 0` or `1`), keep it `allow_failure: true`, and raise
the level deliberately over time once the backlog is triaged — same
reasoning as PrestaShop's scoped-baseline approach.

## Dependency audit findings

No PrestaShop-specific "known bundled CVEs on fresh install" pattern here —
Laravel's own `composer.lock`/`package-lock.json` tend to start cleaner.
Apply the same general rule as PrestaShop's reference file though: compare
the installed version against the project's own version constraint before
hand-patching anything `dependency-audit-php`/`dependency-audit-node` finds.
A patch/minor bump within the existing constraint is low-risk; a fix that
needs crossing a major version boundary is not something to hand-patch
without testing — flag it and ask instead.
