.PHONY: apply-all apply-configmaps apply-secrets apply-deployments clean

apply-all: apply-configmaps apply-secrets apply-deployments

apply-configmaps:
	@echo "Applying ConfigMaps..."
	kubectl apply -f configmaps/

apply-secrets:
	@echo "Applying Secrets..."
	kubectl apply -f secrets/

apply-deployments:
	@echo "Applying Deployments..."
	kubectl apply -f deployments/

clean:
	@echo "Deleting all resources..."
	kubectl delete -f deployments/ --ignore-not-found
	kubectl delete -f secrets/ --ignore-not-found
	kubectl delete -f configmaps/ --ignore-not-found

help:
	@echo "Available targets:"
	@echo "  apply-all          - Apply all ConfigMaps, Secrets, and Deployments"
	@echo "  apply-configmaps   - Apply only ConfigMaps"
	@echo "  apply-secrets      - Apply only Secrets"
	@echo "  apply-deployments  - Apply only Deployments"
	@echo "  clean              - Delete all resources"
	@echo "  help               - Show this help message"
