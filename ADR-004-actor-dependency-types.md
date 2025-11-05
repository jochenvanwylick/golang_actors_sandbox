# ADR-004: Actor Dependency Types - Structs vs Interfaces

## Status

Accepted

## Context

Actors in our system need to resolve dependencies at runtime. We must decide how to represent different types of dependencies. Some dependencies contain data (like configuration), while others represent behavior (like operations). The choice between structs and interfaces affects how dependencies are resolved, tested, and used.

## Decision

Actor dependencies follow this pattern:

- **Data dependencies are structs**: Dependencies that primarily contain data (e.g., `Config`) are represented as structs
- **Behavior dependencies are interfaces**: Dependencies that perform actions or represent behavior (e.g., `IOperation`, `IComplexFlow`) are represented as interfaces

This creates a clear semantic distinction:
- Structs are resolved via actor messages (request-response pattern)
- Interfaces are resolved via factory functions that return implementations

Examples:
- `Config` (struct) - resolved via `resolveConfig()` sending `GetConfig` message
- `IOperation` (interface) - resolved via `resolveInterface()` calling a factory function
- `IComplexFlow` (interface) - resolved via factory function that creates implementation

## Consequences

### Positive

- **Clear Semantics**: The type signals intent - structs are data, interfaces are behavior
- **Testability**: Interfaces allow easy mocking with test doubles
- **Flexibility**: Interface implementations can be swapped without changing callers
- **Type Safety**: Go's type system enforces the distinction

### Negative

- **Consistency**: Different resolution mechanisms for structs vs interfaces
- **Learning Curve**: Developers must understand when to use each pattern
- **Reflection Complexity**: Current implementation uses reflection for both types

## References

- `config.go` - Config struct example
- `operation.go` - IOperation interface example
- `complex_flow.go` - IComplexFlow interface example
- `util.go` - `resolveConfig()` and `resolveInterface()` functions
- "A Philosophy of Software Design" - Chapter 6: General-Purpose Modules Are Deeper
- Effective Go - Interfaces section

