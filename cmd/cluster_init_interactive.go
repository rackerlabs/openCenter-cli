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
    k8sApiPort := fmt.Sprintf("%d", cfg.IAC.K8sAPIPort)
    sshAuthKeys := strings.Join(cfg.IAC.SSHAuthorizedKeys, "\n")
    dnsNameservers := strings.Join(cfg.IAC.Networking.DNSNameservers, "\n")
    ansiblePlaybooks := strings.Join(cfg.Ansible.Playbooks, "\n")
    sopsKeyPath := cfg.Secrets.SopsAgeKeyFile
    verifyNow := true

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
            huh.NewInput().
                Title("Git Branch").
                Description("Branch to push to (defaults to 'main' if empty).").
                Value(&cfg.GitOps.GitBranch),
            huh.NewInput().
                Title("Flux Interval").
                Description("Optional reconciliation interval (e.g., 1m).").
                Value(&cfg.GitOps.Flux.Interval),
            huh.NewConfirm().
                Title("Flux Prune").
                Description("Enable Flux prune (optional).").
                Value(&cfg.GitOps.Flux.Prune),
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
            huh.NewInput().
                Title("Ansible Inventory").
                Description("Optional inventory filename or path.").
                Value(&cfg.Ansible.Inventory),
            huh.NewText().
                Title("Ansible Playbooks").
                Description("Optional playbooks, one per line.").
                Value(&ansiblePlaybooks),
		),

		huh.NewGroup(
            huh.NewNote().Title("IAC / Kubernetes General"),
			huh.NewInput().
				Title("SSH User").
				Description("SSH username for nodes.").
                Value(&cfg.IAC.SSHUser),
			huh.NewInput().
				Title("API Port").
				Description("Kubernetes API server port.").
				Value(&k8sApiPort),
			huh.NewInput().
				Title("Ubuntu Version").
				Description("Version of Ubuntu for nodes.").
                Value(&cfg.IAC.UBVersion),
			huh.NewText().
				Title("SSH Authorized Keys").
				Description("Public SSH keys to authorize, one per line.").
				Value(&sshAuthKeys),
		),

		huh.NewGroup(
            huh.NewNote().Title("IAC / Kubernetes Networking"),
			huh.NewInput().
				Title("Nodes Subnet").
				Description("CIDR for the nodes subnet.").
                Value(&cfg.IAC.Networking.SubnetNodes),
			huh.NewInput().
				Title("Services Subnet").
				Description("CIDR for the services subnet.").
                Value(&cfg.IAC.Networking.SubnetServices),
			huh.NewInput().
				Title("Pods Subnet").
				Description("CIDR for the pods subnet.").
                Value(&cfg.IAC.Networking.SubnetPods),
			huh.NewConfirm().
				Title("Enable Octavia").
				Description("Use Octavia for load balancing.").
                Value(&cfg.IAC.Networking.UseOctavia),
			huh.NewSelect[string]().
				Title("Load Balancer Provider").
				Options(huh.NewOptions("amphora", "ovn")...).
                Value(&cfg.IAC.Networking.LoadbalancerProvider),
		),

	huh.NewGroup(
            huh.NewNote().Title("IAC / DNS"),
			huh.NewConfirm().
				Title("Enable Designate").
				Description("Use Designate for DNS services.").
                Value(&cfg.IAC.Networking.UseDesignate),
			huh.NewInput().
				Title("DNS Zone Name").
				Description("Name of the DNS zone.").
                Value(&cfg.IAC.Networking.DNSZoneName),
			huh.NewText().
				Title("DNS Nameservers").
				Description("List of DNS nameservers, one per line.").
				Value(&dnsNameservers),
	),

		// Secrets and optional verification
		huh.NewGroup(
			huh.NewNote().Title("Secrets & Verification"),
			huh.NewInput().
				Title("SOPS Age Key File").
				Description("Optional path. Leave empty to auto-generate on save.").
				Value(&sopsKeyPath),
			huh.NewConfirm().
				Title("Verify configuration now").
				Description("Run validation and review results before saving.").
				Value(&verifyNow),
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
        cfg.IAC.K8sAPIPort = port
    }
    cfg.IAC.SSHAuthorizedKeys = strings.Split(sshAuthKeys, "\n")
    cfg.IAC.Networking.DNSNameservers = strings.Split(dnsNameservers, "\n")
    if strings.TrimSpace(ansiblePlaybooks) != "" {
        var pbs []string
        for _, line := range strings.Split(ansiblePlaybooks, "\n") {
            line = strings.TrimSpace(line)
            if line != "" {
                pbs = append(pbs, line)
            }
        }
        cfg.Ansible.Playbooks = pbs
    }
    cfg.Secrets.SopsAgeKeyFile = strings.TrimSpace(sopsKeyPath)

	if !cfg.Terraform.Enabled {
		cfg.Terraform.Path = ""
	}
	if !cfg.Ansible.Enabled {
		cfg.Ansible.Path = ""
	}
    if !cfg.IAC.Networking.UseDesignate {
        cfg.IAC.Networking.DNSZoneName = ""
    }

    // Optional verification step
    if verifyNow {
        errs := config.Validate(*cfg)
        summary := "Validation successful."
        if len(errs) > 0 {
            summary = fmt.Sprintf("Validation found %d issue(s):\n- %s", len(errs), strings.Join(errs, "\n- "))
        }
        proceed := true
        vf := huh.NewForm(
            huh.NewGroup(
                huh.NewNote().Title("Verification Results").Description(summary),
                huh.NewConfirm().Title("Proceed to save configuration?").Value(&proceed),
            ),
        )
        if err := vf.Run(); err != nil {
            return fmt.Errorf("verification form failed: %w", err)
        }
        if !proceed {
            return fmt.Errorf("aborted by user after verification")
        }
    }

	return nil
}
