#eng-architecture thread, Oct 14

@maria: We need a store for the new ledger service. DynamoDB or Postgres?
@kofi: Dynamo scales for free and we already run it for sessions. But the ledger needs multi-row transactions and ad-hoc queries for finance.
@maria: Finance runs SQL reports daily, and Dynamo makes that painful. Also we're about 5k writes/min, nowhere near Postgres limits.
@kofi: True. Ops-wise we'd have to run Postgres ourselves though... unless we use RDS.
@lena: RDS is fine, we have the Terraform module already. Backup and failover are covered.
@maria: Okay, so Postgres on RDS. Risk is if we 100x, we'd need to partition. Revisit if we pass 2M writes/min.
@kofi: +1. I'll write it up.
