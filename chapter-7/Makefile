# Makefile for Chapter 7: Rolling Updates & Rollbacks

# Variables
KUBECTL = kubectl
NAMESPACE = default

# Rolling Updates
apply-rolling-update:
	$(KUBECTL) apply -f rolling-updates/books-rolling-update.yaml -n $(NAMESPACE)

# Rollback Mechanisms
apply-v1:
	$(KUBECTL) apply -f rollback/books-deployment-v1.yaml -n $(NAMESPACE)

apply-v2:
	$(KUBECTL) apply -f rollback/books-deployment-v2.yaml -n $(NAMESPACE)

rollback:
	$(KUBECTL) rollout undo deployment/books-deployment -n $(NAMESPACE)

rollback-to-revision:
	$(KUBECTL) rollout undo deployment/books-deployment --to-revision=$(REVISION) -n $(NAMESPACE)

rollout-history:
	$(KUBECTL) rollout history deployment/books-deployment -n $(NAMESPACE)

rollout-status:
	$(KUBECTL) rollout status deployment/books-deployment -n $(NAMESPACE)

# Advanced Health Checking
apply-advanced-health:
	$(KUBECTL) apply -f health-probes/order-processor-advanced-health.yaml -n $(NAMESPACE)

# Demo Scripts
run-rollback-demo:
	cd rollback && ./rollback-demo.sh

# Apply all manifests
apply-all:
	$(KUBECTL) apply -f rolling-updates/ -n $(NAMESPACE)
	$(KUBECTL) apply -f rollback/ -n $(NAMESPACE)
	$(KUBECTL) apply -f health-probes/ -n $(NAMESPACE)

# Delete all resources
delete-all:
	$(KUBECTL) delete -f rolling-updates/ -n $(NAMESPACE) --ignore-not-found
	$(KUBECTL) delete -f rollback/ -n $(NAMESPACE) --ignore-not-found
	$(KUBECTL) delete -f health-probes/ -n $(NAMESPACE) --ignore-not-found

# Help
help:
	@echo "Available targets:"
	@echo "  apply-rolling-update    - Apply advanced rolling update configuration for books service"
	@echo "  apply-v1                - Apply books service v1 deployment"
	@echo "  apply-v2                - Apply books service v2 deployment (with potential issues)"
	@echo "  rollback                - Rollback to the previous deployment version"
	@echo "  rollback-to-revision    - Rollback to a specific revision (use with REVISION=n)"
	@echo "  rollout-history         - Show deployment revision history"
	@echo "  rollout-status          - Show current deployment rollout status"
	@echo "  apply-advanced-health   - Apply deployment with advanced health checking"
	@echo "  run-rollback-demo       - Run the rollback demonstration script"
	@echo "  apply-all               - Apply all manifests"
	@echo "  delete-all              - Delete all resources"
	@echo "  help                    - Show this help message"

.PHONY: apply-rolling-update apply-v1 apply-v2 rollback rollback-to-revision rollout-history rollout-status apply-advanced-health run-rollback-demo apply-all delete-all help
