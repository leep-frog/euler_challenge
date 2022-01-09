package bfs

import (
	"fmt"
)

// TODO: this file still needs work

type CycleState[M, T any] interface {
	Code(*Context[M, T]) string
	AdjacentStates(*Context[M, T]) []T

	// DoneCycle takes the cycle obtained as input and returns whether that cycle is one of the types we are looking for.
	DoneCycle(*Context[M, T]) bool
}

type cycleState[M any, T CycleState[M, T]] struct {
	cs T
}

func (cs *cycleState[M, T]) Code(ctx *Context[*cycleContext[M], *cycleState[M, T]]) string {
	return cs.cs.Code(fromCycleCtx(ctx))
}

func (cs *cycleState[M, T]) String() string {
	return fmt.Sprintf("cycleState(%v)", cs.cs)
}

type cycleContext[M any] struct {
	m       M
	checked map[string]bool
}

func (cc *cycleContext[M]) String() string {
	return fmt.Sprint(cc.checked)
}

func convertCycleStateValue[M any, T CycleState[M, T]](sv *StateValue[*cycleState[M, T]]) *StateValue[T] {
	if sv == nil {
		return nil
	}
	return &StateValue[T]{
		state: sv.state.cs,
		dist:  sv.dist,
		prev: func() *StateValue[T] {
			return convertCycleStateValue[M, T](sv.Prev())
		},
	}
}

func fromCycleCtx[M any, T CycleState[M, T]](ctx *Context[*cycleContext[M], *cycleState[M, T]]) *Context[M, T] {
	return &Context[M, T]{
		ctx.GlobalContext.m,
		convertCycleStateValue[M, T](ctx.StateValue),
	}
}

func (cs *cycleState[M, T]) Done(ctx *Context[*cycleContext[M], *cycleState[M, T]]) bool {
	oCtx := fromCycleCtx[M, T](ctx)
	code := cs.cs.Code(oCtx)
	//fmt.Println(cs.cs.DoneCycle(oCtx))
	//return ctx.GlobalContext.checked[code]
	return ctx.GlobalContext.checked[code] && cs.cs.DoneCycle(oCtx)
}

func (cs *cycleState[M, T]) AdjacentStates(ctx *Context[*cycleContext[M], *cycleState[M, T]]) []*cycleState[M, T] {
	//
	oCtx := fromCycleCtx[M, T](ctx)
	code := cs.cs.Code(oCtx)
	if ctx.GlobalContext.checked[code] {
		return nil
	}

	var css []*cycleState[M, T]
	for _, as := range cs.cs.AdjacentStates(oCtx) {
		css = append(css, &cycleState[M, T]{as})
	}
	return css
}

func CyclePath[M any, T CycleState[M, T]](initState T, globalContext M) ([]T, int) {
	aph := &anyPathHelper[*cycleContext[M], *cycleState[M, T], *cycleState[M, T]]{
		popFunc: func(ctx *Context[*cycleContext[M], *cycleState[M, T]], cs *cycleState[M, T]) {
			//delete(ctx.GlobalContext.checked, cs.cs.Code(fromCycleCtx(ctx)))
		},
		pushFunc: func(ctx *Context[*cycleContext[M], *cycleState[M, T]], cs *cycleState[M, T]) {
			//ctx.GlobalContext.checked[cs.cs.Code(fromCycleCtx(ctx))] = true
		},
		ph: &pathHelper[*cycleContext[M], *cycleState[M, T], *cycleState[M, T]]{
			convFunc: identityConvFunc[*cycleContext[M], *cycleState[M, T]](),
			distFunc: simpleDistFunc[*cycleContext[M], *cycleState[M, T]](),
		},
	}
	css, dist := anyPath(
		&cycleState[M, T]{initState},
		0,
		&cycleContext[M]{
			m:       globalContext,
			checked: map[string]bool{},
		},
		aph,
	)
	var ts []T
	for _, cs := range css {
		ts = append(ts, cs.cs)
	}
	return ts, dist
}
