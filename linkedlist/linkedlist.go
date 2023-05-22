package linkedlist

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

// Numbered returns n nodes starting with value 0.
func Numbered(n int) *Node[int] {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	return NewList(arr...)
}

func CircularRepresentation[T comparable](n *Node[T]) string {
	var r []string
	got := map[T]bool{}
	for k := n; ; k = k.Next {
		if k == nil {
			break
		}
		if got[k.Value] {
			r = append(r, fmt.Sprintf("(%v)", k.Value), "...")
			break
		}
		r = append(r, fmt.Sprintf("%v", k.Value))
		got[k.Value] = true
	}
	return strings.Join(r, " -> ")
}

func End[T comparable](n *Node[T]) *Node[T] {
	if n == nil {
		return nil
	}

	got := map[T]bool{}
	for k := n; ; k = k.Next {
		got[k.Value] = true
		if k.Next == nil || got[k.Next.Value] {
			return k
		}
	}
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func NewList[T any](ts ...T) *Node[T] {
	start := NewCircularList(ts...)
	if start == nil {
		return start
	}

	start.Prev.Next = nil
	start.Prev = nil
	return start
}

func NewCircularList[T any](ts ...T) *Node[T] {
	if len(ts) == 0 {
		return nil
	}

	var start, prev *Node[T]
	for i, t := range ts {
		n := &Node[T]{
			Value: t,
		}
		if i == 0 {
			start = n
		} else {
			prev.Next = n
			n.Prev = prev
		}
		prev = n
	}

	start.Prev = prev
	prev.Next = start

	return start
}
