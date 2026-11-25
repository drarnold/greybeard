---
name: postmortem-writer
description: Write blameless postmortems that capture a clear timeline, contributing factors, and owned action items. Use after an incident, outage, or near miss when documenting what happened and how to prevent it.
---

# Blameless Postmortem

The goal is learning, not punishment. People did what made sense at the time with the information they had. Find out why it made sense.

## Structure
1. **Summary:** What broke, who was affected, for how long. Two or three sentences.
2. **Impact:** Users, revenue, SLO budget burned, data affected.
3. **Timeline:** Timestamped, factual (detection, escalation, mitigation, resolution). Include what was *believed* at each point.
4. **Contributing factors:** Technical and organizational. Avoid "root cause" singular. Incidents are usually several things lining up.
5. **What went well / what went poorly / where we got lucky.** Luck is a finding.
6. **Action items:** Each with an owner, a due date, and a type (prevent, detect, mitigate).

## Guardrails
- Never name a person as the cause. Say "the deploy process allowed X."
- No action item without an owner and a date. "Be more careful" is not an action.
- Ask "how did this reach production?" and "why did detection take so long?" before "who did it?"
- Prefer fixes that make the right thing the easy thing over fixes that rely on memory.
- Follow up. A postmortem whose actions rot is theater.

## Output Format
Use the six sections above. Keep the timeline as a table and the action items as a checklist with owners and dates. Pair with `incident-triage` for the response itself.
