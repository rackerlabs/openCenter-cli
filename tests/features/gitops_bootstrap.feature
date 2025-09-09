Feature: GitOps bootstrap pushes to remote

  Background:
    Given an empty directory "tmp/conf"
    And an empty directory "tmp/repo-dev"
    And a file "tmp/conf/dev.yaml" with content:
      """
      cluster_name: dev
      gitops:
        git_dir: "tmp/repo-dev"
        git_url: ""
      """

  @gitops @bootstrap
  Scenario: bootstrap pushes main to a remote
    Given a bare git repository exists at "tmp/remote.git"
    And I update the YAML "tmp/conf/dev.yaml" to set:
      """
      gitops:
        git_dir: "tmp/repo-dev"
        git_url: "tmp/remote.git"
      """
    And I run "openCenter cluster select dev --config-dir tmp/conf"
    And the exit code should be 0
    And I run "openCenter cluster setup --config-dir tmp/conf"
    And the exit code should be 0
    When I run "openCenter cluster bootstrap --config-dir tmp/conf"
    Then the exit code should be 0
    And the bare repo "tmp/remote.git" should have branch "main"

  @gitops @bootstrap @missing_prereqs
  Scenario: bootstrap errors on missing active cluster, git_dir, or git_url
    Given the file "tmp/conf/active" does not exist
    When I run "openCenter cluster bootstrap --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "no active cluster"
    Given I run "openCenter cluster select dev --config-dir tmp/conf"
    And the exit code should be 0
    And I update the YAML "tmp/conf/dev.yaml" to set:
      """
      gitops:
        git_dir: ""
        git_url: ""
      """
    When I run "openCenter cluster bootstrap --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "git_dir"
    And stderr should contain "git_url"
    And stderr should contain "must be set"

