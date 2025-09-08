# Developer Guide

This document is intended for contributors to the openCenter CLI. It explains how to build the project, run tests, understand the architecture and extend the codebase.

## Building

The CLI is written in Go 1.24. To build a local binary:

```sh
go mod tidy      # download dependencies specified in go.mod
mise run build   # compiles the `openCenter` binary into the repository root
```

The build task defined in `.mise.toml` simply invokes `go build -o openCenter` from the project root.

## Testing and Behaviour‑Driven Development

The project includes a behaviour‑driven test suite using [Godog](https://github.com/cucumber/godog). The `.feature` files reside under `features/` and describe user‑facing behaviours. Each step is implemented in Go in `features/steps/steps_test.go` using the standard `go test` runner.

To run the BDD suite:

```sh
mise run bdd
```

This command compiles the CLI and runs `go test ./features/steps -v`, which in turn executes the Godog scenarios. The steps make use of helper functions in `features/steps/helpers.go` to build the binary once, isolate configuration directories via `OPENCENTER_CONFIG_DIR`, invoke the CLI and assert file or output expectations.

Alternatively, you can use the Godog CLI directly:

```sh
mise run bdd:godog
```

## Generating the JSON Schema

The CLI can export its configuration model as a JSON schema. This is implemented in `internal/config/schema.go` and exposed via `openCenter cluster schema`. Use the `mise run schema` task to write the schema into the `schema/` directory:

```sh
mise run schema
```

## Directory Structure

See `docs/DIRECTORY_STRUCTURE.md` for an overview of the repository layout.

## Extending the CLI

The CLI uses the [Cobra](https://github.com/spf13/cobra) library. New commands should be added under the `cmd/` directory and registered in `newClusterCmd` or via other appropriate parent commands. Configuration loading and saving is centralised in the `internal/config` package; avoid duplicating file or path logic elsewhere.

For template‑related changes, place new files under `internal/gitops/templates/`. They will be embedded at compile time via `go:embed` (see `internal/gitops/embed.go`). The `copy.go` module handles rendering and copying; additional helper functions can be added there.

## Testing Considerations

When adding new features, include scenarios in the existing feature files or create new `.feature` files under `features/`. Each scenario should be expressed in natural language and cover the expected behaviour from the user’s perspective. Steps should avoid external side effects; test helpers isolate configuration directories and use temporary file systems where appropriate.

Use Go’s standard testing features (e.g. `t.Run`, `testing` package) in conjunction with Godog to assert behaviour. Do not call `os.Exit` directly in commands; instead, return errors so that tests can capture non‑zero exit conditions.