Tech debt backlog (everyone's pet items):

1. The invoicing module is a 3,000-line file with no tests and awful naming. Nobody has changed it in 2 years and it has never caused an incident.
2. deploy.sh retries flaky steps by hand. It breaks about once a week and costs someone 1-2 hours each time. Three incidents this quarter.
3. Upgrade the ORM from v3 to v5. v3 is out of security support in 4 months.
4. Rename `utils/` to something meaningful.
5. A few engineers want to rewrite the whole backend in a new language "to clean everything up." Estimate: 14 months.
6. The search service logs 80 GB a day, which costs real money, and the on-call keeps muting its alerts.

We have one sprint of capacity this quarter. Leadership wants a recommendation.
