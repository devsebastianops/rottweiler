package cli

import (
	"github.com/devsebastianops/x/logger"
	"github.com/spf13/cobra"
)

type PersistentFlags struct {
	Verbose bool
	Format  string
	Silent  bool
}

var persistentFlags = PersistentFlags{}

var RootCmd = &cobra.Command{
	Use:   "rottweiler",
	Short: "rottweiler is a declarative policy engine for any YAML or JSON document.",
	Long:  "rottweiler is a declarative policy engine that validates any YAML or JSON document against policies you define.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		isSilentAndVerbose := persistentFlags.Silent && persistentFlags.Verbose
		if isSilentAndVerbose {
			logger.Warn("You used both --silent and --verbose flags. Silent mode will take precedence.")
		}
		logger.SetUp(persistentFlags.Verbose, persistentFlags.Format, persistentFlags.Silent)
	},
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.PersistentFlags().BoolVar(&persistentFlags.Verbose, "verbose", false, "Enable verbose output")
	RootCmd.PersistentFlags().StringVar(&persistentFlags.Format, "format", "pretty", "Set log format ('pretty', 'text' or 'json')")
	RootCmd.PersistentFlags().BoolVar(&persistentFlags.Silent, "silent", false, "Enable silent mode")

	RootCmd.AddCommand(checkCmd)
	RootCmd.Aliases = []string{"rw"}
}

func Run() error {
	return RootCmd.Execute()
}
