package cli

import (
	"github.com/devsebastianops/x/logger"
	"github.com/spf13/cobra"
)

var (
	Version   = "development"
	Commit    = "none"
	BuildTime = "2026-01-01T00:00:00Z"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of rottweiler",
	Long:  "All software has versions. This is rottweiler's.",
	Run: func(cmd *cobra.Command, args []string) {
		logger.Infof("rottweiler version %s, commit %s, built at %s", Version, Commit, BuildTime)
	},
}
