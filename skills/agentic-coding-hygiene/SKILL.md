---
name: agentic-coding-hygiene
description: Practices for safely reviewing, trusting, and directing AI-generated code and coding agents. Use when reviewing agent-written diffs, setting guardrails for coding agents, or deciding how much autonomy to grant.
---

# Agentic Coding Hygiene

AI-written code is a pull request from a fast, confident, tireless junior who has never been paged. Review it accordingly.

## Reviewing Agent Output
1. **Read every line.** "The agent wrote it" is not a review.
2. **Check it actually ran.** Demand evidence: test output, build logs, not "this should work."
3. **Hunt for plausible nonsense:** invented APIs, nonexistent flags, wrong library versions, confident wrong comments.
4. **Watch for scope creep:** unrelated edits, deleted tests, loosened assertions, disabled lint rules.
5. **Look for hardcoded secrets, broad permissions, and swallowed errors.** Agents love `try/except: pass`.
6. **Prefer small diffs.** If you cannot review it, it does not merge.

## Setting Guardrails
- Scope tasks narrowly with clear acceptance criteria and a definition of done.
- Run agents in sandboxes with least privilege. No production credentials.
- Require human approval for destructive, irreversible, or outward-facing actions.
- Let CI be the referee: tests, linters, type checks, and security scans.
- Keep a project instructions file (conventions, commands, don'ts) so the agent stops rediscovering them.

## Guardrails on the Human
- Do not let the agent grade its own homework. Verify independently.
- Own what you merge. Accountability does not transfer to the tool.

## Output Format
A **review checklist** applied to the diff, with findings tagged **[BLOCKER]**, **[VERIFY]** (needs proof it works), or **[NIT]**, and a recommended autonomy level for the task.
