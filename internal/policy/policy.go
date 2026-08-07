package policy

type PolicyFile struct {
	Policies []Policy `json:"policies" yaml:"policies"`
}

type Policy struct {
	Id          string   `json:"id,omitempty" yaml:"id,omitempty"`
	Name        string   `json:"name" yaml:"name"`
	Description string   `json:"description" yaml:"description"`
	Rule        string   `json:"rule" yaml:"rule"`
	Rationale   string   `json:"rationale,omitempty" yaml:"rationale,omitempty"`
	Remediation string   `json:"remediation,omitempty" yaml:"remediation,omitempty"`
	References  []string `json:"references,omitempty" yaml:"references,omitempty"`
	Severity    string   `json:"severity" yaml:"severity"`
	Status      string   `json:"status" yaml:"status"`
}

func NewPolicy(
	id,
	name,
	description,
	rule,
	rationale,
	remediation string,
	references []string,
	severity,
	status string,
) (Policy, error) {

	if name == "" {
		return Policy{}, &InvalidPolicyError{Reason: "name cannot be empty"}
	}

	if len(name) < 3 {
		return Policy{}, &InvalidPolicyError{Reason: "name must be at least 3 characters long"}
	}

	if rule == "" {
		return Policy{}, &InvalidPolicyError{Reason: "rule cannot be empty"}
	}

	if severity == "" {
		severity = SeverityError // Default severity is "error"
	}

	if !isValidSeverity(severity) {
		return Policy{}, &InvalidSeverityError{Severity: severity}
	}

	if status == "" {
		status = PolicyStatusUnchecked // Default status is "unchecked"
	}

	if !isValidStatus(status) {
		return Policy{}, &InvalidStatusError{Status: status}
	}

	return Policy{
		Id:          id,
		Name:        name,
		Description: description,
		Rule:        rule,
		Rationale:   rationale,
		Remediation: remediation,
		References:  references,
		Severity:    severity,
		Status:      status,
	}, nil
}

type InvalidPolicyError struct {
	Reason string
}

func (e *InvalidPolicyError) Error() string {
	return "invalid policy: " + e.Reason
}

type PolicyFatalError struct {
	Policy Policy
}

func (e *PolicyFatalError) Error() string {
	return "fatal policy violation: " + e.Policy.Name + " - " + e.Policy.Description + " - Rule: " + e.Policy.Rule
}
