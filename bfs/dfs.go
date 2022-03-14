package bfs

import (
	"fmt"
	"strings"
)

// TODO: this file still needs work

func permaImp() {
	_ = strings.Join(nil, "")
}

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

/*type dfsPathable[M, T, AS any] interface {
	Done(*Context[M, T]) bool
	AdjacentStates(*Context[M, T]) []AS
}*/

func AnyPath[M any, T State[M, T]](initStates []T, globalContext M) ([]T, int) {
	ph := &pathHelper[M, T, T]{
		distFunc: simpleDistFunc[M, T](),
		convFunc: identityConvFunc[M, T](),
	}
	return searchPath(newDFSSearcher[T](), initStates, nil, globalContext, ph)
}

// anyPath implements a generic depth first search.
func anyPath[M, AS any, T pathable[M, T, AS]](initState T, initDist int, globalContext M, aph *anyPathHelper[M, T, AS]) ([]T, int) {
	ctx := &Context[M, T]{
		GlobalContext: globalContext,
	}
	states := []*popCtx[T]{{
		pop:        false,
		stateValue: &StateValue[T]{initState, initDist, nil},
	}}

	checked := map[string]bool{}

	count := 0
	for len(states) > 0 {
		count++
		if count > 1000 {
			return nil, 777
		}
		svp := states[len(states)-1]
		states = states[:len(states)-1]
		sv := svp.stateValue
		if svp.pop {
			fmt.Println("POPPING", sv)
			aph.popFunc(ctx, sv.state)
			continue
		}
		ctx.StateValue = sv
		// Check done before running checked because sometimes we look for cycles
		if sv.state.Done(ctx) {
			var path []T
			// TODO: make function on stateValue
			fmt.Println("yup", sv)
			for cur := sv; cur != nil; cur = cur.Prev() {
				path = append(path, cur.state)
			}
			return path, sv.dist
		}
		if code := sv.State().Code(ctx); checked[code] {
			continue
		} else {
			checked[code] = true
		}

		var path []string
		for cur := sv; cur != nil; cur = cur.Prev() {
			path = append(path, fmt.Sprintf("%v", cur.state))
		}
		fmt.Println("CHECKING", sv)
		fmt.Println(strings.Join(path, ", "))
		fmt.Println(ctx.GlobalContext)
		fmt.Println("------")
		//fmt.Println("CHECKING", sv)

		states = append(states, &popCtx[T]{true, sv})
		for _, adjState := range sv.state.AdjacentStates(ctx) {
			dist := aph.ph.distFunc(ctx, adjState)
			newT := aph.ph.convFunc(ctx, adjState)
			fmt.Println("ADDING", newT)
			states = append(states, &popCtx[T]{false, &StateValue[T]{newT, dist, func() *StateValue[T] { return sv }}})
		}

		// Push after we check for done and get the adjacent states
		aph.pushFunc(ctx, sv.state)
	}
	return nil, -1
}

// TODO: make this a separate package
// TODO: add this in AnyPath to include search path
// node for linked list
/*type node[T any] struct {
	value T
	next *node[T]
}*/

// AnyPath implements depth first search
/*func AnyPath[M any, T State[M, T]](initStates []T, globalContext M) T {
	checked := map[string]bool{}
	for len(initStates) > 0 {
		cur := initStates[0]
		initStates = initStates[1:]
		if checked[cur].Code() {
			continue
		}
	}

	return nil
}
*/