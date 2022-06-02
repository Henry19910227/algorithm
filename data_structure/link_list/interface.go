package link_list

type Node interface {
	ID() int64
	SetNext(node Node)
	Next() Node
}

type LinkList interface {
	Add(node Node)
}
