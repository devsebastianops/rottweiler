package cli

import (
	"os"

	"github.com/devsebastianops/rottweiler/internal/logger"
	"github.com/devsebastianops/rottweiler/internal/parser"
	"github.com/devsebastianops/rottweiler/internal/policy"
	"github.com/devsebastianops/rottweiler/internal/reporter"
	"github.com/devsebastianops/rottweiler/internal/validator"
	"github.com/spf13/cobra"
)

type CheckOptions struct {
	PolicyFile string
	InputFile  string
	Strict     bool
}

var checkOptions = &CheckOptions{}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check a YAML or JSON document against policies",
	Long:  "Check a YAML or JSON document against policies you define. The command will validate the document and report any policy violations.",
	RunE: func(cmd *cobra.Command, args []string) error {
		return check()
	},
}

func init() {
	checkCmd.Flags().StringVarP(&checkOptions.PolicyFile, "policy", "p", "", "Path to the policy file")
	checkCmd.Flags().StringVarP(&checkOptions.InputFile, "input", "i", "", "Path to the data file (YAML or JSON)")
	checkCmd.Flags().BoolVarP(&checkOptions.Strict, "strict", "s", false, "Enable strict mode (fail on missing keys in conditions)")
	checkCmd.MarkFlagRequired("policy")
	checkCmd.MarkFlagRequired("input")
}

func check() error {
	logger.Debug("Check command executed with policy file: %s and input file: %s", checkOptions.PolicyFile, checkOptions.InputFile)

	// load policy file
	policyFile, err := policy.LoadPoliciesFromFile(checkOptions.PolicyFile)
	if err != nil {
		return err
	}

	logger.Debug("Found policies", "policies", policyFile)

	// load input data
	inputData, err := parser.ParseInput(checkOptions.InputFile)

	if err != nil {
		return err
	}

	validationResult, err := validator.Validate(
		policyFile,
		inputData,
		checkOptions.Strict,
	)

	if err != nil {
		return err
	}

	reporter.Report(
		validationResult,
		persistentFlags.Format,
		persistentFlags.Silent,
	)

	if validationResult.HasErrors() {
		os.Exit(1)
		return nil
	}

	return nil
}
