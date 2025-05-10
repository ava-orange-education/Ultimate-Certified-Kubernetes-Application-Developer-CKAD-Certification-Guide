# AvaKart Kubernetes Manifests

This directory contains Kubernetes manifests for deploying the AvaKart application.

## Directory Structure

- `books/` - Manifests for the Books Service
- `frontend/` - Manifests for the Frontend Service
- `order-processing/` - Manifests for the Order Processor Service
- `storage/` - Manifests for the Storage Service
- `avakart-tls.yaml` - TLS configuration for AvaKart
- `debug-avakart.sh` - Script for debugging AvaKart
- `debug-pod.yaml` - Debug pod configuration
- `storage-class.yaml` - Storage class configuration
- `default-deny-policy.yaml` - Default deny network policy
- `security-README.md` - Documentation on security features
- `apply-security.sh` - Script to apply security manifests
- `cleanup-security.sh` - Script to clean up security manifests

## Service Components

### Books Service

The Books Service provides the API for managing books in the AvaKart application.

- Deployment: `books/deployment.yaml`
- Service Account: `books/books-service-account.yaml`
- RBAC: `books/books-reader-role.yaml`, `books/books-reader-binding.yaml`
- Network Policy: `books/books-network-policy.yaml`

### Storage Service

The Storage Service provides storage functionality for the AvaKart application.

- Deployment: `storage/deployment.yaml`
- Service Account: `storage/storage-service-account.yaml`
- RBAC: `storage/storage-manager-role.yaml`, `storage/storage-manager-binding.yaml`
- Network Policy: `storage/storage-network-policy.yaml`

### Order Processor Service

The Order Processor Service handles order processing for the AvaKart application.

- Deployment: `order-processing/deployment.yaml`
- Service Account: `order-processing/order-processor-service-account.yaml`
- Network Policy: `order-processing/order-processor-network-policy.yaml`

### Frontend Service

The Frontend Service provides the user interface for the AvaKart application.

- Deployment: `frontend/deployment.yaml`
- Service Account: `frontend/frontend-service-account.yaml`
- Network Policy: `frontend/frontend-network-policy.yaml`

## Security Features

The AvaKart application has been secured using the following Kubernetes security features:

1. **Service Accounts**: Dedicated service accounts for each service
2. **Security Contexts**: Pod and container level security settings
3. **RBAC**: Role-Based Access Control for fine-grained permissions
4. **Network Policies**: Control pod-to-pod communication

For detailed information on the security features, see `security-README.md`.

## Deployment

To deploy the AvaKart application with security features:

```bash
# Apply security manifests
./apply-security.sh
```

To clean up the AvaKart application:

```bash
# Clean up security manifests
./cleanup-security.sh
```

## Debugging

For debugging the AvaKart application:

```bash
# Run the debug script
./debug-avakart.sh
```

## TLS Configuration

The AvaKart application uses TLS for secure communication. The TLS configuration is defined in `avakart-tls.yaml`.
