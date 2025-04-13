# AvaKart Helm Charts

This directory contains Helm charts for deploying AvaKart components to Kubernetes.

## Charts

- `books` - Books Service
- `frontend` - Frontend Service
- `order-processor` - Order Processor Service
- `storage` - Storage Service

## Prerequisites

- Kubernetes 1.19+
- Helm 3.2.0+

## Installation

### Add the Helm Repository (Optional)

If you're hosting these charts in a repository:

```bash
helm repo add avakart https://your-helm-repo-url
helm repo update
```

### Install Individual Charts

```bash
# Install Storage Service (dependency for other services)
helm install storage ./storage

# Install Books Service
helm install books ./books

# Install Order Processor
helm install order-processor ./order-processor

# Install Frontend
helm install frontend ./frontend
```

### Install with Custom Values

```bash
helm install storage ./storage -f custom-values.yaml
```

## Configuration

Each chart has its own `values.yaml` file with default configurations. You can override these values by creating your own values file or using the `--set` flag.

### Common Configuration Parameters

#### Image Configuration

```yaml
image:
  repository: your-registry/image-name
  tag: latest
  pullPolicy: IfNotPresent
```

#### Resource Limits

```yaml
resources:
  requests:
    cpu: 100m
    memory: 128Mi
  limits:
    cpu: 200m
    memory: 256Mi
```

#### Persistence

```yaml
persistence:
  enabled: true
  storageClassName: standard
  accessMode: ReadWriteOnce
  size: 1Gi
```

### Service-Specific Configuration

#### Frontend

```yaml
ingress:
  enabled: true
  hosts:
    - host: avakart.example.com
      paths:
        - path: /
          pathType: Prefix
```

#### Order Processor

```yaml
cronJob:
  enabled: true
  schedule: "0 * * * *"
```

## Upgrading

```bash
helm upgrade storage ./storage
```

## Uninstalling

```bash
helm uninstall storage
```

## Development

### Testing Charts Locally

```bash
# Validate chart
helm lint ./storage

# Render templates
helm template storage ./storage

# Dry run
helm install storage ./storage --dry-run
```

### Packaging Charts

```bash
helm package ./storage
