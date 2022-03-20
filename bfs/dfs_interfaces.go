package bfs

// Options:
// Inputs (none, M, Path, M+Path)
// Processors (none, present)
// Interface names:
// 

// noType is used to fill in generic arguments that aren't used
type noType int
const nilNoType noType = 0

type SimpleDepthSearcher[T any] interface {
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
}

type sdsWrapper[T SimpleDepthSearcher[T]] struct {
	state T
}

type sdsBiconverter[T SimpleDepthSearcher[T]] struct {}
func (*sdsBiconverter[T]) To(t T) *sdsWrapper[T] { return &sdsWrapper[T]{t}}
func (*sdsBiconverter[T]) From(sds *sdsWrapper[T]) T { return sds.state }

// TODO: remove these converter funcs
func sdsConverter[T SimpleDepthSearcher[T]]() converter[T, *sdsWrapper[T]] {
	return func(t T) *sdsWrapper[T] {
		return &sdsWrapper[T]{t}
	} 
}

func (sds *sdsWrapper[T]) Code(noType, DFSPath[T]) string { return sds.state.Code() }
func (sds *sdsWrapper[T]) Done(noType, DFSPath[T]) bool { return sds.state.Done() }
func (sds *sdsWrapper[T]) OnPush(noType, DFSPath[T]) {}
func (sds *sdsWrapper[T]) OnPop(noType, DFSPath[T]) {}
func (sds *sdsWrapper[T]) AdjacentStates(noType, DFSPath[T]) []*sdsWrapper[T] {
	return sdsConverter[T]().convertSlice(sds.state.AdjacentStates())
}

func SimpleDFS[T SimpleDepthSearcher[T]](initStates []T, opts ...DFSOption) []T {
	b := &sdsBiconverter[T]{}
	return dfsFinal[noType, T, *sdsWrapper[T]](
		toSlice[T, *sdsWrapper[T]](b, initStates),
		nilNoType,
		b.From, 
		opts...)
}

type PoppableSimpleDepthSearcher[T any] interface {
	SimpleDepthSearcher[T]
	OnPush()	
	OnPop()
}

type poppableSDSWrapper[T PoppableSimpleDepthSearcher[T]] struct {
	state T
}

// TODO: replace SimpleDS with SDS
type poppableSimpleDSBiconverter[T PoppableSimpleDepthSearcher[T]] struct {}
func (*poppableSimpleDSBiconverter[T]) To(t T) *poppableSDSWrapper[T] { return &poppableSDSWrapper[T]{t}}
func (*poppableSimpleDSBiconverter[T]) From(sds *poppableSDSWrapper[T]) T { return sds.state }

func poppableSDSConverter[T PoppableSimpleDepthSearcher[T]]() converter[T, *poppableSDSWrapper[T]] {
	return func(t T) *poppableSDSWrapper[T] {
		return &poppableSDSWrapper[T]{t}
	} 
}

func (sds *poppableSDSWrapper[T]) Code(noType, DFSPath[T]) string { return sds.state.Code() }
func (sds *poppableSDSWrapper[T]) Done(noType, DFSPath[T]) bool { return sds.state.Done() }
func (sds *poppableSDSWrapper[T]) OnPush(noType, DFSPath[T]) { sds.state.OnPush() }
func (sds *poppableSDSWrapper[T]) OnPop(noType, DFSPath[T]) { sds.state.OnPop() }
func (sds *poppableSDSWrapper[T]) AdjacentStates(noType, DFSPath[T]) []*poppableSDSWrapper[T] {
	return poppableSDSConverter[T]().convertSlice(sds.state.AdjacentStates())
}

func PoppableSimpleDFS[T PoppableSimpleDepthSearcher[T]](initStates []T, opts ...DFSOption) []T {
	b := &poppableSimpleDSBiconverter[T]{}
	return dfsFinal[noType, T, *poppableSDSWrapper[T]](
		toSlice[T, *poppableSDSWrapper[T]](b, initStates),
		nilNoType,
		b.From, 
		opts...)
}

type DepthSearcher[M, T any] interface {
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
}

type dsWrapper[M any, T DepthSearcher[M, T]] struct {
	state T
}

type dsBiconverter[M any, T DepthSearcher[M, T]] struct {}
func (*dsBiconverter[M, T]) To(t T) *dsWrapper[M, T] { return &dsWrapper[M, T]{t}}
func (*dsBiconverter[M, T]) From(pdw *dsWrapper[M, T]) T { return pdw.state }

func dsConverter[M any, T DepthSearcher[M, T]]() converter[T, *dsWrapper[M, T]] {
	return func(t T) *dsWrapper[M, T] {
		return &dsWrapper[M, T]{t}
	} 
}

func (ds *dsWrapper[M, T]) Code(m M, _ DFSPath[T]) string { return ds.state.Code(m) }
func (ds *dsWrapper[M, T]) Done(m M, _ DFSPath[T]) bool { return ds.state.Done(m) }
func (ds *dsWrapper[M, T]) OnPush(m M, _ DFSPath[T]) {}
func (ds *dsWrapper[M, T]) OnPop(m M, _ DFSPath[T]) {}
func (ds *dsWrapper[M, T]) AdjacentStates(m M, _ DFSPath[T]) []*dsWrapper[M, T] {
	return dsConverter[M, T]().convertSlice(ds.state.AdjacentStates(m))
}

func DFS[M any, T DepthSearcher[M, T]](initStates []T, m M, opts ...DFSOption) []T {
	b := &dsBiconverter[M, T]{}
	return dfsFinal[M, T, *dsWrapper[M, T]](
		toSlice[T, *dsWrapper[M, T]](b, initStates),
		m,
		b.From, 
		opts...)
}

type PoppableDepthSearcher[M, T any] interface {
	DepthSearcher[M, T]
	OnPush(M)
	OnPop(M)
}

type poppableDSWrapper[M any, T PoppableDepthSearcher[M, T]] struct {
	state T
}

type poppableDSBiconverter[M any, T PoppableDepthSearcher[M, T]] struct {}
func (*poppableDSBiconverter[M, T]) To(t T) *poppableDSWrapper[M, T] { return &poppableDSWrapper[M, T]{t}}
func (*poppableDSBiconverter[M, T]) From(pdw *poppableDSWrapper[M, T]) T { return pdw.state }

func poppableDSConverter[M any, T PoppableDepthSearcher[M, T]]() converter[T, *poppableDSWrapper[M, T]] {
	return func(t T) *poppableDSWrapper[M, T] {
		return &poppableDSWrapper[M, T]{t}
	} 
}

func (pds *poppableDSWrapper[M, T]) Code(m M, _ DFSPath[T]) string { return pds.state.Code(m) }
func (pds *poppableDSWrapper[M, T]) Done(m M, _ DFSPath[T]) bool { return pds.state.Done(m) }
func (pds *poppableDSWrapper[M, T]) OnPush(m M, _ DFSPath[T]) { pds.state.OnPush(m) }
func (pds *poppableDSWrapper[M, T]) OnPop(m M, _ DFSPath[T]) { pds.state.OnPop(m) }
func (pds *poppableDSWrapper[M, T]) AdjacentStates(m M, _ DFSPath[T]) []*poppableDSWrapper[M, T] {
	return poppableDSConverter[M, T]().convertSlice(pds.state.AdjacentStates(m))
}

func PoppableDFS[M any, T PoppableDepthSearcher[M, T]](initStates []T, m M, opts ...DFSOption) []T {
	b := &poppableDSBiconverter[M, T]{}
	return dfsFinal[M, T, *poppableDSWrapper[M, T]](
		toSlice[T, *poppableDSWrapper[M, T]](b, initStates),
		m,
		b.From, 
		opts...)
}

type ContextualDepthSearcher[M, T any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code(M, DFSPath[T]) string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(M, DFSPath[T]) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(M, DFSPath[T]) []T
}

type contextualDSWrapper[M any, T ContextualDepthSearcher[M, T]] struct {
	state T
}

type contextualDSBiconverter[M any, T ContextualDepthSearcher[M, T]] struct {}
func (*contextualDSBiconverter[M, T]) To(t T) *contextualDSWrapper[M, T] { return &contextualDSWrapper[M, T]{t}}
func (*contextualDSBiconverter[M, T]) From(cdw *contextualDSWrapper[M, T]) T { return cdw.state }

func (cds *contextualDSWrapper[M, T]) Code(m M, p DFSPath[T]) string { return cds.state.Code(m, p) }
func (cds *contextualDSWrapper[M, T]) Done(m M, p DFSPath[T]) bool { return cds.state.Done(m, p) }
func (cds *contextualDSWrapper[M, T]) OnPush(m M, p DFSPath[T]) {}
func (cds *contextualDSWrapper[M, T]) OnPop(m M, p DFSPath[T]) {}
func (cds *contextualDSWrapper[M, T]) AdjacentStates(m M, p DFSPath[T]) []*contextualDSWrapper[M, T] {
	return toSlice[T, *contextualDSWrapper[M, T]](&contextualDSBiconverter[M, T]{}, cds.state.AdjacentStates(m, p))
}

func ContextualDFS[M any, T ContextualDepthSearcher[M, T]](initStates []T, m M, opts ...DFSOption) []T {
	b := &contextualDSBiconverter[M, T]{}
	return dfsFinal[M, T, *contextualDSWrapper[M, T]](
		toSlice[T, *contextualDSWrapper[M, T]](b, initStates),
		m,
		b.From, 
		opts...)
}

type PoppableContextualDepthSearcher[M, T any] interface {
	ContextualDepthSearcher[M, T]
	OnPush(M, DFSPath[T])
	OnPop(M, DFSPath[T])
}

func PoppableContextualDFS[M any, T PoppableContextualDepthSearcher[M, T]](initStates []T, m M, opts ...DFSOption) []T {
	return dfsFinal[M, T, T](initStates, m, identityConverter[T](), opts...)
}

// T2 is a different type for more efficient wrappers. Specifically so we don't need
// to create a DFSPath wrapper which would be inefficient since we'd have to
// convert the entire array every time (as opposed to just passing around a
// reference to a single array.
type completeDepthSearcher[M, T, T2 any] interface {
	// A unique code for the current state. This may be called multiple times
	// so this should be cached in the implementing code if computation is expensive.
	Code(M, DFSPath[T2]) string
	// Returns if the given state is in a final position. The input is a contextual variable
	// that is passed along from ShortestPath.
	Done(M, DFSPath[T2]) bool
	// Returns all of the adjacent states. The input is a contextual variable
	// that is passed along from ShortestPath.
	// T should always be State[M], but we cannot do that here without having a recursive type
	AdjacentStates(M, DFSPath[T2]) []T

	// OnPush runs when the node is first visited and added to the stack.
	OnPush(M, DFSPath[T2])
	// OnPop runs when the node is being removed from the stack.
	OnPop(M, DFSPath[T2])
}