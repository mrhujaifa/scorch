package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "flamekit",
	Short:   "Find where your codebase burns",
	Long:    "Flamekit analyzes your git history to find the most dangerous files in your codebase.",
	Version: "0.3.0",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
