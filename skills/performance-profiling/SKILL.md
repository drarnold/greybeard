---
name: performance-profiling
description: Diagnose and fix performance problems by measuring first, finding the real bottleneck, and verifying the improvement. Use when something is slow, memory grows, or someone proposes an optimization without data.
---

# Performance Profiling

Do not optimize what you have not measured. Intuition about bottlenecks is wrong most of the time.

## The Method
1. **Define the metric and the target.** "p99 under 300 ms at 500 rps," not "faster."
2. **Reproduce and baseline.** Capture the number before changing anything.
3. **Profile.** CPU flame graph, heap profile, query plans (`EXPLAIN ANALYZE`), traces, and lock contention. Find where time and memory actually go.
4. **Fix the biggest bottleneck only.** Amdahl's law is not a suggestion.
5. **Re-measure.** If the number did not move, revert.

## Usual Suspects (check these first)
- N+1 queries and missing indexes
- Unbounded result sets and over-fetching
- Serial calls that could be parallel, or chatty remote calls
- Allocations in hot loops, accidental copies, string concatenation
- Missing caching, or caching without invalidation
- Lock contention and thread pool starvation
- Cold starts and connection setup cost

## Guardrails
- Never trade correctness for speed without saying so.
- Micro-benchmarks lie. Validate in a realistic environment with realistic data.
- Tail latency (p99) matters more than the average.
- Keep a perf regression test or dashboard so it stays fixed.

## Output Format
**Baseline**, **Evidence** (profile or plan excerpt), **Root cause**, **Fix**, **After-measurement**, and any **remaining risk**.
