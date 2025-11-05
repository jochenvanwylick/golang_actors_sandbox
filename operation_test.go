package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOperation(t *testing.T) {
	operation := NewOperation()
	result := operation.DoWork(10)
	assert.True(t, result.Success)
}

func TestOperationCancel(t *testing.T) {
	operation := NewOperation()
	operation.Cancel()
	result := operation.DoWork(10)
	time.Sleep(100 * time.Millisecond)
	assert.False(t, result.Success)
}

func TestOperationImplementsIOperation(t *testing.T) {
	var _ IOperation = NewOperation()
}
