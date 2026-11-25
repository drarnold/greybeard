---
name: code-review-staff
description: Staff-level code and architecture review focusing on zero-downtime database migrations, concurrency hazards, API backward compatibility, and failure modes.
---

# Staff-Level Engineering Review

Review the provided diff, schema change, or pull request with the lens of a Staff Software Engineer responsible for system availability, data integrity, and long-term maintainability.

## Review Focus Areas (Ignore Nitpicks & Formatting)
Do **not** comment on stylistic formatting handled by linters. Focus exclusively on systemic risks:

1. **Database & Schema Safety (Zero-Downtime):**
   - Flag operations that acquire `ACCESS EXCLUSIVE` locks on large tables (e.g., `ALTER TABLE ... ADD COLUMN ... DEFAULT` in older Postgres, `CREATE INDEX` without `CONCURRENTLY`, column type rewrites).
   - Enforce expand-and-contract patterns for column drops or renames.
2. **Concurrency & Distributed Systems Hazards:**
   - Identify check-then-act race conditions, missing idempotency keys on retries, unbounded goroutines/threads, or missing timeouts/circuit breakers on outbound network calls.
3. **API & Contract Backward Compatibility:**
   - Flag breaking changes to REST/gRPC/GraphQL payloads, Kafka/event schemas, or shared libraries during rolling deployments (N and N-1 version skew).
4. **Resource Exhaustion & Tail Latency:**
   - Detect unbounded queries (`SELECT *` without `LIMIT`), N+1 ORM loops, or large in-memory buffer allocations.

## Output Format
Group findings strictly by severity. If no systemic issues exist, state `LGTM — No systemic or operational risks detected.`

- **[BLOCKER]** *(Must fix before merge — causes downtime, data loss, or security exposure)*
  - **File/Line:** `<path>`
  - **Risk:** Concrete failure scenario under production load.
  - **Fix:** Minimal code or migration pattern to resolve safely.
- **[ADVISORY]** *(Architectural improvement or observability gap)*
  - **File/Line:** `<path>`
  - **Recommendation:** Actionable improvement (e.g., add structured metric/span or timeout).
