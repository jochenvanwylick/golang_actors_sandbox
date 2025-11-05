package main

import (
	"fmt"
	"os"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

type IOperation interface {
	DoWork(countTo int) OperationResult
	Cancel()
}

type Operation struct {
	cancel chan struct{}
}

func NewOperation() *Operation {
	fmt.Fprintf(os.Stderr, "NewOperation started\n")
	return &Operation{
		cancel: make(chan struct{}),
	}
}

func (o *Operation) DoWork(countTo int) OperationResult {
	fmt.Fprintf(os.Stderr, "Operation started - counting to %d\n", countTo)
	for i := 0; i < countTo; i++ {
		select {
		case <-o.cancel:
			return OperationResult{Success: false, Error: nil}
		case <-time.After(100 * time.Millisecond):
			// Continue counting
		}
	}
	return OperationResult{Success: true}
}

func (o *Operation) Cancel() {
	fmt.Fprintf(os.Stderr, "Operation cancelled\n")
	select {
	case <-o.cancel:
		// Already cancelled
	default:
		close(o.cancel)
	}
}

type OperationResult struct {
	Success bool
	Error   error
}

// OperationActorWrapper implements IOperation by delegating to an OperationActor via messages
type OperationActorWrapper struct {
	pid *actor.PID
	ctx *actor.Context
}
