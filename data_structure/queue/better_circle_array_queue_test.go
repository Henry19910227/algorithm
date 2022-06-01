package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBetterCircleArrayQueue_IsEmpty(t *testing.T) {
	queue := NewBetterCircleArrayQueue(3)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	assert.Equal(t, false, queue.IsEmpty())
	queue.Pop()
	queue.Pop()
	queue.Pop()
	assert.Equal(t, true, queue.IsEmpty())
}

func TestBetterCircleArrayQueue_IsFull(t *testing.T) {
	queue := NewBetterCircleArrayQueue(3)
	assert.NoError(t, queue.Push(1))
	assert.NoError(t, queue.Push(2))
	assert.NoError(t, queue.Push(3))
	assert.Error(t, queue.Push(4))
	_, err := queue.Pop()
	assert.NoError(t, err)
	_, err = queue.Pop()
	assert.NoError(t, err)
	assert.NoError(t, queue.Push(4))
	assert.NoError(t, queue.Push(5))
	assert.Error(t, queue.Push(6))
}

func TestBetterCircleArrayQueue_Head(t *testing.T) {
	queue := NewBetterCircleArrayQueue(3)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	value, _ := queue.Head()
	assert.Equal(t, 1, value)
	queue.Pop()
	queue.Pop()
	value, _ = queue.Head()
	assert.Equal(t, 3, value)
}

func TestBetterCircleArrayQueue_Size(t *testing.T) {
	queue := NewBetterCircleArrayQueue(3)
	queue.Push(1)
	queue.Push(2)
	queue.Push(3)
	assert.Equal(t, 3, queue.Size())
	queue.Pop()
	queue.Pop()
	assert.Equal(t, 1, queue.Size())
	queue.Push(4)
	queue.Push(5)
	assert.Equal(t, 3, queue.Size())
}
