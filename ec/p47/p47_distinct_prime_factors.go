package p47

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P47() *ecmodels.Problem {
	return ecmodels.IntInputNode(47, func(o command.Output, n int) {
		p := generator.Primes()
		var row int
		for i := 1; ; i++ {
			if len(p.PrimeFactors(i)) >= n {
				row++
				if row == n {
					o.Stdoutln(i - (n - 1))
					return
				}
			} else {
				row = 0
			}
		}
	}, []*ecmodels.Execution{
		{
			Args:     []string{"4"},
			Want:     "134043",
			Estimate: 0.35,
		},
		{
			Args: []string{"3"},
			Want: "644",
		},
		{
			Args: []string{"2"},
			Want: "14",
		},
	})
}
