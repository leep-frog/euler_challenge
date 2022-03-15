package bfs

// AnyPath depth-first searches for any valid path. This can also be used to
// search for all paths by having the Done method always return false.
func AnyPath[M any, T State[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: simpleDistFunc[M, T](),
		convFunc: identityConvFunc[M, T](),
	}
	return searchPath(newDFSSearcher[T](), initStates, nil, globalContext, ph)
}
