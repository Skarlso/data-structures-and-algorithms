package chapter16

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeap(t *testing.T) {
	heap := NewHeap([]int{100, 88, 25, 87, 16, 8, 12, 86, 50, 2, 15, 3})
	heap.Delete()
	last, ok := heap.LastNode()
	assert.True(t, ok)
	assert.Equal(t, 15, last)
	assert.Equal(t, []int{88, 87, 25, 86, 16, 8, 12, 3, 50, 2, 15}, heap.data)
}
