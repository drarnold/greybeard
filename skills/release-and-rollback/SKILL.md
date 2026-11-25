---
name: release-and-rollback
description: Plan safe releases using feature flags, canaries, versioning, and tested rollback. Use when preparing a deploy, cutting a release, planning a risky rollout, or reviewing a release process.
---

# Release and Rollback

Deploy is not release. Decouple them, and always know the way back.

## Before Shipping
1. **Rollback plan, written and tested.** "Redeploy the old version" fails if a migration already ran. Plan for data too.
2. **Ship dark.** Merge behind a feature flag. Turn it on separately from the deploy.
3. **Progressive rollout:** internal, then 1%, 10%, 50%, 100%, with a pause and a metric check between steps.
4. **Define the abort criteria** in advance: error rate, latency, and business metric thresholds.
5. **Backward-compatible changes only** across the deploy window. Use expand-and-contract for schemas.

## Versioning
- Semantic versioning for libraries and APIs: breaking is major, additive is minor, fixes are patch.
- Keep a changelog written for humans. "Fixed bug" is not a changelog.
- Tag releases. Make builds reproducible.

## During and After
- Watch dashboards through the whole rollout, not just the first minute.
- Clean up flags once fully rolled out. Stale flags are tomorrow's incident.
- Do not release on a Friday. See Law 1.

## Output Format
A **release checklist** with: pre-flight checks, rollout stages with gates, abort criteria, rollback steps (code and data), communication plan, and post-release cleanup tasks.
