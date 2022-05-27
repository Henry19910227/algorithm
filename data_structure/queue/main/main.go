package main

import "algorithm/data_structure/queue"

func main() {
	arrayQueue := queue.NewArrayQueue(10)
	arrayQueue.Push(1)
	arrayQueue.Push(2)
	arrayQueue.Push(3)
	arrayQueue.Show()
	arrayQueue.Pop()
	arrayQueue.Show()
}
