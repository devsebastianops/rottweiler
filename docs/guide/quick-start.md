# Quick start

This guide will help you get started with Rottweiler quickly.

To make you understand the basics of Rottweiler, we will go through a simple example of validating a simple JSON document against a policy. 

For more complex scenarios, please refer to the [Real-World Examples](/examples/overview).

## 1. Install Rottweiler

::: code-group
```bash [Installer]
curl -sSL https://raw.githubusercontent.com/devsebastianops/rottweiler/main/install.sh | bash
```
```bash [Go]
go install github.com/devsebastianops/rottweiler@latest
```
```bash [From Source]
git clone https://github.com/devsebastianops/rottweiler.git
cd rottweiler
go build -o rtw ./cmd/rtw/main.go
```
:::


::: tip
We recommend to use the installer script, as it will download the latest release and put it in your path.
:::

## 2. Create your input data

::: code-group
```json [JSON]
{
    "name": "my-microservice",
    "version": "1.0.0",
    "image": "myregistry/myimage:latest",
}
```

```yaml [YAML]
name: my-microservice
version: 1.0.0
image: myregistry/myimage:latest
```
:::

## 3. Create your policy

```yaml
policies:
-   name: "no latest"
    description: "We want to ensure that no container is using the 'latest' tag for its image."
    severity: "error"
    rule: "input.image.endsWith(':latest') == false"
```

## 4. Run Rottweiler

```bash
rottweiler check --input input.json --policy policy.yaml --format json
```

### The expected result

```json
{
  "summary": {
    "checkedPolicies": 1,
    "warnings": 0,
    "errors": 1,
    "infos": 0
  }, 
  "errors": [
    {
      "policy": "no latest",
      "description": "We want to ensure that no container is using the 'latest' tag for its image.",
      "severity": "error",
      "rule": "input.image.endsWith(':latest') == false"
    }
  ],
  "warnings": [],
  "infos": []
}
```

