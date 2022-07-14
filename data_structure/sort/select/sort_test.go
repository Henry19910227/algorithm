package _select

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{5, 2, 1, 3, 4}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, Sort(arr))
	arr = []int{1, 2, 3, 4, 5}
	assert.Equal(t, []int{1, 2, 3, 4, 5}, Sort(arr))
}
