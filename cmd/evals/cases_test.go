package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Evals read skills/ and evals/fixtures/ relative to the repo root.
func TestMain(m *testing.M) {
	if err := os.Chdir("../.."); err != nil {
		panic(err)
	}
	os.Exit(m.Run())
}

func TestEverySkillHasAnEval(t *testing.T) {
	entries, err := os.ReadDir("skills")
	if err != nil {
		t.Fatal(err)
	}
	covered := map[string]bool{}
	for _, c := range evalCases {
		covered[c.skill] = true
	}
	for _, e := range entries {
		if e.IsDir() && !covered[e.Name()] {
			t.Errorf("skill %q has no eval case in cmd/evals/cases.go", e.Name())
		}
	}
}

func TestEvalCasesAreWellFormed(t *testing.T) {
	seen := map[string]bool{}
	for _, c := range evalCases {
		if seen[c.skill] {
			t.Errorf("duplicate eval case for %q", c.skill)
		}
		seen[c.skill] = true

		if _, err := os.Stat(filepath.Join("skills", c.skill, "SKILL.md")); err != nil {
			t.Errorf("%s: no such skill: %v", c.skill, err)
		}
		if _, err := os.Stat(filepath.Join("evals", "fixtures", c.fixture)); err != nil {
			t.Errorf("%s: missing fixture: %v", c.skill, err)
		}
		if strings.TrimSpace(c.prompt) == "" {
			t.Errorf("%s: empty prompt", c.skill)
		}
		if len(c.checks) < 3 {
			t.Errorf("%s: want at least 3 checks, got %d", c.skill, len(c.checks))
		}
		for _, ch := range c.checks {
			if ch.name == "" || ch.pass == nil {
				t.Errorf("%s: check with empty name or nil func", c.skill)
			}
		}
	}
}

func TestEveryFixtureIsUsed(t *testing.T) {
	entries, err := os.ReadDir(filepath.Join("evals", "fixtures"))
	if err != nil {
		t.Fatal(err)
	}
	used := map[string]bool{}
	for _, c := range evalCases {
		used[c.fixture] = true
	}
	for _, e := range entries {
		if !used[e.Name()] {
			t.Errorf("fixture %q is not referenced by any eval case", e.Name())
		}
	}
}

func TestHasIsCaseInsensitiveAnyOf(t *testing.T) {
	c := has("x", "Backoff", "jitter")
	if !c.pass("add exponential BACKOFF here") {
		t.Error("expected case-insensitive match")
	}
	if !c.pass("use jitter") {
		t.Error("expected any-of match")
	}
	if c.pass("nothing relevant") {
		t.Error("expected no match")
	}
}

func TestLacksIsInverseOfHas(t *testing.T) {
	c := lacks("x", "dave")
	if c.pass("Dave did it") {
		t.Error("expected failure when term present")
	}
	if !c.pass("the deploy process allowed it") {
		t.Error("expected pass when term absent")
	}
}

func TestSelectCases(t *testing.T) {
	got, err := selectCases(evalCases, "adr-keeper, api-design")
	if err != nil || len(got) != 2 {
		t.Fatalf("got %d cases, err %v", len(got), err)
	}
	if _, err := selectCases(evalCases, "nope"); err == nil {
		t.Error("expected error for unknown skill")
	}
}

func TestUserPromptFencesFixture(t *testing.T) {
	c := evalCase{prompt: "Review:", lang: "go"}
	want := "Review:\n```go\nBODY\n```"
	if got := c.userPrompt("BODY"); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
