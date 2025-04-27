# AvaKart Troubleshooting Guide

This guide provides solutions for common Kubernetes issues you might encounter when running the AvaKart application.

## Pod Startup Issues

### Symptoms:
- Pod stuck in `Pending` state
- Pod stuck in `ContainerCreating` state
- Pod in `CrashLoopBackOff` state

### Debugging Steps:

1. **Check pod status and events**:
   ```bash
   kubectl describe pod <pod-name>
   ```
   Look for events that indicate issues like:
   - Resource constraints
   - Image pull failures
   - Volume mount issues

2. **Check logs**:
   ```bash
   kubectl logs <pod-name>
   # If the container has crashed:
   kubectl logs <pod-name> --previous
   ```

3. **Check resource usage**:
   ```bash
   kubectl top node
   kubectl top pod
   ```

### Common Solutions:

- **Image Pull Errors**: Verify image name and registry access
  ```bash
  # Check image name in deployment
  kubectl get deployment <deployment-name> -o jsonpath='{.spec.template.spec.containers[0].image}'
  
  # Verify registry access from within the cluster
  kubectl run test --image=busybox --rm -it -- wget -q -O- https://registry.example.com/v2/
  ```

- **Resource Constraints**: Adjust resource requests/limits or scale cluster
  ```bash
  # Edit deployment resource requests
  kubectl edit deployment <deployment-name>
  ```

- **Volume Mount Issues**: Check PVC status and storage class
  ```bash
  kubectl get pvc
  kubectl get storageclass
  ```

## Service Connectivity Issues

### Symptoms:
- Services cannot communicate with each other
- External access to services fails

### Debugging Steps:

1. **Verify service exists and has endpoints**:
   ```bash
   kubectl get service <service-name>
   kubectl get endpoints <service-name>
   ```

2. **Test connectivity from within the cluster**:
   ```bash
   kubectl run test-pod --image=busybox --rm -it -- wget -qO- http://<service-name>:<port>/health
   ```

3. **Check network policies**:
   ```bash
   kubectl get networkpolicy
   ```

### Common Solutions:

- **No Endpoints**: Verify selector matches pod labels
  ```bash
  # Check service selector
  kubectl get service <service-name> -o jsonpath='{.spec.selector}'
  
  # Check pod labels
  kubectl get pods --show-labels
  ```

- **DNS Issues**: Check CoreDNS/kube-dns is running
  ```bash
  kubectl get pods -n kube-system -l k8s-app=kube-dns
  ```

- **Network Policy Blocking**: Temporarily disable or modify network policies
  ```bash
  kubectl delete networkpolicy <policy-name>
  ```

## Resource Constraints

### Symptoms:
- Pods evicted
- OOMKilled containers
- Slow application performance

### Debugging Steps:

1. **Check resource usage**:
   ```bash
   kubectl top pod
   ```

2. **Review pod resource requests and limits**:
   ```bash
   kubectl describe pod <pod-name> | grep -A 3 Requests
   kubectl describe pod <pod-name> | grep -A 3 Limits
   ```

3. **Check node capacity**:
   ```bash
   kubectl describe node | grep -A 5 Capacity
   kubectl describe node | grep -A 5 Allocatable
   ```

### Common Solutions:

- **Increase Resource Limits**: Adjust deployment resource configuration
  ```bash
  kubectl edit deployment <deployment-name>
  ```

- **Scale Horizontally**: Increase replicas instead of vertical scaling
  ```bash
  kubectl scale deployment <deployment-name> --replicas=3
  ```

- **Memory Leaks**: Check application logs for memory issues
  ```bash
  kubectl logs <pod-name> | grep -i "memory\|leak\|out of memory"
  ```

## Application-Specific Issues

### Books Service Issues:

- **Cannot retrieve book data**: Check connectivity to storage service
  ```bash
  # From books pod
  kubectl exec -it <books-pod> -- wget -qO- http://storage-service:8083/health
  ```

### Order Processor Issues:

- **Orders not being processed**: Check message queue and storage connectivity
  ```bash
  # Check logs for connection errors
  kubectl logs -l app=order-processor | grep -i "error\|connection\|refused"
  ```

### Storage Service Issues:

- **Data persistence issues**: Check volume mounts and PVC status
  ```bash
  kubectl describe pod <storage-pod>
  kubectl get pvc
  ```

## Using Debug Containers

For advanced debugging, you can attach debug containers to running pods:

```bash
# Add a debug container to a running pod
kubectl debug -it <pod-name> --image=busybox --target=<container-name>

# Create a copy of a pod with debugging tools
kubectl debug <pod-name> -it --image=ubuntu --share-processes --copy-to=<debug-pod-name>

# Debug with a specific shell
kubectl debug -it <pod-name> --image=alpine --target=<container-name> -- sh
```

## Collecting Comprehensive Logs

Use the debug-avakart.sh script to collect comprehensive debugging information:

```bash
./debug-avakart.sh > debug-report.txt
```

This will collect pod status, events, logs, and resource usage information for all AvaKart components.
