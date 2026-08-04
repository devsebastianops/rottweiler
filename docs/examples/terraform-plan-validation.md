# Terraform Plan Validation

When working with Cloud Infrastructure as Code (IaC), policy validation should happen at the Terraform Plan stage before any real resources are provisioned.

## Scenario

- **Cost control**: Disallow heavy, expensive EC2 instance types in non-production environments.
- **Organizational Policies**: Ensure S3 buckets are tagged with an Owner and Environment.

## Input

The input for this example is a Terraform plan file in JSON format. You can generate this file using the following command:

```bash
terraform plan -out=planfile
terraform show -json planfile > tfplan.json
```

For this example, we will asume something like this is part of your terraform plan:
```json
{
  "resource_changes": [
    {
      "type": "aws_instance",
      "name": "dev_server",
      "change": {
        "after": {
          "instance_type": "p3.16xlarge",
          "tags": {
            "Environment": "dev"
          }
        }
      }
    }
  ]
}
```

## Policies

```yaml
policies:
  - name: "tf-disallow-expensive-ec2"
    description: "Prevent creation of high-cost GPU/large EC2 instances in development."
    severity: "error"
    rule: "input.resource_changes.all(rc, rc.type != 'aws_instance' || !rc.change.after.instance_type.startsWith('p3.'))"

  - name: "tf-s3-require-tags"
    description: "S3 buckets must be tagged with an Owner tag for cost allocation."
    severity: "warning"
    rule: "input.resource_changes.all(rc, rc.type != 'aws_s3_bucket' || has(rc.change.after.tags.Owner))"
```

## Rottweiler Check

```bash
rottweiler check --input tfplan.json --policy policies.yaml
```