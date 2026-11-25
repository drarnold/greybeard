# AGENTS.md

Guidance for AI coding agents (OpenCode, Codex, Copilot, Gemini CLI, Cursor, and others) working in this repository.

> **Scope:** This file configures agents working on the greybeard repo itself. The reusable assets are the skills in `skills/`, not this file. Do not copy it into other projects.

## Layout
- `skills/<name>/SKILL.md`: one skill per folder, the only thing users install.
- `cmd/installer`: copies skills into each agent's expected directory.
- `cmd/evals`: static lint plus live LLM evals. Run `go test ./...` and `go run ./cmd/evals --lint-only` before every PR; every skill needs a case in `cmd/evals/cases.go` and a fixture in `evals/fixtures/`.
- `evals/fixtures/`: deliberately bad inputs for evals.
- `CLAUDE.md`: code-comment policy.

## Working on Skills
- Keep each `SKILL.md` under 2,200 estimated tokens (lint enforces this).
- Frontmatter needs `name` (matching the folder) and a one-sentence `description` stating what it does and when to use it.
- Register a new skill in `allSkills` in `cmd/installer/main.go` and list it in the README.
- Add a fixture in `evals/fixtures/` and a case in `cmd/evals/cases.go`. `go test ./...` fails without one.

## Skill Routing
If a task matches a skill and the skill is installed, invoke it rather than improvising.

| Intent | Skill |
|---|---|
| Vague ticket, unclear requirements | `requirements-grilling` |
| Epic too big, plan the slices | `story-slicing` |
| "How long will this take?" | `estimation-sanity` |
| Technical unknown, evaluating a library | `spike-and-prototype` |
| Implement or fix test-first | `tdd-red-green` |
| Restructure without behavior change | `refactor-safely` |
| Old, untested, undocumented code | `legacy-code-archaeology` |
| New or changed endpoint or contract | `api-design` |
| Retries, timeouts, swallowed exceptions | `error-handling` |
| Too many or too verbose comments | `comment-cleanup` |
| What to test, flaky or slow suite | `test-strategy` |
| Auth, input handling, secrets | `security-review` |
| Something is slow or uses too much memory | `performance-profiling` |
| Reviewing a diff or migration | `code-review-staff` |
| Design doc or trade-off decision | `architecture-rfc`, then `adr-keeper` |
| Rollout, flags, rollback | `release-and-rollback` |
| Dockerfile, pipeline, GitHub Actions | `ci-cd-hardening` |
| Alerts, SLOs, telemetry | `observability-slos` |
| Adding or upgrading libraries | `dependency-hygiene` |
| Something is on fire | `incident-triage`, then `postmortem-writer` |
| Debt backlog | `tech-debt-triage` |
| Branching and merge pain | `trunk-based-development` |
| Sprint ceremonies not working | `scrum-ceremonies-sanity` |
| Work stalls, unpredictable cycle time | `kanban-flow` |
| Repo structure | `monorepo-vs-polyrepo` |
| Reviewing agent-written code | `agentic-coding-hygiene` |
| Writing a PR or commit message | `pr-description-writer` |
| README, onboarding, runbooks | `onboarding-docs` |

## Lifecycle
Define (`requirements-grilling`, `spike-and-prototype`) → Plan (`story-slicing`, `estimation-sanity`) → Build (`tdd-red-green`, `refactor-safely`, `api-design`, `error-handling`) → Verify (`test-strategy`, `security-review`, `performance-profiling`) → Review (`code-review-staff`, `pr-description-writer`) → Ship (`release-and-rollback`, `ci-cd-hardening`) → Run (`observability-slos`, `incident-triage`, `postmortem-writer`).
