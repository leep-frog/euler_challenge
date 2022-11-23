package fraction

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

type Fraction struct {
	N int
	D int
}

func New(n, d int) *Fraction {
	absN, absD := maths.Abs(n), maths.Abs(d)
	if n*d < 0 {
		absN = -absN
	}
	return &Fraction{absN, absD}
}

func (f *Fraction) ToRational() *Rational {
	return NewRational(f.N, f.D)
}

func (f *Fraction) Div(that *Fraction) *Fraction {
	return f.Times(that.Reciprocal())
}

func (f *Fraction) Times(that *Fraction) *Fraction {
	return New(f.N*that.N, f.D*that.D)
}

func (f *Fraction) Reciprocal() *Fraction {
	return New(f.D, f.N)
}

func (f *Fraction) Plus(that *Fraction) *Fraction {
	return &Fraction{f.N*that.D + f.D*that.N, f.D * that.D}
}

func (f *Fraction) Minus(that *Fraction) *Fraction {
	return f.Plus(that.Negate())
}

func (f *Fraction) Code() string {
	return f.String()
}

func (f *Fraction) String() string {
	return fmt.Sprintf("%v/%v", f.N, f.D)
}

func (f *Fraction) Negate() *Fraction {
	return New(-f.N, f.D)
}

func (f *Fraction) Copy() *Fraction {
	return New(f.N, f.D)
}

func (f *Fraction) LT(that *Fraction) bool {
	return f.N*that.D < f.D*that.N
}

func Simplify(n, d int, primes *generator.Generator[int]) *Fraction {
	return New(n, d).Simplify(primes)
}

// Return a simplified fraction.
func (f *Fraction) Simplify(primes *generator.Generator[int]) *Fraction {
	if f.D == 0 {
		if f.N == 0 {
			return New(0, 0)
		}
		return New(1, 0)
	}
	if f.N == 0 {
		// we know d isn't 0
		return New(0, 1)
	}

	sign := 1
	if f.D < 0 {
		f.D *= -1
		sign *= -1
	}
	if f.N < 0 {
		f.N *= -1
		sign *= -1
	}

	nfs := generator.MutablePrimeFactors(f.N, primes)
	dfs := generator.MutablePrimeFactors(f.D, primes)

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

func CmpOpts() []cmp.Option {
	return []cmp.Option{
		cmp.Comparer(func(a, b *Rational) bool {
			return a.EQ(b)
		}),
		cmp.Comparer(func(a, b *Fraction) bool {
			return !a.LT(b) && !b.LT(a)
		}),
	}
}
