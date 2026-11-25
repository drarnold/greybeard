---
name: tech-debt-triage
description: Inventory, prioritize, and make the business case for technical debt. Use when the backlog is full of "we should clean that up," when planning debt paydown, or when justifying refactoring time.
---

# Tech Debt Triage

Debt is a loan. Some is smart, some is predatory. Know the interest rate before choosing what to pay down.

## Classify
- **Deliberate and prudent:** a shortcut taken on purpose, with a plan. Fine.
- **Accidental:** we learned better since. Normal.
- **Reckless:** nobody cared. Fix the process too.

## Prioritize by Interest, Not Ugliness
For each item, estimate:
1. **Interest:** how much does it slow the team or cause incidents each month?
2. **Principal:** what does it cost to fix?
3. **Blast radius:** how many teams or customers does it touch?
4. **Trajectory:** is it getting worse, or stable?

High interest and low principal goes first. Ugly but stable code that nobody touches can stay ugly.

## Pay It Down
- Tie work to features: leave the area better than you found it (boy scout rule).
- Reserve a steady capacity (e.g. 15-20%) rather than begging for a "cleanup sprint."
- Make it visible: a debt register with owner, interest, and a target date.

## Guardrails
- "Rewrite" is rarely the answer. Prefer incremental paths (see `legacy-code-archaeology`).
- Speak the business's language: lead time, incident cost, and delivery risk, not "code smells."

## Output Format
A ranked **debt register**: item, interest, principal, recommendation, and the business case in one sentence each.
