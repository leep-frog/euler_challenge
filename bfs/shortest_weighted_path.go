package bfs

// Interface used for automatically running BFS
type WeightedState[M, T any] interface {
	State[M, T]
	// Distance returns the total distance for the given state. The input is a contextual variable
	// that is passed along from ShortestPath.
	Distance(*Context[M, T]) int
}

func ShortestWeightedPath[M any, T WeightedState[M, T]](initState T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: func(ctx *Context[M, T], as T) int {
			return as.Distance(ctx)
		},
		convFunc: identityConvFunc[M, T](),
	}
	return shortestPath(initState, initState.Distance(&Context[M, T]{GlobalContext: globalContext}), globalContext, ph)
}
