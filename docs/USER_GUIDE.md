# User Guide

This guide explains how to use the openCenter CLI to create and manage cluster configurations and GitOps scaffolding. It assumes that you have built the binary (see the README) and have it available as `openCenter` in your PATH or current directory.

## Configuration Directory

By default, openCenter stores configuration files in the operating system’s user configuration directory:

- **Linux/macOS:** `~/.config/openCenter`
- **Windows:** `%APPDATA%\openCenter`

You can override the location with either the `--config-dir` global flag or by setting the `OPENCENTER_CONFIG_DIR` environment variable. The directory is created if it does not exist. An `.active` file in this directory tracks the currently selected cluster.

## Initialising a Cluster

To create a new cluster configuration named `demo`:

```sh
openCenter cluster init demo
```

If you omit the name, openCenter will ask for it and optionally prompt for additional values such as the GitOps directory and Git repository URL. The created YAML file will be located at `${config-dir}/demo.yaml` and contains sensible defaults. Use a text editor to fine‑tune fields like counts, flavors, images, networking and cloud settings.

## Selecting and Inspecting Clusters

List available clusters with:

```sh
openCenter cluster list
```

To mark a cluster as active:

```sh
openCenter cluster select demo
```

If you omit the name and more than one cluster exists, openCenter will present an interactive selection menu. The active cluster persists between commands.

View the current active cluster:

```sh
openCenter cluster current
# or quietly
openCenter cluster current -q
```

Display the full YAML for the active cluster (or another cluster):

```sh
openCenter cluster info           # uses active
openCenter cluster info blue      # override
```

## Validating a Cluster

The `validate` subcommand checks invariants across fields. It returns a non‑zero exit code if any rule is broken:

```sh
openCenter cluster validate

# Example messages:
# kubernetes.networking.use_octavia=true and vrrp_enabled=true are mutually exclusive
```

Validation rules include:

- If `use_octavia` is true, `vrrp_enabled` must be false.
- If `use_octavia` is false, `vrrp_ip` must be set.
- If `vrrp_enabled` is true, `vrrp_ip` must be set.
- If `use_designate` is true, `dns_zone_name` must be set.
- If any of the `counts.master`, `counts.worker`, or `counts.worker_windows` values are > 0, the corresponding `flavors` entry must be set.
- `gitops.git_dir` and `cluster_name` must not be empty.

## Preflight Checks

Run `preflight` to ensure your environment has required tools and provider hints:

```sh
openCenter cluster preflight
```

The command checks for the presence of `git`, `kubectl` and `talosctl`. For OpenStack providers it also checks for the `openstack` CLI and warns if `cloud.openstack.auth_url` is empty.

## Setting up GitOps

The `setup` command copies or renders embedded templates into the directory specified by `gitops.git_dir` and initialises a Git repository if one does not already exist. It also writes a `.opencenter` marker file containing the cluster name.

```sh
openCenter cluster setup --render
```

Use `--render` to process files ending in `.tmpl` via Go’s `text/template` engine with [Sprig](https://masterminds.github.io/sprig/) functions available. Without `--render`, files are copied verbatim. After copying, the command runs `git init -b main`, stages all files and creates an initial commit.

To render templates only without git initialisation, use `render`:

```sh
openCenter cluster render demo
```

## Bootstrapping to a Remote

Once your configuration and templates are in place, set `gitops.git_url` in your cluster YAML to point to a remote Git repository. Then run:

```sh
openCenter cluster bootstrap
```

The command stages and commits any changes, sets or updates the `origin` remote to `gitops.git_url` and pushes the `main` branch. Credentials and SSH keys must be configured externally.

## Exporting the JSON Schema

Use the `schema` subcommand to generate a Draft 2020‑12 JSON schema describing the configuration model. This schema can be used by IDEs for validation and autocompletion:

```sh
openCenter cluster schema --out schema/cluster.schema.json --pretty
```

If `--out` is omitted, the schema is printed to stdout. Passing `--pretty` formats the JSON with indentation.

## Destroying a Cluster

The `destroy` command removes the cluster's configuration file and its associated GitOps directory.

```sh
openCenter cluster destroy demo
```

This command is irreversible and will permanently delete the specified cluster's configuration and GitOps repository from your local machine.

## Integrating with Your Shell Prompt

The `openCenter prompt` command prints `(openCenter:&lt;active&gt;)` only if the current working directory matches `gitops.git_dir` for the active cluster. This is useful for including in your shell’s `PS1` to show which cluster repository you are working in. For example, in Bash:

```sh
PS1='$(openCenter prompt) \u@\h \W \$ '
```