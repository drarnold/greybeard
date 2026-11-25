---
name: dependency-hygiene
description: Audit and manage third-party dependencies for vulnerabilities, licensing, bloat, and upgrade risk. Use when adding a library, reviewing lockfile changes, handling a CVE, or planning upgrades.
---

# Dependency Hygiene

Every dependency is code you did not write, cannot fully review, and will be blamed for.

## Before Adding One
1. Do you need it? A 10-line function beats a 10 MB package.
2. Is it maintained? Check recent releases, open issues, bus factor.
3. License compatible with your use?
4. Size and transitive dependency count acceptable?
5. Any known vulnerabilities?

## Ongoing
- **Commit lockfiles.** Pin exact versions in applications. Use ranges thoughtfully in libraries.
- **Automate updates** (Dependabot or Renovate) in small, frequent PRs. A year of drift is a migration project.
- **Scan continuously** for known vulnerabilities and act on severity and reachability, not just a scanner's scary number.
- **Remove unused dependencies.** Dead weight still expands attack surface.
- **Verify provenance:** beware typosquats and unexpected new maintainers.

## Upgrading
- Read the changelog for breaking changes. Upgrade one major at a time.
- Run the full test suite. Roll out like any other change.

## Guardrails
- Never `curl | bash` in a build. Pin by digest or hash where possible.
- Do not vendor and forget. Vendored code still needs patches.

## Output Format
A table of dependencies reviewed with **risk**, **reason**, and **action** (keep, upgrade, replace, remove). Lead with anything exploitable.
