package p25

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P25() *ecmodels.Problem {
	return ecmodels.IntInputNode(25, func(o command.Output, n int) {
		for g, i := generator.BigFibonaccis(), 1; ; i++ {
			if len(g.Nth(i).String()) >= n {
				o.Stdoutln(i + 1)
				return
			}
		}
	}, []*ecmodels.Execution{
		{
			Args:     []string{"1000"},
			Want:     "4782",
			Estimate: 0.2,
		},
		{
			Args: []string{"2"},
			Want: "7",
		},
		{
			Args: []string{"1"},
			Want: "2",
		},
	})
}
