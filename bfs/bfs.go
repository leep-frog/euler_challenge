package bfs

import (
	"container/heap"
	"fmt"
	"strings"
)

type Distanceable[T any] interface {
	LT(T) bool
	Plus(T) T
}

// TODO: Add peek function
type Path[T any] interface {
	Fetch() []T
	Len() int
}

type converter[F, T any] func(F) T

func joinConverters[A, B, C any](aToB converter[A, B], bToC converter[B, C]) converter[A, C] {
	return func(a A) C {
		return bToC(aToB(a))
	}
}

func (c converter[F, T]) convertPath(p Path[F]) []T {
	if p == nil {
		return nil
	}
	return c.convertSlice(p.Fetch())
}

func (c converter[F, T]) convert(f F) T {
	return c(f)
}

func (c converter[F, T]) convertSlice(fs []F) []T {
	var ts []T
	for _, f := range fs {
		ts = append(ts, c(f))
	}
	return ts
}

type pathWrapper[T, T2 any] struct {
	path      Path[T]
	converter converter[T, T2]
}

func (p *pathWrapper[T, T2]) String() string {
	return fmt.Sprintf("%v", p.path)
}

func (p *pathWrapper[T, T2]) Len() int {
	return p.path.Len()
}

func (p *pathWrapper[T, T2]) Fetch() []T2 {
	return p.converter.convertSlice(p.path.Fetch())
}

type path[DIST Distanceable[DIST], T any] struct {
	tail *StateValue[DIST, T]
}

func (p *path[DIST, T]) String() string {
	return p.tail.PathString()
}

func (p *path[DIST, T]) Len() int {
	if p.tail == nil {
		return 0
	}
	return p.tail.pathLen
}

func (p *path[DIST, T]) Fetch() []T {
	var r []T
	for cur := p.tail; cur != nil; cur = cur.Prev() {
		r = append(r, cur.state)
	}
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return r
}

type Option func(o *option)

type option struct {
	checkDuplicates            bool
	cumulativeDistanceFunction bool
}

func CheckDuplicates() Option {
	return func(o *option) {
		o.checkDuplicates = true
	}
}

// CumulativeDistanceFunction is an `Option` that indicates the node's `Distance`
// function returns the cumulative distance for that node (as opposed to the default
// behavior which assumes the function just returns the edge distance to get to that node
// from the previous node).
func CumulativeDistanceFunction() Option {
	return func(o *option) {
		o.cumulativeDistanceFunction = true
	}
}

func search[RETURN, CTX any, CODE comparable, DIST Distanceable[DIST], T ContextDistancePathNode[CTX, CODE, DIST, T]](ctx CTX, initStates []T, tConverter converter[T, RETURN], opts ...Option) ([]RETURN, DIST) {
	o := &option{}
	for _, opt := range opts {
		opt(o)
	}

	nodes := &bfsHeap[DIST, T]{}
	for _, is := range initStates {
		nodes.PushState(&StateValue[DIST, T]{is, is.Distance(ctx, &path[DIST, T]{nil}), nil, 1})
	}

	checked := map[CODE]bool{}

	for nodes.Len() > 0 {
		sv := nodes.PopState()
		p := &path[DIST, T]{sv}
		if !o.checkDuplicates {
			if code := sv.state.Code(ctx, p); checked[code] {
				continue
			} else {
				checked[code] = true
			}
		}

		if sv.state.Done(ctx, p) {
			pw := &pathWrapper[T, RETURN]{p, tConverter}
			return pw.Fetch(), sv.dist
		}

		for _, neighbor := range sv.state.AdjacentStates(ctx, p) {
			dist := neighbor.Distance(ctx, p)
			if !o.cumulativeDistanceFunction {
				dist = dist.Plus(sv.dist)
			}
			nodes.PushState(&StateValue[DIST, T]{
				neighbor,
				dist,
				func() *StateValue[DIST, T] { return sv },
				sv.pathLen + 1,
			})
		}
	}
	var nill DIST
	return nil, nill
}

type bfsHeap[DIST Distanceable[DIST], T any] struct {
	values []*StateValue[DIST, T]
}

func (bh *bfsHeap[DIST, T]) PushState(sv *StateValue[DIST, T]) {
	heap.Push(bh, sv)
}

func (bh *bfsHeap[DIST, T]) PopState() *StateValue[DIST, T] {
	return heap.Pop(bh).(*StateValue[DIST, T])
}

func (bh *bfsHeap[DIST, T]) popState() *StateValue[DIST, T] {
	return heap.Pop(bh).(*StateValue[DIST, T])
}

// Below functions needed for heap interface
func (bh *bfsHeap[DIST, T]) Len() int {
	return len(bh.values)
}

func (bh *bfsHeap[DIST, T]) Less(i, j int) bool {
	return bh.values[i].dist.LT(bh.values[j].dist)
}

func (bh *bfsHeap[DIST, T]) Push(x interface{}) {
	bh.values = append(bh.values, x.(*StateValue[DIST, T]))
}

func (bh *bfsHeap[DIST, T]) Pop() interface{} {
	r := bh.values[len(bh.values)-1]
	bh.values = bh.values[:len(bh.values)-1]
	return r
}

func (bh *bfsHeap[DIST, T]) Swap(i, j int) {
	bh.values[i], bh.values[j] = bh.values[j], bh.values[i]
}

type StateValue[DIST Distanceable[DIST], T any] struct {
	state T
	// This can be replaced by wrapping type for specific search type wrapper
	dist DIST
	// TODO: this can be replaced by improving container method (to include CurrentPath() function)
	prev func() *StateValue[DIST, T]
	//
	pathLen int
}

func (sv *StateValue[DIST, T]) PathString() string {
	return strings.Join(append(sv.path()), ", ")
}

func (sv *StateValue[DIST, T]) path() []string {
	if sv == nil {
		return []string{}
	}
	return append(sv.Prev().path(), sv.String())
}

func (sv *StateValue[DIST, T]) String() string {
	return fmt.Sprintf("(%v) %v", sv.dist, sv.state)
}

func (sv *StateValue[DIST, T]) State() T {
	return sv.state
}

func (sv *StateValue[DIST, T]) Dist() DIST {
	return sv.dist
}

func (sv *StateValue[DIST, T]) Prev() *StateValue[DIST, T] {
	if sv.prev == nil {
		return nil
	}
	return sv.prev()
}
