package bfs

import (
	"container/heap"
	"fmt"
)

// Breadth first search stuff
type AdjacentState[M, T any] struct {
	State  OffsetState[M, T]
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
	Done(M) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(M) []T
}

type OffsetState[M, T any] interface {
	Code() string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(M) bool
	// Returns all pairs of the adjacent states and those states offsets from this state.
	// The input is a contextual variable that is passed along from ShortestPath.
	AdjacentStates(M) []*AdjacentState[M, T]
	// 
}

// offsetState is a type that converts an OffsetState interface to a State one.
type offsetState[M, T any] struct {
	os   OffsetState[M, T]
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

func (os *offsetState[M, T]) Done(m M) bool {
	return os.os.Done(m)
}

func (os *offsetState[M, T]) AdjacentStates(m M) []*offsetState[M, T] {
	var r []*offsetState[M, T]
	for _, as := range os.os.AdjacentStates(m) {
		r = append(r, &offsetState[M, T]{as.State, os.dist + as.Offset})
	}
	return r
}

func ShortestOffsetPath[M any, T OffsetState[M, T]](initState T, initDist int, globalContext M) (T, int) {
	state := &offsetState[M, T]{initState, initDist}
	r, d := ShortestPath[M, *offsetState[M, T]](state, globalContext)
	return r.os.(T), d
}

func ShortestPath[M any, T State[M, T]](initState T, globalContext M) (T, int) {
	states := &stateSet[M, T]{}
	states.Push(&stateValue[M, T]{initState, initState.Distance(globalContext)})

	checked := map[string]bool{}

	for states.Len() > 0 {
		sv := heap.Pop(states).(*stateValue[M, T])
		if code := sv.state.Code(); checked[code] {
			continue
		} else {
			checked[code] = true
		}

		if sv.state.Done(globalContext) {
			return sv.state, sv.dist
		}

		for _, adjState := range sv.state.AdjacentStates(globalContext) {
			heap.Push(states, &stateValue[M, T]{adjState, adjState.Distance(globalContext)})
		}
	}
	var nill T
	return nill, -1
}

type stateSet[M any, T State[M, T]] []*stateValue[M, T]

func (ss *stateSet[M, T]) Len() int {
	return len(*ss)
}

func (ss *stateSet[M, T]) Less(i, j int) bool {
	return (*ss)[i].dist < (*ss)[j].dist
}

func (ss *stateSet[M, T]) Push(x interface{}) {
	*ss = append(*ss, x.(*stateValue[M, T]))
}

func (ss *stateSet[M, T]) Pop() interface{} {
	r := (*ss)[len(*ss)-1]
	*ss = (*ss)[:len(*ss)-1]
	return r
}

func (ss *stateSet[M, T]) Swap(i, j int) {
	tmp := (*ss)[i]
	(*ss)[i] = (*ss)[j]
	(*ss)[j] = tmp
}

type stateValue[M any, T State[M, T]] struct {
	state T
	dist  int
}

func (sv *stateValue[M, T]) String() string {
	return fmt.Sprintf("(%d) %v", sv.dist, sv.state)
}
