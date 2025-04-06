# Makefile for Chapter 6: Kubernetes Deployment Strategies

# Variables
KUBECTL = kubectl
NAMESPACE = default

# Rolling Updates
apply-rolling-update:
	$(KUBECTL) apply -f rolling-updates/books-rolling-update.yaml -n $(NAMESPACE)

# Blue/Green Deployment
apply-blue-green:
	$(KUBECTL) apply -f blue-green/frontend-blue-green.yaml -n $(NAMESPACE)

switch-to-green:
	$(KUBECTL) patch service frontend-service -n $(NAMESPACE) --type='json' -p='[{"op": "replace", "path": "/spec/selector/version", "value":"green"}]'

switch-to-blue:
	$(KUBECTL) patch service frontend-service -n $(NAMESPACE) --type='json' -p='[{"op": "replace", "path": "/spec/selector/version", "value":"blue"}]'

# Canary Deployment
apply-canary:
	$(KUBECTL) apply -f canary/storage-canary.yaml -n $(NAMESPACE)

scale-canary:
	$(KUBECTL) scale deployment storage-canary -n $(NAMESPACE) --replicas=2

promote-canary:
	$(KUBECTL) scale deployment storage-stable -n $(NAMESPACE) --replicas=0
	$(KUBECTL) scale deployment storage-canary -n $(NAMESPACE) --replicas=3

# Advanced Ingress
apply-ingress:
	$(KUBECTL) apply -f advanced-ingress/frontend-ingress.yaml -n $(NAMESPACE)

# Apply all manifests
apply-all:
	$(KUBECTL) apply -f rolling-updates/ -n $(NAMESPACE)
	$(KUBECTL) apply -f blue-green/ -n $(NAMESPACE)
	$(KUBECTL) apply -f canary/ -n $(NAMESPACE)
	$(KUBECTL) apply -f advanced-ingress/ -n $(NAMESPACE)

# Delete all resources
delete-all:
	$(KUBECTL) delete -f rolling-updates/ -n $(NAMESPACE) --ignore-not-found
	$(KUBECTL) delete -f blue-green/ -n $(NAMESPACE) --ignore-not-found
	$(KUBECTL) delete -f canary/ -n $(NAMESPACE) --ignore-not-found
	$(KUBECTL) delete -f advanced-ingress/ -n $(NAMESPACE) --ignore-not-found

# Help
help:
	@echo "Available targets:"
	@echo "  apply-rolling-update  - Apply rolling update deployment for books service"
	@echo "  apply-blue-green      - Apply blue/green deployment for frontend"
	@echo "  switch-to-green       - Switch traffic to green (new) version"
	@echo "  switch-to-blue        - Switch traffic back to blue (stable) version"
	@echo "  apply-canary          - Apply canary deployment for storage service"
	@echo "  scale-canary          - Increase canary deployment to 2 replicas"
	@echo "  promote-canary        - Promote canary to production (scale stable to 0, canary to 3)"
	@echo "  apply-ingress         - Apply advanced ingress configuration"
	@echo "  apply-all             - Apply all manifests"
	@echo "  delete-all            - Delete all resources"
	@echo "  help                  - Show this help message"

.PHONY: apply-rolling-update apply-blue-green switch-to-green switch-to-blue apply-canary scale-canary promote-canary apply-ingress apply-all delete-all help
