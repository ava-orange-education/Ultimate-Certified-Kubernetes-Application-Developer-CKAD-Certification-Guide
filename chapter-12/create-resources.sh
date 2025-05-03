#!/bin/bash
# Script to demonstrate creating ConfigMaps and Secrets from the command line

# Set -e to exit on error
set -e

echo "Creating ConfigMaps and Secrets from the command line"
echo "===================================================="

# Create a directory for temporary files
mkdir -p temp

# ===== Creating ConfigMaps =====

echo -e "\n1. Creating ConfigMaps from literal values"
kubectl create configmap books-service-config-cli \
  --from-literal=PAGE_SIZE=20 \
  --from-literal=ENABLE_CACHE=true \
  --from-literal=LOG_LEVEL=info \
  --from-literal=API_VERSION=v1 \
  --dry-run=client -o yaml

echo -e "\n2. Creating ConfigMaps from files"
# Create a configuration file
cat > temp/storage-config.properties << EOF
data.path=/data
max.connections=50
timeout.seconds=30
retry.attempts=3
EOF

kubectl create configmap storage-service-config-cli \
  --from-file=temp/storage-config.properties \
  --dry-run=client -o yaml

echo -e "\n3. Creating ConfigMaps from multiple files"
# Create multiple configuration files
mkdir -p temp/config-files
cat > temp/config-files/database.properties << EOF
db.host=postgres
db.port=5432
db.name=avakart
db.pool.size=10
EOF

cat > temp/config-files/cache.properties << EOF
cache.enabled=true
cache.ttl=300
cache.max.size=1000
EOF

kubectl create configmap order-processor-config-cli \
  --from-file=temp/config-files/ \
  --dry-run=client -o yaml

# ===== Creating Secrets =====

echo -e "\n4. Creating Secrets from literal values"
kubectl create secret generic books-db-credentials-cli \
  --from-literal=username=admin \
  --from-literal=password=password123 \
  --dry-run=client -o yaml

echo -e "\n5. Creating Secrets from files"
# Create files with sensitive data
echo -n 'admin' > temp/username.txt
echo -n 'password123' > temp/password.txt

kubectl create secret generic storage-db-credentials-cli \
  --from-file=username=temp/username.txt \
  --from-file=password=temp/password.txt \
  --dry-run=client -o yaml

echo -e "\n6. Creating TLS Secrets"
# Generate self-signed certificate (if openssl is available)
if command -v openssl &> /dev/null; then
  openssl req -x509 -nodes -days 365 -newkey rsa:2048 \
    -keyout temp/tls.key -out temp/tls.crt -subj "/CN=avakart.example.com" \
    -addext "subjectAltName = DNS:avakart.example.com,DNS:avakart.local,IP:127.0.0.1"

  kubectl create secret tls avakart-tls-cli \
    --cert=temp/tls.crt \
    --key=temp/tls.key \
    --dry-run=client -o yaml
else
  echo "OpenSSL not available, skipping TLS Secret creation"
fi

# ===== Cleanup =====
echo -e "\nCleaning up temporary files"
rm -rf temp

echo -e "\nNote: All commands were run with --dry-run=client, so no actual resources were created."
echo "To create the resources, remove the --dry-run=client flag from the commands."
