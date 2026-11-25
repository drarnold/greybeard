# RFC: Move Orders to NovaDB

Status: Draft
Author: J. Rivera

## Context
Our Postgres orders database feels slow and the team wants to use something modern.

## Proposal
Replace Postgres with NovaDB, a new distributed database, for the entire orders domain.
Migrate all 400M rows in a single weekend: freeze writes Saturday 02:00, export, import,
and flip the connection string Sunday morning. NovaDB scales horizontally so we will
never have to think about capacity again, and it has no downsides we are aware of.

## Architecture
The orders service connects directly to NovaDB. Retries on failure are set to infinite
so no order is ever lost.

## Rollout
Saturday night. If it fails we will figure it out.
