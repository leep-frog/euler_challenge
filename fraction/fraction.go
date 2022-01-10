package fraction

import (
	"fmt"

	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

type Fraction struct {
	N int
	D int
}

func New(n, d int) *Fraction {
	return &Fraction{n, d}
}

func (f *Fraction) String() string {
	return fmt.Sprintf("%d/%d", f.N, f.D)
}

func (f *Fraction) Copy() *Fraction {
	return New(f.N, f.D)
}

// Return a fraction to allow for chaining.
func (f *Fraction) Simplify(p *generator.Generator[int]) *Fraction {
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
}