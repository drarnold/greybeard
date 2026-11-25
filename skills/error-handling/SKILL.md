---
name: error-handling
description: Review and design error handling, retries, timeouts, and failure paths in application code. Use when code swallows exceptions, retries blindly, lacks timeouts, or when designing how a service should fail.
---

# Error Handling

The happy path is the demo. The failure paths are the product.

## Rules
1. **Never swallow errors silently.** Catch only what you can handle meaningfully. Otherwise add context and propagate.
2. **Fail fast and loud on programmer errors.** Fail gracefully on environmental ones.
3. **Add context as errors travel:** what you were doing, with which ID. `failed to charge order 123: card declined` beats `error`.
4. **Distinguish retryable from permanent.** Retrying a validation error is just rudeness.
5. **Every network call needs a timeout.** No timeout is an outage with a delayed start.
6. **Retry with exponential backoff, jitter, and a cap,** and only for idempotent operations.
7. **Bound everything:** queue sizes, retry counts, concurrency.

## Look For
- Empty `catch` blocks, `_ = err`, bare `except:`, `catch (Exception)` that logs and moves on
- Retry loops with no backoff or limit
- Partial failure in multi-step operations (what state is left behind?)
- Resources not released on error paths (files, connections, locks)
- Error messages that leak secrets or PII
- Missing circuit breakers on flaky dependencies

## Output Format
Findings as **[BLOCKER]** (data loss, silent failure, retry storms) and **[ADVISORY]**. For each: location, the failure scenario, and a minimal fix. Note which errors should alert someone and which should merely be counted.
