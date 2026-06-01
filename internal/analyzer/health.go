package analyzer

import (
	"time"

	"github.com/mrhujaifa/flamekit/internal/git"
	"github.com/mrhujaifa/flamekit/pkg/models"
)

// HealthReport holds complete project health analysis
type HealthReport struct {
	CurrentScore     int
	TrendScores      []TrendPoint
	MonthlyDecline   int
	PredictedScore   int
	TotalFiles       int
	DangerousFiles   int
	HealthyFiles     int
	WatchFiles       int
	TopDestabilizers []models.FileScore
	BugRate          int
	BugRateTrend     string
	VelocityRisk     string
	WeeksToRisk      int
}

// TrendPoint holds health score at a point in time
type TrendPoint struct {
	Label string
	Score int
	Delta int
}

func CalculateScore(scores []models.FileScore) int {
	penalty := 0
	for _, s := range scores {
		switch {
		case s.FlameScore >= 10:
			penalty += 10
		case s.FlameScore >= 4:
			penalty += 5
		default:
			penalty += 1
		}
	}

	score := 100 - penalty

	if score < 0 {
		return 0
	}

	return score
}

func CalculateHealth(commits []git.CommitData, scores []models.FileScore) HealthReport {
	report := HealthReport{}

	report.TotalFiles = len(scores)

	report.CurrentScore = CalculateScore(scores)

	for _, s := range scores {
		switch {
		case s.FlameScore >= 10:
			report.DangerousFiles++
		case s.FlameScore >= 4:
			report.WatchFiles++
		default:
			report.HealthyFiles++
		}
	}

	// add this line
	report.TopDestabilizers = GetTopDestabilizers(scores, 3)
	report.TrendScores = CalculateTrend(commits)
	// monthly decline
	if len(report.TrendScores) >= 2 {
		total := 0
		for i := 1; i < len(report.TrendScores); i++ {
			total += report.TrendScores[i].Delta
		}
		report.MonthlyDecline = total / (len(report.TrendScores) - 1)
	}

	// prediction
	report.PredictedScore = report.CurrentScore + report.MonthlyDecline
	if report.PredictedScore < 0 {
		report.PredictedScore = 0
	}
	report.BugRate, report.BugRateTrend, report.VelocityRisk, report.WeeksToRisk = CalculateVelocityRisk(commits)
	return report
}

func GetTopDestabilizers(scores []models.FileScore, n int) []models.FileScore {
	sorted := SortByFlame(scores)

	if len(sorted) > n {
		return sorted[:n]
	}

	return sorted
}

func CalculateVelocityRisk(commits []git.CommitData) (bugRate int, trend string, risk string, weeks int) {
	now := time.Now()
	last30 := now.AddDate(0, 0, -30)
	prev30 := now.AddDate(0, 0, -60)

	var recentTotal, recentBugs int
	var prevTotal, prevBugs int

	for _, commit := range commits {
		if commit.Date.After(last30) {
			recentTotal++
			if commit.IsBug {
				recentBugs++
			}
		} else if commit.Date.After(prev30) {
			prevTotal++
			if commit.IsBug {
				prevBugs++
			}
		}
	}

	if recentTotal > 0 {
		bugRate = (recentBugs * 100) / recentTotal
	}

	prevRate := 0
	if prevTotal > 0 {
		prevRate = (prevBugs * 100) / prevTotal
	}

	// compare
	switch {
	case bugRate > prevRate:
		trend = "↑ Increasing"
		risk = "WARNING"
		weeks = 6
	case bugRate < prevRate:
		trend = "↓ Decreasing"
		risk = "STABLE"
	default:
		trend = "→ Stable"
		risk = "STABLE"
	}

	return
}

func CalculateTrend(commits []git.CommitData) []TrendPoint {
	now := time.Now()

	points := []TrendPoint{
		{Label: "3 months ago", Score: scoreAtTime(commits, now.AddDate(0, -3, 0))},
		{Label: "2 months ago", Score: scoreAtTime(commits, now.AddDate(0, -2, 0))},
		{Label: "Last month", Score: scoreAtTime(commits, now.AddDate(0, -1, 0))},
		{Label: "Now", Score: scoreAtTime(commits, now)},
	}

	// delta calculate করো
	for i := 1; i < len(points); i++ {
		points[i].Delta = points[i].Score - points[i-1].Score
	}

	return points
}

func scoreAtTime(commits []git.CommitData, before time.Time) int {
	var filtered []git.CommitData
	for _, c := range commits {
		if c.Date.Before(before) {
			filtered = append(filtered, c)
		}
	}

	// no commits at this time → perfect score
	if len(filtered) == 0 {
		return 100
	}

	scores, _ := CalculateFlameScores(filtered)
	scores = FilterCodeFiles(scores)
	return CalculateScore(scores)
}
