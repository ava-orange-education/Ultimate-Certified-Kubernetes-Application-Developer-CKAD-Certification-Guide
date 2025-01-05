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