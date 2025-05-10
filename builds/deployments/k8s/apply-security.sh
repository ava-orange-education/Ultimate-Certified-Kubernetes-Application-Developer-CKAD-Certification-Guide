#!/bin/bash

# Script to apply all security manifests for AvaKart

# Set script to exit immediately if a command fails
set -e

echo "Applying AvaKart Security Manifests..."

# Create namespace if it doesn't exist
kubectl get namespace default > /dev/null 2>&1 || kubectl create namespace default

# Apply service accounts
echo "Applying Service Accounts..."
kubectl apply -f books/books-service-account.yaml
kubectl apply -f storage/storage-service-account.yaml
kubectl apply -f order-processing/order-processor-service-account.yaml
kubectl apply -f frontend/frontend-service-account.yaml

# Apply RBAC configurations
echo "Applying RBAC configurations..."
kubectl apply -f books/books-reader-role.yaml
kubectl apply -f books/books-reader-binding.yaml
kubectl apply -f storage/storage-manager-role.yaml
kubectl apply -f storage/storage-manager-binding.yaml

# Apply network policies (default deny first)
echo "Applying Network Policies..."
kubectl apply -f default-deny-policy.yaml
kubectl apply -f books/books-network-policy.yaml
kubectl apply -f storage/storage-network-policy.yaml
kubectl apply -f order-processing/order-processor-network-policy.yaml
kubectl apply -f frontend/frontend-network-policy.yaml

# Apply secure deployments
echo "Applying Secure Deployments..."
kubectl apply -f books/deployment.yaml
kubectl apply -f storage/deployment.yaml
kubectl apply -f order-processing/deployment.yaml
kubectl apply -f frontend/deployment.yaml

echo "All security manifests applied successfully!"
echo "To verify the deployments, run: kubectl get pods"
echo "To check service accounts, run: kubectl get serviceaccounts"
echo "To check RBAC configurations, run: kubectl get roles,rolebindings,clusterroles,clusterrolebindings"
echo "To check network policies, run: kubectl get networkpolicies"
