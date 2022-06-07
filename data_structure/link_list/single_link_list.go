package link_list

type singleLinkList struct {
	Head Node
}

func NewSingleLinkList() LinkList {
	head := NewNode(0)
	return &singleLinkList{Head: head}
}

func (s *singleLinkList) Add(node Node) {
	temp := s.Head
	for temp.Next() != nil {
		temp = temp.Next()
	}
	temp.SetNext(node)
}

func (s *singleLinkList) List() []int64 {
	list := make([]int64, 0)
	temp := s.Head
	for temp.Next() != nil {
		temp = temp.Next()
		list = append(list, temp.ID())
	}
	return list
}

func (s *singleLinkList) Insert(index int, node Node) {
	temp := s.Head
	var count int
	for temp.Next() != nil {
		if count == index {
			node.SetNext(temp.Next())
			temp.SetNext(node)
			return
		}
		temp = temp.Next()
		count++
	}
	s.Add(node)
}
