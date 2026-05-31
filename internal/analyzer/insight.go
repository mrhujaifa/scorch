package analyzer

import (
	"fmt"

	"github.com/mrhujaifa/flamekit/pkg/models"
)

// GenerateInsight produces a dynamic, context-aware message
// based on the file's churn and bug fix history
func GenerateInsight(score models.FileScore) string {
	highChurn := score.ChurnCount > 10
	highBug := score.BugFixCount > 3

	switch {
	case highChurn && highBug:
		return fmt.Sprintf(
			"Modified %d times with %d bug fixes — highly unstable. Core logic likely needs redesign.",
			score.ChurnCount, score.BugFixCount,
		)
	case highChurn && !highBug:
		return fmt.Sprintf(
			"Changed %d times but only %d bug fixes — frequently modified yet surprisingly stable. Monitor active development.",
			score.ChurnCount, score.BugFixCount,
		)
	case !highChurn && highBug:
		return fmt.Sprintf(
			"Only %d changes but %d bug fixes — every touch breaks something. Fragile logic, handle with care.",
			score.ChurnCount, score.BugFixCount,
		)
	default:
		return fmt.Sprintf(
			"Low activity with %d changes and %d bug fixes — stable file, no immediate action needed.",
			score.ChurnCount, score.BugFixCount,
		)
	}
}

// CalculateROI returns effort estimate, risk reduction, and priority score
func CalculateROI(score models.FileScore) (effort string, reduction string, priority int) {
	// effort estimate
	switch {
	case score.FlameScore >= 10:
		effort = "~3-4 hours"
	case score.FlameScore >= 4:
		effort = "~1-2 hours"
	default:
		effort = "~30 minutes"
	}

	// risk reduction
	switch {
	case score.BugFixCount >= 10:
		reduction = "~40% fewer bugs"
	case score.BugFixCount >= 5:
		reduction = "~25% fewer bugs"
	case score.BugFixCount >= 2:
		reduction = "~15% fewer bugs"
	default:
		reduction = "~5% fewer bugs"
	}

	// priority score 0-100
	priority = score.FlameScore * 2
	if priority > 100 {
		priority = 100
	}

	return
}

// CalculateTotalROI returns total effort and bug reduction for all files
func CalculateTotalROI(scores []models.FileScore) (totalHours string, totalReduction int) {
	hours := 0
	for _, score := range scores {
		switch {
		case score.FlameScore >= 10:
			hours += 4
		case score.FlameScore >= 4:
			hours += 2
		default:
			hours += 1
		}
		totalReduction += score.BugFixCount
	}

	if hours >= 8 {
		totalHours = fmt.Sprintf("%d days", hours/8)
	} else {
		totalHours = fmt.Sprintf("%d hours", hours)
	}

	// cap at 80%
	if totalReduction > 80 {
		totalReduction = 80
	}

	return
}
