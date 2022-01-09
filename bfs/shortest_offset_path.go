package bfs

type OffsetState[M, T any] interface {
	// Code returns a unique code for a given state. Used to ensure we don't check the same state
	// more than once.
	Code(*Context[M, T]) string
	// Returns if the given state is in a final position. The first input is a contextual variable
	// that is passed along from ShortestPath. The second input is the depth.
	Done(*Context[M, T]) bool
	// Returns all pairs of the adjacent states and those states offsets from this state.
	// The input is a contextual variable that is passed along from ShortestPath.
	// TODO: change this interface to just include a Offset() functio
	AdjacentStates(*Context[M, T]) []*AdjacentState[T]
}

type AdjacentState[T any] struct {
	State  T
	Offset int
}

func ShortestOffsetPath[M any, T OffsetState[M, T]](initState T, initDist int, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, *AdjacentState[T]]{
		distFunc: adjStateDistFunc[M, T](),
		convFunc: adjStateConvFunc[M, T](),
	}
	return shortestPath(initState, initDist, globalContext, ph)
}

func ShortestOffsetPathNonUnique[M any, T OffsetState[M, T]](initState T, initDist int, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, *AdjacentState[T]]{
		distFunc: adjStateDistFunc[M, T](),
		convFunc: adjStateConvFunc[M, T](),
		skipUnique: true,
	}
	return shortestPath(initState, initDist, globalContext, ph)
}
