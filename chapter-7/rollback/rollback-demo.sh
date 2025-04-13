#!/bin/bash
# rollback-demo.sh

# Deploy initial version
echo "Deploying books-service v1..."
kubectl apply -f books-deployment-v1.yaml
kubectl rollout status deployment/books-deployment

# Wait for deployment to stabilize
sleep 10

# Update to v2 (with potential issues)
echo "Updating to books-service v2..."
kubectl apply -f books-deployment-v2.yaml
kubectl rollout status deployment/books-deployment

# Simulate monitoring and detecting issues
echo "Monitoring deployment for issues..."
sleep 5
echo "Issues detected in v2! Initiating rollback..."

# Perform rollback
kubectl rollout undo deployment/books-deployment
echo "Rolling back to previous version..."
kubectl rollout status deployment/books-deployment

# Verify rollback success
echo "Rollback completed. Current deployment status:"
kubectl get pods -l app=books-service
kubectl describe deployment books-deployment | grep -A5 "Events:"

# Show revision history
echo "Deployment revision history:"
kubectl rollout history deployment/books-deployment

# Example of rolling back to a specific revision
echo "Example command to rollback to a specific revision:"
echo "kubectl rollout undo deployment/books-deployment --to-revision=2"
