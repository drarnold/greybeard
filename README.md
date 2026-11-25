# platform-engineering-skills

[![Skill Evals](https://github.com/drarnold/platform-engineering-skills/actions/workflows/evals.yml/badge.svg)](https://github.com/drarnold/platform-engineering-skills/actions/workflows/evals.yml)

Production-grade, evaluation-tested **Agent Skills** (`SKILL.md`) for Staff Engineers, SREs, and Platform teams.

Unlike generic prompt collections, every skill in this repository is engineered as **deterministic developer infrastructure**: strictly scoped to minimize context-window bloat, governed by human-in-the-loop safety guardrails, and validated against automated regression benchmarks.

Compatible with **Claude Code**, **Cursor**, **Windsurf**, and **OpenAI Codex CLI**.

---

## Core Design Principles

1. **Progressive Context Disclosure:** Skills stay compact to preserve working memory for codebase context and stack traces.
2. **Guardrail-First Execution:** Explicit read-only boundaries prevent agents from executing destructive `kubectl`, Terraform, or SQL mutations without human sign-off.
3. **Structured, Auditable Outputs:** Enforces standardized Markdown schemas (e.g., SEV-level matrices, zero-downtime migration checklists) ready to paste directly into Slack, Jira, or GitHub PRs.
4. **Eval-Driven Reliability:** Every skill includes test fixtures in `/evals` and Go tooling in `/cmd` to verify that skills pass lint checks and catch real-world failure modes.

---

## Available Skills

| Skill | Trigger / Use Case | Key Guardrails & Outputs |
| :--- | :--- | :--- |
| **`incident-triage`** | Live outages, noisy alerts, or post-incident reviews | Enforces read-only diagnostics; outputs Timeline, Blast Radius, Root Cause (5 Whys), and Action Items. |
| **`architecture-rfc`** | New service proposals, data store migrations, or scaling bottlenecks | Forces explicit trade-off matrices, failure-mode analysis, and cost/latency estimates. |
| **`ci-cd-hardening`** | Reviewing `Dockerfile` or `.github/workflows/*.yml` changes | Flags privilege escalation, unpinned SHAs, secret leakage, and missing build-cache layers. |
| **`observability-slos`** | Instrumenting new services or tuning noisy PagerDuty alerts | Generates OpenTelemetry semantic conventions and multi-window, multi-burn-rate SLO alerts. |
| **`code-review-staff`** | High-risk PRs touching concurrency, public APIs, or DB schemas | Blocks table-locking migrations, detects N+1/race conditions, and checks backward compatibility. |

---

## Quick Start

Run the Go CLI installer directly:

```bash
# Install all platform skills for Claude Code (.claude/skills)
go run ./cmd/installer --target claude --all

# Install a specific skill for Cursor (.cursor/rules)
go run ./cmd/installer --target cursor incident-triage
```

---

## Running Lints & Evals

```bash
# Static SKILL.md validation (no API key required)
go run ./cmd/evals --lint-only

# Live model evaluation suite
export ANTHROPIC_API_KEY="your-key"
go run ./cmd/evals
```
