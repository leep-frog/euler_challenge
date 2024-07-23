package p7

import (
	"github.com/leep-frog/command/command"
	"github.com/leep-frog/euler_challenge/ec/ecmodels"
	"github.com/leep-frog/euler_challenge/generator"
)

func P7() *ecmodels.Problem {
	return ecmodels.IntInputNode(7, func(o command.Output, n int) {
		p := generator.Primes()
		o.Stdoutln(p.Nth(n - 1))
	}, []*ecmodels.Execution{
		{
			Args: []string{"6"},
			Want: "13",
		},
		{
			Args: []string{"9"},
			Want: "23",
		},
		{
			Args: []string{"10001"},
			Want: "104743",
		},
		{
			Args:     []string{"10000001"},
			Want:     "179424691",
			Estimate: 36,
		},
	})
}
