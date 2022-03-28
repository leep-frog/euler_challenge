package bfs

func ShortestPath[C Comparable[C], T Searchable[C, T]](initStates []T, opts ...Option) ([]T, C) {
	toConverter := toACPConverter[C, T]()
	fromConverter := fromACPConverter[C, T]()
	ts, dist := newSearch[C](toConverter.convertSlice(initStates), 0, opts...)
	return fromConverter.convertPath(ts), dist
}

func ContextualShortestPath[C Comparable[C], M any, T SearchableWithContext[C, M, T]](initStates []T, m M, opts ...Option) ([]T, C) {
	toConverter := toAPWConverter[C, M, T]()
	fromConverter := fromAPWConverter[C, M, T]()
	ts, dist := newSearch[C](toConverter.convertSlice(initStates), m, opts...)
	return fromConverter.convertPath(ts), dist
}

func ContextualShortestPathWithPath[C Comparable[C], M any, T SearchableWithContextAndPath[C, M, T]](initStates []T, m M, opts ...Option) ([]T, C) {
	p, dist := newSearch[C](initStates, m, opts...)
	return p.Fetch(), dist
}

type Int int

func (i Int) Less(j Int) bool { return i < j }
func (i Int) Plus(j Int) Int  { return i + j }
