package datastructures

import "errors"

// Defining a generic Stack data structure in Golang
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
	if len(s.items) == 0 {
		return *new(T), errors.New("Stack is empty")
	}

	// Fetch the last item from the stack
	item := s.items[len(s.items)-1]
	// Remove the last item from the stack
	s.items = s.items[:len(s.items)-1]
	return item, nil

}

func (s *Stack[T]) Peek() (T, error) {
	if len(s.items) == 0 {
		return *new(T), errors.New("Stack is empty")
	}

	return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.items)
}

func (s *Stack[T]) Clear() {
	s.items = []T{}
}
