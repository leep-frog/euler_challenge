package bfs

/*type CycleState[M, T any] interface {
	Code(*Context[M, T]) string
	AdjacentStates(*Context[M, T]) []T
}

type cycleState[M, T any] struct {
	cs CycleState[M, T]
}

type cycleContext[M any] struct {
	m M
	checked map[string]bool
}

func fromCycleCtx[M, T any](ctx *Context[*cycleContext[M], T]) *Context[M, T] {
	return &Context[M, T]{ctx.GlobalContext.m, ctx.StateValue}
}

func (cs *cycleState[M, T]) Done(ctx *Context[*cycleContext[M, T]]) bool {
	return ctx.GlobalContext.checked[cs.cs.Code(ctx.GlobalContext.m)]
}

func (cs *cycleState[M, T]) AdjacentStates(ctx *Context[*cycleContext[M, T], T]) []T {
	return cs.cs.AdjacentStates(&Context[M, T]{ctx.GlobalContext.m, ctx.StateValue})
}

func ShortestCyclePath[M any, T CycleState[M, T]](initState T, globalContext M) ([]T, int) {
	return anyPath(
		&cycleState[M, T]{initState},
		0,
		&Context[*cycleState[M], T]{
			m: globalContext,
			checked: map[string]bool{},
		},
		&anyPathHelper{
			popFunc: func(ctx *Context[*cycleState[M], T], t T) {
				delete(ctx.GlobalContext.checked, t.Code())
			},
			pushFunc: func(ctx *Context[*cycleState[M], T], t T) {
				ctx.GlobalContext.checked[t.Code()] = true
			},
			ph: &pathHelper[M, T, AS]{
				convFunc: identityConvFunc(),
				distFunc: simpleDistFunc(),
			},
		},
	)
}*/
