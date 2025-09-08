Feature: GitOps template rendering
  Background:
    Given an empty directory "<<tmp>>/conf"

  Scenario: Render a simple template
    Given a file "<<tmp>>/conf/demo.yaml" with content:
      """
      cluster_name: demo
      gitops:
        git_dir: "<<tmp>>/opencenter-demo"
      """
    When I run "openCenter cluster render demo --config-dir <<tmp>>/conf"
    Then a file "<<tmp>>/opencenter-demo/KUSTOMIZATION" should exist
    And the file "<<tmp>>/opencenter-demo/KUSTOMIZATION" should contain "resources:"