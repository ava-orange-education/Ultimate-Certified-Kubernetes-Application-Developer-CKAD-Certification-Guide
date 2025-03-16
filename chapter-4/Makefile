.PHONY: apply-sidecar delete-sidecar apply-init delete-init

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
	@echo "Applying books pod with init container..."
	kubectl apply -f pods/books-pod-with-init.yaml

delete-init:
	@echo "Deleting books deployment with init container..."
	kubectl delete -f deployments/books-deploy-with-init.yaml
	@echo "Deleting books pod with init container..."
	kubectl delete -f pods/books-pod-with-init.yaml
