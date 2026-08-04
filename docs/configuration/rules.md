# Rules

Rottweiler uses rules to define the conditions that must be met for a policy to pass or fail.

Rules are typically expressed as boolean expressions that evaluate to true or false based on the input data. When a rule evaluates to `true`, the policy is considered to have passed; when it evaluates to `false`, the policy is considered to have failed.

## Examples

### Basic boolean expression:

```cel
input.image.tag != 'latest'
```

```yaml
image:
  tag: 'v1.0.0' # Will pass
```

```yaml
image: 
  tag: 'latest' # Will fail because it is equal to 'latest'
```

### CEL function usage:

```cel
input.image.tag.startsWith('v')
```

```yaml
image:
  tag: 'v1.0.0' # Will pass
```

```yaml
image: 
  tag: '1.0.0' # Will fail because it does not start with 'v'
```

### Complex expression with logical operators:

```cel
input.image.tag != 'latest' && input.image.tag.startsWith('v')
```

```yaml
image:
  tag: 'v1.0.0' # Will pass
```

```yaml
image: 
  tag: 'latest' # Will fail because first condition is false
```