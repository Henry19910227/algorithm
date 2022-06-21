package circle

type linkList struct {
	first   Node
	current Node
}

func New() LinkList {
	first := NewNode(1)
	first.SetNext(first)
	return &linkList{first: first, current: first}
}

func (l *linkList) Add(node Node) {
	node.SetNext(l.first)
	l.current.SetNext(node)
	l.current = node
}

func (l *linkList) List() []int64 {
	temp := l.first
	outputs := make([]int64, 0)
	outputs = append(outputs, temp.ID())
	for temp.Next().ID() != l.first.ID() {
		temp = temp.Next()
		outputs = append(outputs, temp.ID())
		break
	}
	return outputs
}

func (l *linkList) Delete(id int64) error {
	//TODO implement me
	panic("implement me")
}
