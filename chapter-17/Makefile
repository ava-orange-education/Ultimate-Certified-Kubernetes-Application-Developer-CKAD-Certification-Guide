.PHONY: check-cni install-calico apply-policies test-policies monitor clean status troubleshoot

NAMESPACE ?= avakart
TIMEOUT ?= 300s

check-cni:
	@echo "Checking CNI plugin support for Network Policies..."
	@chmod +x cni-setup/check-cni-support.sh
	@./cni-setup/check-cni-support.sh

install-calico:
	@echo "Installing Calico CNI plugin..."
	@kubectl apply -f cni-setup/calico-installation.yaml
	@echo "Waiting for Calico to be ready..."
	@kubectl wait --for=condition=Ready pod -l k8s-app=calico-node -n calico-system --timeout=$(TIMEOUT)
	@echo "Calico installation completed successfully!"
	
create-namespace:
	@echo "Creating namespace $(NAMESPACE) if it doesn't exist..."
	@kubectl create namespace $(NAMESPACE) --dry-run=client -o yaml | kubectl apply -f -
	
apply-policies: create-namespace
	@echo "Applying all network policies to namespace $(NAMESPACE)..."
	@kubectl apply -f network-policies/ -n $(NAMESPACE) || true
	@echo "Network policies applied successfully!"
	@$(MAKE) status
	
apply-default: create-namespace
	@echo "Applying default deny policies..."
	@kubectl apply -f network-policies/default-deny-all.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/default-deny-ingress.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/default-deny-egress.yaml -n $(NAMESPACE)
	@echo "Default deny policies applied!"
	
apply-services: create-namespace
	@echo "Applying service-specific network policies..."
	@kubectl apply -f network-policies/frontend-ingress-policy.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/frontend-egress-policy.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/books-service-ingress-policy.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/books-service-egress-policy.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/order-processor-ingress-policy.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/order-processor-egress-policy.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/storage-service-ingress-policy.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/storage-service-egress-policy.yaml -n $(NAMESPACE)
	@echo "Service-specific policies applied!"
	
apply-namespace: create-namespace
	@echo "Applying namespace-based network policies..."
	@kubectl apply -f network-policies/allow-monitoring-namespace.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/allow-logging-namespace.yaml -n $(NAMESPACE)
	@kubectl apply -f network-policies/deny-cross-namespace.yaml -n $(NAMESPACE)
	@echo "Namespace-based policies applied!"
	
apply-complete: create-namespace
	@echo "Applying complete network policy set..."
	@kubectl apply -f network-policies/avakart-complete-policy.yaml -n $(NAMESPACE)
	@echo "Complete policy set applied!"
	
test-policies:
	@echo "Running comprehensive network policy tests..."
	@chmod +x testing/test-network-policies.sh
	@./testing/test-network-policies.sh
	
test-pods:
	@echo "Deploying network test pods..."
	@kubectl apply -f testing/network-test-pod.yaml -n $(NAMESPACE)
	@echo "Waiting for test pods to be ready..."
	@kubectl wait --for=condition=Ready pod -l purpose=troubleshooting -n $(NAMESPACE) --timeout=$(TIMEOUT) || true
	@echo "Test pods deployed. Use 'kubectl exec' to test connectivity manually."
	
monitor:
	@echo "Setting up network policy monitoring..."
	@kubectl create namespace monitoring --dry-run=client -o yaml | kubectl apply -f -
	@kubectl apply -f monitoring/network-policy-monitor.yaml
	@echo "Monitoring setup completed!"
	
status:
	@echo "Network Policy Status for namespace $(NAMESPACE):"
	@echo "================================================="
	@echo ""
	@echo "Network Policies:"
	@kubectl get networkpolicies -n $(NAMESPACE) -o wide 2>/dev/null || echo "No network policies found"
	@echo ""
	@echo "Pods and Labels:"
	@kubectl get pods -n $(NAMESPACE) --show-labels 2>/dev/null || echo "No pods found in namespace $(NAMESPACE)"
	@echo ""
	@echo "Services:"
	@kubectl get services -n $(NAMESPACE) 2>/dev/null || echo "No services found in namespace $(NAMESPACE)"

troubleshoot:
	@echo "Running Network Policy Troubleshooting Diagnostics..."
	@echo "====================================================="
	@echo ""
	@echo "1. Checking CNI Plugin Status:"
	@kubectl get pods -n kube-system | grep -E "(calico|cilium|weave|flannel|antrea)" || echo "No common CNI pods found"
	@echo ""
	@echo "2. Network Policy API Availability:"
	@kubectl api-resources | grep networkpolicies || echo "NetworkPolicy API not available"
	@echo ""
	@echo "3. Network Policies in $(NAMESPACE):"
	@kubectl get networkpolicies -n $(NAMESPACE) || echo "No network policies found"
	@echo ""
	@echo "4. Pod Labels (for selector matching):"
	@kubectl get pods -n $(NAMESPACE) --show-labels || echo "No pods found"
	@echo ""
	@echo "5. CNI Plugin Logs (last 10 lines):"
	@kubectl logs -n kube-system -l k8s-app=calico-node --tail=10 2>/dev/null || \
	 kubectl logs -n kube-system -l k8s-app=cilium --tail=10 2>/dev/null || \
	 kubectl logs -n kube-system -l name=weave-net --tail=10 2>/dev/null || \
	 echo "Unable to retrieve CNI plugin logs"
	@echo ""
	@echo "6. Testing Basic Connectivity:"
	@kubectl run test-connectivity --image=busybox --rm -it --restart=Never -n $(NAMESPACE) -- nslookup kubernetes.default.svc.cluster.local 2>/dev/null || echo "Connectivity test failed"

clean:
	@echo "Cleaning up all network policies and test resources..."
	@kubectl delete networkpolicies --all -n $(NAMESPACE) 2>/dev/null || true
	@kubectl delete pods -l purpose=troubleshooting -n $(NAMESPACE) 2>/dev/null || true
	@kubectl delete pods -l purpose=security-testing -n $(NAMESPACE) 2>/dev/null || true
	@kubectl delete namespace monitoring 2>/dev/null || true
	@echo "Cleanup completed!"
	
clean-test:
	@echo "Cleaning up test resources..."
	@kubectl delete pods -l purpose=troubleshooting -n $(NAMESPACE) 2>/dev/null || true
	@kubectl delete pods -l purpose=security-testing -n $(NAMESPACE) 2>/dev/null || true
	@kubectl delete namespace netpol-test 2>/dev/null || true
	@echo "Test resources cleaned up!"
	
validate:
	@echo "Validating network policy configurations..."
	@echo "=========================================="
	@for file in network-policies/*.yaml; do \
		echo "Validating $$file..."; \
		kubectl apply --dry-run=client -f $$file -n $(NAMESPACE) || echo "Validation failed for $$file"; \
	done
	@echo "Validation completed!"
	
describe:
	@echo "Detailed Network Policy Information:"
	@echo "===================================="
	@for policy in $$(kubectl get networkpolicies -n $(NAMESPACE) -o name 2>/dev/null); do \
		echo ""; \
		echo "Describing $$policy:"; \
		echo "-------------------"; \
		kubectl describe $$policy -n $(NAMESPACE); \
	done
	
dev-deploy: create-namespace apply-default apply-services test-pods
	@echo "Development deployment completed!"
	@echo "- Default deny policies applied"
	@echo "- Service-specific policies applied"
	@echo "- Test pods deployed"
	@echo ""
	@echo "Next steps:"
	@echo "1. Run 'make test-policies' to validate policies"
	@echo "2. Use 'make status' to check current state"
	@echo "3. Use 'make troubleshoot' if issues arise"
	
prod-deploy: create-namespace apply-complete monitor
	@echo "Production deployment completed!"
	@echo "- Complete policy set applied"
	@echo "- Monitoring configured"
	@echo ""
	@echo "Next steps:"
	@echo "1. Run 'make test-policies' to validate policies"
	@echo "2. Monitor policy compliance via monitoring dashboard"
	@echo "3. Review logs for any policy violations"
	
logs:
	@echo "Network Policy Monitor Logs:"
	@echo "============================"
	@kubectl logs -n monitoring -l app=network-policy-monitor --tail=50 || echo "Monitor not deployed or no logs available"

dashboard:
	@echo "Setting up port forwarding to monitoring dashboard..."
	@kubectl port-forward -n monitoring service/grafana 3000:3000 &
	@echo "Dashboard available at http://localhost:3000"
	@echo "Press Ctrl+C to stop port forwarding"
	
emergency-remove:
	@echo "EMERGENCY: Removing all network policies..."
	@echo "This will remove all network restrictions!"
	@read -p "Are you sure? Type 'yes' to continue: " confirm && [ "$$confirm" = "yes" ] || exit 1
	@kubectl delete networkpolicies --all -n $(NAMESPACE)
	@echo "All network policies removed!"
	
examples:
	@echo "Network Policy Examples and Recipes:"
	@echo "===================================="
	@echo ""
	@echo "1. Allow all traffic (remove restrictions):"
	@echo "   kubectl apply -f - <<EOF"
	@echo "   apiVersion: networking.k8s.io/v1"
	@echo "   kind: NetworkPolicy"
	@echo "   metadata:"
	@echo "     name: allow-all"
	@echo "     namespace: $(NAMESPACE)"
	@echo "   spec:"
	@echo "     podSelector: {}"
	@echo "     policyTypes:"
	@echo "     - Ingress"
	@echo "     - Egress"
	@echo "     ingress:"
	@echo "     - {}"
	@echo "     egress:"
	@echo "     - {}"
	@echo "   EOF"
	@echo ""
	@echo "2. Test connectivity between pods:"
	@echo "   kubectl exec -n $(NAMESPACE) <source-pod> -- curl -m 5 http://<target-service>:<port>/health"
	@echo ""
	@echo "3. Check DNS resolution:"
	@echo "   kubectl exec -n $(NAMESPACE) <pod> -- nslookup kubernetes.default.svc.cluster.local"
	@echo ""
	@echo "4. View iptables rules (on nodes):"
	@echo "   sudo iptables -L -n | grep -i calico"
	
perf-test:
	@echo "Running network policy performance test..."
	@kubectl run perf-test --image=busybox --rm -it --restart=Never -n $(NAMESPACE) -- sh -c '\
		echo "Testing connection latency..."; \
		time wget -q -O- http://books-service:8081/health; \
		echo "Testing DNS resolution latency..."; \
		time nslookup books-service; \
	' || echo "Performance test completed"
