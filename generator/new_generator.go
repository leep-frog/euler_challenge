package generator

import (
	"strconv"

	"github.com/leep-frog/euler_challenge/maths"
)

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

func Primes() *Generator[int] {
	return newIntGen(&primer{})
}

type primer struct{}

func (p *primer) Next(g *Generator[int]) int {
	if len(g.values) == 0 {
		return 2
	}
	for i := g.Last() + 1; ; i++ {
		newPrime := true
		for _, p := range g.Values() {
			if p*p > i {
				break
			}
			if i%p == 0 {
				newPrime = false
				break
			}
		}
		if newPrime {
			return i
		}
	}
}

func BigPrimes() *Generator[*maths.Int] {
	return newBigGen(&bigPrimer{})
}

type bigPrimer struct{}

func (bp *bigPrimer) Next(g *Generator[*maths.Int]) *maths.Int {
	if len(g.values) == 0 {
		return maths.NewInt(2)
	}
	for i := g.Last().Plus(maths.One()); ; i.PP() {
		newPrime := true
		for _, p := range g.Values() {
			if p.Times(p).GT(i) {
				break
			}
			if i.Mod(p).IsZero() {
				newPrime = false
				break
			}
		}
		if newPrime {
			return i
		}
	}
}

func Fibonaccis() *Generator[int] {
	return newIntGen(&fibs{})
}

type fibs struct{}

func (f *fibs) Next(g *Generator[int]) int {
	if len(g.values) < 2 {
		return 1
	}
	return g.values[len(g.values)-1] + g.values[len(g.values)-2]
}

func BigFibonaccis() *Generator[*maths.Int] {
	return newBigGen(&bigFibs{})
}

type bigFibs struct{}

func (bf *bigFibs) Next(g *Generator[*maths.Int]) *maths.Int {
	if len(g.values) < 2 {
		return maths.One()
	}
	return g.values[len(g.values)-1].Plus(g.values[len(g.values)-2])
}

/*type RightTriangle struct {
	A, B, C int
}

func RightTriangleGenerator() *Generator[*RightTriangle] {
	return &Generator{
		"rightTriangle",
		nil,
		map[string]int{},
		&rightTriangleGeneratable{2, 1},
	}
}*/
