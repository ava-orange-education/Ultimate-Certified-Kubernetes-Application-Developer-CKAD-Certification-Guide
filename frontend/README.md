# Frontend Service

React-based frontend service for the bookstore application.

## Development

```bash
npm install
npm run dev
```

## Docker

The service uses a multi-stage build process:

1. Build stage:
   - Base image: node:18-alpine
   - Installs dependencies and builds the React application

2. Production stage:
   - Base image: nginx:alpine
   - Serves the built static files using Nginx
   - Exposes port 80

### Building the Image

```bash
docker build -t frontend:latest -f builds/dockerfiles/Dockerfile.frontend .
```

## Kubernetes Deployment

The service is deployed using Kubernetes:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: frontend
spec:
  replicas: 2
  template:
    spec:
      containers:
      - name: frontend
        image: frontend:latest
        ports:
        - containerPort: 80
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
- Uses Nginx to serve static content
- Includes health checks for reliability
- Resource limits and requests defined
