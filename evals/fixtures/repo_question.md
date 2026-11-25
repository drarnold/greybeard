We have 12 services across 12 repos, owned by 3 teams (about 25 engineers). A change to
the shared `common-models` library usually takes 5 PRs in 5 repos and a week of
coordination, and versions have drifted: 4 services are two majors behind. CI is a
simple pipeline per repo, about 8 minutes each. Nobody on the team has experience with
Bazel, Nx, or similar tools and we have no platform engineers to spare. A new VP wants
"everything in one repo like Google."

Should we change our repo structure?
