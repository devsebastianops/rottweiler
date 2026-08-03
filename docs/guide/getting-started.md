# Getting Started

## JSON and YAML everywhere

In modern environments, infrastructure and configuration as code are ubiquitous. We manage massive amounts of YAML and JSON files—whether they’re Kubernetes manifests, Terraform state files, or custom application configurations.

And the main challenge from my experience is most of the time **human error.**

Developers can easily make mistakes, regardless of what caused them, when writing these files. This can lead to misconfigurations, security vulnerabilities, and downtime.

So we can no longer rely on syntax validation alone.
A standard linter only checks whether the file is formatted correctly. 

However, it does not prevent a developer from accidentally making an S3 bucket public, defining a Kubernetes deployment without resource limits. Or just forgetting conventional tags, they agreed on two months ago with the team.

## Possible Solution 1: Custom scripts

A common solution is to write custom scripts that check for specific conditions in the files.

However, this approach has several drawbacks, as always, when you write custom code:

It is **Time-consuming**: Once you write a script, you need to maintain it. And if the requirements change, you need to update the script accordingly. That is valuable time that could be spent on other tasks.

And oftentimes, you end up with a **messy collection of scripts** that are hard to maintain and understand. And if you have multiple teams, you will end up with multiple scripts that do the same thing but in different ways.


## Possible Solution 2: A policy engine

That's exactly the reason why "Policy as Code" solutions are on the rise. They allow you to define rules and best practices for your infrastructure and configuration files, ensuring that they meet your organization's standards before they are deployed.

But, however, existing enterprise solutions for Policy as Code are often cumbersome, require users to learn complex, proprietary languages, or entail a massive runtime overhead.

## The solution: Your personal policy watchdog

Rottweiler is a lightweight, declarative policy-as-code engine designed specifically to validate YAML and JSON documents at lightning speed without requiring complex infrastructure.

We developed Rottweiler with the following goals in mind:

- **Simplicity Instead of Overengineering.**
  We did not want to create a completely new language or DSL and stick to the tools, that developers are already familiar with.

- **Declarative and Flexible.**
  We wanted to create a solution that is easy to read and easy to maintain but flexible enough to cover a wide range of use cases.

- **CI/CD Ready.**
  We must be able to run our solution in **any** CI/CD pipeline, because we believe that is the only way to ensure that policies are enforced consistently across all environments.


And with Rottweiler, we achieved all of these goals. It is a declarative policy engine that validates **any** YAML or JSON document against policies you define using Google CEL and YAML.

## Quick Example


### 1. The data file ( JSON or YAML )

First of all we need structured data to validate. Rottweiler works with any JSON or YAML document, so you can use it to validate Kubernetes manifests, Terraform JSON, GitHub Actions workflows, or your own configuration files.

In our example, we will use a simple `deployment.yaml`.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  template:
    spec:
      containers:
      - name: nginx
        image: nginx:latest
```

### 2. The policy file 

Next, we need to define a policy that describes the rules we want to enforce. Policies are defined in YAML format and evaluated using Google CEL.

```yaml
policies:
-   name: "no latest"
    description: "We want to ensure that no container is using the 'latest' tag for its image."
    severity: "error"
    rule: "input.spec.template.spec.containers.all(c, c.image.endsWith(':latest') == false)"
```

### 3. Run Rottweiler

Now we can run Rottweiler to validate our `deployment.yaml` against the policy we defined in `policy.yaml`.

```bash
rottweiler check --input deployment.yaml --policy policy.yaml --format json
```

### Expected result

Because we passed the `--format json` flag, Rottweiler will output the result in JSON format. The output will look like this:

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
      "rule": "input.spec.template.spec.containers.all(c, c.image.endsWith(':latest') == false)"
    }
  ],
  "warnings": [],
  "infos": []
}
```

Please refer to the [Rottweiler CLI Reference](/reference/cli) for more information on the available commands and options or the [Examples](/examples/overview) section for more use cases and examples of how to use Rottweiler in your projects.