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
| `rationale` | `string` | (Optional) The rationale behind the policy. |
| `remediation` | `string` | (Optional) Suggested remediation steps for the policy. |
| `references` | `array` | (Optional) A list of references or links related to the policy. |

## Example Policy File

```yaml
policies:
  - name: "No Public S3 Buckets"
    description: "Ensure that S3 buckets are not publicly accessible."
    rule: "input.resource.type == 'aws_s3_bucket' && input.resource.public_access == true"
    severity: "error"
    rationale: "Public S3 buckets can lead to data leaks and security vulnerabilities."
    remediation: "Restrict public access to the S3 bucket by updating its ACL and bucket policy."
    references:
      - "https://docs.aws.amazon.com/AmazonS3/latest/userguide/access-control-block-public-access.html"

  - name: "EC2 Instance Type Check"
    description: "Ensure that EC2 instances are of a specific type."
    rule: "input.resource.type == 'aws_ec2_instance' && input.resource.instance_type in ['t2.micro', 't2.small']"
    severity: "warning"
```

