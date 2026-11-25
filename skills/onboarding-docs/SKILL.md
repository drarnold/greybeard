---
name: onboarding-docs
description: Create and fix README, onboarding, and runbook documentation so a new engineer can be productive fast. Use when docs are stale or missing, a new hire is joining, or tribal knowledge needs capturing.
---

# Onboarding Docs

Docs are tested by the next person who has to follow them. Be that person.

## The Test
Can someone with zero context clone, build, run, and test the project using only the README, in under an hour? Try it fresh. Every stumble is a doc bug.

## Include
1. **What and why:** One paragraph. What this does and who uses it.
2. **Quickstart:** Exact copy-pasteable commands, prerequisites with versions, and expected output.
3. **Architecture at a glance:** One diagram and a map of the main directories.
4. **Common tasks:** run tests, add a migration, cut a release, debug locally.
5. **Operations:** dashboards, alerts, on-call, and runbooks for known failure modes.
6. **Who to ask:** owners and channels.
7. **Gotchas:** the weird thing that always bites people. Write it down.

## Keep It Alive
- Docs live next to code and change in the same PR.
- Automate what you can: make `make setup` do the steps instead of describing them.
- Date or link-check docs periodically. Delete what is wrong. Wrong docs are worse than none.

## Guardrails
- Do not document what the code makes obvious. Document the why and the surprises.
- Never put secrets in docs.

## Output Format
A drafted or revised document using the structure above, plus a list of **gaps found** that need a human to fill in.
