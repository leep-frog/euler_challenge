package bfs

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/functional"
)

type DistancePathNode[CODE comparable, DIST Distanceable[DIST], T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code(Path[T]) CODE
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(Path[T]) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(Path[T]) []T
	// Distance is the distance it took to get to this node.
	// The behavior can be changed by providing the `TotalDistance` search option.
	Distance(Path[T]) DIST
}

func DistancePathSearch[CODE comparable, DIST Distanceable[DIST], T DistancePathNode[CODE, DIST, T]](_ bool, initStates []T, opts ...Option) ([]T, DIST) {
	convertedStates := functional.Map(initStates, func(t T) *distancePathNodeWrapper[CODE, DIST, T] {
		return &distancePathNodeWrapper[CODE, DIST, T]{t}
	})
	reverter := func(sw *distancePathNodeWrapper[CODE, DIST, T]) T { return sw.state }
	return search[T, bool, CODE, DIST](false, convertedStates, reverter, opts...)
}

type distancePathNodeWrapper[CODE comparable, DIST Distanceable[DIST], T DistancePathNode[CODE, DIST, T]] struct {
	state T
}

func distPathConvert[CODE comparable, DIST Distanceable[DIST], T DistancePathNode[CODE, DIST, T]](w *distancePathNodeWrapper[CODE, DIST, T]) T {
	return w.state
}

func (sc *distancePathNodeWrapper[CODE, DIST, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *distancePathNodeWrapper[CODE, DIST, T]) Code(_ bool, p Path[*distancePathNodeWrapper[CODE, DIST, T]]) CODE {
	return sc.state.Code(&pathWrapper[*distancePathNodeWrapper[CODE, DIST, T], T]{p, distPathConvert[CODE, DIST, T]})
}

func (sc *distancePathNodeWrapper[CODE, DIST, T]) Done(_ bool, p Path[*distancePathNodeWrapper[CODE, DIST, T]]) bool {
	return sc.state.Done(&pathWrapper[*distancePathNodeWrapper[CODE, DIST, T], T]{p, distPathConvert[CODE, DIST, T]})
}

func (sc *distancePathNodeWrapper[CODE, DIST, T]) AdjacentStates(_ bool, p Path[*distancePathNodeWrapper[CODE, DIST, T]]) []*distancePathNodeWrapper[CODE, DIST, T] {
	return functional.Map(sc.state.AdjacentStates(&pathWrapper[*distancePathNodeWrapper[CODE, DIST, T], T]{p, distPathConvert[CODE, DIST, T]}), func(t T) *distancePathNodeWrapper[CODE, DIST, T] {
		return &distancePathNodeWrapper[CODE, DIST, T]{t}
	})
}

func (sc *distancePathNodeWrapper[CODE, DIST, T]) Distance(_ bool, p Path[*distancePathNodeWrapper[CODE, DIST, T]]) DIST {
	return sc.state.Distance(&pathWrapper[*distancePathNodeWrapper[CODE, DIST, T], T]{p, distPathConvert[CODE, DIST, T]})
}
