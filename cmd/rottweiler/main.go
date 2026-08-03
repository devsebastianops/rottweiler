package main

import (
	"os"

	"github.com/devsebastianops/rottweiler/internal/cli"
	"github.com/devsebastianops/rottweiler/internal/logger"
)

func main() {
	err := cli.Run()
	if err != nil {

		logger.Error("rottweiler failed", "error", err.Error())
		os.Exit(1)
	}
}
