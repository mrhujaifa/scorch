package git

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type CommitData struct {
	Hash    string
	Message string
	IsBug   bool
	Files   []string
	Author  string    // ← add
	Date    time.Time // ← add
}

var BugKeywords = []string{
	// Direct
	"fix", "fixes", "fixed",
	"bug", "bugfix",
	"hotfix", "hotpatch",
	"patch", "patched",

	// Solving
	"solve", "solved", "solves",
	"resolve", "resolved", "resolves",
	"repair", "repaired",
	"correct", "corrected",

	// Problems
	"issue", "problem", "defect",
	"crash", "crashes", "crashed",
	"error", "broken", "breaks",
	"fail", "failing", "failed",
	"wrong", "oops", "mistake",

	// Reverting
	"revert", "rollback", "undo",
}

func ParseRepository(repoPath string) ([]CommitData, error) {

	if strings.TrimSpace(repoPath) == "" {
		return nil, fmt.Errorf("repository path cannot be empty")
	}
	openRepo, err := OpenRepository(repoPath)

	if err != nil {
		return nil, fmt.Errorf("failed to openRepos: %w", err)
	}

	getrepo, err := GetCommits(openRepo)

	if err != nil {
		return nil, fmt.Errorf("failed to get commits: %w", err)
	}

	var result []CommitData

	for _, commit := range getrepo {
		files, err := GetChangedFiles(commit)

		if err != nil {
			return nil, fmt.Errorf("failed to get changed files: %w", err)
		}

		data := CommitData{
			Hash:    commit.Hash.String(),
			Message: commit.Message,
			IsBug:   IsBugCommit(commit.Message),
			Files:   files,
			Author:  commit.Author.Name,
			Date:    commit.Author.When,
		}

		result = append(result, data)
	}

	return result, nil

}

func OpenRepository(repoPath string) (*git.Repository, error) {
	if strings.TrimSpace(repoPath) == "" {
		return nil, fmt.Errorf("repository path cannot be empty")
	}

	if _, err := os.Stat(repoPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("path does not exist: %s", repoPath)
	}

	repo, err := git.PlainOpen(repoPath)
	if err != nil {
		return nil, fmt.Errorf("not a valid git repository: %s", repoPath)
	}

	return repo, nil
}

func GetCommits(repo *git.Repository) ([]*object.Commit, error) {
	if repo == nil {
		return nil, fmt.Errorf("repository cannot be nil")
	}

	commits, err := repo.Log(&git.LogOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to read commit history: %w", err)
	}
	defer commits.Close()

	var result []*object.Commit

	err = commits.ForEach(func(c *object.Commit) error {

		if c == nil {
			return nil
		}
		result = append(result, c)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to iterate commits: %w", err)
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("no commits found in repository")
	}

	return result, nil
}

func GetChangedFiles(c *object.Commit) ([]string, error) {
	// validation: nil commit
	if c == nil {
		return nil, fmt.Errorf("commit cannot be nil")
	}

	// first commit এর কোনো parent নেই, skip করো
	if c.NumParents() == 0 {
		return []string{}, nil
	}

	parent, err := c.Parent(0)
	if err != nil {
		return nil, fmt.Errorf("failed to get parent commit: %w", err)
	}

	patch, err := parent.Patch(c)
	if err != nil {
		return nil, fmt.Errorf("failed to get patch: %w", err)
	}

	var files []string

	for _, filePatch := range patch.FilePatches() {
		from, to := filePatch.Files()

		if to != nil {
			files = append(files, to.Path())
		} else if from != nil {
			files = append(files, from.Path())
		}
	}

	return files, nil
}

func IsBugCommit(message string) bool {
	// validation: empty message
	if strings.TrimSpace(message) == "" {
		return false
	}

	lower := strings.ToLower(message)

	for _, keyword := range BugKeywords {
		if strings.Contains(lower, keyword) {
			return true
		}
	}

	return false
}
