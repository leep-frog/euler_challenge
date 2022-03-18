package bfs

import (
	"fmt"
	"reflect"
	"strings"
	"container/heap"
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
	path Path[T]
	converter converter[T, T2]
}

func (p *pathWrapper[T, T2]) Len() int {
	return p.path.Len()
}

func (p *pathWrapper[T, T2]) Fetch() []T2 {
	return p.converter.convertSlice(p.path.Fetch())
}

type path[T any] struct {
	tail *StateValue[T]
}

func (p *path[T]) Len() int {
	if p.tail == nil {
		return 0
	}
	return p.tail.pathLen
}

func (p *path[T]) Fetch() []T {
	var r []T
	for cur := p.tail; cur != nil; cur = cur.Prev() {
		r = append(r, cur.state)
	}
	for i := 0; i < len(r)/2; i++ {
		r[i], r[len(r)-1-i] = r[len(r)-1-i], r[i]
	}
	return r
}

type Searchable[T any] interface {
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
	// TODO: make this return a mathable
	Distance() int
}

type SearchableWithContext[M, T any] interface {
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
	Distance(M) int
}

func toAPWConverter[M any, T SearchableWithContext[M, T]]() converter[T, *addPathWrapper[M, T]] {
	return func(t T) *addPathWrapper[M, T]{
		return &addPathWrapper[M, T]{t}
	}
}

func fromAPWConverter[M any, T SearchableWithContext[M, T]]() converter[*addPathWrapper[M, T], T] {
	return func(apw *addPathWrapper[M, T]) T{
		return apw.state
	}
}

type addPathWrapper[M any, T SearchableWithContext[M, T]] struct {
	state T
}

func (apw *addPathWrapper[M, T]) String() string {
	return fmt.Sprintf("%v", apw.state)
}

func (apw *addPathWrapper[M, T]) Code(m M, _ Path[*addPathWrapper[M, T]]) string {
	return apw.state.Code(m)
}

func (apw *addPathWrapper[M, T]) Done(m M, _ Path[*addPathWrapper[M, T]]) bool {
	return apw.state.Done(m)
}

func (apw *addPathWrapper[M, T]) Distance(m M, _ Path[*addPathWrapper[M, T]]) int {
	return apw.state.Distance(m)
}

func (apw *addPathWrapper[M, T]) AdjacentStates(m M, _ Path[*addPathWrapper[M, T]]) []*addPathWrapper[M, T] {
	var r []*addPathWrapper[M, T]
	for _, n := range apw.state.AdjacentStates(m) {
		r = append(r, &addPathWrapper[M, T]{n})
	}
	return r
}

func toACPConverter[T Searchable[T]]() converter[T, *addContextAndPathWrapper[T]] {
	return func(t T) *addContextAndPathWrapper[T]{
		return &addContextAndPathWrapper[T]{t}
	}
}

func fromACPConverter[T Searchable[T]]() converter[*addContextAndPathWrapper[T], T] {
	return func(apw *addContextAndPathWrapper[T]) T{
		return apw.state
	}
}

type addContextAndPathWrapper[T Searchable[T]] struct {
	state T
}

func (acp *addContextAndPathWrapper[T]) String() string {
	return fmt.Sprintf("%v", acp.state)
}

func (acp *addContextAndPathWrapper[T]) Code(_ int, _ Path[*addContextAndPathWrapper[T]]) string {
	return acp.state.Code()
}

func (acp *addContextAndPathWrapper[T]) Done(_ int, _ Path[*addContextAndPathWrapper[T]]) bool {
	return acp.state.Done()
}

func (acp *addContextAndPathWrapper[T]) Distance(_ int, _ Path[*addContextAndPathWrapper[T]]) int {
	return acp.state.Distance()
}

func (acp *addContextAndPathWrapper[T]) AdjacentStates(_ int, _ Path[*addContextAndPathWrapper[T]]) []*addContextAndPathWrapper[T] {
	var r []*addContextAndPathWrapper[T]
	for _, n := range acp.state.AdjacentStates() {
		r = append(r, &addContextAndPathWrapper[T]{n})
	}
	return r
}

type SearchableWithContextAndPath[M, T any] interface {
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
	Distance(M, Path[T]) int
}

type Option func(o *option)

type option struct {
	checkDuplicates bool
}

func CheckDuplicates() Option {
	return func(o *option) {
		o.checkDuplicates = true
	}
}

func newSearch[M any, T SearchableWithContextAndPath[M, T]](initStates []T, m M, opts ...Option) (Path[T], int) {
	o := &option{}
	for _, opt := range opts {
		opt(o)
	}

	nodes := &bfsHeap[T]{}
	for _, is := range initStates {
		nodes.PushState(&StateValue[T]{is, is.Distance(m, &path[T]{nil}), nil, 1})
	}

	checked := map[string]bool{}

	for nodes.Len() > 0 {
		sv := nodes.PopState()
		p := &path[T]{sv}
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
			nodes.PushState(&StateValue[T]{neighbor, neighbor.Distance(m, p), func() *StateValue[T] { return sv }, sv.pathLen+1})
		}
	}
	return nil, -1
}

type bfsHeap[T any] struct {
	values []*StateValue[T]
}

func (bh *bfsHeap[T]) PushState(sv *StateValue[T]) {
	heap.Push(bh, sv)
}

func (bh *bfsHeap[T]) PopState() *StateValue[T] {
	return heap.Pop(bh).(*StateValue[T])
}

func (bh *bfsHeap[T]) popState() *StateValue[T] {
	return heap.Pop(bh).(*StateValue[T])
}

// Below functions needed for heap interface
func (bh *bfsHeap[T]) Len() int {
	return len(bh.values)
}

func (bh *bfsHeap[T]) Less(i, j int) bool {
	return bh.values[i].dist < bh.values[j].dist
}

func (bh *bfsHeap[T]) Push(x interface{}) {
	bh.values = append(bh.values, x.(*StateValue[T]))
}

func (bh *bfsHeap[T]) Pop() interface{} {
	r := bh.values[len(bh.values)-1]
	bh.values = bh.values[:len(bh.values)-1]
	return r
}

func (bh *bfsHeap[T]) Swap(i, j int) {
	bh.values[i], bh.values[j] = bh.values[j], bh.values[i]
}

type StateValue[T any] struct {
	state T
	// This can be replaced by wrapping type for specific search type wrapper
	dist  int
	// TODO: this can be replaced by improving container method (to include CurrentPath() function)
	prev  func() *StateValue[T]
	// 
	pathLen int
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
