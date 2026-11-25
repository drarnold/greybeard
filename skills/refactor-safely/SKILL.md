---
name: refactor-safely
description: Restructure code in small, behavior-preserving steps backed by tests. Use when cleaning up messy code, renaming or extracting modules, or paying down structural debt without changing what the code does.
---

# Refactor Safely

Refactoring means changing structure without changing behavior. If behavior changes, it is a feature or a bug fix, and it deserves its own commit.

## Before Touching Anything
1. **Lock behavior with tests.** If coverage is thin, write characterization tests that capture what the code does *today*, warts included.
2. **Define the goal.** "Cleaner" is not a goal. "Extract pricing from the controller so it can be unit tested" is.
3. **Keep the build green** at every step.

## The Moves
- Small, named steps: rename, extract function, inline, move, introduce parameter object.
- Run tests after each step. Commit after each green.
- Separate **refactor commits** from **behavior commits**. Reviewers will thank you.
- For large changes, use the strangler or branch-by-abstraction pattern: build the new path beside the old, shift traffic, delete the old.

## Guardrails
- No drive-by refactors inside a bug-fix PR.
- Do not refactor code with no tests and no reason to touch it. Ugly and stable beats pretty and broken.
- Public APIs and serialized formats are not refactorable by yourself. Treat them as contracts.
- If a step breaks tests, revert the step. Do not debug your way forward.

## Output Format
A numbered sequence of small steps, each with its intent, the verification run, and a suggested commit message. Flag any step that risks a behavior change.
