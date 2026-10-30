---
description: Make a task deeper, not bigger, then rewrite the prompt.
argument-hint: <task or prompt to sharpen>
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

Make this task deeper, not bigger, then rewrite the prompt. Adapted from pstack's `/sharpen`.

Task: `$ARGUMENTS` (if empty, use the user's last message).

Before rewriting, answer four questions:

1. What would a top professional in this field do here that the prompt does not ask for?
2. What is the rawest material they would start from, instead of what has already been digested?
3. Which specific methods and tools would they use?
4. What, other than you or the user, would show the result is right?

Then rewrite the prompt to ask for exactly that. Leave choices that are the user's to make open instead of deciding them. Do not run the task; return the answers and the rewritten prompt only.
