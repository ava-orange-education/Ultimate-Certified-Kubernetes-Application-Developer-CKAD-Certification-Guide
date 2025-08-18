.PHONY: create-namespace deploy deploy-services deploy-ingress deploy-ssl deploy-advanced deploy-all clean clean-all status

create-namespace:
	@echo "Creating avakart namespace..."
	kubectl create namespace avakart --dry-run=client -o yaml | kubectl apply -f -
	@echo "Namespace created successfully!"

deploy-services: create-namespace
	@echo "Deploying ClusterIP services..."
	kubectl apply -f services/books-service-clusterip.yaml
	kubectl apply -f services/storage-service-clusterip.yaml
	kubectl apply -f services/order-processor-service-clusterip.yaml
	@echo "Deploying NodePort services..."
	kubectl apply -f services/frontend-service-nodeport.yaml
	kubectl apply -f services/books-service-nodeport.yaml
	@echo "Deploying LoadBalancer services..."
	kubectl apply -f services/frontend-loadbalancer.yaml
	kubectl apply -f services/api-gateway-loadbalancer.yaml
	@echo "Services deployed successfully!"

deploy-stateful: create-namespace
	@echo "Deploying headless services and StatefulSet..."
	kubectl apply -f stateful/database-headless-service.yaml
	kubectl apply -f stateful/database-statefulset.yaml
	@echo "Stateful resources deployed successfully!"

deploy-ingress: create-namespace
	@echo "Deploying NGINX Ingress Controller..."
	kubectl apply -f ingress/nginx-ingress-rbac.yaml
	kubectl apply -f ingress/nginx-ingress-controller.yaml
	kubectl apply -f ingress/nginx-ingress-service.yaml
	@echo "Waiting for Ingress Controller to be ready..."
	kubectl wait --namespace ingress-nginx --for=condition=ready pod --selector=app=nginx-ingress-controller --timeout=300s
	@echo "Deploying Ingress rules..."
	kubectl apply -f ingress/avakart-frontend-ingress.yaml
	kubectl apply -f ingress/avakart-api-ingress.yaml
	@echo "Ingress resources deployed successfully!"

deploy-ssl: create-namespace
	@echo "Deploying SSL/TLS certificates..."
	kubectl apply -f ssl/avakart-tls-secret.yaml
	kubectl apply -f ssl/avakart-api-tls-secret.yaml
	@echo "SSL/TLS certificates deployed successfully!"

deploy-advanced: create-namespace
	@echo "Deploying advanced Ingress features..."
	kubectl apply -f advanced/books-service-canary-ingress.yaml
	kubectl apply -f advanced/avakart-rate-limited-ingress.yaml
	kubectl apply -f advanced/avakart-multi-tenant-ingress.yaml
	@echo "Advanced Ingress features deployed successfully!"

deploy-all: create-namespace deploy-services deploy-stateful deploy-ssl deploy-ingress deploy-advanced
	@echo "All Chapter 16 resources deployed successfully!"
	@echo ""
	@echo "To check the status of your deployments, run: make status"

status:
	@echo "=== Namespace Status ==="
	kubectl get namespace avakart ingress-nginx 2>/dev/null || echo "Namespaces not found"
	@echo ""
	@echo "=== Services Status ==="
	kubectl get services -n avakart 2>/dev/null || echo "No services found in avakart namespace"
	@echo ""
	@echo "=== Ingress Controller Status ==="
	kubectl get pods,services -n ingress-nginx 2>/dev/null || echo "Ingress Controller not found"
	@echo ""
	@echo "=== Ingress Rules Status ==="
	kubectl get ingress -n avakart 2>/dev/null || echo "No ingress rules found"
	@echo ""
	@echo "=== StatefulSet Status ==="
	kubectl get statefulset,pods -n avakart -l app=database 2>/dev/null || echo "No StatefulSet found"
	@echo ""
	@echo "=== SSL/TLS Secrets Status ==="
	kubectl get secrets -n avakart -l component=tls 2>/dev/null || echo "No TLS secrets found"


test-services:
	@echo "Testing service connectivity..."
	@echo "Creating test pod for service connectivity tests..."
	kubectl run test-pod --image=busybox --rm -i --tty --restart=Never -n avakart -- /bin/sh -c "\
		echo 'Testing ClusterIP services:'; \
		echo 'Books service:'; nslookup books-service || echo 'Books service not found'; \
		echo 'Storage service:'; nslookup storage-service || echo 'Storage service not found'; \
		echo 'Order processor service:'; nslookup order-processor-service || echo 'Order processor service not found'; \
		echo 'Database headless service:'; nslookup database-headless || echo 'Database service not found'; \
	" 2>/dev/null || echo "Test completed"

test-ingress:
	@echo "Testing Ingress routing..."
	@echo "Checking Ingress Controller status..."
	kubectl get pods -n ingress-nginx -l app=nginx-ingress-controller
	@echo ""
	@echo "Checking Ingress rules..."
	kubectl describe ingress -n avakart
	@echo ""
	@echo "To test external access, you can use:"
	@echo "  kubectl port-forward -n ingress-nginx service/ingress-nginx 8080:80"
	@echo "  Then access http://localhost:8080 with Host header set to your domain"

clean:
	@echo "Cleaning up Chapter 16 resources..."
	-kubectl delete -f advanced/ --ignore-not-found=true
	-kubectl delete -f ingress/ --ignore-not-found=true
	-kubectl delete -f ssl/ --ignore-not-found=true
	-kubectl delete -f stateful/ --ignore-not-found=true
	-kubectl delete -f services/ --ignore-not-found=true
	@echo "Chapter 16 resources cleaned up!"

clean-all: clean
	@echo "Cleaning up namespaces..."
	-kubectl delete namespace avakart --ignore-not-found=true
	-kubectl delete namespace ingress-nginx --ignore-not-found=true
	@echo "All resources cleaned up!"


deploy: create-namespace deploy-services deploy-ssl
	@echo "Deployment complete!"
	@echo "Services and SSL certificates are ready for testing."
