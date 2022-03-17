package bfs

type DepthState[M, T any] interface {
	State[M, T]
	/*Preprocess(M)
	Postprocess(M)*/
}

// AnyPath depth-first searches for any valid path. This can also be used to
// search for all paths by having the Done method always return false.
func AnyPath[M any, T State[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: simpleDistFunc[M, T](),
		convFunc: identityConvFunc[M, T](),
	}
	return searchPath(newDFSSearcher[T](), initStates, nil, globalContext, ph)
}

type dfsAction[T any] struct {
	popPath *string
	state T	
}

func DFS[M any, T State[M, T]](initStates []T, globalContext M) []T {
	ctx := &Context[M, T]{
		GlobalContext: globalContext,
	}

	var actions []*dfsAction[T]
	for _, is := range initStates {
		actions = append(actions, &dfsAction[T]{nil, is})
	}

	// TODO(option): Check each node at most once vs check unique path at most once
	// For example if A and B both have edge to C, do we want to check
	// C once, or check once for A->C and once for A->C

	checkedNodes := map[string]bool{}
	inPath := map[string]bool{}
	path := []T{}

	for len(actions) > 0 {
		a := actions[len(actions)-1]
		actions = actions[:len(actions)-1]

		// Remove from path if relevant.
		if a.popPath != nil{
			path = path[:len(path)-1]
			delete(inPath, *a.popPath)
			// TODO: postprocess state
			continue
		}

		// Otherwise, process the state
		state := a.state
		c := state.Code(ctx)

		// See if we've visited this node before
		if inPath[c] {
			// TODO: cycleHandler
			continue
		}
		if checkedNodes[c] {
			// TODO: alreadyVisitedHandler
			continue
		}

		// Pop once we examine all of this node's children
		actions = append(actions, &dfsAction[T]{&c, state})

		// Update path variables
		inPath[c] = true
		checkedNodes[c] = true
		path = append(path, state)

		// Check if node is done
		if state.Done(ctx) {
			return path
		}

		// TODO: preprocess state
		for _, neighbor := range state.AdjacentStates(ctx) {
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