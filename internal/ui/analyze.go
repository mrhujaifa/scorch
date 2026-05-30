package ui

import (
	"fmt"
	"strings"

	"github.com/mrhujaifa/scorch/pkg/models"
)

func RenderAnalysis(scores []models.FileScore) {
	high, med, low := countRisks(scores)

	fmt.Println()
	fmt.Println(headerStyle.Render("  SCORCH — Codebase Risk Analysis"))
	fmt.Println(subStyle.Render("  Analyzing git history to find dangerous files"))
	fmt.Println()

	fmt.Printf("  %s  %s\n",
		subStyle.Render("Files scanned:"),
		flameStyle.Render(fmt.Sprintf("%d", len(scores))),
	)
	fmt.Println()

	fmt.Println(strings.Repeat("─", 65))
	fmt.Printf("  %-6s  %-42s  %s\n",
		subStyle.Render("RISK"),
		subStyle.Render("FILE"),
		subStyle.Render("SCORE"),
	)
	fmt.Println(strings.Repeat("─", 65))

	for _, score := range scores {
		risk, riskStyle := GetRisk(score.FlameScore)
		bar := renderBar(score.FlameScore, getMaxFlame(scores))

		fmt.Printf("  %-6s  %-42s  %s %s\n",
			riskStyle.Render(risk),
			fileStyle.Render(trimPath(score.FilePath, 42)),
			flameStyle.Render(fmt.Sprintf("%4d", score.FlameScore)),
			subStyle.Render(bar),
		)
	}

	fmt.Println(strings.Repeat("─", 65))
	fmt.Printf("  %s   %s   %s   %s\n",
		subStyle.Render("Total:"),
		highStyle.Render(fmt.Sprintf("High: %d", high)),
		medStyle.Render(fmt.Sprintf("Med: %d", med)),
		lowStyle.Render(fmt.Sprintf("Low: %d", low)),
	)
	fmt.Println()

	if high > 0 {
		fmt.Println(highStyle.Render("  ! Action needed: Run `scorch suggest` to see refactor priorities"))
	} else {
		fmt.Println(lowStyle.Render("  ✓ Codebase looks healthy"))
	}
	fmt.Println()
}

func renderBar(score, max int) string {
	if max == 0 {
		return ""
	}
	filled := (score * 10) / max
	return strings.Repeat("█", filled) + strings.Repeat("░", 10-filled)
}

func trimPath(path string, max int) string {
	if len(path) <= max {
		return path
	}
	return "..." + path[len(path)-(max-3):]
}

func getMaxFlame(scores []models.FileScore) int {
	max := 0
	for _, s := range scores {
		if s.FlameScore > max {
			max = s.FlameScore
		}
	}
	return max
}

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
