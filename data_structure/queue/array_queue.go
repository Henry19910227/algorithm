package queue

import (
	"errors"
	"fmt"
)

type arrayQueue struct {
	maxSize int //queue的最大容量
	front   int //queue頭部(指向第一筆資料的前一個位置)
	rear    int //queue尾部(指向最後一筆資料的當前位置)
	values  []int
}

func NewArrayQueue(size int) ArrayQueue {
	return &arrayQueue{maxSize: size, front: -1, rear: -1, values: make([]int, size)}
}

func (a *arrayQueue) IsFull() bool {
	return a.rear == a.maxSize-1
}

func (a *arrayQueue) IsEmpty() bool {
	return a.rear == a.front
}

func (a *arrayQueue) Push(value int) error {
	if a.IsFull() {
		return errors.New("queue is full")
	}
	a.rear++
	a.values[a.rear] = value
	return nil
}

func (a *arrayQueue) Pop() (int, error) {
	if a.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	a.front++
	return a.values[a.front], nil
}

func (a *arrayQueue) Show() {
	for i := a.front + 1; i <= a.rear; i++ {
		fmt.Println(a.values[i])
	}
}
