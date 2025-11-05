# Config Tech Spike

A Go prototype exploring actor-based flow management using ProtoActor.

## Overview

This project demonstrates an actor-based architecture for managing workflows with configurable steps, transitions, and callbacks. It uses the ProtoActor framework to implement actor patterns for flow management.

## Features

- Actor-based flow management (ConfigActor, FlowActor, ComplexFlowActor)
- Functional options pattern for configuration
- Step-based workflow with transitions
- Configurable callbacks and timeouts

## Requirements

- Go 1.25.2 or later

## Running

```bash
go run main.go
```

## Testing

```bash
go test ./...
```

## Architecture

The project includes Architecture Decision Records (ADRs) documenting key design decisions:
- ADR-001: Idiomatic Go Guidelines
- ADR-002: Hybrid Architecture
- ADR-003: DAS Integration
- ADR-004: Actor Dependency Types
- ADR-005: Functional Options Pattern

## Dependencies

- [ProtoActor Go](https://github.com/asynkron/protoactor-go) - Actor model framework
- [Testify](https://github.com/stretchr/testify) - Testing toolkit

