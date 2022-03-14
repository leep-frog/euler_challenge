package bfs

type OffsetState[M, T any] interface {
	State[M, T]
	Offsetable[M, T]
}

type Offsetable[M, T any] interface {
	Offset(*Context[M, T]) int
}

func offsetDistFunc[M any, T Offsetable[M, T]](ctx *Context[M, T], t T) int {
	return t.Offset(ctx)
}

func ShortestOffsetPath[M any, T OffsetState[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: adjStateDistFunc[M, T](),
		convFunc: identityConvFunc[M, T](),
	}
	return searchPath(newBFSSearcher[T](), initStates, offsetDistFunc[M, T], globalContext, ph)
}

func ShortestOffsetPathNonUnique[M any, T OffsetState[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc:   adjStateDistFunc[M, T](),
		convFunc:   identityConvFunc[M, T](),
		skipUnique: true,
	}
	return searchPath(newBFSSearcher[T](), initStates, offsetDistFunc[M, T], globalContext, ph)
}
