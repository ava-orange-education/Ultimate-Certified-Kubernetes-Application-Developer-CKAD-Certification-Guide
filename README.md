# Ultimate Certified Kubernetes Application Developer (CKAD) Certification Guide

This repository contains the resources, source code, and deployment configurations for a practical e-commerce bookstore project that complements the **Certified Kubernetes Application Developer (CKAD)** certification guide.

## Overview

The project implements a microservices-based bookstore with the following components:

- **Books Service**: Manages book catalog and metadata
- **Storage Service**: Handles inventory and book quantity management
- **Order Processor**: Processes customer orders and coordinates with storage
- **Frontend**: React-based user interface for browsing and ordering books

These services demonstrate key CKAD concepts including:
- Multi-container deployments
- Service networking and communication
- ConfigMap usage for configuration
- Container image building and management
- Kubernetes deployment strategies

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
- Docker and Docker Compose for local development
- Kubernetes cluster (local or remote) for deployment
- Go 1.23.3 for backend services
- Node.js for frontend development

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
