package link_list

type node struct {
	id   int64
	next Node
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
