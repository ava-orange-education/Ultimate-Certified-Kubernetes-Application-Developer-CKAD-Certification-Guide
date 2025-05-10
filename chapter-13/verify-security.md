# Verifying Kubernetes Security Configurations

This document provides commands and procedures to verify the security configurations implemented in the Chapter 13 manifests.

## 1. Verifying Service Accounts

```bash
# List all service accounts in the default namespace
kubectl get serviceaccounts

# Describe a specific service account to see details
kubectl describe serviceaccount books-service-account
kubectl describe serviceaccount storage-service-account
kubectl describe serviceaccount order-processor-service-account

# Check if a pod is using the correct service account
kubectl get pod <pod-name> -o jsonpath='{.spec.serviceAccountName}'
```

## 2. Verifying Security Contexts

```bash
# Check pod-level security context
kubectl get pod secure-books-pod -o jsonpath='{.spec.securityContext}'

# Check container-level security context
kubectl get pod secure-storage-pod -o jsonpath='{.spec.containers[0].securityContext}'

# Check security context in a deployment
kubectl get deployment secure-books-deployment -o jsonpath='{.spec.template.spec.securityContext}'
kubectl get deployment secure-books-deployment -o jsonpath='{.spec.template.spec.containers[0].securityContext}'
```

## 3. Verifying RBAC Configurations

```bash
# List all roles in the default namespace
kubectl get roles

# Describe a specific role to see its rules
kubectl describe role books-reader

# List all role bindings in the default namespace
kubectl get rolebindings

# Describe a specific role binding to see its subjects and role ref
kubectl describe rolebinding books-reader-binding

# List all cluster roles
kubectl get clusterroles | grep storage-manager

# Describe a specific cluster role
kubectl describe clusterrole storage-manager

# List all cluster role bindings
kubectl get clusterrolebindings | grep storage-manager

# Describe a specific cluster role binding
kubectl describe clusterrolebinding storage-manager-binding
```

## 4. Verifying Permissions

```bash
# Check if a service account can perform specific actions
kubectl auth can-i get pods --as=system:serviceaccount:default:books-service-account
kubectl auth can-i list services --as=system:serviceaccount:default:books-service-account
kubectl auth can-i create deployments --as=system:serviceaccount:default:books-service-account

# Check if a service account can access persistent volumes
kubectl auth can-i get persistentvolumes --as=system:serviceaccount:default:storage-service-account
kubectl auth can-i create persistentvolumeclaims --as=system:serviceaccount:default:storage-service-account
```

## 5. Verifying Pod Security

```bash
# Check if a pod is running as non-root
kubectl exec secure-books-pod -- id

# Check if a pod has the expected capabilities
kubectl exec secure-storage-pod -- capsh --print

# Check if a pod's root filesystem is read-only
kubectl exec secure-storage-pod -- touch /test.txt
# This should fail with a permission denied error if the root filesystem is read-only

# Check if a pod can write to its mounted volumes
kubectl exec secure-storage-pod -- touch /tmp/test.txt
# This should succeed as /tmp is an emptyDir volume
```

## 6. Using Tools for Security Auditing

### Kube-bench

[kube-bench](https://github.com/aquasecurity/kube-bench) is a tool that checks whether Kubernetes is deployed securely by running the checks documented in the CIS Kubernetes Benchmark.

```bash
# Run kube-bench
kubectl apply -f https://raw.githubusercontent.com/aquasecurity/kube-bench/main/job.yaml
kubectl logs -f job.batch/kube-bench
```

### Trivy

[Trivy](https://github.com/aquasecurity/trivy) is a comprehensive security scanner that can scan container images for vulnerabilities.

```bash
# Install Trivy
# For Ubuntu/Debian
apt-get install trivy

# For macOS
brew install aquasecurity/trivy/trivy

# Scan a container image
trivy image books-service:v1
```

### Kubeaudit

[kubeaudit](https://github.com/Shopify/kubeaudit) is a command-line tool that audits Kubernetes clusters for various security concerns.

```bash
# Install kubeaudit
# For macOS
brew install kubeaudit

# For other platforms, download from GitHub releases
# https://github.com/Shopify/kubeaudit/releases

# Audit all resources in the default namespace
kubeaudit all -n default
```

## 7. Continuous Security Monitoring

For continuous security monitoring, consider implementing:

1. **Prometheus and Grafana**: For monitoring and alerting on security metrics
2. **Falco**: For runtime security monitoring
3. **Open Policy Agent (OPA)**: For policy enforcement
4. **Audit Logging**: Enable and analyze Kubernetes audit logs

## Conclusion

Regular verification of security configurations is essential to ensure that your Kubernetes cluster remains secure. Use the commands and tools provided in this document to verify the security configurations implemented in the Chapter 13 manifests.
