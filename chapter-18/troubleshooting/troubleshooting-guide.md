# CRD Troubleshooting Guide

This guide provides common troubleshooting scenarios and solutions for Custom Resource Definitions (CRDs) and custom resources.

## Common Issues and Solutions

### 1. CRD Not Found Error

**Error Message:**
```
error validating data: ValidationError(BookCatalog): unknown field "spec" in io.avakart.v1.BookCatalog
```

**Cause:** The CRD is not installed in the cluster.

**Solution:**
```bash
# Check if CRD exists
kubectl get crd | grep avakart

# Install the CRD
kubectl apply -f crds/bookcatalog-crd.yaml

# Verify CRD installation
kubectl describe crd bookcatalogs.avakart.io
```

### 2. Validation Errors

**Error Message:**
```
error validating data: ValidationError(BookCatalog.spec.price): invalid value: -10, must be greater than or equal to 0
```

**Cause:** The custom resource violates the OpenAPI schema validation rules.

**Common Validation Issues:**
- Required fields missing
- Values outside min/max ranges
- Invalid enum values
- String pattern mismatches
- Wrong data types

**Solution:**
```bash
# Use dry-run to test validation
kubectl apply -f invalid-book.yaml --dry-run=server

# Check the CRD schema
kubectl get crd bookcatalogs.avakart.io -o yaml | grep -A 50 openAPIV3Schema

# Fix the resource and reapply
kubectl apply -f fixed-book.yaml
```

### 3. Custom Resource Stuck in Pending

**Symptoms:**
- Custom resource created but status never updates
- No operator logs or activity

**Possible Causes:**
- Operator not running
- RBAC permissions missing
- Controller not watching the resource

**Troubleshooting Steps:**
```bash
# Check if operator is running
kubectl get pods -n avakart-system -l app=avakart-operator

# Check operator logs
kubectl logs -n avakart-system deployment/avakart-operator

# Verify RBAC permissions
kubectl auth can-i get bookcatalogs --as=system:serviceaccount:avakart-system:avakart-operator

# Check events
kubectl get events --field-selector involvedObject.kind=BookCatalog
```

### 4. Status Not Updating

**Symptoms:**
- Custom resource exists but status field remains empty
- Operator logs show no errors

**Possible Causes:**
- Status subresource not enabled
- Operator lacks status update permissions
- Controller logic issues

**Solution:**
```bash
# Check if status subresource is enabled
kubectl get crd bookcatalogs.avakart.io -o jsonpath='{.spec.versions[0].subresources}'

# Verify status update permissions
kubectl auth can-i update bookcatalogs/status --as=system:serviceaccount:avakart-system:avakart-operator

# Manually update status for testing
kubectl patch bookcatalog kubernetes-guide --type='merge' --subresource=status -p='{"status":{"phase":"available"}}'
```

### 5. Resource Deletion Issues

**Symptoms:**
- Custom resource stuck in "Terminating" state
- Resource won't delete

**Possible Causes:**
- Finalizers preventing deletion
- Dependent resources not cleaned up
- Operator not handling deletion

**Solution:**
```bash
# Check for finalizers
kubectl get bookcatalog kubernetes-guide -o jsonpath='{.metadata.finalizers}'

# Remove finalizers if operator is not working
kubectl patch bookcatalog kubernetes-guide --type='json' -p='[{"op": "remove", "path": "/metadata/finalizers"}]'

# Force delete if necessary (use with caution)
kubectl delete bookcatalog kubernetes-guide --force --grace-period=0
```

## Debugging Commands

### CRD Information
```bash
# List all CRDs
kubectl get crd

# Get CRD details
kubectl describe crd bookcatalogs.avakart.io

# Check CRD versions
kubectl get crd bookcatalogs.avakart.io -o jsonpath='{.spec.versions[*].name}'

# View CRD schema
kubectl get crd bookcatalogs.avakart.io -o yaml
```

### Custom Resource Operations
```bash
# List custom resources
kubectl get bookcatalogs -A

# Get custom resource with custom columns
kubectl get bookcatalogs -o wide

# Describe custom resource
kubectl describe bookcatalog kubernetes-guide

# Get custom resource YAML
kubectl get bookcatalog kubernetes-guide -o yaml

# Check resource status
kubectl get bookcatalog kubernetes-guide -o jsonpath='{.status}'
```

### Events and Logs
```bash
# Get events for custom resources
kubectl get events --field-selector involvedObject.kind=BookCatalog

# Get events for specific resource
kubectl get events --field-selector involvedObject.name=kubernetes-guide

# Watch events in real-time
kubectl get events --watch

# Get operator logs
kubectl logs -f deployment/avakart-operator -n avakart-system
```

### Validation Testing
```bash
# Test resource validation with dry-run
kubectl apply -f resource.yaml --dry-run=server

# Validate against schema
kubectl apply -f resource.yaml --validate=true

# Check resource differences
kubectl diff -f resource.yaml
```

## CKAD Exam Scenarios

### Scenario 1: Working with Existing CRDs
```bash
# Task: Find and work with existing CRDs in the cluster
kubectl get crd
kubectl describe crd <crd-name>
kubectl get <custom-resource-type> -A
```

### Scenario 2: Creating Custom Resources
```bash
# Task: Create a custom resource from given specification
kubectl apply -f custom-resource.yaml
kubectl get <custom-resource-type> <name> -o yaml
```

### Scenario 3: Troubleshooting Validation Issues
```bash
# Task: Fix validation errors in custom resources
kubectl apply -f resource.yaml --dry-run=server
# Fix the issues and reapply
kubectl apply -f resource.yaml
```

### Scenario 4: Understanding Resource Relationships
```bash
# Task: Find resources managed by operators
kubectl get all -l app.kubernetes.io/managed-by=<operator-name>
kubectl describe <resource-type> <name>
```

## Best Practices

1. **Always use dry-run first**
   ```bash
   kubectl apply -f resource.yaml --dry-run=server
   ```

2. **Check CRD installation before creating resources**
   ```bash
   kubectl get crd | grep <api-group>
   ```

3. **Use describe for detailed information**
   ```bash
   kubectl describe crd <crd-name>
   kubectl describe <custom-resource> <name>
   ```

4. **Monitor events for troubleshooting**
   ```bash
   kubectl get events --watch
   ```

5. **Check operator logs when resources aren't processing**
   ```bash
   kubectl logs -f deployment/<operator-name>
   ```

6. **Use labels and selectors for resource management**
   ```bash
   kubectl get <custom-resource> -l <label-selector>
   ```

7. **Understand finalizers and their impact**
   ```bash
   kubectl get <resource> -o jsonpath='{.metadata.finalizers}'
   ```

## Quick Reference

### Essential Commands
```bash
# CRD operations
kubectl get crd
kubectl describe crd <name>
kubectl delete crd <name>

# Custom resource operations
kubectl get <custom-resource-type>
kubectl describe <custom-resource-type> <name>
kubectl delete <custom-resource-type> <name>

# Validation and debugging
kubectl apply --dry-run=server -f <file>
kubectl get events --field-selector involvedObject.kind=<Kind>
kubectl logs <operator-pod>
```

### Common Patterns
```bash
# Find all custom resources in cluster
kubectl api-resources --api-group=<group>

# Get custom resources across all namespaces
kubectl get <custom-resource-type> -A

# Update custom resource status
kubectl patch <custom-resource-type> <name> --subresource=status --type='merge' -p='<status-json>'
