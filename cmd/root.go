package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "scorch",
	Short: "Find where your codebase burns",
	Long:  "Scorch analyzes your git history to find the most dangerous files in your codebase.",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
