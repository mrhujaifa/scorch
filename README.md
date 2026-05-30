# Scorch ⚡

> Analyze your Git history and uncover the files that burn your project.

Scorch is a lightweight CLI tool that identifies high-risk files in your repository by combining code churn with bug-fix history. It helps developers and teams prioritize refactoring, reduce technical debt, and focus on the files that matter most.

## What Scorch Solves

- Find the files that change most often and attract the most bug fixes
- Rank dangerous areas of the codebase by impact
- Turn Git history into actionable refactoring guidance
- Use a zero-setup CLI tool instead of heavy analysis platforms

## Key Features

- `scorch analyze` — build a flame score map from Git history
- `scorch suggest` — get prioritized refactoring recommendations
- Code-aware filtering for common source file types
- CLI-first workflow for fast adoption in developer environments

## Why Use Scorch

- **Minimal setup** — works directly in any Git repository
- **Free and open source** — no server or license barrier
- **Developer-friendly** — designed for fast insights and clear next steps
- **Actionable output** — suggests what to fix first, not just what is broken

## Installation

```bash
go install github.com/mrhujaifa/scorch@latest
```

## Usage

```bash
# Analyze your repository and print the flame map
scorch analyze

# Show the top refactoring suggestions
scorch suggest

# Limit suggestions to only the top 3 files
scorch suggest --limit 3
```

## Example Output

```text
SCORCH — Refactoring Suggestions
Files ranked by impact — fix these first

  #1  internal/analyzer/flame.go
      Flame: 24  Changes: 12  Bug Fixes: 2
      → High change frequency with recurring bugs. Refactor immediately.

  #2  cmd/suggest.go
      Flame: 15  Changes: 10  Bug Fixes: 1
      → Historically bug-prone. Review carefully before next change.
```

## Roadmap

- [x] `scorch analyze` — codebase flame map
- [x] `scorch suggest` — refactoring priorities
- [ ] `scorch health` — project health score
- [ ] `scorch file` — single file deep dive
- [ ] `scorch impact` — blast radius analysis
- [ ] `scorch predict` — bug prediction

## Contributing

Bug reports, feature requests, and pull requests are welcome. If you want to improve Scorch, start by opening an issue or submitting a PR.

## License

MIT © Md. Hujaifa Islam Shanto
