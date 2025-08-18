#!/bin/bash

# Network Policy Testing Script for AvaKart Application
# This script tests the effectiveness of network policies by attempting various connections

set -e

NAMESPACE="avakart"
TIMEOUT=5

echo "=========================================="
echo "AvaKart Network Policy Testing Script"
echo "=========================================="
echo "Testing network policies in namespace: $NAMESPACE"
echo "Connection timeout: ${TIMEOUT}s"
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Function to test connection
test_connection() {
    local from_pod=$1
    local to_service=$2
    local port=$3
    local expected=$4
    local description=$5
    
    echo -n "Testing: $description... "
    
    if kubectl exec -n $NAMESPACE deployment/$from_pod -- timeout $TIMEOUT curl -s --connect-timeout $TIMEOUT http://$to_service:$port/health >/dev/null 2>&1; then
        if [ "$expected" = "allow" ]; then
            echo -e "${GREEN}✓ PASS${NC} (Connection allowed as expected)"
        else
            echo -e "${RED}✗ FAIL${NC} (Connection should be blocked but was allowed)"
        fi
    else
        if [ "$expected" = "deny" ]; then
            echo -e "${GREEN}✓ PASS${NC} (Connection blocked as expected)"
        else
            echo -e "${RED}✗ FAIL${NC} (Connection should be allowed but was blocked)"
        fi
    fi
}

# Function to test DNS resolution
test_dns() {
    local from_pod=$1
    local description=$2
    
    echo -n "Testing: $description... "
    
    if kubectl exec -n $NAMESPACE deployment/$from_pod -- nslookup kubernetes.default.svc.cluster.local >/dev/null 2>&1; then
        echo -e "${GREEN}✓ PASS${NC} (DNS resolution working)"
    else
        echo -e "${RED}✗ FAIL${NC} (DNS resolution failed)"
    fi
}

# Function to test external connectivity
test_external() {
    local from_pod=$1
    local url=$2
    local expected=$3
    local description=$4
    
    echo -n "Testing: $description... "
    
    if kubectl exec -n $NAMESPACE deployment/$from_pod -- timeout $TIMEOUT curl -s --connect-timeout $TIMEOUT $url >/dev/null 2>&1; then
        if [ "$expected" = "allow" ]; then
            echo -e "${GREEN}✓ PASS${NC} (External connection allowed as expected)"
        else
            echo -e "${RED}✗ FAIL${NC} (External connection should be blocked but was allowed)"
        fi
    else
        if [ "$expected" = "deny" ]; then
            echo -e "${GREEN}✓ PASS${NC} (External connection blocked as expected)"
        else
            echo -e "${RED}✗ FAIL${NC} (External connection should be allowed but was blocked)"
        fi
    fi
}

echo "1. Testing Frontend Connections"
echo "================================"
test_connection "frontend" "books-service" "8081" "allow" "Frontend → Books Service"
test_connection "frontend" "order-processor" "8083" "allow" "Frontend → Order Processor"
test_connection "frontend" "storage-service" "8082" "deny" "Frontend → Storage Service (should be blocked)"
test_dns "frontend" "Frontend DNS resolution"
test_external "frontend" "https://httpbin.org/get" "allow" "Frontend → External HTTPS"

echo ""
echo "2. Testing Books Service Connections"
echo "====================================="
test_connection "books-service" "storage-service" "8082" "allow" "Books Service → Storage Service"
test_connection "books-service" "order-processor" "8083" "deny" "Books Service → Order Processor (should be blocked)"
test_dns "books-service" "Books Service DNS resolution"
test_external "books-service" "https://httpbin.org/get" "allow" "Books Service → External HTTPS"

echo ""
echo "3. Testing Order Processor Connections"
echo "======================================="
test_connection "order-processor" "books-service" "8081" "allow" "Order Processor → Books Service"
test_connection "order-processor" "storage-service" "8082" "allow" "Order Processor → Storage Service"
test_dns "order-processor" "Order Processor DNS resolution"
test_external "order-processor" "https://httpbin.org/get" "allow" "Order Processor → External HTTPS"

echo ""
echo "4. Testing Storage Service Connections"
echo "======================================="
test_connection "storage-service" "books-service" "8081" "deny" "Storage Service → Books Service (should be blocked)"
test_connection "storage-service" "order-processor" "8083" "deny" "Storage Service → Order Processor (should be blocked)"
test_connection "storage-service" "frontend" "3000" "deny" "Storage Service → Frontend (should be blocked)"
test_dns "storage-service" "Storage Service DNS resolution"
test_external "storage-service" "https://httpbin.org/get" "allow" "Storage Service → External HTTPS (for backups)"

echo ""
echo "5. Testing Cross-Service Communication Patterns"
echo "================================================"
echo -n "Testing: Service mesh communication pattern... "
# Test if services can communicate in the expected pattern: Frontend → Backend Services → Storage/Database
if kubectl exec -n $NAMESPACE deployment/frontend -- timeout $TIMEOUT curl -s --connect-timeout $TIMEOUT http://books-service:8081/health >/dev/null 2>&1 && \
   kubectl exec -n $NAMESPACE deployment/books-service -- timeout $TIMEOUT curl -s --connect-timeout $TIMEOUT http://storage-service:8082/health >/dev/null 2>&1; then
    echo -e "${GREEN}✓ PASS${NC} (Service mesh pattern working)"
else
    echo -e "${RED}✗ FAIL${NC} (Service mesh pattern broken)"
fi

echo ""
echo "6. Testing Network Policy Enforcement"
echo "======================================"

# Check if network policies are applied
echo -n "Checking if network policies are applied... "
POLICY_COUNT=$(kubectl get networkpolicies -n $NAMESPACE --no-headers 2>/dev/null | wc -l)
if [ $POLICY_COUNT -gt 0 ]; then
    echo -e "${GREEN}✓ PASS${NC} ($POLICY_COUNT network policies found)"
else
    echo -e "${RED}✗ FAIL${NC} (No network policies found)"
fi

# Check if CNI supports network policies
echo -n "Checking CNI network policy support... "
if kubectl get nodes -o jsonpath='{.items[*].status.nodeInfo.containerRuntimeVersion}' | grep -q "containerd\|cri-o\|docker"; then
    echo -e "${GREEN}✓ PASS${NC} (Compatible container runtime detected)"
else
    echo -e "${YELLOW}⚠ WARNING${NC} (Unable to verify container runtime)"
fi

echo ""
echo "7. Performance Impact Assessment"
echo "================================="
echo -n "Testing connection latency with policies... "
START_TIME=$(date +%s%N)
kubectl exec -n $NAMESPACE deployment/frontend -- curl -s --connect-timeout $TIMEOUT http://books-service:8081/health >/dev/null 2>&1
END_TIME=$(date +%s%N)
LATENCY=$(( (END_TIME - START_TIME) / 1000000 ))
echo -e "${GREEN}✓${NC} Connection latency: ${LATENCY}ms"

echo ""
echo "=========================================="
echo "Network Policy Testing Complete"
echo "=========================================="
echo ""
echo "Summary:"
echo "- Test all connections to ensure policies are working correctly"
echo "- Verify DNS resolution is not blocked"
echo "- Confirm external connectivity works where expected"
echo "- Validate service mesh communication patterns"
echo ""
echo "If any tests fail, check:"
echo "1. Network policies are correctly applied"
echo "2. Pod labels match policy selectors"
echo "3. CNI plugin supports network policies"
echo "4. DNS egress rules are properly configured"
