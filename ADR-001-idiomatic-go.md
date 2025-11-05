# ADR-001: Idiomatic Go Guidelines

## Status

Accepted

## Context

As a Go codebase, we need to establish consistent coding standards that align with the Go community's best practices. This ensures code readability, maintainability, and enables effective collaboration. The codebase should be immediately recognizable to Go developers and follow established conventions.

## Decision

We will follow idiomatic Go guidelines throughout the codebase, adhering to:

- **Effective Go** principles: clear naming, simple interfaces, error handling patterns
- **Go Code Review Comments**: standard formatting, naming conventions, and structural patterns
- **Go Best Practices**: package organization, interface design, and concurrency patterns

Key principles include:
- Use interfaces for behavior, structs for data
- Return errors explicitly, don't use exceptions
- Prefer composition over inheritance
- Keep functions small and focused
- Use meaningful variable and function names
- Follow standard package organization
- Write clear, concise code over clever solutions

## Consequences

### Positive

- **Readability**: Code follows familiar patterns that Go developers expect
- **Maintainability**: Consistent style reduces cognitive load when reading code
- **Community Alignment**: New team members familiar with Go can contribute immediately
- **Tooling Support**: Standard patterns work well with Go tooling (gofmt, go vet, etc.)

### Negative

- **Initial Learning Curve**: Developers unfamiliar with Go idioms may need guidance
- **Rigidity**: Some patterns may feel restrictive compared to other languages
- **Enforcement Overhead**: Requires code review vigilance to maintain standards

## References

- [Effective Go](https://go.dev/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go Best Practices](https://golang.org/doc/effective_go)
- "A Philosophy of Software Design" - Chapter 3: Working Code Isn't Enough

