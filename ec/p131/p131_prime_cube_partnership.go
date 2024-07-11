package p131

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
	"github.com/leep-frog/euler_challenge/maths"
)

func P131() *ecmodels.Problem {
	return ecmodels.IntInputNode(131, func(o command.Output, max int) {
		g := generator.Primes()
		t := generator.PowerGenerator(3)

		// Noticed pattern that all solutions are of form:
		// n*n*(p + n) such that
		// A: (n*n) is a perfect cube (say c^3)
		// B: (p + n) = (c+1)^3
		// Basically if (c + 1)^3 - c^3 is prime, then it works
		var count int
		bigMax := maths.NewInt(max)
		for i := 0; t.Nth(i + 1).Minus(t.Nth(i)).LT(bigMax); i++ {
			if g.Contains(t.Nth(i + 1).Minus(t.Nth(i)).ToInt()) {
				count++
			}
		}
		o.Stdoutln(count)
		return
	}, []*ecmodels.Execution{
		{
			Args: []string{"1000000"},
			Want: "173",
		},
		{
			Args: []string{"100"},
			Want: "4",
		},
	})
}
