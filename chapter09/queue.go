package chapter09

import "fmt"

type Queue[T any] struct {
	data []T
}

func NewQueue[T any](data []T) *Queue[T] {
	return &Queue[T]{
		data: data,
	}
}

// Enqueue pushes a value to the end of the Queue.
func (s *Queue[T]) Enqueue(i T) {
	s.data = append(s.data, i)
}

// Dequeue pops an item from the front of the Queue.
func (s *Queue[T]) Dequeue() (T, error) {
	var i T
	if len(s.data) == 0 {
		return i, fmt.Errorf("empty")
	}
	i, s.data = s.data[0], s.data[1:]
	return i, nil
}

// Read reads the last item in the Queue.
func (s *Queue[T]) Read() (T, error) {
	if len(s.data) == 0 {
		var i T
		return i, fmt.Errorf("empty")
	}
	return s.data[0], nil
}

func (s *Queue[T]) Empty() bool {
	return len(s.data) == 0
}
