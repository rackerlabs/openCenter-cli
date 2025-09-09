Feature: Configuration validation rules

  Background:
    Given an empty directory "<<tmp>>/conf"
    And an empty directory "<<tmp>>/repo-bad"

  @validation @octavia_vrrp_conflict
  Scenario: use_octavia=true and vrrp_enabled=true -> error
    Given a file "<<tmp>>/conf/bad-octavia-vrrp.yaml" with content:
      """
      cluster_name: bad-octavia-vrrp
      gitops: { git_dir: "<<tmp>>/repo-bad" }
      iac:
        counts: {}
        flavors: {}
        networking:
          use_octavia: true
          vrrp_enabled: true
      """
    When I run "openCenter cluster info bad-octavia-vrrp --validate"
    Then the exit code should not be 0
    And stderr should contain "iac.networking.use_octavia=true and vrrp_enabled=true are mutually exclusive"

  @validation @vrrp_missing_ip
  Scenario: use_octavia=false and missing vrrp_ip -> error
    Given a file "<<tmp>>/conf/bad-vrrp-ip.yaml" with content:
      """
      cluster_name: bad-vrrp-ip
      gitops: { git_dir: "<<tmp>>/repo-bad" }
      iac:
        counts: {}
        flavors: {}
        networking:
          use_octavia: false
          vrrp_enabled: true
          vrrp_ip: ""
      """
    When I run "openCenter cluster info bad-vrrp-ip --validate"
    Then the exit code should not be 0
    And stderr should contain "iac.networking.use_octavia=false requires vrrp_ip to be set"

  @validation @designate_missing_zone
  Scenario: use_designate=true and missing dns_zone_name -> error
    Given a file "<<tmp>>/conf/bad-designate.yaml" with content:
      """
      cluster_name: bad-designate
      gitops: { git_dir: "<<tmp>>/repo-bad" }
      iac:
        counts: {}
        flavors: {}
        networking:
          use_octavia: true
          vrrp_enabled: false
          use_designate: true
          dns_zone_name: ""
      """
    When I run "openCenter cluster info bad-designate --validate"
    Then the exit code should not be 0
    And stderr should contain "iac.networking.use_designate=true requires dns_zone_name to be set"

  @validation @counts_without_flavors
  Scenario Outline: Node counts > 0 require corresponding flavors
    Given a file "<<tmp>>/conf/<name>.yaml" with content:
      """
      cluster_name: <name>
      gitops: { git_dir: "<<tmp>>/repo-bad" }
      iac:
        counts: { <role>: <count> }
        flavors: { }
        networking: { use_octavia: true, vrrp_enabled: false, use_designate: false }
      """
    When I run "openCenter cluster info <name> --validate"
    Then the exit code should not be 0
    And stderr should contain "iac.counts.<role> > 0 requires iac.flavors.<role> to be set"

    Examples:
      | name           | role   | count |
      | bad-master     | master | 1     |
      | bad-worker     | worker | 2     |
      | bad-windows    | win    | 1     |

