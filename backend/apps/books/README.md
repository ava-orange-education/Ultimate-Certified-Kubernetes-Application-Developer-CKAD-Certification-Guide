# AvaKart Books Service

The books service is a core microservice of the AvaKart e-commerce platform, responsible for managing the book catalog and related operations. Written in Go, it provides RESTful APIs for book management and integrates with other platform services.

## Features

- Book catalog management (CRUD operations)
- Book metadata handling through structured models
- Book purchase initiation and availability checks
- RESTful API endpoints for:
  - Book listings and search
  - Book details retrieval
  - Book inventory management
  - Purchase processing
- Integration with:
  - Storage service for persistence and inventory
  - Order processor service for purchase handling

## Development Setup

### Prerequisites

- Go 1.23.3
- Docker and Docker Compose (for local development)
- Kubernetes cluster (for deployment)

### Local Development

1. Install dependencies:
   ```bash
   go mod download
   ```

2. Run the service:
   ```bash
   go run main.go
   ```

### Docker Build

**Not yet implemeted**

Build the service using the provided Dockerfile:
```bash
docker build -f ../../builds/dockerfiles/Dockerfile.backend -t avakart-books .
```

### Kubernetes Deployment

**Not yet implemeted**

The service can be deployed to Kubernetes using the manifests in `builds/deployments/k8s/backend/`:

```bash
kubectl apply -f builds/deployments/k8s/backend/
```

This will create:
- Deployment with the books service
- Service for network access
- ConfigMap for configuration

## Project Structure

```
.
├── main.go          # Service entry point
├── handlers/        # HTTP request handlers
│   └── books.go     # Books-related endpoint handlers
├── models/          # Data models
│   └── book.go      # Book entity definition
└── services/        # Business logic implementation
    └── books.go     # Books service implementation
```

## Integration Points

- **Storage Service**: 
  - Book data persistence
  - Inventory management
  - Internal quantity checks
- **Order Processor**: 
  - Purchase request handling
  - Order creation and management
- **Frontend**: 
  - Book listing (BookList component)
  - Book addition (AddBook component)

## API Documentation

The service provides the following REST endpoints:

### Book Management
- `GET /books/list` - List all books
- `GET /books/get?id={bookId}` - Get book details by ID
- `POST /books/add` - Add a new book

### Purchase Operations
- `POST /orders/purchase` - Initiate book purchase
  - Validates book availability
  - Creates order through order processing service
  - Manages inventory updates

Each endpoint integrates with appropriate backend services for data persistence and order processing. Detailed API specifications including request/response formats will be maintained in a separate API documentation.
