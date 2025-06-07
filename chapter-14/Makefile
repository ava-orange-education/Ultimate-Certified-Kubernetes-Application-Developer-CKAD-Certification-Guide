# Makefile for Chapter 14: Resource Management and Quotas

.PHONY: apply-basic-resources apply-qos-classes apply-resource-quotas apply-limit-ranges apply-multi-tenant apply-avakart-resources clean-all

# Apply basic resource examples
apply-basic-resources:
	kubectl apply -f basic-resources/

# Apply QoS class examples
apply-qos-classes:
	kubectl apply -f qos-classes/

# Apply ResourceQuota examples
apply-resource-quotas:
	kubectl apply -f resource-quotas/

# Apply LimitRange examples
apply-limit-ranges:
	kubectl apply -f limit-ranges/

# Apply multi-tenant resource management examples
apply-multi-tenant:
	kubectl apply -f multi-tenant/

# Apply AvaKart resource management examples
apply-avakart-resources:
	kubectl apply -f avakart-resources/

# Clean up all resources
clean-all:
	kubectl delete -f avakart-resources/ --ignore-not-found
	kubectl delete -f multi-tenant/ --ignore-not-found
	kubectl delete -f limit-ranges/ --ignore-not-found
	kubectl delete -f resource-quotas/ --ignore-not-found
	kubectl delete -f qos-classes/ --ignore-not-found
	kubectl delete -f basic-resources/ --ignore-not-found

# Apply all examples
apply-all: apply-basic-resources apply-qos-classes apply-resource-quotas apply-limit-ranges apply-multi-tenant apply-avakart-resources
