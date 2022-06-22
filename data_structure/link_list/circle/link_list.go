package circle

type linkList struct {
	first   Node
	current Node
}

func New() LinkList {
	return &linkList{}
}

func (l *linkList) Add(node Node) {
	if l.first == nil {
		l.first = node
		l.current = node
	}
	node.SetNext(l.first)
	l.current.SetNext(node)
	l.current = l.current.Next()
}

func (l *linkList) List() []int64 {
	temp := l.first
	outputs := make([]int64, 0)
	outputs = append(outputs, temp.ID())
	for temp.Next().ID() != l.first.ID() {
		temp = temp.Next()
		outputs = append(outputs, temp.ID())
	}
	return outputs
}

func (l *linkList) Delete(id int64) error {
	//temp := l.first
	//if temp.ID() == id {
	//
	//}
	//for temp.Next().ID() != l.first.ID() {
	//	temp = temp.Next()
	//	outputs = append(outputs, temp.ID())
	//}
	//TODO implement me
	panic("implement me")
}

func Joseph() {

}
