---
name: test-strategy
description: Design a pragmatic test strategy and diagnose flaky, slow, or low-value test suites. Use when deciding what to test and at which level, reviewing test quality, or fixing a suite nobody trusts.
---

# Test Strategy

Tests exist to let you change code with confidence. A suite you don't trust has negative value.

## Choose the Right Level
- **Unit:** pure logic, many, fast, no I/O.
- **Integration:** your code against a real database or queue, fewer, still automated.
- **Contract:** consumer and provider agree on the interface independently.
- **End-to-end:** a handful of critical user journeys, no more.
- Shape the suite as a pyramid or trophy. Never an ice cream cone.

## Review a Suite For
- **Flakiness:** sleeps, shared state, test order dependence, real clocks, real networks. Quarantine, then fix. Do not just retry.
- **Slowness:** the whole suite should give feedback in minutes. Parallelize, share fixtures sensibly, fake slow edges.
- **Over-mocking:** tests that assert the mocks were called prove the implementation, not the behavior.
- **Coverage theater:** high line coverage with no meaningful assertions.
- **Missing risk areas:** money, auth, migrations, concurrency, data deletion.

## Guardrails
- A bug found in production earns a regression test.
- Do not chase 100% coverage. Chase the risks.
- Delete tests that duplicate others or test nothing.

## Output Format
**Current risk map**, **Recommended mix by level**, **Top suite problems** (ranked by pain), and **Concrete next three actions**.
