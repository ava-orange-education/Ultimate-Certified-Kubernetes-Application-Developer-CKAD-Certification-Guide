# Chapter 11: Managing API Deprecations

This directory contains examples and tools for managing Kubernetes API deprecations, as covered in Chapter 11 of the Ultimate Certified Kubernetes Application Developer (CKAD) Certification Guide.

## Directory Structure

- `deprecated-apis/`: Contains manifests using deprecated API versions
- `updated-apis/`: Contains equivalent manifests using current API versions
- `tools/`: Contains scripts for detecting deprecated APIs

## API Versioning in Kubernetes

Kubernetes uses API versioning to indicate the stability and support level of its resources:

- **Alpha (v1alpha1, v1alpha2, etc.)**
  - May be buggy
  - Support may be dropped at any time
  - Not recommended for production
  - Feature gates often required to enable

- **Beta (v1beta1, v1beta2, etc.)**
  - Well-tested
  - Support guaranteed for at least 2 releases
  - Semantics may change in incompatible ways
  - Recommended for non-critical business applications

- **Stable (v1, v2, etc.)**
  - Recommended for production
  - Will appear in released software for many versions

## Deprecation Policy

- API elements may only be removed by incrementing the API version
- API objects must be able to round-trip between API versions without information loss
- Older API versions must be supported after newer versions are introduced
- Deprecated APIs are supported for:
  - Beta features: 3 releases or 9 months (whichever is longer)
  - GA features: 12 months or 3 releases (whichever is longer)

## Common Deprecated APIs and Their Replacements

| Deprecated API | Replacement | Notes |
|----------------|-------------|-------|
| apps/v1beta1 | apps/v1 | For Deployments, StatefulSets, etc. |
| apps/v1beta2 | apps/v1 | For Deployments, StatefulSets, etc. |
| extensions/v1beta1 (Ingress) | networking.k8s.io/v1 | Significant structural changes |
| networking.k8s.io/v1beta1 (Ingress) | networking.k8s.io/v1 | Requires pathType field |
| batch/v1beta1 (CronJob) | batch/v1 | Minimal changes |
| policy/v1beta1 (PodSecurityPolicy) | REMOVED | Replaced by Pod Security Standards |

## Key Changes by Kubernetes Version

### Kubernetes 1.22
- Removed beta APIs:
  - Ingress: extensions/v1beta1 → networking.k8s.io/v1
  - CronJob: batch/v1beta1 → batch/v1
  - PodSecurityPolicy: policy/v1beta1 (removed with no replacement)

### Kubernetes 1.25
- Removed beta APIs:
  - PodSecurityPolicy: policy/v1beta1 (removed completely)

### Kubernetes 1.26
- Deprecated APIs:
  - FlowSchema: flowcontrol.apiserver.k8s.io/v1beta1 → flowcontrol.apiserver.k8s.io/v1beta2
  - PriorityLevelConfiguration: flowcontrol.apiserver.k8s.io/v1beta1 → flowcontrol.apiserver.k8s.io/v1beta2

## Using the Deprecation Check Tool

The `tools/check-deprecated-apis.sh` script can be used to scan your Kubernetes manifests for deprecated APIs:

```bash
# Make the script executable
chmod +x tools/check-deprecated-apis.sh

# Check a specific directory
./tools/check-deprecated-apis.sh path/to/your/manifests

# Check the deprecated-apis directory (example)
./tools/check-deprecated-apis.sh deprecated-apis
```

## Best Practices for API Version Management

1. **Always use the latest stable API version**
   - Prefer GA (v1, v2) versions over Beta (v1beta1) or Alpha (v1alpha1)
   - Avoid using deprecated API versions in new code

2. **Regularly audit your manifests**
   - Run deprecation checks as part of CI/CD
   - Update manifests proactively before deprecations affect you

3. **Test with future Kubernetes versions**
   - Set up a test cluster with the next Kubernetes version
   - Validate all manifests against the newer version

4. **Use version control for manifests**
   - Track API version changes in commit messages
   - Use branches for testing API migrations

5. **Subscribe to Kubernetes release announcements**
   - Join kubernetes-announce mailing list
   - Follow the Kubernetes blog

6. **Plan for API migrations**
   - Schedule regular maintenance windows for API updates
   - Include API version updates in your upgrade planning

7. **Use tools to detect deprecated APIs**
   - Implement automated checks in CI/CD pipelines
   - Use kubectl convert to test migrations

8. **Document API version requirements**
   - Specify the minimum and maximum Kubernetes versions supported
   - Document any special handling for API compatibility

## Examples

This directory contains examples of deprecated APIs and their updated versions:

1. **Deployments**: Migration from apps/v1beta1 to apps/v1
2. **Ingress**: Migration from extensions/v1beta1 to networking.k8s.io/v1
3. **CronJob**: Migration from batch/v1beta1 to batch/v1

Each example demonstrates the specific changes required when migrating from deprecated to current API versions.
