package bfs

// Interface used for automatically running BFS
/*type WeightedState[M, T any] interface {
	State[M, T]
	// Distance returns the total distance for the given state. The input is a contextual variable
	// that is passed along from ShortestPath.
	Distance(*Context[M, T]) int
}

func initDistFunc[M any, T WeightedState[M, T]](ctx *Context[M, T], t T) int {
	return t.Distance(ctx)
}

// TODO: change this and interface to cumulative or something. Weighted should
// be the offset (since edges are actually weighted there).
/*func ShortestWeightedPath[M any, T WeightedState[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: func(ctx *Context[M, T], as T) int {
			return as.Distance(ctx)
		},
		convFunc: identityConvFunc[M, T](),
	}
	return searchPath(newBFSSearcher[T](), initStates, initDistFunc[M, T], globalContext, ph)
}*/

/*func ClosestNode[M any, T WeightedState[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: func(ctx *Context[M, T], as T) int {
			return as.Distance(ctx)
		},
		convFunc: identityConvFunc[M, T](),
	}
	return searchPath(newBFSSearcher[T](), initStates, initDistFunc[M, T], globalContext, ph)
}*/
