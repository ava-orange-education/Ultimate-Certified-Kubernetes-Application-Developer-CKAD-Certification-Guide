# Makefile for Chapter 9: Monitoring and Debugging Applications

.PHONY: apply-all apply-books apply-storage apply-order-processor apply-frontend debug clean

# Apply all manifests
apply-all: apply-books apply-storage apply-order-processor apply-frontend
	@echo "All manifests applied"

# Apply books service manifests
apply-books:
	@echo "Applying books service manifests..."
	kubectl apply -f manifests/books-deployment.yaml

# Apply storage service manifests
apply-storage:
	@echo "Applying storage service manifests..."
	kubectl apply -f manifests/storage-deployment.yaml

# Apply order processor manifests
apply-order-processor:
	@echo "Applying order processor manifests..."
	kubectl apply -f manifests/order-processor-deployment.yaml

# Apply frontend manifests
apply-frontend:
	@echo "Applying frontend manifests..."
	kubectl apply -f manifests/frontend-deployment.yaml

# Run the debug script
debug:
	@echo "Running debug script..."
	chmod +x debugging/debug-avakart.sh
	./debugging/debug-avakart.sh

# Clean up all resources
clean:
	@echo "Cleaning up resources..."
	kubectl delete -f manifests/books-deployment.yaml --ignore-not-found
	kubectl delete -f manifests/storage-deployment.yaml --ignore-not-found
	kubectl delete -f manifests/order-processor-deployment.yaml --ignore-not-found
	kubectl delete -f manifests/frontend-deployment.yaml --ignore-not-found
	@echo "Cleanup complete"

# Help command
help:
	@echo "Available targets:"
	@echo "  apply-all              - Apply all manifests"
	@echo "  apply-books            - Apply books service manifests"
	@echo "  apply-storage          - Apply storage service manifests"
	@echo "  apply-order-processor  - Apply order processor manifests"
	@echo "  apply-frontend         - Apply frontend manifests"
	@echo "  debug                  - Run the debug script"
	@echo "  clean                  - Clean up all resources"
	@echo "  help                   - Show this help message"
