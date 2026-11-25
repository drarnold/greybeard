# Release plan: orders v3

- Date: this Friday, 4:30 PM, so it's live for the weekend sale
- Scope: everything merged since March (about 140 PRs), shipped as one deploy
- DB: migration drops the old `orders.legacy_status` column and renames `orders.total` to `orders.total_cents`
- Rollout: deploy to all pods at once, no flag, no canary
- Rollback plan: "redeploy the previous version if something goes wrong"
- Monitoring: we will watch the dashboard for a few minutes after the deploy
- Owner: Jamie, who is on a flight Saturday morning
