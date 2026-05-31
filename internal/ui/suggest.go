package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/mrhujaifa/flamekit/internal/analyzer"
	"github.com/mrhujaifa/flamekit/pkg/models"
)

func RenderSuggestions(scores []models.FileScore) {
	fmt.Println()
	fmt.Println(headerStyle.Render("  FLAMEKIT — Refactoring Suggestions"))
	fmt.Println(subStyle.Render("  Files ranked by impact — fix these first"))
	fmt.Println()
	fmt.Println(strings.Repeat("─", 70))

	for i, score := range scores {
		rank := fmt.Sprintf("#%d", i+1)
		_, riskStyle := GetRisk(score.FlameScore)

		// rank + file
		fmt.Printf("\n  %s  %s\n",
			flameStyle.Render(rank),
			fileStyle.Render(score.FilePath),
		)

		fmt.Println(strings.Repeat("─", 70))

		// metrics
		fmt.Printf("      %-16s %s  %s  %s\n",
			subStyle.Render("Metrics"),
			subStyle.Render(fmt.Sprintf("Flame: %d", score.FlameScore)),
			subStyle.Render(fmt.Sprintf("Changes: %d", score.ChurnCount)),
			subStyle.Render(fmt.Sprintf("Bug Fixes: %d", score.BugFixCount)),
		)

		// last changed
		if !score.LastChanged.IsZero() {
			lastChanged := formatDate(score.LastChanged)
			warning := ""
			if time.Since(score.LastChanged).Hours() < 72 {
				warning = warnStyle().Render("  ⚠ Recently modified")
			}
			fmt.Printf("      %-16s %s%s\n",
				subStyle.Render("Last changed"),
				subStyle.Render(lastChanged),
				warning,
			)
		}

		// last bug
		if !score.LastBugDate.IsZero() {
			lastBug := formatDate(score.LastBugDate)
			warning := ""
			if time.Since(score.LastBugDate).Hours() < 168 {
				warning = highStyle.Render("  ✕ Active bug area")
			}
			fmt.Printf("      %-16s %s%s\n",
				subStyle.Render("Last bug"),
				subStyle.Render(lastBug),
				warning,
			)
		}

		// top author
		if score.TopAuthor != "" {
			fmt.Printf("      %-16s %s\n",
				subStyle.Render("Best person"),
				subStyle.Render(fmt.Sprintf("%s (%d%% ownership)", score.TopAuthor, score.TopAuthorPct)),
			)
		}

		fmt.Println()

		// effort + reduction + priority
		effort, reduction, priority := analyzer.CalculateROI(score)
		fmt.Printf("      %-16s %s\n",
			subStyle.Render("Est. effort"),
			subStyle.Render(effort),
		)
		fmt.Printf("      %-16s %s\n",
			subStyle.Render("Risk reduction"),
			subStyle.Render(reduction),
		)
		fmt.Printf("      %-16s %s\n",
			subStyle.Render("Priority score"),
			flameStyle.Render(fmt.Sprintf("%d/100", priority)),
		)

		fmt.Println()

		// dynamic insight
		fmt.Printf("      %s\n",
			riskStyle.Render("→ "+analyzer.GenerateInsight(score)),
		)

		// next steps
		fmt.Println()
		fmt.Printf("      %s\n", subStyle.Render("Next steps:"))
		fmt.Printf("        %s\n", subStyle.Render(fmt.Sprintf("$ flamekit file %s", score.FilePath)))
		if score.FlameScore >= 4 {
			fmt.Printf("        %s\n", subStyle.Render(fmt.Sprintf("$ flamekit impact %s", score.FilePath)))
		}
	}

	fmt.Println()
	fmt.Println(strings.Repeat("─", 70))

	// total ROI
	totalHours, totalReduction := analyzer.CalculateTotalROI(scores)
	fmt.Printf("  %s  %s\n",
		subStyle.Render("Total ROI if all fixed:"),
		flameStyle.Render(fmt.Sprintf("~%s work → ~%d%% bug reduction", totalHours, totalReduction)),
	)
	fmt.Println(strings.Repeat("─", 70))
	fmt.Println(subStyle.Render("  Run `flamekit suggest --all` to see all files"))
	fmt.Println()
}

func formatDate(t time.Time) string {
	hours := time.Since(t).Hours()
	switch {
	case hours < 24:
		return "today"
	case hours < 48:
		return "yesterday"
	case hours < 168:
		return fmt.Sprintf("%d days ago", int(hours/24))
	case hours < 720:
		return fmt.Sprintf("%d weeks ago", int(hours/168))
	default:
		return fmt.Sprintf("%d months ago", int(hours/720))
	}
}

func warnStyle() lipgloss.Style {
	return medStyle
}
