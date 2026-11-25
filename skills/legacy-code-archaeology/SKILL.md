---
name: legacy-code-archaeology
description: Safely understand, document, and modify old undocumented code. Use when inheriting a system with no tests, no owners, and no docs, or when asked to change code nobody dares touch.
---

# Legacy Code Archaeology

Legacy code is code without tests, and also code that pays the bills. Respect both facts.

## Dig, Don't Demolish
1. **Map before you move.** Entry points, data flow, side effects, external calls, and what writes to the database. Draw it.
2. **Read history.** `git log` and `git blame` explain why odd code exists. Search old tickets and commit messages before calling something stupid. It may be load-bearing.
3. **Find the seam.** A place where behavior can be changed without editing the surrounding code (an interface, a parameter, a wrapper).
4. **Pin behavior with characterization tests:** feed inputs, record outputs, assert they stay the same.
5. **Change minimally,** then clean up in separate, tested steps (see `refactor-safely`).

## Strategies
- **Strangler fig:** route new behavior through new code, retire old paths one at a time.
- **Sprout method/class:** add new logic in a fresh, tested unit and call it from the old code.
- **Wrap and observe:** add logging and metrics first to learn what actually runs in production. Dead code is common.

## Guardrails
- Do not rewrite from scratch. The rewrite will rediscover every bug the old code fixed.
- Chesterton's fence: do not remove it until you know why it was built.
- Record what you learn as you go (see `onboarding-docs`).

## Output Format
**System map**, **Risky areas**, **Suggested seams**, **First tests to write**, and a **safe change plan** with rollback.
