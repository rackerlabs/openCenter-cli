# CLI Command Reference

This document provides a detailed reference for every command available in the `openCenter` CLI.

## Global Flags

These flags are available on all `openCenter` commands.

| Flag | Description | Default | Environment Variable |
| --- | --- | --- | --- |
| `--config-dir <path>` | Specifies the directory where cluster configuration files are stored. | `~/.config/openCenter` | `OPENCENTER_CONFIG_DIR` |

### Environment Variables
- `OPENCENTER_CONFIG_DIR`: Overrides the configuration directory. Passing `--config-dir` sets this variable for the process.
- `OPENCENTER_PLUGINS_DIR`: Directory searched first for plugin executables before `$OPENCENTER_CONFIG_DIR/plugins` and `PATH`.

---

## `openCenter cluster`

The `cluster` command is the main entry point for managing cluster configurations. It serves as a parent command for many subcommands.

Running `openCenter cluster` by itself will display a help message listing all available subcommands.

### `openCenter cluster list`

(Alias: `ls`)

Lists all available cluster configurations found in the configuration directory.

**Usage**
```bash
openCenter cluster list [flags]
```

---

## `openCenter plugins`

Manage external plugins.

Plugins are executables named `openCenter-<name>` discovered at runtime and exposed as `openCenter <name>`. See `docs/reference/plugins.md` for discovery rules and examples.

### `openCenter plugins list`

Lists discovered plugins and their paths.

Usage
```bash
openCenter plugins list
```

Example
```bash
openCenter plugins list
# kubectl	/Users/alice/.config/openCenter/plugins/openCenter-kubectl
```

**Flags**
| Flag | Description |
| --- | --- |
| `--json` | Output the list of clusters in JSON format. |

**Example**
```bash
$ ./openCenter cluster list
dev
prod

$ ./openCenter cluster ls --json
[
  "dev",
  "prod"
]
```

---

### `openCenter cluster select [name]`

Selects a cluster to be the "active" context for other commands.

If `[name]` is provided, it sets that cluster as active. If no name is provided, it launches an interactive menu to choose from the list of available clusters.

**Usage**
```bash
openCenter cluster select [name]
```

**Example**
```bash
# Select a cluster by name
./openCenter cluster select dev

# Select a cluster interactively
./openCenter cluster select
```

---

### `openCenter cluster current`

Displays the name of the currently active cluster.

**Usage**
```bash
openCenter cluster current
```

**Example**
```bash
$ ./openCenter cluster current
dev
```

---

### `openCenter cluster info [name]`

Displays the full, parsed configuration for a cluster.

If `[name]` is provided, it shows the configuration for that cluster. If no name is provided, it shows the configuration for the active cluster.

**Usage**
```bash
openCenter cluster info [name] [flags]
```

**Flags**
| Flag | Description |
| --- | --- |
| `--json` | Output the full configuration in JSON format. |

**Example**
```bash
# Show info for the active cluster
./openCenter cluster info

# Show info for a specific cluster as JSON
./openCenter cluster info dev --json
```

---

### `openCenter cluster init [name]`

Initializes a new cluster configuration file with default values.

If `[name]` is provided, it creates `<name>.yaml`. If no name is provided, it launches an interactive guide to create the configuration.

**Usage**
```bash
openCenter cluster init [name] [flags]
```

**Flags**
| Flag | Description |
| --- | --- |
| `--force` | Overwrite the configuration file if it already exists. |
| `--strict`| Fail the command if the resulting configuration is not valid. |
| `--no-sops-keygen` | Do not auto-generate a SOPS age key when `secrets.sops_age_key_file` is unset. |

**Dynamic Flags**

The `init` command also accepts flags to override any field in the configuration using dot notation. This is useful for scripting.

**Example**
```bash
# Create a new cluster non-interactively, overriding several fields
./openCenter cluster init my-cluster \
  --gitops.git_dir=/tmp/my-cluster \
  --gitops.git_url=git@github.com:my-org/my-cluster.git \
  --iac.networking.use_octavia=false \
  --iac.networking.vrrp_ip=192.168.1.100

# Create a new cluster interactively
./openCenter cluster init

# Create without auto-generating a SOPS key
./openCenter cluster init my-cluster --no-sops-keygen
 
# During the interactive wizard you can:
# - Provide Git branch, Flux interval and prune
# - Provide Ansible inventory and playbooks
# - Provide a SOPS key path (or leave blank to auto-generate)
# - Choose to verify configuration before saving (shows validation results)
```

---

### `openCenter cluster validate [name]`

Validates the configuration for a given cluster against a set of rules.

**Usage**
```bash
openCenter cluster validate [name]
```

**Example**
```bash
./openCenter cluster validate my-cluster
```

---

### `openCenter cluster preflight [name]`

Runs a series of preflight checks to ensure the environment is ready for setup and bootstrapping. This includes checking for required tools (like `git`, `kubectl`) and provider-specific configurations.

**Usage**
```bash
openCenter cluster preflight [name]
```

---

### `openCenter cluster setup`

Generates the GitOps repository for the active cluster by materializing embedded templates into the `gitops.git_dir` path.

**Usage**
```bash
openCenter cluster setup [flags]
```

**Flags**
| Flag | Description |
| --- | --- |
| `--render` | Process `.tmpl` files using values from the cluster configuration. |
| `--force` | Overwrite the `git_dir` if it already exists. |

**Example**
```bash
# Set up the repo, processing templates
./openCenter cluster setup --render
```

---

### `openCenter cluster render`

Renders the templates for the active cluster without performing the full setup (e.g., does not initialize a git repository). This is useful for inspecting the output of the templating engine.

**Usage**
```bash
openCenter cluster render
```

---

### `openCenter cluster bootstrap`

Commits the contents of the `gitops.git_dir` and pushes them to the configured `gitops.git_url`.

**Usage**
```bash
openCenter cluster bootstrap
```

---

### `openCenter cluster schema`

Generates a JSON Schema (Draft 2020-12) for the cluster configuration. This schema can be used for IDE validation and autocompletion.

**Usage**
```bash
openCenter cluster schema [flags]
```

**Flags**
| Flag | Description |
| --- | --- |
| `--out <path>` | Path to write the schema file to. If omitted, prints to stdout. |
| `--pretty` | Indent the JSON output for readability. |

**Example**
```bash
./openCenter cluster schema --out schema/cluster.schema.json --pretty
```

### `openCenter cluster destroy`

Destroy a cluster.

**Usage**
```bash
openCenter cluster destroy
```

### Sources

*   `cmd/`
*   `tests/features/`
*   `README.md`

---

### `openCenter cluster update [name]`

Updates fields in an existing cluster configuration using dotted flags. If `[name]` is omitted, the active cluster is used.

This mirrors the dynamic flag behavior of `cluster init`, allowing you to script updates succinctly.

**Usage**
```bash
openCenter cluster update [name] [flags]
```

**Flags**
| Flag | Description |
| --- | --- |
| `--strict` | Fail the command if the resulting configuration is not valid. |

**Examples**
```bash
# Update a specific cluster
./openCenter cluster update test01 --iac.counts.master=3

# Update the active cluster
./openCenter cluster update --gitops.git_branch=main --ansible.enabled=false
```

---

## `openCenter secrets`

The `secrets` command group provides helpers for working with secret management tools used alongside GitOps.

Running `openCenter secrets` by itself will display a help message listing available subcommands.

### `openCenter secrets sops-keygen`

Generates a SOPS (age) secret key file suitable for use with SOPS. The file is written with `0600` permissions and contains a key string starting with `AGE-SECRET-KEY-1`.

Note: This helper generates a placeholder key compatible with SOPS. You may replace it with a key generated by the `age-keygen` tool at any time.

**Usage**
```bash
openCenter secrets sops-keygen --out <path>
```

**Flags**
| Flag | Description |
| --- | --- |
| `--out <path>` | Path to write the age key file (required). |

**Example**
```bash
./openCenter secrets sops-keygen --out ~/.config/sops/age/keys.txt

# Configure your cluster to point at the key
# secrets:
#   sops_age_key_file: ~/.config/sops/age/keys.txt
```
