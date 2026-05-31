# Flamekit ⚡

Flamekit is a lightweight CLI tool that identifies high-risk files in your repository by combining code churn with bug-fix history. It helps developers and teams prioritize refactoring, reduce technical debt, and focus on the files that matter most.

[![GitHub release](https://img.shields.io/github/v/release/mrhujaifa/flamekit?style=flat-square&color=orange)](https://github.com/mrhujaifa/flamekit/releases/latest)
[![GitHub stars](https://img.shields.io/github/stars/mrhujaifa/flamekit?style=flat-square&color=yellow)](https://github.com/mrhujaifa/flamekit/stargazers)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![Go version](https://img.shields.io/github/go-mod/go-version/mrhujaifa/flamekit?style=flat-square)](go.mod)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/mrhujaifa/flamekit/pulls)

![demo](demo.gif)

## What Flamekit Solves

- `flamekit analyze` — build a flame score map from Git history
- `flamekit suggest` — get prioritized refactoring recommendations
- Code-aware filtering for common source file types
- CLI-first workflow for fast adoption in developer environments

## Why Use Flamekit

Every developer has faced this challenge: finding the most critical files to refactor in a large codebase. Manual analysis of thousands of commits is impractical. Flamekit solves this by automatically identifying high-risk files based on code churn and bug-fix history.

## How Flamekit Works

Flamekit calculates a **Flame Score** for every file in your repository:

```
Flame Score = Churn × Bug Fixes

Churn     →  how many times the file was changed
Bug Fixes →  how many commits touched this file with fix/bug/hotfix keywords
```

**High Flame Score = dangerous file that needs your attention.**

## Flamekit vs. Other Tools

| Tool | Cost | Setup | Works Offline | CLI |
|------|------|-------|---------------|-----|
| CodeScene | $300/month | Complex | ❌ | ❌ |
| SonarQube | Free/Paid | Server needed | ❌ | ❌ |
| GitLens | Free/Paid | VS Code only | ✅ | ❌ |
| **Flamekit** | **Free** | **Zero config** | **✅** | **✅** |

## Installation

### Option 1 — Go Install (Recommended)

```bash
go install github.com/mrhujaifa/flamekit@latest
```

### Option 2 — Manual Download (No Go required)

**macOS / Linux:**
```bash
# macOS (Apple Silicon)
curl -sSL https://github.com/mrhujaifa/flamekit/releases/latest/download/flamekit_Darwin_arm64.tar.gz | tar -xz
sudo mv flamekit /usr/local/bin/

# macOS (Intel)
curl -sSL https://github.com/mrhujaifa/flamekit/releases/latest/download/flamekit_Darwin_x86_64.tar.gz | tar -xz
sudo mv flamekit /usr/local/bin/

# Linux
curl -sSL https://github.com/mrhujaifa/flamekit/releases/latest/download/flamekit_Linux_x86_64.tar.gz | tar -xz
sudo mv flamekit /usr/local/bin/
```

**Windows:**
1. Go to [Releases](https://github.com/mrhujaifa/flamekit/releases/latest)
2. Download `flamekit_Windows_x86_64.zip`
3. Extract and add `flamekit.exe` to your PATH

**Verify installation:**
```bash
flamekit --version
```

## Usage

Flamekit works with **any git repository** — Go, JavaScript, Python, Java, or any language.

```bash
cd your-project
flamekit analyze
```

## Commands

### `$ flamekit analyze`
Scan your entire codebase and display a flame map — files ranked by risk level.

```bash
flamekit analyze
```

```
  FLAMEKIT — Codebase Risk Analysis
  Analyzing git history to find dangerous files

  Files scanned: 17  |  Commits analyzed via git history

  ─────────────────────────────────────────────────────────────────
  RISK    FILE                                     SCORE
  ─────────────────────────────────────────────────────────────────
  HIGH    internal/app/app.go                         45 ██████████
  MED     internal/auth/middleware.go                 12 ████░░░░░░
  LOW     utils/helpers.go                             1 █░░░░░░░░░
  ─────────────────────────────────────────────────────────────────
  Total:   High: 1   Med: 1   Low: 15

  ! Action needed: Run `flamekit suggest` to see refactor priorities
```

**Risk Levels:**
| Level | Flame Score | Meaning |
|-------|-------------|---------|
| HIGH | ≥ 10 | Frequently changed, historically bug-prone. Act now. |
| MED | ≥ 4 | Moderate risk. Review carefully before touching. |
| LOW | < 4 | Stable. No immediate action needed. |

### `$ flamekit suggest`
Get a prioritized list of files to refactor — ranked by maximum impact.

```bash
# Top 5 suggestions (default)
flamekit suggest

# Custom limit
flamekit suggest --limit 10
flamekit suggest --limit 3
```

```
  FLAMEKIT — Refactoring Suggestions
  Files ranked by impact — fix these first

  ─────────────────────────────────────────────────────────────────
  #1  internal/analyzer/flame.go
      Flame: 24  Changes: 12  Bug Fixes: 2
      → High change frequency with recurring bugs. Refactor immediately.

  #2  internal/auth/middleware.go
      Flame: 12  Changes: 67  Bug Fixes: 8
      → Historically bug-prone. Review carefully before next change.

  #3  services/payment/handler.go
      Flame: 8   Changes: 45  Bug Fixes: 5
      → Mildly active. Monitor but no immediate action needed.
  ─────────────────────────────────────────────────────────────────
  Run `flamekit file <path>` for deep dive into any file
```

**Flags:**
| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--limit` | `-l` | `5` | Number of suggestions to show (1-50) |

## Contributing

Contributions are welcome! Please check out [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

## License

Flamekit is licensed under the MIT License — see [LICENSE](LICENSE) for details.

## Security

For security concerns, please review [SECURITY.md](SECURITY.md).

---

Built with ❤️ by [Muhammad Hujaifa](https://github.com/mrhujaifa)
