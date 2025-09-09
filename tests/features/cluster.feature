Feature: openCenter cluster basics
  Background:
    Given an empty directory "<<tmp>>/conf"

  Scenario: Initialize a cluster with defaults
    When I run "openCenter cluster init demo --config-dir <<tmp>>/conf"
    Then a file "<<tmp>>/conf/demo.yaml" should exist
    And the file "<<tmp>>/conf/demo.yaml" should contain "cluster_name: demo"

  Scenario: Select the cluster
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      """
    When I run "openCenter cluster select demo --config-dir <<tmp>>/conf"
    Then the file "<<tmp>>/conf/.active" should match regex "^demo$"

  Scenario: Show current cluster
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      """
    And I run "openCenter cluster select demo --config-dir <<tmp>>/conf"
    When I run "openCenter cluster current --config-dir <<tmp>>/conf"
    Then stdout should contain "demo"

  Scenario: List clusters
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      """
    And a file "<<tmp>>/conf/blue.yaml" with content:
      """
      cluster_name: blue
      """
    And a file "<<tmp>>/conf/green.yaml" with content:
      """
      cluster_name: green
      """
    When I run "openCenter cluster list --config-dir <<tmp>>/conf"
    Then stdout should contain:
      """
      blue
      demo
      green
      """

  Scenario: List clusters as JSON
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      """
    And a file "<<tmp>>/conf/blue.yaml" with content:
      """
      cluster_name: blue
      """
    And a file "<<tmp>>/conf/green.yaml" with content:
      """
      cluster_name: green
      """
    When I run "openCenter cluster list --json --config-dir <<tmp>>/conf"
    Then stdout should contain '["blue","demo","green"]'

  Scenario: Info for a cluster
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      """
    And I run "openCenter cluster select demo --config-dir <<tmp>>/conf"
    When I run "openCenter cluster info --config-dir <<tmp>>/conf"
    Then stdout should contain "cluster_name: demo"

  Scenario: Validate constraints
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      kubernetes:
        networking:
          use_octavia: false
          vrrp_ip: "10.0.0.10"
          dns_zone_name: "example.com"
      """
    When I run "openCenter cluster validate demo --config-dir <<tmp>>/conf"
    Then the exit code should be 0
    And stdout should contain "Validation successful."

  Scenario: Validate constraints failure
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      kubernetes:
        networking:
          use_octavia: true
          vrrp_enabled: true
      """
    When I run "openCenter cluster validate demo --config-dir <<tmp>>/conf"
    Then exit code should be 1
    And stderr should contain "kubernetes.networking.use_octavia=true and vrrp_enabled=true are mutually exclusive"

  Scenario: Preflight
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      """
    And I run "openCenter cluster select demo --config-dir <<tmp>>/conf"
    When I run "openCenter cluster preflight --config-dir <<tmp>>/conf"
    Then stdout should contain "Preflight complete."

  @hangs
  Scenario: Setup GitOps with rendering
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      gitops:
        git_dir: "<<tmp>>/opencenter-demo"
      kubernetes:
        networking:
          use_designate: false
      """
    When I run "openCenter cluster setup demo --config-dir <<tmp>>/conf"
    Then a directory "<<tmp>>/opencenter-demo" should exist
    And a file "<<tmp>>/opencenter-demo/KUSTOMIZATION" should exist

  @hangs @wip
  @wip
  Scenario: Bootstrap pushes a new commit to a remote repository
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      gitops:
        git_dir: "<<tmp>>/opencenter-demo"
      kubernetes:
        networking:
          use_designate: false
          """

    And a bare git repository exists at "<<tmp>>/remote.git"
    And I update the YAML "<<tmp>>/conf/demo.yaml" to set:
      """
      gitops:
      git_url: "git@localhost:newuser/gitops-repo.git"
      """
    And I run "openCenter cluster setup demo --config-dir <<tmp>>/conf"
    When I run "openCenter cluster bootstrap demo --force --config-dir <<tmp>>/conf"
    Then the command should succeed
    And the remote git repository should contain a "Bootstrap commit"

  @hangs
  Scenario: Setup with provisioning
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      gitops:
        git_dir: "<<tmp>>/opencenter-demo"
      terraform:
        enabled: true
      ansible:
        enabled: true
      kubernetes:
        networking:
          use_designate: false
      """
    When I run "openCenter cluster setup demo --config-dir <<tmp>>/conf"
    Then a directory "<<tmp>>/opencenter-demo/terraform" should exist
    And a file "<<tmp>>/opencenter-demo/terraform/main.tf" should exist
    And a directory "<<tmp>>/opencenter-demo/ansible" should exist
    And a file "<<tmp>>/opencenter-demo/ansible/inventory" should exist

  @skip 
  Scenario: Destroy a cluster
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      gitops:
        git_dir: "<<tmp>>/opencenter-demo"
      """
    When I run "openCenter cluster destroy demo --config-dir <<tmp>>/conf"
    Then the command should succeed
    And a file "<<tmp>>/conf/demo.yaml" should not exist
    And a directory "<<tmp>>/opencenter-demo" should not exist
