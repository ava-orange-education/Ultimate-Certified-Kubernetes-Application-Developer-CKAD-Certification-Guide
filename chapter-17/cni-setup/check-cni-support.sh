#!/bin/bash

# CNI Network Policy Support Check Script
# This script checks if the current CNI plugin supports Network Policies

set -e

echo "=========================================="
echo "CNI Network Policy Support Check"
echo "=========================================="
echo ""

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to check CNI plugin
check_cni_plugin() {
    echo -e "${BLUE}1. Checking CNI Plugin Configuration${NC}"
    echo "======================================"
    
    # Check if CNI config exists
    if kubectl get nodes -o wide >/dev/null 2>&1; then
        echo -e "${GREEN}✓${NC} Kubernetes cluster is accessible"
        
        # Get node information
        echo ""
        echo "Node Information:"
        kubectl get nodes -o custom-columns=NAME:.metadata.name,VERSION:.status.nodeInfo.kubeletVersion,CONTAINER-RUNTIME:.status.nodeInfo.containerRuntimeVersion,OS:.status.nodeInfo.osImage --no-headers
        
        # Check for CNI-related pods
        echo ""
        echo "CNI-related pods in kube-system:"
        kubectl get pods -n kube-system | grep -E "(calico|cilium|weave|flannel|antrea)" || echo "No common CNI pods found"
        
    else
        echo -e "${RED}✗${NC} Cannot access Kubernetes cluster"
        exit 1
    fi
}

# Function to check Network Policy support
check_network_policy_support() {
    echo ""
    echo -e "${BLUE}2. Checking Network Policy Support${NC}"
    echo "===================================="
    
    # Check if NetworkPolicy CRD exists
    if kubectl api-resources | grep -q networkpolicies; then
        echo -e "${GREEN}✓${NC} NetworkPolicy API resource is available"
        
        # Check API version
        API_VERSION=$(kubectl api-versions | grep networking.k8s.io || echo "Not found")
        echo "Available networking API versions: $API_VERSION"
        
    else
        echo -e "${RED}✗${NC} NetworkPolicy API resource is not available"
        return 1
    fi
    
    # Try to create a test network policy
    echo ""
    echo "Testing Network Policy creation..."
    
    cat <<EOF | kubectl apply -f - >/dev/null 2>&1
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: test-network-policy
  namespace: default
spec:
  podSelector:
    matchLabels:
      app: test
  policyTypes:
  - Ingress
EOF
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓${NC} Test NetworkPolicy created successfully"
        kubectl delete networkpolicy test-network-policy -n default >/dev/null 2>&1
        echo -e "${GREEN}✓${NC} Test NetworkPolicy deleted successfully"
    else
        echo -e "${RED}✗${NC} Failed to create test NetworkPolicy"
        return 1
    fi
}

# Function to identify CNI plugin
identify_cni_plugin() {
    echo ""
    echo -e "${BLUE}3. Identifying CNI Plugin${NC}"
    echo "=========================="
    
    # Check for Calico
    if kubectl get pods -n kube-system | grep -q calico; then
        echo -e "${GREEN}✓${NC} Calico CNI detected"
        CALICO_VERSION=$(kubectl get pods -n kube-system -l k8s-app=calico-node -o jsonpath='{.items[0].spec.containers[0].image}' 2>/dev/null | cut -d':' -f2 || echo "unknown")
        echo "  Version: $CALICO_VERSION"
        echo "  Network Policy Support: YES"
        return 0
    fi
    
    # Check for Cilium
    if kubectl get pods -n kube-system | grep -q cilium; then
        echo -e "${GREEN}✓${NC} Cilium CNI detected"
        CILIUM_VERSION=$(kubectl get pods -n kube-system -l k8s-app=cilium -o jsonpath='{.items[0].spec.containers[0].image}' 2>/dev/null | cut -d':' -f2 || echo "unknown")
        echo "  Version: $CILIUM_VERSION"
        echo "  Network Policy Support: YES"
        return 0
    fi
    
    # Check for Weave Net
    if kubectl get pods -n kube-system | grep -q weave; then
        echo -e "${GREEN}✓${NC} Weave Net CNI detected"
        WEAVE_VERSION=$(kubectl get pods -n kube-system -l name=weave-net -o jsonpath='{.items[0].spec.containers[0].image}' 2>/dev/null | cut -d':' -f2 || echo "unknown")
        echo "  Version: $WEAVE_VERSION"
        echo "  Network Policy Support: YES"
        return 0
    fi
    
    # Check for Antrea
    if kubectl get pods -n kube-system | grep -q antrea; then
        echo -e "${GREEN}✓${NC} Antrea CNI detected"
        ANTREA_VERSION=$(kubectl get pods -n kube-system -l app=antrea -o jsonpath='{.items[0].spec.containers[0].image}' 2>/dev/null | cut -d':' -f2 || echo "unknown")
        echo "  Version: $ANTREA_VERSION"
        echo "  Network Policy Support: YES"
        return 0
    fi
    
    # Check for Flannel (does not support Network Policies)
    if kubectl get pods -n kube-system | grep -q flannel; then
        echo -e "${YELLOW}⚠${NC} Flannel CNI detected"
        FLANNEL_VERSION=$(kubectl get pods -n kube-system -l app=flannel -o jsonpath='{.items[0].spec.containers[0].image}' 2>/dev/null | cut -d':' -f2 || echo "unknown")
        echo "  Version: $FLANNEL_VERSION"
        echo -e "  Network Policy Support: ${RED}NO${NC}"
        echo ""
        echo -e "${YELLOW}WARNING:${NC} Flannel does not support Kubernetes Network Policies"
        echo "Consider switching to Calico, Cilium, Weave Net, or Antrea for Network Policy support"
        return 1
    fi
    
    # Unknown CNI
    echo -e "${YELLOW}⚠${NC} Unknown or custom CNI plugin detected"
    echo "  Please verify manually if your CNI plugin supports Network Policies"
    return 1
}

# Function to test Network Policy enforcement
test_network_policy_enforcement() {
    echo ""
    echo -e "${BLUE}4. Testing Network Policy Enforcement${NC}"
    echo "======================================"
    
    # Create test namespace
    kubectl create namespace netpol-test >/dev/null 2>&1 || true
    
    # Create test pods
    echo "Creating test pods..."
    
    cat <<EOF | kubectl apply -f - >/dev/null 2>&1
apiVersion: v1
kind: Pod
metadata:
  name: test-client
  namespace: netpol-test
  labels:
    app: client
spec:
  containers:
  - name: client
    image: busybox
    command: ["sleep", "3600"]
---
apiVersion: v1
kind: Pod
metadata:
  name: test-server
  namespace: netpol-test
  labels:
    app: server
spec:
  containers:
  - name: server
    image: nginx
    ports:
    - containerPort: 80
---
apiVersion: v1
kind: Service
metadata:
  name: test-server-service
  namespace: netpol-test
spec:
  selector:
    app: server
  ports:
  - port: 80
    targetPort: 80
EOF
    
    # Wait for pods to be ready
    echo "Waiting for test pods to be ready..."
    kubectl wait --for=condition=Ready pod/test-client -n netpol-test --timeout=60s >/dev/null 2>&1
    kubectl wait --for=condition=Ready pod/test-server -n netpol-test --timeout=60s >/dev/null 2>&1
    
    if [ $? -eq 0 ]; then
        echo -e "${GREEN}✓${NC} Test pods are ready"
        
        # Test connectivity before applying network policy
        echo "Testing connectivity before Network Policy..."
        if kubectl exec -n netpol-test test-client -- timeout 5 wget -q --spider http://test-server-service >/dev/null 2>&1; then
            echo -e "${GREEN}✓${NC} Connectivity works before Network Policy"
            
            # Apply restrictive network policy
            echo "Applying restrictive Network Policy..."
            cat <<EOF | kubectl apply -f - >/dev/null 2>&1
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: deny-all
  namespace: netpol-test
spec:
  podSelector: {}
  policyTypes:
  - Ingress
  - Egress
EOF
            
            # Wait a moment for policy to take effect
            sleep 5
            
            # Test connectivity after applying network policy
            echo "Testing connectivity after Network Policy..."
            if kubectl exec -n netpol-test test-client -- timeout 5 wget -q --spider http://test-server-service >/dev/null 2>&1; then
                echo -e "${RED}✗${NC} Connectivity still works after Network Policy (Policy not enforced)"
                POLICY_WORKING=false
            else
                echo -e "${GREEN}✓${NC} Connectivity blocked after Network Policy (Policy enforced)"
                POLICY_WORKING=true
            fi
            
        else
            echo -e "${YELLOW}⚠${NC} Initial connectivity test failed"
            POLICY_WORKING="unknown"
        fi
        
    else
        echo -e "${RED}✗${NC} Test pods failed to start"
        POLICY_WORKING="unknown"
    fi
    
    # Cleanup
    echo "Cleaning up test resources..."
    kubectl delete namespace netpol-test >/dev/null 2>&1 || true
    
    return 0
}

# Function to provide recommendations
provide_recommendations() {
    echo ""
    echo -e "${BLUE}5. Recommendations${NC}"
    echo "=================="
    
    if [ "$POLICY_WORKING" = "true" ]; then
        echo -e "${GREEN}✓${NC} Your cluster supports Network Policies and enforcement is working"
        echo ""
        echo "Next steps:"
        echo "1. Review the network policies in chapter-17/network-policies/"
        echo "2. Apply the appropriate policies for your environment"
        echo "3. Test the policies using the provided test scripts"
        echo "4. Set up monitoring for network policy compliance"
        
    elif [ "$POLICY_WORKING" = "false" ]; then
        echo -e "${RED}✗${NC} Network Policies are not being enforced"
        echo ""
        echo "Possible issues:"
        echo "1. CNI plugin doesn't support Network Policies"
        echo "2. Network Policy controller is not running"
        echo "3. CNI plugin is not properly configured"
        echo ""
        echo "Solutions:"
        echo "1. Install a CNI plugin that supports Network Policies (Calico, Cilium, etc.)"
        echo "2. Check CNI plugin logs for errors"
        echo "3. Verify CNI plugin configuration"
        
    else
        echo -e "${YELLOW}⚠${NC} Unable to determine Network Policy support status"
        echo ""
        echo "Manual verification required:"
        echo "1. Check your CNI plugin documentation"
        echo "2. Test Network Policies manually"
        echo "3. Review CNI plugin logs"
    fi
    
    echo ""
    echo "Supported CNI plugins for Network Policies:"
    echo "- Calico (recommended for this chapter)"
    echo "- Cilium"
    echo "- Weave Net"
    echo "- Antrea"
    echo "- Azure CNI (on AKS)"
    echo "- AWS VPC CNI (on EKS with security groups)"
    echo ""
    echo "CNI plugins that do NOT support Network Policies:"
    echo "- Flannel (standalone)"
    echo "- Bridge"
    echo "- Host-local"
}

# Main execution
main() {
    check_cni_plugin
    
    if check_network_policy_support; then
        identify_cni_plugin
        test_network_policy_enforcement
    else
        echo -e "${RED}Network Policy API is not available${NC}"
        POLICY_WORKING="false"
    fi
    
    provide_recommendations
    
    echo ""
    echo "=========================================="
    echo "CNI Network Policy Support Check Complete"
    echo "=========================================="
}

# Run the main function
main
