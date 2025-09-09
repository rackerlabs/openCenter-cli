# Configuration File Reference

The `openCenter` cluster configuration is a single YAML file that serves as the declarative source of truth for a cluster's layout, from its GitOps repository to its cloud provider details.

This document provides a detailed reference for every available field in the configuration file.

For a better editing experience, we highly recommend setting up schema validation in your IDE. See our guide on [How to Configure Your IDE](./../how-to/configure-ide.md).

## Top-Level Fields

These are the main keys at the root of the configuration file.

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `cluster_name` | `string` | Yes | The unique name for the cluster. This is used as the filename. |
| `naming_prefix`| `string` | No | An optional prefix for all named resources. |
| `cluster` | `object` | No | High-level metadata: `name`, `env`, `region`, `status`. |
| `gitops` | `object` | Yes | Configuration for the GitOps repository. |
| `terraform` | `object` | No | Settings for Terraform integration. |
| `ansible` | `object` | No | Settings for Ansible integration. |
| `iac` | `object` | Yes | Infrastructure-as-code and cluster layout settings. |
| `cloud` | `object` | Yes | All cloud provider-specific settings. |
| `secrets` | `object` | No | Secret management settings. |

---

## `gitops`

| Field | Type | Default | Description |
| --- | --- | --- | --- |
| `git_dir` | `string` | `""` | **Required.** The absolute path on your local machine where the GitOps repository will be generated. |
| `git_url` | `string` | `""` | **Required.** The SSH URL of the remote Git repository where the configuration will be pushed. |
| `git_ssh_key`| `string` | `""` | Optional. Path to a specific SSH private key to use for pushing to the remote repository. |
| `git_branch`| `string` | `""` | Optional. Branch to push to (defaults to `main` if unset in bootstrap). |
| `flux.interval`| `string` | `""` | Optional. Reconciliation interval (e.g., `1m`). |
| `flux.prune`| `boolean` | `false` | Optional. Enable pruning in Flux. |

---

## `terraform`

| Field | Type | Default | Description |
| --- | --- | --- | --- |
| `enabled` | `boolean` | `true` | Whether to use Terraform for provisioning. |
| `path` | `string` | `terraform` | The subdirectory within the `git_dir` for Terraform files. |

---

## `ansible`

| Field | Type | Default | Description |
| --- | --- | --- | --- |
| `enabled` | `boolean` | `true` | Whether to use Ansible for provisioning. |
| `path` | `string` | `ansible` | The subdirectory within the `git_dir` for Ansible files. |
| `inventory` | `string` | `""` | Optional. Path or filename for inventory. |
| `playbooks` | `array` | `[]` | Optional. List of playbooks to include. |

---

## `iac`

This section contains all settings related to the Kubernetes cluster itself.

| Field | Type | Default | Description |
| --- | --- | --- | --- |
| `ssh_user` | `string` | `ubuntu` | The SSH user for connecting to the cluster nodes. |
| `k8s_api_port`| `integer` | `443` | The port for the Kubernetes API server. |
| `ub_version`| `string` | `20` | The version of Ubuntu to use for the nodes. |
| `ssh_authorized_keys`| `array` | `[]` | A list of public SSH keys to add to the nodes. |
| `ca_certificates`| `string` | `""` | CA certificates to trust on the nodes. |
| `node_roles`| `map` | See below | A map of role names to their purposes (e.g., `master`, `worker`). |
| `counts` | `map` | See below | A map of node roles to the number of nodes of that type. |
| `images` | `map` | See below | A map of OS types (`linux`, `windows`) to the cloud image to use. |
| `flavors` | `map` | `{}` | A map of node roles to the cloud flavor to use. |
| `storage` | `object` | See below | Configuration for node block storage. |
| `windows` | `object` | See below | Windows-specific node settings. |
| `networking`| `object` | See below | All network-related settings. |

**Default `node_roles`**:
```yaml
node_roles:
  master: master
  worker: worker
  windows: win_wn
```

**Default `counts`**:
```yaml
counts:
  master: 0
  worker: 0
  worker_windows: 0
```
---

### `iac.storage`

| Field | Type | Default | Description |
| --- | --- | --- | --- |
| `master_node_bfv` | `object` | `{size: 100, type: "local"}` | Boot-from-volume settings for master nodes. |
| `worker_node_bfv` | `object` | `{size: 100, type: "local"}` | Boot-from-volume settings for worker nodes. |
| `worker_node_bfv_windows`| `object` | `{size: 0, type: "local"}` | Boot-from-volume settings for Windows worker nodes. |

---

### `iac.networking`

| Field | Type | Default | Description |
| --- | --- | --- | --- |
| `subnet_nodes` | `string` | `10.0.0.0/16` | The CIDR for the node subnet. |
| `allocation_pool_start`| `string` | `""` | The start of the IP allocation pool for the node subnet. |
| `allocation_pool_end` | `string` | `""` | The end of the IP allocation pool for the node subnet. |
| `vrrp_enabled` | `boolean` | `false` | Enable VRRP for control plane load balancing. Incompatible with `use_octavia`. |
| `vrrp_ip` | `string` | `""` | The virtual IP to use for VRRP. Required if `vrrp_enabled` is true or `use_octavia` is false. |
| `subnet_services` | `string` | `10.43.0.0/16` | The CIDR for the Kubernetes services subnet. |
| `subnet_pods` | `string` | `10.42.0.0/16` | The CIDR for the Kubernetes pods subnet. |
| `use_octavia` | `boolean` | `true` | Use OpenStack Octavia for the load balancer. |
| `loadbalancer_provider`| `string` | `amphora` | The Octavia load balancer provider. |
| `use_designate` | `boolean` | `true` | Use OpenStack Designate for DNS. |
| `dns_zone_name` | `string` | `""` | The DNS zone name to use with Designate. Required if `use_designate` is true. |
| `dns_nameservers` | `array` | `["8.8.8.8", "8.8.4.4"]` | A list of DNS nameservers for the cluster. |
| `vlan` | `object` | See below | VLAN settings for the cluster. |

---

## `cloud` and `cloud.openstack`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `provider` | `string` | Yes | The cloud provider to use. Default is `openstack`. |
| `openstack.auth_url` | `string` | Yes | The Keystone authentication URL. |
| `openstack.insecure` | `boolean`| No | If true, allows insecure TLS connections. Default `false`. |
| `openstack.region` | `string` | Yes | The OpenStack region. |
| `openstack.user_name` | `string` | Yes | The OpenStack username. |
| `openstack.user_password` | `string` | Yes | The OpenStack user password. |
| `openstack.project_domain_name`| `string` | Yes | The OpenStack project domain name. |
| `openstack.user_domain_name` | `string` | Yes | The OpenStack user domain name. |
| `openstack.tenant_name` | `string` | Yes | The OpenStack tenant/project name. |
| `openstack.availability_zone`| `string` | Yes | The availability zone for resources. |
| `openstack.floatingip_pool` | `string` | Yes | The floating IP pool to use. |
| `openstack.router_external_network_id` | `string` | Yes | The ID of the external network for the router. |
| `openstack.disable_bastion` | `boolean`| No | If true, do not create a bastion host. Default `false`. |
| `openstack.ca` | `string` | No | Path to a custom CA certificate for the OpenStack endpoint. |
| `openstack.external_network` | `string` | No | Name of the external network. |
| `openstack.use_octavia` | `boolean` | No | Convenience flag for Octavia usage. |
| `openstack.vrrp_ip` | `string` | No | Convenience VRRP IP when not using Octavia. |

## `cloud.aws`

| Field | Type | Required | Description |
| --- | --- | --- | --- |
| `aws.profile` | `string` | No | AWS profile name. |
| `aws.region` | `string` | No | AWS region. |
| `aws.vpc_id` | `string` | No | Target VPC ID. |
| `aws.private_subnets` | `array` | No | Private subnet IDs. |
| `aws.public_subnets` | `array` | No | Public subnet IDs. |

---

## Validation Rules

The `openCenter cluster validate` command enforces these rules:

*   `gitops.git_dir` must be set.
*   If `iac.networking.use_octavia` is `true`, then `vrrp_enabled` must be `false`.
*   If `iac.networking.use_octavia` is `false`, then `vrrp_ip` must be set.
*   If `iac.networking.vrrp_enabled` is `true`, then `vrrp_ip` must be set.
*   If `iac.networking.use_designate` is `true`, then `dns_zone_name` must be set.
*   If a node `count` for a role is greater than 0, a corresponding `flavor` for that role must be set.

## Minimal Example

```yaml
cluster_name: demo
gitops:
  git_dir: /tmp/opencenter-demo
  git_url: git@github.com:example/demo.git
  git_branch: main
  flux:
    interval: 1m
    prune: true
ansible:
  enabled: true
  path: ansible
  inventory: inventory.yaml
  playbooks:
    - hardening.yaml
    - node-setup.yaml
iac:
  engine: terraform
  stack: dev/demo
  counts:
    master: 3
    worker: 5
  flavors:
    master: "c2.medium"
    worker: "m1.large"
  networking:
    use_octavia: false
    vrrp_ip: 10.0.0.10
cloud:
  provider: openstack
  openstack:
    auth_url: https://keystone.example.com/v3
    region: RegionOne
    user_name: "my-user"
    user_password: "my-password"
    project_domain_name: "Default"
    user_domain_name: "Default"
    tenant_name: "my-project"
    availability_zone: "nova"
    floatingip_pool: "public"
    router_external_network_id: "abc-123-def-456"
secrets:
  sops_age_key_file: ~/.config/sops/age/keys.txt
```

## `secrets`

Defines paths and settings related to secret management.

| Field | Type | Default | Description |
| --- | --- | --- | --- |
| `sops_age_key_file` | `string` | `""` | Path to the SOPS age secret key file used for encryption/decryption. |

Notes
- If `sops_age_key_file` is not set at init time, `openCenter` automatically generates a key at `~/.config/openCenter/sops/age/keys/<cluster-name>-key.txt` and updates the saved config accordingly.
- The generated file is written with permissions `0600` and contains a key string starting with `AGE-SECRET-KEY-1`.
 - To disable auto-generation during init, pass `--no-sops-keygen` to `openCenter cluster init`.

### Sources

*   `internal/config/config.go`
*   `internal/config/schema.go`
*   `README.md`
