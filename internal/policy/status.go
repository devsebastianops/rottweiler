package policy

import "slices"

const (
	PolicyStatusUnchecked = "unchecked" // The policy was not checked
	PolicyStatusActive    = "checked"   // The policy was checked and is ok
	PolicyStatusFailed    = "failed"    // The policy was checked and failed
)

var statuses = []string{
	PolicyStatusUnchecked,
	PolicyStatusActive,
	PolicyStatusFailed,
}

func isValidStatus(status string) bool {
	return slices.Contains(statuses, status)
}

type InvalidStatusError struct {
	Status string
}

func (e *InvalidStatusError) Error() string {
	return "invalid status: " + e.Status
}
