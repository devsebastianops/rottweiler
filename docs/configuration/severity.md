# Severities

Rottweiler supports the following severity levels for rules:

- `fatal`: The highest severity level. A violation of a rule with this severity will cause the policy check to fail immediately and halt further processing. 
- `error`: A violation of a rule with this severity will cause the policy check to fail, but processing will continue for other rules.
- `warning`: A violation of a rule with this severity will not cause the policy check to fail, but it will be reported as a warning in the output.
- `info`: The lowest severity level. A violation of a rule with this severity will not cause the policy check to fail, and it will be reported as an informational message in the output.

When there is no severity defined for a rule, the default severity level is `error`.