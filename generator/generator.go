package generator

import (
	"fmt"
	"strconv"

	"github.com/leep-frog/euler_challenge/maths"
)

func copy(m map[int]int) map[int]int {
	c := map[int]int{}
	for k, v := range m {
		c[k] = v
	}
	return c
}

type Geniterator[T any] struct {
	g   *Generator[T]
	Idx int
}

func Iterator[T any](g *Generator[T]) *Geniterator[T] {
	return &Geniterator[T]{g, 0}
}

// TODO: do these instead of Nth
func (gi *Geniterator[T]) Start(startIdx int) T {
	gi.Idx = startIdx + 1
	return gi.g.Nth(gi.Idx - 1)
}

func (gi *Geniterator[T]) Last() T {
	return gi.g.Last()
}

func (gi *Geniterator[T]) Next() T {
	r := gi.g.Nth(gi.Idx)
	gi.Idx++
	return r
}

type GeneratorInterface[T any] interface {
	Next(*Generator[T]) T
}

type Generator[T any] struct {
	values   []T
	set      map[string]bool
	gi       GeneratorInterface[T]
	less     func(T, T) bool
	toString func(T) string
}

func (g *Generator[T]) Last() T {
	return g.values[len(g.values)-1]
}

func (g *Generator[T]) Nth(i int) T {
	for len(g.values) <= i {
		nv := g.gi.Next(g)
		g.values = append(g.values, nv)
		g.set[g.toString(nv)] = true
	}
	return g.values[i]
}

func (g *Generator[T]) Iterator() *Geniterator[T] {
	return Iterator(g)
}

func (g *Generator[T]) Start(idx int) (*Geniterator[T], T) {
	iter := Iterator(g)
	return iter, iter.Start(idx)
}

func (g *Generator[T]) Contains(t T) bool {
	if len(g.values) == 0 {
		g.Nth(0)
	}
	for ; g.less(g.Last(), t); g.Nth(len(g.values)) {
	}
	_, ok := g.set[g.toString(t)]
	return ok
}

func (g *Generator[T]) Values() []T {
	return g.values
}

type CustomGeneratable[T any] interface {
	fmt.Stringer
	maths.Comparable[T]
}

func newCustomGen[T CustomGeneratable[T]](gi GeneratorInterface[T]) *Generator[T] {
	return &Generator[T]{
		gi:       gi,
		less:     func(i, j T) bool { return i.LT(j) },
		toString: func(i T) string { return i.String() },
		set:      map[string]bool{},
	}
}

func newIntGen(gi GeneratorInterface[int]) *Generator[int] {
	return &Generator[int]{
		gi:       gi,
		less:     func(i, j int) bool { return i < j },
		toString: strconv.Itoa,
		set:      map[string]bool{},
	}
}

func newBigGen(gi GeneratorInterface[*maths.Int]) *Generator[*maths.Int] {
	return &Generator[*maths.Int]{
		gi:       gi,
		less:     func(i, j *maths.Int) bool { return i.LT(j) },
		toString: func(i *maths.Int) string { return i.String() },
		set:      map[string]bool{},
	}
}
