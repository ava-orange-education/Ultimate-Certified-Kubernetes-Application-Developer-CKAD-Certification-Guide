# Ultimate Certified Kubernetes Application Developer (CKAD) Certification Guide

This repository contains the resources, source code, and deployment configurations for a practical e-commerce bookstore (**AvaKart**) project that complements the **Certified Kubernetes Application Developer (CKAD)** certification guide.

## Overview

The project implements a microservices-based bookstore (AvaKart) with the following components:

- **Books Service** (`:8081`): 
  - Manages book catalog and metadata
  - Handles book purchase initiation
  - Provides RESTful APIs for book operations
  - Integrates with storage and order processing

- **Order Processor** (`:8082`):
  - Manages complete order lifecycle
  - Handles order status transitions
  - Coordinates with storage for persistence
  - Provides order fulfillment workflow

- **Storage Service** (`:8083`):
  - Centralized data persistence
  - Book inventory management
  - Order data storage
  - Internal APIs for service integration

- **Frontend**:
  - React-based user interface
  - Book browsing and ordering (currently limited to quantity of 1 per purchase)
  - Real-time inventory display
  - Order status tracking
  - Future enhancements planned:
    - Configurable purchase quantities
    - Order cancellation functionality

These services demonstrate key CKAD concepts including:
- Multi-container deployments with service discovery
- Inter-service communication patterns
- ConfigMap usage for service configuration
- Container image building and optimization
- Kubernetes deployment strategies and scaling
- Service mesh integration capabilities

## Directory Structure

- `backend/apps/` - Go-based microservices
  - `books/` - Book catalog service
  - `storage/` - Inventory management service
  - `order-processor/` - Order processing service
- `frontend/` - React-based web interface
- `builds/`
  - `dockerfiles/` - Centralized Dockerfiles for each service
  - `deployments/`
    - `k8s/` - Kubernetes manifests organized by service
    - `docker-compose.yaml` - Local development environment setup
- `scripts/` - Utility scripts for setup and cleanup

## Getting Started

### Prerequisites
- Docker and Docker Compose (v2.0+) for local development
- Kubernetes cluster (v1.24+) for deployment
- Go 1.21 for backend services
- Node.js 18+ for frontend development
- MongoDB for storage service

### Environment Variables

Each service requires specific environment variables for proper configuration:

#### Books Service (8081)
- `STORAGE_SERVICE_URL` - URL of the storage service
- `ORDER_PROCESSOR_URL` - URL of the order processor service
- `SERVER_PORT` - Service port (default: 8081)

#### Order Processor (8082)
- `STORAGE_SERVICE_URL` - URL of the storage service
- `SERVER_PORT` - Service port (default: 8082)

#### Storage Service (8083)
- `MONGODB_URI` - MongoDB connection string
- `SERVER_PORT` - Service port (default: 8083)

#### Frontend
- `VITE_BOOKS_SERVICE_URL` - URL of the books service
- `VITE_ORDER_PROCESSOR_URL` - URL of the order processor service

### Local Development
1. Clone the repository
2. Navigate to the project root
3. Set up environment variables in a `.env` file
4. Run `docker-compose -f builds/deployments/docker-compose.yaml up` to start all services

### Kubernetes Deployment
The `builds/deployments/k8s/` directory contains service-specific manifests:
- ConfigMaps for service configuration
- Deployment specifications
- Service definitions for networking

Each service can be deployed using:
```bash
kubectl apply -f builds/deployments/k8s/<service-name>/
```

### Service Communication

The services communicate with each other using REST APIs:
- Frontend → Books Service: Book catalog and purchase operations
- Books Service → Storage Service: Inventory checks and updates
- Books Service → Order Processor: Order creation
- Order Processor → Storage Service: Order persistence and updates

### Troubleshooting

Common issues and solutions:

1. Service Connection Issues
   - Verify environment variables are correctly set
   - Check if services are running (`docker ps` or `kubectl get pods`)
   - Ensure network policies allow communication

2. MongoDB Connection
   - Verify MongoDB is running and accessible
   - Check MongoDB connection string format
   - Ensure database credentials are correct

3. Build Failures
   - Clear Docker build cache: `docker builder prune`
   - Update dependencies: `go mod tidy` or `npm install`
   - Check for Go/Node.js version compatibility

4. Kubernetes Deployment Issues
   - Verify ConfigMaps are properly created
   - Check pod logs: `kubectl logs <pod-name>`
   - Ensure images are properly tagged and accessible
