package main

import (
	"fmt"
	"os"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

type FlowActor struct {
	f   *Flow
	cfg *Config
}

func NewFlowActor() *FlowActor {
	return &FlowActor{f: NewFlow()}
}

func (fa *FlowActor) Receive(ctx actor.Context) {
	switch msg := ctx.Message().(type) {

	case *actor.Started:
		resolveConfig(&fa.cfg, ctx)

		ConfigureFlow(fa.f,
			WithConfig(fa.cfg),
			WithAutoTransitionTimeout(100*time.Millisecond),
			WithDoneCallback(func(s FlowResult) {
				ctx.Send(ctx.Self(), FlowDoneMsg{
					Result: s,
				})
			}))

	case FlowMsg:
		fa.f.Handle(msg.Input)

	case FlowDoneMsg:
		fmt.Fprintf(os.Stderr, "Flow done: %v\n", msg.Result)
	}
	fmt.Fprintf(os.Stderr, "%s received message of type %T and value %v\n", ctx.Self().Id, ctx.Message(), ctx.Message())
}

type FlowMsg struct {
	Input int
}

type FlowDoneMsg struct {
	Result FlowResult
}
