# Storage Service

The storage service is responsible for managing persistent data storage for the book inventory system. It provides a centralized storage solution for book details and quantity management.

## Features

### Book Management
- Complete book lifecycle handling:
  - Book details storage and retrieval
  - Book metadata management
  - Book quantity tracking
  - Inventory adjustments

### Order Management
- Order data persistence
- Order status tracking
- Order history maintenance
- Integration with order processing service

## API Endpoints

All endpoints are internal, accessible only to other microservices.

### Book Operations
- `GET /internal/books/list` - List all books
- `GET /internal/books/get` - Get book by ID
- `POST /internal/books/add` - Add new book
- `PUT /internal/books/update` - Update book details
- `GET /internal/books/quantity` - Check book quantity
- `PUT /internal/books/update-quantity` - Update book quantity

### Order Operations
- `GET /internal/orders/list` - List all orders
- `GET /internal/orders/get` - Get order by ID
- `POST /internal/orders/add` - Create new order
- `PUT /internal/orders/update-status` - Update order status

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
```
.
├── handlers/
│   ├── book-details.go    # Book data operations
│   ├── book-quantity.go   # Inventory management
│   ├── order.go          # Order operations
│   └── storage.go        # Common handler logic
├── models/
│   └── request.go        # Data transfer objects
├── repository/
│   ├── books.go          # Book data persistence
│   └── orders.go         # Order data persistence
├── services/
│   └── storage.go        # Core business logic
└── main.go               # Service entry point
```

## Service Integration

The service operates on port `:8083` and provides internal APIs for:
- Books Service (`:8081`):
  - Book catalog operations
  - Inventory checks
- Order Processing Service (`:8082`):
  - Order management
  - Status updates
