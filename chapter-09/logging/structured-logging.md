# Structured Logging in AvaKart

This document outlines the structured logging approach implemented in the AvaKart application to facilitate better monitoring and debugging in Kubernetes environments.

## Structured Logging Overview

Structured logging formats log entries as JSON objects rather than plain text, making them easier to parse, filter, and analyze. This approach provides several benefits:

1. **Machine-Readable**: JSON format is easily parsed by log analysis tools
2. **Consistent Fields**: Each log entry contains the same set of fields
3. **Searchable**: Specific fields can be queried efficiently
4. **Metadata-Rich**: Additional context can be included in each log entry

## Implementation in AvaKart

All AvaKart components have been configured to use structured logging with the following features:

### 1. Environment Variables

Each component uses these environment variables to control logging behavior:

- `LOG_LEVEL`: Controls the verbosity of logs (debug, info, warn, error)
- `LOG_FORMAT`: Specifies the output format (json)

### 2. Log Collector Sidecars

Each pod includes a log-collector sidecar container that:

- Collects logs from the main application container
- Processes and forwards logs as needed
- Provides a unified logging interface

### 3. Shared Log Volumes

Components use shared volumes to make logs accessible between containers:

```yaml
volumes:
- name: shared-logs
  emptyDir: {}
```

## Log Format

AvaKart logs follow this JSON structure:

```json
{
  "timestamp": "2025-04-27T10:45:23.123Z",
  "level": "info",
  "message": "Request processed successfully",
  "component": "books-service",
  "traceId": "abc123",
  "requestId": "req-456",
  "method": "GET",
  "path": "/api/books/123",
  "statusCode": 200,
  "responseTime": 45,
  "userId": "user-789",
  "additionalInfo": {
    "bookId": "123",
    "cacheHit": true
  }
}
```

## Common Fields

| Field | Description |
|-------|-------------|
| timestamp | ISO 8601 timestamp of the log event |
| level | Log level (debug, info, warn, error) |
| message | Human-readable log message |
| component | Service that generated the log (books-service, storage-service, etc.) |
| traceId | Unique ID for tracing requests across services |
| requestId | Unique ID for a specific request |
| method | HTTP method (for API requests) |
| path | Request path (for API requests) |
| statusCode | HTTP status code (for API responses) |
| responseTime | Time taken to process request in milliseconds |

## Accessing Logs

### Basic Log Retrieval

```bash
# Get logs from a specific pod
kubectl logs pod/books-pod

# Follow logs in real-time
kubectl logs -f deployment/books-deployment

# Get logs from a specific container in a multi-container pod
kubectl logs pod/storage-pod -c cache-container

# Get logs with timestamps
kubectl logs deployment/order-processor --timestamps

# Get logs from all pods with a specific label
kubectl logs -l app=frontend --all-containers

# Get logs from previous container instance (if it crashed)
kubectl logs pod/order-processor --previous
```

### Filtering Structured Logs

Since logs are in JSON format, you can use tools like `jq` to filter and analyze them:

```bash
# Extract all error logs
kubectl logs deployment/books-deployment | jq 'select(.level == "error")'

# Find slow requests (response time > 500ms)
kubectl logs deployment/books-deployment | jq 'select(.responseTime > 500)'

# Find logs related to a specific user
kubectl logs deployment/books-deployment | jq 'select(.userId == "user-123")'

# Track a request across services using traceId
kubectl logs -l app.kubernetes.io/part-of=avakart --all-containers | jq 'select(.traceId == "abc123")' | jq -s 'sort_by(.timestamp)'
```

## Best Practices

1. **Use Appropriate Log Levels**: 
   - DEBUG: Detailed information for debugging
   - INFO: General operational information
   - WARN: Warning events that might cause issues
   - ERROR: Error events that might still allow the application to continue
   - FATAL: Severe error events that cause the application to terminate

2. **Include Context**: Always include relevant context in logs (user ID, request ID, etc.)

3. **Avoid Sensitive Information**: Never log passwords, tokens, or personal information

4. **Use Correlation IDs**: Include traceId in all logs to track requests across services

5. **Log at Service Boundaries**: Log incoming and outgoing requests at service boundaries

## Troubleshooting with Logs

1. **Identifying Issues**:
   ```bash
   # Find recent errors
   kubectl logs -l app.kubernetes.io/part-of=avakart --since=1h | jq 'select(.level == "error")'
   ```

2. **Tracing Requests**:
   ```bash
   # Extract a specific request flow
   kubectl logs -l app.kubernetes.io/part-of=avakart --all-containers | jq 'select(.requestId == "req-456")' | jq -s 'sort_by(.timestamp)'
   ```

3. **Performance Analysis**:
   ```bash
   # Find slow endpoints
   kubectl logs deployment/books-deployment | jq 'select(.responseTime != null) | {path: .path, responseTime: .responseTime}' | jq -s 'group_by(.path) | map({path: .[0].path, avgResponseTime: (map(.responseTime) | add / length)}) | sort_by(.avgResponseTime) | reverse | .[0:5]'
   ```
