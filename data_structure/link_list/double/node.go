package double

type node struct {
	id       int64
	next     Node
	previous Node
}

func NewNode(id int64) Node {
	return &node{id: id}
}

func (n *node) ID() int64 {
	return n.id
}

func (n *node) SetNext(node Node) {
	n.next = node
}

func (n *node) Next() Node {
	return n.next
}

func (n *node) SetPrevious(node Node) {
	n.previous = node
}

func (n *node) Previous() Node {
	return n.previous
}
