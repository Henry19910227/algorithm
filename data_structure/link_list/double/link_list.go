package double

import "errors"

type linkList struct {
	head Node
}

func New() LinkList {
	head := NewNode(0)
	return &linkList{head: head}
}

func (s *linkList) Head() Node {
	return s.head
}

func (s *linkList) SetHead(node Node) {
	s.head = node
}

func (s *linkList) Add(node Node) {
	temp := s.head
	for temp.Next() != nil {
		temp = temp.Next()
	}
	node.SetPrevious(temp)
	temp.SetNext(node)
}

func (s *linkList) List() []int64 {
	list := make([]int64, 0)
	temp := s.head
	for temp.Next() != nil {
		temp = temp.Next()
		list = append(list, temp.ID())
	}
	return list
}

func (s *linkList) Get(index int) (int64, error) {
	temp := s.head
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

func (s *linkList) Insert(index int, node Node) {
	temp := s.head
	var count int
	for temp.Next() != nil {
		if count == index {
			node.SetNext(temp.Next())
			temp.Next().SetPrevious(node)

			temp.SetNext(node)
			node.SetPrevious(temp)
			return
		}
		temp = temp.Next()
		count++
	}
	s.Add(node)
}

func (s *linkList) Delete(id int64) error {
	temp := s.head.Next()
	for temp != nil {
		if temp.ID() == id {
			temp.Previous().SetNext(temp.Next())
			temp.Next().SetPrevious(temp.Previous())
			return nil
		}
		temp = temp.Next()
	}
	return errors.New("node not found")
}

func (s *linkList) Len() int {
	temp := s.head
	var count int
	for temp.Next() != nil {
		count++
		temp = temp.Next()
	}
	return count
}
