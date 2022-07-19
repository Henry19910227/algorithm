package base

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwap(t *testing.T) {
	a := 10
	b := 20
	Swap(&a, &b)
	assert.Equal(t, 20, a)
	assert.Equal(t, 10, b)
}

func TestMoveBack(t *testing.T) {
	nums := make([]int, 5)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	MoveBack(nums)
	assert.Equal(t, []int{5,1,2,3,4}, nums)
}
