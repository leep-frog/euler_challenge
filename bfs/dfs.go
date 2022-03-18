package bfs

type DepthSearchable[T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code() string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done() bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates() []T
}

type addContextAndPathWrapperDFS[T DepthSearchable[T]] struct {
	state T
}

func (acp *addContextAndPathWrapperDFS[T]) Code(_ int, _ DFSPath[*addContextAndPathWrapperDFS[T]]) string {
	return acp.state.Code()
}

func (acp *addContextAndPathWrapperDFS[T]) Done(_ int, _ DFSPath[*addContextAndPathWrapperDFS[T]]) bool {
	return acp.state.Done()
}

func (acp *addContextAndPathWrapperDFS[T]) AdjacentStates(_ int, _ DFSPath[*addContextAndPathWrapperDFS[T]]) []*addContextAndPathWrapperDFS[T] {
	var r []*addContextAndPathWrapperDFS[T]
	for _, n := range acp.state.AdjacentStates() {
		r = append(r, &addContextAndPathWrapperDFS[T]{n})
	}
	return r
}

type DepthSearchableWithContext[M, T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code(M) string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(M) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(M) []T
}

type addPathWrapperDFS[M any, T DepthSearchableWithContext[M, T]] struct {
	state T
}

func (apw *addPathWrapperDFS[M, T]) Code(m M, _ DFSPath[*addPathWrapperDFS[M, T]]) string {
	return apw.state.Code(m)
}

func (apw *addPathWrapperDFS[M, T]) Done(m M, _ DFSPath[*addPathWrapperDFS[M, T]]) bool {
	return apw.state.Done(m)
}

func (apw *addPathWrapperDFS[M, T]) AdjacentStates(m M, _ DFSPath[*addPathWrapperDFS[M, T]]) []*addPathWrapperDFS[M, T] {
	var r []*addPathWrapperDFS[M, T]
	for _, n := range apw.state.AdjacentStates(m) {
		r = append(r, &addPathWrapperDFS[M, T]{n})
	}
	return r
}

type DepthSearchableWithContextAndPath[M, T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code(M, DFSPath[T]) string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(M, DFSPath[T]) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(M, DFSPath[T]) []T
}

type DFSPath[T any] interface {
	Path() []T
	Len() int
	Contains(string) bool
}

type dfsPath[T any] struct {
	path []T
	// Change value to int so we can keep count of instances?
	set map[string]bool
}

func (dp *dfsPath[T]) pop(s string) {
	dp.path = dp.path[:len(dp.path)-1]
	delete(dp.set, s)
}

func (dp *dfsPath[T]) push(t T, s string) {
	dp.path = append(dp.path, t)
	dp.set[s] = true
}

func (dp *dfsPath[T]) Path() []T {
	return dp.path
}

func (dp *dfsPath[T]) Len() int {
	return len(dp.path)
}

func (dp *dfsPath[T]) Contains(s string) bool {
	return dp.set[s]
}

type dfsAction[T any] struct {
	popPath *string
	state T	
}

func DFSWithContext[M any, T DepthSearchableWithContext[M, T]](initStates []T, m M) []T {
	var r []*addPathWrapperDFS[M, T]
	for _, is := range initStates {
		r = append(r, &addPathWrapperDFS[M, T]{is})
	}
	var ts []T
	for _, wrapped := range DFSWithContextAndPath(r, m) {
		ts = append(ts, wrapped.state)
	}
	return ts
}

func DFS[M any, T DepthSearchable[T]](initStates []T) []T {
	var r []*addContextAndPathWrapperDFS[T]
	for _, is := range initStates {
		r = append(r, &addContextAndPathWrapperDFS[T]{is})
	}
	var ts []T
	for _, wrapped := range DFSWithContextAndPath(r, 0) {
		ts = append(ts, wrapped.state)
	}
	return ts
}

type DFSOption func(o *dfsOption)

type dfsOption struct {
	allowCycles bool
	allowDuplicates bool
}

func AllowDFSCycles() DFSOption {
	return func(o *dfsOption) {o.allowCycles = true}
}

func AllowDFSDuplicates() DFSOption {
	return func(o *dfsOption) {o.allowDuplicates = true}
}

func DFSWithContextAndPath[M any, T DepthSearchableWithContextAndPath[M, T]](initStates []T, m M, opts ...DFSOption) []T {
	opt := &dfsOption{}
	for _, o := range opts {
		o(opt)
	}
	var actions []*dfsAction[T]
	for _, is := range initStates {
		actions = append(actions, &dfsAction[T]{nil, is})
	}

	// TODO(option): Check each node at most once vs check unique path at most once
	// For example if A and B both have edge to C, do we want to check
	// C once, or check once for A->C and once for A->C

	checkedNodes := map[string]bool{}

	dp := &dfsPath[T]{nil, map[string]bool{}}

	for len(actions) > 0 {
		a := actions[len(actions)-1]
		actions = actions[:len(actions)-1]

		// Remove from path if relevant.
		if a.popPath != nil{
			dp.pop(*a.popPath)
			// TODO: postprocess state
			continue
		}

		// Otherwise, process the state
		state := a.state
		c := state.Code(m, dp)

		// See if we've visited this node before
		if !opt.allowCycles && dp.Contains(c) {
			// TODO: cycleHandler
			continue
		}
		if !opt.allowDuplicates && checkedNodes[c] {
			// TODO: alreadyVisitedHandler
			continue
		}

		// Pop once we examine all of this node's children
		actions = append(actions, &dfsAction[T]{&c, state})

		// Update path variables
		dp.push(state, c)
		checkedNodes[c] = true

		// Check if node is done
		if state.Done(m, dp) {
			return dp.Path()
		}

		// TODO: preprocess state
		for _, neighbor := range state.AdjacentStates(m, dp) {
			actions = append(actions, &dfsAction[T]{nil, neighbor})
		}
	}
	return nil
}

/*type cycleState[M any, T Set[M, T]] struct {
	state T
}

func (cs *cycleState) Done() bool {
	// return cs.Code is in 
}

type cycleContext[M any, T Set[M, T]] struct {

}

// CyclePath finds cycles
func CyclePath() {
	//return ShortestPathNonUnique()
}*/