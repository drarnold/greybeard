---
name: tdd-red-green
description: Drive implementation with test-driven development using small red-green-refactor cycles. Use when building a feature or fixing a bug test-first, or when asked to add tests before changing behavior.
---

# TDD: Red, Green, Refactor

Write the test you wish you had, watch it fail for the right reason, then make it pass with the least code that works.

## The Loop
1. **Red:** Write one small failing test for the next behavior. Run it. Confirm it fails *and* that the failure message is the one you expected. A test that never failed proves nothing.
2. **Green:** Write the simplest code that passes. Hardcoding is allowed. Elegance is not required yet.
3. **Refactor:** With tests green, clean up duplication and naming. Run tests after every change.
4. Repeat. Commit at green.

## Bug Fixes
Reproduce the bug as a failing test first. Then fix it. The test is the regression guard and the proof.

## What Makes a Good Test
- Tests **behavior through the public interface**, not private implementation details.
- One reason to fail. Descriptive name: `rejects_expired_token`, not `test1`.
- Fast, deterministic, independent. No sleeping, no shared state, no network.
- Mock at the boundary you do not own, not everything you can.

## Guardrails
- Never write production code without a failing test driving it (spikes excepted, see `spike-and-prototype`).
- Never edit a test to make it pass unless the requirement changed.
- Resist testing the framework or trivial getters.

## Output Format
Show each cycle briefly: the failing test, the minimal change, and any refactor. End with the full test run result.
