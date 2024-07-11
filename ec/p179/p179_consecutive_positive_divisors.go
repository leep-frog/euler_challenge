package p179

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P179() *ecmodels.Problem {
	return ecmodels.IntInputNode(179, func(o command.Output, n int) {
		p := generator.Primes()

		var sum int
		prev := 1
		for i := 2; i < 10_000_000; i++ {
			fn := p.FactorCount(i)
			if fn == prev {
				sum++
			}
			prev = fn
		}
		o.Stdoutln(sum)
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1"},
			Want:     "986262",
			Estimate: 7,
		},
	})
}
