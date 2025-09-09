Feature: Update existing cluster configuration via dotted flags

  Background:
    Given an empty directory "tmp/conf"
    And a cluster "upd1" exists
    And a cluster "upd2" exists

  @config @update @by_name
  Scenario: Update a named cluster using dotted flags
    When I run "openCenter cluster update upd1 --iac.counts.master=3 --iac.k8s_api_port=6444 --config-dir tmp/conf"
    Then the exit code should be 0
    And the cluster configuration "upd1" should have "iac.counts.master" set to "3"
    And the cluster configuration "upd1" should have "iac.k8s_api_port" set to "6444"

  @config @update @active
  Scenario: Update the active cluster when name is omitted
    Given the active cluster is "upd2"
    When I run "openCenter cluster update --gitops.git_branch=dev --config-dir tmp/conf"
    Then the exit code should be 0
    And the cluster configuration "upd2" should have "gitops.git_branch" set to "dev"

