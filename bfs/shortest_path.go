package bfs

type State[M, T any] interface {
	// A unique code that represents the current state of the world.
	Code(*Context[M, T]) string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(*Context[M, T]) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(*Context[M, T]) []T
}

func ShortestPath[M any, T State[M, T]](initState T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: simpleDistFunc[M, T](),
		convFunc: identityConvFunc[M, T](),
	}
	return shortestPath(initState, nil, globalContext, ph)
}

func ShortestPathNonUnique[M any, T State[M, T]](initState T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc:   simpleDistFunc[M, T](),
		convFunc:   identityConvFunc[M, T](),
		skipUnique: true,
	}
	return shortestPath(initState, nil, globalContext, ph)
}
