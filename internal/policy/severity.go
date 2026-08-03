package policy

import "slices"

const (
	SeverityFatal = "fatal" // Aborts the execution of the workflow and marks it as failed.
	SeverityError = "error" // Marks the workflow as failed but continues the execution of the workflow.
	SeverityWarn  = "warn"  // Marks the workflow as successful but logs a warning.
	SeverityInfo  = "info"  // Logs an informational message but does not affect the workflow's success or failure.
)

var severities = []string{
	SeverityFatal,
	SeverityError,
	SeverityWarn,
	SeverityInfo,
}

func isValidSeverity(severity string) bool {
	return slices.Contains(severities, severity)
}

type InvalidSeverityError struct {
	Severity string
}

func (e *InvalidSeverityError) Error() string {
	return "invalid severity: " + e.Severity
}
