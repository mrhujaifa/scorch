package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/mrhujaifa/scorch/pkg/models"
)

var (
	headerStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15"))
	subStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))
	highStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("1"))
	medStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("3"))
	lowStyle    = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))
	fileStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("12"))
	flameStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("15"))
)

func RenderAnalysis(scores []models.FileScore) {
	high, med, low := countRisks(scores)

	// Header
	fmt.Println()
	fmt.Println(headerStyle.Render("  SCORCH — Codebase Risk Analysis"))
	fmt.Println(subStyle.Render("  Analyzing git history to find dangerous files"))
	fmt.Println()

	// Summary bar
	fmt.Printf("  %s  %s  %s  %s\n",
		subStyle.Render("Files scanned:"),
		flameStyle.Render(fmt.Sprintf("%d", len(scores))),
		subStyle.Render("  |  Commits analyzed via git history"),
		"",
	)
	fmt.Println()

	// Table header
	fmt.Println(strings.Repeat("─", 65))
	fmt.Printf("  %-6s  %-42s  %s\n",
		subStyle.Render("RISK"),
		subStyle.Render("FILE"),
		subStyle.Render("SCORE"),
	)
	fmt.Println(strings.Repeat("─", 65))

	// Rows
	for _, score := range scores {
		risk, riskStyle := getRisk(score.FlameScore)
		bar := renderBar(score.FlameScore, getMaxFlame(scores))

		fmt.Printf("  %-6s  %-42s  %s %s\n",
			riskStyle.Render(risk),
			fileStyle.Render(trimPath(score.FilePath, 42)),
			flameStyle.Render(fmt.Sprintf("%4d", score.FlameScore)),
			subStyle.Render(bar),
		)
	}

	// Footer
	fmt.Println(strings.Repeat("─", 65))
	fmt.Printf("  %s   %s   %s   %s\n",
		subStyle.Render("Total:"),
		highStyle.Render(fmt.Sprintf("High %d", high)),
		medStyle.Render(fmt.Sprintf("Med %d", med)),
		lowStyle.Render(fmt.Sprintf("Low %d", low)),
	)
	fmt.Println()

	// Suggestion
	if high > 0 {
		fmt.Println(highStyle.Render("  ! Action needed: Run `scorch suggest` to see refactor priorities"))
	} else {
		fmt.Println(lowStyle.Render("  ✓ Codebase looks healthy"))
	}
	fmt.Println()
}

// getRisk returns risk level and style based on flame score
func getRisk(flame int) (string, lipgloss.Style) {
	switch {
	case flame >= 10:
		return "HIGH", highStyle
	case flame >= 4:
		return "MED ", medStyle
	default:
		return "LOW ", lowStyle
	}
}

// renderBar creates a simple visual bar based on score
func renderBar(score, max int) string {
	if max == 0 {
		return ""
	}
	filled := (score * 10) / max
	return strings.Repeat("█", filled) + strings.Repeat("░", 10-filled)
}

// trimPath shortens long file paths
func trimPath(path string, max int) string {
	if len(path) <= max {
		return path
	}
	return "..." + path[len(path)-(max-3):]
}

// getMaxFlame finds highest flame score for bar scaling
func getMaxFlame(scores []models.FileScore) int {
	max := 0
	for _, s := range scores {
		if s.FlameScore > max {
			max = s.FlameScore
		}
	}
	return max
}

// countRisks counts files by risk level
func countRisks(scores []models.FileScore) (high, med, low int) {
	for _, s := range scores {
		switch {
		case s.FlameScore >= 10:
			high++
		case s.FlameScore >= 4:
			med++
		default:
			low++
		}
	}
	return
}
