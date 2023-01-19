package maths

type Stack[T any] struct {
	items []T
}

func NewStack[T any](items ...T) *Stack[T] {
	return &Stack[T]{items}
}

func (s *Stack[T]) Push(t T) {
	s.items = append(s.items, t)
}

func (s *Stack[T]) Peek() T {
	return s.items[len(s.items)-1]
}

func (s *Stack[T]) Pop() T {
	t := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return t
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}
