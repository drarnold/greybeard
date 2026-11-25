---
name: adr-keeper
description: Record architecture decisions as short, durable ADRs capturing context, options, and consequences. Use when a significant technical decision is made, reversed, or when someone asks why the system is built this way.
---

# ADR Keeper

Future engineers will ask "why on earth did we do this?" Leave them an answer.

## When to Write One
A decision is ADR-worthy if it is expensive to reverse, affects multiple teams, or someone later will plausibly ask why. Choosing a database, a messaging pattern, an auth approach, or a build tool all qualify. Variable naming does not.

## Format (keep it to one page)
- **Title:** Short, numbered: `ADR-0007: Use Postgres for the ledger`.
- **Status:** Proposed, Accepted, Superseded by ADR-00XX, or Deprecated.
- **Context:** The forces at play, including constraints and what we knew at the time.
- **Options considered:** Two or three real alternatives with honest pros and cons.
- **Decision:** What we chose, in one or two sentences.
- **Consequences:** Good, bad, and the follow-up work. Admit the costs.

## Habits
- Store in the repo (`docs/adr/`), reviewed like code.
- ADRs are append-only. Never rewrite history. Supersede with a new ADR.
- Write it when deciding, not three months later from fuzzy memory.
- Link to related ADRs, RFCs (see `architecture-rfc`), and spikes (see `spike-and-prototype`).

## Guardrails
- Capture the options you rejected, or the next person will reinvent them.
- No decision theater. If it was already decided, record that honestly.

## Output Format
A ready-to-commit Markdown ADR using the format above, with a filename suggestion.
