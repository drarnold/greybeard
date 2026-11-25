Raw notes from the Tuesday outage:

- 09:12 Dave pushed a config change to prod without testing it first. This is on him.
- 09:14 Error rate on checkout hit 60%. Nobody was paged because the alert threshold for checkout was set to 90%.
- 09:31 A customer tweeted about it. Support pinged the channel.
- 09:40 Priya found the bad value (pool size 0) and reverted it by hand on each of 12 pods. Took 25 minutes because there was no one-click rollback.
- 10:05 Error rate back to normal.
- Config changes are not validated or staged anywhere; anyone with prod access can edit the file directly.
- About 4,100 orders failed to complete. Customer support is still reaching out.
