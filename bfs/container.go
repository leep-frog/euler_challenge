package bfs

import (
	"container/heap"
)

type stateContainer[T any] interface {
	PushState(*StateValue[T])
  PopState() *StateValue[T]
	Len() int
	CurrentPath() []*StateValue[T]
}

// newBFSSearcher returns a searcher for depth first search.
func newDFSSearcher[T any]() stateContainer[T] {
	return &dfsSearcher[T]{}
}

type dfsSearcher[T any] struct {
	//stack []*StateValue[T]
	stack []*dfsItem[T]
	path []*StateValue[T]
}

type dfsItem[T any] struct {
	popPath bool
	item *StateValue[T]
}

func (ds *dfsSearcher[T]) PushState(sv *StateValue[T]) {
	ds.pushState(false, sv)
}

func (ds *dfsSearcher[T]) pushState(popPath bool, sv *StateValue[T]) {
	ds.stack = append(ds.stack, &dfsItem[T]{popPath, sv})
}

func (ds *dfsSearcher[T]) PopState() *StateValue[T] {
	item := ds.popState()
	for ; item.popPath; item = ds.popState() {
		ds.path = ds.path[:len(ds.path)-1]
	}
	ds.pushState(true, nil)
	ds.path = append(ds.path, item.item)
	return item.item
}

func (ds *dfsSearcher[T]) popState() *dfsItem[T] {
	r := ds.stack[len(ds.stack)-1]
	ds.stack = ds.stack[:len(ds.stack)-1]
	return r
}

func (ds *dfsSearcher[T]) Len() int {
	return len(ds.stack)
}

func (ds *dfsSearcher[T]) CurrentPath() []*StateValue[T] {
	return ds.path
}

// newBFSSearcher returns a searcher for breadth first search.
func newBFSSearcher[T any]() stateContainer[T] {
	return &bfsSearcher[T]{}
}

type bfsSearcher[T any] struct {
	values []*StateValue[T]
	path []T
}

type bfsItem[T any] struct {
	item T
	newPath []T
}

func (bs *bfsSearcher[T]) PushState(sv *StateValue[T]) {
	heap.Push(bs, sv)
}

func (bs *bfsSearcher[T]) PopState() *StateValue[T] {
	/*item := bs.popState()
	for ; */
	return heap.Pop(bs).(*StateValue[T])
}

func (bs *bfsSearcher[T]) CurrentPath() []*StateValue[T] {
	return nil
}

func (bs *bfsSearcher[T]) popState() *StateValue[T] {
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