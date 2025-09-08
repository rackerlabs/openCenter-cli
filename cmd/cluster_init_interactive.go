// Copyright 2024 Victor Palma <victor.palma@rackspace.com>
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

package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/rackerlabs/openCenter/internal/config"
)

func runInitInteractive(cfg *config.Config) error {
	// Handlers for complex types
	k8sApiPort := fmt.Sprintf("%d", cfg.Kubernetes.K8sAPIPort)
	sshAuthKeys := strings.Join(cfg.Kubernetes.SSHAuthorizedKeys, "\n")
	dnsNameservers := strings.Join(cfg.Kubernetes.Networking.DNSNameservers, "\n")

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewNote().Title("Cluster Information"),
			huh.NewInput().
				Title("Cluster Name").
				Description("What is the name of your cluster?").
				Value(&cfg.ClusterName),
			huh.NewInput().
				Title("Naming Prefix").
				Description("Optional prefix for all resources.").
				Value(&cfg.NamingPrefix),
		),

		huh.NewGroup(
			huh.NewNote().Title("GitOps Configuration"),
			huh.NewInput().
				Title("Git Directory").
				Description("Local directory for GitOps repository.").
				Value(&cfg.GitOps.GitDir),
			huh.NewInput().
				Title("Git URL").
				Description("URL of the GitOps repository.").
				Value(&cfg.GitOps.GitURL),
			huh.NewInput().
				Title("Git SSH Key").
				Description("Path to the SSH key for Git access.").
				Value(&cfg.GitOps.GitSSHKey),
		),

		huh.NewGroup(
			huh.NewNote().Title("Terraform Settings"),
			huh.NewConfirm().
				Title("Enable Terraform").
				Value(&cfg.Terraform.Enabled),
			huh.NewInput().
				Title("Terraform Path").
				Description("Path to Terraform configuration files.").
				Value(&cfg.Terraform.Path),
		),

		huh.NewGroup(
			huh.NewNote().Title("Ansible Settings"),
			huh.NewConfirm().
				Title("Enable Ansible").
				Value(&cfg.Ansible.Enabled),
			huh.NewInput().
				Title("Ansible Path").
				Description("Path to Ansible configuration files.").
				Value(&cfg.Ansible.Path),
		),

		huh.NewGroup(
			huh.NewNote().Title("Kubernetes General"),
			huh.NewInput().
				Title("SSH User").
				Description("SSH username for nodes.").
				Value(&cfg.Kubernetes.SSHUser),
			huh.NewInput().
				Title("API Port").
				Description("Kubernetes API server port.").
				Value(&k8sApiPort),
			huh.NewInput().
				Title("Ubuntu Version").
				Description("Version of Ubuntu for nodes.").
				Value(&cfg.Kubernetes.UBVersion),
			huh.NewText().
				Title("SSH Authorized Keys").
				Description("Public SSH keys to authorize, one per line.").
				Value(&sshAuthKeys),
		),

		huh.NewGroup(
			huh.NewNote().Title("Kubernetes Networking"),
			huh.NewInput().
				Title("Nodes Subnet").
				Description("CIDR for the nodes subnet.").
				Value(&cfg.Kubernetes.Networking.SubnetNodes),
			huh.NewInput().
				Title("Services Subnet").
				Description("CIDR for the services subnet.").
				Value(&cfg.Kubernetes.Networking.SubnetServices),
			huh.NewInput().
				Title("Pods Subnet").
				Description("CIDR for the pods subnet.").
				Value(&cfg.Kubernetes.Networking.SubnetPods),
			huh.NewConfirm().
				Title("Enable Octavia").
				Description("Use Octavia for load balancing.").
				Value(&cfg.Kubernetes.Networking.UseOctavia),
			huh.NewSelect[string]().
				Title("Load Balancer Provider").
				Options(huh.NewOptions("amphora", "ovn")...).
				Value(&cfg.Kubernetes.Networking.LoadbalancerProvider),
		),

		huh.NewGroup(
			huh.NewNote().Title("Kubernetes DNS"),
			huh.NewConfirm().
				Title("Enable Designate").
				Description("Use Designate for DNS services.").
				Value(&cfg.Kubernetes.Networking.UseDesignate),
			huh.NewInput().
				Title("DNS Zone Name").
				Description("Name of the DNS zone.").
				Value(&cfg.Kubernetes.Networking.DNSZoneName),
			huh.NewText().
				Title("DNS Nameservers").
				Description("List of DNS nameservers, one per line.").
				Value(&dnsNameservers),
		),
		huh.NewGroup(
			huh.NewNote().Title("Cloud Provider"),
			huh.NewSelect[string]().
				Title("Cloud Provider").
				Options(huh.NewOptions("openstack")...).
				Value(&cfg.Cloud.Provider),
		),
	)

	if err := form.Run(); err != nil {
		return fmt.Errorf("failed to run interactive form: %w", err)
	}

	// Post-processing for complex types
	if port, err := strconv.Atoi(k8sApiPort); err == nil {
		cfg.Kubernetes.K8sAPIPort = port
	}
	cfg.Kubernetes.SSHAuthorizedKeys = strings.Split(sshAuthKeys, "\n")
	cfg.Kubernetes.Networking.DNSNameservers = strings.Split(dnsNameservers, "\n")

	if !cfg.Terraform.Enabled {
		cfg.Terraform.Path = ""
	}
	if !cfg.Ansible.Enabled {
		cfg.Ansible.Path = ""
	}
	if !cfg.Kubernetes.Networking.UseDesignate {
		cfg.Kubernetes.Networking.DNSZoneName = ""
	}

	return nil
}
