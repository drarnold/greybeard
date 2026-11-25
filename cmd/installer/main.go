package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

const rawBaseURL = "https://raw.githubusercontent.com/drarnold/greybeard/main"

var allSkills = []string{
	"incident-triage",
	"architecture-rfc",
	"ci-cd-hardening",
	"observability-slos",
	"code-review-staff",
	// discover and plan
	"requirements-grilling",
	"story-slicing",
	"estimation-sanity",
	"spike-and-prototype",
	// build
	"tdd-red-green",
	"refactor-safely",
	"legacy-code-archaeology",
	"api-design",
	"error-handling",
	"comment-cleanup",
	// verify
	"test-strategy",
	"security-review",
	"performance-profiling",
	// ship and run
	"release-and-rollback",
	"dependency-hygiene",
	"tech-debt-triage",
	"postmortem-writer",
	// process and style
	"trunk-based-development",
	"scrum-ceremonies-sanity",
	"kanban-flow",
	"monorepo-vs-polyrepo",
	"agentic-coding-hygiene",
	// people
	"pr-description-writer",
	"onboarding-docs",
	"adr-keeper",
}

// skillDirTargets install a SKILL.md into <dir>/<skill>/SKILL.md.
var skillDirTargets = map[string]string{
	"claude":      filepath.Join(".claude", "skills"),
	"codex":       filepath.Join(".codex", "skills"),
	"gemini":      filepath.Join(".gemini", "skills"),
	"copilot":     filepath.Join(".github", "skills"),
	"opencode":    filepath.Join(".opencode", "skills"),
	"kiro":        filepath.Join(".kiro", "skills"),
	"commandcode": filepath.Join(".commandcode", "skills"),
	"agents":      filepath.Join(".agents", "skills"),
}

// ruleFileTargets install a single flat file <dir>/<skill><ext>.
var ruleFileTargets = map[string]struct{ dir, ext string }{
	"cursor":   {filepath.Join(".cursor", "rules"), ".mdc"},
	"windsurf": {filepath.Join(".windsurf", "rules"), ".md"},
}

func supportedTargets() string {
	names := make([]string, 0, len(skillDirTargets)+len(ruleFileTargets))
	for n := range skillDirTargets {
		names = append(names, n)
	}
	for n := range ruleFileTargets {
		names = append(names, n)
	}
	sort.Strings(names)
	return strings.Join(names, ", ")
}

func main() {
	target := flag.String("target", "claude", "Agent target: "+supportedTargets())
	installAll := flag.Bool("all", false, "Install all skills")
	flag.Parse()

	var selected []string
	if *installAll {
		selected = allSkills
	} else if flag.NArg() == 1 {
		selected = []string{flag.Arg(0)}
	} else {
		fmt.Fprintf(os.Stderr, "Usage: installer [--target %s] [--all | <skill-name>]\n", strings.ReplaceAll(supportedTargets(), ", ", "|"))
		os.Exit(1)
	}

	client := &http.Client{Timeout: 15 * time.Second}

	for _, skill := range selected {
		if err := installSkill(client, *target, skill); err != nil {
			fmt.Fprintf(os.Stderr, "Error installing %s: %v\n", skill, err)
			os.Exit(1)
		}
	}
}

func destPathFor(target, skill string) (string, error) {
	t := strings.ToLower(target)
	if dir, ok := skillDirTargets[t]; ok {
		return filepath.Join(dir, skill, "SKILL.md"), nil
	}
	if r, ok := ruleFileTargets[t]; ok {
		return filepath.Join(r.dir, skill+r.ext), nil
	}
	return "", fmt.Errorf("unsupported target %q (expected one of: %s)", target, supportedTargets())
}

func installSkill(client *http.Client, target, skill string) error {
	dest, err := destPathFor(target, skill)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("%s/skills/%s/SKILL.md", rawBaseURL, skill)
	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("fetch %s: %w", url, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected HTTP %d fetching %s", resp.StatusCode, url)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return fmt.Errorf("create dir: %w", err)
	}

	if err := os.WriteFile(dest, body, 0o644); err != nil {
		return fmt.Errorf("write %s: %w", dest, err)
	}

	fmt.Printf("Installed %s -> %s\n", skill, dest)
	return nil
}
