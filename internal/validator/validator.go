package validator

import (
	"github.com/devsebastianops/rottweiler/internal/policy"
	"github.com/google/cel-go/cel"
	"github.com/google/cel-go/ext"
)

func Validate(policyFile policy.PolicyFile, inputData map[string]interface{}, strict bool) (ValidationResult, error) {

	// Prepare CEL env
	env, _ := cel.NewEnv(
		cel.Variable("input", cel.MapType(cel.StringType, cel.AnyType)),
		cel.Macros(cel.StandardMacros...),
		cel.OptionalTypes(),
		ext.Encoders(),
		ext.Lists(),
		ext.Strings(),
		ext.Sets(),
		ext.Regex(),
	)

	// New empty validation result
	validationResult := emptyValidationResult()

	// Check policies
	for _, p := range policyFile.Policies {
		// Evaluate the policy rule using CEL
		result, err := evalCelCondition(p.Rule, env, inputData, strict)

		if err != nil {
			return validationResult, err
		}

		if !result {
			finding := NewFinding(p)
			validationResult.AddFinding(finding)
			if p.Severity == policy.SeverityFatal {
				return validationResult, &policy.PolicyFatalError{Policy: p}
			}
		} else {
			validationResult.AddCheckedPolicy()
		}
	}

	return validationResult, nil
}
