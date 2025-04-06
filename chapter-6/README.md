# Chapter 6: Kubernetes Deployment Strategies

This directory contains Kubernetes manifests demonstrating various deployment strategies for the AvaKart application.

## Deployment Strategies

### 1. Rolling Updates (`rolling-updates/`)

Rolling updates allow for zero-downtime deployments by gradually replacing old pods with new ones. The `books-rolling-update.yaml` manifest demonstrates:

- Basic rolling update configuration
- Control over update parameters (maxSurge, maxUnavailable)
- Health checks for safe deployments
- Version labeling for tracking

### 2. Blue/Green Deployments (`blue-green/`)

Blue/Green deployments maintain two identical environments (blue for current version, green for new version) and switch traffic between them. The `frontend-blue-green.yaml` manifest demonstrates:

- Maintaining parallel deployments
- Service selector switching for traffic routing
- Zero-downtime cutover process
- Version labeling for tracking

### 3. Canary Releases (`canary/`)

Canary deployments allow testing new versions with a subset of users before full rollout. The `storage-canary.yaml` manifest demonstrates:

- Gradual traffic shifting with replica count
- Common service for both versions
- Monitoring canary deployment health
- Progressive rollout strategy

### 4. Advanced Deployment Patterns with Ingress (`advanced-ingress/`)

Advanced patterns using Ingress for more sophisticated traffic routing. The `frontend-ingress.yaml` manifest demonstrates:

- Advanced traffic splitting with percentages
- Host-based routing
- Ingress-based deployment strategies
- Progressive traffic shifting

## Usage

Use the provided Makefile to apply and manage the deployment strategies:

```bash
# Apply rolling update deployment
make apply-rolling-update

# Apply blue/green deployment
make apply-blue-green

# Switch traffic to green (new) version
make switch-to-green

# Apply canary deployment
make apply-canary

# See all available commands
make help
```

## Key Concepts

1. **Zero-downtime deployments**: All strategies ensure the application remains available during updates
2. **Controlled rollout**: Different strategies offer varying levels of control over the update process
3. **Traffic management**: Directing user traffic to different versions based on deployment strategy
4. **Health monitoring**: Using readiness and liveness probes to ensure safe deployments
5. **Rollback capability**: All strategies support quick rollback to previous versions if issues are detected
