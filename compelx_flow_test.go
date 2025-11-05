package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestComplexFlowWithFake(t *testing.T) {
	fake := NewFakeOperation()

	flow := NewComplexFlow(fake)
	assert.Equal(t, flow.GetCurrentStep().Name, "Start")
	flow.Start()
	flow.Handle(4)

	assert.Equal(t, flow.GetCurrentStep().Name, "Working")
	assert.True(t, fake.doWorkCalled)
}

func TestComplexFlowWithRealOperation(t *testing.T) {
	operation := NewOperation()

	flow := NewComplexFlow(operation)
	assert.Equal(t, flow.GetCurrentStep().Name, "Start")
	flow.Start()
	flow.Handle(4)
	assert.Equal(t, flow.GetCurrentStep().Name, "Working")
}

// Fake
type FakeOperation struct {
	doWorkCalled bool
}

func NewFakeOperation() *FakeOperation {
	return &FakeOperation{doWorkCalled: false}
}

func (f *FakeOperation) DoWork(countTo int) OperationResult {
	f.doWorkCalled = true
	return OperationResult{Success: true}
}

func (f *FakeOperation) Cancel() {
}
