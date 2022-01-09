package bfs

type CycleState[M, T any] interface {
	Code(*Context[M, T]) string
	AdjacentStates(*Context[M, T]) []T
}

type cycleState[M any, T CycleState[M, T]] struct {
	cs T
}

type cycleContext[M any] struct {
	m       M
	checked map[string]bool
}

func convertCycleStateValue[M any, T CycleState[M, T]](sv *StateValue[*cycleState[M, T]]) *StateValue[T] {
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
	code := cs.cs.Code(fromCycleCtx[M, T](ctx))
	return ctx.GlobalContext.checked[code]
}

func (cs *cycleState[M, T]) AdjacentStates(ctx *Context[*cycleContext[M], T]) []T {
	return cs.cs.AdjacentStates(&Context[M, T]{ctx.GlobalContext.m, ctx.StateValue})
}

func ShortestCyclePath[M any, T CycleState[M, T]](initState T, globalContext M) ([]T, int) {
	aph := &anyPathHelper[*cycleContext[M], *cycleState[M, T], T]{
		popFunc: func(ctx *Context[*cycleContext[M], *cycleState[M, T]], cs *cycleState[M, T]) {
			delete(ctx.GlobalContext.checked, cs.cs.Code(fromCycleCtx(ctx)))
		},
		/*pushFunc: func(ctx *Context[*cycleState[M], T], t T) {
			ctx.GlobalContext.checked[t.Code()] = true
		},
		ph: &pathHelper[M, T, AS]{
			convFunc: identityConvFunc(),
			distFunc: simpleDistFunc(),
		},*/
	}
	_ = aph
	return nil, 0
	/*return anyPath(
		&cycleState[M, T]{initState},
		0,
		&cycleContext[M]{
			m: globalContext,
			checked: map[string]bool{},
			},

	)*/
}
