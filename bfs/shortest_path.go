package bfs

/*type State[M, T any] interface {
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
}*/

func ShortestPath[T Searchable[T]](initStates []T, opts ...Option) ([]T, int) {
	toConverter := toACPConverter[T]()
	fromConverter := fromACPConverter[T]()
	ts, dist := newSearch(toConverter.convertSlice(initStates), 0, opts...)
	return fromConverter.convertPath(ts), dist
}

func ShortestOffsetPath[T Searchable[T]](initStates []T, opts ...Option) ([]T, int) {
	toConverter := toACPConverter[T]()
	var input []*offsetWrapper[int, *addContextAndPathWrapper[T]]
	for _, is := range initStates {
		input = append(input, &offsetWrapper[int, *addContextAndPathWrapper[T]]{toConverter.convert(is), is.Distance()})
	}
	fromConverter := joinConverters(fromOffsetConverter[int, *addContextAndPathWrapper[T]](), fromACPConverter[T]())
	ts, dist := newSearch(input, 0, opts...)
	return fromConverter.convertPath(ts), dist
}

func ContextualShortestPath[M any, T SearchableWithContext[M, T]](initStates []T, m M, opts ...Option) ([]T, int) {
	toConverter := toAPWConverter[M, T]()
	fromConverter := fromAPWConverter[M, T]()
	ts, dist := newSearch(toConverter.convertSlice(initStates), m, opts...)
	return fromConverter.convertPath(ts), dist
}

func ContextualShortestPathWithPath[M any, T SearchableWithContextAndPath[M, T]](initStates []T, m M, opts ...Option) ([]T, int) {
	p, dist := newSearch(initStates, m, opts...)
	return p.Fetch(), dist
}

func ContextualShortestOffsetPath[M any, T SearchableWithContext[M, T]](initStates []T, m M, opts ...Option) ([]T, int) {
	toConverter := toAPWConverter[M, T]()
	var input []*offsetWrapper[M, *addPathWrapper[M, T]]
	for _, is := range initStates {
		input = append(input, &offsetWrapper[M, *addPathWrapper[M, T]]{toConverter.convert(is), is.Distance(m)})
	}
	fromConverter := joinConverters(fromOffsetConverter[M, *addPathWrapper[M, T]](), fromAPWConverter[M, T]())
	ts, dist := newSearch(input, m, opts...)
	return fromConverter.convertPath(ts), dist
}

func ContextualShortestOffsetPathWithPath[M any, T SearchableWithContextAndPath[M, T]](initStates []T, m M, opts ...Option) ([]T, int) {
	var input []*offsetWrapper[M, T]
	for _, is := range initStates {
		input = append(input, &offsetWrapper[M, T]{is, is.Distance(m, nil)})
	}
	fromConverter := fromOffsetConverter[M, T]()
	ts, dist := newSearch(input, m, opts...)
	return fromConverter.convertPath(ts), dist
}
