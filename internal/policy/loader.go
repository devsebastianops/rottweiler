package policy

import (
	"os"

	"github.com/devsebastianops/rottweiler/internal/logger"
	"gopkg.in/yaml.v3"
)

func LoadPoliciesFromFile(filePath string) (PolicyFile, error) {
	if filePath == "" {
		return PolicyFile{}, &InvalidPolicyError{Reason: "policy file path cannot be empty"}
	}

	// File existance check
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return PolicyFile{}, &InvalidPolicyError{Reason: "policy file does not exist"}
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return PolicyFile{}, err
	}

	logger.Debugf("File content of %s: %s", filePath, string(fileContent))

	//
	var policyFile PolicyFile

	err = yaml.Unmarshal(fileContent, &policyFile)
	if err != nil {
		return PolicyFile{}, err
	}

	logger.Debugf("Loaded policies from file: %s", filePath)
	logger.Debugf("Policies: %+v", policyFile)

	return policyFile, nil
}
