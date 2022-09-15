package binary_search_tree

type tree struct {
	root Node
}

func NewTree(root Node) Tree {
	return &tree{root: root}
}

func (t *tree) Root() Node {
	return t.root
}

func (t *tree) Find(id int64) Node {
	return t.find(t.root, id)
}

func (t *tree) FindMin() Node {
	return t.findMin(t.root)
}

func (t *tree) FindMax() Node {
	return t.findMax(t.root)
}

func (t *tree) Insert(node Node) {
	t.insert(t.root, node)
}

func (t *tree) InOrder() []Node {
	list := make([]Node, 0)
	t.inOrder(&list, t.root)
	return list
}

func (t *tree) find(node Node, id int64) Node {
	if node == nil {
		return nil
	}
	if node.ID() > id {
		return t.find(node.Left(), id)
	}
	if node.ID() < id {
		return t.find(node.Right(), id)
	}
	return node
}

func (t *tree) findMin(node Node) Node {
	if node == nil {
		return nil
	}
	if minNode := t.findMin(node.Left()); minNode != nil {
		return minNode
	}
	return node
}
func (t *tree) findMax(node Node) Node {
	if node == nil {
		return nil
	}
	if maxNode := t.findMax(node.Right()); maxNode != nil {
		return maxNode
	}
	return node
}
func (t *tree) insert(root Node, node Node) Node {
	if root == nil {
		return node
	}
	if node.ID() > root.ID() {
		root.SetRight(t.insert(root.Right(), node))
	}
	if node.ID() < root.ID() {
		root.SetLeft(t.insert(root.Left(), node))
	}
	return root
}

func (t *tree) inOrder(list *[]Node, node Node) {
	if node == nil {
		return
	}
	t.inOrder(list, node.Left())
	*list = append(*list, node)
	t.inOrder(list, node.Right())
}