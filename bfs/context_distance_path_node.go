package bfs

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/functional"
)

// Change order of these for best inference from search functions.
type ContextDistancePathNode[CTX any, CODE comparable, DIST Distanceable[DIST], T any] interface {
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

func ContextDistancePathSearch[CTX any, CODE comparable, DIST Distanceable[DIST], T ContextDistancePathNode[CTX, CODE, DIST, T]](ctx CTX, initStates []T, opts ...Option) ([]T, DIST) {
	convertedStates := functional.Map(initStates, func(t T) *contextDistPathNodeWrapper[CTX, CODE, DIST, T] {
		return &contextDistPathNodeWrapper[CTX, CODE, DIST, T]{t}
	})
	reverter := func(sw *contextDistPathNodeWrapper[CTX, CODE, DIST, T]) T { return sw.state }
	return search[T, CTX, CODE, DIST](ctx, convertedStates, reverter, opts...)
}

type contextDistPathNodeWrapper[CTX any, CODE comparable, DIST Distanceable[DIST], T ContextDistancePathNode[CTX, CODE, DIST, T]] struct {
	state T
}

func cdpnConvert[CTX any, CODE comparable, DIST Distanceable[DIST], T ContextDistancePathNode[CTX, CODE, DIST, T]](w *contextDistPathNodeWrapper[CTX, CODE, DIST, T]) T {
	return w.state
}

func (sc *contextDistPathNodeWrapper[CTX, CODE, DIST, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *contextDistPathNodeWrapper[CTX, CODE, DIST, T]) Code(ctx CTX, p Path[*contextDistPathNodeWrapper[CTX, CODE, DIST, T]]) CODE {
	return sc.state.Code(ctx, &pathWrapper[*contextDistPathNodeWrapper[CTX, CODE, DIST, T], T]{p, cdpnConvert[CTX, CODE, DIST, T]})
}

func (sc *contextDistPathNodeWrapper[CTX, CODE, DIST, T]) Done(ctx CTX, p Path[*contextDistPathNodeWrapper[CTX, CODE, DIST, T]]) bool {
	return sc.state.Done(ctx, &pathWrapper[*contextDistPathNodeWrapper[CTX, CODE, DIST, T], T]{p, cdpnConvert[CTX, CODE, DIST, T]})
}

func (sc *contextDistPathNodeWrapper[CTX, CODE, DIST, T]) AdjacentStates(ctx CTX, p Path[*contextDistPathNodeWrapper[CTX, CODE, DIST, T]]) []*contextDistPathNodeWrapper[CTX, CODE, DIST, T] {
	return functional.Map(sc.state.AdjacentStates(ctx, &pathWrapper[*contextDistPathNodeWrapper[CTX, CODE, DIST, T], T]{p, cdpnConvert[CTX, CODE, DIST, T]}), func(t T) *contextDistPathNodeWrapper[CTX, CODE, DIST, T] {
		return &contextDistPathNodeWrapper[CTX, CODE, DIST, T]{t}
	})
}

func (sc *contextDistPathNodeWrapper[CTX, CODE, DIST, T]) Distance(ctx CTX, p Path[*contextDistPathNodeWrapper[CTX, CODE, DIST, T]]) DIST {
	return sc.state.Distance(ctx, &pathWrapper[*contextDistPathNodeWrapper[CTX, CODE, DIST, T], T]{p, cdpnConvert[CTX, CODE, DIST, T]})
}
