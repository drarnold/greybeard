---
name: monorepo-vs-polyrepo
description: Evaluate repository structure trade-offs between monorepo and polyrepo, including tooling, ownership, and CI. Use when starting a multi-service project, splitting or merging repos, or fighting cross-repo change pain.
---

# Monorepo vs Polyrepo

Neither is correct. Each trades one kind of pain for another.

## Trade-Offs
**Monorepo**
- Atomic cross-project changes, easy refactors, shared tooling, one dependency version.
- Needs build caching and affected-only CI, code ownership files, and discipline about boundaries. Without them: slow CI and tangled coupling.

**Polyrepo**
- Clear ownership, independent release cadence, small clones, natural access control.
- Cross-repo changes become multi-PR choreography, versions drift, and shared libraries become diplomacy.

## Decide With These Questions
1. How often does one change span several components?
2. Do teams release independently, or in lockstep?
3. Is there tooling investment available (Bazel, Nx, Turborepo, or similar)?
4. What are the access-control and compliance needs?
5. How large is it, and how large will it be in two years?

## Guardrails
- A monorepo is not permission to ignore module boundaries. Enforce them.
- Do not split repos to dodge a coupling problem. It just hides it behind a network.
- Whatever you pick, standardize CI, versioning, and CODEOWNERS.
- Migrating is expensive. Decide for the next two years, not the next two weeks.

## Output Format
**Recommendation**, the **top three reasons**, **Costs you are accepting**, and a **migration outline** if the current setup differs.
