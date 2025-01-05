# Development Commands
.PHONY: dev-frontend dev-backend dev-order dev-storage dev-all install-deps

install-deps:
	@echo "Installing frontend dependencies..."
	cd frontend && npm install
	@echo "Installing backend dependencies..."
	cd backend/apps/books && go mod download
	cd backend/apps/order-processor && go mod download
	cd backend/apps/storage && go mod download

dev-frontend:
	@echo "Starting frontend development server..."
	cd frontend && npm run dev

dev-backend:
	@echo "Starting backend service..."
	cd backend/apps/books && go run main.go

dev-order:
	@echo "Starting order processing service..."
	cd backend/apps/order-processor && go run main.go

dev-storage:
	@echo "Starting storage service..."
	cd backend/apps/storage && go run main.go

dev-all:
	@echo "Starting all services in development mode..."
	make dev-storage & make dev-backend & make dev-order & make dev-frontend

# Docker Commands
.PHONY: docker-build docker-up docker-down docker-logs docker-clean

docker-build:
	@echo "Building all Docker images..."
	docker-compose -f builds/deployments/docker-compose.yaml build

docker-up:
	@echo "Starting all services in Docker..."
	docker-compose -f builds/deployments/docker-compose.yaml up -d

docker-down:
	@echo "Stopping all services..."
	docker-compose -f builds/deployments/docker-compose.yaml down

docker-logs:
	@echo "Showing logs from all services..."
	docker-compose -f builds/deployments/docker-compose.yaml logs -f

docker-clean:
	@echo "Cleaning up Docker resources..."
	docker-compose -f builds/deployments/docker-compose.yaml down --rmi all --volumes --remove-orphans

# Clean Commands
.PHONY: clean clean-deps clean-all

clean:
	@echo "Cleaning build artifacts..."
	rm -rf frontend/dist
	find . -name "*.log" -type f -delete

clean-deps:
	@echo "Cleaning dependencies..."
	rm -rf frontend/node_modules
	find . -name "go.sum" -type f -delete

clean-all: clean clean-deps docker-clean

.PHONY: docker-build-frontend docker-build-books docker-build-order docker-build-storage docker-push-all docker-push-frontend docker-push-books docker-push-order docker-push-storage

docker-build-frontend:
	@echo "Building Docker image for the frontend service..."
	docker build -t frontend:latest -f builds/dockerfiles/Dockerfile.frontend frontend

docker-build-books:
	@echo "Building Docker image for the Books Service..."
	docker build -t books-service:latest -f builds/dockerfiles/Dockerfile.books backend/apps/books

docker-build-order:
	@echo "Building Docker image for the Order Processing Service..."
	docker build -t order-processor:latest -f builds/dockerfiles/Dockerfile.order-processing backend/apps/order-processor

docker-build-storage:
	@echo "Building Docker image for the Storage Service..."
	docker build -t storage-service:latest -f builds/dockerfiles/Dockerfile.storage backend/apps/storage

docker-push-frontend:
	@echo "Pushing Docker image for the frontend service..."
	docker push frontend:latest

docker-push-books:
	@echo "Pushing Docker image for the Books Service..."
	docker push books-service:latest

docker-push-order:
	@echo "Pushing Docker image for the Order Processing Service..."
	docker push order-processor:latest

docker-push-storage:
	@echo "Pushing Docker image for the Storage Service..."
	docker push storage-service:latest

# Helper Commands
.PHONY: help

help:
	@echo "Available commands:"
	@echo ""
	@echo "Development Commands:"
	@echo "  make install-deps    - Install all dependencies"
	@echo "  make dev-frontend    - Run frontend development server"
	@echo "  make dev-backend     - Run backend service"
	@echo "  make dev-order       - Run order processing service"
	@echo "  make dev-storage     - Run storage service"
	@echo "  make dev-all         - Run all services in development mode"
	@echo ""
	@echo "Docker Commands:"
	@echo "  make docker-build    - Build all Docker images"
	@echo "  make docker-up       - Start all services in Docker"
	@echo "  make docker-down     - Stop all services"
	@echo "  make docker-logs     - Show logs from all services"
	@echo "  make docker-clean    - Clean up Docker resources"
	@echo "  make docker-build-frontend  - Build Docker image for frontend service"
	@echo "  make docker-build-books     - Build Docker image for Books Service"
	@echo "  make docker-build-order     - Build Docker image for Order Processing Service"
	@echo "  make docker-build-storage   - Build Docker image for Storage Service"
	@echo "  make docker-push-frontend   - Push Docker image for frontend service"
	@echo "  make docker-push-books      - Push Docker image for Books Service"
	@echo "  make docker-push-order      - Push Docker image for Order Processing Service"
	@echo "  make docker-push-storage    - Push Docker image for Storage Service"
	@echo ""
	@echo "Clean Commands:"
	@echo "  make clean           - Clean build artifacts"
	@echo "  make clean-deps      - Clean dependencies"
	@echo "  make clean-all       - Clean everything (artifacts, deps, docker)"
	@echo ""
	@echo "Helper Commands:"
	@echo "  make help            - Show this help message"

# Default target
.DEFAULT_GOAL := help
