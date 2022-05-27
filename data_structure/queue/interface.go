package queue

type ArrayQueue interface {
	// IsFull 判斷queue是否滿了
	IsFull() bool
	// IsEmpty 判斷queue是否為空
	IsEmpty() bool
	// Push 從queue尾部存入數據
	Push(value int) error
	// Pop 從queue頭部取出數據
	Pop() (int, error)
	// Show 印出整個queue
	Show()
}
