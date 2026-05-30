package cmd

import (
	"fmt"
	"os"

	"github.com/mrhujaifa/scorch/internal/analyzer"
	"github.com/mrhujaifa/scorch/internal/git"
	"github.com/mrhujaifa/scorch/internal/ui"
	"github.com/spf13/cobra"
)

var suggestLimit int

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Get prioritized refactoring suggestions",
	Long:  "Analyzes git history and suggests which files to refactor first for maximum impact.",
	Example: `  scorch suggest
  scorch suggest --limit 10
  scorch suggest --limit 3`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// validation
		if suggestLimit < 1 {
			return fmt.Errorf("limit must be greater than 0")
		}

		if suggestLimit > 50 {
			return fmt.Errorf("limit cannot exceed 50")
		}

		// parse repository
		commits, err := git.ParseRepository(".")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		// calculate scores
		scores, err := analyzer.CalculateFlameScores(commits)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}

		// filter and sort
		scores = analyzer.FilterCodeFiles(scores)
		scores = analyzer.SortByFlame(scores)

		// apply limit
		if len(scores) > suggestLimit {
			scores = scores[:suggestLimit]
		}

		// render
		ui.RenderSuggestions(scores)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(suggestCmd)
	suggestCmd.Flags().IntVarP(
		&suggestLimit,
		"limit", "l",
		5,
		"Number of suggestions to show (1-50)",
	)
}
