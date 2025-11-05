package main

import (
	"fmt"
	"os"
	"time"
)

type IComplexFlow interface {
	Start()
	Handle(input int)
	GetCurrentStep() Step
	GetComplexFlowResult() ComplexFlowResult
}

type ComplexFlow struct {
	tansitionNumber       int
	steps                 []Step
	currentStepIndex      int
	autoTransitionTimeout time.Duration
	doneCallback          func(ComplexFlowResult)
	cfg                   *Config
	operation             IOperation
}

type ComplexFlowResult struct {
	success bool
}

func NewComplexFlow(operation IOperation) *ComplexFlow {
	steps := []Step{
		{Name: "Start"},
		{Name: "Working"},
		{Name: "Done"},
	}
	transitionNumber := 4
	// Default to 10 seconds; can be overridden via options (e.g., tests use 10ms)
	autoTransitionTimeout := 1000 * time.Millisecond

	f := &ComplexFlow{
		steps:                 steps,
		tansitionNumber:       transitionNumber,
		currentStepIndex:      0,
		autoTransitionTimeout: autoTransitionTimeout,
		operation:             operation,
	}

	return f
}

func (f *ComplexFlow) Start() {
	go func() {
		f.operation.DoWork(10)
		time.Sleep(f.autoTransitionTimeout)
		f.jumpToStep(2)
	}()
}

func (f *ComplexFlow) Handle(input int) {
	// if inuput == transitionNumber - go to next step
	if input == f.tansitionNumber && f.currentStepIndex < len(f.steps)-1 {
		fmt.Fprintf(os.Stderr, "Transitioning to next step - :D - from %s to %s\n", f.steps[f.currentStepIndex].Name, f.steps[f.currentStepIndex+1].Name)
		f.operation.Cancel()
		f.operation.DoWork(input * 5)
		f.jumpToStep(f.currentStepIndex + 1)
	} else if input == f.tansitionNumber && f.currentStepIndex == len(f.steps)-1 {
		fmt.Fprintln(os.Stderr, "In last step ... no transition")
	} else {
		fmt.Fprintf(os.Stderr, "Wrong number ... no transition\n")
	}
}

func (f *ComplexFlow) GetCurrentStep() Step {
	return f.steps[f.currentStepIndex]
}

func (f *ComplexFlow) GetComplexFlowResult() ComplexFlowResult {
	result := ComplexFlowResult{success: false}
	currentStep := f.GetCurrentStep()
	if currentStep.Name == "Done" {
		result.success = true
	}
	return result
}

func (f *ComplexFlow) jumpToStep(stepIndex int) {
	fmt.Fprintf(os.Stderr, "Auto-transitioning to step: %s\n", f.steps[stepIndex].Name)
	f.currentStepIndex = stepIndex

	if f.currentStepIndex == len(f.steps)-1 && f.doneCallback != nil {
		f.doneCallback(f.GetComplexFlowResult())
	}
}
