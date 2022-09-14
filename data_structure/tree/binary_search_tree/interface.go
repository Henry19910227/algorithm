package binary_search_tree

type Node interface {
	ID() int64
	Name() string
	Left() Node
	Right() Node
	SetLeft(node Node)
	SetRight(node Node)
}

type Tree interface {
	Root() Node
	Find(id int64) Node
	FindMin() Node
	FindMax() Node
}
