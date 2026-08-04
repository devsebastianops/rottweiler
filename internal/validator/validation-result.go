package validator

import (
	"encoding/json"

	"github.com/devsebastianops/rottweiler/internal/policy"
)

type ValidationResult struct {
	Summary  Summary   `json:"summary" yaml:"summary"`
	Errors   []Finding `json:"errors" yaml:"errors"`
	Warnings []Finding `json:"warnings" yaml:"warnings"`
	Infos    []Finding `json:"infos" yaml:"infos"`
}

type Summary struct {
	CheckedPolicies int `json:"checkedPolicies" yaml:"checkedPolicies"`
	Warnings        int `json:"warnings" yaml:"warnings"`
	Errors          int `json:"errors" yaml:"errors"`
	Infos           int `json:"infos" yaml:"infos"`
}

type Finding struct {
	Policy      string `json:"policy" yaml:"policy"`
	Description string `json:"description" yaml:"description"`
	Severity    string `json:"severity" yaml:"severity"`
	Rule        string `json:"rule" yaml:"rule"`
}

func NewFinding(policy policy.Policy) Finding {

	return Finding{
		Policy:      policy.Name,
		Description: policy.Description,
		Severity:    policy.Severity,
		Rule:        policy.Rule,
	}
}

func emptyValidationResult() ValidationResult {
	return ValidationResult{
		Summary: Summary{
			CheckedPolicies: 0,
			Warnings:        0,
			Errors:          0,
			Infos:           0,
		},
		Errors:   []Finding{},
		Warnings: []Finding{},
		Infos:    []Finding{},
	}
}

func (vr *ValidationResult) AddFinding(finding Finding) {
	vr.Summary.CheckedPolicies++
	switch finding.Severity {
	case policy.SeverityFatal:
		vr.Errors = append(vr.Errors, finding)
		vr.Summary.Errors++
	case policy.SeverityError:
		vr.Errors = append(vr.Errors, finding)
		vr.Summary.Errors++
	case policy.SeverityWarn:
		vr.Warnings = append(vr.Warnings, finding)
		vr.Summary.Warnings++
	case policy.SeverityInfo:
		vr.Infos = append(vr.Infos, finding)
		vr.Summary.Infos++
	}
}

func (vr *ValidationResult) AddCheckedPolicy() {
	vr.Summary.CheckedPolicies++
}

func (vr *ValidationResult) HasFindings() bool {
	return vr.HasErrors() || vr.HasWarnings() || vr.HasInfos()
}

func (vr *ValidationResult) HasErrors() bool {
	return vr.Summary.Errors > 0
}

func (vr *ValidationResult) HasWarnings() bool {
	return vr.Summary.Warnings > 0
}

func (vr *ValidationResult) HasInfos() bool {
	return vr.Summary.Infos > 0
}

func (vr *ValidationResult) ToJSON() ([]byte, error) {
	return json.Marshal(vr)
}
