# Chapter 7: Rolling Updates & Rollbacks

This directory contains Kubernetes manifests demonstrating advanced rolling update configurations and rollback mechanisms for the AvaKart application.

## Key Concepts

### 1. Advanced Rolling Update Configurations (`rolling-updates/`)

The `books-rolling-update.yaml` manifest demonstrates advanced rolling update configurations:

- Percentage-based update parameters (maxSurge, maxUnavailable)
- Revision history configuration for rollbacks
- Detailed health check configurations
- Resource constraints for consistent performance
- Change-cause annotation for tracking deployment reasons
- Version labeling for tracking

### 2. Rollback Mechanisms and Strategies (`rollback/`)

The rollback directory contains files demonstrating rollback mechanisms:

- `books-deployment-v1.yaml`: Initial version of the books service
- `books-deployment-v2.yaml`: Updated version with potential issues
- `rollback-demo.sh`: Script demonstrating the rollback process

Key rollback operations:
- Viewing deployment history
- Rolling back to the previous version
- Rolling back to a specific revision
- Monitoring rollback status
- Verifying rollback success

### 3. Advanced Health Checking with Probes (`health-probes/`)

The `order-processor-advanced-health.yaml` manifest demonstrates comprehensive health checking:

- Startup probes to give applications time to initialize
- Readiness probes to control when pods receive traffic
- Liveness probes to detect and restart unhealthy pods
- Configuring appropriate timing parameters for each probe type

## Usage

Use the provided Makefile to apply and manage the deployment strategies:

```bash
# Apply advanced rolling update configuration
make apply-rolling-update

# Apply books service v1 deployment
make apply-v1

# Apply books service v2 deployment (with potential issues)
make apply-v2

# Rollback to the previous deployment version
make rollback

# Show deployment revision history
make rollout-history

# Run the rollback demonstration script
make run-rollback-demo

# See all available commands
make help
```

## Best Practices

1. **Revision History Management**: Configure `revisionHistoryLimit` to store an appropriate number of revisions for rollback
2. **Change Tracking**: Use annotations like `kubernetes.io/change-cause` to document deployment changes
3. **Health Check Configuration**: Implement appropriate startup, readiness, and liveness probes with suitable timing parameters
4. **Resource Management**: Set appropriate resource requests and limits for consistent performance
5. **Rollback Testing**: Regularly test rollback procedures to ensure they work as expected
6. **Monitoring**: Implement monitoring to detect issues quickly and trigger rollbacks when necessary
7. **Version Labeling**: Use consistent version labeling to track deployments and rollbacks
