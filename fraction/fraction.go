package fraction

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/maths"
	"github.com/leep-frog/euler_challenge/generator"
)

type Fraction[T any] struct {
	N T
	D T
	plus func(T, T) T
	times func(T, T) T
	lt func(T, T) bool
}

func New(n, d int) *Fraction[int] {
	return &Fraction[int]{
		n, 
		d,
		func(a, b int) int { return a + b },
		func(a, b int) int { return a * b },
		func(a, b int) bool { return a < b },
	}
}

func NewBig(n, d *maths.Int) *Fraction[*maths.Int] {
	return &Fraction[*maths.Int]{
		n, 
		d,
		func(a, b *maths.Int) *maths.Int { return a.Plus(b) },
		func(a, b *maths.Int) *maths.Int { return a.Times(b) },
		func(a, b *maths.Int) bool { return a.LT(b) },
	}
}

func (f *Fraction[T]) Invert() *Fraction[T] {
	tmp := f.N
	f.N = f.D
	f.D = tmp
	return f
}

func (f *Fraction[T]) Plus(t T) *Fraction[T] {
	f.N = f.plus(f.N, f.times(f.D, t))
	return f
}

func (f *Fraction[T]) String() string {
	return fmt.Sprintf("%v/%v", f.N, f.D)
}

func (f *Fraction[T]) Copy() *Fraction[T] {
	return &Fraction[T]{f.N, f.D, f.plus, f.times, f.lt}
}

func (f *Fraction[T]) LT(that *Fraction[T]) bool {
	return f.lt(f.times(f.N, that.D), f.times(f.D, that.N))
}

// Return a fraction to allow for chaining.
func Simplify(n, d int, p *generator.Generator[int]) *Fraction[int] {
	nfs := generator.MutablePrimeFactors(n, p)	
	dfs := generator.MutablePrimeFactors(d, p)	

	for k, v := range nfs {
		if dv, ok := dfs[k]; ok {
			m := maths.Min(v, dv)
			nfs[k] -= m
			dfs[k] -= m
		}
	}

	newN, newD  := 1, 1
	for k, v := range nfs {
		for i := 0; i < v; i++ {
			newN *= k
		}
	}
	for k, v := range dfs {
		for i := 0; i < v; i++ {
			newD *= k
		}
	}
	return New(newN, newD)
}
