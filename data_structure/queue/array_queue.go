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

func NewArrayQueue(size int) Queue {
	return &arrayQueue{maxSize: size, front: 0, rear: 0, values: make([]int, size)}
}

func (a *arrayQueue) IsFull() bool {
	return a.rear == a.maxSize
}

func (a *arrayQueue) IsEmpty() bool {
	return a.rear == a.front
}

func (a *arrayQueue) Push(value int) error {
	if a.IsFull() {
		return errors.New("queue is full")
	}
	a.values[a.rear] = value
	a.rear++
	return nil
}

func (a *arrayQueue) Pop() (int, error) {
	if a.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	value := a.values[a.front]
	a.front++
	return value, nil
}

func (a *arrayQueue) Head() (int, error) {
	if a.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return a.values[a.front], nil
}

func (a *arrayQueue) Size() int {
	return a.rear - a.front
}

func (a *arrayQueue) Show() {
	for i := a.front + 1; i <= a.rear; i++ {
		fmt.Println(a.values[i])
	}
}
