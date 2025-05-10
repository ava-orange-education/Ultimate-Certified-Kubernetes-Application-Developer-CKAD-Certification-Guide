#!/bin/bash

# Script to clean up all security manifests for AvaKart

# Set script to exit immediately if a command fails
set -e

echo "Cleaning up AvaKart Security Manifests..."

# Clean up in reverse order of application

# Clean up deployments
echo "Cleaning up Deployments..."
kubectl delete -f frontend/deployment.yaml --ignore-not-found
kubectl delete -f order-processing/deployment.yaml --ignore-not-found
kubectl delete -f storage/deployment.yaml --ignore-not-found
kubectl delete -f books/deployment.yaml --ignore-not-found

# Clean up network policies
echo "Cleaning up Network Policies..."
kubectl delete -f frontend/frontend-network-policy.yaml --ignore-not-found
kubectl delete -f order-processing/order-processor-network-policy.yaml --ignore-not-found
kubectl delete -f storage/storage-network-policy.yaml --ignore-not-found
kubectl delete -f books/books-network-policy.yaml --ignore-not-found
kubectl delete -f default-deny-policy.yaml --ignore-not-found

# Clean up RBAC configurations
echo "Cleaning up RBAC configurations..."
kubectl delete -f storage/storage-manager-binding.yaml --ignore-not-found
kubectl delete -f storage/storage-manager-role.yaml --ignore-not-found
kubectl delete -f books/books-reader-binding.yaml --ignore-not-found
kubectl delete -f books/books-reader-role.yaml --ignore-not-found

# Clean up service accounts
echo "Cleaning up Service Accounts..."
kubectl delete -f frontend/frontend-service-account.yaml --ignore-not-found
kubectl delete -f order-processing/order-processor-service-account.yaml --ignore-not-found
kubectl delete -f storage/storage-service-account.yaml --ignore-not-found
kubectl delete -f books/books-service-account.yaml --ignore-not-found

echo "All security manifests cleaned up successfully!"
