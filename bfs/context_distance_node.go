package bfs

import (
	"fmt"

	"github.com/leep-frog/functional"
)

type ContextDistanceNode[CODE comparable, DIST Distanceable[DIST], CTX any, T any] interface {
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

func ContextDistanceSearch[CODE comparable, DIST Distanceable[DIST], CTX any, T ContextDistanceNode[CODE, DIST, CTX, T]](ctx CTX, initStates []T, opts ...Option) ([]T, DIST) {
	convertedStates := functional.Map(initStates, func(t T) *contextDistanceNodeWrapper[CODE, DIST, CTX, T] {
		return &contextDistanceNodeWrapper[CODE, DIST, CTX, T]{t}
	})
	reverter := func(sw *contextDistanceNodeWrapper[CODE, DIST, CTX, T]) T { return sw.state }
	return search[T, CTX, CODE, DIST](ctx, convertedStates, reverter, opts...)
}

type contextDistanceNodeWrapper[CODE comparable, DIST Distanceable[DIST], CTX any, T ContextDistanceNode[CODE, DIST, CTX, T]] struct {
	state T
}

func (sc *contextDistanceNodeWrapper[CODE, DIST, CTX, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *contextDistanceNodeWrapper[CODE, DIST, CTX, T]) Code(ctx CTX, _ Path[*contextDistanceNodeWrapper[CODE, DIST, CTX, T]]) CODE {
	return sc.state.Code(ctx)
}

func (sc *contextDistanceNodeWrapper[CODE, DIST, CTX, T]) Done(ctx CTX, _ Path[*contextDistanceNodeWrapper[CODE, DIST, CTX, T]]) bool {
	return sc.state.Done(ctx)
}

func (sc *contextDistanceNodeWrapper[CODE, DIST, CTX, T]) AdjacentStates(ctx CTX, _ Path[*contextDistanceNodeWrapper[CODE, DIST, CTX, T]]) []*contextDistanceNodeWrapper[CODE, DIST, CTX, T] {
	return functional.Map(sc.state.AdjacentStates(ctx), func(t T) *contextDistanceNodeWrapper[CODE, DIST, CTX, T] {
		return &contextDistanceNodeWrapper[CODE, DIST, CTX, T]{t}
	})
}

func (sc *contextDistanceNodeWrapper[CODE, DIST, CTX, T]) Distance(ctx CTX, _ Path[*contextDistanceNodeWrapper[CODE, DIST, CTX, T]]) DIST {
	return sc.state.Distance(ctx)
}

func (sc *contextDistanceNodeWrapper[CODE, DIST, CTX, T]) AStarEstimate(ctx CTX, _ Path[*contextDistanceNodeWrapper[CODE, DIST, CTX, T]]) DIST {
	var zero DIST
	return zero
}
