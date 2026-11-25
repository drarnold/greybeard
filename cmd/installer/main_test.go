package main

import (
	"path/filepath"
	"testing"
)

func TestDestPathFor(t *testing.T) {
	cases := map[string]string{
		"claude":      filepath.Join(".claude", "skills", "s", "SKILL.md"),
		"codex":       filepath.Join(".codex", "skills", "s", "SKILL.md"),
		"gemini":      filepath.Join(".gemini", "skills", "s", "SKILL.md"),
		"copilot":     filepath.Join(".github", "skills", "s", "SKILL.md"),
		"opencode":    filepath.Join(".opencode", "skills", "s", "SKILL.md"),
		"kiro":        filepath.Join(".kiro", "skills", "s", "SKILL.md"),
		"commandcode": filepath.Join(".commandcode", "skills", "s", "SKILL.md"),
		"agents":      filepath.Join(".agents", "skills", "s", "SKILL.md"),
		"cursor":      filepath.Join(".cursor", "rules", "s.mdc"),
		"windsurf":    filepath.Join(".windsurf", "rules", "s.md"),
		"GEMINI":      filepath.Join(".gemini", "skills", "s", "SKILL.md"),
	}
	for target, want := range cases {
		got, err := destPathFor(target, "s")
		if err != nil {
			t.Errorf("%s: unexpected error: %v", target, err)
			continue
		}
		if got != want {
			t.Errorf("%s: got %q, want %q", target, got, want)
		}
	}
}

func TestDestPathForUnknownTarget(t *testing.T) {
	if _, err := destPathFor("emacs", "s"); err == nil {
		t.Fatal("expected error for unsupported target")
	}
}
