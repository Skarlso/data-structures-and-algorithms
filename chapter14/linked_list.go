package chapter14

// Node is a single node in a linked list.
type Node[T comparable] struct {
	Next *Node[T]
	Data T
}

type Tree[T comparable] struct {
	FirstNode *Node[T]
}

func NewLinkedList[T comparable](firstNode *Node[T]) *Tree[T] {
	return &Tree[T]{
		FirstNode: firstNode,
	}
}

// NewListFromSlice creates a new linked list using a list of any items.
func NewLinkedListFromSlice[T comparable](list []T) *Tree[T] {
	if len(list) == 0 {
		return &Tree[T]{
			FirstNode: &Node[T]{},
		}
	}

	start := &Node[T]{
		Data: list[0],
	}
	last := &Node[T]{}
	start.Next = last
	for i := 1; i < len(list)-1; i++ {
		last.Data = list[i]
		last.Next = &Node[T]{}
		last = last.Next
	}

	last.Data = list[len(list)-1]
	return &Tree[T]{FirstNode: start}
}

// Traverse creates a slice from the linked list.
func (t *Tree[T]) Traverse() []T {
	n := t.FirstNode
	result := append([]T{}, n.Data)
	node := n.Next
	for node != nil {
		result = append(result, node.Data)
		node = node.Next
	}

	return result
}

func (t *Tree[T]) Search(val T) int {
	index := 0
	curr := t.FirstNode
	for curr != nil {
		if curr.Data == val {
			return index
		}
		curr = curr.Next
		index++
	}

	return -1
}

func (t *Tree[T]) Insert(index int, val T) {
	node := &Node[T]{
		Data: val,
	}
	if index == 0 {
		node.Next = t.FirstNode
		t.FirstNode = node
		return
	}
	currNode := t.FirstNode
	currIndex := 0

	for currIndex < index-1 {
		currNode = currNode.Next
		currIndex++
	}
	node.Next = currNode.Next
	currNode.Next = node
}

func (t *Tree[T]) Read(index int) *Node[T] {
	i := 0
	result := t.FirstNode
	for i < index {
		result = result.Next
		i++

		if result == nil {
			return nil
		}
	}

	return result
}
