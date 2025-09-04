Feature: Initialize an openCenter cluster configuration

  As a user of the openCenter CLI
  I want to initialize a cluster configuration
  So that I have a cluster.yaml with all required settings

  Scenario: Running openCenter cluster init
    Given I have the openCenter CLI installed
    When I run "openCenter cluster init"
    Then a file named "cluster.yaml" should be created
    And the file should contain the cluster name
    And the file should contain an SSH key reference
    And the file should contain an S3 bucket for openTofu state
    And the file should contain FluxCD configuration details
    And the file should contain a location for the kubeconfig

