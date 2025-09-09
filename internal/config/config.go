// Copyright 2025 Victor Palma <victor.palma@rackspace.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package config

import (
    "encoding/json"
    "errors"
    "fmt"
    "io/fs"
    "os"
    "path/filepath"
    "runtime"
    "strings"

    "gopkg.in/yaml.v3"
)

// Config represents the root configuration for a cluster. It groups
// GitOps, Kubernetes, and Cloud configuration under nested fields.
//
// When marshalled to YAML, the field names follow the dotted notation as
// defined in the specification.
//
// Some fields have default values; these defaults are applied when
// loading or initializing a configuration. Required fields are
// validated by the Validate function.
//
// See docs/ARCHITECTURE.md for a complete description of the data model.
type Config struct {
    ClusterName   string       `yaml:"cluster_name" json:"cluster_name"`
    NamingPrefix  string       `yaml:"naming_prefix,omitempty" json:"naming_prefix,omitempty"`
    Cluster       ClusterMeta  `yaml:"cluster,omitempty" json:"cluster,omitempty"`
    GitOps        GitOpsConfig `yaml:"gitops" json:"gitops"`
    Terraform     Terraform    `yaml:"terraform" json:"terraform"`
    Ansible       Ansible      `yaml:"ansible" json:"ansible"`
    IAC           IAC          `yaml:"iac" json:"iac"`
    Cloud         Cloud        `yaml:"cloud" json:"cloud"`
    Secrets       Secrets      `yaml:"secrets" json:"secrets"`
}

// ClusterMeta holds high-level metadata about the cluster.
type ClusterMeta struct {
    Name   string `yaml:"name" json:"name"`
    Env    string `yaml:"env" json:"env"`
    Region string `yaml:"region" json:"region"`
    Status string `yaml:"status" json:"status"`
}

// Terraform holds Terraform-specific settings.
type Terraform struct {
    Enabled bool   `yaml:"enabled" json:"enabled"`
    Path    string `yaml:"path" json:"path"`
}

// Ansible holds Ansible-specific settings.
type Ansible struct {
    Enabled   bool     `yaml:"enabled" json:"enabled"`
    Path      string   `yaml:"path" json:"path"`
    Inventory string   `yaml:"inventory,omitempty" json:"inventory,omitempty"`
    Playbooks []string `yaml:"playbooks,omitempty" json:"playbooks,omitempty"`
}

// GitOpsConfig holds configuration related to GitOps scaffolding and repositories.
type GitOpsConfig struct {
    GitDir     string     `yaml:"git_dir" json:"git_dir"`
    GitURL     string     `yaml:"git_url" json:"git_url"`
    GitSSHKey  string     `yaml:"git_ssh_key,omitempty" json:"git_ssh_key,omitempty"`
    GitBranch  string     `yaml:"git_branch,omitempty" json:"git_branch,omitempty"`
    Flux       GitOpsFlux `yaml:"flux,omitempty" json:"flux,omitempty"`
}

// GitOpsFlux holds optional FluxCD settings for reconciliation behavior.
type GitOpsFlux struct {
    Interval string `yaml:"interval" json:"interval"`
    Prune    bool   `yaml:"prune" json:"prune"`
}

// Kubernetes groups settings for the Kubernetes cluster.
// It nests further objects for counts, images, flavors, and networking.
// Default values are applied at load time.
// IAC groups settings for infrastructure-as-code driven cluster provisioning.
// It retains the detailed node/layout fields and adds engine/stack selectors.
type IAC struct {
    // New fields for engine selection
    Engine string `yaml:"engine" json:"engine"`
    Stack  string `yaml:"stack" json:"stack"`
    // Optional high-level k8s selectors
    Version string `yaml:"version" json:"version"`
    CNI     string `yaml:"cni" json:"cni"`
    Ingress string `yaml:"ingress" json:"ingress"`
    // Detailed fields retained from previous Kubernetes struct
    SSHUser           string            `yaml:"ssh_user" json:"ssh_user"`
    K8sAPIPort        int               `yaml:"k8s_api_port" json:"k8s_api_port"`
    UBVersion         string            `yaml:"ub_version" json:"ub_version"`
    SSHAuthorizedKeys []string          `yaml:"ssh_authorized_keys" json:"ssh_authorized_keys"`
    CACertificates    string            `yaml:"ca_certificates" json:"ca_certificates"`
    NodeRoles         map[string]string `yaml:"node_roles" json:"node_roles"`
    Counts            map[string]int    `yaml:"counts" json:"counts"`
    Images            map[string]string `yaml:"images" json:"images"`
    Flavors           map[string]string `yaml:"flavors" json:"flavors"`
    Storage           Storage           `yaml:"storage" json:"storage"`
    Windows           Windows           `yaml:"windows" json:"windows"`
    Networking        Networking        `yaml:"networking" json:"networking"`
}

// Storage describes block volume storage sizes and types for various node roles.
type Storage struct {
    MasterNodeBFV        BFV `yaml:"master_node_bfv" json:"master_node_bfv"`
    WorkerNodeBFV        BFV `yaml:"worker_node_bfv" json:"worker_node_bfv"`
    WorkerNodeBFVWindows BFV `yaml:"worker_node_bfv_windows" json:"worker_node_bfv_windows"`
}

// BFV describes the block volume size and type.
type BFV struct {
    Size int    `yaml:"size" json:"size"`
    Type string `yaml:"type" json:"type"`
}

// Windows holds Windows-specific settings.
type Windows struct {
    User         string `yaml:"user" json:"user"`
    AdminPassword string `yaml:"admin_password" json:"admin_password"`
}

// Networking groups network settings and options around VRRP and service networks.
type Networking struct {
    SubnetNodes           string   `yaml:"subnet_nodes" json:"subnet_nodes"`
    AllocationPoolStart   string   `yaml:"allocation_pool_start" json:"allocation_pool_start"`
    AllocationPoolEnd     string   `yaml:"allocation_pool_end" json:"allocation_pool_end"`
    VRRPEnabled           bool     `yaml:"vrrp_enabled" json:"vrrp_enabled"`
    VRRPIP                string   `yaml:"vrrp_ip" json:"vrrp_ip"`
    SubnetServices        string   `yaml:"subnet_services" json:"subnet_services"`
    SubnetPods            string   `yaml:"subnet_pods" json:"subnet_pods"`
    UseOctavia            bool     `yaml:"use_octavia" json:"use_octavia"`
    LoadbalancerProvider  string   `yaml:"loadbalancer_provider" json:"loadbalancer_provider"`
    UseDesignate          bool     `yaml:"use_designate" json:"use_designate"`
    DNSZoneName           string   `yaml:"dns_zone_name" json:"dns_zone_name"`
    DNSNameservers        []string `yaml:"dns_nameservers" json:"dns_nameservers"`
    VLAN                  VLAN     `yaml:"vlan" json:"vlan"`
}

// VLAN describes VLAN settings for the cluster.
type VLAN struct {
    ID       string `yaml:"id" json:"id"`
    MTU      int    `yaml:"mtu" json:"mtu"`
    Provider string `yaml:"provider" json:"provider"`
}

// Cloud holds provider-specific configuration. Currently, only OpenStack is supported.
type Cloud struct {
    Provider  string         `yaml:"provider" json:"provider"`
    OpenStack OpenStackCloud `yaml:"openstack" json:"openstack"`
    AWS       AWSCloud       `yaml:"aws" json:"aws"`
}

// OpenStackCloud contains options for connecting to an OpenStack deployment.
type OpenStackCloud struct {
    AuthURL            string `yaml:"auth_url" json:"auth_url"`
    Insecure           bool   `yaml:"insecure" json:"insecure"`
    Region             string `yaml:"region" json:"region"`
    UserName           string `yaml:"user_name" json:"user_name"`
    UserPassword       string `yaml:"user_password" json:"user_password"`
    AdminPassword      string `yaml:"admin_password" json:"admin_password"`
    ProjectDomainName  string `yaml:"project_domain_name" json:"project_domain_name"`
    UserDomainName     string `yaml:"user_domain_name" json:"user_domain_name"`
    TenantName         string `yaml:"tenant_name" json:"tenant_name"`
    AvailabilityZone   string `yaml:"availability_zone" json:"availability_zone"`
    FloatingIPPool     string `yaml:"floatingip_pool" json:"floatingip_pool"`
    RouterExternalNetworkID string `yaml:"router_external_network_id" json:"router_external_network_id"`
    DisableBastion     bool   `yaml:"disable_bastion" json:"disable_bastion"`
    CA                 string `yaml:"ca" json:"ca"`
    ExternalNetwork    string `yaml:"external_network" json:"external_network"`
    UseOctavia         bool   `yaml:"use_octavia" json:"use_octavia"`
    VRRPIP             string `yaml:"vrrp_ip" json:"vrrp_ip"`
}

// AWSCloud contains options for connecting to AWS environments.
type AWSCloud struct {
    Profile        string   `yaml:"profile" json:"profile"`
    Region         string   `yaml:"region" json:"region"`
    VPCID          string   `yaml:"vpc_id" json:"vpc_id"`
    PrivateSubnets []string `yaml:"private_subnets" json:"private_subnets"`
    PublicSubnets  []string `yaml:"public_subnets" json:"public_subnets"`
}

// Secrets holds paths or settings for secret management tools.
type Secrets struct {
    SopsAgeKeyFile string `yaml:"sops_age_key_file" json:"sops_age_key_file"`
}

// defaultConfig returns a Config pre-populated with the default
// values defined in the specification. This function can be used to
// initialise new cluster configurations.
func defaultConfig(name string) Config {
    cfg := Config{
        ClusterName:  name,
        NamingPrefix: "",
        Cluster: ClusterMeta{},
        GitOps: GitOpsConfig{
            GitDir: "",
            GitURL: "",
            GitBranch: "",
            Flux: GitOpsFlux{},
        },
        Terraform: Terraform{
            Enabled: true,
            Path:    "terraform",
        },
        Ansible: Ansible{
            Enabled: true,
            Path:    "ansible",
            Inventory: "",
            Playbooks: []string{},
        },
        IAC: IAC{
            Engine:            "",
            Stack:             "",
            Version:           "",
            CNI:               "",
            Ingress:           "",
            SSHUser:           "ubuntu",
            K8sAPIPort:        443,
            UBVersion:         "20",
            SSHAuthorizedKeys: []string{},
            CACertificates:    "",
            NodeRoles:         map[string]string{"master": "master", "worker": "worker", "windows": "win_wn"},
            Counts:            map[string]int{"master": 0, "worker": 0, "worker_windows": 0},
            Images:            map[string]string{"linux": "", "windows": ""},
            Flavors:           map[string]string{},
            Storage: Storage{
                MasterNodeBFV:        BFV{Size: 100, Type: "local"},
                WorkerNodeBFV:        BFV{Size: 100, Type: "local"},
                WorkerNodeBFVWindows: BFV{Size: 0, Type: "local"},
            },
            Windows: Windows{
                User:          "Administrator",
                AdminPassword: "",
            },
            Networking: Networking{
                SubnetNodes:          "10.0.0.0/16",
                AllocationPoolStart:  "",
                AllocationPoolEnd:    "",
                VRRPEnabled:          false,
                VRRPIP:               "",
                SubnetServices:       "10.43.0.0/16",
                SubnetPods:           "10.42.0.0/16",
                UseOctavia:           true,
                LoadbalancerProvider: "amphora",
                UseDesignate:         true,
                DNSZoneName:          "",
                DNSNameservers:       []string{"8.8.8.8", "8.8.4.4"},
                VLAN: VLAN{
                    ID:       "",
                    MTU:      1440,
                    Provider: "",
                },
            },
        },
        Cloud: Cloud{
            Provider: "openstack",
            OpenStack: OpenStackCloud{
                AuthURL:          "",
                Insecure:         false,
                Region:           "",
                UserName:         "",
                UserPassword:     "",
                AdminPassword:    "",
                ProjectDomainName: "",
                UserDomainName:    "",
                TenantName:        "",
                AvailabilityZone:  "",
                FloatingIPPool:    "",
                RouterExternalNetworkID: "",
                DisableBastion:    false,
                CA:               "",
                ExternalNetwork:   "",
                UseOctavia:        false,
                VRRPIP:            "",
            },
            AWS: AWSCloud{},
        },
        Secrets: Secrets{},
    }
    return cfg
}

// NewDefault returns a Config initialized with the default values for the given cluster name.
//
// Inputs:
//   - name: The name of the cluster.
//
// Outputs:
//   - Config: A new Config object with default values.
func NewDefault(name string) Config {
    return defaultConfig(name)
}

// ResolveConfigDir resolves the configuration directory based on the OPENCENTER_CONFIG_DIR
// environment variable. If the variable is not set, it falls back to the user's
// standard config directory (e.g., ~/.config/openCenter on Linux).
// The directory is created if it does not exist.
//
// Outputs:
//   - string: The absolute path to the configuration directory.
//   - error: An error if one occurred.
func ResolveConfigDir() (string, error) {
	var err error
	dir := os.Getenv("OPENCENTER_CONFIG_DIR")
	if dir == "" {
		// Determine OS-specific config directory
		switch runtime.GOOS {
		case "windows":
			base := os.Getenv("APPDATA")
			if base == "" {
				base = os.Getenv("LOCALAPPDATA")
			}
			if base == "" {
				base = os.Getenv("USERPROFILE")
			}
			dir = filepath.Join(base, "openCenter")
		default:
			home, herr := os.UserHomeDir()
			if herr != nil {
				err = herr
				return "", err
			}
			dir = filepath.Join(home, ".config", "openCenter")
		}
	}
	// Ensure absolute path
	if !filepath.IsAbs(dir) {
		dir, err = filepath.Abs(dir)
		if err != nil {
			return "", err
		}
	}
	// Create directory if not exists
	if mkErr := os.MkdirAll(dir, 0o755); mkErr != nil {
		err = mkErr
		return "", err
	}
	return dir, err
}

// ConfigPath returns the absolute path to a cluster's configuration file.
//
// Inputs:
//   - name: The name of the cluster (without the .yaml extension).
//
// Outputs:
//   - string: The absolute path to the configuration file.
//   - error: An error if one occurred.
func ConfigPath(name string) (string, error) {
    dir, err := ResolveConfigDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(dir, name+".yaml"), nil
}

// Load reads and unmarshals a YAML configuration file for the given cluster name.
// Default values are applied for any omitted fields.
//
// Inputs:
//   - name: The name of the cluster.
//
// Outputs:
//   - Config: The loaded configuration.
//   - error: An error if the file does not exist or cannot be parsed.
func Load(name string) (Config, error) {
    path, err := ConfigPath(name)
    if err != nil {
        return Config{}, err
    }
    data, readErr := os.ReadFile(path)
    if readErr != nil {
        return Config{}, fmt.Errorf("failed to read %s: %w", path, readErr)
    }
    // Unmarshal YAML then overlay onto default config
    cfg := defaultConfig(name)
    if unmarshalErr := yaml.Unmarshal(data, &cfg); unmarshalErr != nil {
        return Config{}, fmt.Errorf("failed to parse YAML: %w", unmarshalErr)
    }
    return cfg, nil
}

// Save writes the configuration to a YAML file. The file is saved with 0600
// permissions to protect sensitive data.
//
// Inputs:
//   - cfg: The configuration to save.
//
// Outputs:
//   - error: An error if the configuration cannot be saved.
func Save(cfg Config) error {
    if cfg.ClusterName == "" {
        return errors.New("cluster_name must not be empty")
    }
    path, err := ConfigPath(cfg.ClusterName)
    if err != nil {
        return err
    }
    // Marshal YAML with indentation
    data, marshalErr := yaml.Marshal(&cfg)
    if marshalErr != nil {
        return marshalErr
    }
    // Write with 0600 permissions
    if writeErr := os.WriteFile(path, data, 0o600); writeErr != nil {
        return writeErr
    }
    return nil
}

// List returns a sorted list of cluster names from the configuration directory.
// It ignores any files that do not have a .yaml extension.
//
// Outputs:
//   - []string: A list of cluster names.
//   - error: An error if the directory cannot be read.
func List() ([]string, error) {
    dir, err := ResolveConfigDir()
    if err != nil {
        return nil, err
    }
    entries, readErr := os.ReadDir(dir)
    if readErr != nil {
        return nil, readErr
    }
    var names []string
    for _, entry := range entries {
        if entry.IsDir() {
            continue
        }
        name := entry.Name()
        if strings.HasSuffix(name, ".yaml") {
            names = append(names, strings.TrimSuffix(name, ".yaml"))
        }
    }
    // Sort lexically
    if len(names) > 1 {
        sortStrings(names)
    }
    return names, nil
}

// simple string sorter to avoid pulling in a larger dependency.
func sortStrings(s []string) {
    for i := 0; i < len(s); i++ {
        for j := i + 1; j < len(s); j++ {
            if s[j] < s[i] {
                s[i], s[j] = s[j], s[i]
            }
        }
    }
}

// activeClusterPath returns the absolute path to the file tracking
// the active cluster. This file stores the cluster name as plain
// text.
func activeClusterPath() (string, error) {
    dir, err := ResolveConfigDir()
    if err != nil {
        return "", err
    }
    return filepath.Join(dir, ".active"), nil
}

// SetActive writes the given cluster name into the active marker file.
// If the name is empty, the marker file is removed.
//
// Inputs:
//   - name: The name of the cluster to set as active.
//
// Outputs:
//   - error: An error if the file cannot be written.
func SetActive(name string) error {
    path, err := activeClusterPath()
    if err != nil {
        return err
    }
    if name == "" {
        return os.Remove(path)
    }
    return os.WriteFile(path, []byte(name), 0o600)
}

// GetActive reads the active cluster name from the marker file.
// If the file does not exist or is empty, it returns an empty string.
//
// Outputs:
//   - string: The active cluster name.
//   - error: An error if the file cannot be read.
func GetActive() (string, error) {
    path, err := activeClusterPath()
    if err != nil {
        return "", err
    }
    data, readErr := os.ReadFile(path)
    if readErr != nil {
        if errors.Is(readErr, fs.ErrNotExist) {
            return "", nil
        }
        return "", readErr
    }
    return strings.TrimSpace(string(data)), nil
}

// Validate performs a set of invariant checks on the configuration.
//
// Inputs:
//   - cfg: The configuration to validate.
//
// Outputs:
//   - []string: A list of error messages describing any validation failures.
//     If the list is empty, the configuration is valid.
func Validate(cfg Config) []string {
    var errs []string
    // Required cluster name and gitops.git_dir
    if cfg.ClusterName == "" {
        errs = append(errs, "cluster_name must be set")
    }
    if cfg.GitOps.GitDir == "" {
        errs = append(errs, "gitops.git_dir must be set")
    }
    n := cfg.IAC.Networking
    // If use_octavia is true then vrrp_enabled must be false
    if n.UseOctavia && n.VRRPEnabled {
        errs = append(errs, "iac.networking.use_octavia=true and vrrp_enabled=true are mutually exclusive")
    }
    // If use_octavia is false, vrrp_ip must be set
    if !n.UseOctavia {
        if n.VRRPIP == "" {
            errs = append(errs, "iac.networking.use_octavia=false requires vrrp_ip to be set")
        }
    }
    // If vrrp_enabled is true, vrrp_ip must be set
    if n.VRRPEnabled && n.VRRPIP == "" {
        errs = append(errs, "iac.networking.vrrp_enabled=true requires vrrp_ip to be set")
    }
    // If use_designate is true, dns_zone_name must be set
    if n.UseDesignate && n.DNSZoneName == "" {
        errs = append(errs, "iac.networking.use_designate=true requires dns_zone_name to be set")
    }
    // If counts > 0, corresponding flavors must be set
    for role, count := range cfg.IAC.Counts {
        if count > 0 {
            if _, ok := cfg.IAC.Flavors[role]; !ok || cfg.IAC.Flavors[role] == "" {
                errs = append(errs, fmt.Sprintf("iac.counts.%s > 0 requires iac.flavors.%s to be set", role, role))
            }
        }
    }
    return errs
}

// ToJSON marshals the configuration to JSON. This is used for generating
// the JSON schema and for other tools that consume JSON.
//
// Outputs:
//   - []byte: The JSON-encoded configuration.
//   - error: An error if the configuration cannot be marshaled.
func (c Config) ToJSON() ([]byte, error) {
    return json.MarshalIndent(c, "", "  ")
}
