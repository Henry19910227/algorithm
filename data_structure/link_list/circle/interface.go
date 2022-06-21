package circle

type LinkList interface {
	Add(node Node)
	List() []int64
	Delete(id int64) error
}

type Node interface {
	ID() int64
	SetNext(node Node)
	Next() Node
}
