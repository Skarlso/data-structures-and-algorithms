package chapter14

type DoubleNode[T comparable] struct {
	NextNode     *DoubleNode[T]
	PreviousNode *DoubleNode[T]
	Data         T
}

type DoublyLinkedList[T comparable] struct {
	FirstNode *DoubleNode[T]
	LastNode  *DoubleNode[T]
}

func NewDoublyLinkedList[T comparable](firstNode, lastNode *DoubleNode[T]) *DoublyLinkedList[T] {
	return &DoublyLinkedList[T]{
		FirstNode: firstNode,
		LastNode:  lastNode,
	}
}

func (t *DoublyLinkedList[T]) InsertAtEnd(val T) {
	newNode := &DoubleNode[T]{
		Data: val,
	}
	if t.FirstNode == nil {
		t.FirstNode = newNode
		t.LastNode = newNode
	} else {
		// Update the last node to be the new node and update the previously last node's link to point to this
		// newly inserted node.
		newNode.PreviousNode = t.LastNode
		t.LastNode.NextNode = newNode
		t.LastNode = newNode
	}
}

func (t *DoublyLinkedList[T]) DeleteFromFront() *DoubleNode[T] {
	removedNode := t.FirstNode
	t.FirstNode = t.FirstNode.NextNode
	return removedNode
}

func (t *DoublyLinkedList[T]) ReverseTraverse() []T {
	var result []T
	curr := t.LastNode
	for curr != nil {
		result = append(result, curr.Data)
		curr = curr.PreviousNode
	}
	return result
}

type Queue[T comparable] struct {
	queue *DoublyLinkedList[T]
}

func NewQueue[T comparable]() *Queue[T] {
	return &Queue[T]{
		queue: NewDoublyLinkedList[T](nil, nil),
	}
}

func (q *Queue[T]) Enqueue(val T) {
	q.queue.InsertAtEnd(val)
}

func (q *Queue[T]) Dequeue() T {
	removedNode := q.queue.DeleteFromFront()
	return removedNode.Data
}

func (q *Queue[T]) Read() T {
	if q.queue.FirstNode == nil {
		return q.queue.FirstNode.Data
	}
	var t T
	return t
}
