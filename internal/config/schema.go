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
    "fmt"
)

// GenerateSchema returns a JSON document representing the Draft 2020-12
// schema for the Config structure.
//
// The schema describes the nested properties of the configuration, and can
// be used for validation and documentation purposes.
//
// Inputs:
//   - pretty: If true, the output JSON will be indented.
//
// Outputs:
//   - []byte: The JSON schema document.
//   - error: An error if the schema cannot be generated.
func GenerateSchema(pretty bool) ([]byte, error) {
    schema := map[string]any{
        "$schema": "https://json-schema.org/draft/2020-12/schema",
        "title":   "openCenter Cluster Configuration",
        "type":    "object",
        "properties": map[string]any{
            "cluster_name": map[string]any{"type": "string"},
            "naming_prefix": map[string]any{"type": "string"},
            "gitops": map[string]any{
                "type":       "object",
                "properties": map[string]any{
                    "git_dir":    map[string]any{"type": "string"},
                    "git_url":    map[string]any{"type": "string"},
                    "git_ssh_key": map[string]any{"type": "string"},
                },
            },
            "terraform": map[string]any{
                "type":       "object",
                "properties": map[string]any{
                    "enabled": map[string]any{"type": "boolean"},
                    "path":    map[string]any{"type": "string"},
                },
            },
            "ansible": map[string]any{
                "type":       "object",
                "properties": map[string]any{
                    "enabled": map[string]any{"type": "boolean"},
                    "path":    map[string]any{"type": "string"},
                },
            },
            "kubernetes": map[string]any{
                "type":       "object",
                "properties": map[string]any{
                    "ssh_user":           map[string]any{"type": "string"},
                    "k8s_api_port":       map[string]any{"type": "integer"},
                    "ub_version":         map[string]any{"type": "string"},
                    "ssh_authorized_keys": map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
                    "ca_certificates":    map[string]any{"type": "string"},
                    "node_roles":          map[string]any{"type": "object", "additionalProperties": map[string]any{"type": "string"}},
                    "counts":             map[string]any{"type": "object", "additionalProperties": map[string]any{"type": "integer"}},
                    "images":             map[string]any{"type": "object", "additionalProperties": map[string]any{"type": "string"}},
                    "flavors":            map[string]any{"type": "object", "additionalProperties": map[string]any{"type": "string"}},
                    "storage": map[string]any{
                        "type": "object",
                        "properties": map[string]any{
                            "master_node_bfv": map[string]any{
                                "type":       "object",
                                "properties": map[string]any{"size": map[string]any{"type": "integer"}, "type": map[string]any{"type": "string"}},
                            },
                            "worker_node_bfv": map[string]any{
                                "type":       "object",
                                "properties": map[string]any{"size": map[string]any{"type": "integer"}, "type": map[string]any{"type": "string"}},
                            },
                            "worker_node_bfv_windows": map[string]any{
                                "type":       "object",
                                "properties": map[string]any{"size": map[string]any{"type": "integer"}, "type": map[string]any{"type": "string"}},
                            },
                        },
                    },
                    "windows": map[string]any{
                        "type":       "object",
                        "properties": map[string]any{"user": map[string]any{"type": "string"}, "admin_password": map[string]any{"type": "string"}},
                    },
                    "networking": map[string]any{
                        "type":       "object",
                        "properties": map[string]any{
                            "subnet_nodes":          map[string]any{"type": "string"},
                            "allocation_pool_start": map[string]any{"type": "string"},
                            "allocation_pool_end":   map[string]any{"type": "string"},
                            "vrrp_enabled":          map[string]any{"type": "boolean"},
                            "vrrp_ip":               map[string]any{"type": "string"},
                            "subnet_services":       map[string]any{"type": "string"},
                            "subnet_pods":           map[string]any{"type": "string"},
                            "use_octavia":           map[string]any{"type": "boolean"},
                            "loadbalancer_provider": map[string]any{"type": "string"},
                            "use_designate":         map[string]any{"type": "boolean"},
                            "dns_zone_name":         map[string]any{"type": "string"},
                            "dns_nameservers":       map[string]any{"type": "array", "items": map[string]any{"type": "string"}},
                            "vlan": map[string]any{
                                "type":       "object",
                                "properties": map[string]any{"id": map[string]any{"type": "string"}, "mtu": map[string]any{"type": "integer"}, "provider": map[string]any{"type": "string"}},
                            },
                        },
                    },
                },
            },
            "cloud": map[string]any{
                "type":       "object",
                "properties": map[string]any{
                    "provider": map[string]any{"type": "string"},
                    "openstack": map[string]any{
                        "type": "object",
                        "properties": map[string]any{
                            "auth_url":          map[string]any{"type": "string"},
                            "insecure":          map[string]any{"type": "boolean"},
                            "region":            map[string]any{"type": "string"},
                            "user_name":         map[string]any{"type": "string"},
                            "user_password":     map[string]any{"type": "string"},
                            "admin_password":    map[string]any{"type": "string"},
                            "project_domain_name": map[string]any{"type": "string"},
                            "user_domain_name":    map[string]any{"type": "string"},
                            "tenant_name":        map[string]any{"type": "string"},
                            "availability_zone":  map[string]any{"type": "string"},
                            "floatingip_pool":     map[string]any{"type": "string"},
                            "router_external_network_id": map[string]any{"type": "string"},
                            "disable_bastion":     map[string]any{"type": "boolean"},
                            "ca":                  map[string]any{"type": "string"},
                        },
                    },
                },
            },
        },
        "required": []string{"cluster_name", "gitops.git_dir"},
    }
    var data []byte
    var err error
    if pretty {
        data, err = json.MarshalIndent(schema, "", "  ")
    } else {
        data, err = json.Marshal(schema)
    }
    if err != nil {
        return nil, fmt.Errorf("failed to marshal schema: %w", err)
    }
    return data, nil
}
