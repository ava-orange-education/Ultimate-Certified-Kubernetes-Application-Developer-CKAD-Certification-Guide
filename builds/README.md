# AvaKart Build Configurations

This directory contains centralized Dockerfiles and Kubernetes manifests for deploying AvaKart.

## Structure

- `dockerfiles/` - Contains Dockerfiles for each AvaKart service
  - `Dockerfile.books` - Books service multi-stage build
  - `Dockerfile.frontend` - Frontend service with Nginx
  - `Dockerfile.order-processing` - Order processor service
  - `Dockerfile.storage` - Storage service
- `deployments/`
  - `k8s/` - Kubernetes manifests for each service
  - `docker-compose.yaml` - Local development setup

## Building Services

### Backend Services (Go)

All backend services use multi-stage builds with Go 1.21:

```bash
# Books Service
docker build -t books-service:latest -f dockerfiles/Dockerfile.books ../

# Order Processor
docker build -t order-processor:latest -f dockerfiles/Dockerfile.order-processing ../

# Storage Service
docker build -t storage-service:latest -f dockerfiles/Dockerfile.storage ../
```

### Frontend Service (React)

```bash
# Frontend with Nginx
docker build -t frontend:latest -f dockerfiles/Dockerfile.frontend ../
```

## Local Development

Use Docker Compose for local development environment:

```bash
docker-compose -f deployments/docker-compose.yaml up
```

Services will be available at:
- Frontend: http://localhost:3000
- Books Service: http://localhost:8081
- Order Processor: http://localhost:8082
- Storage Service: http://localhost:8083

## Kubernetes Deployment

Each service has its own Kubernetes manifests in `deployments/k8s/<service>/`:

1. Create ConfigMaps for environment variables:
```bash
kubectl create configmap books-config --from-env-file=../backend/apps/books/.env
kubectl create configmap order-config --from-env-file=../backend/apps/order-processor/.env
kubectl create configmap storage-config --from-env-file=../backend/apps/storage/.env
kubectl create configmap frontend-config --from-env-file=../frontend/.env
```

2. Deploy services:
```bash
kubectl apply -f deployments/k8s/storage/
kubectl apply -f deployments/k8s/books/
kubectl apply -f deployments/k8s/order-processing/
kubectl apply -f deployments/k8s/frontend/
```

3. Verify deployments:
```bash
kubectl get pods
kubectl get services
```

## Configuration

### Resource Requirements

Each service has defined resource limits:

```yaml
resources:
  requests:
    cpu: "100m"
    memory: "128Mi"
  limits:
    cpu: "200m"
    memory: "256Mi"
```

### Security Context

Backend services run with security constraints:

```yaml
securityContext:
  runAsNonRoot: true
  runAsUser: 1000
  allowPrivilegeEscalation: false
```

### Health Checks

All services implement health endpoints:
- Backend services: `/health`
- Frontend: `/` with Nginx stub_status
