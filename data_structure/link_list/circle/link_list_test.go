package circle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkList_Add(t *testing.T) {
	linkList := New()
	linkList.Add(NewNode(1))
	linkList.Add(NewNode(2))
	linkList.Add(NewNode(3))
	assert.Equal(t, []int64{1}, linkList.List())
}
