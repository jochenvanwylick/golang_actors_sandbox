# ADR-002: Proto-Actor and Vanilla Go Hybrid Architecture

## Status

Accepted

## Context

We need to balance two competing concerns: keeping core business logic testable and independent, while leveraging the actor model for concurrency and message-passing. The core library should not be tightly coupled to the actor framework, but it also doesn't need to be a standalone service that manages its own network connections.

## Decision

We will use a hybrid architecture where:

1. **Core library** (`operation.go`, `flow.go`, `complex_flow.go`) is written in vanilla Go with no proto-actor dependencies
2. **Actor layer** (`*_actor.go` files) orchestrates and coordinates the core library using the actor model
3. **Core library does not run standalone**: It does not listen for TCP messages or manage its own network layer

The core library contains pure business logic that can be tested independently. The actor layer provides the orchestration, message-passing, and concurrency management. This separation allows us to:
- Test core logic without actor system complexity
- Swap actor implementations if needed
- Maintain clear boundaries between business logic and infrastructure

## Consequences

### Positive

- **Testability**: Core logic can be unit tested without actor system overhead
- **Separation of Concerns**: Business logic is independent of infrastructure
- **Flexibility**: Actor layer can be modified or replaced without touching core logic
- **Clean Architecture**: Aligns with dependency inversion - core doesn't depend on framework

### Negative

- **Code Duplication Risk**: Some patterns may need to be duplicated in both layers
- **Coordination Complexity**: Need to carefully manage the boundary between layers
- **Not Fully Generic**: Core library is designed for actor orchestration, limiting reusability in other contexts

## References

- `operation.go` - Vanilla Go implementation
- `flow.go` - Vanilla Go implementation
- `complex_flow.go` - Vanilla Go implementation
- `operation_actor.go` - Actor orchestration layer
- `flow_actor.go` - Actor orchestration layer
- Clean Architecture by Robert C. Martin
- "A Philosophy of Software Design" - Chapter 4: Modules Should Be Deep

