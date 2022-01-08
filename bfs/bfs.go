package bfs

import (
	"container/heap"
	"fmt"
)

// Breadth first search stuff
type Context[M, T any] struct {
	GlobalContext M
	StateValue    StateValue[M, T]
}

type StateValue[M, T any] interface {
	State() T
	Dist() int
	Prev() StateValue[M, T]
}

type pathable[M, T, AS any] interface {
	Code(*Context[M, T]) string
	Done(*Context[M, T]) bool
	AdjacentStates(*Context[M, T]) []AS
}

/*type pathOrPop[T any] struct {
	t T
	pop bool
	popCode string
}

/*func AnyPath[M any, T State[M, T]](initState T, globalContext M) []T {
	pops := []pathOrPop[T]{{initState, false}}
	path := []T{}
	inPath := map[string]bool{}
	for len(pops) > 0 {
		pop := pops[len(path)-1]
		pops = pops[:len(pops)-1]

		if pop.pop {
			path = path[:len(path)-1]
			delete(inPath, pop.popCode)
			continue
		}

		state := pop.t
		path = append(path, state)
		inPath[state.Code()] = true

		if state.Done(globalContext) {
			return path
		}

		pops = append(pops, &pathOrPop{nil, true})
		for _, adjState := range state.AdjacentStates(globalContext) {
			pops = append(pops, &pathOrPop{adjState, false})
		}
	}
}*/

type pathHelper[M, T, AS any] struct {
	distFunc   func(*Context[M, T], AS) int
	convFunc   func(*Context[M, T], AS) T
	skipUnique bool
}

func shortestPath[M any, AS any, T pathable[M, T, AS]](initState T, initDist int, globalContext M, ph *pathHelper[M, T, AS]) ([]T, int) {
	ctx := &Context[M, T]{
		GlobalContext: globalContext,
	}
	states := &stateSet[M, T]{}
	states.Push(&stateValue[M, T]{initState, initDist, nil})

	checked := map[string]bool{}

	for states.Len() > 0 {
		sv := heap.Pop(states).(*stateValue[M, T])
		ctx.StateValue = sv
		if !ph.skipUnique {
			if code := sv.state.Code(ctx); checked[code] {
				continue
			} else {
				checked[code] = true
			}
		}

		if sv.state.Done(ctx) {
			var path []T
			for cur := sv; cur != nil; cur = cur.prev {
				path = append(path, cur.state)
			}
			return path, sv.dist
		}

		for _, adjState := range sv.state.AdjacentStates(ctx) {
			dist := ph.distFunc(ctx, adjState)
			newT := ph.convFunc(ctx, adjState)
			heap.Push(states, &stateValue[M, T]{newT, dist, sv})
		}
	}
	return nil, -1
}

type stateSet[M, T any] struct {
	values []*stateValue[M, T]
}

func (ss *stateSet[M, T]) Len() int {
	return len(ss.values)
}

func (ss *stateSet[M, T]) Less(i, j int) bool {
	return ss.values[i].dist < ss.values[j].dist
}

func (ss *stateSet[M, T]) Push(x interface{}) {
	ss.values = append(ss.values, x.(*stateValue[M, T]))
}

func (ss *stateSet[M, T]) Pop() interface{} {
	r := ss.values[len(ss.values)-1]
	ss.values = ss.values[:len(ss.values)-1]
	return r
}

func (ss *stateSet[M, T]) Swap(i, j int) {
	tmp := ss.values[i]
	ss.values[i] = ss.values[j]
	ss.values[j] = tmp
}

type stateValue[M, T any] struct {
	state T
	dist  int
	prev  *stateValue[M, T]
}

func (sv *stateValue[M, T]) String() string {
	return fmt.Sprintf("(%d) %v", sv.dist, sv.state)
}

func (sv *stateValue[M, T]) State() T {
	return sv.state
}

func (sv *stateValue[M, T]) Dist() int {
	return sv.dist
}

func (sv *stateValue[M, T]) Prev() StateValue[M, T] {
	return sv.prev
}
