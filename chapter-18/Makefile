.PHONY: install-crds create-namespaces deploy-examples verify test-validation cleanup clean-resources clean-namespaces clean-crds status

install-crds:
	@echo "Installing Custom Resource Definitions..."
	kubectl apply -f crds/bookcatalog-crd.yaml
	kubectl apply -f crds/order-crd.yaml
	kubectl apply -f crds/avakartconfig-crd.yaml
	@echo "Waiting for CRDs to be established..."
	kubectl wait --for condition=established --timeout=60s crd/bookcatalogs.avakart.io
	kubectl wait --for condition=established --timeout=60s crd/orders.avakart.io
	kubectl wait --for condition=established --timeout=60s crd/avakartconfigs.avakart.io
	@echo "CRDs installed successfully!"
	
create-namespaces:
	@echo "Creating namespaces..."
	kubectl create namespace avakart --dry-run=client -o yaml | kubectl apply -f -
	kubectl create namespace avakart-dev --dry-run=client -o yaml | kubectl apply -f -
	kubectl create namespace avakart-staging --dry-run=client -o yaml | kubectl apply -f -
	kubectl create namespace avakart-ha --dry-run=client -o yaml | kubectl apply -f -
	@echo "Namespaces created successfully!"
	
deploy-examples:
	@echo "Deploying example custom resources..."
	kubectl apply -f custom-resources/bookcatalog-examples.yaml
	kubectl apply -f custom-resources/order-examples.yaml
	kubectl apply -f custom-resources/avakartconfig-examples.yaml
	@echo "Example resources deployed successfully!"
	
verify:
	@echo "Verifying CRD installation..."
	@echo "=== Custom Resource Definitions ==="
	kubectl get crd | grep avakart || echo "No AvaKart CRDs found"
	@echo ""
	@echo "=== BookCatalog Resources ==="
	kubectl get bookcatalogs -A || echo "No BookCatalog resources found"
	@echo ""
	@echo "=== Order Resources ==="
	kubectl get orders -A || echo "No Order resources found"
	@echo ""
	@echo "=== AvaKartConfig Resources ==="
	kubectl get avakartconfigs -A || echo "No AvaKartConfig resources found"
	@echo ""
	@echo "=== Namespaces ==="
	kubectl get namespaces | grep avakart || echo "No AvaKart namespaces found"
	
test-validation:
	@echo "Testing CRD validation with invalid examples..."
	@echo "The following commands should fail with validation errors:"
	@echo ""
	@echo "Testing invalid BookCatalog (empty title)..."
	-kubectl apply --dry-run=server -f - <<EOF || echo "✓ Validation correctly rejected empty title"
	apiVersion: avakart.io/v1
	kind: BookCatalog
	metadata:
	  name: invalid-empty-title
	  namespace: avakart
	spec:
	  title: ""
	  author: "Valid Author"
	  isbn: "978-1234567890"
	  price: 29.99
	  category: "technical"
	EOF
	@echo ""
	@echo "Testing invalid Order (invalid zip code)..."
	-kubectl apply --dry-run=server -f - <<EOF || echo "✓ Validation correctly rejected invalid zip code"
	apiVersion: avakart.io/v1
	kind: Order
	metadata:
	  name: invalid-zipcode
	  namespace: avakart
	spec:
	  customerId: "customer-12345"
	  itemCount: 1
	  items:
	  - bookId: "kubernetes-guide"
	    quantity: 1
	    price: 49.99
	  shippingAddress:
	    street: "123 Tech Street"
	    city: "San Francisco"
	    state: "CA"
	    zipCode: "invalid-zip"
	    country: "US"
	  paymentMethod: "credit-card"
	EOF
	@echo ""
	@echo "Validation testing completed!"
	
status:
	@echo "=== Chapter 18 Resource Status ==="
	@echo ""
	@echo "Custom Resource Definitions:"
	kubectl get crd | grep avakart || echo "No AvaKart CRDs found"
	@echo ""
	@echo "Namespaces:"
	kubectl get namespaces | grep avakart || echo "No AvaKart namespaces found"
	@echo ""
	@echo "BookCatalog Resources:"
	kubectl get bookcatalogs -A -o wide || echo "No BookCatalog resources found"
	@echo ""
	@echo "Order Resources:"
	kubectl get orders -A -o wide || echo "No Order resources found"
	@echo ""
	@echo "AvaKartConfig Resources:"
	kubectl get avakartconfigs -A -o wide || echo "No AvaKartConfig resources found"

cleanup: clean-resources clean-namespaces clean-crds
	@echo "Complete cleanup finished!"
	
clean-resources:
	@echo "Cleaning up custom resources..."
	-kubectl delete -f custom-resources/ --ignore-not-found=true
	@echo "Custom resources cleaned up!"
	
clean-namespaces:
	@echo "Cleaning up namespaces..."
	-kubectl delete namespace avakart --ignore-not-found=true
	-kubectl delete namespace avakart-dev --ignore-not-found=true
	-kubectl delete namespace avakart-staging --ignore-not-found=true
	-kubectl delete namespace avakart-ha --ignore-not-found=true
	@echo "Namespaces cleaned up!"
	
clean-crds:
	@echo "Cleaning up Custom Resource Definitions..."
	@echo "WARNING: This will delete all custom resources of these types!"
	-kubectl delete crd bookcatalogs.avakart.io --ignore-not-found=true
	-kubectl delete crd orders.avakart.io --ignore-not-found=true
	-kubectl delete crd avakartconfigs.avakart.io --ignore-not-found=true
	@echo "CRDs cleaned up!"
	
dev-setup: install-crds create-namespaces
	@echo "Environment setup complete!"
	
demo: install-crds create-namespaces deploy-examples
	@echo "Demo environment ready!"
	@echo ""
	@echo "Try these commands:"
	@echo "  kubectl get bookcatalogs -n avakart"
	@echo "  kubectl get orders -n avakart -o wide"
	@echo "  kubectl get avakartconfigs -A"
	@echo "  kubectl describe bookcatalog kubernetes-guide -n avakart"

test-all-validation:
	@echo "Testing all validation scenarios..."
	@echo "Note: All of these should fail with validation errors"
	-kubectl apply -f examples/invalid-examples.yaml --dry-run=server
	@echo "Validation testing completed!"

show-schemas:
	@echo "=== BookCatalog CRD Schema ==="
	kubectl get crd bookcatalogs.avakart.io -o jsonpath='{.spec.versions[0].schema.openAPIV3Schema}' | jq '.' || echo "jq not available, showing raw output:"
	kubectl get crd bookcatalogs.avakart.io -o jsonpath='{.spec.versions[0].schema.openAPIV3Schema}'
	@echo ""
	@echo "=== Order CRD Schema ==="
	kubectl get crd orders.avakart.io -o jsonpath='{.spec.versions[0].schema.openAPIV3Schema}' | jq '.' || echo "jq not available, showing raw output:"
	kubectl get crd orders.avakart.io -o jsonpath='{.spec.versions[0].schema.openAPIV3Schema}'
	@echo ""
	@echo "=== AvaKartConfig CRD Schema ==="
	kubectl get crd avakartconfigs.avakart.io -o jsonpath='{.spec.versions[0].schema.openAPIV3Schema}' | jq '.' || echo "jq not available, showing raw output:"
	kubectl get crd avakartconfigs.avakart.io -o jsonpath='{.spec.versions[0].schema.openAPIV3Schema}'
