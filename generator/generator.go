package generator

import (
	"constraints"
	"fmt"

	"github.com/leep-frog/euler_challenge/maths"
)

type Generator[T any] struct {
	values []T
	set    map[string]bool

	f   func(*Generator[T]) T
	idx int
}

func (g *Generator[T]) last() T {
	return g.values[len(g.values)-1]
}

func (g *Generator[T]) len() int {
	return len(g.values)
}

func (g *Generator[T]) Nth(i int) T {
	for g.len() <= i {
		g.getNext()
	}
	return g.values[i]
}

func (g *Generator[T]) Next() T {
	g.idx++
	return g.Nth(g.idx - 1)
}

func (g *Generator[T]) getNext() T {
	i := g.f(g)
	g.values = append(g.values, i)
	if g.set == nil {
		g.set = map[string]bool{}
	}
	g.set[fmt.Sprintf("%v", i)] = true
	return i
}

func SystemContains[T constraints.Ordered](g *Generator[T], t T) bool {
	for ; g.len() == 0 || g.last() <= t; g.getNext() {
	}
	return g.set[fmt.Sprintf("%v", t)]
}

type Comparable[T any] interface {
	LTE(T) bool
}

func Contains[T Comparable[T]](g *Generator[T], t T) bool {
	for ; g.len() == 0 || g.last().LTE(t); g.getNext() {
	}
	return g.set[fmt.Sprintf("%v", t)]
}

// TODO: cache stuff (every 1000?)

func NewGenericator[T any](start T, f func(*Generator[T]) T) *Generator[T] {
	return &Generator[T]{
		f: func(g *Generator[T]) T {
			if len(g.values) == 0 {
				return start
			}
			return f(g)
		},
	}
}

func PrimeFactors(n int, p *Generator[int]) map[int]int {
	r := map[int]int{}
	for i := 0; ; i++ {
		pi := int(p.Nth(i))
		for n%pi == 0 {
			r[pi]++
			n = n / pi
			if n == 1 {
				return r
			}
		}
	}
}

func Primes() *Generator[int] {
	return NewGenericator(2, func(g *Generator[int]) int {
		for i := g.last() + 1; ; i++ {
			newPrime := true
			for _, p := range g.values {
				if rem := i % p; rem == 0 {
					newPrime = false
					break
				}
			}
			if newPrime {
				return i
			}
		}
	})
}

func Fibonaccis() *Generator[int] {
	a, b := 1, 1
	return NewGenericator(1, func(g *Generator[int]) int {
		r := b
		b = a + b
		a = r
		return int(a)
	})
}

func BigFibonaccis() *Generator[*maths.Int] {
	a, b := maths.One(), maths.One()
	return NewGenericator(maths.One(), func(g *Generator[*maths.Int]) *maths.Int {
		r := b
		b = a.Plus(b)
		a = r
		return a
	})
}

func Triangulars() *Generator[int] {
	i := 1
	return NewGenericator(1, func(g *Generator[int]) int {
		i++
		return g.last() + int(i)
	})
}
