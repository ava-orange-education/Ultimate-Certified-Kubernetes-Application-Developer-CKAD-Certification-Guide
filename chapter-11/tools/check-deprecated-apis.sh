#!/bin/bash
# check-deprecated-apis.sh

# Define known deprecated APIs
declare -A DEPRECATED_APIS=(
  ["apps/v1beta1"]="apps/v1"
  ["apps/v1beta2"]="apps/v1"
  ["extensions/v1beta1,Deployment"]="apps/v1"
  ["extensions/v1beta1,Ingress"]="networking.k8s.io/v1"
  ["networking.k8s.io/v1beta1,Ingress"]="networking.k8s.io/v1"
  ["batch/v1beta1,CronJob"]="batch/v1"
  ["policy/v1beta1,PodDisruptionBudget"]="policy/v1"
  ["policy/v1beta1,PodSecurityPolicy"]="REMOVED"
  ["autoscaling/v2beta1"]="autoscaling/v2"
  ["autoscaling/v2beta2"]="autoscaling/v2"
)

# Check all YAML files in the specified directory
check_directory() {
  local dir=$1
  echo "Checking for deprecated APIs in $dir..."
  
  for file in $(find $dir -name "*.yaml" -o -name "*.yml"); do
    echo "Analyzing $file..."
    
    # Extract apiVersion and kind
    api_version=$(grep -E "^apiVersion:" $file | head -1 | awk '{print $2}')
    kind=$(grep -E "^kind:" $file | head -1 | awk '{print $2}')
    
    if [ -n "$api_version" ] && [ -n "$kind" ]; then
      # Check for direct match
      replacement=${DEPRECATED_APIS["$api_version"]}
      
      # Check for specific kind match
      if [ -z "$replacement" ]; then
        replacement=${DEPRECATED_APIS["$api_version,$kind"]}
      fi
      
      if [ -n "$replacement" ]; then
        if [ "$replacement" == "REMOVED" ]; then
          echo "WARNING: $file uses $api_version $kind which has been removed with no direct replacement"
        else
          echo "WARNING: $file uses deprecated API $api_version for $kind, should use $replacement"
        fi
      fi
    fi
  done
}

# Main execution
if [ -z "$1" ]; then
  echo "Usage: $0 <directory>"
  exit 1
fi

check_directory "$1"
echo "Deprecation check complete."
