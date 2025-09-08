# Directory Structure

Below is an overview of the repository layout. The `openCenter-cli/` directory is the project root when unpacked.

```
openCenter-cli/
├─ cmd/                    # Cobra commands grouped by file
│  ├─ root.go              # Root command and CLI setup
│  ├─ prompt.go            # Command for shell prompt integration
│  ├─ cluster.go           # Cluster parent command
│  ├─ cluster_list.go      # Lists configured clusters
│  ├─ cluster_select.go    # Selects the active cluster
│  ├─ cluster_current.go   # Shows the current cluster
│  ├─ cluster_info.go      # Prints cluster YAML
│  ├─ cluster_init.go      # Initialises a new cluster configuration
│  ├─ cluster_validate.go  # Checks invariants across fields
│  ├─ cluster_preflight.go # Checks prerequisites like git and kubectl
│  ├─ cluster_setup.go     # Populates GitOps directory and initialises git
│  ├─ cluster_render.go    # Renders templates without git initialisation
│  ├─ cluster_bootstrap.go # Commits and pushes to remote git
│  └─ cluster_schema.go    # Exports Draft 2020‑12 JSON schema
├─ internal/
│  ├─ config/              # YAML model, schema, validation
│  │  ├─ config.go         # Config structs, load/save/list, validation
│  │  └─ schema.go         # JSON schema generation
│  ├─ gitops/              # Template embedding and rendering helpers
│  │  ├─ embed.go          # Uses go:embed to include templates
│  │  ├─ copy.go           # Copies or renders templates into gitops.git_dir
│  │  └─ templates/
│  │     ├─ README.md      # Explanation of rendering rules
│  │     └─ KUSTOMIZATION.tmpl # Example template for Kustomize
│  └─ cloud/openstack/
│     └─ preflight.go      # Provider‑specific preflight checks
├─ tests/
│  └─ features/               # Behaviour‑driven test specifications
│     ├─ cluster.feature      # Scenarios covering cluster commands
│     ├─ gitops.feature       # Scenarios for template rendering
│     └─ steps/
│        ├─ helpers.go        # Test scaffolding and binary invocation
│        └─ steps_test.go     # Godog step implementations
├─ docs/
│  ├─ USER_GUIDE.md        # End‑user instructions and workflows
│  ├─ DEVELOPER.md         # Guidance for contributors
│  ├─ ARCHITECTURE.md      # High‑level design and flows
│  └─ DIRECTORY_STRUCTURE.md # This file
├─ schema/                 # Generated JSON schemas (created at runtime)
├─ .mise.toml              # Task definitions for build, test and schema generation
├─ go.mod                  # Go module metadata
├─ main.go                 # Entrypoint delegating to the root command
└─ README.md               # Project overview and quick start
```