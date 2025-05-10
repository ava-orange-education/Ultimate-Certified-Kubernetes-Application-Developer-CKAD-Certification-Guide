# Kubernetes Security Exam Preparation Questions

This document contains sample questions and answers related to Kubernetes security to help prepare for the CKAD certification exam.

## Service Accounts

### Question 1
**What is the purpose of a Service Account in Kubernetes?**

**Answer:**
A Service Account provides an identity for processes running in a pod, allowing the pod to interact with the Kubernetes API server. It enables pod-to-API server authentication and allows for fine-grained access control for applications running in pods.

### Question 2
**How do you create a Service Account and assign it to a pod?**

**Answer:**
To create a Service Account:
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: my-service-account
  namespace: default
```

To assign it to a pod:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  serviceAccountName: my-service-account
  containers:
  - name: my-container
    image: nginx
```

### Question 3
**What happens if you don't specify a Service Account for a pod?**

**Answer:**
If you don't specify a Service Account for a pod, Kubernetes automatically assigns the `default` Service Account from the namespace where the pod is created. This default Service Account typically has limited permissions.

## Security Contexts

### Question 4
**What is a Security Context in Kubernetes and at what levels can it be applied?**

**Answer:**
A Security Context in Kubernetes defines privilege and access control settings for pods and containers. It can be applied at two levels:
1. Pod level: Settings apply to all containers in the pod
2. Container level: Settings apply to a specific container and override pod-level settings

### Question 5
**How do you configure a pod to run as a non-root user with a specific user ID?**

**Answer:**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: security-context-demo
spec:
  securityContext:
    runAsNonRoot: true
    runAsUser: 1000
  containers:
  - name: nginx
    image: nginx
```

### Question 6
**How do you prevent privilege escalation in a container?**

**Answer:**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: no-privilege-escalation
spec:
  containers:
  - name: nginx
    image: nginx
    securityContext:
      allowPrivilegeEscalation: false
```

## RBAC (Role-Based Access Control)

### Question 7
**What are the four main API objects used in Kubernetes RBAC?**

**Answer:**
1. Role: Defines permissions within a namespace
2. ClusterRole: Defines permissions across the entire cluster
3. RoleBinding: Binds roles to users, groups, or service accounts within a namespace
4. ClusterRoleBinding: Binds cluster roles to users, groups, or service accounts across the cluster

### Question 8
**How do you create a Role that allows reading pods and services in a namespace?**

**Answer:**
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: pod-reader
rules:
- apiGroups: [""]
  resources: ["pods", "services"]
  verbs: ["get", "list", "watch"]
```

### Question 9
**How do you bind a Role to a Service Account?**

**Answer:**
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-pods-binding
  namespace: default
subjects:
- kind: ServiceAccount
  name: my-service-account
  namespace: default
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```

## Network Policies

### Question 10
**What is a Network Policy in Kubernetes?**

**Answer:**
A Network Policy is a specification of how groups of pods are allowed to communicate with each other and other network endpoints. Network policies provide a way to control the traffic flow at the IP address or port level.

### Question 11
**How do you create a Network Policy that allows incoming traffic only from pods with a specific label?**

**Answer:**
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: allow-from-frontend
spec:
  podSelector:
    matchLabels:
      app: backend
  policyTypes:
  - Ingress
  ingress:
  - from:
    - podSelector:
        matchLabels:
          app: frontend
```

### Question 12
**How do you create a default deny Network Policy for all pods in a namespace?**

**Answer:**
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: default-deny-all
spec:
  podSelector: {}  # Selects all pods in the namespace
  policyTypes:
  - Ingress
  - Egress
  # No ingress or egress rules specified, which means deny all traffic
```

## Pod Security

### Question 13
**How do you configure a read-only root filesystem for a container?**

**Answer:**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: readonly-pod
spec:
  containers:
  - name: nginx
    image: nginx
    securityContext:
      readOnlyRootFilesystem: true
    volumeMounts:
    - name: tmp-volume
      mountPath: /tmp
  volumes:
  - name: tmp-volume
    emptyDir: {}
```

### Question 14
**How do you add or drop Linux capabilities for a container?**

**Answer:**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: capabilities-demo
spec:
  containers:
  - name: nginx
    image: nginx
    securityContext:
      capabilities:
        add: ["NET_BIND_SERVICE"]
        drop: ["ALL"]
```

### Question 15
**What is the purpose of the `fsGroup` field in a pod's security context?**

**Answer:**
The `fsGroup` field specifies a supplemental group ID that will be added to the containers in the pod. Any files created by the containers will be owned by this group ID. This is useful for controlling access to shared storage volumes.

## Secrets and ConfigMaps

### Question 16
**How do you create a Secret and mount it as a volume in a pod?**

**Answer:**
Create a Secret:
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: my-secret
type: Opaque
data:
  username: dXNlcm5hbWU=  # base64 encoded "username"
  password: cGFzc3dvcmQ=  # base64 encoded "password"
```

Mount it as a volume:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: secret-pod
spec:
  containers:
  - name: nginx
    image: nginx
    volumeMounts:
    - name: secret-volume
      mountPath: /etc/secrets
      readOnly: true
  volumes:
  - name: secret-volume
    secret:
      secretName: my-secret
```

### Question 17
**How do you create a ConfigMap and use it as environment variables in a pod?**

**Answer:**
Create a ConfigMap:
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  APP_ENV: production
  APP_DEBUG: "false"
```

Use it as environment variables:
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: config-pod
spec:
  containers:
  - name: app
    image: my-app
    envFrom:
    - configMapRef:
        name: app-config
```

## Practical Exam Tips

### Question 18
**How do you check if a Service Account has permission to perform a specific action?**

**Answer:**
Use the `kubectl auth can-i` command with the `--as` flag:
```bash
kubectl auth can-i get pods --as=system:serviceaccount:default:my-service-account
```

### Question 19
**How do you verify the security context of a running pod?**

**Answer:**
```bash
kubectl get pod <pod-name> -o jsonpath='{.spec.securityContext}'
kubectl get pod <pod-name> -o jsonpath='{.spec.containers[0].securityContext}'
```

### Question 20
**How do you create a temporary pod with specific security settings for debugging?**

**Answer:**
```bash
kubectl run debug-pod --image=busybox --restart=Never --rm -it \
  --overrides='{"spec":{"securityContext":{"runAsUser":1000,"runAsNonRoot":true}}}' \
  -- sh
```

## Conclusion

These questions cover the key security concepts in Kubernetes that are relevant for the CKAD certification exam. Practice implementing these security features in a Kubernetes environment to gain hands-on experience.
