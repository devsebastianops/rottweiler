# Docker Compose Check

Docker Compose files define multi-container Docker applications. Misconfigurations can lead to issues.

## Scenario

- **No exposed database ports**: Database ports should not be exposed to the public. This policy checks that no database service has an exposed port.
- **Healthchecks mandatory**: All services should have a healthcheck defined. This policy checks that each service has a healthcheck.

## Input

```yaml
version: '3.8'
services:
  app:
    image: myapp:v1.0.0
    ports:
      - "8080:8080"
    healthcheck:
      test: ["CMD", "curl", "-f", "http://localhost:8080/health"]

  postgres:
    image: postgres:15
    ports:
      - "5432:5432" # Danger: Database directly exposed to the host!
```

## Policies

```yaml
policies:
  - name: "compose-no-db-host-ports"
    description: "Database ports like 5432 should not be bound to host interfaces."
    severity: "error"
    rule: "input.services.values().all(s, !has(s.ports) || s.ports.all(p, !p.startsWith('5432:')))"

  - name: "compose-require-healthcheck"
    description: "Every service in compose must define a healthcheck block."
    severity: "warning"
    rule: "input.services.values().all(s, has(s.healthcheck))"
```

## Rottweiler Check

```bash
rottweiler check --input docker-compose.yaml --policy policy.yaml
```