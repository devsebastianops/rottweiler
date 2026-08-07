package reporter

import (
	"fmt"

	"github.com/devsebastianops/rottweiler/internal/logger"
	"github.com/devsebastianops/rottweiler/internal/validator"
)

func Report(validationResult validator.ValidationResult, logFormat string, silent bool) {

	if silent {
		return
	}

	logger.Debug("Reporting validation results...")
	logger.Debugf("Validation Result: %+v", validationResult)
	logger.Debugf("Log Format: %s", logFormat)

	if logFormat == logger.LOG_FORMAT_JSON {
		reportJson(validationResult)
	} else {
		reportTextOrPretty(validationResult)
	}

}

func reportJson(validationResult validator.ValidationResult) {
	jsonData, err := validationResult.ToJSON()
	if err != nil {
		fmt.Printf("Error converting validation result to JSON: %v\n", err)
		return
	}
	fmt.Printf("%s\n", jsonData)
}

func reportTextOrPretty(validationResult validator.ValidationResult) {
	logger.Infof("Validation Summary:")
	logger.Infof("Checked Policies: %d", validationResult.Summary.CheckedPolicies)
	logger.Infof("Errors: %d", validationResult.Summary.Errors)
	logger.Infof("Warnings: %d", validationResult.Summary.Warnings)
	logger.Infof("Infos: %d", validationResult.Summary.Infos)

	if validationResult.HasFindings() {
		logger.Infof("Detailed Findings:")
	}

	if validationResult.HasErrors() {

		reportFindings(validationResult.Errors, "Errors")
	}

	if validationResult.HasWarnings() {

		reportFindings(validationResult.Warnings, "Warnings")
	}

	if validationResult.HasInfos() {

		reportFindings(validationResult.Infos, "Infos")
	}
}

func reportFindings(findings []validator.Finding, title string) {
	if len(findings) == 0 {
		return
	}

	logger.Infof("%s:", title)
	for _, finding := range findings {
		logger.Infof("- Policy: %s\n  Description: %s\n  Severity: %s\n  Rule: %s\n", finding.Policy, finding.Description, finding.Severity, finding.Rule)
		if finding.Rationale != "" {
			logger.Infof("  Rationale: %s\n", finding.Rationale)
		}
		if finding.Remediation != "" {
			logger.Infof("  Remediation: %s\n", finding.Remediation)
		}
		if len(finding.References) > 0 {
			logger.Infof("  References:")
			for _, ref := range finding.References {
				logger.Infof("    - %s", ref)
			}
		}
	}
}
