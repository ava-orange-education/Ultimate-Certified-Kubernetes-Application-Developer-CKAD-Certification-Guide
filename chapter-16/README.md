# Chapter 16: Services and Ingress

This chapter demonstrates Kubernetes Services and Ingress resources for the AvaKart application, covering different service types, Ingress controllers, and advanced networking features.

## Overview

Kubernetes Services provide stable network endpoints for pods and enable service discovery, while Ingress resources manage external HTTP/HTTPS access to services. This chapter covers:

- **Service Types**: ClusterIP, NodePort, LoadBalancer, and Headless services
- **Ingress Controllers**: NGINX Ingress Controller setup and configuration
- **SSL/TLS Termination**: Certificate management and HTTPS configuration
- **Advanced Features**: Canary deployments, rate limiting, and multi-tenant routing

## Directory Structure

```
chapter-16/
├── services/                    # Service manifests
│   ├── books-service-clusterip.yaml
│   ├── storage-service-clusterip.yaml
│   ├── order-processor-service-clusterip.yaml
│   ├── frontend-service-nodeport.yaml
│   ├── books-service-nodeport.yaml
│   ├── frontend-loadbalancer.yaml
│   └── api-gateway-loadbalancer.yaml
├── ingress/                     # Ingress Controller and rules
│   ├── nginx-ingress-controller.yaml
│   ├── nginx-ingress-rbac.yaml
│   ├── nginx-ingress-service.yaml
│   ├── avakart-frontend-ingress.yaml
│   └── avakart-api-ingress.yaml
├── ssl/                         # SSL/TLS certificates
│   ├── avakart-tls-secret.yaml
│   └── avakart-api-tls-secret.yaml
├── advanced/                    # Advanced Ingress features
│   ├── books-service-canary-ingress.yaml
│   ├── avakart-rate-limited-ingress.yaml
│   └── avakart-multi-tenant-ingress.yaml
├── stateful/                    # Stateful resources
│   ├── database-headless-service.yaml
│   └── database-statefulset.yaml
├── Makefile                     # Deployment automation
└── README.md                    # This file
```

## Prerequisites

- Kubernetes cluster (v1.20+)
- kubectl configured to access your cluster
- Basic understanding of Kubernetes networking concepts

## Quick Start

### 1. Deploy All Resources

```bash
# Deploy everything at once
make deploy-all

# Or deploy step by step
make create-namespace
make deploy-services
make deploy-ingress
make deploy-ssl
make deploy-advanced
```

### 2. Check Status

```bash
make status
```

### 3. Test Services

```bash
make test-services
make test-ingress
```

## Service Types Explained

### ClusterIP Services (Internal Communication)

ClusterIP services provide internal cluster communication between pods:

```yaml
# books-service-clusterip.yaml
apiVersion: v1
kind: Service
metadata:
  name: books-service
  namespace: avakart
spec:
  type: ClusterIP
  selector:
    app: books-service
  ports:
  - name: http
    port: 80
    targetPort: 8081
```

**Use Cases:**
- Internal microservice communication
- Database connections
- Backend API calls

### NodePort Services (External Access via Node IPs)

NodePort services expose applications on each node's IP at a static port:

```yaml
# frontend-service-nodeport.yaml
apiVersion: v1
kind: Service
metadata:
  name: frontend-service
  namespace: avakart
spec:
  type: NodePort
  selector:
    app: frontend
  ports:
  - name: http
    port: 80
    targetPort: 3000
    nodePort: 30080
```

**Access:** `http://<node-ip>:30080`

**Use Cases:**
- Development and testing
- Simple external access
- When LoadBalancer is not available

### LoadBalancer Services (Cloud Provider Integration)

LoadBalancer services provision external load balancers in cloud environments:

```yaml
# frontend-loadbalancer.yaml
apiVersion: v1
kind: Service
metadata:
  name: frontend-loadbalancer
  namespace: avakart
  annotations:
    service.beta.kubernetes.io/aws-load-balancer-type: "nlb"
spec:
  type: LoadBalancer
  selector:
    app: frontend
  ports:
  - name: http
    port: 80
    targetPort: 3000
```

**Use Cases:**
- Production external access
- Cloud-native deployments
- Automatic load balancer provisioning

### Headless Services (Direct Pod Access)

Headless services return individual pod IPs instead of a single service IP:

```yaml
# database-headless-service.yaml
apiVersion: v1
kind: Service
metadata:
  name: database-headless
  namespace: avakart
spec:
  clusterIP: None  # Makes it headless
  selector:
    app: database
  ports:
  - name: postgres
    port: 5432
    targetPort: 5432
```

**Use Cases:**
- StatefulSet applications
- Database clusters
- Direct pod-to-pod communication

## Ingress Configuration

### NGINX Ingress Controller

The NGINX Ingress Controller processes Ingress resources and routes traffic:

**Components:**
- **Controller Deployment**: Main NGINX controller pod
- **RBAC**: Permissions for controller to read Kubernetes resources
- **Service**: Exposes the controller (LoadBalancer type)
- **IngressClass**: Defines the controller type

### Basic Ingress Rules

#### Frontend Ingress

```yaml
# avakart-frontend-ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: avakart-frontend-ingress
  namespace: avakart
  annotations:
    nginx.ingress.kubernetes.io/ssl-redirect: "true"
spec:
  ingressClassName: nginx
  tls:
  - hosts:
    - avakart.example.com
    secretName: avakart-tls
  rules:
  - host: avakart.example.com
    http:
      paths:
      - path: /
        pathType: Prefix
        backend:
          service:
            name: frontend-service
            port:
              number: 80
```

#### API Services Ingress (Path-based Routing)

```yaml
# avakart-api-ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: avakart-api-ingress
  namespace: avakart
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$2
    nginx.ingress.kubernetes.io/use-regex: "true"
spec:
  ingressClassName: nginx
  rules:
  - host: api.avakart.example.com
    http:
      paths:
      - path: /books(/|$)(.*)
        pathType: Prefix
        backend:
          service:
            name: books-service
            port:
              number: 80
      - path: /storage(/|$)(.*)
        pathType: Prefix
        backend:
          service:
            name: storage-service
            port:
              number: 80
```

## SSL/TLS Configuration

### TLS Secrets

TLS secrets store SSL certificates for HTTPS termination:

```yaml
# avakart-tls-secret.yaml
apiVersion: v1
kind: Secret
metadata:
  name: avakart-tls
  namespace: avakart
type: kubernetes.io/tls
data:
  tls.crt: <base64-encoded-certificate>
  tls.key: <base64-encoded-private-key>
```

### Certificate Management

**Manual Certificates:**
1. Obtain certificates from a Certificate Authority
2. Base64 encode the certificate and private key
3. Create TLS secrets with encoded data

**Automatic Certificates (cert-manager):**
```yaml
apiVersion: cert-manager.io/v1
kind: Certificate
metadata:
  name: avakart-cert
  namespace: avakart
spec:
  secretName: avakart-tls
  issuerRef:
    name: letsencrypt-prod
    kind: ClusterIssuer
  dnsNames:
  - avakart.example.com
```

## Advanced Ingress Features

### Canary Deployments

Route a percentage of traffic to a new version:

```yaml
# books-service-canary-ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: books-service-canary
  namespace: avakart
  annotations:
    nginx.ingress.kubernetes.io/canary: "true"
    nginx.ingress.kubernetes.io/canary-weight: "10"  # 10% traffic
    nginx.ingress.kubernetes.io/canary-by-header: "X-Canary"
spec:
  ingressClassName: nginx
  rules:
  - host: api.avakart.example.com
    http:
      paths:
      - path: /books
        pathType: Prefix
        backend:
          service:
            name: books-service-canary
            port:
              number: 80
```

**Testing Canary:**
```bash
# Normal traffic (90% to main service)
curl -H "Host: api.avakart.example.com" http://ingress-ip/books

# Force canary (100% to canary service)
curl -H "Host: api.avakart.example.com" -H "X-Canary: always" http://ingress-ip/books
```

### Rate Limiting

Protect services from excessive requests:

```yaml
# avakart-rate-limited-ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: avakart-rate-limited-ingress
  namespace: avakart
  annotations:
    nginx.ingress.kubernetes.io/rate-limit: "100"        # 100 requests
    nginx.ingress.kubernetes.io/rate-limit-window: "1m"  # per minute
    nginx.ingress.kubernetes.io/rate-limit-connections: "10"  # max connections
```

### Multi-tenant Routing

Route traffic based on URL patterns:

```yaml
# avakart-multi-tenant-ingress.yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: avakart-multi-tenant-ingress
  namespace: avakart
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$2
    nginx.ingress.kubernetes.io/configuration-snippet: |
      more_set_headers "X-Tenant-ID: $1";
spec:
  ingressClassName: nginx
  rules:
  - host: avakart.example.com
    http:
      paths:
      - path: /tenant-([a-z0-9]+)(/|$)(.*)
        pathType: Prefix
        backend:
          service:
            name: frontend-service
            port:
              number: 80
```

## StatefulSet with Headless Service

For stateful applications like databases:

```yaml
# database-statefulset.yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: database
  namespace: avakart
spec:
  serviceName: database-headless  # Links to headless service
  replicas: 3
  selector:
    matchLabels:
      app: database
  template:
    metadata:
      labels:
        app: database
    spec:
      containers:
      - name: postgres
        image: postgres:13
        ports:
        - containerPort: 5432
        env:
        - name: POSTGRES_DB
          value: avakart
        volumeMounts:
        - name: data
          mountPath: /var/lib/postgresql/data
  volumeClaimTemplates:
  - metadata:
      name: data
    spec:
      accessModes: ["ReadWriteOnce"]
      resources:
        requests:
          storage: 10Gi
```

## Service Discovery

### DNS-based Discovery

Kubernetes provides automatic DNS resolution:

```bash
# Service FQDN format
<service-name>.<namespace>.svc.cluster.local

# Examples
books-service.avakart.svc.cluster.local
storage-service.avakart.svc.cluster.local

# Short names (within same namespace)
books-service
storage-service
```

### Environment Variables

Kubernetes injects service environment variables:

```bash
BOOKS_SERVICE_SERVICE_HOST=10.96.1.100
BOOKS_SERVICE_SERVICE_PORT=80
STORAGE_SERVICE_SERVICE_HOST=10.96.1.101
STORAGE_SERVICE_SERVICE_PORT=80
```

## Testing and Troubleshooting

### Service Connectivity Tests

```bash
# Test DNS resolution
kubectl run test-pod --image=busybox --rm -it --restart=Never -n avakart -- nslookup books-service

# Test HTTP connectivity
kubectl run test-pod --image=busybox --rm -it --restart=Never -n avakart -- wget -qO- books-service

# Check service endpoints
kubectl get endpoints books-service -n avakart
```

### Ingress Troubleshooting

```bash
# Check Ingress Controller logs
kubectl logs -n ingress-nginx deployment/nginx-ingress-controller

# Describe Ingress resource
kubectl describe ingress avakart-frontend-ingress -n avakart

# Test with port-forward
kubectl port-forward -n ingress-nginx service/ingress-nginx 8080:80

# Test with curl
curl -H "Host: avakart.example.com" http://localhost:8080
```

### Common Issues

1. **Service not accessible**
   - Check selector matches pod labels
   - Verify pods are running and ready
   - Ensure ports match container ports

2. **Ingress not routing**
   - Verify Ingress Controller is running
   - Check Ingress rules syntax
   - Ensure backend services exist

3. **SSL/TLS issues**
   - Verify certificate secret exists
   - Check certificate validity
   - Ensure proper certificate chain

## Best Practices

### Service Design

1. **Use appropriate service types**
   - ClusterIP for internal communication
   - LoadBalancer for production external access
   - NodePort for development/testing

2. **Implement health checks**
   - Configure readiness probes
   - Use liveness probes for unhealthy pods
   - Monitor service endpoint health

3. **Plan for scalability**
   - Use headless services for StatefulSets
   - Implement proper session affinity
   - Consider service mesh for complex routing

### Ingress Best Practices

1. **Security**
   - Always use HTTPS in production
   - Implement proper CORS policies
   - Use rate limiting to prevent abuse

2. **Performance**
   - Enable HTTP/2 support
   - Configure appropriate timeouts
   - Use CDN for static content

3. **Monitoring**
   - Implement comprehensive logging
   - Set up metrics collection
   - Configure alerting for failures

## Cleanup

```bash
# Remove all chapter 16 resources
make clean

# Remove everything including namespaces
make clean-all
```

## Additional Resources

- [Kubernetes Services Documentation](https://kubernetes.io/docs/concepts/services-networking/service/)
- [Kubernetes Ingress Documentation](https://kubernetes.io/docs/concepts/services-networking/ingress/)
- [NGINX Ingress Controller](https://kubernetes.github.io/ingress-nginx/)
- [cert-manager Documentation](https://cert-manager.io/docs/)

## Learning Objectives

After completing this chapter, you should understand:

1. Different Kubernetes service types and their use cases
2. How to configure and deploy Ingress controllers
3. SSL/TLS termination and certificate management
4. Advanced Ingress features like canary deployments and rate limiting
5. Service discovery mechanisms in Kubernetes
6. Best practices for production networking configurations

This knowledge is essential for the CKAD certification and building production-ready Kubernetes applications.
