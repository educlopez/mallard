---
description: Catch me up on what changed in this chat and what is waiting on me.
version: "0.1.0"
metadata:
  author: Eduardo Calvo
---

Catch the user up on what changed in this chat since their last message before `/ketchup`. Adapted from pstack's `/ketchup`, text only. Read-only: change no code, start no app or job, post nothing.

Reply in the user's language, in plain words and short sentences. Say what changed for the user, not how you did it. No hashes, file paths or command names unless the user needs them to act.

1. **Open** with one sentence naming what waits on the user, or saying nothing does. If nothing happened and nothing waits, say so and stop.
2. **What happened**: under 100 words, one bullet per thing the user asked for, with its state: done, running, failed, or blocked. Leave out work that changed nothing for them.
3. **What waits on you**: every call only the user can make (a choice, an approval, an irreversible action, access you lack), including ones raised before their last message. For each: the call, your pick, what it changes for users, the evidence behind your pick, and its risk. Write "unknown" for a fact you lack. Work you can still do yourself is not an action item.
