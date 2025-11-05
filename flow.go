package main

import (
	"fmt"
	"os"
	"time"
)

type Flow struct {
	tansitionNumber       int
	steps                 []Step
	currentStepIndex      int
	autoTransitionTimeout time.Duration
	doneCallback          func(FlowResult)
	cfg                   *Config
}

type FlowResult struct {
	success bool
}

type Step struct {
	Name string
}

// FlowOption configures a Flow instance.
type FlowOption func(*Flow)

// WithAutoTransitionTimeout overrides the default auto transition timeout.
func WithAutoTransitionTimeout(d time.Duration) FlowOption {
	return func(f *Flow) {
		f.autoTransitionTimeout = d
	}
}

func WithDoneCallback(callback func(FlowResult)) FlowOption {
	return func(f *Flow) {
		f.doneCallback = callback
	}
}

func WithConfig(c *Config) FlowOption {
	return func(f *Flow) {
		f.cfg = c
	}
}

func NewFlow(opts ...FlowOption) *Flow {
	steps := []Step{
		{Name: "Start"},
		{Name: "Working"},
		{Name: "Done"},
	}
	transitionNumber := 4
	// Default to 10 seconds; can be overridden via options (e.g., tests use 10ms)
	autoTransitionTimeout := 1000 * time.Millisecond

	f := &Flow{
		steps:                 steps,
		tansitionNumber:       transitionNumber,
		currentStepIndex:      0,
		autoTransitionTimeout: autoTransitionTimeout,
	}
	ConfigureFlow(f, opts...)
	return f
}

func ConfigureFlow(f *Flow, opts ...FlowOption) {
	for _, opt := range opts {
		opt(f)
	}
}

func (f *Flow) Start() {
	go func() {
		time.Sleep(f.autoTransitionTimeout)
		f.jumpToStep(2)
	}()
}

func (f *Flow) Handle(input int) {
	// if inuput == transitionNumber - go to next step
	if input == f.tansitionNumber && f.currentStepIndex < len(f.steps)-1 {
		fmt.Fprintf(os.Stderr, "Transitioning to next step - :D - from %s to %s\n", f.steps[f.currentStepIndex].Name, f.steps[f.currentStepIndex+1].Name)
		f.jumpToStep(f.currentStepIndex + 1)
	} else if input == f.tansitionNumber && f.currentStepIndex == len(f.steps)-1 {
		fmt.Fprintln(os.Stderr, "In last step ... no transition")
	} else {
		fmt.Fprintf(os.Stderr, "Wrong number ... no transition\n")
	}

}

func (f *Flow) GetCurrentStep() Step {
	return f.steps[f.currentStepIndex]
}

func (f *Flow) GetFlowResult() FlowResult {
	result := FlowResult{success: false}
	currentStep := f.GetCurrentStep()
	if currentStep.Name == "Done" {
		result.success = true
	}
	return result
}

func (f *Flow) jumpToStep(stepIndex int) {
	fmt.Fprintf(os.Stderr, "Auto-transitioning to step: %s\n", f.steps[stepIndex].Name)
	f.currentStepIndex = stepIndex

	if f.currentStepIndex == len(f.steps)-1 && f.doneCallback != nil {
		f.doneCallback(f.GetFlowResult())
	}
}
