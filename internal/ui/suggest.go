package ui

import (
	"fmt"
	"strings"

	"github.com/mrhujaifa/scorch/pkg/models"
)

func RenderSuggestions(scores []models.FileScore) {
	fmt.Println()
	fmt.Println(headerStyle.Render("  SCORCH — Refactoring Suggestions"))
	fmt.Println(subStyle.Render("  Files ranked by impact — fix these first"))
	fmt.Println()
	fmt.Println(strings.Repeat("─", 65))

	for i, score := range scores {
		rank := fmt.Sprintf("#%d", i+1)
		_, riskStyle := GetRisk(score.FlameScore)

		fmt.Printf("  %s  %s\n",
			flameStyle.Render(rank),
			fileStyle.Render(score.FilePath),
		)

		fmt.Printf("      %s  %s  %s\n",
			subStyle.Render(fmt.Sprintf("Flame: %d", score.FlameScore)),
			subStyle.Render(fmt.Sprintf("Changes: %d", score.ChurnCount)),
			subStyle.Render(fmt.Sprintf("Bug Fixes: %d", score.BugFixCount)),
		)

		fmt.Printf("      %s\n",
			riskStyle.Render(getSuggestion(score)),
		)
		fmt.Println()
	}

	fmt.Println(strings.Repeat("─", 65))
	fmt.Println(subStyle.Render("  Run `scorch file <path>` for deep dive into any file"))
	fmt.Println()
}

func getSuggestion(score models.FileScore) string {
	switch {
	case score.FlameScore >= 10:
		return "→ High change frequency with recurring bugs. Refactor immediately."
	case score.FlameScore >= 4:
		return "→ Historically bug-prone. Review carefully before next change."
	default:
		return "→ Mildly active. Monitor but no immediate action needed."
	}
}
