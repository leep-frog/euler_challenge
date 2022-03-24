package generator

import "github.com/leep-frog/euler_challenge/maths"

func PowerGenerator(power int) *Generator[*maths.Int] {
	return newBigGen(&powerGen{power})
}

type powerGen struct {
	power int
}

func (pg *powerGen) Next(g *Generator[*maths.Int]) *maths.Int {
	return maths.BigPow(len(g.values)+1, pg.power)
}
