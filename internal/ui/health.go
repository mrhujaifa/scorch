package ui

import (
	"fmt"
	"strings"

	"github.com/mrhujaifa/flamekit/internal/analyzer"
)

// RenderHealth displays the complete project health report
func RenderHealth(report analyzer.HealthReport) {

	// header
	fmt.Println()
	fmt.Println(headerStyle.Render("  FLAMEKIT — Project Health Report"))
	fmt.Println(subStyle.Render("  Behavioral analysis from git history · 100% local & private"))
	fmt.Println(strings.Repeat("─", 65))
	fmt.Println()

	// health score
	scoreLabel := getScoreLabel(report.CurrentScore)
	scoreBar := renderBar(report.CurrentScore, 100)
	fmt.Printf("  Health Score   %s  %s  %s\n",
		flameStyle.Render(fmt.Sprintf("%d/100", report.CurrentScore)),
		subStyle.Render(scoreBar),
		scoreLabel,
	)
	fmt.Println()

	// trend
	fmt.Println(strings.Repeat("─", 65))
	fmt.Println(subStyle.Render("  TREND  (last 4 months)"))
	fmt.Println(strings.Repeat("─", 65))

	for i, point := range report.TrendScores {
		bar := renderBar(point.Score, 100)
		delta := ""
		if i > 0 && point.Delta != 0 {
			if point.Delta < 0 {
				delta = highStyle.Render(fmt.Sprintf("  ↓ %d", point.Delta))
			} else {
				delta = lowStyle.Render(fmt.Sprintf("  ↑ +%d", point.Delta))
			}
		}
		fmt.Printf("  %-14s  %3d  %s%s\n",
			subStyle.Render(point.Label),
			point.Score,
			subStyle.Render(bar),
			delta,
		)
	}

	fmt.Println()

	// decline or stable message
	if report.MonthlyDecline < 0 {
		fmt.Printf("  %s\n",
			highStyle.Render(fmt.Sprintf(
				"! Declining %d points/month on average",
				-report.MonthlyDecline,
			)),
		)
	} else {
		fmt.Printf("  %s\n",
			lowStyle.Render("✓ Health is stable or improving"),
		)
	}

	// prediction
	if report.PredictedScore < 60 {
		fmt.Printf("  %s\n",
			highStyle.Render(fmt.Sprintf(
				"! Predicted next month: %d/100  ⚠ Approaching CRITICAL",
				report.PredictedScore,
			)),
		)
	} else {
		fmt.Printf("  %s\n",
			subStyle.Render(fmt.Sprintf(
				"Predicted next month: %d/100",
				report.PredictedScore,
			)),
		)
	}

	fmt.Println()

	// velocity risk
	fmt.Println(strings.Repeat("─", 65))
	fmt.Println(subStyle.Render("  VELOCITY RISK  (last 30 days)"))
	fmt.Println(strings.Repeat("─", 65))

	fmt.Printf("  %-20s %d%%\n",
		subStyle.Render("Bug rate:"),
		report.BugRate,
	)
	fmt.Printf("  %-20s %s\n",
		subStyle.Render("Trend:"),
		subStyle.Render(report.BugRateTrend),
	)

	// show risk warning only when critical
	if report.WeeksToRisk > 0 {
		fmt.Printf("  %-20s %s\n",
			subStyle.Render("Risk threshold:"),
			highStyle.Render(fmt.Sprintf("CRITICAL in ~%d weeks", report.WeeksToRisk)),
		)
		fmt.Println()
		fmt.Printf("  %s\n",
			highStyle.Render("! Recommendation: Freeze new features, focus on stability"),
		)
	}

	fmt.Println()

	// stability index
	fmt.Println(strings.Repeat("─", 65))
	fmt.Println(subStyle.Render("  STABILITY INDEX"))
	fmt.Println(strings.Repeat("─", 65))

	fmt.Printf("  %-20s %d\n",
		subStyle.Render("Files analyzed:"),
		report.TotalFiles,
	)
	fmt.Printf("  %-20s %s\n",
		subStyle.Render("Dangerous files:"),
		highStyle.Render(fmt.Sprintf("%d", report.DangerousFiles)),
	)
	fmt.Printf("  %-20s %d\n",
		subStyle.Render("Watch list:"),
		report.WatchFiles,
	)
	fmt.Printf("  %-20s %d\n",
		subStyle.Render("Healthy files:"),
		report.HealthyFiles,
	)

	// top destabilizers only when significant
	if len(report.TopDestabilizers) > 0 &&
		report.TopDestabilizers[0].FlameScore >= 4 {
		fmt.Println()
		fmt.Println(subStyle.Render("  Top destabilizers:"))
		for _, f := range report.TopDestabilizers {
			fmt.Printf("    %s  %s\n",
				fileStyle.Render(f.FilePath),
				highStyle.Render(fmt.Sprintf("→ Flame: %d", f.FlameScore)),
			)
		}
	}

	// footer
	fmt.Println()
	fmt.Println(strings.Repeat("─", 65))
	fmt.Println(subStyle.Render("  Run `flamekit suggest` to see refactoring priorities"))
	fmt.Println()
}

// getScoreLabel returns colored label based on health score
func getScoreLabel(score int) string {
	switch {
	case score >= 80:
		return lowStyle.Render("✓ HEALTHY")
	case score >= 60:
		return medStyle.Render("⚠ WARNING")
	case score >= 40:
		return medStyle.Render("⚠ DECLINING")
	default:
		return highStyle.Render("✕ CRITICAL")
	}
}
