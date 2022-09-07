package chapter14

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseTraverse(t *testing.T) {
	lastNode := &DoubleNode[int]{
		Data: 3,
	}
	firstNode := &DoubleNode[int]{Data: 1}
	nextNode := &DoubleNode[int]{
		Data:         2,
		PreviousNode: firstNode,
		NextNode:     lastNode,
	}
	firstNode.NextNode = nextNode
	lastNode.PreviousNode = nextNode
	n := NewDoublyLinkedList(firstNode, lastNode)
	list := n.ReverseTraverse()
	assert.Equal(t, []int{3, 2, 1}, list)
}
