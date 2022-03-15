package bfs

type State[M, T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code(*Context[M, T]) string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(*Context[M, T]) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(*Context[M, T]) []T
}

func ShortestPath[M any, T State[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: simpleDistFunc[M, T](),
		convFunc: identityConvFunc[M, T](),
	}
	return searchPath(newBFSSearcher[T](), initStates, nil, globalContext, ph)
}

// TODO: change this to accept interface without Code
func ShortestPathNonUnique[M any, T State[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: simpleDistFunc[M, T](),
		convFunc: identityConvFunc[M, T](),
		skipUnique: true,
	}
	return searchPath(newBFSSearcher[T](), initStates, nil, globalContext, ph)
}