package bfs

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/functional"
)

// ContextNode defines an interface that can be used by the `ContextSearch` function.
// Note the `CTX` generic type is at the end because it can be inferred.
type ContextNode[CODE comparable, CTX, T any] interface {
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

func ContextSearch[CODE comparable, CTX any, T ContextNode[CODE, CTX, T]](ctx CTX, initStates []T, opts ...Option) ([]T, int) {
	convertedStates := functional.Map(initStates, func(t T) *contextNodeWrapper[CODE, CTX, T] {
		return &contextNodeWrapper[CODE, CTX, T]{t}
	})
	reverter := func(sw *contextNodeWrapper[CODE, CTX, T]) T { return sw.state }
	p, d := search[T, CTX, CODE, Int](ctx, convertedStates, reverter, append(opts, ignoreInitStateDistance())...)
	return p, int(d)
}

type contextNodeWrapper[CODE comparable, CTX any, T ContextNode[CODE, CTX, T]] struct {
	state T
}

func (sc *contextNodeWrapper[CODE, CTX, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *contextNodeWrapper[CODE, CTX, T]) Code(ctx CTX, _ Path[*contextNodeWrapper[CODE, CTX, T]]) CODE {
	return sc.state.Code(ctx)
}

func (sc *contextNodeWrapper[CODE, CTX, T]) Done(ctx CTX, _ Path[*contextNodeWrapper[CODE, CTX, T]]) bool {
	return sc.state.Done(ctx)
}

func (sc *contextNodeWrapper[CODE, CTX, T]) AdjacentStates(ctx CTX, _ Path[*contextNodeWrapper[CODE, CTX, T]]) []*contextNodeWrapper[CODE, CTX, T] {
	return functional.Map(sc.state.AdjacentStates(ctx), func(t T) *contextNodeWrapper[CODE, CTX, T] {
		return &contextNodeWrapper[CODE, CTX, T]{t}
	})
}

func (sc *contextNodeWrapper[CODE, CTX, T]) Distance(CTX, Path[*contextNodeWrapper[CODE, CTX, T]]) Int {
	return 1
}
