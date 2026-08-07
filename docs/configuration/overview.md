# Overview

Rottweiler uses a YAML configuration file to define which policies are applied to which resources, also referred as a "policyfile".

This section outlines the basic structure of the policyfile and provides an overview of the available configuration options.

## Policyfile Structure

Here is a minimal example of a policyfile:

```yaml
policies:
- name: No latest
  description: Disallow the use of the "latest" tag in container images
  rule: input.image.tag != 'latest'
```

## Inside a Policy

Each policy consists of the following fields:

- `name`: A unique name for the policy.
- `description`: A brief description of what the policy does.
- `rule`: The rule that defines the policy. This is typically a boolean expression that evaluates to true or false based on the input data.
- `severity`: The severity level of the policy violation. This can be `fatal`, `error`, `warning`, or `info`. When there is no severity defined, the default is `error`.

### Enhanced Policy Definitions

Additionally, policies can include the following optional fields to provide more context and guidance:

- `rationale`: A brief explanation of why the policy exists and what it aims to prevent or enforce.
- `remediation`: Suggested steps or actions to take when the policy is violated, helping users understand how to resolve issues.
- `references`: A list of external links or references that provide additional information or context related to the policy. This can include documentation, best practices, or relevant articles.