package e2e

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/devsebastianops/rottweiler/internal/logger"
	"github.com/devsebastianops/rottweiler/internal/parser"
	"github.com/devsebastianops/rottweiler/internal/policy"
	"github.com/devsebastianops/rottweiler/internal/validator"
)

func TestE2EExamples(t *testing.T) {
	exampleDir := "../../example"

	entries, err := os.ReadDir(exampleDir)
	if err != nil {
		t.Fatalf("failed to read example directory: %v", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		exampleName := entry.Name()
		examplePath := filepath.Join(exampleDir, exampleName)

		t.Run(exampleName, func(t *testing.T) {
			runE2ETest(t, examplePath)
		})
	}
}

// runE2ETest runs a single example through the full pipeline
func runE2ETest(t *testing.T, examplePath string) {
	exampleName := filepath.Base(examplePath)

	// Find input file (either .json or .yaml)
	var inputFile string
	for _, ext := range []string{".json", ".yaml", ".yml"} {
		candidate := filepath.Join(examplePath, "input"+ext)
		if _, err := os.Stat(candidate); err == nil {
			inputFile = candidate
			break
		}
	}
	if inputFile == "" {
		t.Fatalf("no input file found in %s", examplePath)
	}

	policyFileLocation := filepath.Join(examplePath, "policy.yaml")
	if _, err := os.Stat(policyFileLocation); err != nil {
		t.Fatalf("policy file not found: %s", policyFileLocation)
	}

	expectedOutputFile := filepath.Join(examplePath, "expected_output.json")
	if _, err := os.Stat(expectedOutputFile); err != nil {
		t.Fatalf("expected output file not found: %s", expectedOutputFile)
	}

	// load policy file
	policyFile, err := policy.LoadPoliciesFromFile(policyFileLocation)
	if err != nil {
		t.Fatalf("failed to load policy file: %v", err)
	}

	logger.Debug("Found policies", "policies", policyFile)

	// load input data
	inputData, err := parser.ParseInput(inputFile)

	if err != nil {
		t.Fatalf("failed to parse input file: %v", err)
	}

	result, err := validator.Validate(
		policyFile,
		inputData,
		true,
	)

	if err != nil {
		t.Fatalf("validation failed: %v", err)
	}

	var expected validator.ValidationResult
	expectedData, err := os.ReadFile(expectedOutputFile)
	if err != nil {
		t.Fatalf("failed to read expected output file: %v", err)
	}
	if err := json.Unmarshal(expectedData, &expected); err != nil {
		t.Fatalf("failed to parse expected output: %v", err)
	}

	// 5. Compare
	if result.Summary.CheckedPolicies != expected.Summary.CheckedPolicies {
		t.Errorf("CheckedPolicies mismatch: expected %d, got %d", expected.Summary.CheckedPolicies, result.Summary.CheckedPolicies)
	}

	if result.Summary.Errors != expected.Summary.Errors {
		t.Errorf("Errors mismatch: expected %d, got %d", expected.Summary.Errors, result.Summary.Errors)
	}

	if result.Summary.Warnings != expected.Summary.Warnings {
		t.Errorf("Warnings mismatch: expected %d, got %d", expected.Summary.Warnings, result.Summary.Warnings)
	}

	if result.Summary.Infos != expected.Summary.Infos {
		t.Errorf("Infos mismatch: expected %d, got %d", expected.Summary.Infos, result.Summary.Infos)
	}

	if len(result.Errors) != len(expected.Errors) {
		t.Errorf("Errors length mismatch: expected %d, got %d", len(expected.Errors), len(result.Errors))
	}

	if len(result.Warnings) != len(expected.Warnings) {
		t.Errorf("Warnings length mismatch: expected %d, got %d", len(expected.Warnings), len(result.Warnings))
	}

	if len(result.Infos) != len(expected.Infos) {
		t.Errorf("Infos length mismatch: expected %d, got %d", len(expected.Infos), len(result.Infos))
	}

	if !compareFindings(result.Errors, expected.Errors) {
		t.Errorf("Errors mismatch: expected %+v, got %+v", expected.Errors, result.Errors)
	}

	if !compareFindings(result.Warnings, expected.Warnings) {
		t.Errorf("Warnings mismatch: expected %+v, got %+v", expected.Warnings, result.Warnings)
	}

	if !compareFindings(result.Infos, expected.Infos) {
		t.Errorf("Infos mismatch: expected %+v, got %+v", expected.Infos, result.Infos)
	}

	t.Logf("✓ %s passed", exampleName)

}

func compareFindings(a, b []validator.Finding) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
