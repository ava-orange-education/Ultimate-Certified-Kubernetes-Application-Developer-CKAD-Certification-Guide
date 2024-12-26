# Storage Service

Go-based service for managing storage and inventory in the bookstore application.

## API Endpoints

### Book Management
- `GET /books` - List all books with inventory
- `GET /books/{id}` - Get book details and stock
- `PUT /books/{id}/quantity` - Update book quantity
  ```json
  {
    "quantity": 1
  }
  ```

### Order Management
- `POST /orders` - Create order record
  ```json
  {
    "orderId": "string",
    "bookId": "string",
    "quantity": 1,
    "userId": "string",
    "status": "pending"
  }
  ```
- `GET /orders/{id}` - Get order details
- `PUT /orders/{id}/status` - Update order status
  ```json
  {
    "status": "completed"
  }
  ```

### Health Check
- `GET /health` - Service health status

## Development

### Prerequisites
- Go 1.21
- MongoDB instance

### Environment Variables
```
MONGODB_URI=mongodb://localhost:27017/avakart
SERVER_PORT=8083
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
   - Exposes port 8083

### Building the Image

```bash
docker build -t storage-service:latest -f builds/dockerfiles/Dockerfile.storage .
```

## Kubernetes Deployment

The service is deployed using Kubernetes:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: storage-service
spec:
  replicas: 2
  template:
    spec:
      containers:
      - name: storage-service
        image: storage-service:latest
        ports:
        - containerPort: 8083
        env:
        - name: MONGODB_URI
          valueFrom:
            configMapKeyRef:
              name: storage-config
              key: MONGODB_URI
        - name: SERVER_PORT
          valueFrom:
            configMapKeyRef:
              name: storage-config
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

### Data Models

#### Book
```json
{
  "id": "string",
  "title": "string",
  "author": "string",
  "price": 0.00,
  "quantity": 0
}
```

#### Order
```json
{
  "id": "string",
  "bookId": "string",
  "userId": "string",
  "quantity": 0,
  "status": "string",
  "createdAt": "timestamp",
  "updatedAt": "timestamp"
}
```

### Key Features
- Runs 2 replicas for high availability
- Secure configuration with non-root user
- Health checks at /health endpoint
- Resource limits and requests defined
- Container security context configured
- ConfigMap-based configuration
- MongoDB persistence
