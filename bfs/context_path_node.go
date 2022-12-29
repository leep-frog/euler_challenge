package bfs

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/parse"
)

type ContextPathNode[CTX any, CODE comparable, T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code(CTX, Path[T]) CODE
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(CTX, Path[T]) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(CTX, Path[T]) []T
}

func ContextPathSearch[CTX any, CODE comparable, T ContextPathNode[CTX, CODE, T]](ctx CTX, initStates []T, opts ...Option) ([]T, int) {
	convertedStates := parse.Map(initStates, func(t T) *contextPathNodeWrapper[CTX, CODE, T] {
		return &contextPathNodeWrapper[CTX, CODE, T]{t}
	})
	reverter := func(sw *contextPathNodeWrapper[CTX, CODE, T]) T { return sw.state }
	ts, d := search[T, CTX, CODE, Int](ctx, convertedStates, reverter, opts...)
	return ts, int(d)
}

type contextPathNodeWrapper[CTX any, CODE comparable, T ContextPathNode[CTX, CODE, T]] struct {
	state T
}

func ctswConvert[CTX any, CODE comparable, T ContextPathNode[CTX, CODE, T]](w *contextPathNodeWrapper[CTX, CODE, T]) T {
	return w.state
}

func (sc *contextPathNodeWrapper[CTX, CODE, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *contextPathNodeWrapper[CTX, CODE, T]) Code(ctx CTX, p Path[*contextPathNodeWrapper[CTX, CODE, T]]) CODE {
	return sc.state.Code(ctx, &pathWrapper[*contextPathNodeWrapper[CTX, CODE, T], T]{p, ctswConvert[CTX, CODE, T]})
}

func (sc *contextPathNodeWrapper[CTX, CODE, T]) Done(ctx CTX, p Path[*contextPathNodeWrapper[CTX, CODE, T]]) bool {
	return sc.state.Done(ctx, &pathWrapper[*contextPathNodeWrapper[CTX, CODE, T], T]{p, ctswConvert[CTX, CODE, T]})
}

func (sc *contextPathNodeWrapper[CTX, CODE, T]) AdjacentStates(ctx CTX, p Path[*contextPathNodeWrapper[CTX, CODE, T]]) []*contextPathNodeWrapper[CTX, CODE, T] {
	return parse.Map(sc.state.AdjacentStates(ctx, &pathWrapper[*contextPathNodeWrapper[CTX, CODE, T], T]{p, ctswConvert[CTX, CODE, T]}), func(t T) *contextPathNodeWrapper[CTX, CODE, T] {
		return &contextPathNodeWrapper[CTX, CODE, T]{t}
	})
}

func (sc *contextPathNodeWrapper[CTX, CODE, T]) Distance(CTX, Path[*contextPathNodeWrapper[CTX, CODE, T]]) Int {
	return 1
}
