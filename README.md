<div align="center">

<img src="https://i.postimg.cc/htBBLfKZ/Gemini-Generated-Image-v3zjoxv3zjoxv3zj-(1)-jukebox-bg-removed.png" alt="FlameKit Logo" width="150" />



**Find where your codebase burns.**

flamekit analyzes your git history to identify the most dangerous files in your codebase — files that change frequently and cause the most bugs, ranked by risk.

<br>

[![GitHub release](https://img.shields.io/github/v/release/mrhujaifa/flamekit?style=flat-square&color=orange)](https://github.com/mrhujaifa/flamekit/releases/latest)
[![GitHub stars](https://img.shields.io/github/stars/mrhujaifa/flamekit?style=flat-square&color=yellow)](https://github.com/mrhujaifa/flamekit/stargazers)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![Go version](https://img.shields.io/github/go-mod/go-version/mrhujaifa/flamekit?style=flat-square)](go.mod)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/mrhujaifa/flamekit/pulls)


</div>

---

## The Problem

Every developer has faced this:

> *"This codebase keeps breaking. But where do I start fixing it?"*

Finding high-risk areas manually means reading thousands of commits — nobody does that. flamekit does it in milliseconds.

---

## How It Works

flamekit calculates a **Flame Score** for every file in your repository:

```
Flame Score = Churn × Bug Fixes

Churn     →  how many times the file was changed
Bug Fixes →  how many commits touched this file with fix/bug/hotfix keywords
```

**High Flame Score = dangerous file that needs your attention.**

---

## Why flamekit

| Tool | Cost | Setup | Works Offline | CLI |
|------|------|-------|---------------|-----|
| CodeScene | $300/month | Complex | ❌ | ❌ |
| SonarQube | Free/Paid | Server needed | ❌ | ❌ |
| GitLens | Free/Paid | VS Code only | ✅ | ❌ |
| **flamekit** | **Free** | **Zero config** | **✅** | **✅** |

---

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

---

## Usage

flamekit works with **any git repository** — Go, JavaScript, Python, Java, or any language.

```bash
cd your-project
flamekit analyze
```

---

## Commands

### `$ flamekit analyze`
Scan your entire codebase and display a flame map — files ranked by risk level.

```bash
$ flamekit analyze
```

```
  flamekit — Codebase Risk Analysis
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

### `$ flamekit suggest`
Get a prioritized list of files to refactor — ranked by maximum impact, with effort estimates and ownership data.

```bash
# Top 5 suggestions (default)
$ flamekit suggest

# Custom limit
$ flamekit suggest --limit 10

# Show all risky files
$ flamekit suggest --all

# Analyze specific repository
$ flamekit suggest --path /path/to/repo
```

```
  Analyzing repository...

  FLAMEKIT — Refactoring Suggestions
  Files ranked by impact — fix these first

  ──────────────────────────────────────────────────────────────────────
  #1  internal/app/app.go
  ──────────────────────────────────────────────────────────────────────
      Metrics        Flame: 45  Changes: 120  Bug Fixes: 23
      Last changed   2 days ago  ⚠ Recently modified
      Last bug       2 days ago  ✕ Active bug area
      Best person    Rahim (78% ownership)

      Est. effort    ~3-4 hours
      Risk reduction ~40% fewer bugs
      Priority score 90/100

      → Modified 120 times with 23 bug fixes — highly unstable.
        Core logic likely needs redesign.

      Next steps:
        $ flamekit file internal/app/app.go
        $ flamekit impact internal/app/app.go

  #2  internal/auth/middleware.go
  ──────────────────────────────────────────────────────────────────────
      Metrics        Flame: 12  Changes: 67  Bug Fixes: 8
      Last changed   5 days ago
      Last bug       5 days ago  ⚠ Recent bug
      Best person    Karim (65% ownership)

      Est. effort    ~1-2 hours
      Risk reduction ~25% fewer bugs
      Priority score 24/100

      → Only 67 changes but 8 bug fixes — every touch breaks something.
        Fragile logic, handle with care.

      Next steps:
        $ flamekit file internal/auth/middleware.go
        $ flamekit impact internal/auth/middleware.go

  ──────────────────────────────────────────────────────────────────────
  Total ROI if all fixed:  ~6 hours work → ~65% bug reduction
  ──────────────────────────────────────────────────────────────────────
  Run `flamekit suggest --all` to see all risky files
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--limit` | `-l` | `5` | Number of suggestions to show (1-100) |
| `--path` | `-p` | `.` | Path to git repository |
| `--all` | — | `false` | Show all risky files without limit |

---
### `$ flamekit health`
Analyze your project's overall health score, trend, and velocity risk — all from git history.

```bash
# Current directory
$ flamekit health

# Specific repository
$ flamekit health --path /path/to/repo
```

```
  FLAMEKIT — Project Health Report
  Behavioral analysis from git history · 100% local & private
  ─────────────────────────────────────────────────────────────────

  Health Score   58/100  █████░░░░░  ⚠ WARNING

  ─────────────────────────────────────────────────────────────────
  TREND  (last 4 months)
  ─────────────────────────────────────────────────────────────────
  3 months ago    82  ████████░░
  2 months ago    71  ███████░░░  ↓ -11
  Last month      63  ██████░░░░  ↓ -8
  Now             58  █████░░░░░  ↓ -5

  ! Declining 8 points/month on average
  ! Predicted next month: 50/100  ⚠ Approaching CRITICAL

  ─────────────────────────────────────────────────────────────────
  VELOCITY RISK  (last 30 days)
  ─────────────────────────────────────────────────────────────────
  Bug rate:            35%
  Trend:               ↑ Increasing
  Risk threshold:      CRITICAL in ~6 weeks

  ! Recommendation: Freeze new features, focus on stability

  ─────────────────────────────────────────────────────────────────
  STABILITY INDEX
  ─────────────────────────────────────────────────────────────────
  Files analyzed:      142
  Dangerous files:     3
  Watch list:          50
  Healthy files:       89

  Top destabilizers:
    internal/auth/middleware.go   → Flame: 45
    services/payment/handler.go  → Flame: 32
    internal/user/service.go     → Flame: 18

  ─────────────────────────────────────────────────────────────────
  Run `flamekit suggest` to see refactoring priorities
```

**Flags:**

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--path` | `-p` | `.` | Path to git repository |

**Health Score Levels:**

| Score | Status | Meaning |
|-------|--------|---------|
| 80-100 | ✓ HEALTHY | Low risk, codebase is stable |
| 60-79 | ⚠ WARNING | Some risk areas, monitor closely |
| 40-59 | ⚠ DECLINING | Health dropping, action needed |
| 0-39 | ✕ CRITICAL | High risk, immediate action required |

## Roadmap

| Command | Status | Description |
|---------|--------|-------------|
| `flamekit analyze` | ✅ Available | Codebase flame map |
| `flamekit suggest` | ✅ Available | Refactoring priorities |
| `flamekit health`  | ✅ Available | Project health score 0-100 |
| `flamekit file <path>` | 🔨 In Progress | Single file deep dive |
| `flamekit coupling` | 📋 Planned | Hidden dependency detection |
| `flamekit who` | 📋 Planned | Knowledge ownership map |
| `flamekit impact <file>` | 📋 Planned | Blast radius before touching a file |
| `flamekit predict` | 📋 Planned | 30-day bug probability forecast |
| `flamekit onboard` | 📋 Planned | New developer codebase guide |
| `flamekit debt` | 📋 Planned | Technical debt estimator |
| `flamekit bus-factor` | 📋 Planned | Team knowledge risk analysis |
| `flamekit guardian` | 🤖 AI Feature | Real-time change safety net |
| `flamekit explain <file>` | 🤖 AI Feature | AI-powered risk explanation |

---

## Who Is flamekit For

- **Junior Developers** — Understand a new codebase in minutes, not weeks
- **Senior Developers** — Make data-driven refactoring decisions
- **Tech Leads** — Identify team knowledge silos and technical debt
- **Code Reviewers** — Get historical context before reviewing a PR
- **Open Source Contributors** — Find where help is needed most

---

## Supported Languages

flamekit analyzes any git repository regardless of language. File filtering supports:

`.go` `.js` `.jsx` `.ts` `.tsx` `.py` `.java` `.rs` `.c` `.cpp` `.rb` `.php` `.swift` `.kt` `.vue` `.svelte`

---

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feat/amazing-feature`)
3. Commit your changes (`git commit -m 'feat: add amazing feature'`)
4. Push to the branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

---

## License

MIT © [Md. Hujaifa Islam Shanto](https://github.com/mrhujaifa)

---

<div align="center">

**If flamekit helped you, please consider giving it a ⭐**

[Report Bug](https://github.com/mrhujaifa/flamekit/issues) · [Request Feature](https://github.com/mrhujaifa/flamekit/issues) · [Discussions](https://github.com/mrhujaifa/flamekit/discussions)

</div>
