# AvaKart Order Processing Service

The order processing service for AvaKart, handling order lifecycle events from submission to fulfillment.

## Architecture

The service follows a layered architecture:
- `handlers` - HTTP endpoints for order management
- `services` - Core business logic for order processing
- `repository` - Data persistence layer
- `models` - Data structures and storage service integration

## Features

- Manages order status and coordinates with backend and storage services
- Integrates with storage service for inventory management
- Handles order persistence and state transitions
- Designed to support scaling for high-volume order handling

## Development

### Prerequisites
- Go 1.23.3
- Access to storage service

### Project Structure
```
.
├── handlers/
│   └── order.go       # HTTP endpoints
├── models/
│   ├── order.go       # Order data structures
│   └── storage-svc.go # Storage service integration
├── repository/
│   └── orders.go      # Data persistence
├── services/
│   └── order-processor.go # Business logic
├── main.go            # Application entry point
└── README.md
```

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
