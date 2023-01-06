package bfs

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/functional"
)

type ContextDistanceNode[CTX any, CODE comparable, DIST Distanceable[DIST], T any] interface {
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
	// Distance is the distance it took to get to this node.
	// The behavior can be changed by providing the `TotalDistance` search option.
	Distance(CTX) DIST
}

func ContextDistanceSearch[CTX any, CODE comparable, DIST Distanceable[DIST], T ContextDistanceNode[CTX, CODE, DIST, T]](ctx CTX, initStates []T, opts ...Option) ([]T, DIST) {
	convertedStates := functional.Map(initStates, func(t T) *contextDistanceNodeWrapper[CTX, CODE, DIST, T] {
		return &contextDistanceNodeWrapper[CTX, CODE, DIST, T]{t}
	})
	reverter := func(sw *contextDistanceNodeWrapper[CTX, CODE, DIST, T]) T { return sw.state }
	return search[T, CTX, CODE, DIST](ctx, convertedStates, reverter, opts...)
}

type contextDistanceNodeWrapper[CTX any, CODE comparable, DIST Distanceable[DIST], T ContextDistanceNode[CTX, CODE, DIST, T]] struct {
	state T
}

func (sc *contextDistanceNodeWrapper[CTX, CODE, DIST, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *contextDistanceNodeWrapper[CTX, CODE, DIST, T]) Code(ctx CTX, _ Path[*contextDistanceNodeWrapper[CTX, CODE, DIST, T]]) CODE {
	return sc.state.Code(ctx)
}

func (sc *contextDistanceNodeWrapper[CTX, CODE, DIST, T]) Done(ctx CTX, _ Path[*contextDistanceNodeWrapper[CTX, CODE, DIST, T]]) bool {
	return sc.state.Done(ctx)
}

func (sc *contextDistanceNodeWrapper[CTX, CODE, DIST, T]) AdjacentStates(ctx CTX, _ Path[*contextDistanceNodeWrapper[CTX, CODE, DIST, T]]) []*contextDistanceNodeWrapper[CTX, CODE, DIST, T] {
	return functional.Map(sc.state.AdjacentStates(ctx), func(t T) *contextDistanceNodeWrapper[CTX, CODE, DIST, T] {
		return &contextDistanceNodeWrapper[CTX, CODE, DIST, T]{t}
	})
}

func (sc *contextDistanceNodeWrapper[CTX, CODE, DIST, T]) Distance(ctx CTX, _ Path[*contextDistanceNodeWrapper[CTX, CODE, DIST, T]]) DIST {
	return sc.state.Distance(ctx)
}
