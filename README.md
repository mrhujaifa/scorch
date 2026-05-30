# Scorch ⚡

> Find where your codebase burns — Git history analyzer for developers

![demo](demo.gif)

Scorch analyzes your git history to find the most dangerous files
in your codebase. Files that change frequently and cause the most
bugs — ranked by risk.

## Why Scorch

| Tool       | Cost       | Setup    | CLI |
| ---------- | ---------- | -------- | --- |
| CodeScene  | $300/month | Complex  | ❌  |
| SonarQube  | Free/Paid  | Server   | ❌  |
| **Scorch** | **Free**   | **Zero** | ✅  |

## Installation

```bash
go install github.com/mrhujaifa/scorch@latest
```

## Usage

```bash
# Analyze your codebase
scorch analyze

# Get refactoring suggestions
scorch suggest

# Top 3 suggestions
scorch suggest --limit 3
```

## Roadmap

- [x] `scorch analyze` — codebase flame map
- [x] `scorch suggest` — refactoring priorities
- [ ] `scorch health` — project health score
- [ ] `scorch file` — single file deep dive
- [ ] `scorch impact` — blast radius analysis
- [ ] `scorch predict` — bug prediction

## License

MIT © Md. Hujaifa Islam Shanto
