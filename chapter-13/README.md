# Chapter 13: Security Best Practices

This directory contains Kubernetes manifests that demonstrate security best practices for the AvaKart application. The focus is on implementing security at various levels within Kubernetes deployments, from container-level security to cluster-wide access controls.

## Directory Structure

- **service-accounts/**: Service account definitions for different services
  - `books-service-account.yaml`: Basic service account for the books service
  - `storage-service-account.yaml`: Service account with image pull secrets for the storage service
  - `order-processor-service-account.yaml`: Service account with annotations for the order processor service

- **security-contexts/**: Examples of security contexts at pod and container levels
  - `secure-books-pod.yaml`: Pod with pod-level security context
  - `secure-storage-pod.yaml`: Pod with container-level security context

- **rbac/**: Role-Based Access Control configurations
  - **roles/**: Namespace-scoped role definitions
    - `books-reader-role.yaml`: Role for reading pods, services, and deployments
  - **role-bindings/**: Bindings for namespace-scoped roles
    - `books-reader-binding.yaml`: Binding for the books-reader role
  - **cluster-roles/**: Cluster-wide role definitions
    - `storage-manager-role.yaml`: ClusterRole for managing persistent volumes
  - **cluster-role-bindings/**: Bindings for cluster-wide roles
    - `storage-manager-binding.yaml`: Binding for the storage-manager role

- **deployments/**: Secure deployment configurations
  - `secure-books-deployment.yaml`: Secure deployment for the books service
  - `secure-storage-deployment.yaml`: Secure deployment for the storage service

- **network-policies/**: Network policies for controlling pod-to-pod communication
  - `default-deny-policy.yaml`: Default deny policy for all pods
  - `books-network-policy.yaml`: Network policy for the books service
  - `storage-network-policy.yaml`: Network policy for the storage service
  - `README.md`: Documentation on network policies

## Security Features Implemented

1. **Service Accounts**: Provide identity for processes running in pods
2. **Security Contexts**: Define privilege and access control settings for pods and containers
3. **RBAC (Role-Based Access Control)**: Control access to Kubernetes API resources
4. **Principle of Least Privilege**: Run containers as non-root users with minimal capabilities
5. **Read-Only Root Filesystem**: Prevent modifications to the container filesystem
6. **Volume Management**: Properly configure volume mounts for temporary and persistent storage
7. **Network Policies**: Control pod-to-pod communication with ingress and egress rules

## Usage

These manifests can be applied to a Kubernetes cluster using the `kubectl apply` command:

```bash
kubectl apply -f chapter-13/service-accounts/
kubectl apply -f chapter-13/rbac/roles/
kubectl apply -f chapter-13/rbac/role-bindings/
kubectl apply -f chapter-13/rbac/cluster-roles/
kubectl apply -f chapter-13/rbac/cluster-role-bindings/
kubectl apply -f chapter-13/deployments/
kubectl apply -f chapter-13/network-policies/
```

Alternatively, you can use the provided script:

```bash
./chapter-13/apply-security.sh
```

## Best Practices

- Always run containers as non-root users
- Drop unnecessary capabilities and only add those that are required
- Use read-only root filesystems where possible
- Implement RBAC with the principle of least privilege
- Create dedicated service accounts for each application
- Use security contexts to enforce security policies
- Implement network policies to control pod-to-pod communication
- Start with a default deny policy and then add specific allow rules

## Documentation

- `security-best-practices.md`: Detailed explanation of Kubernetes security best practices
- `verify-security.md`: Commands and procedures to verify security configurations
- `network-policies/README.md`: Documentation on network policies
- `exam-prep-questions.md`: Sample questions and answers for CKAD exam preparation
