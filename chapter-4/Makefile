.PHONY: apply-sidecar delete-sidecar apply-init delete-init apply-ambassador delete-ambassador

# Sidecar Pattern
apply-sidecar:
	@echo "Applying storage deployment with logging sidecar..."
	kubectl apply -f deployments/storage-deploy-with-sidecar.yaml

delete-sidecar:
	@echo "Deleting storage deployment with logging sidecar..."
	kubectl delete -f deployments/storage-deploy-with-sidecar.yaml

# Init Container Pattern
apply-init:
	@echo "Applying books deployment with init container..."
	kubectl apply -f deployments/books-deploy-with-init.yaml

delete-init:
	@echo "Deleting books deployment with init container..."
	kubectl delete -f deployments/books-deploy-with-init.yaml

# Ambassador Pattern
apply-ambassador:
	@echo "Applying storage ambassador ConfigMap..."
	kubectl apply -f deployments/storage-ambassador-configmap.yaml
	@echo "Applying order processor deployment with ambassador..."
	kubectl apply -f deployments/order-processor-with-ambassador.yaml

delete-ambassador:
	@echo "Deleting order processor deployment with ambassador..."
	kubectl delete -f deployments/order-processor-with-ambassador.yaml
	@echo "Deleting storage ambassador ConfigMap..."
	kubectl delete -f deployments/storage-ambassador-configmap.yaml
