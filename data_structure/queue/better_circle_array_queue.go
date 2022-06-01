package queue

import "errors"

type betterCircleArrayQueue struct {
	isIncome bool
	maxSize  int
	front    int
	rear     int
	values   []int
}

func NewBetterCircleArrayQueue(size int) Queue {
	return &betterCircleArrayQueue{isIncome: false, maxSize: size, front: 0, rear: 0, values: make([]int, size)}
}

func (b *betterCircleArrayQueue) IsFull() bool {
	return (b.rear == b.front) && b.isIncome
}

func (b *betterCircleArrayQueue) IsEmpty() bool {
	return (b.rear == b.front) && !b.isIncome
}

func (b *betterCircleArrayQueue) Push(value int) error {
	if b.IsFull() {
		return errors.New("queue is full")
	}
	b.isIncome = true
	b.values[b.rear] = value
	b.rear = (b.rear + 1) % b.maxSize
	return nil
}

func (b *betterCircleArrayQueue) Pop() (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	b.isIncome = false
	value := b.values[b.rear]
	b.front = (b.front + 1) % b.maxSize
	return value, nil
}

func (b *betterCircleArrayQueue) Head() (int, error) {
	if b.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return b.values[b.front], nil
}

func (b *betterCircleArrayQueue) Size() int {
	size := (b.rear + b.maxSize - b.front) % b.maxSize
	if b.isIncome {
		return b.maxSize
	}
	return size
}

func (b *betterCircleArrayQueue) Show() {
	//TODO implement me
	panic("implement me")
}
