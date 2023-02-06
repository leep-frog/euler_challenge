package bfs

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/functional"
)

type ContextPathNode[CODE comparable, CTX, T any] interface {
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

func ContextPathSearch[CODE comparable, CTX any, T ContextPathNode[CODE, CTX, T]](ctx CTX, initStates []T, opts ...Option) ([]T, int) {
	convertedStates := functional.Map(initStates, func(t T) *contextPathNodeWrapper[CODE, CTX, T] {
		return &contextPathNodeWrapper[CODE, CTX, T]{t}
	})
	reverter := func(sw *contextPathNodeWrapper[CODE, CTX, T]) T { return sw.state }
	ts, d := search[T, CTX, CODE, Int](ctx, convertedStates, reverter, append(opts, ignoreInitStateDistance())...)
	return ts, int(d)
}

type contextPathNodeWrapper[CODE comparable, CTX any, T ContextPathNode[CODE, CTX, T]] struct {
	state T
}

func ctswConvert[CODE comparable, CTX any, T ContextPathNode[CODE, CTX, T]](w *contextPathNodeWrapper[CODE, CTX, T]) T {
	return w.state
}

func (sc *contextPathNodeWrapper[CODE, CTX, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *contextPathNodeWrapper[CODE, CTX, T]) Code(ctx CTX, p Path[*contextPathNodeWrapper[CODE, CTX, T]]) CODE {
	return sc.state.Code(ctx, &pathWrapper[*contextPathNodeWrapper[CODE, CTX, T], T]{p, ctswConvert[CODE, CTX, T]})
}

func (sc *contextPathNodeWrapper[CODE, CTX, T]) Done(ctx CTX, p Path[*contextPathNodeWrapper[CODE, CTX, T]]) bool {
	return sc.state.Done(ctx, &pathWrapper[*contextPathNodeWrapper[CODE, CTX, T], T]{p, ctswConvert[CODE, CTX, T]})
}

func (sc *contextPathNodeWrapper[CODE, CTX, T]) AdjacentStates(ctx CTX, p Path[*contextPathNodeWrapper[CODE, CTX, T]]) []*contextPathNodeWrapper[CODE, CTX, T] {
	return functional.Map(sc.state.AdjacentStates(ctx, &pathWrapper[*contextPathNodeWrapper[CODE, CTX, T], T]{p, ctswConvert[CODE, CTX, T]}), func(t T) *contextPathNodeWrapper[CODE, CTX, T] {
		return &contextPathNodeWrapper[CODE, CTX, T]{t}
	})
}

func (sc *contextPathNodeWrapper[CODE, CTX, T]) Distance(CTX, Path[*contextPathNodeWrapper[CODE, CTX, T]]) Int {
	return 1
}
