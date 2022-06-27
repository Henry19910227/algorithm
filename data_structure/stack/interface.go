package stack

type Stack interface {
	Push(value int) error
	Pop() (value int, err error)
	List() []int
}
