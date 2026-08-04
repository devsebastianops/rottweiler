# Reporting

Rottweiler supports multiple output formats for reporting policy violations.

## Pretty

By default, Rottweiler outputs the results in a human-readable format. This format is suitable for quick reviews and debugging.

Explicitly you can specify the `pretty` format using the `--format` flag:

```bash
rottweiler check --input <input_file> --policy <policy_file> --format pretty
```

Example output:

```text
INFO Validation Summary:
INFO Checked Policies: 1
INFO Errors: 0
INFO Warnings: 0
INFO Infos: 0
```

## Text

The `text` format provides a simple, human and machine readable output. It might be useful for logging and further processing of the results.

You can specify the `text` format using the `--format` flag:

```bash
rottweiler check --input <input_file> --policy <policy_file> --format text
```

Example output:

```text
time=2026-08-04T08:00:44.998+02:00 level=INFO msg="Validation Summary:"
time=2026-08-04T08:00:45.030+02:00 level=INFO msg="Checked Policies: 1"
time=2026-08-04T08:00:45.030+02:00 level=INFO msg="Errors: 0"
time=2026-08-04T08:00:45.030+02:00 level=INFO msg="Warnings: 0"
time=2026-08-04T08:00:45.030+02:00 level=INFO msg="Infos: 0"
```

## JSON

The `json` format provides a structured output that can be easily parsed and processed by other tools or scripts. This format is suitable for integration with CI/CD pipelines or other automated systems.

You can specify the `json` format using the `--format` flag:

```bash
rottweiler check --input <input_file> --policy <policy_file> --format json
```

Example output:

```json
{
  "summary": {
    "checked_policies": 1,
    "errors": 0,
    "warnings": 0,
    "infos": 0
  },
  "errors": [],
  "warnings": [],
  "infos": []
}
```