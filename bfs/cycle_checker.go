package bfs

/*type cycleChecker[M any, T SearchableWithContext[M, T]] struct {
	state T
	remainingDepth int
	startCode string
}

func (cc *cycleChecker[M, T]) convertedPath(p Path[*cycleChecker[M, T]]) Path[T] {
	return &pathWrapper[*cycleChecker[M, T], T]{p, fromCycleConverter[M, T]()}
}

func (cc *cycleChecker[M, T]) Code(m M, p Path[*cycleChecker[M, T]]) string {
	return cc.state.Code(m, cc.convertedPath(p))
}

func (cc *cycleChecker[M, T]) Distance(m M) int {
	return 0
	//return cc.state.Distance(m, m, cc.convertedPath(p))
}

func (cc *cycleChecker[M, T]) Done(m M) bool {
	if cc.remainingDepth != 0 {
		return false
	}

	p := cc.convertedPath(p)
	for _, n := range cc.state.AdjacentStates(m) {
		if n.Code(m, p) == cc.startCode {
			return true
		}
	}
	return false
}

func (cc *cycleChecker[M, T]) AdjacentStates(m M) []*cycleChecker[M, T] {
	var r []*cycleChecker[M, T]
	for _, n := range cc.state.AdjacentStates(m) {
		r = append(r, &cycleChecker[M, T]{n, cc.remainingDepth-1, cc.startCode})
	}
	return r
}

func toCycleConverter[M any, T SearchableWithContext[M, T]](cycleLen int, m M) converter[T, *cycleChecker[M, T]] {
	return func(t T) *cycleChecker[M, T] {
		return &cycleChecker[M, T]{t, cycleLen - 1, t.Code(m)}
	}
}

func fromCycleConverter[M any, T SearchableWithContext[M, T]]() converter[*cycleChecker[M, T], T] {
	return func(cc *cycleChecker[M, T]) T {
		return cc.state
	}
}

// Check for cycles via BFS
func WeightedCycleChecker[M any, T SearchableWithContext[M, T]](initStates []T, m M, cycleLen int) ([]T, int) {
	toConverter := toCycleConverter[M, T](2, m)
	fromConverter := fromCycleConverter[M, T]()
	ts, dist := ContextualShortestOffsetPath(toConverter.convertSlice(initStates), m)
	return fromConverter.convertSlice(ts), dist
}*/
