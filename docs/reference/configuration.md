# Configuration

The Rottweiler policy file is the core definition of your rules. It follows a structured format that allows you to define policies, rules, and conditions in a clear and organized manner.

## Top Level Structure

| Key | Type | Description |
| --- | --- | --- |
| `policies` | `array` | A list of policy definitions. Each policy can contain multiple rules. |

## Policy structure

| Key | Type | Description |
| --- | --- | --- |
| `name` | `string` | The name of the policy. |
| `description` | `string` | A brief description of the policy. |
| `rule` | `string` | The CEL expression that defines the rule. |
| `severity` | `string` | The severity level of the policy (e.g., `error`, `info`, `warning`, `fatal`). |

## Example Policy File

```yaml
policies:
  - name: "No Public S3 Buckets"
    description: "Ensure that S3 buckets are not publicly accessible."
    rule: "input.resource.type == 'aws_s3_bucket' && input.resource.public_access == true"
    severity: "error"

  - name: "EC2 Instance Type Check"
    description: "Ensure that EC2 instances are of a specific type."
    rule: "input.resource.type == 'aws_ec2_instance' && input.resource.instance_type in ['t2.micro', 't2.small']"
    severity: "warning"
```

