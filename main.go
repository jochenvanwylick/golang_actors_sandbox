package main

import (
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

func main() {
	sys := actor.NewActorSystem()
	defer sys.Shutdown()

	sys.Root.SpawnNamed(actor.PropsFromProducer(func() actor.Actor { return NewConfigActor() }), string(ActorConfig))
	sys.Root.SpawnNamed(actor.PropsFromProducer(func() actor.Actor { return NewFlowActor() }), string(ActorFlow))

	sys.Root.SpawnNamed(actor.PropsFromProducer(func() actor.Actor { return NewComplexFlowActor() }), "whatever")

	time.Sleep(10 * time.Second)
}
