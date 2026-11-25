package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const rawBaseURL = "https://raw.githubusercontent.com/drarnold/platform-engineering-skills/main"

var allSkills = []string{
	"incident-triage",
	"architecture-rfc",
	"ci-cd-hardening",
	"observability-slos",
	"code-review-staff",
}

func main() {
	target := flag.String("target", "claude", "Agent target: claude, cursor, windsurf, or codex")
	installAll := flag.Bool("all", false, "Install all skills")
	flag.Parse()

	var selected []string
	if *installAll {
		selected = allSkills
	} else if flag.NArg() == 1 {
		selected = []string{flag.Arg(0)}
	} else {
		fmt.Fprintln(os.Stderr, "Usage: installer [--target claude|cursor|windsurf|codex] [--all | <skill-name>]")
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
	switch strings.ToLower(target) {
	case "claude":
		return filepath.Join(".claude", "skills", skill, "SKILL.md"), nil
	case "cursor":
		return filepath.Join(".cursor", "rules", skill+".mdc"), nil
	case "windsurf":
		return filepath.Join(".windsurf", "rules", skill+".md"), nil
	case "codex":
		return filepath.Join(".codex", "skills", skill, "SKILL.md"), nil
	default:
		return "", fmt.Errorf("unsupported target %q (expected claude, cursor, windsurf, or codex)", target)
	}
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
