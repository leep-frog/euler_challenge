package maths

func Smallest[T Mathable]() *Bester[T] {
	return &Bester[T]{
		better: func(i, j T) bool {
			return i < j
		},
	}
}

func Largest[T Mathable]() *Bester[T] {
	return &Bester[T]{
		better: func(i, j T) bool {
			return i > j
		},
	}
}

func SmallestT[T Operable[T]]() *Bester[T] {
	return &Bester[T]{
		better: LT[T],
	}
}

func LargestT[T Operable[T]]() *Bester[T] {
	return &Bester[T]{
		better: GT[T],
	}
}

type Bester[T any] struct {
	better func(T, T) bool
	best   T
	bestI  int

	set bool
}

func (b *Bester[T]) Best() T {
	return b.best
}

func (b *Bester[T]) BestIndex() int {
	return b.bestI
}

func (b *Bester[T]) Check(v T) {
	if !b.set || b.better(v, b.best) {
		b.best = v
		b.set = true
	}
}

func (b *Bester[T]) IndexCheck(idx int, v T) {
	if !b.set || b.better(v, b.best) {
		b.best = v
		b.bestI = idx
		b.set = true
	}
}

type IncrementalBester struct {
	b *Bester[int]
	m map[int]int
}

func (ib *IncrementalBester) Best() int {
	return ib.b.best
}

func (ib *IncrementalBester) BestIndex() int {
	return ib.b.bestI
}

func (ib *IncrementalBester) Increment(v int) {
	ib.m[v]++
	ib.b.IndexCheck(v, ib.m[v])
}

func LargestIncremental() *IncrementalBester {
	return &IncrementalBester{
		b: Largest[int](),
		m: map[int]int{},
	}
}