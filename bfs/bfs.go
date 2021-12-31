package bfs

import (
	"container/heap"
	"fmt"
)

// Breadth first search stuff
type AdjacentState struct {
	State  OffsetState
	Offset int
}

// Interface used for automatically running BFS
type State interface {
	// A unique code that represents the current state of the world.
	Code() string
	// Distance returns the total distance for the given state. The input is a contextual variable
	// that is passed along from ShortestPath.
	Distance(interface{}) int
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(interface{}) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	AdjacentStates(interface{}) []State
}

type OffsetState interface {
	Code() string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(interface{}) bool
	// Returns all pairs of the adjacent states and those states offsets from this state.
	// The input is a contextual variable that is passed along from ShortestPath.
	AdjacentStates(interface{}) []*AdjacentState
}

type offsetState struct {
	os   OffsetState
	dist int
}

func (os *offsetState) String() string {
	return fmt.Sprintf("(%d) %v", os.dist, os.os)
}

func (os *offsetState) Code() string {
	return os.os.Code()
}

func (os *offsetState) Distance(interface{}) int {
	return os.dist
}

func (os *offsetState) Done(i interface{}) bool {
	return os.os.Done(i)
}

func (os *offsetState) AdjacentStates(i interface{}) []State {
	var r []State
	for _, as := range os.os.AdjacentStates(i) {
		r = append(r, &offsetState{as.State, os.dist + as.Offset})
	}
	return r
}

func ShortestOffsetPath(initState OffsetState, initDist int, globalContext interface{}) (OffsetState, int) {
	r, d := ShortestPath(&offsetState{initState, initDist}, globalContext)
	return r.(*offsetState).os, d
}

func ShortestPath(initState State, globalContext interface{}) (State, int) {
	states := &stateSet{}
	states.Push(&stateValue{initState, initState.Distance(globalContext)})

	checked := map[string]bool{}

	for states.Len() > 0 {
		sv := heap.Pop(states).(*stateValue)
		if code := sv.state.Code(); checked[code] {
			continue
		} else {
			checked[code] = true
		}

		if sv.state.Done(globalContext) {
			return sv.state, sv.dist
		}

		for _, adjState := range sv.state.AdjacentStates(globalContext) {
			heap.Push(states, &stateValue{adjState, adjState.Distance(globalContext)})
		}
	}
	return nil, -1
}

type stateSet []*stateValue

func (ss *stateSet) Len() int {
	return len(*ss)
}

func (ss *stateSet) Less(i, j int) bool {
	return (*ss)[i].dist < (*ss)[j].dist
}

func (ss *stateSet) Push(x interface{}) {
	*ss = append(*ss, x.(*stateValue))
}

func (ss *stateSet) Pop() interface{} {
	r := (*ss)[len(*ss)-1]
	*ss = (*ss)[:len(*ss)-1]
	return r
}

func (ss *stateSet) Swap(i, j int) {
	tmp := (*ss)[i]
	(*ss)[i] = (*ss)[j]
	(*ss)[j] = tmp
}

type stateValue struct {
	state State
	dist  int
}

func (sv *stateValue) String() string {
	return fmt.Sprintf("(%d) %v", sv.dist, sv.state)
}
