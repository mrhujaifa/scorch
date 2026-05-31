# Contributing to Flamekit

Thanks for your interest in improving Flamekit! This document explains how to contribute, report issues, and submit changes.

## How to Contribute

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature-name
   ```

````
3. Make your changes.
4. Run tests and linters:
   ```bash
go test ./...
go vet ./...
````

5. Commit with a clear message.
6. Open a pull request with a short description of the change.

## Issue Reporting

When opening an issue, please include:

- A clear problem statement
- Steps to reproduce
- Expected behavior
- Actual behavior
- Relevant logs or output

## Pull Request Guidelines

- Keep PRs focused and small.
- Describe the problem and your solution clearly.
- Include tests for bug fixes and new features when possible.
- Ensure all existing tests pass.
- Avoid committing generated files or build artifacts.

## Code Style

- Follow Go conventions.
- Use `gofmt` on all changed Go files.
- Keep code readable and consistent.

## Security and Sensitive Data

- Do not commit secrets, credentials, or private keys.
- If your change affects security, mention it in the PR description.
- Use `SECURITY.md` for vulnerability reporting guidelines.

## Development Workflow

- `flamekit analyze` and `flamekit suggest` are the current working commands.
- New commands should follow the existing cobra-based CLI structure.
- Add documentation for any new feature in `README.md`.

## Thank You

Contributions make Flamekit better for everyone. Your fixes, ideas, and reviews are appreciated.
