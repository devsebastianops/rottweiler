# Core concepts

To understand how Rottweiler works, it is important to understand the core concepts behind it. 

Rottweiler is a policy engine that evaluates input data against a set of policies defined in a policy file and designed to solve a pain: Ensuring that your rules and standards are being followed.

## 1. Declarative YAML policies

Rottweiler uses `YAML` to define policies.

A policyfile contains of a list of policies, each with a name, description, severity, and a rule. 

```yaml
policies:
-   name: "no latest"
    description: "We want to ensure that no container is using the 'latest' tag for its image."
    severity: "error"
    rule: "input.image.endsWith(':latest') == false"
```

## 2. Heavy lifting with CEL (Common Expression Language)

What sets Rottweiler apart from other policy engines is that it uses Google CEL (Common Expression Language) to evaluate the rules defined in the policy file. CEL is a powerful expression language that allows you to write complex expressions in a simple and concise way.

So you do not need to learn a new language to write policies, you can use CEL expressions to define the rules in your policy file.

## Afterword: Start playing with Rottweiler

Rottweiler is a powerful companion and you now know the core concepts behind it. Start playing with it and see how it can help you ensure that your rules and standards are being followed.

In case you want to learn more about Rottweiler, check out the [Configuration Overview](/configuration/overview) or the [Real world Examples](/examples/overview).
