package p44

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P44() *ecmodels.Problem {
	return ecmodels.NoInputNode(44, func(o command.Output) {
		best := -1
		pents := generator.Pentagonals()
		for i := 0; ; i++ {
			pi := pents.Nth(i)
			for j := i - 1; j >= 0 && (best == -1 || pi-pents.Nth(j) < best); j-- {
				pj := pents.Nth(j)
				if generator.IsPentagonal(pi-pj) && generator.IsPentagonal(pi+pj) && (best == -1 || pi-pj < best) {
					best = pi - pj
				}
			}
			// We can't do better if the next difference is bigger than the current best.
			if best != -1 && pents.Nth(i+1)-pents.Nth(i) >= best {
				break
			}
		}
		o.Stdoutln(best)
	}, &ecmodels.Execution{
		Want:     "5482660",
		Estimate: 1,
	})
}
