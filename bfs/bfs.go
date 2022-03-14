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

func searchPath[M, AS any, T pathable[M, T, AS]](container stateContainer[T], initStates []T, initDistFunc func(*Context[M, T], T) int, globalContext M, ph *pathHelper[M, T, AS]) ([]T, int) {
	ctx := &Context[M, T]{
		GlobalContext: globalContext,
	}
	//states := &stateSet[T]{}
	for _, initState := range initStates {
		var initDist int
		if initDistFunc != nil {
			initDist = initDistFunc(ctx, initState)
		}
		container.PushState(&StateValue[T]{initState, initDist, nil})
	}

	checked := map[string]bool{}

	for container.Len() > 0 {
		sv := container.PopState()
		//sv := heap.Pop(states).(*StateValue[T])
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
			container.PushState(&StateValue[T]{newT, dist, func() *StateValue[T] { return sv }})
		}
	}
	return nil, -1
}

type stateContainer[T any] interface {
	PushState(sv *StateValue[T])
  PopState() *StateValue[T]
	Len() int
}

// newBFSSearcher returns a searcher for depth first search.
func newDFSSearcher[T any]() stateContainer[T] {
	return &dfsSearcher[T]{}
}

type dfsSearcher[T any] struct {
	stack []*StateValue[T]
}

func (ds *dfsSearcher[T]) PushState(sv *StateValue[T]) {
	ds.stack = append(ds.stack, sv)
}

func (ds *dfsSearcher[T]) PopState() *StateValue[T] {
	r := ds.stack[len(ds.stack)-1]
	ds.stack = ds.stack[:len(ds.stack)-1]
	return r
}

func (ds *dfsSearcher[T]) Len() int {
	return len(ds.stack)
}

// newBFSSearcher returns a searcher for breadth first search.
func newBFSSearcher[T any]() stateContainer[T] {
	return &bfsSearcher[T]{}
}

type bfsSearcher[T any] struct {
	values []*StateValue[T]
}

func (bs *bfsSearcher[T]) PushState(sv *StateValue[T]) {
	heap.Push(bs, sv)
}

func (bs *bfsSearcher[T]) PopState() *StateValue[T] {
	return heap.Pop(bs).(*StateValue[T])
}

// Below functions needed for heap interface
func (bs *bfsSearcher[T]) Len() int {
	return len(bs.values)
}

func (bs *bfsSearcher[T]) Less(i, j int) bool {
	return bs.values[i].dist < bs.values[j].dist
}

func (bs *bfsSearcher[T]) Push(x interface{}) {
	bs.values = append(bs.values, x.(*StateValue[T]))
}

func (bs *bfsSearcher[T]) Pop() interface{} {
	r := bs.values[len(bs.values)-1]
	bs.values = bs.values[:len(bs.values)-1]
	return r
}

func (bs *bfsSearcher[T]) Swap(i, j int) {
	tmp := bs.values[i]
	bs.values[i] = bs.values[j]
	bs.values[j] = tmp
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
