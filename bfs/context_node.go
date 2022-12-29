package bfs

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/parse"
)

type ContextNode[CTX any, CODE comparable, T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code(CTX) CODE
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(CTX) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(CTX) []T
}

func ContextSearch[CTX any, CODE comparable, T ContextNode[CTX, CODE, T]](ctx CTX, initStates []T, opts ...Option) ([]T, int) {
	convertedStates := parse.Map(initStates, func(t T) *contextNodeWrapper[CTX, CODE, T] {
		return &contextNodeWrapper[CTX, CODE, T]{t}
	})
	reverter := func(sw *contextNodeWrapper[CTX, CODE, T]) T { return sw.state }
	p, d := search[T, CTX, CODE, Int](ctx, convertedStates, reverter, append(opts, ignoreInitStateDistance())...)
	return p, int(d)
}

type contextNodeWrapper[CTX any, CODE comparable, T ContextNode[CTX, CODE, T]] struct {
	state T
}

func (sc *contextNodeWrapper[CTX, CODE, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *contextNodeWrapper[CTX, CODE, T]) Code(ctx CTX, _ Path[*contextNodeWrapper[CTX, CODE, T]]) CODE {
	return sc.state.Code(ctx)
}

func (sc *contextNodeWrapper[CTX, CODE, T]) Done(ctx CTX, _ Path[*contextNodeWrapper[CTX, CODE, T]]) bool {
	return sc.state.Done(ctx)
}

func (sc *contextNodeWrapper[CTX, CODE, T]) AdjacentStates(ctx CTX, _ Path[*contextNodeWrapper[CTX, CODE, T]]) []*contextNodeWrapper[CTX, CODE, T] {
	return parse.Map(sc.state.AdjacentStates(ctx), func(t T) *contextNodeWrapper[CTX, CODE, T] {
		return &contextNodeWrapper[CTX, CODE, T]{t}
	})
}

func (sc *contextNodeWrapper[CTX, CODE, T]) Distance(CTX, Path[*contextNodeWrapper[CTX, CODE, T]]) Int {
	return 1
}
