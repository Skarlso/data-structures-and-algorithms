package chapter14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	items := []int{1, 2, 3, 4, 5, 6}
	n := NewListFromSlice(items)
	d := n.Data
	assert.Equal(t, 1, d)
	node := n.Read(5)
	assert.NotNil(t, node)
	assert.Equal(t, 6, node.Data)
	trav := n.Traverse()
	assert.Equal(t, items, trav)
}
