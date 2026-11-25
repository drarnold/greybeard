---
name: comment-cleanup
description: Clean up code comments by removing AI-style over-commenting, narration, and noise while keeping comments that carry real context. Use when asked to clean, tidy, prune, or audit comments, when comments are too verbose, or before committing an implementation.
---

# Comment Cleanup

A comment-only editing pass. Walk every comment in the target, keep the few that carry context the code cannot, delete or rewrite the rest.

## Scope
**Edit comments only. Never change code.** No renames, extractions, or logic changes, even when a comment exists only because the code is unclear. That comment may survive and gets flagged for a separate refactor. Tidy orphaned blank lines, but the code must be otherwise byte-identical.

## Core Principle
Code shows **how**. A comment earns its place only by carrying **why**: a non-obvious constraint, deliberate deviation, gotcha, workaround, or the reason a tempting simpler version is wrong. In doubt about a *why*, keep it. In doubt about a *how*, delete it. Wrongly deleting a real warning costs more than leaving a mediocre comment.

## The Pass (every comment, in order)
1. **Narration?** Restates code, steps, names, types, or signatures, or marks block ends. **Delete.**
2. **Change narration?** "Updated to v2", "fixed the off-by-one", "added per feedback". Would it make sense to a fresh reader a year from now? If not, **delete**. It belongs in the commit message.
3. **Moving target?** Points at a spec, requirements doc, design doc, or any section number ("spec §7"). **Delete**, or **rewrite** to encode the substance. Durable breadcrumbs are fine: ticket IDs, RFCs, permalinks, and a maintained README at a stable path. Even then the comment must stand alone with the link removed. This applies inside docstrings too.
4. **Bloated or misplaced *why*?**
   - If a maintained doc already carries the system-level narrative, **delete the inline copy** and leave no pointer-only signpost. If no doc exists, keep it and **flag** it as extraction-worthy.
   - Trim throat-clearing and anything the code shows.
   - Split long header blocks into short comments next to the lines they govern.
   - Apply Occam's razor to every survivor: the one non-obvious fact, in the fewest words. "Carries a real why" and "is worded minimally" are separate tests. A five-line block rarely survives intact. The razored answer is sometimes zero.
5. **Stale?** Describes code or behavior that no longer exists. **Delete** or correct. A wrong comment is worse than none.
6. **None of the above?** **Keep**, and do not reword comments that are already tight.

## Easy-to-Misread Keeps
- **Cross-file sync pointers:** "keep in sync with X". The link is the why.
- **Data-literal semantics:** units, order, or convention of a literal (`# pence, not pounds`).
- **Format contracts:** output pinned to an external expectation.
- **Section banners** (`# --- Name ---`): navigation, not narration. Keep them.

## TODOs
Keep TODO/FIXME without issue IDs. Delete one if the work is visibly done. If a TODO is deferral in disguise (`# TODO: handle errors` on a path that should already), keep it and **flag** it.

## Docstrings
All six rules apply. Delete docstrings that only re-emit name and types. Keep contract the signature cannot show: units, ranges, side effects, failure modes, invariants. Condense to one line only when nothing beyond restatement is lost. A one-line summary on a public function or endpoint is legitimate.

## Output
1. **Counts:** deleted, rewritten, kept, per file.
2. **Flags:** comments compensating for unclear code, deferral TODOs, comments whose correctness could not be verified.
3. **Judgment calls:** borderline decisions, one line each, so the user can overrule.

Keep the report brief. Do not list every deleted narration comment.

## Examples
```python
# Delete (narration on clean code)
# Parse the JSON body of the request into a dict
info = json.loads(request.data)

# Rewrite (real why, trimmed to one line)
# Before: 5-line block on log-filter prefixes and a future pipeline
# After:  Stable first-key prefix ({"stream":"bi.telemetry") for log filtering.

# Keep (surprising but essential)
# nextValue() may return a value that equals() null; both checks are required
if value is None or value == NULL_SENTINEL:
    return None
```
