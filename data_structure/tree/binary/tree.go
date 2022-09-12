package binary

type tree struct {
	root Node
}

func NewTree(root Node) Tree {
	return &tree{root: root}
}

func (t *tree) Root() Node {
	return t.root
}

func (t *tree) PreOrder() []Node {
	list := make([]Node, 0)
	t.preOrder(&list, t.root)
	return list
}

func (t *tree) InOrder() []Node {
	list := make([]Node, 0)
	t.inOrder(&list, t.root)
	return list
}

func (t *tree) PostOrder() []Node {
	list := make([]Node, 0)
	t.postOrder(&list, t.root)
	return list
}

func (t *tree) preOrder(list *[]Node, node Node) {
	if node == nil {
		return
	}
	*list = append(*list, node)
	t.preOrder(list, node.Left())
	t.preOrder(list, node.Right())
}

func (t *tree) inOrder(list *[]Node, node Node) {
	if node == nil {
		return
	}
	t.inOrder(list, node.Left())
	*list = append(*list, node)
	t.inOrder(list, node.Right())
}

func (t *tree) postOrder(list *[]Node, node Node) {
	if node == nil {
		return
	}
	t.postOrder(list, node.Left())
	t.postOrder(list, node.Right())
	*list = append(*list, node)
}
