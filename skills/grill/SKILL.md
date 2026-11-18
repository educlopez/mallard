---
name: grill
description: >
  Stress-test a plan, estimate or design with a decision-tree interview before
  any work starts. Use when the user says "hazme preguntas", "grill", "revisemos
  el plan", "antes de empezar", "estima esta tarea", or when `/task` returns a
  thin or ambiguous brief.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

# grill

Adapted from Matt Pocock's `grilling` and `grill-me`. It interviews, it does not build: no file edits, no Intervals writes, nothing sent to a client. For questions only the client can answer, hand off to `client-questionnaire`.

## 1. State the decision

Read the plan, the brief (engram `task/<localid>/brief`) or what the user described. Done when you can say the decision in one sentence and the user does not correct it.

## 2. Map the open decisions

List every decision the plan depends on as a tree. The frontier is the set of decisions whose prerequisites are already settled. Ask only about the frontier.

## 3. Look up facts yourself

Anything the repo or environment answers is not a question for the user: `composer.json`, `AGENTS.md`, existing tests, the module or route involved, the framework version. For wide lookups use a read-only subagent when one is available.

## 4. Ask the frontier

One round = the whole frontier, at most five questions. Per question:

```
Qn - <title>
<what is open and why it matters, two lines at most>
Recommended: <your answer>
```

Word each question so that "sí" accepts your recommendation.

## 5. Repeat until the frontier is empty

Each answer can open or close decisions: redraw the tree and ask the new frontier. Done when no decision is open.

## 6. Close

Print decisions, assumptions, and the questions left for the client or PM. If an estimate was asked, give a range in hours and the assumption that moves it most. Then wait for the user's OK.

Next: `research-brief` for a fact you could not find, `client-questionnaire` for client-only questions, then the route that `/task` named.

It does not start the work, log time or touch Intervals.
