package linkedlist

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Value T
	Next *Node[T]
	Prev *Node[T]
}

// List with hashed positions
//type List

func (n *Node[T]) String() string {
	var r []string
	for k := n; k != nil; k = k.Next {
		r = append(r, fmt.Sprintf("%v", k.Value))
	}
	return strings.Join(r, " -> ")
}

func NewList[T any](ts []T) *Node[T] {
	if len(ts) == 0 {
		return nil
	}
	n := &Node[T]{
		Value: ts[0],
	}
	rest := NewList(ts[1:])
	if rest == nil {
		return n
	}
	n.Next = rest
	rest.Prev = n
	return n
}
