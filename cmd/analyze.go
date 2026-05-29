package cmd

import (
	"fmt"

	"github.com/mrhujaifa/scorch/internal/analyzer"
	"github.com/mrhujaifa/scorch/internal/git"
	"github.com/mrhujaifa/scorch/internal/ui"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze your codebase and show flame scores",
	Run: func(cmd *cobra.Command, args []string) {
		parseRepo, err := git.ParseRepository(".")

		if err != nil {
			fmt.Println("failed to parse repository")
			return
		}

		calcFlames, err := analyzer.CalculateFlameScores(parseRepo)

		if err != nil {
			fmt.Println("failed to calc falme score")
		}

		calcFlames = analyzer.FilterCodeFiles(calcFlames)
		calcFlames = analyzer.SortByFlame(calcFlames)

		ui.RenderAnalysis(calcFlames)
	},
}

func init() {
	rootCmd.AddCommand(analyzeCmd)
}
