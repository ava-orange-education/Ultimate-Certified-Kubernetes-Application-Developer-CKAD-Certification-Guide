# Network Policies for AvaKart Application

This directory contains network policy examples for the AvaKart application. Network policies are a key component of Kubernetes security, allowing you to control the traffic flow between pods.

## Network Policies Overview

Network policies in Kubernetes:
- Define how pods can communicate with each other and other network endpoints
- Are implemented by the network plugin (e.g., Calico, Cilium, Weave Net)
- Follow a deny-by-default approach when policies are applied
- Are namespace-scoped resources

## Included Network Policies

### 1. Default Deny Policy (`default-deny-policy.yaml`)

This policy denies all ingress and egress traffic for all pods in the namespace. It serves as a baseline security measure, ensuring that only explicitly allowed traffic can flow.

Key features:
- Selects all pods in the namespace with an empty pod selector (`{}`)
- Denies all ingress and egress traffic by not specifying any allow rules
- Provides a secure foundation for more specific policies

### 2. Books Service Network Policy (`books-network-policy.yaml`)

This policy controls traffic to and from the Books Service pods.

Ingress rules:
- Allow incoming traffic only from pods with the label `app: frontend`
- Allow traffic only on port 8081 (Books Service API port)

Egress rules:
- Allow outgoing traffic only to pods with the label `app: storage-service`
- Allow traffic only on port 8083 (Storage Service API port)
- Allow DNS resolution by permitting traffic to kube-dns on ports 53/TCP and 53/UDP

### 3. Storage Service Network Policy (`storage-network-policy.yaml`)

This policy controls traffic to and from the Storage Service pods.

Ingress rules:
- Allow incoming traffic from pods with labels `app: books-service` or `app: order-processor`
- Allow traffic only on port 8083 (Storage Service API port)

Egress rules:
- Allow DNS resolution by permitting traffic to kube-dns on ports 53/TCP and 53/UDP
- Allow outgoing traffic to the database subnet (10.0.0.0/24) on port 5432 (PostgreSQL)

## Implementation Strategy

The network policies in this directory follow a layered approach:

1. **Default Deny**: Start with a default deny policy to block all traffic
2. **Selective Allow**: Create specific policies for each service to allow only necessary traffic
3. **Least Privilege**: Follow the principle of least privilege by allowing only required ports and protocols

## Applying Network Policies

Apply these network policies using kubectl:

```bash
# Apply default deny policy first
kubectl apply -f network-policies/default-deny-policy.yaml

# Apply service-specific network policies
kubectl apply -f network-policies/books-network-policy.yaml
kubectl apply -f network-policies/storage-network-policy.yaml
```

## Verifying Network Policies

Verify that the network policies are working correctly:

```bash
# List all network policies
kubectl get networkpolicies

# Describe a specific network policy
kubectl describe networkpolicy books-network-policy

# Test connectivity between pods
kubectl exec -it <source-pod> -- curl <destination-pod-ip>:<port>
```

## Best Practices

1. **Start with Default Deny**: Begin with a default deny policy and then add specific allow rules
2. **Be Specific**: Use specific pod selectors and namespace selectors rather than broad rules
3. **Allow DNS**: Always allow DNS traffic (port 53 UDP/TCP) to kube-dns
4. **Test Thoroughly**: Test connectivity between pods after applying network policies
5. **Document Policies**: Document the purpose and rules of each network policy
6. **Regular Review**: Regularly review and update network policies as application requirements change
