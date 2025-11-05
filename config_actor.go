package main

import (
	"fmt"
	"os"

	"github.com/asynkron/protoactor-go/actor"
)

type ConfigActor struct {
	c *Config
}

func NewConfigActor() *ConfigActor {
	return &ConfigActor{c: NewConfig()}
}

func (ca *ConfigActor) Receive(ctx actor.Context) {
	switch ctx.Message().(type) {
	case *actor.Started:
		fmt.Fprintf(os.Stdout, "ConfigActor started\n")
	case GetConfig:
		ctx.Respond(ca.c)
	}

	fmt.Fprintf(os.Stderr, "%s received message of type %T and value %v\n", ctx.Self().Id, ctx.Message(), ctx.Message())
}

type GetConfig struct {
}
