package maths

import "fmt"

type Bestable interface {
	~int | ~int64 | ~uint64 | ~float64 | ~string
}

func Smallest[I any, T Bestable]() *Bester[I, T] {
	return &Bester[I, T]{
		better: func(i, j T) bool {
			return i < j
		},
	}
}

func Closest[I any, T ~int | ~int64 | ~float64](target T) *Bester[I, T] {
	return &Bester[I, T]{
		better: func(i, j T) bool {
			return Abs(target-i) < Abs(target-j)
		},
	}
}

func Largest[I any, T Bestable]() *Bester[I, T] {
	return &Bester[I, T]{
		better: func(i, j T) bool {
			return i > j
		},
	}
}

/*func Closest[I any, T Bestable](center T) *Bester[I, T] {
	return &Bester[I, T]{
		better: func(i, j T) bool {
			return Abs(center-i) < Abs(center-j)
		},
	}
}*/

func SmallestT[I any, T Comparable[T]]() *Bester[I, T] {
	return &Bester[I, T]{
		better: LT[T],
	}
}

func LargestT[I any, T Comparable[T]]() *Bester[I, T] {
	return &Bester[I, T]{
		better: GT[T],
	}
}

// T for comparable type, I for index type.
type Bester[I, T any] struct {
	better func(T, T) bool
	best   T
	bestI  I

	set bool
}

func (b *Bester[I, T]) String() string {
	if !b.set {
		return "{}"
	}
	return fmt.Sprintf("{Best: %v, Index: %v}", b.best, b.bestI)
}

func (b *Bester[I, T]) Best() T {
	return b.best
}

func (b *Bester[I, T]) Set() bool {
	return b.set
}

func (b *Bester[I, T]) BestIndex() I {
	return b.bestI
}

func (b *Bester[I, T]) Check(v T) bool {
	if !b.set || b.better(v, b.best) {
		b.best = v
		b.set = true
		return true
	}
	return false
}

func (b *Bester[I, T]) IndexCheck(idx I, v T) bool {
	if !b.set || b.better(v, b.best) {
		b.best = v
		b.bestI = idx
		b.set = true
		return true
	}
	return false
}

type IncrementalBester struct {
	b *Bester[int, int]
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
		b: Largest[int, int](),
		m: map[int]int{},
	}
}
