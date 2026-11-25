Our test suite situation:

- Full run takes 42 minutes. People batch changes because of it.
- 30 end-to-end browser tests, 40 integration tests, about 150 unit tests.
- Roughly 1 in 6 CI runs fails for no reason. We have set CI to retry failures three times.
- Many tests use `sleep(5)` to wait for things. Some depend on the order they run in.
- Line coverage is 92%, which leadership likes. Most unit tests mock every collaborator
  and assert that the mocks were called with certain arguments.
- Last quarter's two production incidents were in payment rounding and a data-deletion job.
  Neither had tests.
- The team wants to add 200 more end-to-end tests to feel safer.
