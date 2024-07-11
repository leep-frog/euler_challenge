package p133

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/ec/p129"
	"github.com/leep-frog/euler_challenge/generator"
)

func P133() *ecmodels.Problem {
	return ecmodels.IntInputNode(133, func(o command.Output, n int) {
		g := generator.Primes()
		var sum int
		for i, pi := 0, g.Nth(0); pi < n; i, pi = i+1, g.Nth(i+1) {
			if !p129.Repunitable(pi) {
				sum += pi
				continue
			}
			factors := g.PrimeFactors(p129.RepunitSmallest(pi))

			for k := range factors {
				if k != 2 && k != 5 {
					sum += pi
					break
				}
			}
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"100000"},
			Want:     "453647705",
			Estimate: 2.5,
		},
		{
			Args: []string{"100"},
			Want: "918",
		},
	})
}
