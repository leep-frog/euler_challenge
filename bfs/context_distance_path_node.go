package bfs

import (
	"fmt"

	"github.com/leep-frog/functional"
)

// Change order of these for best inference from search functions.
type ContextDistancePathNode[CODE comparable, DIST Distanceable[DIST], CTX any, T any] interface {
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
	// Distance is the distance it took to get to this node.
	// The behavior can be changed by providing the `TotalDistance` search option.
	Distance(CTX, Path[T]) DIST
}

func ContextDistancePathSearch[CODE comparable, DIST Distanceable[DIST], CTX any, T ContextDistancePathNode[CODE, DIST, CTX, T]](ctx CTX, initStates []T, opts ...Option) ([]T, DIST) {
	convertedStates := functional.Map(initStates, func(t T) *contextDistPathNodeWrapper[CODE, DIST, CTX, T] {
		return &contextDistPathNodeWrapper[CODE, DIST, CTX, T]{t}
	})
	reverter := func(sw *contextDistPathNodeWrapper[CODE, DIST, CTX, T]) T { return sw.state }
	return search[T, CTX, CODE, DIST](ctx, convertedStates, reverter, opts...)
}

type contextDistPathNodeWrapper[CODE comparable, DIST Distanceable[DIST], CTX any, T ContextDistancePathNode[CODE, DIST, CTX, T]] struct {
	state T
}

func cdpnConvert[CODE comparable, DIST Distanceable[DIST], CTX any, T ContextDistancePathNode[CODE, DIST, CTX, T]](w *contextDistPathNodeWrapper[CODE, DIST, CTX, T]) T {
	return w.state
}

func (sc *contextDistPathNodeWrapper[CODE, DIST, CTX, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *contextDistPathNodeWrapper[CODE, DIST, CTX, T]) Code(ctx CTX, p Path[*contextDistPathNodeWrapper[CODE, DIST, CTX, T]]) CODE {
	return sc.state.Code(ctx, &pathWrapper[*contextDistPathNodeWrapper[CODE, DIST, CTX, T], T]{p, cdpnConvert[CODE, DIST, CTX, T]})
}

func (sc *contextDistPathNodeWrapper[CODE, DIST, CTX, T]) Done(ctx CTX, p Path[*contextDistPathNodeWrapper[CODE, DIST, CTX, T]]) bool {
	return sc.state.Done(ctx, &pathWrapper[*contextDistPathNodeWrapper[CODE, DIST, CTX, T], T]{p, cdpnConvert[CODE, DIST, CTX, T]})
}

func (sc *contextDistPathNodeWrapper[CODE, DIST, CTX, T]) AdjacentStates(ctx CTX, p Path[*contextDistPathNodeWrapper[CODE, DIST, CTX, T]]) []*contextDistPathNodeWrapper[CODE, DIST, CTX, T] {
	return functional.Map(sc.state.AdjacentStates(ctx, &pathWrapper[*contextDistPathNodeWrapper[CODE, DIST, CTX, T], T]{p, cdpnConvert[CODE, DIST, CTX, T]}), func(t T) *contextDistPathNodeWrapper[CODE, DIST, CTX, T] {
		return &contextDistPathNodeWrapper[CODE, DIST, CTX, T]{t}
	})
}

func (sc *contextDistPathNodeWrapper[CODE, DIST, CTX, T]) Distance(ctx CTX, p Path[*contextDistPathNodeWrapper[CODE, DIST, CTX, T]]) DIST {
	return sc.state.Distance(ctx, &pathWrapper[*contextDistPathNodeWrapper[CODE, DIST, CTX, T], T]{p, cdpnConvert[CODE, DIST, CTX, T]})
}

func (sc *contextDistPathNodeWrapper[CODE, DIST, CTX, T]) AStarEstimate(ctx CTX, p Path[*contextDistPathNodeWrapper[CODE, DIST, CTX, T]]) DIST {
	var zero DIST
	return zero
}
