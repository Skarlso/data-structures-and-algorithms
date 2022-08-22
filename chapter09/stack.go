package chapter09

import "fmt"

type Stack[T any] struct {
	data []T
}

func NewStack[T any](data []T) *Stack[T] {
	return &Stack[T]{
		data: data,
	}
}

// Push pushes a value to the end of the stack.
func (s *Stack[T]) Push(i T) {
	s.data = append(s.data, i)
}

// Pop pops an item from the stack.
func (s *Stack[T]) Pop() (T, error) {
	var i T
	if len(s.data) == 0 {
		return i, fmt.Errorf("empty")
	}
	i, s.data = s.data[len(s.data)-1], s.data[:len(s.data)-1]
	return i, nil
}

// Read reads the last item in the stack.
func (s *Stack[T]) Read() (T, error) {
	var i T
	if len(s.data) == 0 {
		return i, fmt.Errorf("empty")
	}
	return s.data[len(s.data)-1], nil
}
