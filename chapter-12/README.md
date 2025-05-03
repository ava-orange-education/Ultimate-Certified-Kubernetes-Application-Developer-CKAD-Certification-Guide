# Chapter 12: ConfigMaps and Secrets

This chapter demonstrates how to manage environment-specific configurations and sensitive data in Kubernetes applications using ConfigMaps and Secrets.

## Directory Structure

```
chapter-12/
├── configmaps/
│   ├── books-service-config.yaml      # ConfigMap for Books Service
│   ├── storage-service-config.yaml    # ConfigMap for Storage Service
│   └── order-processor-config.yaml    # ConfigMap for Order Processor
├── secrets/
│   ├── books-db-credentials.yaml      # Secret for Books Service DB credentials
│   ├── storage-db-credentials.yaml    # Secret for Storage Service DB credentials
│   ├── order-processor-credentials.yaml # Secret for Order Processor credentials
│   └── avakart-tls.yaml               # TLS Secret for AvaKart
└── deployments/
    ├── books-deployment.yaml          # Books Deployment using ConfigMap as env vars and Secret as env vars
    ├── storage-deployment.yaml        # Storage Deployment using ConfigMap as volume and Secret as env vars
    └── order-processor-deployment.yaml # Order Processor using ConfigMap and Secret as volumes
```

## ConfigMaps

ConfigMaps are used to store non-sensitive configuration data. This chapter demonstrates different methods for creating and consuming ConfigMaps:

1. **Creating ConfigMaps from literal values** (books-service-config.yaml)
2. **Creating ConfigMaps from files** (storage-service-config.yaml)
3. **Creating ConfigMaps from multiple files** (order-processor-config.yaml)

## Secrets

Secrets are used to store sensitive information. This chapter demonstrates different methods for creating and consuming Secrets:

1. **Creating Secrets from literal values** (books-db-credentials.yaml)
2. **Creating Secrets for TLS** (avakart-tls.yaml)

## Consuming ConfigMaps and Secrets

This chapter demonstrates different methods for consuming ConfigMaps and Secrets in Pods:

1. **Using ConfigMap as Environment Variables** (books-deployment.yaml)
   - Individual ConfigMap values as environment variables

2. **Using ConfigMap as Volume** (storage-deployment.yaml, order-processor-deployment.yaml)
   - Mounting ConfigMap as a volume

3. **Using Secret as Environment Variables** (books-deployment.yaml, storage-deployment.yaml)
   - Individual Secret values as environment variables
   - All Secret values as environment variables

4. **Using Secret as Volume** (order-processor-deployment.yaml)
   - Mounting Secret as a volume

## Application Code

The application code in the backend services has been designed to read configuration from environment variables and files:

1. **Books Service** (backend/apps/books/main.go)
   - Reads configuration from environment variables

2. **Storage Service** (backend/apps/storage/main.go)
   - Reads configuration from files mounted as volumes

3. **Order Processor** (backend/apps/order-processor/main.go)
   - Reads basic configuration from environment variables
   - Checks for the presence of ConfigMap and Secret volumes
   - Logs when configuration files and secrets are found

## Usage

To apply these manifests, use the following commands:

```bash
# Create ConfigMaps
kubectl apply -f chapter-12/configmaps/

# Create Secrets
kubectl apply -f chapter-12/secrets/

# Create Deployments
kubectl apply -f chapter-12/deployments/
```

## Best Practices

1. **Separate configuration from code**
   - Store all environment-specific configuration in ConfigMaps
   - Avoid hardcoding configuration values in container images

2. **Enhance Secret security**
   - Enable encryption at rest for Secrets
   - Use RBAC to restrict access to Secrets
   - Consider external secret management solutions for production

3. **Minimize Secret exposure**
   - Mount Secrets as volumes instead of environment variables when possible
   - Use specific Secret keys rather than exposing all keys
   - Set volumes to readOnly

4. **Handle configuration updates**
   - Be aware that volume-mounted ConfigMaps update automatically (with some delay)
   - Environment variables from ConfigMaps require pod restart to update
