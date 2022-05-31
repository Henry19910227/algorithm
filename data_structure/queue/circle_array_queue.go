package queue

import "errors"

//http://epaper.gotop.com.tw/pdf/AEE032400.pdf
//為了解決 rear == front 時無法分辨是空佇列還是滿佇列的問題
//因此必須浪費一個空格來判斷是空的還是滿的，rear到達front前一位時就代表滿位

type circleArrayQueue struct {
	maxSize int
	front   int
	rear    int
	values  []int
}

func NewCircleArrayQueue(size int) Queue {
	return &circleArrayQueue{maxSize: size + 1, front: 0, rear: 0, values: make([]int, size+1)}
}

func (c *circleArrayQueue) IsFull() bool {
	return (c.rear+1)%c.maxSize == c.front
}

func (c *circleArrayQueue) IsEmpty() bool {
	return c.rear == c.front
}

func (c *circleArrayQueue) Push(value int) error {
	if c.IsFull() {
		return errors.New("queue is full")
	}
	c.values[c.rear] = value
	c.rear = (c.rear + 1) % c.maxSize
	return nil
}

func (c *circleArrayQueue) Pop() (int, error) {
	if c.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	value := c.values[c.front]
	c.front = (c.front + 1) % c.maxSize
	return value, nil
}

func (c *circleArrayQueue) Head() (int, error) {
	if c.IsEmpty() {
		return 0, errors.New("queue is empty")
	}
	return c.values[c.front], nil
}

func (c *circleArrayQueue) Size() int {
	return (c.rear + c.maxSize - c.front) % c.maxSize
}

func (c *circleArrayQueue) Show() {
	//TODO implement me
	panic("implement me")
}
