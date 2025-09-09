Feature: Cluster initialisation
  As a user, I want to initialise a new cluster configuration using
  the `init` command, so that I can start defining my cluster layout.

  Scenario: Initialise a new cluster with default settings
    When I run "openCenter cluster init test-cluster"
    Then a cluster configuration "test-cluster" should exist
    And the cluster configuration "test-cluster" should have "cluster_name" set to "test-cluster"

  Scenario: Initialise a cluster and override string settings from flags
    When I run "openCenter cluster init test-cluster --iac.ssh_user=debian --gitops.git_dir=/opt/openCenter/test-cluster"
    Then a cluster configuration "test-cluster" should exist
    And the cluster configuration "test-cluster" should have "iac.ssh_user" set to "debian"
    And the cluster configuration "test-cluster" should have "gitops.git_dir" set to "/opt/openCenter/test-cluster"

  Scenario: Initialise a cluster and override boolean settings from flags
    When I run "openCenter cluster init test-cluster --ansible.enabled=false --iac.networking.vrrp_enabled=true"
    Then a cluster configuration "test-cluster" should exist
    And the cluster configuration "test-cluster" should have "ansible.enabled" set to "false"
    And the cluster configuration "test-cluster" should have "iac.networking.vrrp_enabled" set to "true"

  Scenario: Initialise a cluster and override integer settings from flags
    When I run "openCenter cluster init test-cluster --iac.k8s_api_port=6443 --iac.counts.master=3"
    Then a cluster configuration "test-cluster" should exist
    And the cluster configuration "test-cluster" should have "iac.k8s_api_port" set to "6443"
    And the cluster configuration "test-cluster" should have "iac.counts.master" set to "3"

  Scenario: Init generates a SOPS key when not provided
    When I run "openCenter cluster init demo --gitops.git_dir=<<tmp>>/repo-demo"
    Then a file "~/.config/openCenter/sops/age/keys/demo-key.txt" should exist
    And the file "~/.config/openCenter/sops/age/keys/demo-key.txt" should contain "AGE-SECRET-KEY-1"

  Scenario: Init does not generate a SOPS key when disabled
    When I run "openCenter cluster init demo2 --gitops.git_dir=<<tmp>>/repo-demo2 --no-sops-keygen"
    Then the file "~/.config/openCenter/sops/age/keys/demo2-key.txt" should not exist
    And the cluster configuration "demo2" should have "secrets.sops_age_key_file" set to ""
