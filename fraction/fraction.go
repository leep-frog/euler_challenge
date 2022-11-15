package fraction

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

// TODO: Fraction (maths.Mathable) and FractionI (from interface)

// Have this implement mathable
type Fraction[T any] struct {
	N     T
	D     T
	plus  func(T, T) T
	times func(T, T) T
	lt    func(T, T) bool
}

func New[T maths.Mathable](n, d T) *Fraction[T] {
	return &Fraction[T]{
		n,
		d,
		func(a, b T) T { return a + b },
		func(a, b T) T { return a * b },
		func(a, b T) bool { return a < b },
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

func (f *Fraction[T]) Reciprocal() *Fraction[T] {
	return &Fraction[T]{f.D, f.N, f.plus, f.times, f.lt}
}

func (f *Fraction[T]) Plus(that *Fraction[T]) *Fraction[T] {
	n := f.plus(f.times(f.N, that.D), f.times(that.N, f.D))
	d := f.times(f.D, that.D)
	return &Fraction[T]{n, d, f.plus, f.times, f.lt}
}

func (f *Fraction[T]) Code() string {
	return f.String()
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
	if d == 0 {
		if n == 0 {
			return New(0, 0)
		}
		return New(1, 0)
	}
	if n == 0 {
		// we know d isn't 0
		return New(0, 1)
	}

	sign := 1
	if d < 0 {
		d *= -1
		sign *= -1
	}
	if n < 0 {
		n *= -1
		sign *= -1
	}

	nfs := generator.MutablePrimeFactors(n, p)
	dfs := generator.MutablePrimeFactors(d, p)

	for k, v := range nfs {
		if dv, ok := dfs[k]; ok {
			m := maths.Min(v, dv)
			nfs[k] -= m
			dfs[k] -= m
		}
	}

	newN, newD := 1, 1
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
	return New(sign*newN, newD)
}
