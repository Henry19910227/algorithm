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
