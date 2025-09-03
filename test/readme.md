# Test Scenario Submission Guide

## Quick Start

This guide helps you write and submit test scenarios for our openCenter repository.

## How to Submit a Test Scenario

### 1. Choose the Right Location

Submit your test scenarios to the appropriate directory:
```
/test-scenarios/
├── kubernetes/          # K8s deployment, scaling, networking tests
├── openstack/          # VM, storage, network management tests  
├── sitecore/           # CMS functionality, performance tests
├── integration/        # Cross-system interaction tests
└── infrastructure/     # General infrastructure tests
```

### 2. Use This Simple Template

```gherkin
Feature: [What are you testing - keep it short]

Scenario: [Describe the specific test case]
  Given [the starting situation]
  When [the action you're testing]  
  Then [what should happen]
```

## Real Examples for Our Systems

### Kubernetes Tests

```gherkin
Feature: Pod Deployment

Scenario: Deploy application pod successfully
  Given the Kubernetes cluster is running
  When I deploy a new application pod with valid configuration
  Then the pod should be in "Running" status within 60 seconds
  And the application should respond to health checks
```

### OpenStack Tests

```gherkin
Feature: VM Management

Scenario: Create virtual machine
  Given OpenStack is available with sufficient resources
  When I create a VM with 2 CPUs and 4GB RAM
  Then the VM should be created successfully
  And I should be able to SSH into the VM
```

### Sitecore Tests

```gherkin
Feature: Content Publishing

Scenario: Publish content item
  Given I have a content item in draft status
  When I publish the content item
  Then the item should appear on the live website
  And the item should be accessible via its URL
```

### Integration Tests

```gherkin
Feature: Application Deployment Pipeline

Scenario: Deploy application from CI/CD to Kubernetes
  Given the application image is built and pushed to registry
  When the deployment pipeline is triggered
  Then the application should deploy to Kubernetes cluster
  And the load balancer should route traffic to the new pods
  And OpenStack storage should be mounted correctly
```

## Writing Guidelines

### ✅ Do This
- Use **simple, clear language**
- Be **specific** about expected results
- Include **timing expectations** (e.g., "within 30 seconds")
- Test **one main thing** per scenario
- Use **realistic data** (actual resource sizes, timeouts)

### ❌ Avoid This
- Technical jargon that business users won't understand
- Testing multiple unrelated things in one scenario  
- Vague expectations like "should work properly"
- Implementation details like API endpoints or database queries

## Common Patterns for Our Infrastructure

### Resource Management
```gherkin
Given [system] has [specific resources available]
When I [create/scale/delete] [resource type]
Then [specific outcome with metrics]
```

### Error Handling  
```gherkin
Given [system] is in [specific state]
When [failure condition occurs]
Then [system should handle gracefully]
And [specific recovery action should happen]
```

### Performance Testing
```gherkin
Given [system] is under [specific load]
When [action is performed]
Then [response time should be] less than [X seconds]
And [system resources should remain] below [Y%]
```

## File Naming Convention

Name your files descriptively:
- `kubernetes-pod-scaling.feature`
- `openstack-vm-backup.feature`  
- `sitecore-content-sync.feature`
- `integration-app-deployment.feature`

## Submission Process

1. **Create** your `.feature` file in the appropriate directory
2. **Follow** the template and examples above
3. **Submit** via pull request with description:
   ```
   Test Scenario: [Brief description]
   
   - System(s): [K8s/OpenStack/Sitecore/etc.]
   - Test Type: [Functional/Performance/Integration]
   - Priority: [High/Medium/Low]
   ```

## Example Scenarios You Can Submit

### Infrastructure Scenarios
- VM creation and deletion
- Storage volume mounting/unmounting
- Network connectivity between services
- Load balancer configuration
- Backup and restore operations

### Application Scenarios  
- Application deployment and rollback
- Scaling applications up/down
- Service discovery and communication
- Configuration updates
- Health monitoring and alerting

### Integration Scenarios
- Data synchronization between systems
- Authentication across services  
- Monitoring and logging aggregation
- Disaster recovery procedures

## Getting Help

- **Unclear about syntax?** Check the examples above
- **Don't know which directory?** Ask in the team chat
- **Need technical details?** The development team will help implement
- **Want to discuss the scenario?** Schedule a quick review meeting

## Review Process

1. **Automated checks** verify your syntax is correct
2. **Team review** ensures the scenario makes business sense  
3. **Technical review** confirms feasibility with our infrastructure
4. **Implementation** by the development team (Go/Python)
5. **Integration** into our test suite

---

**Remember**: You don't need to know how to code! Just describe what you want to test in simple terms. The development team will handle the technical implementation.
