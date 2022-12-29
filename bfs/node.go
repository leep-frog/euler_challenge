package bfs

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/parse"
)

type Node[CODE, T any] interface {
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
}

func Search[CODE comparable, T Node[CODE, T]](initStates []T, opts ...Option) ([]T, int) {
	convertedStates := parse.Map(initStates, func(t T) *nodeWrapper[CODE, T] {
		return &nodeWrapper[CODE, T]{t}
	})
	reverter := func(sw *nodeWrapper[CODE, T]) T { return sw.state }
	p, d := search[T, bool, CODE, Int](false, convertedStates, reverter, append(opts, ignoreInitStateDistance())...)
	return p, int(d)
}

type nodeWrapper[CODE any, T Node[CODE, T]] struct {
	state T
}

func (sc *nodeWrapper[CODE, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *nodeWrapper[CODE, T]) Code(bool, Path[*nodeWrapper[CODE, T]]) CODE {
	return sc.state.Code()
}

func (sc *nodeWrapper[CODE, T]) Done(bool, Path[*nodeWrapper[CODE, T]]) bool {
	return sc.state.Done()
}

func (sc *nodeWrapper[CODE, T]) AdjacentStates(bool, Path[*nodeWrapper[CODE, T]]) []*nodeWrapper[CODE, T] {
	return parse.Map(sc.state.AdjacentStates(), func(t T) *nodeWrapper[CODE, T] {
		return &nodeWrapper[CODE, T]{t}
	})
}

func (sc *nodeWrapper[CODE, T]) Distance(bool, Path[*nodeWrapper[CODE, T]]) Int {
	return 1
}

type Int int

func (i Int) LT(j Int) bool  { return i < j }
func (i Int) Plus(j Int) Int { return i + j }
