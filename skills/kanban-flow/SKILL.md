---
name: kanban-flow
description: Improve continuous-flow delivery with WIP limits, flow metrics, and bottleneck analysis. Use when work stalls in review or testing, cycle time is unpredictable, or a team runs Kanban or wants less sprint overhead.
---

# Kanban Flow

Stop starting, start finishing. The fastest way to ship more is to work on less at once.

## Core Mechanics
1. **Visualize the real workflow.** Every column reflects an actual state, including "waiting for review."
2. **Limit WIP.** A limit that never bites is decoration. When blocked, swarm the blocker instead of pulling new work.
3. **Manage flow, not people.** Look at where work waits, not who is "slow."
4. **Make policies explicit:** definition of done, pull criteria, and class of service (expedite lane, normal, fixed-date).

## Measure
- **Cycle time:** start to done. Track the distribution and the 85th percentile, not just the mean.
- **Throughput:** items finished per week.
- **WIP and work item age:** old items in progress are the early warning.
- **Cumulative flow diagram:** widening bands mean a bottleneck.

## Forecasting
Use historical cycle time and throughput to give probabilistic dates ("85% by the 14th"). See `estimation-sanity`.

## Guardrails
- Do not let the expedite lane become the normal lane.
- Keep items small. Giant cards wreck flow metrics.
- Improve one bottleneck at a time, then re-measure.

## Output Format
**Where work waits**, **Suggested WIP limits** with reasoning, **Metrics to track**, and **one bottleneck experiment** with an expected effect.
