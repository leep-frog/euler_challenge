package bfs

import (
	"fmt"
)

/*type OffsetState[M, T any] interface {
	State[M, T]
	Offsetable[M, T]
}

type Offsetable[M, T any] interface {
	Offset(*Context[M, T]) int
}*/

/*type offsetableWrapper[M any, T OffsetState[M, T]] struct {
	state T
	dist int
}

func (ow *offsetableWrapper[M, T]) Code(ctx *Context[M, T]) string {
	return ow.state.Code(ctx)
}

func (ow *offsetableWrapper[M, T]) Done(ctx *Context[M, T]) bool {
	return ow.state.Done(ctx)
}

func (ow *offsetableWrapper[M, T]) Distance(ctx *Context[M, T]) int {
	return ow.dist
}

func (ow *offsetableWrapper[M, T]) AdjacentStates(ctx *Context[M, T]) []*offsetableWrapper[M, T] {
	var r []*offsetableWrapper[M, T]
	for _, as := range ow.state.AdjacentStates(ctx) {
		r = append(r, &offsetableWrapper[M, T]{as, ow.dist + as.Offset(ctx)})
	}
	return r
}*/

/*func offsetDistFunc[M any, T Offsetable[M, T]](ctx *Context[M, T], t T) int {
	return t.Offset(ctx)
}*/

/*func ShortestOffsetPath[M any, T OffsetState[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: adjStateDistFunc[M, T](),
		convFunc: identityConvFunc[M, T](),
	}
	return searchPath(newBFSSearcher[T](), initStates, offsetDistFunc[M, T], globalContext, ph)
}

func ShortestOffsetPathNonUnique[M any, T OffsetState[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc:   adjStateDistFunc[M, T](),
		convFunc:   identityConvFunc[M, T](),
		skipUnique: true,
	}
	return searchPath(newBFSSearcher[T](), initStates, offsetDistFunc[M, T], globalContext, ph)
}
*/

/*func toOffsetConverter[M any, T SearchableWithContextAndPath[M, T]]() converter[T, *offset[M, T]] {
	return func(t T) *offsetWrapper[M, T]{
		return &offset[M, T]{t, t.Distance}
	}
}*/

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