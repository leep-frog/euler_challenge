package bfs

type biconverter[T, T2 any] interface {
	To(T) T2
	From(T2) T
}

func toSlice[T, T2 any](b biconverter[T, T2], ts []T) []T2 {
	var t2s []T2
	for _, t := range ts {
		t2s = append(t2s, b.To(t))
	}
	return t2s
}

func fromSlice[T, T2 any](b biconverter[T, T2], t2s []T2) []T {
	var ts []T
	for _, t := range t2s {
		ts = append(ts, b.From(t))
	}
	return ts
}

/*func DFSWithContext[M any, T DepthSearchableWithContext[M, T]](initStates []T, m M, opts ...DFSOption) []T {
	var r []*addProcessWrapperDFS[M,*addPathWrapperDFS[M, T]]
	for _, is := range initStates {
		r = append(r, &addProcessWrapperDFS[M, *addPathWrapperDFS[M, T]]{&addPathWrapperDFS[M, T]{is}})
	}
	//c := func(apw *addPathWrapperDFS[M, T]) T { return apw.state}
	c := func(apw *addProcessWrapperDFS[M, *addPathWrapperDFS[M, T]]) T { return apw.T.state }
	return dfsFinal(r, m, c, opts...)
}

func DFS[M any, T DepthSearchable[T]](initStates []T, opts ...DFSOption) []T {
	var r []*addContextAndPathWrapperDFS[T]
	for _, is := range initStates {
		r = append(r, &addContextAndPathWrapperDFS[T]{is})
	}
	c := func(apw *addContextAndPathWrapperDFS[T]) T { return apw.state}
	return dfsFinal(r, 0, c, opts...)
}

func DFSWithContextAndPath[M any, T DepthSearchableWithContextAndPath[M, T]](initStates []T, m M, opts ...DFSOption) []T {
	return dfsFinal(initStates, m, identityConverter[T](), opts...)
}

type addContextAndPathWrapperDFS[T DepthSearchable[T]] struct {
	state T
}

func (acp *addContextAndPathWrapperDFS[T]) Code(_ int, _ DFSPath[T]) string {
	return acp.state.Code()
}

func (acp *addContextAndPathWrapperDFS[T]) Done(_ int, _ DFSPath[T]) bool {
	return acp.state.Done()
}

func (acp *addContextAndPathWrapperDFS[T]) AdjacentStates(_ int, _ DFSPath[T]) []*addContextAndPathWrapperDFS[T] {
	var r []*addContextAndPathWrapperDFS[T]
	for _, n := range acp.state.AdjacentStates() {
		r = append(r, &addContextAndPathWrapperDFS[T]{n})
	}
	return r
}

type addPathWrapperDFS[M any, T DepthSearchableWithContext[M, T]] struct {
	state T
}

func (apw *addPathWrapperDFS[M, T]) Code(m M, _ DFSPath[T]) string {
	return apw.state.Code(m)
}

func (apw *addPathWrapperDFS[M, T]) Done(m M, _ DFSPath[T]) bool {
	return apw.state.Done(m)
}

func (apw *addPathWrapperDFS[M, T]) AdjacentStates(m M, _ DFSPath[T]) []*addPathWrapperDFS[M, T] {
	var r []*addPathWrapperDFS[M, T]
	for _, n := range apw.state.AdjacentStates(m) {
		r = append(r, &addPathWrapperDFS[M, T]{n})
	}
	return r
}

type processWrapperDFS[M any, T DepthSearchable[M, T]] struct {
	state T
}

func (apw *addProcessWrapperDFS[M, T]) PreProcess(M, DFSPath[T]) { }
func (apw *addProcessWrapperDFS[M, T]) PostProcess(M, DFSPath[T]) {}*/

func identityConverter[T any]() converter[T, T] {
	return func(t T) T { return t }
}

type dfsAction[T any] struct {
	popPath *string
	state   T
}

func dfsFinal[M, T2 any, T completeDepthSearcher[M, T, T2]](initStates []T, m M, convert converter[T, T2], opts ...DFSOption) []T2 {
	opt := &dfsOption{}
	for _, o := range opts {
		o(opt)
	}
	var actions []*dfsAction[T]
	for _, is := range initStates {
		actions = append(actions, &dfsAction[T]{nil, is})
	}

	checkedNodes := map[string]bool{}

	dp := &dfsPath[T2]{nil, map[string]bool{}}

	for len(actions) > 0 {
		a := actions[len(actions)-1]
		actions = actions[:len(actions)-1]

		// Remove from path if relevant.
		if a.popPath != nil {
			// TODO: move OnPop into dfsPath implementation
			dp.pop(*a.popPath)
			a.state.OnPop(m, dp)
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
		dp.push(convert(state), c)
		state.OnPush(m, dp)
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
