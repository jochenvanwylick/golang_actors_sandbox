package main

import (
	"fmt"
	"os"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

type OperationActor struct {
	o *Operation
}

func NewOperationActor() *OperationActor {
	return &OperationActor{o: NewOperation()}
}

func (oa *OperationActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {
	case *actor.Started:
		fmt.Fprintf(os.Stderr, "OperationActor started\n")
	case DoWorkRequest:
		result := oa.o.DoWork(msg.CountTo)
		ctx.Respond(result)
	case CancelRequest:
		oa.o.Cancel()
	}
	fmt.Fprintf(os.Stderr, "%s received message of type %T and value %v\n", ctx.Self().Id, ctx.Message(), ctx.Message())
}

// Message types for communication with OperationActor
type DoWorkRequest struct {
	CountTo int
}

type CancelRequest struct{}

// NewOperationActorWrapper creates a new wrapper that spawns an OperationActor and uses it via messages
func NewOperationActorWrapper(ctx *actor.Context) *OperationActorWrapper {
	pid, err := (*ctx).SpawnNamed(actor.PropsFromProducer(func() actor.Actor { return NewOperationActor() }), "operation-actor-wrapper")
	if err != nil {
		panic(fmt.Sprintf("Failed to spawn OperationActor: %v", err))
	}
	return &OperationActorWrapper{
		pid: pid,
		ctx: ctx,
	}
}

func (w *OperationActorWrapper) DoWork(countTo int) OperationResult {
	future := (*w.ctx).RequestFuture(w.pid, DoWorkRequest{CountTo: countTo}, 10*time.Second)
	result, err := future.Result()
	if err != nil {
		return OperationResult{Success: false, Error: err}
	}
	return result.(OperationResult)
}

func (w *OperationActorWrapper) Cancel() {
	(*w.ctx).Send(w.pid, CancelRequest{})
}
