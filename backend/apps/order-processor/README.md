# AvaKart Order Processing Service

The order processing service for AvaKart, handling order lifecycle events from submission to fulfillment.

## Architecture

The service follows a layered architecture:
- `handlers` - HTTP endpoints for order management
- `services` - Core business logic for order processing
- `repository` - Data persistence layer
- `models` - Data structures and storage service integration

## Features

- Complete order lifecycle management:
  - Order creation with automatic ID generation
  - Order status tracking and updates
  - Order fulfillment coordination
- Integration with storage service for:
  - Inventory verification
  - Order persistence
  - Status synchronization
- Robust error handling and validation
- Designed for high-volume order processing

## Configuration

The service can be configured using environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | The port number the service listens on | 8082 |
| STORAGE_SERVICE_URL | URL of the storage service | http://localhost:8083 |

## Development

### Prerequisites
- Go 1.23.3
- Access to storage service

## API Endpoints

### Order Management
- `POST /orders/create` - Create a new order
  - Validates book availability
  - Generates unique order ID
  - Initializes order status
- `PUT /orders/update-status` - Update order status
  - Manages order state transitions
  - Synchronizes with storage service
  - Handles fulfillment workflow

## Project Structure
```
.
├── handlers/
│   └── order.go           # HTTP endpoint handlers
├── models/
│   ├── order.go           # Order data structures
│   └── storage-svc.go     # Storage service integration
├── services/
│   └── order-processor.go # Core business logic
├── main.go                # Service entry point
└── README.md
```

## Service Integration

The service operates on port `:8082` and integrates with:
- Storage Service (`:8083`) for:
  - Inventory verification
  - Order persistence
  - Status management
- Books Service (`:8081`) for:
  - Order initiation
  - Book availability checks

### Docker Build

**Not yet implemeted**

Build the service using the provided Dockerfile:
```bash
docker build -f ../../builds/dockerfiles/Dockerfile.order-processing -t order-processing.svc.avakart .
```
### Kubernetes Deployment

**Not yet implemeted**

The service can be deployed using Kubernetes:
```bash
kubectl apply -f builds/deployments/k8s/order-processing/
```

Configuration is managed through Kubernetes ConfigMaps and can be customized by modifying:
- `builds/deployments/k8s/order-processing/configmap.yaml`
- `builds/deployments/k8s/order-processing/deployment.yaml`
- `builds/deployments/k8s/order-processing/service.yaml`
