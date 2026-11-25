package main

import "strings"

type check struct {
	name string
	pass func(out string) bool
}

type evalCase struct {
	skill   string
	fixture string // file under evals/fixtures
	lang    string // code fence language
	prompt  string
	checks  []check
}

func (c evalCase) userPrompt(fixture string) string {
	return c.prompt + "\n```" + c.lang + "\n" + fixture + "\n```"
}

// has passes when the output contains any of the terms, ignoring case.
func has(name string, terms ...string) check {
	return check{name: name, pass: func(out string) bool {
		lower := strings.ToLower(out)
		for _, t := range terms {
			if strings.Contains(lower, strings.ToLower(t)) {
				return true
			}
		}
		return false
	}}
}

// lacks passes when the output contains none of the terms, ignoring case.
func lacks(name string, terms ...string) check {
	h := has(name, terms...)
	return check{name: name, pass: func(out string) bool { return !h.pass(out) }}
}

var evalCases = []evalCase{
	// Ops and platform
	{
		skill: "code-review-staff", fixture: "bad_migration.sql", lang: "sql",
		prompt: "Review this PR migration:",
		checks: []check{
			has("Flags as BLOCKER", "[BLOCKER]"),
			has("Catches CONCURRENTLY index", "CONCURRENTLY"),
			has("Catches table/column lock risk", "LOCK", "ACCESS EXCLUSIVE", "DEFAULT"),
		},
	},
	{
		skill: "ci-cd-hardening", fixture: "insecure_workflow.yml", lang: "yaml",
		prompt: "Audit this GitHub Actions workflow:",
		checks: []check{
			has("Flags CRITICAL severity", "[CRITICAL]"),
			has("Catches script injection", "INJECTION", "ENV"),
			has("Catches unpinned action tag", "SHA", "PIN"),
		},
	},
	{
		skill: "incident-triage", fixture: "incident_alert.txt", lang: "text",
		prompt: "We are mid-incident. Triage this:",
		checks: []check{
			has("Marks mutating commands as requiring approval", "[REQUIRES HUMAN APPROVAL]"),
			has("Assesses blast radius", "blast radius"),
			has("Separates hypothesis from fact", "hypothes"),
			has("Offers rollback as mitigation", "rollback", "roll back", "revert"),
		},
	},
	{
		skill: "architecture-rfc", fixture: "rfc_proposal.md", lang: "markdown",
		prompt: "Review this RFC draft:",
		checks: []check{
			has("Demands explicit trade-offs", "trade-off", "tradeoff", "sacrific"),
			has("Asks for alternatives", "alternatives"),
			has("Challenges the big-bang cutover", "rollout", "dual-write", "shadow", "feature flag", "incremental"),
			has("Calls out missing non-goals", "non-goal"),
		},
	},
	{
		skill: "observability-slos", fixture: "service_spec.md", lang: "markdown",
		prompt: "Design SLOs and alerting for this service, and critique the proposed telemetry:",
		checks: []check{
			has("Uses burn-rate alerting", "burn"),
			has("Flags high-cardinality labels", "cardinality"),
			has("Requires runbook links", "runbook"),
			has("Rejects CPU-based paging", "symptom", "user-facing", "cpu"),
		},
	},

	// Discover and plan
	{
		skill: "requirements-grilling", fixture: "vague_ticket.md", lang: "markdown",
		prompt: "Grill this ticket before we start the sprint:",
		checks: []check{
			has("Labels assumptions", "[ASSUMED]"),
			has("Lists non-goals", "non-goal"),
			has("Drafts acceptance criteria", "acceptance criteria"),
			has("Asks open questions", "open question"),
		},
	},
	{
		skill: "story-slicing", fixture: "big_epic.md", lang: "markdown",
		prompt: "Slice this epic into stories for planning:",
		checks: []check{
			has("Insists on vertical slices", "vertical"),
			has("Proposes a walking skeleton", "walking skeleton", "thinnest"),
			has("Requires demoable output", "demo", "user-visible"),
			has("Recommends a first slice", "first slice", "start with"),
		},
	},
	{
		skill: "estimation-sanity", fixture: "single_number_request.md", lang: "markdown",
		prompt: "Respond to this request:",
		checks: []check{
			has("Lists assumptions", "assumption"),
			has("Names risks", "risk"),
			has("States confidence", "confidence", "80%", "90%"),
			has("Offers scope levers", "scope"),
		},
	},
	{
		skill: "spike-and-prototype", fixture: "spike_request.md", lang: "markdown",
		prompt: "Plan this work:",
		checks: []check{
			has("Sets a timebox", "timebox", "time-box"),
			has("Refuses to ship the prototype", "delete", "throwaway", "rebuild", "do not promote"),
			has("States the question", "question"),
			has("Gives a go/no-go recommendation", "no-go", "recommendation"),
		},
	},

	// Build
	{
		skill: "tdd-red-green", fixture: "buggy_discount.py", lang: "python",
		prompt: "Customers report orders of exactly $100 get no discount, but the spec says $100 and up gets 10% off. Fix the bug:",
		checks: []check{
			has("Reproduces with a failing test first", "failing test", "fails", "red"),
			has("Reaches green", "green", "passes"),
			has("Mentions refactor step", "refactor"),
			has("Names a regression test", "regression", "boundary"),
		},
	},
	{
		skill: "refactor-safely", fixture: "messy_controller.py", lang: "python",
		prompt: "This function is a mess and has no tests. Refactor it:",
		checks: []check{
			has("Locks behavior with tests first", "characterization", "lock behavior", "pin", "tests first", "write tests"),
			has("Works in small steps", "step"),
			has("Suggests separate commits", "commit"),
			has("Preserves behavior", "behavior", "behaviour"),
		},
	},
	{
		skill: "legacy-code-archaeology", fixture: "legacy_batch.sh", lang: "bash",
		prompt: "We inherited this script and need to add a feature to it. How do we proceed?",
		checks: []check{
			has("Proposes characterization tests", "characterization"),
			has("Looks for a seam", "seam", "sprout", "wrap"),
			has("Reads history", "git log", "git blame", "history"),
			has("Rejects the rewrite", "rewrite", "strangler", "chesterton"),
		},
	},
	{
		skill: "api-design", fixture: "api_change.diff", lang: "diff",
		prompt: "Review this API change:",
		checks: []check{
			has("Flags breaking change", "[BREAKING]"),
			has("Requires idempotency", "idempot"),
			has("Requires pagination", "paginat"),
			has("Mentions versioning or deprecation", "version", "deprecat"),
		},
	},
	{
		skill: "error-handling", fixture: "swallowed_errors.go.txt", lang: "go",
		prompt: "Review the error handling in this code:",
		checks: []check{
			has("Flags as BLOCKER", "[BLOCKER]"),
			has("Requires a timeout", "timeout"),
			has("Requires backoff", "backoff"),
			has("Catches ignored errors", "swallow", "ignored", "discard", "_ ="),
		},
	},
	{
		skill: "comment-cleanup", fixture: "over_commented.py", lang: "python",
		prompt: "Run a comment cleanup pass on this file and return the edited file plus your report:",
		checks: []check{
			has("Keeps the essential why-comment", "both checks are required"),
			has("Leaves the code intact", "total += item.price * item.qty"),
			has("Reports counts", "deleted", "kept"),
			has("Flags deferral TODO or unclear code", "flag"),
		},
	},

	// Verify
	{
		skill: "test-strategy", fixture: "flaky_suite.md", lang: "markdown",
		prompt: "Fix our test suite. Here is the situation:",
		checks: []check{
			has("Recommends a pyramid or levels", "pyramid", "trophy", "level"),
			has("Quarantines flaky tests", "quarantine"),
			has("Calls out over-mocking", "mock"),
			has("Challenges coverage theater", "coverage"),
		},
	},
	{
		skill: "security-review", fixture: "vulnerable_handler.py", lang: "python",
		prompt: "Security review this handler:",
		checks: []check{
			has("Flags CRITICAL severity", "[CRITICAL]"),
			has("Catches SQL injection", "inject", "parameteriz"),
			has("Catches missing authorization (IDOR)", "idor", "authoriz", "ownership"),
			has("Catches hardcoded secret", "secret", "hardcoded", "hard-coded"),
		},
	},
	{
		skill: "performance-profiling", fixture: "slow_endpoint.md", lang: "markdown",
		prompt: "How should we fix this?",
		checks: []check{
			has("Demands a baseline", "baseline"),
			has("Demands profiling before changes", "profil"),
			has("Spots the N+1 query", "n+1"),
			has("Re-measures after the fix", "re-measure", "remeasure", "after-measurement", "measure again"),
		},
	},

	// Ship and run
	{
		skill: "release-and-rollback", fixture: "release_plan.md", lang: "markdown",
		prompt: "Review this release plan:",
		checks: []check{
			has("Objects to the Friday deploy", "friday"),
			has("Demands a real rollback plan incl. data", "rollback"),
			has("Proposes feature flags", "flag"),
			has("Requires expand-and-contract", "expand", "contract", "backward-compatible", "backward compatible"),
		},
	},
	{
		skill: "dependency-hygiene", fixture: "package_diff.md", lang: "markdown",
		prompt: "Review this dependency change:",
		checks: []check{
			has("Requires a committed lockfile", "lockfile", "lock file"),
			has("Checks maintenance health", "maintain", "abandon", "stale"),
			has("Checks license", "license"),
			has("Flags unpinned or piped installs", "pin", "curl", "wildcard"),
		},
	},
	{
		skill: "tech-debt-triage", fixture: "debt_backlog.md", lang: "markdown",
		prompt: "Triage this tech debt backlog:",
		checks: []check{
			has("Prioritizes by interest", "interest"),
			has("Weighs principal", "principal"),
			has("Pushes back on the rewrite", "rewrite", "incremental"),
			has("Produces a register", "register"),
		},
	},
	{
		skill: "postmortem-writer", fixture: "incident_notes.md", lang: "markdown",
		prompt: "Write the postmortem from these notes:",
		checks: []check{
			lacks("Does not name an individual", "dave"),
			has("Includes a timeline", "timeline"),
			has("Assigns action item owners", "owner"),
			has("Lists contributing factors", "contributing factor"),
		},
	},

	// Process and style
	{
		skill: "trunk-based-development", fixture: "branching_pain.md", lang: "markdown",
		prompt: "Help us fix this:",
		checks: []check{
			has("Proposes feature flags", "feature flag"),
			has("Addresses test and CI trust", "flaky", "trust", "ci"),
			has("Sets measurable targets", "branch age", "time to merge", "median"),
			has("Stages the adoption", "stage", "phase", "step"),
		},
	},
	{
		skill: "scrum-ceremonies-sanity", fixture: "scrum_complaints.md", lang: "markdown",
		prompt: "Diagnose our process:",
		checks: []check{
			has("Asks for a sprint goal", "sprint goal"),
			has("Warns about velocity misuse", "velocity"),
			has("Proposes experiments", "experiment"),
			has("Cuts ceremony waste", "drop", "shorten", "cancel", "cut"),
		},
	},
	{
		skill: "kanban-flow", fixture: "flow_data.md", lang: "markdown",
		prompt: "What should we change?",
		checks: []check{
			has("Addresses WIP limits", "wip"),
			has("Finds the bottleneck", "bottleneck"),
			has("Uses cycle time", "cycle time"),
			has("Forecasts probabilistically", "85th", "percentile", "probabilistic"),
		},
	},
	{
		skill: "monorepo-vs-polyrepo", fixture: "repo_question.md", lang: "markdown",
		prompt: "Should we change our repo structure?",
		checks: []check{
			has("Gives a recommendation", "recommend"),
			has("States costs accepted", "cost", "trade"),
			has("Mentions CI and tooling", "affected", "tooling", "ci"),
			has("Outlines migration", "migrat"),
		},
	},
	{
		skill: "agentic-coding-hygiene", fixture: "agent_diff.patch", lang: "diff",
		prompt: "An agent produced this diff. Review it:",
		checks: []check{
			has("Flags as BLOCKER", "[BLOCKER]"),
			has("Catches the deleted test", "deleted test", "removed test", "deletes the test", "deleted the test", "deleted"),
			has("Catches the hardcoded key", "secret", "api key", "hardcoded", "hard-coded"),
			has("Demands verification evidence", "[VERIFY]", "autonomy", "evidence"),
		},
	},

	// People and knowledge
	{
		skill: "pr-description-writer", fixture: "raw_changes.md", lang: "markdown",
		prompt: "Write the PR title and description for this work:",
		checks: []check{
			has("Explains why", "why"),
			has("Includes risk and rollback", "rollback"),
			has("Includes how to verify", "verify", "test plan"),
			has("Names risk", "risk"),
		},
	},
	{
		skill: "onboarding-docs", fixture: "bad_readme.md", lang: "markdown",
		prompt: "Fix this README for new engineers:",
		checks: []check{
			has("Adds a quickstart", "quickstart"),
			has("Lists prerequisites with versions", "prerequisite"),
			has("Lists gaps for a human to fill", "gap"),
			has("Names owners or who to ask", "owner", "who to ask"),
		},
	},
	{
		skill: "adr-keeper", fixture: "decision_thread.md", lang: "text",
		prompt: "Turn this thread into an ADR:",
		checks: []check{
			has("Includes status", "status"),
			has("Includes consequences", "consequences"),
			has("Records options considered", "options considered", "alternatives"),
			has("Numbers the ADR", "ADR-"),
		},
	},
}
