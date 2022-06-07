package link_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleLinkList_Add(t *testing.T) {
	linkList := NewSingleLinkList()
	linkList.Add(NewNode(1))
	linkList.Add(NewNode(2))
	linkList.Add(NewNode(3))
	assert.Equal(t, []int64{1, 2, 3}, linkList.List())
}

func TestSingleLinkList_Insert(t *testing.T) {
	linkList := NewSingleLinkList()
	linkList.Add(NewNode(10))
	linkList.Add(NewNode(20))
	linkList.Add(NewNode(30))
	linkList.Insert(0, NewNode(5))
	linkList.Insert(2, NewNode(15))
	linkList.Insert(10, NewNode(100))
	assert.Equal(t, []int64{5, 10, 15, 20, 30, 100}, linkList.List())
}
