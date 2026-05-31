package analyzer

import (
	"fmt"
	"path/filepath"
	"sort"

	"github.com/mrhujaifa/flamekit/internal/git"
	"github.com/mrhujaifa/flamekit/pkg/models"
)

func CalculateFlameScores(commits []git.CommitData) ([]models.FileScore, error) {
	if len(commits) == 0 {
		return nil, fmt.Errorf("no commits provided")
	}

	fileMap := make(map[string]*models.FileScore)
	authorMap := make(map[string]map[string]int)

	for _, commit := range commits {
		for _, file := range commit.Files {
			// initialize if new file
			if fileMap[file] == nil {
				fileMap[file] = &models.FileScore{
					FilePath: file,
				}
				authorMap[file] = make(map[string]int)
			}

			// churn count
			fileMap[file].ChurnCount++

			// bug fix count
			if commit.IsBug {
				fileMap[file].BugFixCount++

				// last bug date
				if commit.Date.After(fileMap[file].LastBugDate) {
					fileMap[file].LastBugDate = commit.Date
				}
			}

			// last changed date
			if commit.Date.After(fileMap[file].LastChanged) {
				fileMap[file].LastChanged = commit.Date
			}

			// author tracking
			authorMap[file][commit.Author]++

			// flame score
			fileMap[file].FlameScore =
				fileMap[file].ChurnCount *
					fileMap[file].BugFixCount
		}
	}

	// calculate top author
	var results []models.FileScore
	for filePath, score := range fileMap {
		topAuthor, topCount := getTopAuthor(authorMap[filePath])
		score.TopAuthor = topAuthor
		if score.ChurnCount > 0 {
			score.TopAuthorPct = (topCount * 100) / score.ChurnCount
		}
		results = append(results, *score)
	}

	return results, nil
}

// getTopAuthor finds the author with most commits for a file
func getTopAuthor(authors map[string]int) (string, int) {
	topAuthor := ""
	topCount := 0

	for author, count := range authors {
		if count > topCount {
			topCount = count
			topAuthor = author
		}
	}

	return topAuthor, topCount
}

func SortByFlame(results []models.FileScore) []models.FileScore {
	sort.Slice(results, func(i, j int) bool {
		return results[i].FlameScore > results[j].FlameScore
	})

	return results
}

var extention = []string{
	".go", ".js", ".ts", ".tsx", ".jsx", ".php", "swift", ".kt", ".vue", ".svelte", ".rb", ".py", ".java", ".rs", ".cpp", ".c",
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
