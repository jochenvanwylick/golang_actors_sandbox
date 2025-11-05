# ADR-003: DAS Trading Platform Integration

## Status

Accepted

## Context

The core library needs to integrate with DAS (Direct Access System) trading platforms. We must decide whether to keep the core library generic and abstract DAS-specific concepts, or allow DAS knowledge to permeate the core library. A fully generic approach would require extensive abstraction layers, while a DAS-specific approach trades reusability for simplicity and clarity.

## Decision

We will allow DAS trading platform knowledge to leak into the core library. The core library is purpose-built for DAS integration and is not intended to be a generic trading library.

This means:
- DAS-specific concepts, types, and behaviors can be used directly in the core library
- No need for abstract interfaces that hide DAS specifics
- The library is optimized for DAS workflows rather than generic trading operations

## Consequences

### Positive

- **Simplicity**: No unnecessary abstraction layers reduce complexity
- **Clarity**: Code directly expresses DAS concepts without indirection
- **Performance**: Direct integration avoids abstraction overhead
- **Faster Development**: Can use DAS-specific features without abstraction work

### Negative

- **Limited Reusability**: Core library cannot be easily adapted for other trading platforms
- **Tight Coupling**: Changes to DAS may require core library changes
- **Domain Lock-in**: The library is tightly bound to DAS ecosystem

## References

- "A Philosophy of Software Design" - Chapter 2: The Nature of Complexity (discusses strategic vs tactical complexity)
- Clean Architecture - Chapter 22: The Clean Architecture (discusses when to allow framework coupling)

