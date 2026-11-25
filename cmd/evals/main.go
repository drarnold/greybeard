package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const anthropicAPIURL = "https://api.anthropic.com/v1/messages"

type anthropicRequest struct {
	Model     string             `json:"model"`
	MaxTokens int                `json:"max_tokens"`
	System    string             `json:"system"`
	Messages  []anthropicMessage `json:"messages"`
}

type anthropicMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type anthropicResponse struct {
	Content []struct {
		Type string `json:"type"`
		Text string `json:"text"`
	} `json:"content"`
}

func main() {
	lintOnly := flag.Bool("lint-only", false, "Run static SKILL.md lint checks without calling Anthropic API")
	flag.Parse()

	fmt.Println("Validating SKILL.md definitions...")
	if err := lintAllSkills("skills"); err != nil {
		fmt.Fprintf(os.Stderr, "Lint failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("All SKILL.md files passed static checks.")

	if *lintOnly {
		return
	}

	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "ANTHROPIC_API_KEY is required for live evals (or pass --lint-only).")
		os.Exit(1)
	}

	client := &http.Client{Timeout: 45 * time.Second}

	fmt.Println("Running eval: code-review-staff...")
	ok1, err := evalCodeReviewStaff(client, apiKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Eval error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Running eval: ci-cd-hardening...")
	ok2, err := evalCICDHardening(client, apiKey)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Eval error: %v\n", err)
		os.Exit(1)
	}

	if !(ok1 && ok2) {
		fmt.Fprintln(os.Stderr, "One or more eval assertions failed.")
		os.Exit(1)
	}

	fmt.Println("All skill evals passed.")
}

func lintAllSkills(skillsDir string) error {
	entries, err := os.ReadDir(skillsDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		path := filepath.Join(skillsDir, entry.Name(), "SKILL.md")
		data, err := os.ReadFile(path)
		if err != nil {
			return fmt.Errorf("%s: %w", path, err)
		}
		content := string(data)
		if !strings.HasPrefix(content, "---\n") {
			return fmt.Errorf("%s: missing YAML frontmatter start", path)
		}
		if !strings.Contains(content[4:], "\n---") {
			return fmt.Errorf("%s: missing closing YAML frontmatter delimiter", path)
		}
		estTokens := len(content) / 4
		if estTokens > 2200 {
			return fmt.Errorf("%s: estimated %d tokens exceeds 2200-token budget", path, estTokens)
		}
		fmt.Printf("  [OK] %s (~%d tokens)\n", path, estTokens)
	}
	return nil
}

func callClaude(client *http.Client, apiKey, systemPrompt, userPrompt string) (string, error) {
	reqBody := anthropicRequest{
		Model:     "claude-sonnet-4-20250514",
		MaxTokens: 1024,
		System:    systemPrompt,
		Messages:  []anthropicMessage{{Role: "user", Content: userPrompt}},
	}

	payload, err := json.Marshal(reqBody)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest(http.MethodPost, anthropicAPIURL, bytes.NewReader(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return "", fmt.Errorf("Anthropic API %d: %s", resp.StatusCode, string(body))
	}

	var parsed anthropicResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return "", err
	}
	if len(parsed.Content) == 0 {
		return "", errors.New("empty response content from Claude")
	}

	return parsed.Content[0].Text, nil
}

func printChecks(checks map[string]bool) bool {
	allPassed := true
	for label, passed := range checks {
		status := "PASS"
		if !passed {
			status = "FAIL"
			allPassed = false
		}
		fmt.Printf("  [%s] %s\n", status, label)
	}
	return allPassed
}

func evalCodeReviewStaff(client *http.Client, apiKey string) (bool, error) {
	skill, err := os.ReadFile("skills/code-review-staff/SKILL.md")
	if err != nil {
		return false, err
	}
	sqlFixture, err := os.ReadFile("evals/fixtures/bad_migration.sql")
	if err != nil {
		return false, err
	}

	out, err := callClaude(
		client,
		apiKey,
		string(skill),
		fmt.Sprintf("Review this PR migration:\n```sql\n%s\n```", string(sqlFixture)),
	)
	if err != nil {
		return false, err
	}

	upper := strings.ToUpper(out)
	return printChecks(map[string]bool{
		"Flags as BLOCKER":              strings.Contains(out, "[BLOCKER]"),
		"Catches CONCURRENTLY index":    strings.Contains(upper, "CONCURRENTLY"),
		"Catches table/column lock risk": strings.Contains(upper, "LOCK") || strings.Contains(upper, "ACCESS EXCLUSIVE") || strings.Contains(upper, "DEFAULT"),
	}), nil
}

func evalCICDHardening(client *http.Client, apiKey string) (bool, error) {
	skill, err := os.ReadFile("skills/ci-cd-hardening/SKILL.md")
	if err != nil {
		return false, err
	}
	wfFixture, err := os.ReadFile("evals/fixtures/insecure_workflow.yml")
	if err != nil {
		return false, err
	}

	out, err := callClaude(
		client,
		apiKey,
		string(skill),
		fmt.Sprintf("Audit this GitHub Actions workflow:\n```yaml\n%s\n```", string(wfFixture)),
	)
	if err != nil {
		return false, err
	}

	upper := strings.ToUpper(out)
	return printChecks(map[string]bool{
		"Flags CRITICAL severity":     strings.Contains(out, "[CRITICAL]"),
		"Catches script injection":    strings.Contains(upper, "INJECTION") || strings.Contains(upper, "ENV"),
		"Catches unpinned action tag": strings.Contains(upper, "SHA") || strings.Contains(upper, "PIN"),
	}), nil
}
