package analyzer

import (
	"path/filepath"
	"sort"

	"github.com/mrhujaifa/flamekit/internal/git"
	"github.com/mrhujaifa/flamekit/pkg/models"
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

func SortByFlame(results []models.FileScore) []models.FileScore {
	sort.Slice(results, func(i, j int) bool {
		return results[i].FlameScore > results[j].FlameScore
	})

	return results
}

var extention = []string{
	".go", ".js", ".ts", ".py", ".java", ".rs", ".cpp", ".c",
}

func FilterCodeFiles(results []models.FileScore) []models.FileScore {

	var filtered []models.FileScore
	for _, result := range results {
		ext := filepath.Ext(result.FilePath)

		for _, exten := range extention {
			if ext == exten {
				filtered = append(filtered, result)
				break
			}
		}

	}

	return filtered
}
