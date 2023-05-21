package bfs

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/functional"
)

type DistanceNode[CODE comparable, DIST Distanceable[DIST], T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code() CODE
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done() bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates() []T
	// Distance is the distance it took to get to this node.
	// The behavior can be changed by providing the `TotalDistance` search option.
	Distance() DIST
}

func DistanceSearch[CODE comparable, DIST Distanceable[DIST], T DistanceNode[CODE, DIST, T]](initStates []T, opts ...Option) ([]T, DIST) {
	convertedStates := functional.Map(initStates, func(t T) *DistanceNodeWrapper[CODE, DIST, T] {
		return &DistanceNodeWrapper[CODE, DIST, T]{t}
	})
	reverter := func(sw *DistanceNodeWrapper[CODE, DIST, T]) T { return sw.state }
	return search[T, bool, CODE, DIST](false, convertedStates, reverter, opts...)
}

type DistanceNodeWrapper[CODE comparable, DIST Distanceable[DIST], T DistanceNode[CODE, DIST, T]] struct {
	state T
}

func (sc *DistanceNodeWrapper[CODE, DIST, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *DistanceNodeWrapper[CODE, DIST, T]) Code(bool, Path[*DistanceNodeWrapper[CODE, DIST, T]]) CODE {
	return sc.state.Code()
}

func (sc *DistanceNodeWrapper[CODE, DIST, T]) Done(bool, Path[*DistanceNodeWrapper[CODE, DIST, T]]) bool {
	return sc.state.Done()
}

func (sc *DistanceNodeWrapper[CODE, DIST, T]) AdjacentStates(bool, Path[*DistanceNodeWrapper[CODE, DIST, T]]) []*DistanceNodeWrapper[CODE, DIST, T] {
	return functional.Map(sc.state.AdjacentStates(), func(t T) *DistanceNodeWrapper[CODE, DIST, T] {
		return &DistanceNodeWrapper[CODE, DIST, T]{t}
	})
}

func (sc *DistanceNodeWrapper[CODE, DIST, T]) Distance(bool, Path[*DistanceNodeWrapper[CODE, DIST, T]]) DIST {
	return sc.state.Distance()
}

func (sc *DistanceNodeWrapper[CODE, DIST, T]) AStarEstimate(bool, Path[*DistanceNodeWrapper[CODE, DIST, T]]) DIST {
	var zero DIST
	return zero
}
