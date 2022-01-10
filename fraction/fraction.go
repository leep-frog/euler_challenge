package fraction

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/maths"
)

type Fraction[T any] struct {
	N T
	D T
	plus func(T, T) T
	times func(T, T) T
}

func New(n, d int) *Fraction[int] {
	return &Fraction[int]{
		n, 
		d,
		func(a, b int) int { return a + b },
		func(a, b int) int { return a * b },
	}
}

func NewBig(n, d *maths.Int) *Fraction[*maths.Int] {
	return &Fraction[*maths.Int]{
		n, 
		d,
		func(a, b *maths.Int) *maths.Int { return a.Plus(b) },
		func(a, b *maths.Int) *maths.Int { return a.Times(b) },
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
	return &Fraction[T]{f.N, f.D, f.plus, f.times}
}

// Return a fraction to allow for chaining.
/*func (f *Fraction[T]) Simplify(p *generator.Generator[T]) *Fraction[T] {
	nfs := generator.PrimeFactors(f.N, p)	
	dfs := generator.PrimeFactors(f.D, p)	

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
	f.N, f.D = newN, newD
	return f
}*/