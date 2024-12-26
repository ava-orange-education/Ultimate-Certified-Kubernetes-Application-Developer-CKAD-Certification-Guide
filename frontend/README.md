# Frontend Service

React-based frontend service for the bookstore application.

## Features
- Book catalog browsing
- Real-time inventory display
- Book purchase functionality
- Order status tracking

## Development

### Prerequisites
- Node.js 18+
- Books Service running on port 8081
- Order Processor running on port 8082

### Environment Variables
```
VITE_BOOKS_SERVICE_URL=http://localhost:8081
VITE_ORDER_PROCESSOR_URL=http://localhost:8082
```

### Running Locally
```bash
npm install
npm run dev
```

The development server will start at http://localhost:3000

## API Integration

### Books Service Integration
- `GET /books` - Fetch book catalog
- `GET /books/{id}` - Get book details
- `POST /books/purchase` - Purchase book

### Order Processor Integration
- `GET /orders/{id}` - Track order status
- `GET /orders/user/{userId}` - View user orders

## Project Structure
```
frontend/
├── src/
│   ├── components/      # React components
│   │   ├── AddBook.jsx
│   │   └── BookList.jsx
│   ├── services/        # API integration
│   │   ├── api.js
│   │   └── mockData.js
│   ├── styles/          # CSS modules
│   │   └── App.module.css
│   ├── App.jsx         # Main application
│   └── main.jsx        # Entry point
├── public/             # Static assets
└── index.html         # HTML template
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
        env:
        - name: VITE_BOOKS_SERVICE_URL
          valueFrom:
            configMapKeyRef:
              name: frontend-config
              key: VITE_BOOKS_SERVICE_URL
        - name: VITE_ORDER_PROCESSOR_URL
          valueFrom:
            configMapKeyRef:
              name: frontend-config
              key: VITE_ORDER_PROCESSOR_URL
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
- ConfigMap-based configuration
- Environment-specific service URLs
