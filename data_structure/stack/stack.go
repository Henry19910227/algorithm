package stack

import "errors"

type stack struct {
	values  []int
	top     int
	maxSize int
}

func New(maxSize int) Stack {
	return &stack{values: make([]int, maxSize), top: -1, maxSize: maxSize}
}

func (s *stack) Push(value int) error {
	if s.isFull() {
		return errors.New("stack is full")
	}
	s.top++
	s.maxSize = value
	s.values[s.top] = value
	return nil
}

func (s *stack) Pop() (value int, err error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	value = s.values[s.top]
	s.top--
	return value, nil
}

func (s *stack) List() []int {
	values := make([]int, 0)
	for _, value := range s.values {
		values = append(values, value)
	}
	return values
}

func (s *stack) isFull() bool {
	return s.top == s.maxSize-1
}

func (s *stack) isEmpty() bool {
	return s.top == -1
}
