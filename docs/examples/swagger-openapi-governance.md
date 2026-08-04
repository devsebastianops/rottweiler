# Swagger / OpenAPI Governance

When working with APIs you might want to enforce certain rules on your Swagger or OpenAPI specifications. This example shows how to use Rottweiler to validate an OpenAPI specification against a set of policies.

## Scenario

- **HTTPS only**: All endpoints should be served over HTTPS. This policy checks that the scheme is set to HTTPS.
- **Version path naming**: All endpoints should include a version in the path (e.g., /v1/). This policy checks that the path includes a version prefix.

## Input

```json
{
  "openapi": "3.0.0",
  "info": {
    "title": "User Service",
    "version": "1.0.0"
  },
  "servers": [
    {
      "url": "http://api.example.com"
    }
  ],
  "paths": {
    "/users": {
      "get": {
        "summary": "Get users"
      }
    }
  }
}
```

## Policies

```yaml
policies:
  - name: "api-enforce-https-servers"
    description: "All OpenAPI server definitions must use HTTPS."
    severity: "error"
    rule: "input.servers.all(s, s.url.startsWith('https://'))"

  - name: "api-enforce-versioned-paths"
    description: "All endpoint paths must start with a version prefix like /v1/ or /v2/."
    severity: "warning"
    rule: "input.paths.keys().all(path, path.matches('^/v[0-9]+/.*'))"
```

## Rottweiler Check

```bash
rottweiler check --input openapi.json --policy policies.yaml
```