package binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTree_PreOrder(t *testing.T) {
	node1 := NewNode(1, "")
	node2 := NewNode(2, "")
	node3 := NewNode(3, "")
	node4 := NewNode(4, "")
	node5 := NewNode(5, "")
	node6 := NewNode(6, "")
	node7 := NewNode(7, "")

	node1.SetLeft(node2)
	node1.SetRight(node3)
	node2.SetLeft(node4)
	node2.SetRight(node5)
	node3.SetLeft(node6)
	node3.SetRight(node7)

	tree := NewTree(node1)
	result := make([]int64, 0)
	for _, node := range tree.PreOrder() {
		result = append(result, node.ID())
	}
	assert.Equal(t, []int64{1, 2, 4, 5, 3, 6, 7}, result)
}

func TestTree_InOrder(t *testing.T) {
	node1 := NewNode(1, "")
	node2 := NewNode(2, "")
	node3 := NewNode(3, "")
	node4 := NewNode(4, "")
	node5 := NewNode(5, "")
	node6 := NewNode(6, "")
	node7 := NewNode(7, "")

	node1.SetLeft(node2)
	node1.SetRight(node3)
	node2.SetLeft(node4)
	node2.SetRight(node5)
	node3.SetLeft(node6)
	node3.SetRight(node7)

	tree := NewTree(node1)
	result := make([]int64, 0)
	for _, node := range tree.InOrder() {
		result = append(result, node.ID())
	}
	assert.Equal(t, []int64{4, 2, 5, 1, 6, 3, 7}, result)
}

func TestTree_PostOrder(t *testing.T) {
	node1 := NewNode(1, "")
	node2 := NewNode(2, "")
	node3 := NewNode(3, "")
	node4 := NewNode(4, "")
	node5 := NewNode(5, "")
	node6 := NewNode(6, "")
	node7 := NewNode(7, "")

	node1.SetLeft(node2)
	node1.SetRight(node3)
	node2.SetLeft(node4)
	node2.SetRight(node5)
	node3.SetLeft(node6)
	node3.SetRight(node7)

	tree := NewTree(node1)
	result := make([]int64, 0)
	for _, node := range tree.InOrder() {
		result = append(result, node.ID())
	}
	assert.Equal(t, []int64{4, 5, 2, 6, 7, 3, 1}, result)
}
