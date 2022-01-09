package bfs

import (
	"fmt"
)

type Set[M, T any] interface {
	Code(*Context[M, T]) string
	AdjacentStates(*Context[M, T]) []T
	BiggerThan(T) bool
	HasEdge(*Context[M, T], T) bool
	Offset(*Context[M, T]) int
}

type setState[M any, T Set[M, T]] struct {
	s              T
	remainingDepth int
}

func (cs *setState[M, T]) String() string {
	return fmt.Sprintf("setState(%v)", cs.s)
}

func (cs *setState[M, T]) Code(ctx *Context[M, *setState[M, T]]) string {
	return cs.s.Code(fromSetContext(ctx))
}

func (cs *setState[M, T]) Offset(ctx *Context[M, *setState[M, T]]) int {
	return cs.s.Offset(fromSetContext(ctx))
}

func (cs *setState[M, T]) Done(ctx *Context[M, *setState[M, T]]) bool {
	return ctx != nil && ctx.StateValue != nil && ctx.StateValue.State().remainingDepth <= 0
}

func (cs *setState[M, T]) AdjacentStates(ctx *Context[M, *setState[M, T]]) []*setState[M, T] {
	// If already bigger than size, then no need to check more
	if ctx != nil && ctx.StateValue != nil && ctx.StateValue.State().remainingDepth <= 0 {
		return nil
	}
	var r []*setState[M, T]
	for _, as := range cs.s.AdjacentStates(fromSetContext(ctx)) {
		if !as.BiggerThan(cs.s) {
			continue
		}
		edgeToAll := true
		for cur := ctx.StateValue; cur != nil; cur = cur.Prev() {
			if !as.HasEdge(fromSetContext(ctx), cur.State().s) {
				edgeToAll = false
				break
			}
		}
		if edgeToAll {
			r = append(r, &setState[M, T]{as, cs.remainingDepth - 1})
		}
	}
	return r
}

func convertSetState[M any, T Set[M, T]](sv *StateValue[*setState[M, T]]) *StateValue[T] {
	if sv == nil {
		return nil
	}
	return &StateValue[T]{
		state: sv.state.s,
		dist:  sv.dist,
		prev: func() *StateValue[T] {
			return convertSetState[M, T](sv.Prev())
		},
	}
}

func fromSetContext[M any, T Set[M, T]](ctx *Context[M, *setState[M, T]]) *Context[M, T] {
	return &Context[M, T]{
		ctx.GlobalContext,
		convertSetState[M, T](ctx.StateValue),
	}
}

func CompleteSets[M any, T Set[M, T]](sets []T, globalContext M, size int) []T {
	var setStates []*setState[M, T]
	for _, s := range sets {
		setStates = append(setStates, &setState[M, T]{s, size})
	}
	path, _ := ShortestOffsetPathNonUnique(setStates, globalContext)
			var ts []T
			for _, p := range path {
				ts = append(ts, p.s)
			}
			return ts
}
