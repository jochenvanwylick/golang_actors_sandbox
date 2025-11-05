# ADR-005: Functional Options Pattern for Configuration

## Status

Accepted

## Context

We need a flexible way to configure structs during initialization. Traditional approaches like constructor overloading or large parameter lists become unwieldy as the number of configuration options grows. We also need sensible defaults that work for most cases, but allow overrides for special cases or testing.

## Decision

We will use the **Functional Options Pattern** (also known as the Options Pattern) for configuring structs. This pattern follows the idiomatic Go approach established by Rob Pike and Dave Cheney.

Key principles:
1. **Sensible defaults**: All configuration options have reasonable defaults set in the constructor
2. **Optional overrides**: Options can be provided via variadic function arguments
3. **Extensibility**: New options can be added without breaking existing code
4. **Testing support**: Options enable easy test-specific overrides (e.g., shorter timeouts)

The pattern structure:
- Define an option type: `type FlowOption func(*Flow)`
- Provide option constructors: `func WithConfig(c *Config) FlowOption`
- Accept options in constructor: `func NewFlow(opts ...FlowOption) *Flow`
- Support post-construction configuration: `func ConfigureFlow(f *Flow, opts ...FlowOption)`

## Consequences

### Positive

- **Flexibility**: Options can be provided in any order, any subset
- **Extensibility**: New options don't break existing code or require constructor changes
- **Testability**: Easy to override defaults for faster tests (e.g., `WithAutoTransitionTimeout(100*time.Millisecond)`)
- **Readability**: Options are self-documenting (`WithConfig(cfg)`)
- **Backward Compatible**: Adding new options doesn't require changes to existing callers

### Negative

- **Slight Overhead**: Small function call overhead for each option
- **Learning Curve**: Developers unfamiliar with pattern may need guidance
- **Default Discovery**: Developers must read constructor to discover defaults

## Examples

```go
// Use defaults
flow := NewFlow()

// Override for testing
flow := NewFlow(WithAutoTransitionTimeout(100 * time.Millisecond))

// Configure in actor after initialization
ConfigureFlow(fa.f,
    WithConfig(fa.cfg),
    WithAutoTransitionTimeout(100*time.Millisecond),
    WithDoneCallback(func(s FlowResult) { ... }))
```

## References

- [Functional Options Pattern - Rob Pike](https://commandcenter.blogspot.com/2014/01/self-referential-functions-and-design.html)
- [Functional Options for Friendly APIs - Dave Cheney](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
- `flow.go` - `FlowOption`, `WithConfig`, `WithAutoTransitionTimeout`, `WithDoneCallback`
- `flow_actor.go:26-33` - Example usage in actor layer
- `flow_test.go:21-27` - Example usage in tests
- "A Philosophy of Software Design" - Chapter 6: General-Purpose Modules Are Deeper

