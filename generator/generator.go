package generator

import "github.com/leep-frog/euler_challenge/maths"

// TODO: cache stuff (every 1000?)
type Generator struct {
	values []*maths.Int
	set    map[string]bool

	f func(*Generator) *maths.Int
}

func (g *Generator) LastBig() *maths.Int {
	return g.values[len(g.values)-1]
}

func (g *Generator) Last() int {
	return g.LastBig().ToInt()
}

func (g *Generator) Len() int {
	return len(g.values)
}

func (g *Generator) BigNth(i int) *maths.Int {
	for len(g.values) <= i {
		g.Next()
	}
	return g.values[i]
}

func (g *Generator) Nth(i int) int {
	return g.BigNth(i).ToInt()
}

func (g *Generator) NextBig() *maths.Int {
	i := g.f(g)
	g.values = append(g.values, i)
	if g.set == nil {
		g.set = map[string]bool{}
	}
	g.set[i.String()] = true
	return i
}

func (g *Generator) Next() int {
	return g.NextBig().ToInt()
}

// Note: this assumes that the cycles are strictly increasing.
func (g *Generator) Contains(i int) bool {
	return g.ContainsBig(maths.NewInt(int64(i)))
}

func (g *Generator) ContainsBig(i *maths.Int) bool {
	for ; g.Len() == 0 || g.LastBig().LTE(i); g.Next() {
	}
	return g.set[i.String()]
}

func NewGenerator(start *maths.Int, f func(*Generator) *maths.Int) *Generator {
	return &Generator{
		f: func(g *Generator) *maths.Int {
			if len(g.values) == 0 {
				return start
			}
			return f(g)
		},
	}
}

func PrimeFactors(n int, p *Generator) map[int]int {
	r := map[int]int{}
	for i := 0; ; i++ {
		pi := p.Nth(i)
		for n%pi == 0 {
			r[pi]++
			n = n / pi
			if n == 1 {
				return r
			}
		}
	}
}

func Primes() *Generator {
	return NewGenerator(maths.NewInt(2), func(g *Generator) *maths.Int {
		for i := g.LastBig().Plus(maths.One()); ; i.PP() {
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

func ShortFibonaccis() *Generator {
	return fibonaccisInt(maths.NewInt(1), maths.NewInt(2))
}

func Fibonaccis() *Generator {
	return fibonaccisInt(maths.NewInt(1), maths.NewInt(1))
}

func fibonaccisInt(a, b *maths.Int) *Generator {
	return NewGenerator(maths.NewInt(1), func(g *Generator) *maths.Int {
		r := b
		b = a.Plus(b)
		a = r
		return a
	})
}

func Triangulars() *Generator {
	i := maths.One()
	return NewGenerator(maths.One(), func(g *Generator) *maths.Int {
		i.PP()
		return g.LastBig().Plus(i)
	})
}
