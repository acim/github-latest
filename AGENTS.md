# Agent instructions

- Before pushing Go changes, run `make check` (the local mirror of the CI
  `go-check` gates: lint, govulncheck, `go fix -diff`) plus the tests affected
  by the change, and fix any failure it exposes.
