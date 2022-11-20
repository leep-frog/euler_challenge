package fraction

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

type Fraction[T maths.Mathable] struct {
	N T
	D T
}

type FractionI[T maths.Operable[T]] struct {
	N T
	D T
}

func New[T maths.Mathable](n, d T) *Fraction[T] {
	absN, absD := maths.Abs(n), maths.Abs(d)
	if n*d < 0 {
		absN = -absN
	}
	return &Fraction[T]{absN, absD}
}

func NewI[T maths.Operable[T]](n, d T) *FractionI[T] {
	return &FractionI[T]{n, d}
}

func (f *Fraction[T]) Times(that *Fraction[T]) *Fraction[T] {
	return New(f.N*that.N, f.D*that.D)
}

func (f *Fraction[T]) Reciprocal() *Fraction[T] {
	return New(f.D, f.N)
}

func (f *FractionI[T]) Reciprocal() *FractionI[T] {
	return NewI(f.D, f.N)
}

func (f *Fraction[T]) Plus(that *Fraction[T]) *Fraction[T] {
	return &Fraction[T]{f.N*that.D + f.D*that.N, f.D * that.D}
}

func (f *FractionI[T]) Plus(that *FractionI[T]) *FractionI[T] {
	return &FractionI[T]{f.N.Times(that.D).Plus(f.D.Times(that.N)), f.D.Times(that.D)}
}

func (f *Fraction[T]) Code() string {
	return f.String()
}

func (f *FractionI[T]) Code() string {
	return f.String()
}

func (f *Fraction[T]) String() string {
	return fmt.Sprintf("%v/%v", f.N, f.D)
}

func (f *FractionI[T]) String() string {
	return fmt.Sprintf("%v/%v", f.N, f.D)
}

func (f *Fraction[T]) Copy() *Fraction[T] {
	return New(f.N, f.D)
}

func (f *FractionI[T]) Copy() *FractionI[T] {
	return NewI(f.N, f.D)
}

func (f *Fraction[T]) LT(that *Fraction[T]) bool {
	return f.N*that.D < f.D*that.N
}

func (f *FractionI[T]) LT(that *FractionI[T]) bool {
	return f.N.Times(that.D).LT(f.D.Times(that.N))
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
