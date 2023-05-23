package linkedlist

import (
	"fmt"
	"strings"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

func (n *Node[T]) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (n *Node[T]) ToSlice() []T {
	var ts []T
	iterate(n, nil, nil, func(k *Node[T]) {
		ts = append(ts, k.Value)
	})
	return ts
}

// TODO: Test with cycles
func (n *Node[T]) Nth(k int) *Node[T] {
	if k < 0 {
		r := n
		for i := 0; r != nil && i > k; i-- {
			r = r.Prev
		}
		return r
	} else {
		// To the right
		r := n
		for i := 0; r != nil && i < k; i++ {
			r = r.Next
		}
		return r
	}
}

func (n *Node[T]) Index(t T) (*Node[T], int) {
	for k, times := n, 0; k != nil; k, times = k.Next, times+1 {
		if k.Value == t {
			return k, times
		}
	}
	var r *Node[T]
	return r, 0
}

func iterate[T comparable](root *Node[T], nilBreakFunc func(prev *Node[T]), repeatBreakFunc func(prev, cur *Node[T]), iterFunc func(*Node[T])) {
	got := map[T]bool{}
	var prev *Node[T]
	for k := root; ; k = k.Next {
		if k == nil {
			if nilBreakFunc != nil {
				nilBreakFunc(prev)
			}
			break
		}
		if got[k.Value] {
			if repeatBreakFunc != nil {
				repeatBreakFunc(prev, k)
			}
			break
		}
		if iterFunc != nil {
			iterFunc(k)
		}
		got[k.Value] = true
		prev = k
	}
}

func (n *Node[T]) PushAt(position int, k *Node[T]) {
	if position < 0 {
		panic("Negative positions not supported")
	}

	if position == 0 {
		prev := n.Prev
		if n.Prev != nil {
			prev.Next = k
			k.Prev = prev

			n.Prev = k
			k.Next = n
		}
		return
	}

	prev := n.Nth(position - 1)
	if prev == nil {
		return
	}
	if prev.Next == nil {
		prev.Next = k
		k.Prev = prev
		return
	}

	next := prev.Next
	prev.Next = k
	k.Prev = prev
	k.Next = next
	next.Prev = k
}

func (n *Node[T]) PopAt(position int) *Node[T] {
	k := n.Nth(position)
	if k == nil {
		return nil
	}

	if k.Prev == nil && k.Next == nil {
		return k
	} else if k.Prev == nil {
		// k.Next is not nil from first if condition
		k.Next.Prev = nil
		k.Next = nil
		return k
	} else if k.Next == nil {
		// k.Prev is not nil from first if condition
		k.Prev.Next = nil
		k.Prev = nil
		return k
	}

	// Both are not nil so we need to stitch
	prev, next := k.Prev, k.Next
	prev.Next = next
	next.Prev = prev
	k.Next, k.Prev = nil, nil
	return k
}

func (n *Node[T]) PopNext() *Node[T] {
	if n == nil {
		return nil
	}
	toPop := n.Next
	if toPop == nil || n.Value == toPop.Value {
		return toPop
	}

	if toPop.Next != nil {
		toPop.Next.Prev = n
	}
	n.Next = toPop.Next

	// Unset popped node neighbors
	toPop.Prev = nil
	toPop.Next = nil

	return toPop
}

// Numbered returns n nodes starting with value 0.
func Numbered(n int) *Node[int] {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	return NewList(arr...)
}

func CircularNumbered(n int) *Node[int] {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, i)
	}
	return NewCircularList(arr...)
}

func CircularRepresentation[T comparable](n *Node[T]) string {
	var r []string
	iterate(n, nil,
		func(prev, k *Node[T]) {
			r = append(r, fmt.Sprintf("(%v)", k.Value), "...")
		},
		func(k *Node[T]) {
			if k.Prev != nil {
				if k.Prev.Next.Value != k.Value {
					panic(fmt.Sprintf("Broken link: k=%v, prev=%v, prev.next=%v", k.Value, k.Prev.Value, k.Prev.Next.Value))
				}
			}
			if k.Next != nil {
				if k.Next.Prev.Value != k.Value {
					panic(fmt.Sprintf("Broken link: k=%v, next=%v, next.prev=%v", k.Value, k.Next.Value, k.Next.Prev.Value))
				}
			}
			r = append(r, fmt.Sprintf("%v", k.Value))
		},
	)
	return strings.Join(r, " -> ")
}

func End[T comparable](n *Node[T]) *Node[T] {
	var r *Node[T]
	iterate(n, func(prev *Node[T]) {
		r = prev
	}, func(prev, cur *Node[T]) {
		r = prev
	}, nil)
	return r
}

func Len[T comparable](n *Node[T]) int {
	var r int
	iterate(n, nil, nil, func(n *Node[T]) { r++ })
	return r
}

func NewList[T comparable](ts ...T) *Node[T] {
	start := NewCircularList(ts...)
	if start == nil {
		return start
	}

	start.Prev.Next = nil
	start.Prev = nil
	return start
}

func NewCircularList[T comparable](ts ...T) *Node[T] {
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
