package bfs

import (
	"fmt"

	"github.com/leep-frog/functional"
)

type PathNode[CODE comparable, T any] interface {
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
}

func PathSearch[CODE comparable, T PathNode[CODE, T]](initStates []T, opts ...Option) ([]T, int) {
	convertedStates := functional.Map(initStates, func(t T) *pathNodeWrapper[CODE, T] {
		return &pathNodeWrapper[CODE, T]{t}
	})
	reverter := func(sw *pathNodeWrapper[CODE, T]) T { return sw.state }
	p, d := search[T, bool, CODE, Int](false, convertedStates, reverter, append(opts, ignoreInitStateDistance())...)
	return p, int(d)
}

type pathNodeWrapper[CODE comparable, T PathNode[CODE, T]] struct {
	state T
}

func pathNodeConvert[CODE comparable, T PathNode[CODE, T]](w *pathNodeWrapper[CODE, T]) T {
	return w.state
}

func (sc *pathNodeWrapper[CODE, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *pathNodeWrapper[CODE, T]) Code(_ bool, p Path[*pathNodeWrapper[CODE, T]]) CODE {
	return sc.state.Code(&pathWrapper[*pathNodeWrapper[CODE, T], T]{p, pathNodeConvert[CODE, T]})
}

func (sc *pathNodeWrapper[CODE, T]) Done(_ bool, p Path[*pathNodeWrapper[CODE, T]]) bool {
	return sc.state.Done(&pathWrapper[*pathNodeWrapper[CODE, T], T]{p, pathNodeConvert[CODE, T]})
}

func (sc *pathNodeWrapper[CODE, T]) AdjacentStates(_ bool, p Path[*pathNodeWrapper[CODE, T]]) []*pathNodeWrapper[CODE, T] {
	return functional.Map(sc.state.AdjacentStates(&pathWrapper[*pathNodeWrapper[CODE, T], T]{p, pathNodeConvert[CODE, T]}), func(t T) *pathNodeWrapper[CODE, T] {
		return &pathNodeWrapper[CODE, T]{t}
	})
}

func (sc *pathNodeWrapper[CODE, T]) Distance(bool, Path[*pathNodeWrapper[CODE, T]]) Int {
	return 1
}

func (sc *pathNodeWrapper[CODE, T]) AStarEstimate(bool, Path[*pathNodeWrapper[CODE, T]]) Int {
	return 0
}
