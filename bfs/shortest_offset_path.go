package bfs

import (
	"fmt"
)

func ShortestOffsetPath[T Searchable[T]](initStates []T, opts ...Option) ([]T, int) {
	toConverter := toACPConverter[T]()
	var input []*offsetWrapper[int, *addContextAndPathWrapper[T]]
	for _, is := range initStates {
		input = append(input, &offsetWrapper[int, *addContextAndPathWrapper[T]]{toConverter.convert(is), is.Distance()})
	}
	fromConverter := joinConverters(fromOffsetConverter[int, *addContextAndPathWrapper[T]](), fromACPConverter[T]())
	ts, dist := newSearch(input, 0, opts...)
	return fromConverter.convertPath(ts), dist
}

func ContextualShortestOffsetPath[M any, T SearchableWithContext[M, T]](initStates []T, m M, opts ...Option) ([]T, int) {
	toConverter := toAPWConverter[M, T]()
	var input []*offsetWrapper[M, *addPathWrapper[M, T]]
	for _, is := range initStates {
		input = append(input, &offsetWrapper[M, *addPathWrapper[M, T]]{toConverter.convert(is), is.Distance(m)})
	}
	fromConverter := joinConverters(fromOffsetConverter[M, *addPathWrapper[M, T]](), fromAPWConverter[M, T]())
	ts, dist := newSearch(input, m, opts...)
	return fromConverter.convertPath(ts), dist
}

func ContextualShortestOffsetPathWithPath[M any, T SearchableWithContextAndPath[M, T]](initStates []T, m M, opts ...Option) ([]T, int) {
	var input []*offsetWrapper[M, T]
	for _, is := range initStates {
		input = append(input, &offsetWrapper[M, T]{is, is.Distance(m, nil)})
	}
	fromConverter := fromOffsetConverter[M, T]()
	ts, dist := newSearch(input, m, opts...)
	return fromConverter.convertPath(ts), dist
}


func fromOffsetConverter[M any, T SearchableWithContextAndPath[M, T]]() converter[*offsetWrapper[M, T], T] {
	return func(ow *offsetWrapper[M, T]) T {
		return ow.state
	}
}

// this treats distance as an offset from the previous state (as opposed to the total distance from the root)
type offsetWrapper[M any, T SearchableWithContextAndPath[M, T]] struct {
	state T
	dist int
}

func (ow *offsetWrapper[M, T]) convertedPath(p Path[*offsetWrapper[M, T]]) Path[T] {
	return &pathWrapper[*offsetWrapper[M, T], T]{p, fromOffsetConverter[M, T]()}
}

func (ow *offsetWrapper[M, T]) Code(m M, p Path[*offsetWrapper[M, T]]) string {
	return ow.state.Code(m, ow.convertedPath(p))
}

func (ow *offsetWrapper[M, T]) Done(m M, p Path[*offsetWrapper[M, T]]) bool {
	return ow.state.Done(m, ow.convertedPath(p))
}

func (ow *offsetWrapper[M, T]) String() string {
	return fmt.Sprintf("%v", ow.state)
}

func (ow *offsetWrapper[M, T]) Distance(m M, p Path[*offsetWrapper[M, T]]) int {
	return ow.dist
}

func (ow *offsetWrapper[M, T]) AdjacentStates(m M, p Path[*offsetWrapper[M, T]]) []*offsetWrapper[M, T] {
	newP := ow.convertedPath(p)
	var r []*offsetWrapper[M, T]
	for _, n := range ow.state.AdjacentStates(m, newP) {
		r = append(r, &offsetWrapper[M, T]{n, ow.dist + n.Distance(m, newP)})
	}
	return r
}