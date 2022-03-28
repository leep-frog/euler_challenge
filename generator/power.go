package generator

import "github.com/leep-frog/euler_challenge/maths"

func PowerGenerator(power int) *Generator[*maths.Int] {
	return newBigGen(&powerGen{power})
}

func SmallPowerGenerator(power int) *Generator[int] {
	return newIntGen(&smallPowerGen{power})
}

type powerGen struct {
	power int
}

func (pg *powerGen) Next(g *Generator[*maths.Int]) *maths.Int {
	return maths.BigPow(len(g.values)+1, pg.power)
}

type smallPowerGen struct {
	power int
}

func (spg *smallPowerGen) Next(g *Generator[int]) int {
	return maths.Pow(len(g.values)+1, spg.power)
}
