package main

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFlow(t *testing.T) {

	flow := NewFlow()
	flow.Start()

	time.Sleep(1100 * time.Millisecond)
	assert.Equal(t, flow.GetCurrentStep().Name, "Done")
}

func TestFlowWithFastAutoTransition(t *testing.T) {
	flow := NewFlow(WithAutoTransitionTimeout(100 * time.Millisecond))
	flow.Start()

	time.Sleep(110 * time.Millisecond)
	assert.Equal(t, flow.GetCurrentStep().Name, "Done")
}

func TestFlowRandomTransitions(t *testing.T) {
	flow := NewFlow()
	flow.Start()
	// for 10 times - handle a random number
	for i := 0; i < 100; i++ {
		flow.Handle(rand.Intn(10))
	}

	time.Sleep(1100 * time.Millisecond)
	assert.Equal(t, flow.GetCurrentStep().Name, "Done")
}

func TestFlowWithDoneCallback(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	flow := NewFlow(WithDoneCallback(func(s FlowResult) {
		assert.True(t, s.success)
		wg.Done()
	}))

	flow.Start()
	time.Sleep(1100 * time.Millisecond)
	wg.Wait()
}
