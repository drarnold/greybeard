---
name: spike-and-prototype
description: Run timeboxed spikes and throwaway prototypes that answer one specific question without turning into production code. Use when facing a technical unknown, evaluating a library, or exploring a design before committing.
---

# Spike and Prototype

A spike buys information, not software. If it ships, you did the wrong thing.

## Before You Write a Line
1. **State the question.** One sentence: "Can library X handle 10k events/sec?" A spike without a question is just coding with extra steps.
2. **Set the timebox.** Hours or a day or two, written down. When it expires, report what you learned, even if the answer is "unclear."
3. **Define the exit.** What result means go, no-go, or needs more data?

## While Spiking
- Hardcode, skip auth, ignore edge cases. Speed is the point.
- Keep notes as you go: surprises, dead ends, numbers.
- Work on a throwaway branch. Name it `spike/...` so nobody merges it by accident.

## When It Ends
- Write up the **finding**, not the code: question, answer, evidence, recommendation.
- Delete the code, or label it clearly throwaway. If the answer is "build it," rebuild it properly with tests. Do not promote the prototype.
- Record the decision (see `adr-keeper`).

## Guardrails
- Prototypes that "just need a little cleanup" become production. This is how legacy is born.
- Do not spike what you can answer by reading the docs for 20 minutes.

## Output Format
**Question**, **Timebox used**, **Finding** (with numbers or evidence), **Recommendation** (go / no-go / more data), **What to do with the code** (delete, archive, or reference).
