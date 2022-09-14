package binary_search_tree

type node struct {
	id    int64
	name  string
	left  Node
	right Node
}

func NewNode(id int64, name string) Node {
	return &node{id: id, name: name}
}

func (n *node) ID() int64 {
	return n.id
}

func (n *node) Name() string {
	return n.name
}

func (n *node) Left() Node {
	return n.left
}

func (n *node) Right() Node {
	return n.right
}

func (n *node) SetLeft(node Node) {
	n.left = node
}

func (n *node) SetRight(node Node) {
	n.right = node
}
