package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

type ComplexFlowActor struct {
	c *Config
	o IOperation
	f IComplexFlow
}

func NewComplexFlowActor() *ComplexFlowActor {
	return &ComplexFlowActor{}
}

func (oa *ComplexFlowActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {

	case *actor.Started:
		resolveConfig(&oa.c, ctx)
		resolveInterface(&oa.o, func() IOperation { return NewOperationActorWrapper(&ctx) })
		resolveInterface(&oa.f, func() IComplexFlow { return NewComplexFlow(oa.o) })

		oa.f.Start()

		for i := 0; i < 100; i++ {

			oa.f.Handle(rand.Intn(10))
			time.Sleep(500 * time.Millisecond)
		}

	}

	fmt.Fprintf(os.Stderr, "%s received message of type %T and value %v\n", ctx.Self().Id, ctx.Message(), ctx.Message())
}
