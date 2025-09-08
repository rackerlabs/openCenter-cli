# tests/features/git_push.feature
# Tests pushing the local GitOps repo to a git remote via `openCenter cluster bootstrap`.

Feature: Push GitOps repository to remote

  Background:
    Given an empty directory "tmp/conf"
    And an empty directory "tmp/repo-demo"
    And a file "tmp/conf/demo.yaml" with content:
      """
      cluster_name: demo
      gitops:
        git_dir: tmp/repo-demo
        git_url: ""
      kubernetes:
        counts: {}
        flavors: {}
        networking:
          use_octavia: true
          vrrp_enabled: false
      """
    And I run "openCenter cluster select demo --config-dir tmp/conf"
    And the exit code should be 0
    # Ensure repo has content to commit
    And I run "openCenter cluster setup --config-dir tmp/conf"
    And the exit code should be 0

  @bootstrap @success
  Scenario: Bootstrap pushes main branch to a bare remote repository
    Given a bare git repository exists at "tmp/remote.git"
    And I update the YAML "tmp/conf/demo.yaml" to set:
      """
      gitops:
        git_url: tmp/remote.git
      """
    When I run "openCenter cluster bootstrap --config-dir tmp/conf"
    Then the exit code should be 0
    And the bare repo "tmp/remote.git" should have branch "main"

  @bootstrap @error @missing_git_url
  Scenario: Bootstrap fails with a descriptive error when git_url is missing
    Given I update the YAML "tmp/conf/demo.yaml" to set:
      """
      gitops:
        git_url: ""
      """
    When I run "openCenter cluster bootstrap --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "git_url"
    And stderr should contain "must be set"

  @bootstrap @error @invalid_remote
  Scenario: Bootstrap fails when the remote repository is unreachable
    Given I update the YAML "tmp/conf/demo.yaml" to set:
      """
      gitops:
        git_url: tmp/does-not-exist.git
      """
    When I run "openCenter cluster bootstrap --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "remote"
    And stderr should contain "not found"

  @bootstrap @idempotent_push
  Scenario: A second bootstrap after changes pushes a new commit
    Given a bare git repository exists at "tmp/remote.git"
    And I update the YAML "tmp/conf/demo.yaml" to set:
      """
      gitops:
        git_url: tmp/remote.git
      """
    And a file "tmp/repo-demo/CHANGELOG.md" with content:
      """
      - chore: test commit for second push
      """
    When I run "openCenter cluster bootstrap --config-dir tmp/conf"
    Then the exit code should be 0
    And the bare repo "tmp/remote.git" should have branch "main"

