# Books Service

Go-based backend service for managing books in the bookstore application.

## API Endpoints

### Books Management
- `GET /books` - List all books
- `GET /books/{id}` - Get book details
- `POST /books/purchase` - Initiate book purchase
  ```json
  {
    "bookId": "string",
    "quantity": 1
  }
  ```

### Health Check
- `GET /health` - Service health status

## Development

### Prerequisites
- Go 1.21
- Storage Service running on port 8083
- Order Processor running on port 8082

### Environment Variables
```
STORAGE_SERVICE_URL=http://localhost:8083
ORDER_PROCESSOR_URL=http://localhost:8082
SERVER_PORT=8081
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
   - Exposes port 8081

### Building the Image

```bash
docker build -t books-service:latest -f builds/dockerfiles/Dockerfile.books .
```

## Kubernetes Deployment

The service is deployed using Kubernetes:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: books-service
spec:
  replicas: 2
  template:
    spec:
      containers:
      - name: books-service
        image: books-service:latest
        ports:
        - containerPort: 8081
        env:
        - name: STORAGE_SERVICE_URL
          valueFrom:
            configMapKeyRef:
              name: books-config
              key: STORAGE_SERVICE_URL
        - name: ORDER_PROCESSOR_URL
          valueFrom:
            configMapKeyRef:
              name: books-config
              key: ORDER_PROCESSOR_URL
        - name: SERVER_PORT
          valueFrom:
            configMapKeyRef:
              name: books-config
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
   - Book inventory management
   - Book metadata storage
2. Order Processor
   - Order creation and management

### Key Features
- Runs 2 replicas for high availability
- Secure configuration with non-root user
- Health checks at /health endpoint
- Resource limits and requests defined
- Container security context configured
- ConfigMap-based configuration
