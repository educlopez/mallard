---
description: Ask the prestashop-expert agent — PS 8/9 core (Symfony BO, Twig, Smarty, hooks, modules, themes, migration).
argument-hint: <your question for the PrestaShop core expert>
version: "0.1.3"
metadata:
  author: Eduardo Calvo
---

Invoke the `prestashop-experts:prestashop-expert` subagent to answer the following question. Forward the question verbatim and let the agent apply its CWD-first protocol (it will detect the project shape — `install-dev/` vs `install/`, missing `vendor/`, themes folder layout — and adjust its response accordingly).

If `$ARGUMENTS` is empty, first ask the user what they need from the PrestaShop core expert (e.g. *"¿Sobre qué necesitas al prestashop-expert? Hook, módulo, theme parent-child, Smarty, migración 8→9…?"*) before invoking the agent.

User's question:

$ARGUMENTS
