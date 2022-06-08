package link_list

import "errors"

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

func (s *singleLinkList) Get(index int) (int64, error) {
	temp := s.Head
	count := 0
	for temp.Next() != nil {
		temp = temp.Next()
		if count == index {
			return temp.ID(), nil
		}
		count++
	}
	return 0, errors.New("linklist overflow")
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

func (s *singleLinkList) Delete(id int64) error {
	temp := s.Head
	for temp.Next() != nil {
		if temp.Next().ID() == id {
			temp.SetNext(temp.Next().Next())
			return nil
		}
		temp = temp.Next()
	}
	return errors.New("node not found")
}

func (s *singleLinkList) Len() int {
	temp := s.Head
	var count int
	for temp.Next() != nil {
		count++
		temp = temp.Next()
	}
	return count
}

// FindReverseIndex 查詢 single_link_list 倒數第K個節點
func FindReverseIndex(linkList LinkList, k int) (int64, error) {
	return linkList.Get(linkList.Len() - k)
}
