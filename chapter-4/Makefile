.PHONY: apply-sidecar delete-sidecar

# Sidecar Pattern
apply-sidecar:
	@echo "Applying storage deployment with logging sidecar..."
	kubectl apply -f deployments/storage-deploy-with-sidecar.yaml

delete-sidecar:
	@echo "Deleting storage deployment with logging sidecar..."
	kubectl delete -f deployments/storage-deploy-with-sidecar.yaml
