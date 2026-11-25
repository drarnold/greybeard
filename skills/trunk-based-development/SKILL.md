---
name: trunk-based-development
description: Adopt and coach trunk-based development with short-lived branches, feature flags, and continuous integration. Use when branches live too long, merges are painful, or a team wants to move toward continuous delivery.
---

# Trunk-Based Development

Integrate early, integrate often. Long-lived branches are deferred merge conflicts with interest.

## The Practice
1. **Small, frequent merges** to main, at least daily. Branches live hours, not weeks.
2. **Main is always releasable.** Broken main is a stop-the-line event. Fix or revert within minutes.
3. **Hide unfinished work behind feature flags,** not branches.
4. **Fast CI.** If the pipeline takes an hour, people batch changes and the whole thing collapses.
5. **Branch by abstraction** for large changes: introduce the seam, switch over, remove the old path.
6. **Review quickly** (hours, not days), or pair and skip the queue.

## Prerequisites (be honest about these)
- Reliable automated tests the team trusts
- Feature flag tooling and a flag cleanup habit
- A fast build and an easy rollback

## When It Gets Hard
- Mandatory release branches are allowed for shipped versions needing patches, but fixes flow trunk-first.
- Code freezes mean the process lacks confidence. Fix the confidence.

## Guardrails
- Do not mandate it without the test and CI foundation. It will hurt.
- Delete flags after rollout (see `release-and-rollback`).

## Output Format
**Current state diagnosis** (branch age, merge pain, CI time), **Gaps to close first**, and a **staged adoption plan** with measurable targets such as median branch age and time to merge.
