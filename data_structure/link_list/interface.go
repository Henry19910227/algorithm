package link_list

type Node interface {
	ID() int64
	SetNext(node Node)
	Next() Node
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
