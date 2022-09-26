# Node-Based Data Structures

Now we get to the interesting part. Linked lists.

A linked list is a collection of items where Node 1 points to Node 2 points to Node 3... points to Node N.
It is not the best performance-wise in reading and searching, but it shines tremendously when deleting or inserting in
certain scenarios.

The structure also maps nicely to specific algorithms and usages. A linked list, for example, is a perfect way to
represent queues, as we'll see later.
Another neat thing is that lists need a continuous chunk of memory; whereas for linked lists, the data can be located anywhere in memory as long as there is a point to it from one of the nodes.

## Singly Linked Lists

Let's look at implementing a linked list.

```go
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

```

I threw in the `NewListFromSlice` to make it convenient to test. We need a `comparable` type because, for searching, we
need the ability to compare the data values. That said, practically, it could be anything. Because we can
just call `Compare` on any respective type that can be compered. But for that, we would need some very specific API.
For example, in case of strings, like book titles, we could use an alphabetical order. Or better. a hash!

### Searching

Searching has a complexity of O(n) since we need to look through all the items. This search returns the index of an
item. We can use the index of the item in the `Read` operation, but that would require another O(n) run.

```go
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
```

Trivial.

### Insertion

Insertion can take O(n) for a regular array. But for a linked list, inserting to the beginning is O(1). To insert
anywhere else, we need O(n) because we need to get to the right node to insert at.

```go
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
```

### Deletion

Deletion efficiency is defined by the following table:

| Situation           | Array        | Linked List  |
| ------------------- | ------------ | ------------ |
| Delete at beginning | Worst Case   | Best Case    |
| Delete at middle    | Average Case | Average Case |
| Delete at end       | Best Case    | Worst case   |

Code:

```go

func (t *Tree[T]) Delete(index int) {
	// Deleting the first item is easy.
	if index == 0 {
		t.FirstNode = t.FirstNode.Next
		return
	}

	currNode := t.FirstNode
	currIndex := 0
	// We loop through the list to find the item we are looking for.
	for currIndex < index-1 {
		currNode = currNode.Next
		currIndex++
	}
	// Save the next node after the one we will be deleting.
	nodeAfterDeletedNode := currNode.Next.Next
	// Update the link to the node to point to the next node leaving out the current node of the link.
	currNode.Next = nodeAfterDeletedNode
}
```

### Efficiency of Linked ist

| Operation | Array              | Linked List              |
| --------- | ------------------ | ------------------------ |
| Reading   | O(1)               | O(n)                     |
| Search    | O(n)               | O(n)                     |
| Insertion | O(n) (O(1) at end) | O(n) (O(1) at beginning) |
| Deletion  | O(n) (O(1) at end) | O(n) (O(1) at beginning) |

Linked lists are super effective if you encounter data that is being manipulated, inserting and deleting, often as you
read through it.

## Doubly Linked Lists

A doubly linked list is like a singly linked list just that the nodes point back to the previous node as well as the
next node.

To implement a doubly linked list we have to update the `Node` to have a `Previous` one and update our `Tree` to have a
`LastNode` item as well.

```go

type DoubleNode[T comparable] struct {
	Next         *Node[T]
	PreviousNode *Node[T]
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
```

We can now delete and insert in O(1) from the end and the beginning as well.

```go
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
		t.LastNode.Next = newNode
		t.LastNode = newNode
	}
}
```

But bare in mind, that never connect the last node with the first node. The last node needs to be `nil` for NextNode.
Otherwise, it would create an infinite loop when traversing the list.

### Queues with Doubly Linked List

Turns out, Doubly Linked Lists are _perfect_ for queues.

```go
type DoubleNode[T comparable] struct {
	Next         *DoubleNode[T]
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
		t.LastNode.Next = newNode
		t.LastNode = newNode
	}
}

func (t *DoublyLinkedList[T]) DeleteFromFront() *DoubleNode[T] {
	removedNode := t.FirstNode
	t.FirstNode = t.FirstNode.Next
	return removedNode
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
```

Implementing the queue this way, we can now insert and delete from the queue in O(1).

## Exercises

1. Add a method to the classic LinkedList that prints all the elements of the list.

Answer: Hah, I already added Traverse.

```go
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
```

2. Add a method to DoublyLinkedList that print all the elements of the list in _reverse_ order.

Answer:

```go
func (t *DoublyLinkedList[T]) ReverseTraverse() []T {
	var result []T
	curr := t.LastNode
	for curr != nil {
		result = append(result, curr.Data)
		curr = curr.PreviousNode
	}
	return result
}
```

3. Add a method to LinkedList returns the last element in the list.

```go
func (t *Tree[T]) LastNode() *Node[T] {
	last := t.FirstNode

	for last.Next != nil {
		last = last.Next
	}
	return last
}
```

4. Tricky one. Add a method to LinkedList that reverses the list in place.

Answer:

```go
func (t *Tree[T]) Reverse() {
	var previousNode *Node[T]
	currentNode := t.FirstNode

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	t.FirstNode = previousNode
}
```

5. Delete an item from the list which when you only have a node in the middle of the list.

Answer:

```go
func DeleteMiddle(node *Node[T]) {
	node.Data = node.Next.Data
	node.Next = node.Next.Next
}
```