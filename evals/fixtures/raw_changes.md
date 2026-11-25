Ticket: PAY-318 "Payment provider sometimes times out and customers get double charged"

Commits on the branch:
- wip
- fix stuff
- more fixes
- address comments
- ok now

What changed (author's notes, rough):
- payment client now retries timeouts up to 3 times with exponential backoff
- added an idempotency key column to `payments` (new migration, nullable, no backfill)
- behind flag `payments_retry_v2`, off by default
- unit tests for the retry loop; didn't try against the sandbox PSP
- touches the checkout path, which is high traffic
