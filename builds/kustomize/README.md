# Kustomize Templates for Bookstore Application

This directory contains Kustomize templates for deploying the Bookstore application components to Kubernetes.

## Components

The application consists of the following components:

1. **Books Service**: Backend service for book management
2. **Frontend**: User interface for the bookstore
3. **Order Processor**: Service for processing book orders
4. **Storage Service**: Service for storing book data and order information

## Directory Structure

```
kustomize/
├── base/                  # Base configurations for all components
│   ├── books/             # Books service base configuration
│   ├── frontend/          # Frontend base configuration
│   ├── order-processor/   # Order processor base configuration
│   └── storage/           # Storage service base configuration
└── overlays/              # Environment-specific overlays
    ├── dev/               # Development environment
    ├── staging/           # Staging environment
    └── prod/              # Production environment
```

## Usage

### Prerequisites

- Kubernetes cluster
- kubectl installed
- kustomize installed (or use kubectl built-in kustomize)

### Deploying to Different Environments

#### Development Environment

```bash
# Preview the resources
kubectl kustomize builds/kustomize/overlays/dev

# Apply the resources
kubectl apply -k builds/kustomize/overlays/dev
```

#### Staging Environment

```bash
# Preview the resources
kubectl kustomize builds/kustomize/overlays/staging

# Apply the resources
kubectl apply -k builds/kustomize/overlays/staging
```

#### Production Environment

```bash
# Preview the resources
kubectl kustomize builds/kustomize/overlays/prod

# Apply the resources
kubectl apply -k builds/kustomize/overlays/prod
```

## Environment Differences

### Development
- Single replica for each component
- Uses `dev` tagged images
- Deployed to `bookstore-dev` namespace

### Staging
- Two replicas for each component
- Uses `staging` tagged images
- Deployed to `bookstore-staging` namespace

### Production
- Three replicas for each component
- Increased resource requests and limits
- Uses `prod` tagged images
- Deployed to `bookstore-prod` namespace

## Customization

To customize the deployment for a specific environment, modify the corresponding overlay kustomization.yaml file. You can:

- Change the number of replicas
- Adjust resource requests and limits
- Update image tags
- Add or modify environment variables
- Add additional resources or patches
