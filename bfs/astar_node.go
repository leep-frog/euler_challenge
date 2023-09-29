package bfs

import (
	"fmt"

	"github.com/leep-frog/functional"
)

type AStarNode[CODE, T any] interface {
	Node[CODE, T]
	// AStarEstimate is the estimate length of the *remaining* distance.
	// It must be less than or equal to the actual solution distance.
	AStarEstimate() int
}

func AStarSearch[CODE comparable, T AStarNode[CODE, T]](initStates []T, opts ...Option) ([]T, int) {
	convertedStates := functional.Map(initStates, func(t T) *astarNodeWrapper[CODE, T] {
		return &astarNodeWrapper[CODE, T]{t}
	})
	reverter := func(sw *astarNodeWrapper[CODE, T]) T { return sw.state }
	p, d := search[T, bool, CODE, Int](false, convertedStates, reverter, append(opts, ignoreInitStateDistance())...)
	return p, int(d)
}

type astarNodeWrapper[CODE any, T AStarNode[CODE, T]] struct {
	state T
}

func (anw *astarNodeWrapper[CODE, T]) String() string {
	return fmt.Sprintf("%v", anw.state)
}

func (anw *astarNodeWrapper[CODE, T]) Code(bool, Path[*astarNodeWrapper[CODE, T]]) CODE {
	return anw.state.Code()
}

func (anw *astarNodeWrapper[CODE, T]) Done(bool, Path[*astarNodeWrapper[CODE, T]]) bool {
	return anw.state.Done()
}

func (anw *astarNodeWrapper[CODE, T]) AdjacentStates(bool, Path[*astarNodeWrapper[CODE, T]]) []*astarNodeWrapper[CODE, T] {
	return functional.Map(anw.state.AdjacentStates(), func(t T) *astarNodeWrapper[CODE, T] {
		return &astarNodeWrapper[CODE, T]{t}
	})
}

func (anw *astarNodeWrapper[CODE, T]) Distance(bool, Path[*astarNodeWrapper[CODE, T]]) Int {
	return 1
}

func (anw *astarNodeWrapper[CODE, T]) AStarEstimate(bool, Path[*astarNodeWrapper[CODE, T]]) Int {
	return Int(anw.state.AStarEstimate())
}
