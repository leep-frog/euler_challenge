package bfs

import (
	"fmt"
)

func ShortestOffsetPath[C OffsetComparable[C], T Searchable[C, T]](initStates []T, opts ...Option) ([]T, C) {
	toConverter := toACPConverter[C, T]()
	var input []*offsetWrapper[C, int, *addContextAndPathWrapper[C, T]]
	for _, is := range initStates {
		input = append(input, &offsetWrapper[C, int, *addContextAndPathWrapper[C, T]]{toConverter.convert(is), is.Distance()})
	}
	fromConverter := joinConverters(fromOffsetConverter[C, int, *addContextAndPathWrapper[C, T]](), fromACPConverter[C, T]())
	ts, dist := newSearch[C](input, 0, opts...)
	return fromConverter.convertPath(ts), dist
}

func ContextualShortestOffsetPath[C OffsetComparable[C], M any, T SearchableWithContext[C, M, T]](initStates []T, m M, opts ...Option) ([]T, C) {
	toConverter := toAPWConverter[C, M, T]()
	var input []*offsetWrapper[C, M, *addPathWrapper[C, M, T]]
	for _, is := range initStates {
		input = append(input, &offsetWrapper[C, M, *addPathWrapper[C, M, T]]{toConverter.convert(is), is.Distance(m)})
	}
	fromConverter := joinConverters(fromOffsetConverter[C, M, *addPathWrapper[C, M, T]](), fromAPWConverter[C, M, T]())
	ts, dist := newSearch[C](input, m, opts...)
	return fromConverter.convertPath(ts), dist
}

func ContextualShortestOffsetPathWithPath[C OffsetComparable[C], M any, T SearchableWithContextAndPath[C, M, T]](initStates []T, m M, opts ...Option) ([]T, C) {
	var input []*offsetWrapper[C, M, T]
	for _, is := range initStates {
		input = append(input, &offsetWrapper[C, M, T]{is, is.Distance(m, nil)})
	}
	fromConverter := fromOffsetConverter[C, M, T]()
	ts, dist := newSearch[C](input, m, opts...)
	return fromConverter.convertPath(ts), dist
}

func fromOffsetConverter[C OffsetComparable[C], M any, T SearchableWithContextAndPath[C, M, T]]() converter[*offsetWrapper[C, M, T], T] {
	return func(ow *offsetWrapper[C, M, T]) T {
		return ow.state
	}
}

// this treats distance as an offset from the previous state (as opposed to the total distance from the root)
type offsetWrapper[C OffsetComparable[C], M any, T SearchableWithContextAndPath[C, M, T]] struct {
	state T
	dist  C
}

func (ow *offsetWrapper[C, M, T]) convertedPath(p Path[*offsetWrapper[C, M, T]]) Path[T] {
	return &pathWrapper[*offsetWrapper[C, M, T], T]{p, fromOffsetConverter[C, M, T]()}
}

func (ow *offsetWrapper[C, M, T]) Code(m M, p Path[*offsetWrapper[C, M, T]]) string {
	return ow.state.Code(m, ow.convertedPath(p))
}

func (ow *offsetWrapper[C, M, T]) Done(m M, p Path[*offsetWrapper[C, M, T]]) bool {
	return ow.state.Done(m, ow.convertedPath(p))
}

func (ow *offsetWrapper[C, M, T]) String() string {
	return fmt.Sprintf("%v", ow.state)
}

func (ow *offsetWrapper[C, M, T]) Distance(m M, p Path[*offsetWrapper[C, M, T]]) C {
	return ow.dist
}

func (ow *offsetWrapper[C, M, T]) AdjacentStates(m M, p Path[*offsetWrapper[C, M, T]]) []*offsetWrapper[C, M, T] {
	newP := ow.convertedPath(p)
	var r []*offsetWrapper[C, M, T]
	for _, n := range ow.state.AdjacentStates(m, newP) {
		r = append(r, &offsetWrapper[C, M, T]{n, ow.dist.Plus(n.Distance(m, newP))})
	}
	return r
}

type OffsetComparable[C any] interface {
	Comparable[C]
	Plus(C) C
}
