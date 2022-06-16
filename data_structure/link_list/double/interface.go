package double

type Node interface {
	ID() int64
	SetNext(node Node)
	SetPrevious(node Node)
	Next() Node
	Previous() Node
}

type LinkList interface {
	Head() Node
	SetHead(node Node)
	Add(node Node)
	List() []int64
	Get(index int) (int64, error)
	Insert(index int, node Node)
	Delete(id int64) error
	Len() int
}
