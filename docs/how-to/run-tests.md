# How-To: Run the BDD Test Suite

`openCenter` uses Behavior-Driven Development (BDD) to ensure the CLI behaves as expected. The tests are written in Gherkin (`.feature` files) and executed with Godog. These tests serve as both a safety net against regressions and as living documentation for the tool's functionality.

This guide explains how to run the test suite.

## Who is this for?

*   **Contributors** who want to validate their changes before submitting a pull request.
*   **Developers** who want to understand the expected behavior of a specific command.

## What you'll achieve

*   Run the entire BDD test suite.
*   Run specific feature files or scenarios for targeted testing.
*   Understand the basic structure of the tests.

---

### Prerequisites

*   You have cloned the `openCenter` repository.
*   You have [Mise](https://mise.jdx.dev/) installed.
*   You have installed the project's development tools:

    ```bash
    mise install
    ```

### Running the Test Suite

The primary way to run the BDD tests is by using the `godog` task defined in the project's `.mise.toml` file.

```bash
# This is the primary command for running all regression tests.
mise run godog
```

This task executes the Godog test suite via `go test`. By default, it runs all `.feature` files in the `tests/features/` directory.

### Running Specific Features

The `godog` task can accept arguments passed to the underlying `go test` command. To run a specific feature, you can use the `--godog.paths` flag.

```bash
# Run only the tests in the workflow.feature file
mise run godog -- --godog.paths=tests/features/workflow.feature
```

Note the `--` which separates the `mise` command from the arguments being passed to the task.

There is also a `test` task, which runs unit tests for the internal packages:
```bash
mise run test
```

### Understanding the Test Structure

The tests are organized in the `tests/features/` directory:

*   **`.feature` files**: These files contain the human-readable test scenarios written in Gherkin syntax (`Given`, `When`, `Then`). Each file typically corresponds to a specific feature of the CLI (e.g., `cluster_init.feature`, `gitops.feature`).
*   **`steps/` directory**: This directory (`tests/features/steps/`) contains the Go code that implements the steps defined in the feature files. For example, when a test says `When I run "openCenter cluster list"`, the code in the `steps/` directory is what actually executes that command and captures its output.

By reading the `.feature` files, you can get a clear understanding of what `openCenter` is designed to do and how it should behave in various situations.
