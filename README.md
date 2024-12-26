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
- Go 1.23.3 for backend services
- Node.js 18+ for frontend development
- MongoDB for storage service

### Local Development
1. Clone the repository
2. Navigate to the project root
3. Run `docker-compose -f builds/deployments/docker-compose.yaml up` to start all services

### Kubernetes Deployment
The `builds/deployments/k8s/` directory contains service-specific manifests:
- ConfigMaps for service configuration
- Deployment specifications
- Service definitions for networking

Each service can be deployed using:
```bash
kubectl apply -f builds/deployments/k8s/<service-name>/
```
