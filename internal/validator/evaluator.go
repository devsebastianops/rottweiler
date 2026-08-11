package validator

import (
	"fmt"
	"strings"

	"github.com/devsebastianops/x/logger"
	"github.com/google/cel-go/cel"
)

func evalCelCondition(expr string, env *cel.Env, inputData map[string]any, strict bool) (bool, error) {
	ast, iss := env.Compile(expr)
	if iss.Err() != nil {
		return false, fmt.Errorf("syntax error: %v", iss.Err())
	}
	program, err := env.Program(ast)
	if err != nil {
		return false, err
	}
	out, _, err := program.Eval(map[string]any{
		"input": inputData,
	})
	if err != nil {
		// Handle missing key errors based on strict mode
		if isMissingKeyErrorInCondition(err) {
			if strict {
				return false, err
			}
			// In lenient mode, log warning and return false (condition not met)
			logger.Warn("missing key in condition, treating as false", "expression", expr, "error", err.Error())
			return false, nil
		}
		return false, err
	}

	boolVal, ok := out.Value().(bool)
	if !ok {
		return false, fmt.Errorf("expression did not return a boolean value")
	}
	return boolVal, nil
}

func isMissingKeyErrorInCondition(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return strings.Contains(errMsg, "no such key") ||
		strings.Contains(errMsg, "no such field") ||
		strings.Contains(errMsg, "undefined reference")
}
