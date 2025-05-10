#!/bin/bash

# Script to apply all security manifests for Chapter 13

# Set script to exit immediately if a command fails
set -e

echo "Applying Chapter 13 Security Manifests..."

# Create namespace if it doesn't exist
kubectl get namespace default > /dev/null 2>&1 || kubectl create namespace default

# Apply service accounts
echo "Applying Service Accounts..."
kubectl apply -f service-accounts/

# Apply RBAC configurations
echo "Applying RBAC configurations..."
kubectl apply -f rbac/roles/
kubectl apply -f rbac/role-bindings/
kubectl apply -f rbac/cluster-roles/
kubectl apply -f rbac/cluster-role-bindings/

# Apply security context examples
echo "Applying Security Context examples..."
kubectl apply -f security-contexts/

# Apply secure deployments
echo "Applying Secure Deployments..."
kubectl apply -f deployments/

# Apply network policies
echo "Applying Network Policies..."
kubectl apply -f network-policies/

echo "All security manifests applied successfully!"
echo "To verify the deployments, run: kubectl get pods"
echo "To check service accounts, run: kubectl get serviceaccounts"
echo "To check RBAC configurations, run: kubectl get roles,rolebindings,clusterroles,clusterrolebindings"
echo "To check network policies, run: kubectl get networkpolicies"
