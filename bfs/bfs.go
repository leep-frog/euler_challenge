package bfs

import (
	"container/heap"
	"fmt"
)

// Breadth first search stuff
type AdjacentState[M any] struct {
	State  OffsetState[M]
	Offset int
}

// TODO: use parameters for this
// Interface used for automatically running BFS
type State[M any] interface {
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
	AdjacentStates(M) []State[M]
}

type OffsetState[M any] interface {
	Code() string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(M) bool
	// Returns all pairs of the adjacent states and those states offsets from this state.
	// The input is a contextual variable that is passed along from ShortestPath.
	AdjacentStates(M) []*AdjacentState[M]
	// 
}

// offsetState is a type that converts an OffsetState interface to a State one.
type offsetState[M any] struct {
	os   OffsetState[M]
	dist int
}

func (os *offsetState[M]) String() string {
	return fmt.Sprintf("(%d) %v", os.dist, os.os)
}

func (os *offsetState[M]) Code() string {
	return os.os.Code()
}

func (os *offsetState[M]) Distance(M) int {
	return os.dist
}

func (os *offsetState[M]) Done(m M) bool {
	return os.os.Done(m)
}

func (os *offsetState[M]) AdjacentStates(m M) []State[M] {
	var r []State[M]
	for _, as := range os.os.AdjacentStates(m) {
		r = append(r, &offsetState[M]{as.State, os.dist + as.Offset})
	}
	return r
}

func ShortestOffsetPath[M any, T OffsetState[M]](initState T, initDist int, globalContext M) (State[M], int) {
	var state State[M]
	state = &offsetState[M]{initState, initDist}
	r, d := ShortestPath[M](state, globalContext)
	return r, d
}

func ShortestPath[M any](initState State[M], globalContext M) (State[M], int) {
	states := &stateSet[M]{}
	states.Push(&stateValue[M]{initState, initState.Distance(globalContext)})

	checked := map[string]bool{}

	for states.Len() > 0 {
		sv := heap.Pop(states).(*stateValue[M])
		if code := sv.state.Code(); checked[code] {
			continue
		} else {
			checked[code] = true
		}

		if sv.state.Done(globalContext) {
			return sv.state, sv.dist
		}

		for _, adjState := range sv.state.AdjacentStates(globalContext) {
			heap.Push(states, &stateValue[M]{adjState, adjState.Distance(globalContext)})
		}
	}
	var nill State[M]
	return nill, -1
}

type stateSet[M any] []*stateValue[M]

func (ss *stateSet[M]) Len() int {
	return len(*ss)
}

func (ss *stateSet[M]) Less(i, j int) bool {
	return (*ss)[i].dist < (*ss)[j].dist
}

func (ss *stateSet[M]) Push(x interface{}) {
	*ss = append(*ss, x.(*stateValue[M]))
}

func (ss *stateSet[M]) Pop() interface{} {
	r := (*ss)[len(*ss)-1]
	*ss = (*ss)[:len(*ss)-1]
	return r
}

func (ss *stateSet[M]) Swap(i, j int) {
	tmp := (*ss)[i]
	(*ss)[i] = (*ss)[j]
	(*ss)[j] = tmp
}

type stateValue[M any] struct {
	state State[M]
	dist  int
}

func (sv *stateValue[M]) String() string {
	return fmt.Sprintf("(%d) %v", sv.dist, sv.state)
}
