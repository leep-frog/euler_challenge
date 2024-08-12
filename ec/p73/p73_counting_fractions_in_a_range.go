package p73

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P73() *ecmodels.Problem {
	return ecmodels.IntInputNode(73, func(o command.Output, n int) {
		p := generator.Primes()
		var unique int
		for den := 4; den <= n; den++ {
			for num := den / 3; num*2 < den; num++ {
				if num*3 <= den {
					continue
				}
				if !p.Coprimes(num, den) {
					continue
				}
				unique++
			}
		}
		o.Stdoutln(unique)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"12000"},
			Want:     "7295372",
			Estimate: 8,
		},
		{
			Args: []string{"8"},
			Want: "3",
		},
	})
}
