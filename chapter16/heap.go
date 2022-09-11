package chapter16

// Number covers all numeric types.
type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// Heap based on numbers. This could be anything that could be compared with using Equal if implemented that way.
type Heap[T Number] struct {
	data []T
}

func (h *Heap[T]) RootNode() (T, bool) {
	var t T
	if len(h.data) == 0 {
		return t, false
	}
	return h.data[0], true
}

func (h *Heap[T]) LastNode() (T, bool) {
	var t T
	if len(h.data) == 0 {
		return t, false
	}
	return h.data[len(h.data)-1], true
}

func (h *Heap[T]) LeftChild(index int) (int, bool) {
	i := (index * 2) + 1
	if i > len(h.data)-1 {
		return -1, false
	}
	return i, true
}

func (h *Heap[T]) RightChild(index int) (int, bool) {
	i := (index * 2) + 2
	if i > len(h.data)-1 {
		return -1, false
	}
	return i, true
}

func (h *Heap[T]) Parent(index int) (int, bool) {
	i := (index - 1) / 2
	if i < 0 {
		return -1, false
	}
	return i, true
}

func (h *Heap[T]) Insert(val T) {
	// Add the value as the last node
	h.data = append(h.data, val)

	// The index of the new node
	index := len(h.data) - 1

	// Trickle up

	// If the new node is not root and greater than current node:
	parentIndex, ok := h.Parent(index)
	if !ok {
		return
	}
	for index > 0 && h.data[index] > h.data[parentIndex] {
		// Do the swap
		h.data[parentIndex], h.data[index] = h.data[index], h.data[parentIndex]
		// Update the index of the new node to the swapped node
		index = parentIndex
	}
}

// Delete will delete the root node. The only delete ever allowed.
// For posterity it ignores if the data value is of 0 length.
func (h *Heap[T]) Delete() {
	// Make the last node the new root node
	var lastNode T
	// Pop
	lastNode, h.data = h.data[len(h.data)-1], h.data[:len(h.data)-1]
	h.data[0] = lastNode

	trickleIndex := 0

	// Loop until there is a child which has a higher value.
	for h.HasGreaterChild(trickleIndex) {
		// Get the largest child
		largerChildIndex := h.CalculateLargerChildIndex(trickleIndex)

		// Do the swap
		h.data[trickleIndex], h.data[largerChildIndex] = h.data[largerChildIndex], h.data[trickleIndex]

		// update the tracking index
		trickleIndex = largerChildIndex
	}
}

func (h *Heap[T]) HasGreaterChild(index int) bool {
	leftIndex, lok := h.LeftChild(index)
	rightIndex, rok := h.RightChild(index)
	return (lok && h.data[leftIndex] > h.data[index]) || (rok && h.data[rightIndex] > h.data[index])
}

func (h *Heap[T]) CalculateLargerChildIndex(index int) int {
	leftIndex, _ := h.LeftChild(index)
	rightIndex, rok := h.RightChild(index)
	if !rok {
		return leftIndex
	}
	if h.data[rightIndex] > h.data[leftIndex] {
		return rightIndex
	}
	return leftIndex
}

// NewHeap creates a new heap.
func NewHeap[T Number](data []T) *Heap[T] {
	return &Heap[T]{
		data: data,
	}
}
