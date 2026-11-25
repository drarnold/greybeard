---
name: pr-description-writer
description: Write clear pull request descriptions and commit messages that explain why, risk, and how to verify. Use when opening a PR, writing a commit message, or improving a description reviewers keep asking questions about.
---

# PR Description Writer

The diff shows what changed. The description explains what the diff cannot: why, what could go wrong, and how to check.

## Structure
1. **Title:** Imperative, specific, under ~70 characters. `Add retry with backoff to payment client`, not `fixes`.
2. **Why:** The problem or ticket, in a few sentences. Link the issue.
3. **What changed:** The approach in plain words. Mention alternatives you rejected, briefly.
4. **Risk and rollback:** Blast radius, flags, migrations, and how to undo it.
5. **How to verify:** Tests run, manual steps, screenshots or output for UI and behavior changes.
6. **Reviewer notes:** Where to start, what is mechanical, what needs a hard look.

## Commit Messages
- Subject line imperative and short, body explains *why*, not just *what*.
- One logical change per commit. Separate refactors from behavior.

## Guardrails
- Keep PRs small. If the description needs sections for five unrelated things, split the PR.
- Never write "misc fixes" or "WIP" in anything that will be merged.
- Do not claim tests passed unless you ran them.
- No secrets or customer data in descriptions or screenshots.

## Output Format
Return the finished title and description in Markdown, ready to paste, with any open questions for the author listed separately.
