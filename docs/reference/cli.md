# Rottweiler CLI Reference

## `rottweiler` Command

The global `rottweiler` command is the entry point for all Rottweiler CLI operations. 

### Global Options

| Option | Type | Description |
| --- | --- | --- |
| `--verbose` | boolean | Enable verbose output for debugging purposes. |
| `--format` | string | Specify the output format (e.g., `json`, `text`, `pretty`). |
| `--silent` | boolean | Suppress all output except for errors and warnings. |

## `rottweiler check`

`rottweiler check` is used to validate policies against a given set of rules and configurations.

### Usage

```bash
rottweiler check [options]
```

### Options

| Option | Type | Description |
| --- | --- | --- |
| `--policy` | string | Path to the policy file to be checked. |
| `--input` | string | Path to the input data file for policy evaluation. |
| `--strict` | boolean | Enables strict mode (halts on missing keys or syntax errors). |

### Example

A typical usage of the `rottweiler check` command might look like this:

```bash
rottweiler check --policy ./policies/my_policy.yaml --input ./data/input.json --format json
```
