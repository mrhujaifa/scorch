package cmd

import (
	"fmt"
	"os"

	"github.com/mrhujaifa/flamekit/internal/analyzer"
	"github.com/mrhujaifa/flamekit/internal/git"
	"github.com/mrhujaifa/flamekit/internal/ui"
	"github.com/mrhujaifa/flamekit/pkg/models"
	"github.com/spf13/cobra"
)

var (
	suggestLimit   int
	suggestPath    string
	suggestShowAll bool
)

var suggestCmd = &cobra.Command{
	Use:   "suggest",
	Short: "Get prioritized refactoring suggestions",
	Long:  "Analyzes git history and suggests which files to refactor first for maximum impact.",
	Example: `  flamekit suggest
  flamekit suggest --limit 10
  flamekit suggest --path /other/project
  flamekit suggest --all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// validation: limit range
		if suggestLimit < 1 {
			return fmt.Errorf("--limit must be greater than 0, got %d", suggestLimit)
		}
		if suggestLimit > 100 {
			return fmt.Errorf("--limit cannot exceed 100, got %d", suggestLimit)
		}

		// validation: path exists
		if _, err := os.Stat(suggestPath); os.IsNotExist(err) {
			return fmt.Errorf("path does not exist: %s", suggestPath)
		}

		// loading indicator
		fmt.Fprintf(os.Stderr, "  Analyzing repository...\n\n")

		// parse repository
		commits, err := git.ParseRepository(suggestPath)
		if err != nil {
			return fmt.Errorf("failed to read git history: %w", err)
		}

		// validation: no commits
		if len(commits) == 0 {
			return fmt.Errorf("no commits found in repository: %s", suggestPath)
		}

		// calculate scores
		scores, err := analyzer.CalculateFlameScores(commits)
		if err != nil {
			return fmt.Errorf("failed to calculate flame scores: %w", err)
		}

		// filter code files only
		scores = analyzer.FilterCodeFiles(scores)
		scores = analyzer.SortByFlame(scores)

		// filter risky files only (flame > 0)
		scores = filterRiskyFiles(scores)

		// validation: no risky files
		if len(scores) == 0 {
			fmt.Println()
			fmt.Println("  No risky files found.")
			fmt.Println("  Codebase looks healthy! Keep up the good work.")
			fmt.Println()
			return nil
		}

		// apply limit unless --all flag
		if !suggestShowAll && len(scores) > suggestLimit {
			scores = scores[:suggestLimit]
		}

		// render
		ui.RenderSuggestions(scores)

		return nil
	},
}

// filterRiskyFiles removes files with flame score of 0
func filterRiskyFiles(scores []models.FileScore) []models.FileScore {
	var risky []models.FileScore
	for _, s := range scores {
		if s.FlameScore > 0 {
			risky = append(risky, s)
		}
	}
	return risky
}

func init() {
	rootCmd.AddCommand(suggestCmd)
	suggestCmd.Flags().IntVarP(
		&suggestLimit,
		"limit", "l",
		5,
		"Number of suggestions to show (1-100)",
	)
	suggestCmd.Flags().StringVarP(
		&suggestPath,
		"path", "p",
		".",
		"Path to git repository (default: current directory)",
	)
	suggestCmd.Flags().BoolVar(
		&suggestShowAll,
		"all",
		false,
		"Show all risky files without limit",
	)
}
