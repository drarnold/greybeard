---
name: requirements-grilling
description: Stress-test a vague ticket, user story, or feature request before anyone writes code. Use when requirements are ambiguous, a spec feels too easy, or work is about to start without acceptance criteria.
---

# Requirements Grilling

A ticket that says "make it faster" or "add export" is a question wearing a trench coat. Interrogate it before the sprint commits to it.

## Ask, in this order
1. **Why:** What problem does this solve, for whom, and how will we know it worked? No measurable outcome means no done.
2. **Who:** Which user, which role, which permission level? "Users" is not a persona.
3. **Edges:** Empty state, max size, concurrent edits, partial failure, timezone, locale, and the user who does the weird thing on purpose.
4. **Non-goals:** What are we explicitly *not* building? Write it down so scope creep has to argue with a document.
5. **Constraints:** Deadline, compliance, backwards compatibility, budget, who gets paged.
6. **Cheaper path:** Is there a config flag, existing feature, or manual process that solves 80% of this?

## Guardrails
- Ask the 3 questions that would change the design most. Do not carpet-bomb with 40.
- Separate **facts** (stated in the ticket) from **assumptions** (yours). Label every assumption.
- Never invent requirements to fill silence. Flag the silence.

## Output Format
- **Restated problem:** One sentence, in your own words.
- **Open questions:** Ranked by how much the answer changes the build.
- **Assumptions:** Each one marked `[ASSUMED]` for the stakeholder to confirm or reject.
- **Draft acceptance criteria:** Testable, Given/When/Then where it helps.
- **Non-goals:** Explicit list.
