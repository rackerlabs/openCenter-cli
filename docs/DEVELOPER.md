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


## Documentation
The following template should be used when documenting code.

- The First Line is Key: The initial summary line is the most important part. Tools like go doc and IDEs use it for quick reference lists. Make sure it's concise and accurately describes what the function does.
- Adapt and Omit: This template includes everything you might need. Don't force a section if it doesn't apply. If a function doesn't panic, has no parameters, or isn't deprecated, simply delete those sections. The goal is clear and relevant documentation, not just filling out a form.
- Runnable Examples: Go tooling can actually compile and run the code in your Example: blocks as tests. This ensures your documentation is always accurate and provides immense value to anyone using your code.

### Golang documentation temlate

```go
// [FunctionName] [starts with an active verb and provides a one-sentence summary].
//
// [This section provides a more detailed explanation. Describe the function's
// behavior, its purpose, the algorithm it uses, and any important context or
// side effects. Use full paragraphs for clarity. Explain *why* this function
// exists if it's not immediately obvious.]
//
// Parameters:
//   - [param1]: [Description of the first parameter, its purpose, expected
//     values, and behavior with nil or zero-values.]
//   - [param2]: [Description of the second parameter.]
//
// Returns:
//   - [returnValue1]: [Description of what this return value represents.]
//   - [error]: [Describe the exact conditions under which an error is returned.]
//
// Example:
//   [Provide a clear, runnable code snippet demonstrating a common use case.
//   This is one of the most valuable parts of documentation.]
//   err := [FunctionName]([param1_value], [param2_value])
//   if err != nil {
//       log.Printf("An error occurred: %v", err)
//   }
//
// Panics:
//   [Clearly state the exact conditions that will cause this function to panic.
//   If the function never panics, omit this section.]
//
// Deprecated: [If the function is deprecated, state the reason why and what
// function or method developers should use instead. This tag is formally
// recognized by Go tooling.]
//
// See Also:
//   - [RelatedFunction1]
//   - [AnotherPackage.RelevantFunction]
func [FunctionName](param1 Type1, param2 Type2) (returnValue1 Type3, err error) {
	// function implementation
}
```
