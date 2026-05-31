<div align="center">

# ⚡ flamekit

**Find where your codebase burns.**

flamekit analyzes your git history to identify the most dangerous files in your codebase — files that change frequently and cause the most bugs, ranked by risk.

[![GitHub release](https://img.shields.io/github/v/release/mrhujaifa/flamekit?style=flat-square&color=orange)](https://github.com/mrhujaifa/flamekit/releases/latest)
[![GitHub stars](https://img.shields.io/github/stars/mrhujaifa/flamekit?style=flat-square&color=yellow)](https://github.com/mrhujaifa/flamekit/stargazers)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg?style=flat-square)](LICENSE)
[![Go version](https://img.shields.io/github/go-mod/go-version/mrhujaifa/flamekit?style=flat-square)](go.mod)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square)](https://github.com/mrhujaifa/flamekit/pulls)

![demo](demo.gif)

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

### `flamekit analyze`
Scan your entire codebase and display a flame map — files ranked by risk level.

```bash
flamekit analyze
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

**Risk Levels:**
| Level | Flame Score | Meaning |
|-------|-------------|---------|
| HIGH | ≥ 10 | Frequently changed, historically bug-prone. Act now. |
| MED | ≥ 4 | Moderate risk. Review carefully before touching. |
| LOW | < 4 | Stable. No immediate action needed. |

---

### `flamekit suggest`
Get a prioritized list of files to refactor — ranked by maximum impact.

```bash
# Top 5 suggestions (default)
flamekit suggest

# Custom limit
flamekit suggest --limit 10
flamekit suggest --limit 3
```

```
  flamekit — Refactoring Suggestions
  Files ranked by impact — fix these first

  ─────────────────────────────────────────────────────────────────
  #1  internal/app/app.go
      Flame: 45  Changes: 120  Bug Fixes: 23
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

---

## Roadmap

| Command | Status | Description |
|---------|--------|-------------|
| `flamekit analyze` | ✅ Available | Codebase flame map |
| `flamekit suggest` | ✅ Available | Refactoring priorities |
| `flamekit health` | 🔨 In Progress | Project health score 0-100 |
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
