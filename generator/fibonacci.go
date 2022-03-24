package generator

import "github.com/leep-frog/euler_challenge/maths"

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
