# tests/features/cluster_commands.feature
# Expected behavior for the "openCenter cluster" command group:
# - Parent "cluster" prints help & subcommands
# - list/ls scans config_dir for *.yaml and prints names (no .yaml); --json outputs JSON
# - select (by name & interactive), writes active_pointer; header when CWD == git_dir
# - info (active & named), human summary; --json prints full JSON; helpful errors
# - init (guided & non-interactive), does not overwrite unless --force; prints next steps
# - setup (materialize embedded templates into git_dir), idempotent, --force overwrites
# - bootstrap (git init/commit/remote/push) with actionable errors on missing prereqs

Feature: Cluster command group

  Background:
    Given an empty directory "tmp/conf"
    And an empty directory "tmp/repo-dev"
    And an empty directory "tmp/repo-prod"
    And a file "tmp/conf/dev.yaml" with content:
      """
      cluster_name: dev
      git_dir: tmp/repo-dev
      git_url: ""
      """
    And a file "tmp/conf/prod.yaml" with content:
      """
      cluster_name: prod
      git_dir: tmp/repo-prod
      git_url: ""
      """

  # ---------------------------------------------------------------------------
  # Parent: help shows subcommands
  # ---------------------------------------------------------------------------
  @help
  Scenario: "openCenter cluster" prints help with all subcommands
    When I run "openCenter cluster --config-dir tmp/conf"
    Then the exit code should be 0
    And stdout should contain "openCenter cluster list"
    And stdout should contain "openCenter cluster select"
    And stdout should contain "openCenter cluster info"
    And stdout should contain "openCenter cluster init"
    And stdout should contain "openCenter cluster setup"
    And stdout should contain "openCenter cluster bootstrap"

  # ---------------------------------------------------------------------------
  # list / ls
  # ---------------------------------------------------------------------------
  @list
  Scenario: Listing clusters shows file basenames without .yaml
    When I run "openCenter cluster list --config-dir tmp/conf"
    Then the exit code should be 0
    And stdout should contain "dev"
    And stdout should contain "prod"
    And stdout should not contain ".yaml"

  @list @json
  Scenario: Listing clusters as JSON
    When I run "openCenter cluster ls --json --config-dir tmp/conf"
    Then the exit code should be 0
    And stdout should contain "[" 
    And stdout should contain "\"dev\""
    And stdout should contain "\"prod\""

  @list @missing_dir
  Scenario: If config_dir does not exist, create it and print no entries
    Given the directory "<<tmp>>/fresh-conf" does not exist
    When I run "openCenter cluster list --config-dir <<tmp>>/fresh-conf"
    Then the exit code should be 0
    And the directory "<<tmp>>/fresh-conf" should exist
    And stdout should be empty

  # ---------------------------------------------------------------------------
  # select
  # ---------------------------------------------------------------------------
  @select @by_name
  Scenario: Selecting a cluster by name verifies file and writes active_pointer
    When I run "openCenter cluster select dev --config-dir tmp/conf"
    Then the exit code should be 0
    And the file "tmp/conf/.active" should match regex "^dev$"
    And stdout should contain "Selected cluster: dev"

  @select @interactive
  Scenario: Selecting a cluster interactively
    When I run interactively "openCenter cluster select --config-dir tmp/conf"
    And I choose "prod" from the prompt
    Then the exit code should be 0
    And the file "tmp/conf/.active" should match regex "^prod$"
    And stdout should contain "Selected cluster: prod"

  @select @missing
  Scenario: Selecting a non-existent cluster yields a helpful error
    When I run "openCenter cluster select missing --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "cluster 'missing' not found"
    And stderr should contain "openCenter cluster list"

  @select @header_in_git_dir
  Scenario: When CWD equals selected cluster's git_dir, subsequent commands show an active header
    Given I run "openCenter cluster select dev --config-dir tmp/conf"
    And the exit code should be 0
    And I cd to "tmp/repo-dev"
    When I run "openCenter cluster info --config-dir ../conf"
    Then the exit code should be 0
    And the first line of stdout should start with "Active cluster: dev"

  # ---------------------------------------------------------------------------
  # info
  # ---------------------------------------------------------------------------
  @info @active
  Scenario: Info without argument reads active_pointer
    Given I run "openCenter cluster select prod --config-dir tmp/conf"
    And the exit code should be 0
    When I run "openCenter cluster info --config-dir tmp/conf"
    Then the exit code should be 0
    And stdout should contain "cluster_name: prod"
    And stdout should contain "git_dir: tmp/repo-prod"

  @info @named @json
  Scenario: Info for a named cluster with --json prints full parsed config
    When I run "openCenter cluster info dev --json --config-dir tmp/conf"
    Then the exit code should be 0
    And stdout should contain "\"cluster_name\":\"dev\""
    And stdout should contain "\"git_dir\":\"tmp/repo-dev\""

  @info @unset_active
  Scenario: Info without active cluster set yields helpful message
    Given the file "tmp/conf/active" does not exist
    When I run "openCenter cluster info --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "no active cluster"

  @info @invalid_yaml
  Scenario: Invalid YAML is surfaced as a parse error
    Given a file "tmp/conf/bad.yaml" with content:
      """
      : not: yaml:
      """
    When I run "openCenter cluster info bad --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "parse"
    And stderr should contain "yaml"

  # ---------------------------------------------------------------------------
  # init
  # ---------------------------------------------------------------------------
  @init @by_name
  Scenario: init <cluster-name> creates a YAML with defaults; does not overwrite unless --force
    Given the file "tmp/conf/newone.yaml" does not exist
    When I run "openCenter cluster init newone --config-dir tmp/conf"
    Then the exit code should be 0
    And the file "tmp/conf/newone.yaml" should exist
    And stdout should contain "Created"
    When I run "openCenter cluster init newone --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "exists"
    When I run "openCenter cluster init newone --force --config-dir tmp/conf"
    Then the exit code should be 0

  @init @guided
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

  # ---------------------------------------------------------------------------
  # setup
  # ---------------------------------------------------------------------------
  @setup @materialize
  Scenario: setup materializes embedded templates into git_dir
    Given I run "openCenter cluster select dev --config-dir tmp/conf"
    And the exit code should be 0
    When I run "openCenter cluster setup --config-dir tmp/conf"
    Then the exit code should be 0
    And the directory "tmp/repo-dev" should contain a file matching "README.md"
    And stdout should contain "Created GitOps repo at"
    And stdout should contain "tmp/repo-dev"

  @setup @idempotent
  Scenario: setup is idempotent when run repeatedly
    Given I run "openCenter cluster select dev --config-dir tmp/conf"
    And the exit code should be 0
    And I run "openCenter cluster setup --config-dir tmp/conf"
    And the exit code should be 0
    When I run "openCenter cluster setup --config-dir tmp/conf"
    Then the exit code should be 0
    And stdout should contain "already initialized"

  @setup @force
  Scenario: setup --force overwrites existing files
    Given I run "openCenter cluster select dev --config-dir tmp/conf"
    And the exit code should be 0
    And a file "tmp/repo-dev/README.md" with content:
      """
      local edits that should be replaced
      """
    When I run "openCenter cluster setup --force --config-dir tmp/conf"
    Then the exit code should be 0
    And the file "tmp/repo-dev/README.md" should not contain "local edits that should be replaced"

  @setup @missing_prereqs
  Scenario: setup errors when no active cluster or git_dir is missing
    Given the file "tmp/conf/active" does not exist
    When I run "openCenter cluster setup --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "no active cluster"
    Given a file "tmp/conf/nogit.yaml" with content:
      """
      cluster_name: nogit
      kubernetes:
        networking:
          use_designate: false
      """
    When I run "openCenter cluster setup nogit --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "git_dir"
    And stderr should contain "must be set"

  # ---------------------------------------------------------------------------
  # bootstrap
  # ---------------------------------------------------------------------------
  @bootstrap
  Scenario: bootstrap pushes main to a remote
    Given a bare git repository exists at "tmp/remote.git"
    And I update the YAML "tmp/conf/dev.yaml" to set:
      """
      git_dir: tmp/repo-dev
      git_url: tmp/remote.git
      """
    And I run "openCenter cluster select dev --config-dir tmp/conf"
    And the exit code should be 0
    And I run "openCenter cluster setup --config-dir tmp/conf"
    And the exit code should be 0
    When I run "openCenter cluster bootstrap --config-dir tmp/conf"
    Then the exit code should be 0
    And the bare repo "tmp/remote.git" should have branch "main"

  @bootstrap @missing_prereqs
  Scenario: bootstrap errors on missing active cluster, git_dir, or git_url
    Given the file "tmp/conf/active" does not exist
    When I run "openCenter cluster bootstrap --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "no active cluster"
    Given I run "openCenter cluster select prod --config-dir tmp/conf"
    And the exit code should be 0
    And I update the YAML "tmp/conf/prod.yaml" to set:
      """
      git_dir: ""
      git_url: ""
      """
    When I run "openCenter cluster bootstrap --config-dir tmp/conf"
    Then the exit code should not be 0
    And stderr should contain "git_dir"
    And stderr should contain "git_url"
    And stderr should contain "must be set"

