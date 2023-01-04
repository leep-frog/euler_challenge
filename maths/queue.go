package maths

type Queue[T any] struct {
	items []T
}

func NewQueue[T any](items []T) *Queue[T] {
	return &Queue[T]{items}
}

func (q *Queue[T]) Push(t T) {
	q.items = append(q.items, t)
}

func (q *Queue[T]) Peek() T {
	return q.items[0]
}

func (q *Queue[T]) Pop() T {
	t := q.items[0]
	q.items = q.items[1:]
	return t
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}
