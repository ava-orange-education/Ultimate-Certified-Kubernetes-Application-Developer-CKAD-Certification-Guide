# Storage Service

Go-based service for managing storage and inventory in the bookstore application.

## Development

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
   - Exposes port 8080

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
        - containerPort: 8080
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

### Key Features
- Runs 2 replicas for high availability
- Secure configuration with non-root user
- Health checks at /health endpoint
- Resource limits and requests defined
- Container security context configured
