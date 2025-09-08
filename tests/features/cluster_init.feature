Feature: Cluster initialisation
  As a user, I want to initialise a new cluster configuration using
  the `init` command, so that I can start defining my cluster layout.

  Scenario: Initialise a new cluster with default settings
    When I run "openCenter cluster init test-cluster"
    Then a cluster configuration "test-cluster" should exist
    And the cluster configuration "test-cluster" should have "cluster_name" set to "test-cluster"

  Scenario: Initialise a cluster and override string settings from flags
    When I run "openCenter cluster init test-cluster --kubernetes.ssh_user=debian --gitops.git_dir=/opt/openCenter/test-cluster"
    Then a cluster configuration "test-cluster" should exist
    And the cluster configuration "test-cluster" should have "kubernetes.ssh_user" set to "debian"
    And the cluster configuration "test-cluster" should have "gitops.git_dir" set to "/opt/openCenter/test-cluster"

  Scenario: Initialise a cluster and override boolean settings from flags
    When I run "openCenter cluster init test-cluster --ansible.enabled=false --kubernetes.networking.vrrp_enabled=true"
    Then a cluster configuration "test-cluster" should exist
    And the cluster configuration "test-cluster" should have "ansible.enabled" set to "false"
    And the cluster configuration "test-cluster" should have "kubernetes.networking.vrrp_enabled" set to "true"

  Scenario: Initialise a cluster and override integer settings from flags
    When I run "openCenter cluster init test-cluster --kubernetes.k8s_api_port=6443 --kubernetes.counts.master=3"
    Then a cluster configuration "test-cluster" should exist
    And the cluster configuration "test-cluster" should have "kubernetes.k8s_api_port" set to "6443"
    And the cluster configuration "test-cluster" should have "kubernetes.counts.master" set to "3"
