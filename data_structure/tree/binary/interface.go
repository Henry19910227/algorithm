package binary

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
	PreOrder() []Node
	InOrder() []Node
	PostOrder() []Node
}
