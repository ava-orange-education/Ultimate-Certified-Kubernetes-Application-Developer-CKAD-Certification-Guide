# Makefile for Chapter 13: Security Best Practices

.PHONY: apply-all apply-service-accounts apply-security-contexts apply-rbac apply-deployments apply-network-policies clean

# Apply all manifests
apply-all: apply-service-accounts apply-rbac apply-deployments apply-network-policies

# Apply service account manifests
apply-service-accounts:
	@echo "Applying service account manifests..."
	kubectl apply -f service-accounts/

# Apply security context examples
apply-security-contexts:
	@echo "Applying security context examples..."
	kubectl apply -f security-contexts/

# Apply RBAC manifests
apply-rbac:
	@echo "Applying RBAC manifests..."
	kubectl apply -f rbac/roles/
	kubectl apply -f rbac/role-bindings/
	kubectl apply -f rbac/cluster-roles/
	kubectl apply -f rbac/cluster-role-bindings/

# Apply secure deployment manifests
apply-deployments:
	@echo "Applying secure deployment manifests..."
	kubectl apply -f deployments/

# Apply network policy manifests
apply-network-policies:
	@echo "Applying network policy manifests..."
	kubectl apply -f network-policies/

# Clean up all resources
clean:
	@echo "Cleaning up resources..."
	kubectl delete -f network-policies/ --ignore-not-found
	kubectl delete -f deployments/ --ignore-not-found
	kubectl delete -f rbac/cluster-role-bindings/ --ignore-not-found
	kubectl delete -f rbac/cluster-roles/ --ignore-not-found
	kubectl delete -f rbac/role-bindings/ --ignore-not-found
	kubectl delete -f rbac/roles/ --ignore-not-found
	kubectl delete -f security-contexts/ --ignore-not-found
	kubectl delete -f service-accounts/ --ignore-not-found
	@echo "Cleanup complete."

# Display help information
help:
	@echo "Available targets:"
	@echo "  apply-all              - Apply all manifests"
	@echo "  apply-service-accounts - Apply service account manifests"
	@echo "  apply-security-contexts - Apply security context examples"
	@echo "  apply-rbac             - Apply RBAC manifests"
	@echo "  apply-deployments      - Apply secure deployment manifests"
	@echo "  apply-network-policies - Apply network policy manifests"
	@echo "  clean                  - Clean up all resources"
	@echo "  help                   - Display this help information"
