package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArrayQueue_IsEmpty(t *testing.T) {
	queue := NewArrayQueue(3)
	assert.Equal(t, true, queue.IsEmpty())
	queue.Push(1)
	assert.Equal(t, false, queue.IsEmpty())
	queue.Push(2)
	assert.Equal(t, false, queue.IsEmpty())
	queue.Push(3)
	assert.Equal(t, false, queue.IsEmpty())
	queue.Pop()
	assert.Equal(t, false, queue.IsEmpty())
	queue.Pop()
	assert.Equal(t, false, queue.IsEmpty())
	queue.Pop()
	assert.Equal(t, true, queue.IsEmpty())
}

func TestArrayQueue_IsFull(t *testing.T) {
	queue := NewArrayQueue(3)
	assert.Equal(t, false, queue.IsFull())
	queue.Push(1)
	assert.Equal(t, false, queue.IsFull())
	queue.Push(2)
	assert.Equal(t, false, queue.IsFull())
	queue.Push(3)
	assert.Equal(t, true, queue.IsFull())
}

func TestArrayQueue_Head(t *testing.T) {
	queue := NewArrayQueue(2)
	_, err := queue.Head()
	assert.EqualError(t, err, "queue is empty")
	queue.Push(1)
	queue.Push(2)
	value, _ := queue.Head()
	assert.Equal(t, 1, value)
}

func TestArrayQueue_Size(t *testing.T) {
	queue := NewArrayQueue(3)
	assert.Equal(t, 0, queue.Size())
	queue.Push(1)
	assert.Equal(t, 1, queue.Size())
	queue.Push(2)
	assert.Equal(t, 2, queue.Size())
	queue.Push(3)
	assert.Equal(t, 3, queue.Size())
	queue.Pop()
	assert.Equal(t, 2, queue.Size())
	queue.Pop()
	assert.Equal(t, 1, queue.Size())
	queue.Pop()
	assert.Equal(t, 0, queue.Size())
}
