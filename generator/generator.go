package generator

import (
	"bufio"
	"strconv"

	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/parse"
)

const (
	primesName = "primes"
	fibName = "fibonaccis"
	triName = "triangulars"
)

type Generatable[T any] interface {
	LTE(T, T) bool
	String(T) string
	FromString(string) T
}

func newBigGeneratable() Generatable[*maths.Int] {
	return &bigGeneratable{}
}

type bigGeneratable struct {}

func (bg *bigGeneratable) LTE(this, that *maths.Int) bool {
	return this.LTE(that)
}

func (bg *bigGeneratable) String(i *maths.Int) string {
	return i.String()
}

func (bg *bigGeneratable) FromString(s string)*maths.Int {
	return maths.MustIntFromString(s)
}

func newIntGeneratable() Generatable[int] {
	return &intGeneratable{}
}

type intGeneratable struct {}

func (ig *intGeneratable) LTE(this, that int) bool {
	return this <= that
}

func (ig *intGeneratable) String(i int) string {
	return strconv.Itoa(i)
}

func (ig *intGeneratable) FromString(s string) int {
	return parse.Atoi(s)
}

type Generator[T any] struct {
	values []T
	set    map[string]bool

	g Generatable[T]

	f   func(*Generator[T]) T
	idx int

	scanner *bufio.Scanner
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
	if g.scanner != nil && g.scanner.Scan() {
		g.values = append(g.values, g.g.FromString(g.scanner.Text()))
	}
	i := g.f(g)
	g.values = append(g.values, i)
	if g.set == nil {
		g.set = map[string]bool{}
	}
	g.set[g.g.String(i)] = true
	return i
}

func (g *Generator[T]) Contains(t T) bool {
	for ; g.len() == 0 || g.g.LTE(g.last(), t); g.getNext() {
	}
	return g.set[g.g.String(t)]
}

// TODO: cache stuff (every 1000?)

func NewGenericator[T any](start T, g Generatable[T], f func(*Generator[T]) T) *Generator[T] {
	return &Generator[T]{
		g: g,
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
	return NewGenericator(2, newIntGeneratable(), func(g *Generator[int]) int {
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

func BigPrimes() *Generator[*maths.Int] {
	return NewGenericator(maths.NewInt(2), newBigGeneratable(), func(g *Generator[*maths.Int]) *maths.Int {
		for i := g.last().Plus(maths.One()); ; i.PP() {
			newPrime := true
			for _, p := range g.values {
				if _, rem := i.Div(p); rem.IsZero() {
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
	return NewGenericator(1, newIntGeneratable(), func(g *Generator[int]) int {
		r := b
		b = a + b
		a = r
		return int(a)
	})
}

func BigFibonaccis() *Generator[*maths.Int] {
	a, b := maths.One(), maths.One()
	return NewGenericator(maths.One(), newBigGeneratable(), func(g *Generator[*maths.Int]) *maths.Int {
		r := b
		b = a.Plus(b)
		a = r
		return a
	})
}

func Triangulars() *Generator[int] {
	i := 1
	return NewGenericator(1, newIntGeneratable(), func(g *Generator[int]) int {
		i++
		return g.last() + int(i)
	})
}
