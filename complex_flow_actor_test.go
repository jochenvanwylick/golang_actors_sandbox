package main

import (
	"testing"
	"time"

	"github.com/asynkron/protoactor-go/actor"
	"github.com/stretchr/testify/assert"
)

func TestComplexFlowActorInIsolation(t *testing.T) {
	system := actor.NewActorSystem()
	defer system.Shutdown()

	complexFlowActor := NewComplexFlowActor()

	complexFlowActor.c = NewConfig()
	complexFlowActor.o = NewFakeOperation()
	fakeFlow := NewFakeComplexFlow()
	complexFlowActor.f = fakeFlow

	props := actor.PropsFromProducer(func() actor.Actor { return complexFlowActor })
	system.Root.Spawn(props)
	time.Sleep(1 * time.Second)

	assert.True(t, fakeFlow.startCalled)
}

func TestComplexFlowActorWithContext(t *testing.T) {
	system := actor.NewActorSystem()
	defer system.Shutdown()

	cfgProps := actor.PropsFromProducer(func() actor.Actor { return NewConfigActor() })
	system.Root.SpawnNamed(cfgProps, string(ActorConfig))

	props := actor.PropsFromProducer(func() actor.Actor { return NewComplexFlowActor() })
	system.Root.Spawn(props)
	time.Sleep(1 * time.Second)
}

// Fake
type FakeComplexFlow struct {
	startCalled bool
}

func NewFakeComplexFlow() *FakeComplexFlow {
	return &FakeComplexFlow{}
}

func (f *FakeComplexFlow) GetCurrentStep() Step {
	return Step{Name: "Fake"}
}

func (f *FakeComplexFlow) GetComplexFlowResult() ComplexFlowResult {
	return ComplexFlowResult{success: true}
}

func (f *FakeComplexFlow) Handle(input int) {
}

func (f *FakeComplexFlow) Start() {
	f.startCalled = true
}
