# Chapter 10: Probes and Health Checks
# Makefile for applying and managing Kubernetes manifests

.PHONY: apply-basic apply-advanced apply-comprehensive apply-all delete-basic delete-advanced delete-comprehensive delete-all

# Apply basic probe implementations
apply-basic:
	kubectl apply -f basic-probes/books-deployment.yaml
	kubectl apply -f basic-probes/storage-deployment.yaml
	kubectl apply -f basic-probes/order-processor-deployment.yaml
	@echo "Basic probe implementations applied"

# Apply advanced probe configurations
apply-advanced:
	kubectl apply -f advanced-probes/frontend-deployment.yaml
	kubectl apply -f advanced-probes/storage-tcp-probe.yaml
	kubectl apply -f advanced-probes/books-exec-probe.yaml
	@echo "Advanced probe configurations applied"

# Apply comprehensive health check implementation
apply-comprehensive:
	kubectl apply -f comprehensive/order-processor-advanced-health.yaml
	@echo "Comprehensive health check implementation applied"

# Apply all manifests
apply-all: apply-basic apply-advanced apply-comprehensive
	@echo "All Chapter 10 manifests applied"

# Delete basic probe implementations
delete-basic:
	kubectl delete -f basic-probes/books-deployment.yaml
	kubectl delete -f basic-probes/storage-deployment.yaml
	kubectl delete -f basic-probes/order-processor-deployment.yaml
	@echo "Basic probe implementations deleted"

# Delete advanced probe configurations
delete-advanced:
	kubectl delete -f advanced-probes/frontend-deployment.yaml
	kubectl delete -f advanced-probes/storage-tcp-probe.yaml
	kubectl delete -f advanced-probes/books-exec-probe.yaml
	@echo "Advanced probe configurations deleted"

# Delete comprehensive health check implementation
delete-comprehensive:
	kubectl delete -f comprehensive/order-processor-advanced-health.yaml
	@echo "Comprehensive health check implementation deleted"

# Delete all manifests
delete-all: delete-basic delete-advanced delete-comprehensive
	@echo "All Chapter 10 manifests deleted"

# Check the status of the deployments
status:
	@echo "Checking status of deployments..."
	kubectl get deployments -l app=books-service
	kubectl get deployments -l app=storage-service
	kubectl get deployments -l app=order-processor
	kubectl get deployments -l app=frontend

# Describe the deployments to see probe configurations
describe:
	@echo "Describing deployments to see probe configurations..."
	kubectl describe deployments -l app=books-service
	kubectl describe deployments -l app=storage-service
	kubectl describe deployments -l app=order-processor
	kubectl describe deployments -l app=frontend
