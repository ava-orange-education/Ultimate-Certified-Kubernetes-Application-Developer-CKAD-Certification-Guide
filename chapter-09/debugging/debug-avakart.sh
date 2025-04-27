#!/bin/bash
# debug-avakart.sh - Comprehensive debugging script for AvaKart application

# Check overall application health
echo "Checking AvaKart application health..."
kubectl get pods -l app.kubernetes.io/part-of=avakart
kubectl get services -l app.kubernetes.io/part-of=avakart

# Check for any failed pods
echo -e "\nChecking for failed pods..."
kubectl get pods -l app.kubernetes.io/part-of=avakart | grep -v "Running\|Completed"

# Check recent events
echo -e "\nChecking recent events..."
kubectl get events --sort-by='.lastTimestamp' | tail -n 20

# Check logs for each component
for component in frontend books-service order-processor storage-service
do
  echo -e "\nChecking logs for $component..."
  kubectl logs -l app=$component --tail=50
done

# Check resource usage
echo -e "\nChecking resource usage..."
kubectl top pods -l app.kubernetes.io/part-of=avakart

# Check pod details for any issues
echo -e "\nChecking pod details for potential issues..."
for pod in $(kubectl get pods -l app.kubernetes.io/part-of=avakart -o name)
do
  echo -e "\nDescribing $pod..."
  kubectl describe $pod | grep -A 5 "State:\|Last State:\|Reason:\|Message:\|Events:"
done

# Check service endpoints
echo -e "\nChecking service endpoints..."
kubectl get endpoints -l app.kubernetes.io/part-of=avakart

echo -e "\nDebug information collected. Check output for potential issues."
