package p27

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P27() *ecmodels.Problem {
	return ecmodels.IntInputNode(27, func(o command.Output, n int) {
		p := generator.Primes()

		var max, maxI int
		for a := -n + 1; a < n; a++ {
			for b := -n; b <= n; b++ {
				// Try positive direction
				k := 0

				for ; p.Contains(k*k + a*k + b); k++ {
				}
				if k > max {
					max = k
					maxI = a * b
				}

				// Try negative direction
				k = 0
				for ; p.Contains(k*k + a*k + b); k-- {
				}
				if k > max {
					max = k
					maxI = a * b
				}
			}
		}
		o.Stdoutln(maxI)
	}, []*ecmodels.Execution{
		{
			Args: []string{"1000"},
			Want: "-59231",
		},
	})
}
