package reporter

import (
	"fmt"

	"github.com/devsebastianops/rottweiler/internal/validator"
	"github.com/devsebastianops/x/logger"
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

	if validationResult.HasErrors() {

		reportFindings(validationResult.Errors, "Error")
	}

	if validationResult.HasWarnings() {

		reportFindings(validationResult.Warnings, "Warning")
	}

	if validationResult.HasInfos() {

		reportFindings(validationResult.Infos, "Info")
	}
}

func reportFindings(findings []validator.Finding, title string) {
	if len(findings) == 0 {
		return
	}

	for _, finding := range findings {
		logger.Info(title+" in policy", "policy", finding.Policy, "description", finding.Description, "severity", finding.Severity, "rule", finding.Rule)

		if finding.Id != "" {
			logger.Info("Policy ID: " + finding.Id)
		}
		if finding.Rationale != "" {
			logger.Info("Rationale: " + finding.Rationale)
		}
		if finding.Remediation != "" {
			logger.Info("Remediation: " + finding.Remediation)
		}
		if len(finding.References) > 0 {
			for i, ref := range finding.References {
				logger.Infof("Reference %d: %s", i+1, ref)
			}
		}
	}
}
