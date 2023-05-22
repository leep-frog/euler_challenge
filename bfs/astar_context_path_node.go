package bfs

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/functional"
)

type AStarContextPathNode[CODE comparable, CTX, T any] interface {
	ContextPathNode[CODE, CTX, T]
	// AStarEstimate is the estimate length of the *remaining* distance.
	// It must be less than or equal to the actual solution distance.
	AStarEstimate(CTX, Path[T]) int
}

func AStarContextPathSearch[CODE comparable, CTX any, T AStarContextPathNode[CODE, CTX, T]](ctx CTX, initStates []T, opts ...Option) ([]T, int) {
	convertedStates := functional.Map(initStates, func(t T) *aStarContextPathNodeWrapper[CODE, CTX, T] {
		return &aStarContextPathNodeWrapper[CODE, CTX, T]{t}
	})
	reverter := func(sw *aStarContextPathNodeWrapper[CODE, CTX, T]) T { return sw.state }
	ts, d := search[T, CTX, CODE, Int](ctx, convertedStates, reverter, append(opts, ignoreInitStateDistance())...)
	return ts, int(d)
}

type aStarContextPathNodeWrapper[CODE comparable, CTX any, T AStarContextPathNode[CODE, CTX, T]] struct {
	state T
}

func actswConvert[CODE comparable, CTX any, T AStarContextPathNode[CODE, CTX, T]](w *aStarContextPathNodeWrapper[CODE, CTX, T]) T {
	return w.state
}

func (sc *aStarContextPathNodeWrapper[CODE, CTX, T]) String() string {
	return fmt.Sprintf("%v", sc.state)
}

func (sc *aStarContextPathNodeWrapper[CODE, CTX, T]) Code(ctx CTX, p Path[*aStarContextPathNodeWrapper[CODE, CTX, T]]) CODE {
	return sc.state.Code(ctx, &pathWrapper[*aStarContextPathNodeWrapper[CODE, CTX, T], T]{p, actswConvert[CODE, CTX, T]})
}

func (sc *aStarContextPathNodeWrapper[CODE, CTX, T]) Done(ctx CTX, p Path[*aStarContextPathNodeWrapper[CODE, CTX, T]]) bool {
	return sc.state.Done(ctx, &pathWrapper[*aStarContextPathNodeWrapper[CODE, CTX, T], T]{p, actswConvert[CODE, CTX, T]})
}

func (sc *aStarContextPathNodeWrapper[CODE, CTX, T]) AdjacentStates(ctx CTX, p Path[*aStarContextPathNodeWrapper[CODE, CTX, T]]) []*aStarContextPathNodeWrapper[CODE, CTX, T] {
	return functional.Map(sc.state.AdjacentStates(ctx, &pathWrapper[*aStarContextPathNodeWrapper[CODE, CTX, T], T]{p, actswConvert[CODE, CTX, T]}), func(t T) *aStarContextPathNodeWrapper[CODE, CTX, T] {
		return &aStarContextPathNodeWrapper[CODE, CTX, T]{t}
	})
}

func (sc *aStarContextPathNodeWrapper[CODE, CTX, T]) Distance(CTX, Path[*aStarContextPathNodeWrapper[CODE, CTX, T]]) Int {
	return 1
}

func (sc *aStarContextPathNodeWrapper[CODE, CTX, T]) AStarEstimate(ctx CTX, p Path[*aStarContextPathNodeWrapper[CODE, CTX, T]]) Int {
	return Int(sc.state.AStarEstimate(ctx, &pathWrapper[*aStarContextPathNodeWrapper[CODE, CTX, T], T]{p, actswConvert[CODE, CTX, T]}))
}
