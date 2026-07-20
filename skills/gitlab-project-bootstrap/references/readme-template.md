Copy-pasteable README skeleton. Fill in placeholders, delete sections that don't apply (e.g. no Local Setup section if there's no real local dev flow to document).

```markdown
Project Name
============

[![Stack](https://img.shields.io/badge/Stack-Version-COLOR?style=flat-square)](https://link-to-stack-docs)
[![PHP](https://img.shields.io/badge/PHP-%3E%3D%208.1-8892BF?style=flat-square)](https://php.net/)
[![Theme](https://img.shields.io/badge/theme-Name-COLOR?style=flat-square)](themes/theme-dir)

One-line description of what this project actually is (audience: a new dev or a client stakeholder skimming GitLab).

<p align="center">
  <img src="docs/screenshot-home.webp" alt="Project homepage" width="900"/>
</p>

Stack
-----

- **Framework/Platform + version**
- **Theme/frontend**, if applicable — name the parent theme explicitly
- Anything architecturally notable: fork-tracking strategy (upstream remote + edition branches), override system, custom modules, page-builder module, monorepo/workspace structure, etc.

Environments
------------

- **Production:** https://real-url.example.com/
- **Pre/Staging:** https://pre-url.example.com/ (note "currently in maintenance" etc. if applicable, don't just omit)

Local setup
-----------

Only include if there's a real, documented local dev flow (Lando, Docker, etc.) — don't invent one that doesn't exist.

```bash
lando start
```

Theme/asset build
------------------

Only include if the project has its own build pipeline (e.g. a CSS/JS bundler for a theme).

```bash
cd path/to/theme
pnpm install
pnpm run build
```

Conventions
-----------

See [CLAUDE.md](CLAUDE.md) for full architecture notes, conventions, and workflow. Don't duplicate its content here — just point to it.
```
