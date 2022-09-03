package chapter14

// Node is a single node in a linked list.
type Node[T any] struct {
	Next *Node[T]
	Data T
}

// NewListFromSlice creates a new linked list using a list of any items.
func NewListFromSlice[T any](list []T) *Node[T] {
	if len(list) == 0 {
		return &Node[T]{}
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
	return start
}

// Traverse creates a slice from the linked list.
func (n *Node[T]) Traverse() []T {
	result := append([]T{}, n.Data)
	node := n.Next
	for node != nil {
		result = append(result, node.Data)
		node = node.Next
	}

	return result
}

func (n *Node[T]) Read(index int) *Node[T] {
	i := 0
	result := n
	for i < index {
		result = result.Next
		i++

		if result == nil {
			return nil
		}
	}

	return result
}
