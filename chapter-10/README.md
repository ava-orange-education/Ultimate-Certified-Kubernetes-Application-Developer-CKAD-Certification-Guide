# Chapter 10: Probes and Health Checks

This chapter focuses on implementing Kubernetes probes and health checks for the AvaKart application components. Proper health check implementation is essential for ensuring application reliability, availability, and self-healing capabilities in Kubernetes.

## Directory Structure

- `basic-probes/`: Basic implementations of the three types of Kubernetes probes
  - `books-deployment.yaml`: Liveness probe implementation
  - `storage-deployment.yaml`: Readiness probe implementation
  - `order-processor-deployment.yaml`: Startup probe implementation

- `advanced-probes/`: Advanced probe configurations
  - `frontend-deployment.yaml`: HTTP probe with headers and parameters
  - `storage-tcp-probe.yaml`: TCP socket probe implementation
  - `books-exec-probe.yaml`: Exec command probe implementation

- `comprehensive/`: Comprehensive health check implementation
  - `order-processor-advanced-health.yaml`: Deployment with all three probe types working together

## Types of Kubernetes Probes

### 1. Liveness Probe
- **Purpose**: Checks if the application is running properly
- **Behavior**: If the probe fails, Kubernetes restarts the container
- **Use Case**: Detect application deadlocks, infinite loops, or other states where the application is running but unable to make progress
- **Example**: `/health/live` endpoint in the Books Service

### 2. Readiness Probe
- **Purpose**: Checks if the pod can receive traffic
- **Behavior**: If the probe fails, Kubernetes removes the pod from service endpoints
- **Use Case**: Prevent traffic from being sent to pods that are not ready to handle requests
- **Example**: `/health/ready` endpoint in the Storage Service

### 3. Startup Probe
- **Purpose**: Gives the application time to start up
- **Behavior**: Disables liveness and readiness checks until the application is ready
- **Use Case**: Applications with slow startup times
- **Example**: `/health/startup` endpoint in the Order Processor Service

## Advanced Probe Configurations

### HTTP Probe with Headers
- Allows sending custom HTTP headers with the probe request
- Useful for services that require authentication or specific headers

### TCP Socket Probe
- Checks if a TCP socket can be opened to a specified port
- Useful for services that don't have HTTP endpoints

### Exec Command Probe
- Executes a command inside the container
- Useful for custom health checks that require shell commands

## Comprehensive Health Check Implementation

The `order-processor-advanced-health.yaml` demonstrates a comprehensive health check strategy with all three probe types working together:

1. **Startup Probe**: Gives the application time to start up
2. **Readiness Probe**: Controls when the pod receives traffic
3. **Liveness Probe**: Restarts the pod if it becomes unhealthy

## Usage

Use the provided Makefile to apply and manage the Kubernetes manifests:

```bash
# Apply basic probe implementations
make apply-basic

# Apply advanced probe configurations
make apply-advanced

# Apply comprehensive health check implementation
make apply-comprehensive

# Apply all manifests
make apply-all

# Delete all manifests
make delete-all

# Check the status of the deployments
make status

# Describe the deployments to see probe configurations
make describe
```

## Best Practices

1. **Implement all three probe types when appropriate**
2. **Keep probes lightweight** to minimize resource usage
3. **Set appropriate timing parameters** based on application characteristics
4. **Configure failure thresholds carefully** to avoid premature restarts
5. **Use success thresholds for readiness** to prevent flapping services
