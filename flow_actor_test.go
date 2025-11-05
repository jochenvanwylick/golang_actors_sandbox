package main

import (
	"testing"
	"time"

	"github.com/asynkron/protoactor-go/actor"
)

func TestFlowActorInIsolation(t *testing.T) {
	system := actor.NewActorSystem()
	defer system.Shutdown()

	flowActor := NewFlowActor()
	flowActor.cfg = NewConfig()

	props := actor.PropsFromProducer(func() actor.Actor { return flowActor })
	system.Root.Spawn(props)
	time.Sleep(1 * time.Second)
}

func TestFlowActorWithContext(t *testing.T) {
	system := actor.NewActorSystem()
	defer system.Shutdown()

	cfgProps := actor.PropsFromProducer(func() actor.Actor { return NewConfigActor() })
	system.Root.SpawnNamed(cfgProps, string(ActorConfig))

	props := actor.PropsFromProducer(func() actor.Actor { return NewFlowActor() })
	system.Root.Spawn(props)

	time.Sleep(1 * time.Second)
}
