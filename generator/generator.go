package generator

import "github.com/leep-frog/euler_challenge/maths"

// TODO: Use go1.18 parameters to implement this with int and maths.Int
type IntGenerator struct {
	values []*maths.Int

	f func(*IntGenerator) *maths.Int
}

func (ig *IntGenerator) Last() *maths.Int {
	return ig.values[len(ig.values)-1]
}

func (ig *IntGenerator) Len() int {
	return len(ig.values)
}

func (ig *IntGenerator) Nth(i int) *maths.Int {
	for len(ig.values) <= i {
		ig.Next()
	}
	return ig.values[i]
}

func (ig *IntGenerator) Next() *maths.Int {
	i := ig.f(ig)
	ig.values = append(ig.values, i)
	return i
}

type Generator struct {
	values []int

	f func(*Generator) int
}

func (g *Generator) Last() int {
	return g.values[len(g.values)-1]
}

func (g *Generator) Len() int {
	return len(g.values)
}

func (g *Generator) Nth(i int) int {
	for len(g.values) <= i {
		g.Next()
	}
	return g.values[i]
}

func (g *Generator) Next() int {
	i := g.f(g)
	g.values = append(g.values, i)
	return i
}

func NewGenerator(start int, f func(*Generator) int) *Generator {
	return &Generator{
		f: func(g *Generator) int {
			if len(g.values) == 0 {
				return start
			}
			return f(g)
		},
	}
}

func NewIntGenerator(start *maths.Int, f func(*IntGenerator) *maths.Int) *IntGenerator {
	return &IntGenerator{
		f: func(g *IntGenerator) *maths.Int {
			if len(g.values) == 0 {
				return start
			}
			return f(g)
		},
	}
}

func Primes() *Generator {
	return NewGenerator(2, func(g *Generator) int {
		for i := g.Last() + 1; ; i++ {
			newPrime := true
			for _, p := range g.values {
				if i%p == 0 {
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

func PrimesInt() *IntGenerator {
	return NewIntGenerator(maths.NewInt(2), func(g *IntGenerator) *maths.Int {
		for i := g.Last().Plus(maths.One()); ; i.PP() {
			newPrime := true
			for _, p := range g.values {
				if _, rem := i.Div(p); rem.EQ(maths.Zero()) {
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

func Fibonaccis() *Generator {
	a := 1
	b := 2
	return NewGenerator(1, func(g *Generator) int {
		r := b
		b = a + b
		a = r
		return a
	})
}

func FibonaccisInt() *IntGenerator {
	a := maths.NewInt(1)
	b := maths.NewInt(2)
	return NewIntGenerator(maths.NewInt(1), func(ig *IntGenerator) *maths.Int {
		r := b
		b = a.Plus(b)
		a = r
		return a
	})
}

func Triangulars() *Generator {
	i := 1
	return NewGenerator(1, func(g *Generator) int {
		i++
		return g.Last() + i
	})
}

func TriangularsInt() *IntGenerator {
	i := maths.One()
	return NewIntGenerator(maths.One(), func(ig *IntGenerator) *maths.Int {
		i.PP()
		return ig.Last().Plus(i)
	})
}
