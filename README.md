<div align="center">
  <img src="./docs/public/assets/rottweiler_mascott.png" alt="Rottweiler Mascot" width="180" />
  <h1>Rottweiler</h1>
  <p><strong>Your personal policy watchdog.</strong></p>
  <p>
    <a href="https://github.com/devsebastianops/rottweiler/actions"><img src="https://img.shields.io/github/actions/workflow/status/devsebastianops/rottweiler/on-pull-request-ci.yaml?branch=main&style=flat-square" alt="CI Status"></a>
    <a href="https://github.com/devsebastianops/rottweiler/releases"><img src="https://img.shields.io/github/v/release/devsebastianops/rottweiler?style=flat-square" alt="Latest Release"></a>
    <a href="https://github.com/devsebastianops/rottweiler/blob/main/LICENSE"><img src="https://img.shields.io/github/license/devsebastianops/rottweiler?style=flat-square" alt="License"></a>
  </p>
  <h3>
    <a href="https://rottweiler.dev">Documentation</a>
    <span> | </span>
    <a href="https://rottweiler.dev/guide/quick-start">Quick Start</a>
    <span> | </span>
    <a href="https://rottweiler.dev/examples/overview">Real-World Examples</a>
  </h3>
</div>

---

**Rottweiler** is a declarative policy engine that validates _any_ `YAML` or `JSON` document against policies you define.

Unlike policy engines that are tightly coupled to a specific ecosystem, Rottweiler works with **any structured data**. Whether you're validating Kubernetes manifests, Terraform JSON, GitHub Actions workflows, or your own configuration files, the workflow stays exactly the same.

Policies are defined in `YAML` format and evaluated using Google CEL, combining a familiar configuration format with a powerful rule language - **without writing custom code or learning a new DSL**. 

The result is a policy engine that is easy to read, easy to maintain, and simple to integrate into existing workflows.

---

## Why Rottweiler? 🐾

- **No proprietary policy language**: Write policies in YAML and express conditions with Google CEL.

- **Works with any structured data**: Terraform, Kubernetes, Docker Compose, Helm, GitHub Actions, APIs—or your own JSON.

- **Integrates anywhere**: Run it locally, in CI/CD pipelines, or embed it into your own tooling.

- **Fast by Design**: CEL expressions evaluate in milliseconds.

- **Developer friendly**: Readable policies, useful error messages, minimal boilerplate.


### Use Cases

- Enforce Kubernetes conventions
- Validate Terraform JSON
- Check GitHub Actions workflows
- Validate application configuration
- Enforce company standards in CI/CD
- ...and much more! 


---

## Quick Look: How it Works 

### 1. Simple Data ( JSON or YAML )

Any structured data can be validated by Rottweiler.

```json
{
  "service": {
    "name": "microservice-orders",
  },
  "app": {
    "image": "ghcr.io/examplecom/ms-orders",
    "version": "v1.0.0",
    "port": 8080
  },
  "database": {
    "type": "postgres",
    "version": "12",
    "size": "20GB"
  }
}
```

### 2. The policy file ( YAML)

The policy file defines what is allowed and what is not. It uses Google CEL expressions to define the rules.

```yaml
policies:
  - name: "Name must obey naming conventions"
    description: "We require that service names follow the naming convention: ms-<service-name>."
    rule: "service.name.matches('^ms-[a-z0-9-]+$')"
    severity: "error"

  - name: "Only allow valid, non-blocked, port numbers"
    description: "Valid, non-blocked port numbers are 1024-65535."
    rule: "app.port >= 1024 && app.port <= 65535"
    severity: "error"

  - name: "Expect size end with Gi"
    description: "We expect that database size ends with Gi."
    rule: "database.size.endsWith('Gi')"
    severity: "warning"
```

### 3. Run Rottweiler

```bash
rottweiler check --input input.json --policy policy.yaml --format json
```

#### Example Result

```json
{
  "summary": {
    "checkedPolicies": 3,
    "warnings": 1,
    "errors": 1
  },
  "warnings": [
    {
      "policy": "Expect size end with Gi",
      "description": "We expect that database size ends with Gi.",
      "severity": "warning",
      "rule": "database.size.endsWith('Gi')",
    }
  ],
  "errors": [
    {
      "policy": "Name must obey naming conventions",
      "description": "We require that service names follow the naming convention: ms-<service-name>.",
      "severity": "error",
      "rule": "service.name.matches('^ms-[a-z0-9-]+$')",
    }
  ]
}
```

---

## Getting Started

The fastest way to install and try out Rottweiler locally:

```bash
curl -sSL https://raw.githubusercontent.com/devsebastianops/rottweiler/main/install.sh | bash
rottweiler --help
```

For more installation methods and the full documentation check out [rottweiler.dev](https://rottweiler.dev).

---

## Examples 

All configurations in the `/examples` directory are actively tested in our E2E pipeline on every commit.

You can use these to understand how Rottweiler works, or as a starting point for your own blueprints.

> [!TIP]
> Have a look into our real world examples at [rottweiler.dev/examples/overview](https://rottweiler.dev/examples/overview).

---

## Community, Feedback & Support

We ❤️ contributions! Whether you found a bug, want to propose a new feature, or improve the codebase:

*   **Report Bugs:** Found an issue? [Open an Issue on GitHub](https://github.com/devsebastianops/rottweiler/issues).
*   **Contribute:** Check out our [Contribution Guidelines](https://github.com/devsebastianops/rottweiler/blob/main/CONTRIBUTING.md) to see how to get started with the codebase and documentation.
*   **Security:** If you discovered a security vulnerability, please refer to our [Security Policy](https://github.com/devsebastianops/rottweiler/blob/main/SECURITY.md) for responsible disclosure instructions.

---

## License

This project is licensed under the **MIT License** - see the [LICENSE](https://github.com/devsebastianops/rottweiler/blob/main/LICENSE) file for details.


---

<p align="center">
  Made with ❤️ by <a href="https://github.com/devsebastianops">devsebastianops</a>. We bite for good configs! 🐾
</p>
