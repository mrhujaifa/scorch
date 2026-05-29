package models

// FileScore holds the calculated risk metrics for a single file
type FileScore struct {
	FilePath    string // relative path: "auth/middleware.go"
	ChurnCount  int    // total number of times file was changed
	BugFixCount int    // commits containing fix/bug/hotfix/patch
	FlameScore  int    // ChurnCount × BugFixCount = risk level
}

// ProjectHealth represents the overall health of the repository
type ProjectHealth struct {
	HealthScore    int // 0-100, higher is better
	TotalFiles     int // total files analyzed
	DangerousFiles int // files with high flame score
	HealthyFiles   int // files with low flame score
}
