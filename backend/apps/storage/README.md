# Storage Service

The storage service is responsible for managing persistent data storage for the book inventory system. It provides a centralized storage solution for book details and quantity management.

## Features

- Book Details Management
  - Store and retrieve book information
  - Update book details
  - Delete book records

- Book Quantity Management
  - Track book inventory levels
  - Update book quantities
  - Handle inventory adjustments

## API Endpoints

### Book Details
- `GET /books/{id}` - Retrieve book details by ID
- `POST /books` - Create new book record
- `PUT /books/{id}` - Update book details
- `DELETE /books/{id}` - Remove book record

### Book Quantity
- `GET /books/{id}/quantity` - Get current book quantity
- `PUT /books/{id}/quantity` - Update book quantity

## Development

### Prerequisites
- Go 1.23.3
- Access to a MongoDB instance (configured via environment variables)

### Running Locally
1. Set required environment variables
2. Run `go mod download` to install dependencies
3. Execute `go run main.go` to start the service

### Docker Build

**Not yet implemeted**

Build the service using the provided Dockerfile:
```bash
docker build -f ../../builds/dockerfiles/Dockerfile.storage -t storage.svc.avakart .
```
### Kubernetes Deployment

**Not yet implemeted**

The service can be deployed using Kubernetes:
```bash
kubectl apply -f builds/deployments/k8s/storage/
```

Configuration is managed through Kubernetes ConfigMaps and can be customized by modifying:
- `builds/deployments/k8s/storage/configmap.yaml`
- `builds/deployments/k8s/storage/deployment.yaml`
- `builds/deployments/k8s/storage/service.yaml`

## Architecture

The service follows a clean architecture pattern:
- `handlers/` - HTTP request handlers
- `services/` - Business logic layer
- `repository/` - Data persistence layer
- `models/` - Data models and DTOs
