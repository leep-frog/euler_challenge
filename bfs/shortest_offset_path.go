package bfs

type OffsetState[M, T any] interface {
	State[M, T]
	Offset(*Context[M, T]) int
}

func ShortestOffsetPath[M any, T OffsetState[M, T]](initState T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: adjStateDistFunc[M, T](),
		convFunc: identityConvFunc[M, T](),
	}
	return shortestPath(initState, initState.Offset, globalContext, ph)
}

func ShortestOffsetPathNonUnique[M any, T OffsetState[M, T]](initState T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc:   adjStateDistFunc[M, T](),
		convFunc:   identityConvFunc[M, T](),
		skipUnique: true,
	}
	return shortestPath(initState, initState.Offset, globalContext, ph)
}
