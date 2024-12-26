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

## Configuration

The service can be configured using environment variables:

| Variable | Description | Default |
|----------|-------------|---------|
| PORT | The port number the service listens on | 8081 |
| STORAGE_SERVICE_URL | URL of the storage service | http://localhost:8083 |
| ORDER_PROCESSOR_URL | URL of the order processor service | http://localhost:8082 |

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
docker build -f ../../builds/dockerfiles/Dockerfile.books -t books.svc.avakart .
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

The service provides the following REST endpoints under `/api/books`:

### Book Management
- `GET /list` - List all available books
- `GET /details?id={bookId}` - Get detailed book information by ID
- `POST /add` - Add a new book to the catalog
  - Automatically generates UUID-based book ID
  - Associates book with seller ID

### Purchase Operations
- `POST /purchase` - Initiate book purchase
  - Associates purchase with buyer ID
  - Validates book availability through storage service
  - Creates order through order processing service
  - Manages inventory updates

Each endpoint integrates with appropriate backend services:
- Storage Service (`:8083`) for data persistence and inventory
- Order Processing Service (`:8082`) for order management

Detailed request/response formats and examples are available in the API documentation.
