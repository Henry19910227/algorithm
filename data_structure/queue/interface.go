package queue

type Queue interface {
	// IsFull 判斷queue是否滿了
	IsFull() bool
	// IsEmpty 判斷queue是否為空
	IsEmpty() bool
	// Push 從queue尾部存入數據
	Push(value int) error
	// Pop 從queue頭部取出數據
	Pop() (int, error)
	// Head 取得 head value
	Head() (int, error)
	// Size 當前個數
	Size() int
	// Show 印出整個queue
	Show()
}
