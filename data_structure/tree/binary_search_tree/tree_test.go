package binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_Find(t *testing.T) {
	node1 := NewNode(1, "one")
	node2 := NewNode(2, "two")
	node3 := NewNode(3, "three")
	node4 := NewNode(4, "four")
	node6 := NewNode(6, "six")
	node8 := NewNode(8, "eight")

	node6.SetLeft(node2)
	node6.SetRight(node8)
	node2.SetLeft(node1)
	node2.SetRight(node4)
	node4.SetLeft(node3)

	tree := NewTree(node6)
	assert.Equal(t, "three", tree.Find(3).Name())
	assert.Equal(t, "eight", tree.Find(8).Name())
	assert.Equal(t, "one", tree.Find(1).Name())
}

func TestTree_FindMin(t *testing.T) {
	node1 := NewNode(1, "one")
	node2 := NewNode(2, "two")
	node3 := NewNode(3, "three")
	node4 := NewNode(4, "four")
	node6 := NewNode(6, "six")
	node8 := NewNode(8, "eight")

	node6.SetLeft(node2)
	node6.SetRight(node8)
	node2.SetLeft(node1)
	node2.SetRight(node4)
	node4.SetLeft(node3)

	tree := NewTree(node6)
	assert.Equal(t, int64(1), tree.FindMin().ID())
}

func TestTree_FindMax(t *testing.T) {
	node1 := NewNode(1, "one")
	node2 := NewNode(2, "two")
	node3 := NewNode(3, "three")
	node4 := NewNode(4, "four")
	node6 := NewNode(6, "six")
	node8 := NewNode(8, "eight")

	node6.SetLeft(node2)
	node6.SetRight(node8)
	node2.SetLeft(node1)
	node2.SetRight(node4)
	node4.SetLeft(node3)

	tree := NewTree(node6)
	assert.Equal(t, int64(8), tree.FindMax().ID())
}

func TestTree_Insert(t *testing.T) {
	node1 := NewNode(1, "one")
	node2 := NewNode(2, "two")
	node3 := NewNode(3, "three")
	node4 := NewNode(4, "four")
	node6 := NewNode(6, "six")
	node8 := NewNode(8, "eight")

	node6.SetLeft(node2)
	node6.SetRight(node8)
	node2.SetLeft(node1)
	node2.SetRight(node4)
	node4.SetLeft(node3)

	tree := NewTree(node6)
	tree.Insert(NewNode(5, "five"))
	result := make([]int64, 0)
	for _, node := range tree.InOrder() {
		result = append(result, node.ID())
	}
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 8}, result)

	tree.Insert(NewNode(1, "one"))
	result = make([]int64, 0)
	for _, node := range tree.InOrder() {
		result = append(result, node.ID())
	}
	assert.Equal(t, []int64{1, 2, 3, 4, 5, 6, 8}, result)
}
