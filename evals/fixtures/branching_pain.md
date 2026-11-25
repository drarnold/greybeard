How we work today:

- Feature branches live 2 to 4 weeks. Merging them takes a day of conflict resolution.
- Every release has a two-week code freeze on `main` and a release branch we cherry-pick into.
- CI takes 55 minutes and fails randomly about 15% of the time.
- Test coverage is patchy. People say "we can't merge small things, the tests are too flaky to trust."
- We have no feature flag system.
- Leadership wants us to "do trunk-based development by next sprint."
