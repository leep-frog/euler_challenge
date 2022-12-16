package maths

import (
	"container/heap"
)

// Heap is strictly typed heap interface
type Heap[T any] struct {
	ih *internalHeap[T]
}

func (h *Heap[T]) Len() int {
	return h.ih.Len()
}

func (h *Heap[T]) Push(t T) {
	heap.Push(h.ih, t)
}

func (h *Heap[T]) Pop() T {
	return heap.Pop(h.ih).(T)
}

// It is very important to note that we point to the same object
// that is still in the heap (we do not create a copy of it).
func (h *Heap[T]) Peek() T {
	t := h.ih.Pop().(T)
	h.Push(t)
	return t
}

func NewHeap[T any](lt func(T, T) bool) *Heap[T] {
	return &Heap[T]{&internalHeap[T]{nil, lt}}
}

// internalHeap is an internal interface used to satisfy the heap.Interface.
type internalHeap[T any] struct {
	items []T
	lt    func(T, T) bool
}

// Needed by heap.Interface
func (h *internalHeap[T]) Len() int {
	return len(h.items)
}

// Needed by heap.Interface
func (h *internalHeap[T]) Less(i, j int) bool {
	return h.lt(h.items[i], h.items[j])
}

// Needed by heap.Interface
func (h *internalHeap[T]) Swap(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
}

func (h *internalHeap[T]) Push(t any) {
	h.items = append(h.items, t.(T))
}

func (h *internalHeap[T]) Pop() any {
	last := h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	return last
}
