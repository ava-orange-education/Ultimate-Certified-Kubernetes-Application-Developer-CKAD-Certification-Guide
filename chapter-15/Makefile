.PHONY: apply-all clean-all apply-service-accounts apply-rbac apply-admission-control apply-deployments apply-validation-policies

apply-all: apply-service-accounts apply-rbac apply-admission-control apply-deployments apply-validation-policies

clean-all:
	kubectl delete -f service-accounts/ --ignore-not-found
	kubectl delete -f rbac/roles/ --ignore-not-found
	kubectl delete -f rbac/role-bindings/ --ignore-not-found
	kubectl delete -f rbac/cluster-roles/ --ignore-not-found
	kubectl delete -f rbac/cluster-role-bindings/ --ignore-not-found
	kubectl delete -f admission-control/ --ignore-not-found
	kubectl delete -f deployments/ --ignore-not-found
	kubectl delete -f validation-policies/ --ignore-not-found

apply-service-accounts:
	kubectl apply -f service-accounts/

apply-rbac:
	kubectl apply -f rbac/roles/
	kubectl apply -f rbac/role-bindings/
	kubectl apply -f rbac/cluster-roles/
	kubectl apply -f rbac/cluster-role-bindings/

apply-admission-control:
	kubectl apply -f admission-control/

apply-deployments:
	kubectl apply -f deployments/

apply-validation-policies:
	kubectl apply -f validation-policies/
