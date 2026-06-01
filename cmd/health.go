package cmd

import (
	"fmt"
	"os"

	"github.com/mrhujaifa/flamekit/internal/analyzer"
	"github.com/mrhujaifa/flamekit/internal/git"
	"github.com/mrhujaifa/flamekit/internal/ui"
	"github.com/spf13/cobra"
)

var healthPath string

var healthCmd = &cobra.Command{
	Use:   "health",
	Short: "Show overall project health score and trend",
	Long:  "Analyzes git history to calculate project health score, trend, and velocity risk.",
	Example: `  flamekit health
  flamekit health --path /other/project`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// validation
		if _, err := os.Stat(healthPath); os.IsNotExist(err) {
			return fmt.Errorf("path does not exist: %s", healthPath)
		}

		// loading
		fmt.Fprintf(os.Stderr, "  Analyzing repository...\n\n")

		// parse
		commits, err := git.ParseRepository(healthPath)
		if err != nil {
			return fmt.Errorf("failed to read git history: %w", err)
		}

		if len(commits) == 0 {
			return fmt.Errorf("no commits found in repository")
		}

		// calculate scores
		scores, err := analyzer.CalculateFlameScores(commits)
		if err != nil {
			return fmt.Errorf("failed to calculate scores: %w", err)
		}

		scores = analyzer.FilterCodeFiles(scores)

		// calculate health
		report := analyzer.CalculateHealth(commits, scores)

		// render
		ui.RenderHealth(report)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(healthCmd)
	healthCmd.Flags().StringVarP(
		&healthPath,
		"path", "p",
		".",
		"Path to git repository",
	)
}
