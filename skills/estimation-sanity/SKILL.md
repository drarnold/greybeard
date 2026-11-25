---
name: estimation-sanity
description: Produce honest estimates as ranges with named risks and assumptions, and push back on single-number deadlines. Use when asked how long something will take, when sizing a roadmap, or when a date was promised before the work was understood.
---

# Estimation Sanity

An estimate is a probability distribution, not a promise. Anyone who gives you one number is guessing or selling.

## Method
1. **Decompose** until each piece is understandable. If you cannot describe it, you cannot size it.
2. **Estimate each piece as a range:** best case, likely, and "everything went sideways."
3. **Add the invisible work:** code review, testing, deploy, docs, meetings, on-call interruptions, and the thing nobody listed.
4. **Name the unknowns.** Each unknown is either a spike (timeboxed) or a risk line with a multiplier.
5. **Compare to history.** Similar past work beats gut feel. Say what you compared it to.

## Guardrails
- Never convert an estimate into a commitment without the stakeholder seeing the range.
- Effort is not duration. Convert using realistic focus time (usually 4-5 productive hours a day, not 8).
- "Add people" does not shorten a late project. Say so politely.
- Cut scope before cutting quality. Offer the scope options explicitly.

## Output Format
- **Estimate:** `X–Y days/weeks`, with a confidence level (e.g. 80%).
- **Assumptions:** What must be true for this range to hold.
- **Top risks:** Each with its impact on the range.
- **Unknowns to spike first:** With timebox.
- **Scope levers:** What to cut or defer if the date is fixed.
