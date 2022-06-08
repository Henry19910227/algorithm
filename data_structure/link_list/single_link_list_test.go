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

func TestSingleLinkList_Delete(t *testing.T) {
	linkList := NewSingleLinkList()
	assert.Error(t, linkList.Delete(100))
	linkList.Add(NewNode(10))
	linkList.Add(NewNode(20))
	linkList.Add(NewNode(30))
	assert.NoError(t, linkList.Delete(20))
	assert.Equal(t, []int64{10, 30}, linkList.List())
	assert.Error(t, linkList.Delete(100))
}

func TestSingleLinkList_Len(t *testing.T) {
	linkList := NewSingleLinkList()
	assert.Equal(t, 0, linkList.Len())
	linkList.Add(NewNode(10))
	linkList.Add(NewNode(20))
	linkList.Add(NewNode(30))
	assert.Equal(t, 3, linkList.Len())
}

func TestSingleLinkList_Get(t *testing.T) {
	linkList := NewSingleLinkList()
	linkList.Add(NewNode(10))
	linkList.Add(NewNode(20))
	linkList.Add(NewNode(30))
	value, err := linkList.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), value)
	value, err = linkList.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, int64(20), value)
	value, err = linkList.Get(2)
	assert.NoError(t, err)
	assert.Equal(t, int64(30), value)
	value, err = linkList.Get(3)
	assert.Error(t, err)
	assert.Equal(t, int64(0), value)
}

// 面試題1：查詢 single_link_list 倒數第K個節點
func TestFindReverseIndex(t *testing.T) {
	linkList := NewSingleLinkList()
	linkList.Add(NewNode(10))
	linkList.Add(NewNode(20))
	linkList.Add(NewNode(30))
	value, err := FindReverseIndex(linkList, 1)
	assert.NoError(t, err)
	assert.Equal(t, int64(30), value)
	value, err = FindReverseIndex(linkList, 2)
	assert.NoError(t, err)
	assert.Equal(t, int64(20), value)
	value, err = FindReverseIndex(linkList, 3)
	assert.NoError(t, err)
	assert.Equal(t, int64(10), value)
	value, err = FindReverseIndex(linkList, 4)
	assert.Error(t, err)
	assert.Equal(t, int64(0), value)
}
