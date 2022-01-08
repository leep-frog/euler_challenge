package bfs

import (
	"container/heap"
	"fmt"
)

// Breadth first search stuff
type AdjacentState[T any] struct {
	State  T
	Offset int
}

// TODO: use parameters for this
// Interface used for automatically running BFS
type State[M, T any] interface {
	// A unique code that represents the current state of the world.
	Code() string
	// Distance returns the total distance for the given state. The input is a contextual variable
	// that is passed along from ShortestPath.
	Distance(M) int
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(*Context[M, T]) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(M) []T
}

type Context[M, T any] struct {
	GlobalContext M
	StateValue StateValue[M, T]
}

type StateValue[M, T any] interface {
	State() T
	Dist() int
	Prev() StateValue[M, T]
}

type OffsetState[M, T any] interface {
	Code() string
	// Returns if the given state is in a final position. The first input is a contextual variable
	// that is passed along from ShortestPath. The second input is the depth.
	Done(*Context[M, T]) bool
	// Returns all pairs of the adjacent states and those states offsets from this state.
	// The input is a contextual variable that is passed along from ShortestPath.
	AdjacentStates(M) []*AdjacentState[T]
	// 
}

// offsetState is a type that converts an OffsetState interface to a State one.
type offsetState[M any, T OffsetState[M, T]] struct {
	os   T
	dist int
}

func (os *offsetState[M, T]) String() string {
	return fmt.Sprintf("(%d) %v", os.dist, os.os)
}

func (os *offsetState[M, T]) Code() string {
	return os.os.Code()
}

func (os *offsetState[M, T]) Distance(M) int {
	return os.dist
}

func (os *offsetState[M, T]) Done(m *Context[M, *offsetState[M, T]]) bool {
	ctx := &Context[M, T]{
		GlobalContext: m.GlobalContext,
	}
	return os.os.Done(ctx)
}

func (os *offsetState[M, T]) AdjacentStates(m M) []*offsetState[M, T] {
	var r []*offsetState[M, T]
	for _, as := range os.os.AdjacentStates(m) {
		r = append(r, &offsetState[M, T]{as.State, os.dist + as.Offset})
	}
	return r
}

func ShortestOffsetPath[M any, T OffsetState[M, T]](initState T, initDist int, globalContext M) ([]T, int) {
	state := &offsetState[M, T]{initState, initDist}
	r, d := ShortestPath[M, *offsetState[M, T]](state, globalContext)
	var path []T
	for _, s := range r {
		path = append(path, s.os)
	}
	return path, d
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

func ShortestPath[M any, T State[M, T]](initState T, globalContext M) ([]T, int) {
	states := &stateSet[M, T]{}
	states.Push(&stateValue[M, T]{initState, initState.Distance(globalContext), nil})

	checked := map[string]bool{}
	ctx := &Context[M, T]{
		GlobalContext: globalContext,
	}

	for states.Len() > 0 {
		sv := heap.Pop(states).(*stateValue[M, T])
		ctx.StateValue = sv
		if code := sv.state.Code(); checked[code] {
			continue
		} else {
			checked[code] = true
		}

		if sv.state.Done(ctx) {
			var path []T
			for cur := sv; cur != nil; cur = cur.prev {
				path = append(path, cur.state)
			}
			return path, sv.dist
		}

		for _, adjState := range sv.state.AdjacentStates(globalContext) {
			heap.Push(states, &stateValue[M, T]{adjState, adjState.Distance(globalContext), sv})
		}
	}
	return nil, -1
}

type stateSet[M , T any] struct {
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
	prev *stateValue[M, T]
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
