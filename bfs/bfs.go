package bfs

import (
	"container/heap"
	"fmt"
	"reflect"
	"strings"
)

var (
	debug = false
)

// Breadth first search stuff
type Context[M, T any] struct {
	GlobalContext M
	StateValue    *StateValue[T]
}

type pathable[M, T, AS any] interface {
	Code(*Context[M, T]) string
	Done(*Context[M, T]) bool
	AdjacentStates(*Context[M, T]) []AS
}

type pathHelper[M, T, AS any] struct {
	distFunc   func(*Context[M, T], AS) int
	convFunc   func(*Context[M, T], AS) T
	skipUnique bool
}

func identityConvFunc[M, T any]() func(*Context[M, T], T) T {
	return func(_ *Context[M, T], as T) T {
		return as
	}
}

func adjStateDistFunc[M any, T OffsetState[M, T]]() func(*Context[M, T], T) int {
	return func(ctx *Context[M, T], os T) int {
		if ctx.StateValue == nil {
			return os.Offset(ctx)
		}
		return ctx.StateValue.Dist() + os.Offset(ctx)
	}
}

func simpleDistFunc[M, T any]() func(*Context[M, T], T) int {
	return func(ctx *Context[M, T], as T) int {
		if ctx.StateValue == nil {
			return 0
		}
		return ctx.StateValue.Dist() + 1
	}
}

func shortestPath[M, AS any, T pathable[M, T, AS]](initStates []T, initDistFunc func(*Context[M, T], T) int, globalContext M, ph *pathHelper[M, T, AS]) ([]T, int) {
	ctx := &Context[M, T]{
		GlobalContext: globalContext,
	}
	states := &stateSet[T]{}
	for _, initState := range initStates {
		var initDist int
		if initDistFunc != nil {
			initDist = initDistFunc(ctx, initState)
		}
		states.Push(&StateValue[T]{initState, initDist, nil})
	}

	checked := map[string]bool{}

	for states.Len() > 0 {
		sv := heap.Pop(states).(*StateValue[T])
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
			for cur := sv; cur != nil; cur = cur.Prev() {
				path = append(path, cur.state)
			}
			return path, sv.dist
		}

		for _, adjState := range sv.state.AdjacentStates(ctx) {
			dist := ph.distFunc(ctx, adjState)
			newT := ph.convFunc(ctx, adjState)
			heap.Push(states, &StateValue[T]{newT, dist, func() *StateValue[T] { return sv }})
		}
	}
	return nil, -1
}

type stateSet[T any] struct {
	values []*StateValue[T]
}

func (ss *stateSet[T]) Len() int {
	return len(ss.values)
}

func (ss *stateSet[T]) Less(i, j int) bool {
	return ss.values[i].dist < ss.values[j].dist
}

func (ss *stateSet[T]) Push(x interface{}) {
	ss.values = append(ss.values, x.(*StateValue[T]))
}

func (ss *stateSet[T]) Pop() interface{} {
	r := ss.values[len(ss.values)-1]
	ss.values = ss.values[:len(ss.values)-1]
	return r
}

func (ss *stateSet[T]) Swap(i, j int) {
	tmp := ss.values[i]
	ss.values[i] = ss.values[j]
	ss.values[j] = tmp
}

type StateValue[T any] struct {
	state T
	dist  int
	prev  func() *StateValue[T]
}

func (sv *StateValue[T]) PathString() string {
	return strings.Join(append(sv.path(), reflect.TypeOf(sv.state).String()), ", ")
}

func (sv *StateValue[T]) path() []string {
	if sv == nil {
		return []string{}
	}
	return append(sv.Prev().path(), sv.String())
}

func (sv *StateValue[T]) String() string {
	return fmt.Sprintf("(%d) %v", sv.dist, sv.state)
}

func (sv *StateValue[T]) State() T {
	return sv.state
}

func (sv *StateValue[T]) Dist() int {
	return sv.dist
}

func (sv *StateValue[T]) Prev() *StateValue[T] {
	if sv.prev == nil {
		return nil
	}
	return sv.prev()
}
