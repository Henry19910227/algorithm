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
