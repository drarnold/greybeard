# platform-engineering-skills

[![Skill Evals](https://github.com/drarnold/platform-engineering-skills/actions/workflows/evals.yml/badge.svg)](https://github.com/drarnold/platform-engineering-skills/actions/workflows/evals.yml)

**Opinionated, eval-tested AI agent skills for Staff Engineers, SREs, and Platform teams.**

Stop treating agent instructions like loose chat snippets. Every skill in this repo is packaged as a focused `SKILL.md` module, paired with a **Go 1.27** installer and an automated evaluation harness so agents follow production guardrails instead of guessing.

Works with **Claude Code**, **Cursor**, **Windsurf**, and **OpenAI Codex CLI**.

---

## Why This Exists

Most AI coding setups fall apart in production because agents:
- **Optimize for clever code over blast radius**—missing table locks, retry storms, or breaking rolling deploys.
- **Mutate state too eagerly**—suggesting or running destructive commands before confirming evidence.
- **Drift silently over time**—prompt tweaks break edge cases with zero regression checks.

This repository treats agent workflows like real platform infrastructure: **tightly scoped context**, **explicit guardrails**, **predictable output formats**, and **automated lint + LLM evals**.

---

## Quick Start

Install all skills or pick just the workflow you need:

```bash
# Install all skills for Claude Code (.claude/skills)
go run ./cmd/installer --target claude --all

# Install a single skill for Cursor (.cursor/rules)
go run ./cmd/installer --target cursor incident-triage

# Also supports: windsurf (.windsurf/rules) and codex (.codex/skills)
go run ./cmd/installer --target codex --all