---
description: Ask the panda-expert agent — Panda theme + Easy Builder + st* modules + SunnyToo demos.
argument-hint: <your question for the Panda expert>
version: "0.1.3"
metadata:
  author: Eduardo Calvo
---

Invoke the `prestashop-experts:panda-expert` subagent to answer the following question. Forward the question verbatim and let the agent apply its CWD-first protocol (it will detect whether the current working directory is a PrestaShop project with Panda installed, and adjust its response accordingly).

If `$ARGUMENTS` is empty, first ask the user what they need from the Panda expert (e.g. *"¿Sobre qué necesitas al panda-expert? Módulos, layout, Easy Builder, configuración de un st*…?"*) before invoking the agent.

User's question:

$ARGUMENTS
