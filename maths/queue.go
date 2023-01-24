package maths

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/bread"
)

type Queue[T any] struct {
	items []T
}

func NewQueue[T any](items ...T) *Queue[T] {
	return &Queue[T]{items}
}

func (q *Queue[T]) Push(ts ...T) {
	q.items = append(q.items, ts...)
}

func (q *Queue[T]) SliceCopy(start, end int) *Queue[T] {
	return &Queue[T]{bread.Copy(q.items[start:end])}
}

func (q *Queue[T]) Copy() *Queue[T] {
	return &Queue[T]{bread.Copy(q.items)}
}

func (q *Queue[T]) Peek() T {
	return q.items[0]
}

func (q *Queue[T]) Pop() T {
	t := q.items[0]
	q.items = q.items[1:]
	return t
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) Len() int {
	return len(q.items)
}

func (q *Queue[T]) String() string {
	return fmt.Sprintf("%v", q.items)
}
