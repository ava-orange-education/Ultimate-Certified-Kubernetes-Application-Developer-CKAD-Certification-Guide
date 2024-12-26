# Order Processor Service

Go-based service for processing orders in the bookstore application.

## API Endpoints

### Order Management
- `POST /orders` - Create new order
  ```json
  {
    "bookId": "string",
    "quantity": 1,
    "userId": "string"
  }
  ```
- `GET /orders/{id}` - Get order status
- `GET /orders/user/{userId}` - List user orders

### Health Check
- `GET /health` - Service health status

## Development

### Prerequisites
- Go 1.21
- Storage Service running on port 8083

### Environment Variables
```
STORAGE_SERVICE_URL=http://localhost:8083
SERVER_PORT=8082
```

### Running Locally
```bash
go mod download
go run main.go
```

## Docker

The service uses a multi-stage build process:

1. Build stage:
   - Base image: golang:1.21-alpine
   - Compiles the Go application with CGO disabled
   - Produces a statically linked binary

2. Production stage:
   - Base image: alpine:3.18
   - Runs as non-root user for security
   - Exposes port 8082

### Building the Image

```bash
docker build -t order-processor:latest -f builds/dockerfiles/Dockerfile.order-processing .
```

## Kubernetes Deployment

The service is deployed using Kubernetes:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: order-processor
spec:
  replicas: 2
  template:
    spec:
      containers:
      - name: order-processor
        image: order-processor:latest
        ports:
        - containerPort: 8082
        env:
        - name: STORAGE_SERVICE_URL
          valueFrom:
            configMapKeyRef:
              name: order-config
              key: STORAGE_SERVICE_URL
        - name: SERVER_PORT
          valueFrom:
            configMapKeyRef:
              name: order-config
              key: SERVER_PORT
        securityContext:
          runAsNonRoot: true
          runAsUser: 1000
          allowPrivilegeEscalation: false
        resources:
          requests:
            cpu: "100m"
            memory: "128Mi"
          limits:
            cpu: "200m"
            memory: "256Mi"
```

### Service Dependencies
1. Storage Service
   - Order data persistence
   - Inventory updates

### Order Processing Flow
1. Receive order request from Books Service
2. Validate order details
3. Create order record in Storage Service
4. Update inventory in Storage Service
5. Return order confirmation

### Key Features
- Runs 2 replicas for high availability
- Secure configuration with non-root user
- Health checks at /health endpoint
- Resource limits and requests defined
- Container security context configured
- ConfigMap-based configuration
