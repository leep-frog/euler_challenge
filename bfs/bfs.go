package bfs

import (
	"container/heap"
	"fmt"
	"strings"

	"github.com/leep-frog/euler_challenge/maths"
)

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

type path[C maths.Comparable[C], T any] struct {
	tail *StateValue[C, T]
}

func (p *path[C, T]) String() string {
	return p.tail.PathString()
}

func (p *path[C, T]) Len() int {
	if p.tail == nil {
		return 0
	}
	return p.tail.pathLen
}

func (p *path[C, T]) Fetch() []T {
	var r []T
	for cur := p.tail; cur != nil; cur = cur.Prev() {
		r = append(r, cur.state)
	}
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return r
}

type Searchable[C maths.Comparable[C], T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code() string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done() bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates() []T
	// Distance is the total distance
	// TODO: make this return an heap-able interface
	Distance() C
}

type SearchableWithContext[C maths.Comparable[C], M, T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code(M) string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(M) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(M) []T
	// Distance is the total distance
	// TODO: make this return a mathable
	Distance(M) C
}

func toAPWConverter[C maths.Comparable[C], M any, T SearchableWithContext[C, M, T]]() converter[T, *addPathWrapper[C, M, T]] {
	return func(t T) *addPathWrapper[C, M, T] {
		return &addPathWrapper[C, M, T]{t}
	}
}

func fromAPWConverter[C maths.Comparable[C], M any, T SearchableWithContext[C, M, T]]() converter[*addPathWrapper[C, M, T], T] {
	return func(apw *addPathWrapper[C, M, T]) T {
		return apw.state
	}
}

type addPathWrapper[C maths.Comparable[C], M any, T SearchableWithContext[C, M, T]] struct {
	state T
}

func (apw *addPathWrapper[C, M, T]) String() string {
	return fmt.Sprintf("%v", apw.state)
}

func (apw *addPathWrapper[C, M, T]) Code(m M, _ Path[*addPathWrapper[C, M, T]]) string {
	return apw.state.Code(m)
}

func (apw *addPathWrapper[C, M, T]) Done(m M, _ Path[*addPathWrapper[C, M, T]]) bool {
	return apw.state.Done(m)
}

func (apw *addPathWrapper[C, M, T]) Distance(m M, _ Path[*addPathWrapper[C, M, T]]) C {
	return apw.state.Distance(m)
}

func (apw *addPathWrapper[C, M, T]) AdjacentStates(m M, _ Path[*addPathWrapper[C, M, T]]) []*addPathWrapper[C, M, T] {
	var r []*addPathWrapper[C, M, T]
	for _, n := range apw.state.AdjacentStates(m) {
		r = append(r, &addPathWrapper[C, M, T]{n})
	}
	return r
}

func toACPConverter[C maths.Comparable[C], T Searchable[C, T]]() converter[T, *addContextAndPathWrapper[C, T]] {
	return func(t T) *addContextAndPathWrapper[C, T] {
		return &addContextAndPathWrapper[C, T]{t}
	}
}

func fromACPConverter[C maths.Comparable[C], T Searchable[C, T]]() converter[*addContextAndPathWrapper[C, T], T] {
	return func(apw *addContextAndPathWrapper[C, T]) T {
		return apw.state
	}
}

type addContextAndPathWrapper[C maths.Comparable[C], T Searchable[C, T]] struct {
	state T
}

func (acp *addContextAndPathWrapper[C, T]) String() string {
	return fmt.Sprintf("%v", acp.state)
}

func (acp *addContextAndPathWrapper[C, T]) Code(_ int, _ Path[*addContextAndPathWrapper[C, T]]) string {
	return acp.state.Code()
}

func (acp *addContextAndPathWrapper[C, T]) Done(_ int, _ Path[*addContextAndPathWrapper[C, T]]) bool {
	return acp.state.Done()
}

func (acp *addContextAndPathWrapper[C, T]) Distance(_ int, _ Path[*addContextAndPathWrapper[C, T]]) C {
	return acp.state.Distance()
}

func (acp *addContextAndPathWrapper[C, T]) AdjacentStates(_ int, _ Path[*addContextAndPathWrapper[C, T]]) []*addContextAndPathWrapper[C, T] {
	var r []*addContextAndPathWrapper[C, T]
	for _, n := range acp.state.AdjacentStates() {
		r = append(r, &addContextAndPathWrapper[C, T]{n})
	}
	return r
}

type SearchableWithContextAndPath[C maths.Comparable[C], M, T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code(M, Path[T]) string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(M, Path[T]) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(M, Path[T]) []T
	// Distance is the total distance
	// TODO: make this return a mathable
	Distance(M, Path[T]) C
}

type Option func(o *option)

type option struct {
	checkDuplicates bool
}

// Change to DontCheckDuplicates?
func CheckDuplicates() Option {
	return func(o *option) {
		o.checkDuplicates = true
	}
}

func newSearch[C maths.Comparable[C], M any, T SearchableWithContextAndPath[C, M, T]](initStates []T, m M, opts ...Option) (Path[T], C) {
	o := &option{}
	for _, opt := range opts {
		opt(o)
	}

	nodes := &bfsHeap[C, T]{}
	for _, is := range initStates {
		nodes.PushState(&StateValue[C, T]{is, is.Distance(m, &path[C, T]{nil}), nil, 1})
	}

	checked := map[string]bool{}

	for nodes.Len() > 0 {
		sv := nodes.PopState()
		p := &path[C, T]{sv}
		if !o.checkDuplicates {
			if code := sv.state.Code(m, p); checked[code] {
				continue
			} else {
				checked[code] = true
			}
		}

		if sv.state.Done(m, p) {
			return p, sv.dist
		}

		for _, neighbor := range sv.state.AdjacentStates(m, p) {
			nodes.PushState(&StateValue[C, T]{neighbor, neighbor.Distance(m, p), func() *StateValue[C, T] { return sv }, sv.pathLen + 1})
		}
	}
	var nill C
	return nil, nill
}

type bfsHeap[C maths.Comparable[C], T any] struct {
	values []*StateValue[C, T]
}

func (bh *bfsHeap[C, T]) PushState(sv *StateValue[C, T]) {
	heap.Push(bh, sv)
}

func (bh *bfsHeap[C, T]) PopState() *StateValue[C, T] {
	return heap.Pop(bh).(*StateValue[C, T])
}

func (bh *bfsHeap[C, T]) popState() *StateValue[C, T] {
	return heap.Pop(bh).(*StateValue[C, T])
}

// Below functions needed for heap interface
func (bh *bfsHeap[C, T]) Len() int {
	return len(bh.values)
}

func (bh *bfsHeap[C, T]) Less(i, j int) bool {
	return bh.values[i].dist.LT(bh.values[j].dist)
}

func (bh *bfsHeap[C, T]) Push(x interface{}) {
	bh.values = append(bh.values, x.(*StateValue[C, T]))
}

func (bh *bfsHeap[C, T]) Pop() interface{} {
	r := bh.values[len(bh.values)-1]
	bh.values = bh.values[:len(bh.values)-1]
	return r
}

func (bh *bfsHeap[C, T]) Swap(i, j int) {
	bh.values[i], bh.values[j] = bh.values[j], bh.values[i]
}

type StateValue[C maths.Comparable[C], T any] struct {
	state T
	// This can be replaced by wrapping type for specific search type wrapper
	dist C
	// TODO: this can be replaced by improving container method (to include CurrentPath() function)
	prev func() *StateValue[C, T]
	//
	pathLen int
}

func (sv *StateValue[C, T]) PathString() string {
	return strings.Join(append(sv.path()), ", ")
}

func (sv *StateValue[C, T]) path() []string {
	if sv == nil {
		return []string{}
	}
	return append(sv.Prev().path(), sv.String())
}

func (sv *StateValue[C, T]) String() string {
	return fmt.Sprintf("(%v) %v", sv.dist, sv.state)
}

func (sv *StateValue[C, T]) State() T {
	return sv.state
}

func (sv *StateValue[C, T]) Dist() C {
	return sv.dist
}

func (sv *StateValue[C, T]) Prev() *StateValue[C, T] {
	if sv.prev == nil {
		return nil
	}
	return sv.prev()
}
