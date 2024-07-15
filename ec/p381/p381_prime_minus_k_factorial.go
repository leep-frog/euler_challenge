package p381

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P381() *ecmodels.Problem {
	return ecmodels.IntInputNode(381, func(o command.Output, n int) {
		o.Stdoutln(solve(maths.Pow(10, n)))
	}, []*ecmodels.Execution{
		{
			Args: []string{"2"},
			Want: "480",
		},
		{
			Args:     []string{"8"},
			Want:     "139602943319822",
			Estimate: 30,
		},
	})
}

func solve(n int) *maths.Int {
	g := generator.Primes()

	// We start at 7 because 5 includes 0! which is weird. The answer for 5 is 4, so just start with that here
	total := maths.NewInt(4)
	for i := 3; g.Nth(i) < n; i++ {
		p := g.Nth(i)

		// Pattern is [X Y (p-1)/2, 1, (p-1)]

		// z * PROD(p-1 to p-5) mod p == p - 1
		// z * (     prod     ) mod p == p - 1

		prod := maths.NewInt(p - 1).TimesInt(p - 2).TimesInt(p - 3).TimesInt(p - 4).TimesInt(p - 5)
		topMod := prod.ModInt(p)
		z := maths.SolveMod(topMod, p, p-1)
		res := prod.TimesInt(z)

		var ns int
		for j := 0; j < 5; j++ {
			ns += res.ModInt(p)
			res = res.DivInt(p - 1 - j)
		}
		total = total.PlusInt(ns % p)
	}
	return total
}
