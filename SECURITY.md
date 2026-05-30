# Security Policy

## Reporting a Vulnerability

If you find a security issue in Scorch, please report it responsibly:

1. Open an issue with the label `security` and mark it as private if the platform supports it.
2. If you prefer private communication, send a message to the project maintainer via the repository contact or email.

Do not post security issues publicly until a fix is available.

## Supported Versions

We actively support the latest released version of Scorch and the previous minor release. Security fixes are prioritized for these supported versions.

## Response Process

- Acknowledge receipt within 48 hours.
- Investigate the issue and communicate expected timelines.
- Publish a fix or mitigation as soon as possible.
- Notify affected users once the fix is ready.

## Security Practices

- All code changes should pass `go test ./...` and `go vet ./...` before release.
- Dependencies should be checked using `govulncheck ./...` or equivalent tooling.
- Secrets and credentials must never be committed to the repository.

## Reporting Guidelines

Please include:

- A clear description of the issue
- Steps to reproduce
- Expected and actual behavior
- Affected versions
- Any suggested mitigation
