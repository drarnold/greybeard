---
name: security-review
description: Application-level security review covering injection, authn/authz, secrets, input validation, and common OWASP-class flaws. Use when reviewing a diff or feature that handles user input, auth, data access, or secrets.
---

# Application Security Review

Assume every input is hostile and every dependency is a stranger.

## Check For
1. **Injection:** SQL, command, template, LDAP, and path traversal. Parameterize. Never concatenate untrusted input into an interpreter.
2. **Broken access control:** missing authz on routes, IDOR (user A fetches user B's record by changing an ID), privilege escalation, authorization done only in the UI.
3. **Authentication:** weak session handling, missing expiry, tokens in URLs, no rate limiting on login, hand-rolled crypto.
4. **Secrets:** keys in code, config, logs, or error messages. Rotate anything that leaked.
5. **XSS and CSRF:** unescaped output, missing CSRF protection on state-changing requests, permissive CORS.
6. **Unsafe deserialization and uploads:** untrusted pickle/YAML/Java objects, unvalidated file types, no size limits.
7. **SSRF:** server fetches a user-supplied URL with no allowlist.
8. **Sensitive data:** PII in logs, missing encryption in transit or at rest, overly broad responses.
9. **Dependencies:** known-vulnerable versions (see `dependency-hygiene`).

## Guardrails
- Authorization must be enforced server-side, per object, every time.
- Use vetted libraries for crypto and auth. Do not invent.
- Defense in depth: validate at the edge, constrain at the data layer.
- Only report what you can point to. No vague "might be insecure."

## Output Format
Findings tagged **[CRITICAL]**, **[HIGH]**, **[MEDIUM]**. Each with: location, attack scenario in one sentence, and the specific fix.
