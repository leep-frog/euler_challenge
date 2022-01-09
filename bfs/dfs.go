package bfs

type popCtx[T any] struct {
	pop        bool
	stateValue *StateValue[T]
}

type anyPathHelper[M, T, AS any] struct {
	popFunc  func(*Context[M, T], T)
	pushFunc func(*Context[M, T], T)
	ph       *pathHelper[M, T, AS]
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

/*func AnyPath[M, T State[M, T]](initState T, globalContext M) ([]T, int) {
	return anyPath(
		initState,
		 0,
		  globalContext,
			 &anyPathHelper[M, T]{
				 ph:
			 })
}*/

type dfsPathable[M, T, AS any] interface {
	Done(*Context[M, T]) bool
	AdjacentStates(*Context[M, T]) []AS
}

// anyPath implements a generic depth first search.
func anyPath[M, AS any, T dfsPathable[M, T, AS]](initState T, initDist int, globalContext M, aph *anyPathHelper[M, T, AS]) ([]T, int) {
	ctx := &Context[M, T]{
		GlobalContext: globalContext,
	}
	states := []*popCtx[T]{{
		pop:        false,
		stateValue: &StateValue[T]{initState, initDist, nil},
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
			for cur := sv; cur != nil; cur = cur.Prev() {
				path = append(path, cur.state)
			}
			return path, sv.dist
		}

		states = append(states, &popCtx[T]{true, sv})
		for _, adjState := range sv.state.AdjacentStates(ctx) {
			dist := aph.ph.distFunc(ctx, adjState)
			newT := aph.ph.convFunc(ctx, adjState)
			states = append(states, &popCtx[T]{false, &StateValue[T]{newT, dist, func() *StateValue[T] { return sv }}})
		}
	}
	return nil, -1
}
