# GitHub Workflow Validation

CI/CD workflows have high privileges. Vulnerable step definitions can open the door to supply-chain attacks.

## Scenario

We want to lint GitHub Actions workflows before they are committed:

- **No wildcard permissions**: Workflows must not use permissions: write-all.
- **Timeouts**: Every job must explicitly set a timeout-minutes threshold to avoid hanging jobs consuming pipeline minutes.

## Input

```yaml
name: CI Pipeline
on: [push]

jobs:
  build:
    runs-on: ubuntu-latest
    permissions: write-all # Dangerous!
    steps:
      - uses: actions/checkout@v4
      - run: npm test
```

## Policies

```yaml
policies:
  - name: "gha-no-write-all-permissions"
    description: "Workflows must not use write-all permissions. Apply least privilege instead."
    severity: "error"
    rule: "!has(input.permissions) || input.permissions != 'write-all'"

  - name: "gha-require-job-timeout"
    description: "All jobs must explicitly declare 'timeout-minutes' to prevent hung runners."
    severity: "warning"
    rule: "input.jobs.values().all(job, has(job['timeout-minutes']))"
```

## Rottweiler Check

```bash
rottweiler check --input .github/workflows/ci.yaml --policy policy.yaml
```