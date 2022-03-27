package bfs

func ShortestPath[T Searchable[T]](initStates []T, opts ...Option) ([]T, int) {
	toConverter := toACPConverter[T]()
	fromConverter := fromACPConverter[T]()
	ts, dist := newSearch(toConverter.convertSlice(initStates), 0, opts...)
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
