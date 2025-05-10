# AvaKart Kubernetes Security Features

This document outlines the security features implemented in the AvaKart Kubernetes manifests.

## Security Features Overview

The AvaKart application has been secured using the following Kubernetes security features:

1. **Service Accounts**: Dedicated service accounts for each service
2. **Security Contexts**: Pod and container level security settings
3. **RBAC**: Role-Based Access Control for fine-grained permissions
4. **Network Policies**: Control pod-to-pod communication

## Service Accounts

Each service has its own dedicated service account:

- `books-service-account`: Basic service account for the books service
- `storage-service-account`: Service account with image pull secrets for the storage service
- `order-processor-service-account`: Service account with cloud provider annotations for the order processor service
- `frontend-service-account`: Basic service account for the frontend service

## Security Contexts

Security contexts are applied at both pod and container levels:

### Pod-Level Security Context

```yaml
securityContext:
  fsGroup: 2000
  runAsNonRoot: true
```

### Container-Level Security Context

```yaml
securityContext:
  runAsNonRoot: true
  runAsUser: 1000  # Or 101 for nginx-based containers
  allowPrivilegeEscalation: false
  readOnlyRootFilesystem: true
  capabilities:
    drop: ["ALL"]
    add: ["NET_BIND_SERVICE"]  # Only where needed
```

## RBAC (Role-Based Access Control)

RBAC is implemented using roles, role bindings, cluster roles, and cluster role bindings:

### Books Service

- `books-reader` role: Allows reading pods, services, and deployments
- `books-reader-binding`: Binds the role to the books service account

### Storage Service

- `storage-manager` cluster role: Allows managing persistent volumes and claims
- `storage-manager-binding`: Binds the cluster role to the storage service account

## Network Policies

Network policies control pod-to-pod communication:

### Default Deny Policy

A default deny policy blocks all traffic by default:

```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny-all
spec:
  podSelector: {}  # Selects all pods
  policyTypes:
  - Ingress
  - Egress
```

### Service-Specific Network Policies

Each service has its own network policy that allows only necessary traffic:

- `books-network-policy`: Allows ingress from frontend and egress to storage service
- `storage-network-policy`: Allows ingress from books and order processor services
- `order-processor-network-policy`: Allows ingress from books service and egress to storage service
- `frontend-network-policy`: Allows ingress from any source (typically ingress controller) and egress to books service

## Applying Security Features

To apply all security features:

```bash
# Apply service accounts
kubectl apply -f books/books-service-account.yaml
kubectl apply -f storage/storage-service-account.yaml
kubectl apply -f order-processing/order-processor-service-account.yaml
kubectl apply -f frontend/frontend-service-account.yaml

# Apply RBAC
kubectl apply -f books/books-reader-role.yaml
kubectl apply -f books/books-reader-binding.yaml
kubectl apply -f storage/storage-manager-role.yaml
kubectl apply -f storage/storage-manager-binding.yaml

# Apply network policies (default deny first)
kubectl apply -f default-deny-policy.yaml
kubectl apply -f books/books-network-policy.yaml
kubectl apply -f storage/storage-network-policy.yaml
kubectl apply -f order-processing/order-processor-network-policy.yaml
kubectl apply -f frontend/frontend-network-policy.yaml

# Apply deployments
kubectl apply -f books/deployment.yaml
kubectl apply -f storage/deployment.yaml
kubectl apply -f order-processing/deployment.yaml
kubectl apply -f frontend/deployment.yaml
```

## Security Best Practices

The following security best practices are implemented:

1. **Principle of Least Privilege**: Services have only the permissions they need
2. **Non-Root Containers**: All containers run as non-root users
3. **Read-Only Root Filesystem**: Prevents modifications to the container filesystem
4. **Capability Management**: Unnecessary capabilities are dropped
5. **Network Segmentation**: Network policies restrict pod-to-pod communication
6. **Dedicated Service Accounts**: Each service has its own service account
7. **RBAC**: Fine-grained access control for API resources

## Verifying Security Settings

To verify the security settings:

```bash
# Check service accounts
kubectl get serviceaccounts

# Check RBAC
kubectl get roles,rolebindings,clusterroles,clusterrolebindings

# Check network policies
kubectl get networkpolicies

# Check security contexts in running pods
kubectl get pod <pod-name> -o jsonpath='{.spec.securityContext}'
kubectl get pod <pod-name> -o jsonpath='{.spec.containers[0].securityContext}'
