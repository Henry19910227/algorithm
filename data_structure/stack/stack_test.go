package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	stack := New(3)
	assert.NoError(t, stack.Push(10))
	assert.NoError(t, stack.Push(20))
	assert.NoError(t, stack.Push(30))
	assert.Error(t, stack.Push(40))
}

func TestStack_Pop(t *testing.T) {
	stack := New(3)
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	_, err := stack.Pop()
	assert.NoError(t, err)
	_, err = stack.Pop()
	assert.NoError(t, err)
	_, err = stack.Pop()
	assert.NoError(t, err)
	_, err = stack.Pop()
	assert.Error(t, err)
}
