# Package Governance ( NPM )

NPM uses `package.json` files to define the dependencies and scripts for a Node.js project. This can be used as input for Rottweiler to enforce certain rules on the package.json file.

## Scenario

- **Private packages only**: Ensures that the package is private and not published to the public NPM registry. This policy checks that the `private` field is set to `true`.
- **No wildcard dependencies**: Ensures that dependencies do not use wildcard versions. This policy checks that the version strings for dependencies do not contain `*`.

## Input

```json
{
  "name": "@my-company/internal-core",
  "version": "2.4.0",
  "private": false,
  "dependencies": {
    "express": "*",
    "lodash": "^4.17.21"
  }
}
```

## Policies

```yaml
policies:
  - name: "npm-enforce-private-flag"
    description: "Internal packages must have 'private: true' to prevent npm leakages."
    severity: "error"
    rule: "has(input.private) && input.private == true"

  - name: "npm-no-star-dependencies"
    description: "Dependencies cannot use wildcard versions."
    severity: "error"
    rule: "!has(input.dependencies) || input.dependencies.values().all(ver, ver != '*')"
```

## Rottweiler Check

```bash
rottweiler check --input package.json --policy policies.yaml
```