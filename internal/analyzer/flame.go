package analyzer

import (
	"github.com/mrhujaifa/scorch/internal/git"
	"github.com/mrhujaifa/scorch/pkg/models"
)

func CalculateFlameScores(commits []git.CommitData) ([]models.FileScore, error) {
	fileMap := make(map[string]*models.FileScore)
	for _, commit := range commits {
		for _, file := range commit.Files {
			// আগে check করো file আছে কিনা
			if fileMap[file] == nil {
				fileMap[file] = &models.FileScore{
					FilePath: file,
				}
			}

			// তারপর count বাড়াও
			fileMap[file].ChurnCount++

			if commit.IsBug {
				fileMap[file].BugFixCount++
			}

			fileMap[file].FlameScore =
				fileMap[file].ChurnCount *
					fileMap[file].BugFixCount
		}
	}

	var results []models.FileScore

	for _, score := range fileMap {
		results = append(results, *score)
	}

	return results, nil
}
