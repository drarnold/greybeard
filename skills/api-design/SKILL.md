---
name: api-design
description: Design and review HTTP, gRPC, and library APIs for consistency, evolvability, and safe versioning. Use when creating a new endpoint, changing a contract, or reviewing API shape and error handling.
---

# API Design

An API is a promise made to people you will never meet. Make fewer promises and keep them forever.

## Principles
1. **Design from the consumer's side.** Write the call you wish you could make, then build to it.
2. **Consistency beats cleverness.** Naming, pagination, filtering, and errors work the same everywhere.
3. **Additive changes only.** Add optional fields. Never remove, rename, or retype in place. Breaking changes require a new version and a deprecation window.
4. **Be strict in what you send, liberal in what you accept,** but validate at the boundary.

## Checklist
- **Resources and verbs:** correct HTTP methods, idempotent `PUT`/`DELETE`, idempotency keys on unsafe retries of `POST`.
- **Errors:** stable machine-readable codes, human messages, correct status codes, no stack traces.
- **Pagination:** cursor-based for large or changing data. Always bounded.
- **Auth and limits:** authN/authZ on every route, rate limits, request size limits.
- **Timeouts and retries:** documented behavior and safe-to-retry semantics.
- **Versioning:** policy stated up front (path, header, or additive-only).
- **Docs and examples:** OpenAPI or proto as the source of truth, with real request/response samples.

## Guardrails
- Never leak internal IDs, table names, or implementation details through the contract.
- Do not ship an endpoint with no owner and no deprecation plan.

## Output Format
Group findings as **[BREAKING]**, **[RISK]**, and **[POLISH]**. For each, show the current shape, the problem, and a concrete fix.
