Feature: Interactive init wizard

  Background:
    Given an empty directory "tmp/conf"

  @wizard @init
  Scenario: Guided init prompts for cluster_name and other defaults
    When I run interactively "openCenter cluster init --config-dir tmp/conf"
    And I answer the prompts with:
      | prompt         | answer        |
      | cluster_name   | guided        |
      | git_dir        | tmp/repo-guid |
    Then the exit code should be 0
    And the file "tmp/conf/guided.yaml" should exist
    And the file "tmp/conf/guided.yaml" should contain "cluster_name: guided"
    And the file "tmp/conf/guided.yaml" should contain "git_dir: tmp/repo-guid"
    And stdout should contain "openCenter cluster select guided"

