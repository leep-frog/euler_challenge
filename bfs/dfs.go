package bfs

type popCtx[M, T any] struct {
	pop bool
	stateValue *stateValue[M, T]
}

type anyPathHelper[M, T any] struct {
	popFunc func(*Context[M, T], T)
	pushFunc func(*Context[M, T], T)
}

/*type CycleState[M, T any] interface {
	Code(*Context[M, T]) string
	AdjacentStates(*Context[M, T]) []T
}

func CyclePath[map[string]bool, T State[map[string]bool, T]](initState T) ([]T, int) {
	return anyPath(
		initState,
		 0,
		  globalContext,
			 &anyPathHelper[M, T]{
				popFunc
			 },
			)
}*/

func AnyPath[M, T State[M, T]](initState T, globalContext M) ([]T, int) {
	return anyPath(initState, 0, globalContext, &anyPathHelper[M, T]{})
}

// anyPath implements a generic depth first search.
func anyPath[M any, T pathable[M, T, T]](initState T, initDist int, globalContext M, aph *anyPathHelper[M, T]) ([]T, int) {
	ctx := &Context[M, T]{
		GlobalContext: globalContext,
	}
	states := []*popCtx[M, T]{{
		pop: false,
		stateValue: &stateValue[M, T]{initState, initDist, nil},
	}}
	
	for len(states) > 0 {
		svp := states[len(states)-1]
		states = states[:len(states)-1]
		sv := svp.stateValue
		if svp.pop {
			aph.popFunc(ctx, sv.state)
			continue
		}
		aph.pushFunc(ctx, sv.state)
		
		ctx.StateValue = sv
		if sv.state.Done(ctx) {
			var path []T
			// TODO: make function on stateValue
			for cur := sv; cur != nil; cur = cur.prev {
				path = append(path, cur.state)
			}
			return path, sv.dist
		}

		states = append(states, &popCtx[M, T]{true, sv})
		for _, adjState := range sv.state.AdjacentStates(ctx) {
			states = append(states, &popCtx[M, T]{false, &stateValue[M, T]{adjState, sv.dist + 1, sv}})
		}
	}
	return nil, -1
}
