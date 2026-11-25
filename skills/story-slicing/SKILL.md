---
name: story-slicing
description: Break large epics and stories into thin vertical slices that each deliver user-visible value. Use when a story is too big to finish in a sprint, when work is split by technical layer, or when planning an iterative rollout.
---

# Story Slicing

Big stories hide risk. Slice until each piece ships on its own and teaches you something.

## Rules of the Knife
1. **Slice vertically.** Each slice touches UI, logic, and storage just enough to work end to end. "Build the database layer" is not a story. It is a delayed integration failure.
2. **Each slice is demoable.** If nobody can see or use it, it is not done, it is inventory.
3. **Walking skeleton first.** The thinnest possible path through the whole system, even if it is ugly and hardcoded.
4. **Peel by variation:** happy path, then other roles, then other data types, then error handling, then performance.

## Splitting Patterns (try them in order)
- By workflow step (create, then edit, then delete)
- By business rule (one pricing tier first)
- By data variation (CSV before Excel)
- By interface (API before UI, only if the API has a real consumer)
- By quality attribute (works, then fast, then pretty)
- By spike: if the unknown is the risk, timebox a spike and slice after

## Guardrails
- A slice larger than ~3 days of work gets sliced again.
- Never split into "frontend" and "backend" stories assigned to different people unless a contract is agreed first.
- Order slices by risk and learning, not by ease.

## Output Format
A numbered list of slices. Each has: **Title**, **User-visible outcome**, **Acceptance criteria**, **Rough size (S/M/L)**, and **What it unblocks or de-risks**. Close with the recommended first slice and why.
