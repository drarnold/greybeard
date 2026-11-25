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
	"sync"
	"time"
)

const (
	anthropicAPIURL = "https://api.anthropic.com/v1/messages"
	defaultModel    = "claude-sonnet-4-20250514"
	maxConcurrent   = 4
	maxAttempts     = 3
)

type anthropicRequest struct {
	Model       string             `json:"model"`
	MaxTokens   int                `json:"max_tokens"`
	Temperature float64            `json:"temperature"`
	System      string             `json:"system"`
	Messages    []anthropicMessage `json:"messages"`
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

type evalResult struct {
	c      evalCase
	checks []checkResult
	err    error
}

type checkResult struct {
	name   string
	passed bool
}

func main() {
	lintOnly := flag.Bool("lint-only", false, "Run static SKILL.md lint checks without calling Anthropic API")
	only := flag.String("only", "", "Comma-separated skill names to evaluate (default: all)")
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

	cases, err := selectCases(evalCases, *only)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	model := os.Getenv("EVAL_MODEL")
	if model == "" {
		model = defaultModel
	}

	client := &http.Client{Timeout: 90 * time.Second}
	results := runCases(client, apiKey, model, cases)

	failed := 0
	for _, r := range results {
		fmt.Printf("Eval: %s\n", r.c.skill)
		if r.err != nil {
			fmt.Printf("  [ERROR] %v\n", r.err)
			failed++
			continue
		}
		ok := true
		for _, cr := range r.checks {
			status := "PASS"
			if !cr.passed {
				status = "FAIL"
				ok = false
			}
			fmt.Printf("  [%s] %s\n", status, cr.name)
		}
		if !ok {
			failed++
		}
	}

	if failed > 0 {
		fmt.Fprintf(os.Stderr, "%d of %d evals failed.\n", failed, len(results))
		os.Exit(1)
	}
	fmt.Printf("All %d skill evals passed.\n", len(results))
}

func selectCases(all []evalCase, only string) ([]evalCase, error) {
	if only == "" {
		return all, nil
	}
	want := map[string]bool{}
	for _, name := range strings.Split(only, ",") {
		want[strings.TrimSpace(name)] = true
	}
	var picked []evalCase
	for _, c := range all {
		if want[c.skill] {
			picked = append(picked, c)
			delete(want, c.skill)
		}
	}
	for name := range want {
		return nil, fmt.Errorf("no eval defined for skill %q", name)
	}
	return picked, nil
}

func runCases(client *http.Client, apiKey, model string, cases []evalCase) []evalResult {
	results := make([]evalResult, len(cases))
	sem := make(chan struct{}, maxConcurrent)
	var wg sync.WaitGroup
	for i, c := range cases {
		wg.Add(1)
		go func(i int, c evalCase) {
			defer wg.Done()
			sem <- struct{}{}
			defer func() { <-sem }()
			results[i] = runCase(client, apiKey, model, c)
		}(i, c)
	}
	wg.Wait()
	return results
}

func runCase(client *http.Client, apiKey, model string, c evalCase) evalResult {
	res := evalResult{c: c}

	skill, err := os.ReadFile(filepath.Join("skills", c.skill, "SKILL.md"))
	if err != nil {
		res.err = err
		return res
	}
	fixture, err := os.ReadFile(filepath.Join("evals", "fixtures", c.fixture))
	if err != nil {
		res.err = err
		return res
	}

	out, err := callClaude(client, apiKey, model, string(skill), c.userPrompt(string(fixture)))
	if err != nil {
		res.err = err
		return res
	}

	for _, ch := range c.checks {
		res.checks = append(res.checks, checkResult{name: ch.name, passed: ch.pass(out)})
	}
	return res
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

func callClaude(client *http.Client, apiKey, model, systemPrompt, userPrompt string) (string, error) {
	payload, err := json.Marshal(anthropicRequest{
		Model:       model,
		MaxTokens:   1500,
		Temperature: 0,
		System:      systemPrompt,
		Messages:    []anthropicMessage{{Role: "user", Content: userPrompt}},
	})
	if err != nil {
		return "", err
	}

	var lastErr error
	for attempt := 0; attempt < maxAttempts; attempt++ {
		if attempt > 0 {
			time.Sleep(time.Duration(1<<attempt) * 2 * time.Second)
		}
		out, retryable, err := doRequest(client, apiKey, payload)
		if err == nil {
			return out, nil
		}
		lastErr = err
		if !retryable {
			break
		}
	}
	return "", lastErr
}

func doRequest(client *http.Client, apiKey string, payload []byte) (out string, retryable bool, err error) {
	req, err := http.NewRequest(http.MethodPost, anthropicAPIURL, bytes.NewReader(payload))
	if err != nil {
		return "", false, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := client.Do(req)
	if err != nil {
		return "", true, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", true, err
	}
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		retry := resp.StatusCode == http.StatusTooManyRequests || resp.StatusCode >= 500
		return "", retry, fmt.Errorf("Anthropic API %d: %s", resp.StatusCode, string(body))
	}

	var parsed anthropicResponse
	if err := json.Unmarshal(body, &parsed); err != nil {
		return "", false, err
	}
	if len(parsed.Content) == 0 {
		return "", false, errors.New("empty response content from Claude")
	}
	return parsed.Content[0].Text, false, nil
}
